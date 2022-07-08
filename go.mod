module github.com/logicmonitor/lm-telemetry-sdk-go

go 1.16

require (
	cloud.google.com/go v0.88.0
	github.com/aws/aws-lambda-go v1.27.1
	github.com/aws/aws-sdk-go v1.40.41
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.10.3 // indirect
	github.com/stretchr/testify v1.7.1
	go.opentelemetry.io/contrib/detectors/aws v0.22.0
	go.opentelemetry.io/contrib/detectors/aws/ecs v0.22.0
	go.opentelemetry.io/contrib/detectors/gcp v0.22.0
	go.opentelemetry.io/otel v1.7.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.7.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp v1.7.0
	go.opentelemetry.io/otel/sdk v1.7.0
	go.opentelemetry.io/otel/trace v1.7.0
	go.opentelemetry.io/proto/otlp v0.18.0
	golang.org/x/net v0.0.0-20220706163947-c90051bbdb60 // indirect
	golang.org/x/sys v0.0.0-20220704084225-05e143d24a9e // indirect
	google.golang.org/genproto v0.0.0-20220707150051-590a5ac7bee1 // indirect
	google.golang.org/grpc v1.47.0
	google.golang.org/protobuf v1.28.0
)
