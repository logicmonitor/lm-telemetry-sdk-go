package utils

import (
	"os"
	"strings"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
)

const (
	serviceNameKey        = "service.name"
	serviceNamespaceKey   = "service.namespace"
	resourceAttributesKey = "OTEL_RESOURCE_ATTRIBUTES"

	commaSeperator = ","
	equalSeperator = "="
)

var GetServiceDetails = func() map[string]string {
	resourceAttributes := os.Getenv(resourceAttributesKey)
	if resourceAttributes == "" {
		return nil
	}
	attributeKVPairs := strings.Split(resourceAttributes, commaSeperator)

	attribute := make(map[string]string)

	for _, kv := range attributeKVPairs {
		keyValueList := strings.Split(kv, equalSeperator)
		attribute[keyValueList[0]] = keyValueList[1]
	}

	return attribute
}

var AddEnvResAttributes = func(res *resource.Resource, attributeMap map[string]string) (*resource.Resource, error) {
	attributes := make([]attribute.KeyValue, 0, 1)

	servieName, ok := attributeMap[serviceNameKey]
	if ok {
		serviceNameAttr := attribute.Key(serviceNameKey).String(servieName)
		attributes = append(attributes, serviceNameAttr)
	}

	servieNameSpace, ok := attributeMap[serviceNamespaceKey]
	if ok {
		serviceNameAttr := attribute.Key(serviceNamespaceKey).String(servieNameSpace)
		attributes = append(attributes, serviceNameAttr)
	}

	envRes := resource.NewSchemaless(attributes...)
	return resource.Merge(res, envRes)

}
