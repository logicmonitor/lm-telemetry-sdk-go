package gke

import (
	"context"

	"go.opentelemetry.io/contrib/detectors/gcp"
	"go.opentelemetry.io/otel/sdk/resource"
)

type GKE struct {
	gke resource.Detector
}

func (computeEngine *GKE) Detect(ctx context.Context) (*resource.Resource, error) {
	res, err := computeEngine.gke.Detect(ctx)
	return res, err
}

func NewResourceDetector() resource.Detector {
	return &GKE{
		gke: &gcp.GKE{},
	}
}
