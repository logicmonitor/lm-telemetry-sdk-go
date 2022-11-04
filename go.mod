module github.com/logicmonitor/lm-telemetry-sdk-go

go 1.16

require (
	cloud.google.com/go/compute v1.10.0
	github.com/aws/aws-lambda-go v1.27.1
	github.com/aws/aws-sdk-go v1.44.114
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.11.3 // indirect
	github.com/stretchr/testify v1.7.1
	go.opentelemetry.io/contrib/detectors/aws v0.22.0
	go.opentelemetry.io/contrib/detectors/aws/ecs v0.22.0
	go.opentelemetry.io/contrib/detectors/gcp v0.22.0
	go.opentelemetry.io/otel v1.11.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.11.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp v1.11.0
	go.opentelemetry.io/otel/sdk v1.11.0
	go.opentelemetry.io/otel/trace v1.11.0
	go.opentelemetry.io/proto/otlp v0.19.0
	golang.org/x/net v0.0.0-20221012135044-0b7e1fb9d458 // indirect
	golang.org/x/sys v0.0.0-20221010170243-090e33056c14 // indirect
	golang.org/x/text v0.3.8 // indirect
	google.golang.org/genproto v0.0.0-20221010155953-15ba04fc1c0e // indirect
	google.golang.org/grpc v1.50.0
	google.golang.org/protobuf v1.28.1
)
