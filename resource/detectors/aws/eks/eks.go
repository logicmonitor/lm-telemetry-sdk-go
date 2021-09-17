package eks

import (
	"context"

	otelcontribeks "go.opentelemetry.io/contrib/detectors/aws/ecs"
	"go.opentelemetry.io/otel/sdk/resource"
)

type EKS struct {
	otelEKSDetector resource.Detector
}

func (eks *EKS) Detect(ctx context.Context) (*resource.Resource, error) {
	res, err := eks.otelEKSDetector.Detect(ctx)
	return res, err
}

func NewResourceDetector() resource.Detector {
	return &EKS{
		otelEKSDetector: otelcontribeks.NewResourceDetector(),
	}
}
