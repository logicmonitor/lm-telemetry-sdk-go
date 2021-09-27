package config

import "github.com/logicmonitor/lm-telemetry-sdk-go/resource/detectors/aws/ec2"

/* WithAttributes returns a config option which adds
custom attributes to resource */
func WithAttributes(attributes map[string]string) Option {
	return func(c *Config) {
		c.UserResourceAttributes = attributes
	}
}

/* WithAWSEC2Detector returns a config option which sets
config detector to ec2 detector
*/
func WithAWSEC2Detector() Option {
	return func(c *Config) {
		c.Detector = ec2.NewResourceDetector()
	}
}

/* WithHttpTraceEndpoint returns a config option which sets
trace endpoint
*/
func WithHttpTraceEndpoint(endpoint string) Option {
	return func(c *Config) {
		c.TraceEndpoint = endpoint
	}
}
