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
		"key1": "value1",
	}

	err := telemetry.SetupTelemetry(ctx,
		config.WithAWSEC2Detector(),
		config.WithAttributes(customAttributes),
		config.WithHttpTraceEndpoint("localhost:55681"),
	)
	if err != nil {
		log.Fatalf("error in setting up telemetry: %s", err.Error())
	}

	tracer := otel.Tracer("tracer-1")

	ctx, parentSpan := tracer.Start(ctx, "parent span")
	defer parentSpan.End()
}

func childFunc(ctx context.Context) {
	_, childSpan := tracer.Start(ctx, "child span")
	defer childSpan.End()
}
