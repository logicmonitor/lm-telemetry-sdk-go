package main

import (
	"context"
	"log"

	"github.com/logicmonitor/lm-telemetry-sdk-go/config"
	"github.com/logicmonitor/lm-telemetry-sdk-go/telemetry"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

var (
	tracer trace.Tracer
)

func main() {
	ctx := context.Background()

	customAttributes := map[string]string{
		"service.namespace": "sample-namespace",
		"service.name":      "sample-service",
	}

	err := telemetry.SetupTelemetry(ctx,
		config.WithAttributes(customAttributes),
		config.WithHTTPTraceEndpoint("localhost:4318"),
		config.WithSimlpeSpanProcessor(),
	)

	if err != nil {
		log.Fatalf("error in setting up telemetry: %s", err.Error())
	}

	tracer = otel.Tracer("tracer-1")

	ctx, parentSpan := tracer.Start(ctx, "parent span")
	defer parentSpan.End()

	childFunc(ctx)

}

func childFunc(ctx context.Context) {
	_, childSpan := tracer.Start(ctx, "child span")
	defer childSpan.End()
}
