package mock

import (
	"github.com/aws/aws-sdk-go/service/lambda"
)

type LambdaMock struct {
	Output *lambda.GetFunctionOutput
	Err    error
}

func (lm LambdaMock) GetFunction(input *lambda.GetFunctionInput) (*lambda.GetFunctionOutput, error) {
	return lm.Output, lm.Err
}
