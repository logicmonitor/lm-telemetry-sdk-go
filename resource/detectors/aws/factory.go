package aws

import (
	"github.com/logicmonitor/lm-telemetry-sdk-go/resource/detectors/aws/ec2"
	"github.com/logicmonitor/lm-telemetry-sdk-go/resource/detectors/aws/ecs"
	"github.com/logicmonitor/lm-telemetry-sdk-go/resource/detectors/aws/eks"
	"github.com/logicmonitor/lm-telemetry-sdk-go/resource/detectors/aws/lambda"
	"go.opentelemetry.io/otel/sdk/resource"
)

//AWSDetectors is a list of resource detector for AWS
var AWSDetectors []resource.Detector

func init() {
	AWSDetectors = make([]resource.Detector, 0, 1)
	AWSDetectors = append(AWSDetectors,
		eks.NewResourceDetector(),
		ecs.NewResourceDetector(),
		ec2.NewResourceDetector(),
		lambda.NewResourceDetector(),
	)
}
