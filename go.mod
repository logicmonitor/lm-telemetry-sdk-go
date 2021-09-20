module github.com/logicmonitor/lm-telemetry-sdk-go

go 1.16

require (
	cloud.google.com/go v0.88.0
	github.com/aws/aws-sdk-go v1.40.41
	github.com/stretchr/testify v1.7.0
	github.com/wadey/gocovmerge v0.0.0-20160331181800-b5bfa59ec0ad // indirect
	go.opentelemetry.io/contrib/detectors/aws v0.22.0
	go.opentelemetry.io/contrib/detectors/aws/ecs v0.22.0
	go.opentelemetry.io/contrib/detectors/gcp v0.22.0
	go.opentelemetry.io/otel v1.0.0-RC3
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp v1.0.0-RC3 // indirect
	go.opentelemetry.io/otel/sdk v1.0.0-RC3
	go.opentelemetry.io/otel/trace v1.0.0-RC3 // indirect
)
