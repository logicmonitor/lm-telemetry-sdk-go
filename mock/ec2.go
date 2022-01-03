package mock

import "github.com/aws/aws-sdk-go/aws/ec2metadata"

type MockClient struct {
	available                   func() bool
	getInstanceIdentityDocument func() (ec2metadata.EC2InstanceIdentityDocument, error)
	getMetadata                 func(p string) (string, error)
}

func (mc *MockClient) Available() bool {
	if mc.available != nil {
		return mc.available()
	}
	return true
}

func (mc *MockClient) GetInstanceIdentityDocument() (ec2metadata.EC2InstanceIdentityDocument, error){
	if mc.getInstanceIdentityDocument != nil{
		return mc.getInstanceIdentityDocument()
	}
}


func (mc *MockClient)
