package config

import (
	"context"

	lmresource "github.com/logicmonitor/lm-telemetry-sdk-go/resource"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

//Config represents opentelemetry configurations
type Config struct {
	UserResourceAttributes   map[string]string
	Detector                 resource.Detector
	TraceEndpoint            string
	InAppExporter            *sdkTraceExporter
	SpanProcessor            func(sdktrace.SpanExporter) sdktrace.SpanProcessor
	HTTPOption               []otlptracehttp.Option
	GRPCOption               []otlptracegrpc.Option
	IsGRPCExporterConfigured bool
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
		Detector: &defaultDetector{},
	}
}

type defaultDetector struct {
}

func (dd *defaultDetector) Detect(ctx context.Context) (*resource.Resource, error) {
	return lmresource.New(ctx)
}
