package gke

import (
	"context"
	"testing"

	"github.com/logicmonitor/lm-telemetry-sdk-go/mock"
	"go.opentelemetry.io/otel/sdk/resource"
)

func TestDetect(t *testing.T) {
	gkeMock := mock.DetectorMock{
		Res: resource.Empty(),
		Err: nil,
	}

	ecs := GKE{
		gke: gkeMock,
	}

	t.Run("Success case", func(t *testing.T) {
		_, err := ecs.Detect(context.Background())
		if err != gkeMock.Err {
			t.Fatal("unexpected error:", err)
		}
	})
}
