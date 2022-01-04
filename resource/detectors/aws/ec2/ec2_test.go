package ec2

import (
	"errors"
	"testing"

	"github.com/logicmonitor/lm-telemetry-sdk-go/mock"
	"go.opentelemetry.io/otel/sdk/resource"
)

var (
	localIP = "127.0.0.1"
	errTest = errors.New("test error")
)

func TestDetect(t *testing.T) {
	type input struct {
		client Client
	}
	type want struct {
		err error
		res *resource.Resource
	}

	testTable := map[string]struct {
		Input input
		Want  want
	}{
		"client unavailable": {
			Input: input{
				client: &mock.MockClient{
					AvailableValue: false,
				},
			},
		},
	}
}
