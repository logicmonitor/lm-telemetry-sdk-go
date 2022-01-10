package config

import (
	"github.com/logicmonitor/lm-telemetry-sdk-go/resource/detectors/aws/ec2"
	"github.com/logicmonitor/lm-telemetry-sdk-go/resource/detectors/aws/lambda"
	"github.com/logicmonitor/lm-telemetry-sdk-go/resource/detectors/gcp/cloudfunction"
)

/*WithAttributes returns a config option which adds
custom attributes to resource */
func WithAttributes(attributes map[string]string) Option {
	return func(c *Config) {
		c.UserResourceAttributes = attributes
	}
}

/*WithAWSEC2Detector returns a config option which sets
config detector to ec2 detector
*/
func WithAWSEC2Detector() Option {
	return func(c *Config) {
		c.Detector = ec2.NewResourceDetector()
	}
}

/*WithAWSLambdaDetector returns a config option which sets
config detector to ec2 detector
*/
func WithAWSLambdaDetector() Option {
	return func(c *Config) {
		c.Detector = lambda.NewResourceDetector()
	}
}

func WithGCPcloudFunctionDetector() Option {
	return func(c *Config) {
		c.Detector = cloudfunction.NewResourceDetector()
	}
}
