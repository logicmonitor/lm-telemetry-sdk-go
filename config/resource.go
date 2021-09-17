package config

import "github.com/logicmonitor/lm-telemetry-sdk-go/resource/detectors/aws/ec2"

func WithAttributes(attributes map[string]string) Option {
	return func(c *Config) {
		c.UserResourceAttributes = attributes
	}
}

func WithAWSEC2Detector() Option {
	return func(c *Config) {
		c.Detector = ec2.NewResourceDetector()
	}
}

func WithHttpTraceEndpoint(endpoint string) Option {
	return func(c *Config) {
		c.TraceEndpoint = endpoint
	}
}
