package cloudfunction

import (
	"context"
	"errors"
	"os"
	"strings"

	"cloud.google.com/go/compute/metadata"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

const (
	gcpFunctionNameKey = "K_SERVICE"
)

var (
	errNotOnGoogleCloudFunction = errors.New("cannot detect environment variables from Google Cloud Function")
)

//NewResourceDetector will return an implementation for gcp cloud function resource detector
func NewResourceDetector() resource.Detector {
	return &Function{
		client: gcpImpl{},
	}
}

type gcpClient interface {
	gcpProjectID() (string, error)
	gcpRegion() (string, error)
}

type gcpImpl struct{}

func (gi gcpImpl) gcpProjectID() (string, error) {
	return metadata.ProjectID()
}

func (gi gcpImpl) gcpRegion() (string, error) {
	var region string
	zone, err := metadata.Zone()
	if zone != "" {
		splitArr := strings.SplitN(zone, "-", 3)
		if len(splitArr) == 3 {
			region = strings.Join(splitArr[0:2], "-")
		}
	}
	return region, err
}

// Function implements resource.Detector interface for google cloud-function
type Function struct {
	client gcpClient
}

// Detect detects associated resources when running in  cloud function.
func (f *Function) Detect(ctx context.Context) (*resource.Resource, error) {

	functionName, ok := f.googleCloudFunctionName()
	if !ok {
		return nil, errNotOnGoogleCloudFunction
	}

	projectID, err := f.client.gcpProjectID()
	if err != nil {
		return nil, err
	}

	region, err := f.client.gcpRegion()
	if err != nil {
		return nil, err
	}

	attributes := []attribute.KeyValue{
		semconv.CloudProviderGCP,
		semconv.CloudPlatformGCPCloudFunctions,
		attribute.String(string(semconv.FaaSNameKey), functionName),
		semconv.CloudAccountIDKey.String(projectID),
		semconv.CloudRegionKey.String(region),
	}
	return resource.NewSchemaless(attributes...), nil
}

func (f *Function) googleCloudFunctionName() (string, bool) {
	return os.LookupEnv(gcpFunctionNameKey)
}
