package mock

import (
	"context"

	"go.opentelemetry.io/otel/sdk/resource"
)

type DetectorMock struct {
	Res *resource.Resource
	Err error
}

func (detector DetectorMock) Detect(ctx context.Context) (*resource.Resource, error) {
	return detector.Res, detector.Err
}
