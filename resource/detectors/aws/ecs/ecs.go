package ecs

import (
	"context"

	otelcontribecs "go.opentelemetry.io/contrib/detectors/aws/ecs"
	"go.opentelemetry.io/otel/sdk/resource"
)

//NewResourceDetector will return an implementation for aws ecs resource detector
func NewResourceDetector() resource.Detector {
	return &ECS{
		otelECSDetector: otelcontribecs.NewResourceDetector(),
	}
}

//ECS implements resource.Detector interface for aws ecs
type ECS struct {
	otelECSDetector resource.Detector
}

/*
Detect will return a resource instance which will have attributes describing,
an ecs
*/
func (ecs *ECS) Detect(ctx context.Context) (*resource.Resource, error) {
	res, err := ecs.otelECSDetector.Detect(ctx)
	return res, err
}
