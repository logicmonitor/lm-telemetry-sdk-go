package gce

import (
	"context"

	"github.com/logicmonitor/lm-telemetry-sdk-go/utils"
	"go.opentelemetry.io/contrib/detectors/gcp"
	"go.opentelemetry.io/otel/sdk/resource"
)

type GCE struct {
	gce resource.Detector
}

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
	return mergedRes, err
}

func NewResourceDetector() resource.Detector {
	return &GCE{
		gce: &gcp.GCE{},
	}
}
