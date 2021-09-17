package gcp

import (
	"github.com/logicmonitor/lm-telemetry-sdk-go/resource/detectors/gcp/cloudfunction"
	"github.com/logicmonitor/lm-telemetry-sdk-go/resource/detectors/gcp/gce"
	"github.com/logicmonitor/lm-telemetry-sdk-go/resource/detectors/gcp/gke"
	"go.opentelemetry.io/otel/sdk/resource"
)

var GCPDetectors []resource.Detector

func init() {
	GCPDetectors = make([]resource.Detector, 0, 1)
	GCPDetectors = append(GCPDetectors,
		cloudfunction.NewResourceDetector(),
		gce.NewResourceDetector(),
		gke.NewResourceDetector(),
	)
}
