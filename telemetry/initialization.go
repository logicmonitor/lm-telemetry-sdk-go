package telemetry

import (
	"context"

	"github.com/logicmonitor/lm-telemetry-sdk-go/config"
	"github.com/logicmonitor/lm-telemetry-sdk-go/exporter/otlphttpexporter"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

/*SetupTelemetry initializes opentelemetry configurations,
and takes in context and config.Option(s) as input params.
Configuration options could be found in config package. */
func SetupTelemetry(ctx context.Context, opts ...config.Option) error {
	c := config.NewConfig()

	for _, opt := range opts {
		opt(c)
	}

	res, err := c.Detector.Detect(ctx)
	if err != nil {
		return err
	}

	if c.UserResourceAttributes != nil {
		attributes := make([]attribute.KeyValue, 0, 1)
		for key, value := range c.UserResourceAttributes {
			attributes = append(attributes, attribute.String(key, value))
		}
		attrRes := resource.NewSchemaless(attributes...)
		res, err = resource.Merge(attrRes, res)
		if err != nil {
			return err
		}
	}

	var traceExporter sdktrace.SpanExporter
	if c.InAppCollector != nil {
		traceExporter, err = otlphttpexporter.NewOtlpHttpExporter(c.InAppCollector.TraceEndpoint, c.InAppCollector.Headers)
	} else {
		traceExporter, err = otlptracehttp.New(ctx,
			otlptracehttp.WithInsecure(),
			otlptracehttp.WithEndpoint(c.TraceEndpoint),
		)
	}
	if err != nil {
		return err
	}

	bsp := sdktrace.NewBatchSpanProcessor(traceExporter)
	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(res),
		sdktrace.WithSpanProcessor(bsp),
	)
	otel.SetTracerProvider(tracerProvider)
	otel.SetTextMapPropagator(propagation.TraceContext{})
	return nil
}
