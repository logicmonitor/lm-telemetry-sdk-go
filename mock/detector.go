package mock

import (
	"context"

	"go.opentelemetry.io/otel/sdk/resource"
)

//DetectorMock mocks resource.Detector implementation
type DetectorMock struct {
	Res *resource.Resource
	Err error
}

//Detect detects a mocked resource
func (detector DetectorMock) Detect(ctx context.Context) (*resource.Resource, error) {
	return detector.Res, detector.Err
}
