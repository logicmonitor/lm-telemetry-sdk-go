package config

import (
	"context"

	lmresource "github.com/logicmonitor/lm-telemetry-sdk-go/resource"
	"go.opentelemetry.io/otel/sdk/resource"
)

//Config represents opentelemetry configurations
type Config struct {
	UserResourceAttributes map[string]string
	Detector               resource.Detector
	TraceEndpoint          string
}

type Option func(*Config)

// NewConfig returns new instance of config
func NewConfig() *Config {
	return &Config{
		Detector: &defaultDetector{},
	}
}

type defaultDetector struct {
}

func (dd *defaultDetector) Detect(ctx context.Context) (*resource.Resource, error) {
	return lmresource.New(ctx), nil
}
