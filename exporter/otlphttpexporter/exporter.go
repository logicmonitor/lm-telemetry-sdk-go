package otlphttpexporter

import (
	"bytes"
	"context"
	"net/http"
	"net/url"

	//"go.opentelemetry.io/otel/exporters/otlp/otlptrace/internal/tracetransform"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	coltracepb "go.opentelemetry.io/proto/otlp/collector/trace/v1"
	"google.golang.org/protobuf/proto"
)

const (
	traceSuffix = "/v1/traces"
)

type Exporter struct {
	endpoint string
	headers  map[string]string
}

func NewOtlpHttpExporter(endpoint string, headers map[string]string) (*Exporter, error) {
	_, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}
	return &Exporter{
		endpoint: endpoint,
		headers:  headers,
	}, nil
}

func (e *Exporter) ExportSpans(ctx context.Context, spans []tracesdk.ReadOnlySpan) error {
	protoSpans := Spans(spans)
	if len(protoSpans) == 0 {
		return nil
	}
	pbRequest := &coltracepb.ExportTraceServiceRequest{
		ResourceSpans: protoSpans,
	}
	rawRequest, err := proto.Marshal(pbRequest)
	if err != nil {
		return err
	}
	reader := bytes.NewBuffer(rawRequest)

	req, err := http.NewRequest(http.MethodPost, e.endpoint+traceSuffix, reader)
	if err != nil {
		return err
	}

	for key, value := range e.headers {
		req.Header.Set(key, value)
	}
	req.Header.Set("Content-Type", "application/x-protobuf")
	client := &http.Client{}
	_, err = client.Do(req)
	return err
}

func (e *Exporter) Shutdown(ctx context.Context) error {
	return nil
}
