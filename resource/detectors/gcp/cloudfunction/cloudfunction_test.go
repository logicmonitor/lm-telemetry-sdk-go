package cloudfunction

import (
	"context"
	"errors"
	"os"
	"testing"
)

var (
	functionName = "function-1"
	errTest      = errors.New("test error")
)

type gcpMock struct {
	ProjectID    string
	ProjectIDErr error
	Region       string
	RegionErr    error
}

func (gi gcpMock) gcpProjectID() (string, error) {
	return gi.ProjectID, gi.ProjectIDErr
}

func (gi gcpMock) gcpRegion() (string, error) {
	return gi.Region, gi.RegionErr
}

func TestDetect(t *testing.T) {
	err := os.Setenv(gcpFunctionNameKey, functionName)
	if err != nil {
		t.Fatal("error in setting function name :", err)
	}
	defer func() {
		err := os.Unsetenv(gcpFunctionNameKey)
		if err != nil {
			t.Fatal("error in clearing env var :", err)
		}
	}()

	t.Run("Success Case", func(t *testing.T) {
		detector := Function{
			client: gcpMock{
				ProjectID:    "someID",
				ProjectIDErr: nil,
				Region:       "some-region",
				RegionErr:    nil,
			},
		}

		_, err := detector.Detect(context.Background())
		if err != nil {
			t.Fatal("unexpected error :", err)
		}
	})

	t.Run("error in getting projectID", func(t *testing.T) {
		mockClient := gcpMock{
			ProjectID:    "someID",
			ProjectIDErr: errTest,
		}
		detector := Function{
			client: mockClient,
		}

		_, err := detector.Detect(context.Background())
		if err != mockClient.ProjectIDErr {
			t.Fatal("unexpected error :", err)
		}
	})

	t.Run("error in getting region", func(t *testing.T) {
		mockClient := gcpMock{
			Region:    "",
			RegionErr: errTest,
		}
		detector := Function{
			client: mockClient,
		}

		_, err := detector.Detect(context.Background())
		if err != mockClient.RegionErr {
			t.Fatal("unexpected error :", err)
		}
	})

	t.Run("functionName not set", func(t *testing.T) {
		oldKey := os.Getenv(gcpFunctionNameKey)
		defer func() {
			if oldKey != "" {
				os.Setenv(gcpFunctionNameKey, oldKey)
			}
		}()
		os.Unsetenv(gcpFunctionNameKey)

		detector := Function{}
		_, err := detector.Detect(context.Background())
		if err != errNotOnGoogleCloudFunction {
			t.Fatal("unexpected error :", err)
		}
	})

}
