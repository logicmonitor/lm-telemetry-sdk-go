package vm

import (
	"context"

	"github.com/logicmonitor/lm-telemetry-sdk-go/utils"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

//NewResourceDetector will return an implementation for aws ec2 resource detector
func NewResourceDetector() resource.Detector {
	return &VM{
		provider: NewProvider(),
	}
}

type VM struct {
	provider Provider
}

func (v *VM) Detect(ctx context.Context) (*resource.Resource, error) {
	compute, err := v.provider.Metadata(ctx)
	if err != nil {
		return nil, err
	}
	attributes := []attribute.KeyValue{
		semconv.CloudProviderAzure,
		semconv.CloudPlatformAzureVM,
		semconv.HostIDKey.String(compute.VMID),
		semconv.CloudRegionKey.String(compute.Location),
		semconv.CloudAccountIDKey.String(compute.SubscriptionID),
		semconv.NetHostNameKey.String(compute.Name),
	}

	res := resource.NewSchemaless(attributes...)
	envAttributes := utils.GetServiceDetails()
	mergedRes, err := utils.AddEnvResAttributes(res, envAttributes)
	if err != nil {
		return res, err
	}

	return mergedRes, nil
}
