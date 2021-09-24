package gke

import (
	"context"

	"go.opentelemetry.io/contrib/detectors/gcp"
	"go.opentelemetry.io/otel/sdk/resource"
)

// GKE implements resource.Detector interface for GKE
type GKE struct {
	gke resource.Detector
}

// Detect detects associated resources when running in GKE environment.
func (computeEngine *GKE) Detect(ctx context.Context) (*resource.Resource, error) {
	res, err := computeEngine.gke.Detect(ctx)
	return res, err
}

// NewResourceDetector will return an implementation for gcp gce resource detector
func NewResourceDetector() resource.Detector {
	return &GKE{
		gke: &gcp.GKE{},
	}
}
