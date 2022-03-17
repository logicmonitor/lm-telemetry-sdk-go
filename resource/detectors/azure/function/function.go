package function

import (
	"context"
	"errors"
	"os"
	"strings"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
)

const (
	azureFunctionWorkingDirectory = "FUNCTIONS_WORKER_DIRECTORY"
	azureFunctions                = "azure-functions"
	websiteDeploymentID           = "WEBSITE_DEPLOYMENT_ID"
)

var (
	errNotOnAzureFunction = errors.New("process is not on Azure-Function, cannot detect environment variables from Azure-Function")
	errFaasIDNotFound     = errors.New("faas.Id not found")
)

//NewResourceDetector will return an implementation for aws ec2 resource detector
func NewResourceDetector() resource.Detector {
	return &AzureFunction{}
}

type AzureFunction struct {
}

func (afunc *AzureFunction) Detect(ctx context.Context) (*resource.Resource, error) {
	if !isAzureFunction() {
		return resource.Empty(), errNotOnAzureFunction
	}

	faasID := getFaasID()

	attributes := []attribute.KeyValue{
		semconv.CloudProviderAzure,
		semconv.CloudPlatformAzureFunctions,
	}

	if faasID == "" {
		return resource.Empty(), errFaasIDNotFound
	}

	attributes = append(attributes, attribute.String(string(semconv.FaaSIDKey), faasID))
	return resource.NewSchemaless(attributes...), nil
}

func isAzureFunction() bool {
	workingDir, ok := os.LookupEnv(azureFunctionWorkingDirectory)
	if !ok {
		return false
	}

	if strings.Contains(workingDir, azureFunctions) {
		return true
	}
	return false
}

func getFaasID() string {
	return os.Getenv(websiteDeploymentID)
}
