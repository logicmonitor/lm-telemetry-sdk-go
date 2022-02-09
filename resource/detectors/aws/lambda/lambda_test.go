package lambda

import (
	"context"
	"os"
	"testing"

	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/logicmonitor/lm-telemetry-sdk-go/mock"
)

var (
	sampleARN       = "arn:aws:lambda:us-east-2:123456789012:function:my-function:1"
	sampleAccountID = "1234567890"
)

func CreateIsAWSLambdaMock(islambda bool) func() bool {
	return func() bool {
		return islambda
	}
}

func getAWSLambdaARNMock(ctx context.Context) string {
	return sampleARN
}

var getLambdaClientMock = func() Client {
	return mock.LambdaMock{
		Output: &sts.GetCallerIdentityOutput{
			Account: &sampleAccountID,
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

	t.Run("AWS arn not found", func(t *testing.T) {
		oldgetAWSLambdaARN := getAWSLambdaARN
		getAWSLambdaARN = func(ctx context.Context) string {
			return ""
		}
		defer func() {
			getAWSLambdaARN = oldgetAWSLambdaARN
		}()
		_, err := lambdaDetector.Detect(context.Background())
		if err != nil {
			t.Fatalf("Getting error: %v", err)
		}

	})
}

func TestGetAWSLambdaARN(t *testing.T) {

	t.Run("with lambdaContext", func(t *testing.T) {
		lambdaContext := lambdacontext.LambdaContext{
			InvokedFunctionArn: sampleARN,
		}
		arn := getAWSLambdaARN(lambdacontext.NewContext(context.Background(), &lambdaContext))
		if arn != sampleARN {
			t.Fatalf("expected arn=%s, rcvd arn = %s", sampleARN, arn)
		}
	})

	t.Run("without lambdaContext", func(t *testing.T) {
		arn := getAWSLambdaARN(context.Background())
		if arn != "" {
			t.Fatalf("expected arn=%s, rcvd arn = %s", sampleARN, arn)
		}
	})
}

func TestGetAWSAccountID(t *testing.T) {
	oldgetClient := getClient
	defer func() {
		getClient = oldgetClient
	}()
	getClient = getLambdaClientMock

	t.Run("success", func(t *testing.T) {
		accountID, err := getAWSAccountID()
		if err != nil || accountID != sampleAccountID {
			t.Fatalf("expected account_id = %s, rcvd account_id = %s ; expected error = %v, rcvd error = %v", sampleAccountID, accountID, nil, err)
		}
	})

}
