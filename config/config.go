package config

import (
	"context"

	lmresource "github.com/logicmonitor/lm-telemetry-sdk-go/resource"
	"go.opentelemetry.io/otel/sdk/resource"
)

type Config struct {
	UserResourceAttributes map[string]string
	Detector               resource.Detector
	TraceEndpoint          string
}

type Option func(*Config)

func NewConfig() *Config {
	return &Config{
		Detector: &DefaultDetector{},
	}
}

type DefaultDetector struct {
}

func (dd *DefaultDetector) Detect(ctx context.Context) (*resource.Resource, error) {
	return lmresource.New(ctx), nil
}
