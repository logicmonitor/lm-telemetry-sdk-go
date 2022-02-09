package mock

import (
	"github.com/aws/aws-sdk-go/service/sts"
)

//LambdaMock mocks GetFunction behaviour of aws lambda sdk
type LambdaMock struct {
	Output *sts.GetCallerIdentityOutput
	Err    error
}

//GetFunction returns mocked lambda details
func (lm LambdaMock) GetCallerIdentity(input *sts.GetCallerIdentityInput) (*sts.GetCallerIdentityOutput, error) {
	return lm.Output, lm.Err
}
