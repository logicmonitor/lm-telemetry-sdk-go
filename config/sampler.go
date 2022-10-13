package config

import (
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func WithTraceIDRatioBasedSampler(fraction float64) Option {
	return func(c *Config) {
		c.Sampler = sdktrace.TraceIDRatioBased(fraction)
	}
}
