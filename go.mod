module github.com/logicmonitor/lm-telemetry-sdk-go

go 1.16

require (
	cloud.google.com/go v0.88.0
	github.com/aws/aws-lambda-go v1.27.1
	github.com/aws/aws-sdk-go v1.40.41
	github.com/stretchr/testify v1.7.0
	go.opentelemetry.io/contrib/detectors/aws v0.22.0
	go.opentelemetry.io/contrib/detectors/aws/ecs v0.22.0
	go.opentelemetry.io/contrib/detectors/gcp v0.22.0
	go.opentelemetry.io/otel v1.5.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp v1.0.0-RC3
	go.opentelemetry.io/otel/sdk v1.0.0-RC3
	go.opentelemetry.io/otel/trace v1.5.0
	go.opentelemetry.io/proto/otlp v0.9.0
	google.golang.org/protobuf v1.27.1
)
