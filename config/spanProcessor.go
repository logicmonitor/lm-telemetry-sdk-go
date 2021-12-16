package config

import (
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func WithSimlpeSpanProcessor() Option {
	return func(c *Config) {
		c.SpanProcessor = sdktrace.NewSimpleSpanProcessor
	}
}
