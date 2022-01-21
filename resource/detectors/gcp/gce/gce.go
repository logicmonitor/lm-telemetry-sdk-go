package gce

import (
	"context"

	"github.com/logicmonitor/lm-telemetry-sdk-go/utils"
	"go.opentelemetry.io/contrib/detectors/gcp"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

// GCE implements resource.Detector interface for GCE instances
type GCE struct {
	gce resource.Detector
}

// Detect detects associated resources when running on GCE hosts.
func (computeEngine *GCE) Detect(ctx context.Context) (*resource.Resource, error) {
	res, err := computeEngine.gce.Detect(ctx)
	if err != nil {
		return res, err
	}
	envAttributes := utils.GetServiceDetails()
	mergedRes, utilserr := utils.AddEnvResAttributes(res, envAttributes)
	if utilserr != nil {
		return res, err
	}

	attributes := make([]attribute.KeyValue, 0, 1)
	attributes = append(attributes, semconv.CloudPlatformGCPComputeEngine)
	attrRes := resource.NewSchemaless(attributes...)

	return resource.Merge(attrRes, mergedRes)
}

//NewResourceDetector will return an implementation for gcp gce resource detector
func NewResourceDetector() resource.Detector {
	return &GCE{
		gce: &gcp.GCE{},
	}
}
