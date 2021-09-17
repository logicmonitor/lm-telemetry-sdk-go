package ecs

import (
	"context"
	"testing"

	"github.com/logicmonitor/lm-telemetry-sdk-go/mock"
	"go.opentelemetry.io/otel/sdk/resource"
)

func TestDetect(t *testing.T) {
	ecsMock := mock.DetectorMock{
		Res: resource.Empty(),
		Err: nil,
	}

	ecs := ECS{
		otelECSDetector: ecsMock,
	}

	t.Run("Success case", func(t *testing.T) {
		_, err := ecs.Detect(context.Background())
		if err != ecsMock.Err {
			t.Fatal("unexpected error:", err)
		}
	})
}
