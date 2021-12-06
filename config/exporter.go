package config

import (
	"fmt"
	"os"
)

const (
	lmAccountKey     = "LOGICMONITOR_ACCOUNT"
	lmBearerTokenKey = "LOGICMONITOR_BEARER_TOKEN"

	defaultEndpoint    = "https://%s.logicmonitor.com/rest/api" //https://${LOGICMONITOR_ACCOUNT}.logicmonitor.com/rest/api
	lmTokenHeaderValue = "Bearer %s"

	authorizationHeaderKey = "Authorization"
	xLMAccountHeaderKey    = "x-logicmonitor-account"
)

/*WithHTTPTraceEndpoint returns a config option which sets
trace endpoint
*/
func WithHTTPTraceEndpoint(endpoint string) Option {
	return func(c *Config) {
		c.TraceEndpoint = endpoint
	}
}

/*WithInAppExporter will send spans directly to HTTP destination
of choice. Content-Type would be of application/x-protobuf
*/
func WithInAppExporter(endpoint string, headers map[string]string) Option {
	return func(c *Config) {
		c.InAppExporter = &sdkTraceExporter{
			TraceEndpoint: endpoint,
			Headers:       headers,
		}
	}
}

func WithDefaultInAppExporter() Option {
	return func(c *Config) {
		if c.InAppExporter == nil {
			lmAccount := os.Getenv(lmAccountKey)
			lmBearerToken := os.Getenv(lmBearerTokenKey)

			endpoint := fmt.Sprintf(defaultEndpoint, lmAccount)
			headers := map[string]string{
				authorizationHeaderKey: fmt.Sprintf(lmTokenHeaderValue, lmBearerToken),
				xLMAccountHeaderKey:    lmAccount,
			}
			c.InAppExporter = &sdkTraceExporter{
				TraceEndpoint: endpoint,
				Headers:       headers,
			}
		}
	}
}
