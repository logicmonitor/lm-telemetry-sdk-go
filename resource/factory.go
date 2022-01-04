package resource

import (
	"context"
	"errors"
	"os"
	"strings"

	"github.com/logicmonitor/lm-telemetry-sdk-go/resource/detectors/aws"
	"github.com/logicmonitor/lm-telemetry-sdk-go/resource/detectors/aws/ec2"
	"github.com/logicmonitor/lm-telemetry-sdk-go/resource/detectors/aws/ecs"
	"github.com/logicmonitor/lm-telemetry-sdk-go/resource/detectors/aws/eks"
	"github.com/logicmonitor/lm-telemetry-sdk-go/resource/detectors/aws/lambda"
	"github.com/logicmonitor/lm-telemetry-sdk-go/resource/detectors/gcp"
	"github.com/logicmonitor/lm-telemetry-sdk-go/resource/detectors/gcp/cloudfunction"
	"github.com/logicmonitor/lm-telemetry-sdk-go/resource/detectors/gcp/gce"
	"github.com/logicmonitor/lm-telemetry-sdk-go/resource/detectors/gcp/gke"
	"go.opentelemetry.io/otel/sdk/resource"
)

const (
	resourceAttrKey = "LM_RESOURCE_DETECTOR"
)

//New will auto detect and return a resource
func New(ctx context.Context) (*resource.Resource, error) {
	var res *resource.Resource
	var detector resource.Detector
	var err error

	detectorType := strings.TrimSpace(os.Getenv(resourceAttrKey))

	switch detectorType {
	case AWS_EC2:
		detector = ec2.NewResourceDetector()
	case AWS_ECS:
		detector = ecs.NewResourceDetector()
	case AWS_EKS:
		detector = eks.NewResourceDetector()
	case AWS_LAMBDA:
		detector = lambda.NewResourceDetector()
	case GCP_GCE:
		detector = gce.NewResourceDetector()
	case GCP_GKE:
		detector = gke.NewResourceDetector()
	case GCP_CLOUD_FUNCTIONS:
		detector = cloudfunction.NewResourceDetector()
	}

	if detector != nil {
		return detector.Detect(ctx)
	}

	//AWS
	for _, awsDetector := range aws.AWSDetectors {
		res, err = awsDetector.Detect(ctx)
		if res != nil && err == nil {
			return res, nil
		}
	}

	//GCP
	for _, gcpDetector := range gcp.GCPDetectors {
		res, err = gcpDetector.Detect(ctx)
		if res != nil && err == nil {
			return res, nil
		}
	}

	return nil, errors.New("resouce cannot be detected")
}
