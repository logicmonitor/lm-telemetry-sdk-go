package cloudfunction

import (
	"context"
	"errors"
	"os"

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

func NewResourceDetector() resource.Detector {
	return &Function{
		client: gcpImpl{},
	}
}

type gcpClient interface {
	gcpProjectID() (string, error)
}

type gcpImpl struct{}

func (gi gcpImpl) gcpProjectID() (string, error) {
	return metadata.ProjectID()
}

type Function struct {
	client gcpClient
}

func (f *Function) Detect(ctx context.Context) (*resource.Resource, error) {

	functionName, ok := f.googleCloudFunctionName()
	if !ok {
		return nil, errNotOnGoogleCloudFunction
	}

	projectID, err := f.client.gcpProjectID()
	if err != nil {
		return nil, err
	}

	attributes := []attribute.KeyValue{
		semconv.CloudProviderGCP,
		attribute.String(string(semconv.FaaSNameKey), functionName),
		attribute.String("gcp.projectID", projectID),
	}
	return resource.NewSchemaless(attributes...), nil
}

func (f *Function) googleCloudFunctionName() (string, bool) {
	return os.LookupEnv(gcpFunctionNameKey)
}
