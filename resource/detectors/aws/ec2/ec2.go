package ec2

import (
	"context"
	"fmt"

	otelcontribaws "go.opentelemetry.io/contrib/detectors/aws"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"

	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/logicmonitor/lm-telemetry-sdk-go/utils"
	"go.opentelemetry.io/otel/attribute"
)

const (
	ec2ArnFormat = "arn:aws:ec2:%s:%s:instance/%s" //arn:aws:ec2:<REGION>:<ACCOUNT_ID>:instance/<instance-id>
	awsARN       = "aws.arn"
)

func NewResourceDetector() resource.Detector {
	return &EC2{
		otelAWS: &otelcontribaws.AWS{},
	}
}

type EC2 struct {
	otelAWS resource.Detector
}

func (ec2 *EC2) Detect(ctx context.Context) (*resource.Resource, error) {
	awsRes, err := ec2.otelAWS.Detect(ctx)
	if err != nil {
		return awsRes, err
	} else if awsRes == nil {
		return nil, err
	}
	attributes := make([]attribute.KeyValue, 0, 1)
	identityDocument, err := getEc2InstanceIdentityDocument()
	if err == nil {
		privateIPattribute := attribute.Key("private.IP").String(identityDocument.PrivateIP)
		attributes = append(attributes, privateIPattribute)
	}

	attributes = append(attributes, semconv.CloudPlatformAWSEC2)

	ec2Arn := fmt.Sprintf(ec2ArnFormat, identityDocument.Region, identityDocument.AccountID, identityDocument.InstanceID)
	arnAttribute := attribute.Key(awsARN).String(ec2Arn)
	attributes = append(attributes, arnAttribute)

	res := resource.NewSchemaless(attributes...)

	envAttributes := utils.GetServiceDetails()
	mergedRes, err := utils.AddEnvResAttributes(res, envAttributes)
	if err != nil {
		return resource.Merge(awsRes, res)
	}

	return resource.Merge(awsRes, mergedRes)
}

var getEc2InstanceIdentityDocument = func() (ec2metadata.EC2InstanceIdentityDocument, error) {
	s, err := session.NewSession()
	if err != nil {
		return ec2metadata.EC2InstanceIdentityDocument{}, err
	}

	return ec2metadata.New(s).GetInstanceIdentityDocument()
}
