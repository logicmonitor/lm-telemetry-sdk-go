# LM-telemetry-sdk-go, go sdk for OpenTelemetry by LogicMonitor

_NOTE: This is in private beta._

### LM-telemetry-sdk-go

1. LM-telemetry-sdk-go is a go SDK  
2. Aims to minimize adding initialization code for opentelemetry tracing, and assumes default values
3. It has implementation for cloud specific resource detectors

### Installation

```bash
go get github.com/logicmonitor/lm-telemetry-sdk-go
```

### Usage

#### Resource detection for cloud environments

##### AWS 

###### EC2
``` go

import github.com/logicmonitor/lm-telemetry-sdk-go/resource/detectors/aws/ec2
...
lmEc2Detector := ec2.NewResourceDetector()
ctx := context.Background()
resource,err := lmEc2Detector.Detect(ctx)

```


###### ECS

``` go

import github.com/logicmonitor/lm-telemetry-sdk-go/resource/detectors/aws/ecs
...
lmEcsDetector := ecs.NewResourceDetector()
ctx := context.Background()
resource,err := lmEcsDetector.Detect(ctx)

```

###### EKS

``` go

import github.com/logicmonitor/lm-telemetry-sdk-go/resource/detectors/aws/eks
...
lmEksDetector := eks.NewResourceDetector()
ctx := context.Background()
resource,err := lmEksDetector.Detect(ctx)

```

###### Lambda

``` go

import github.com/logicmonitor/lm-telemetry-sdk-go/resource/detectors/aws/lambda
...
lmLambdaDetector := lambda.NewResourceDetector()
ctx := context.Background()
resource,err := lmLambdaDetector.Detect(ctx)

```