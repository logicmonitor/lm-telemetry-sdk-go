package gce

import (
	"context"
	"errors"
	"testing"

	"github.com/logicmonitor/lm-telemetry-sdk-go/mock"
	"github.com/logicmonitor/lm-telemetry-sdk-go/utils"
	"go.opentelemetry.io/otel/sdk/resource"
)

var (
	testError = errors.New("test error")
)

func TestDetect(t *testing.T) {
	oldutilsGetServiceDetails := utils.GetServiceDetails
	oldAddEnvResAttributes := utils.AddEnvResAttributes
	defer func() {
		utils.GetServiceDetails = oldutilsGetServiceDetails
		utils.AddEnvResAttributes = oldAddEnvResAttributes
	}()
	utils.GetServiceDetails = mock.CreateGetServiceDetailsMock(map[string]string{"attrib1": "value1"})
	utils.AddEnvResAttributes = mock.CreateAddEnvResAttributesMock(resource.Empty(), nil)

	gceMock := mock.DetectorMock{
		Res: resource.Empty(),
		Err: nil,
	}

	gce := GCE{
		gce: gceMock,
	}

	t.Run("Success case", func(t *testing.T) {
		_, err := gce.Detect(context.Background())
		if err != gceMock.Err {
			t.Fatalf("unexpected error:%v", err)
		}
	})

	t.Run("Error in adding environment resource attribute", func(t *testing.T) {
		oldAddEnvResAttributes := utils.AddEnvResAttributes
		defer func() {
			utils.AddEnvResAttributes = oldAddEnvResAttributes
		}()
		utils.AddEnvResAttributes = mock.CreateAddEnvResAttributesMock(resource.Empty(), testError)

		_, err := gce.Detect(context.Background())
		if err != gceMock.Err {
			t.Fatalf("unexpected error:%v", err)
		}
	})

	t.Run("Error in resource detection", func(t *testing.T) {
		gceMock := mock.DetectorMock{
			Res: resource.Empty(),
			Err: testError,
		}

		gce := GCE{
			gce: gceMock,
		}

		_, err := gce.Detect(context.Background())
		if err != gceMock.Err {
			t.Fatalf("unexpected error:%v", err)
		}
	})
}
