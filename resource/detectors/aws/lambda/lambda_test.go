package lambda

import (
	"context"
	"os"
	"testing"
)

var (
	sampleARN = "arn:aws:lambda:us-east-2:123456789012:function:my-function:1"
)

func CreateIsAWSLambdaMock(islambda bool) func() bool {
	return func() bool {
		return islambda
	}
}

func getAWSLambdaARNMock(ctx context.Context, functionName *string) string {
	return sampleARN
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
