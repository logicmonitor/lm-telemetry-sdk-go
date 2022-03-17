package azure

import (
	"github.com/logicmonitor/lm-telemetry-sdk-go/resource/detectors/azure/function"
	"github.com/logicmonitor/lm-telemetry-sdk-go/resource/detectors/azure/vm"
	"go.opentelemetry.io/otel/sdk/resource"
)

//AzureDetectors is a list of resource detector for AWS
var AzureDetectors []resource.Detector

func init() {
	AzureDetectors = make([]resource.Detector, 0, 1)
	AzureDetectors = append(AzureDetectors,
		vm.NewResourceDetector(),
		function.NewResourceDetector(),
	)
}
