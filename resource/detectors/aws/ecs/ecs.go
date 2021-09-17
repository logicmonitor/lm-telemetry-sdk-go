package ecs

import (
	"context"

	otelcontribecs "go.opentelemetry.io/contrib/detectors/aws/ecs"
	"go.opentelemetry.io/otel/sdk/resource"
)

func NewResourceDetector() resource.Detector {
	return &ECS{
		otelECSDetector: otelcontribecs.NewResourceDetector(),
	}
}

type ECS struct {
	otelECSDetector resource.Detector
}

func (ecs *ECS) Detect(ctx context.Context) (*resource.Resource, error) {
	res, err := ecs.otelECSDetector.Detect(ctx)
	return res, err
}
