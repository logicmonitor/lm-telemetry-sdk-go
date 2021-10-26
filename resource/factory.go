package resource

import (
	"context"

	"github.com/logicmonitor/lm-telemetry-sdk-go/resource/detectors/aws"
	"github.com/logicmonitor/lm-telemetry-sdk-go/resource/detectors/gcp"
	"go.opentelemetry.io/otel/sdk/resource"
)

//New will auto detect and return a resource
func New(ctx context.Context) *resource.Resource {
	var res *resource.Resource
	var err error
	//AWS
	for _, awsDetector := range aws.AWSDetectors {
		res, err = awsDetector.Detect(ctx)
		if res != nil && err == nil {
			return res
		}
	}

	//GCP
	for _, gcpDetector := range gcp.GCPDetectors {
		res, err = gcpDetector.Detect(ctx)
		if res != nil && err == nil {
			return res
		}
	}

	return nil
}
