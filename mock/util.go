package mock

import "go.opentelemetry.io/otel/sdk/resource"

func CreateGetServiceDetailsMock(attributes map[string]string) func() map[string]string {
	return func() map[string]string {
		return attributes
	}
}

func CreateAddEnvResAttributesMock(retRes *resource.Resource, err error) func(res *resource.Resource, attributeMap map[string]string) (*resource.Resource, error) {
	return func(res *resource.Resource, attributeMap map[string]string) (*resource.Resource, error) {
		return retRes, err
	}
}
