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
	ProjectID string
	Err       error
}

func (gi gcpMock) gcpProjectID() (string, error) {
	return gi.ProjectID, gi.Err
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
				ProjectID: "someID",
				Err:       nil,
			},
		}

		_, err := detector.Detect(context.Background())
		if err != nil {
			t.Fatal("unexpected error :", err)
		}
	})

	t.Run("error in getting projectID", func(t *testing.T) {
		mockClient := gcpMock{
			ProjectID: "someID",
			Err:       errTest,
		}
		detector := Function{
			client: mockClient,
		}

		_, err := detector.Detect(context.Background())
		if err != mockClient.Err {
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
