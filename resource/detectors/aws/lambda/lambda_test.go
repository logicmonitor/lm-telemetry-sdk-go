package lambda

import (
	"context"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/logicmonitor/lm-telemetry-sdk-go/mock"
)

var (
	sampleARN = "arn:aws:lambda:us-east-2:123456789012:function:my-function:1"
)

func CreateIsAWSLambdaMock(islambda bool) func() bool {
	return func() bool {
		return islambda
	}
}

func getAWSLambdaARNMock(ctx context.Context) string {
	return sampleARN
}

var getLambdaClientMock = func(p client.ConfigProvider, cfgs ...*aws.Config) lambdaClient {
	return mock.LambdaMock{
		Output: &lambda.GetFunctionOutput{
			Configuration: &lambda.FunctionConfiguration{
				FunctionArn: &sampleARN,
			},
		},
	}
}

func TestDetect(t *testing.T) {
	lambdaDetector := Lambda{}
	Name := "function-1"
	os.Setenv(functionName, Name)

	oldgetAWSLambdaARN := getAWSLambdaARN
	defer func() {
		err := os.Unsetenv(functionName)
		if err != nil {
			t.Fatal("error in unsetting env var:", err)
		}
		getAWSLambdaARN = oldgetAWSLambdaARN
	}()
	//isAWSLambda = CreateIsAWSLambdaMock(true)
	getAWSLambdaARN = getAWSLambdaARNMock

	t.Run("Not AWS Lambda", func(t *testing.T) {
		oldisAWSLambda := isAWSLambda
		defer func() {
			isAWSLambda = oldisAWSLambda
		}()
		isAWSLambda = CreateIsAWSLambdaMock(false)
		_, err := lambdaDetector.Detect(context.Background())
		if err != errNotOnLambda {
			t.Fatalf("Should through Not on lambda error but throwing %v", err)
		}
	})
	t.Run("Success Case", func(t *testing.T) {

		_, err := lambdaDetector.Detect(context.Background())
		if err != nil {
			t.Fatalf("Getting error: %v", err)
		}
	})
}

func TestGetAWSLambdaARN(t *testing.T) {
	// oldgetLambdaClient := getLambdaClient
	// defer func() {
	// 	getLambdaClient = oldgetLambdaClient
	// }()
	// getLambdaClient = getLambdaClientMock

	// // t.Run("ARN in  context", func(t *testing.T) {
	// // 	ctx := context.WithValue(context.Background(), arnKey, sampleARN)
	// // 	functionName := "function-1"
	// // 	arn := getAWSLambdaARN(ctx, &functionName)
	// // 	if arn != sampleARN {
	// // 		t.Fatal("ARN value is not equal to expected one")
	// // 	}
	// // })

	// t.Run("Successful ARN", func(t *testing.T) {
	// 	//functionName := "function-1"
	// 	getAWSLambdaARN(context.Background())
	// })
}
