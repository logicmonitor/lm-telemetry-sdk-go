package config

import (
	"context"

	lmresource "github.com/logicmonitor/lm-telemetry-sdk-go/resource"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

//Config represents opentelemetry configurations
type Config struct {
	UserResourceAttributes   map[string]string
	Detector                 resource.Detector
	TraceEndpoint            string
	InAppExporter            *sdkTraceExporter
	SpanProcessor            func(sdktrace.SpanExporter) sdktrace.SpanProcessor
	IsGRPCExporterConfigured bool
	Credential               credentials.TransportCredentials
	SecureHTTP               bool
}

type sdkTraceExporter struct {
	TraceEndpoint string
	Headers       map[string]string
}

//Option option for configuring telemetry setup
type Option func(*Config)

// NewConfig returns new instance of config
func NewConfig() *Config {
	return &Config{
		Detector:   &defaultDetector{},
		Credential: insecure.NewCredentials(),
		SecureHTTP: true,
	}
}

type defaultDetector struct {
}

func (dd *defaultDetector) Detect(ctx context.Context) (*resource.Resource, error) {
	return lmresource.New(ctx)
}
