package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/otel/sdk/resource"
)

var (
	testVariables = "service.namespace=sample-namespace,service.name=sample-service"
	attribMap     = map[string]string{
		"service.namespace": "sample-namespace",
		"service.name":      "sample-service",
	}
)

func TestGetServiceDetails(t *testing.T) {
	os.Setenv(resourceAttributesKey, testVariables)
	defer func() {
		os.Unsetenv(resourceAttributesKey)
	}()

	t.Run("Success case", func(t *testing.T) {
		attributes := GetServiceDetails()
		assert.Equal(t, attributes, attribMap)
	})

	t.Run("no env var set", func(t *testing.T) {
		os.Unsetenv(resourceAttributesKey)
		attributes := GetServiceDetails()
		assert.Nil(t, attributes)
	})
}

func TestAddEnvResAttributes(t *testing.T) {
	os.Setenv(resourceAttributesKey, testVariables)
	defer func() {
		os.Unsetenv(resourceAttributesKey)
	}()

	t.Run("Success case", func(t *testing.T) {
		attributes := GetServiceDetails()
		_, err := AddEnvResAttributes(resource.Empty(), attributes)
		if err != nil {
			t.Fatal("unexpected error :", err)
		}
	})
}
