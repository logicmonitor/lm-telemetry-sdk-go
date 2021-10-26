package mock

import (
	"github.com/aws/aws-sdk-go/service/lambda"
)

//LambdaMock mocks GetFunction behaviour of aws lambda sdk
type LambdaMock struct {
	Output *lambda.GetFunctionOutput
	Err    error
}

//GetFunction returns mocked lambda details
func (lm LambdaMock) GetFunction(input *lambda.GetFunctionInput) (*lambda.GetFunctionOutput, error) {
	return lm.Output, lm.Err
}
