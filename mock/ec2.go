package mock

import "github.com/aws/aws-sdk-go/aws/ec2metadata"

type MockClient struct {
	AvailableValue bool

	InstanceIdentityDocumentValue ec2metadata.EC2InstanceIdentityDocument
	InstanceIdentityDocumentErr   error

	MetadataValue string
	MetadataErr   error
}

func (mc *MockClient) Available() bool {
	return mc.AvailableValue
}

func (mc *MockClient) GetInstanceIdentityDocument() (ec2metadata.EC2InstanceIdentityDocument, error) {
	return mc.InstanceIdentityDocumentValue, mc.InstanceIdentityDocumentErr
}

func (mc *MockClient) GetMetadata(p string) (string, error) {
	return mc.MetadataValue, mc.MetadataErr
}
