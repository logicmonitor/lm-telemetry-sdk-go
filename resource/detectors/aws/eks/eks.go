package eks

import (
	"context"

	otelcontribeks "go.opentelemetry.io/contrib/detectors/aws/ecs"
	"go.opentelemetry.io/otel/sdk/resource"
)

//EKS implements resource.Detector interface for aws eks
type EKS struct {
	otelEKSDetector resource.Detector
}

/*
Detect will return a resource instance which will have attributes describing,
an eks
*/
func (eks *EKS) Detect(ctx context.Context) (*resource.Resource, error) {
	res, err := eks.otelEKSDetector.Detect(ctx)
	return res, err
}

//NewResourceDetector will return an implementation for aws eks resource detector
func NewResourceDetector() resource.Detector {
	return &EKS{
		otelEKSDetector: otelcontribeks.NewResourceDetector(),
	}
}
