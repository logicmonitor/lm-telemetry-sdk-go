package eks

import (
	"context"
	"testing"

	"github.com/logicmonitor/lm-telemetry-sdk-go/mock"
	"go.opentelemetry.io/otel/sdk/resource"
)

func TestDetect(t *testing.T) {
	eksMock := mock.DetectorMock{
		Res: resource.Empty(),
		Err: nil,
	}

	eks := EKS{
		otelEKSDetector: eksMock,
	}

	t.Run("Success case", func(t *testing.T) {
		_, err := eks.Detect(context.Background())
		if err != eksMock.Err {
			t.Fatal("unexpected error:", err)
		}
	})
}
