package ec2

import (
	"context"
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/logicmonitor/lm-telemetry-sdk-go/utils"
	otelcontribaws "go.opentelemetry.io/contrib/detectors/aws"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

const (
	ec2ArnFormat = "arn:aws:ec2:%s:%s:instance/%s" //arn:aws:ec2:<REGION>:<ACCOUNT_ID>:instance/<instance-id>
	awsARN       = "aws.arn"
)

//NewResourceDetector will return an implementation for aws ec2 resource detector
func NewResourceDetector() resource.Detector {
	return &EC2{
		otelAWS: &otelcontribaws.AWS{},
	}
}

// Client implements methods to capture EC2 environment metadata information
type Client interface {
	Available() bool
	GetInstanceIdentityDocument() (ec2metadata.EC2InstanceIdentityDocument, error)
	GetMetadata(p string) (string, error)
}

//EC2 implements resource.Detector interface, for an ec2 instance
type EC2 struct {
	c       Client
	otelAWS resource.Detector
}

/*
Detect will return a resource instance which will have attributes describing,
an ec2 instance
*/
func (detector *EC2) Detect(ctx context.Context) (*resource.Resource, error) {
	client, err := detector.client()
	if err != nil {
		return nil, err
	}

	if !client.Available() {
		return nil, nil
	}

	doc, err := client.GetInstanceIdentityDocument()
	if err != nil {
		return nil, err
	}

	attributes := []attribute.KeyValue{
		semconv.CloudProviderAWS,
		semconv.CloudPlatformAWSEC2,
		semconv.CloudRegionKey.String(doc.Region),
		semconv.CloudAvailabilityZoneKey.String(doc.AvailabilityZone),
		semconv.CloudAccountIDKey.String(doc.AccountID),
		semconv.HostIDKey.String(doc.InstanceID),
		semconv.HostImageIDKey.String(doc.ImageID),
		semconv.HostTypeKey.String(doc.InstanceType),
	}

	m := &metadata{client: client}
	m.add(semconv.HostNameKey, "hostname")

	attributes = append(attributes, m.attributes...)

	if len(m.errs) > 0 {
		err = fmt.Errorf("%w: %s", resource.ErrPartialResource, m.errs)
	}

	ec2Arn := fmt.Sprintf(ec2ArnFormat, doc.Region, doc.AccountID, doc.InstanceID)
	arnAttribute := attribute.Key(awsARN).String(ec2Arn)
	attributes = append(attributes, arnAttribute)

	envAttributes := utils.GetServiceDetails()
	attributes = append(attributes, utils.GetAttributesfromMap(envAttributes)...)
	attributes = append(attributes, semconv.CloudPlatformAWSEC2)

	return resource.NewWithAttributes(semconv.SchemaURL, attributes...), err

}

func (detector *EC2) client() (Client, error) {
	if detector.c != nil {
		return detector.c, nil
	}

	s, err := session.NewSession()
	if err != nil {
		return nil, err
	}

	return ec2metadata.New(s), nil
}

type metadata struct {
	client     Client
	errs       []error
	attributes []attribute.KeyValue
}

func (m *metadata) add(k attribute.Key, n string) {
	v, err := m.client.GetMetadata(n)
	if err == nil {
		m.attributes = append(m.attributes, k.String(v))
		return
	}

	rf, ok := err.(awserr.RequestFailure)
	if !ok {
		m.errs = append(m.errs, fmt.Errorf("%q: %w", n, err))
		return
	}

	if rf.StatusCode() == http.StatusNotFound {
		return
	}

	m.errs = append(m.errs, fmt.Errorf("%q: %d %s", n, rf.StatusCode(), rf.Code()))
}
