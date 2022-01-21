package lambda

import (
	"context"
	"errors"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

const (
	executionEnvironmentKey = "AWS_EXECUTION_ENV"
	functionName            = "AWS_LAMBDA_FUNCTION_NAME"
	region                  = "AWS_REGION"
	functionVersion         = "AWS_LAMBDA_FUNCTION_VERSION"
	colonSeperator          = ":"
)

var (
	errNotOnLambda = errors.New("process is not on Lambda, cannot detect environment variables from lambda")
)

type lambdaClient interface {
	GetFunction(input *lambda.GetFunctionInput) (*lambda.GetFunctionOutput, error)
}

var getLambdaClient = func(p client.ConfigProvider, cfgs ...*aws.Config) lambdaClient {
	return lambda.New(p, cfgs...)
}

//Lambda implements, resource.Detector for aws lambda
type Lambda struct {
}

//Detect will return a resource instance which will have attributes describing lambda
func (lm *Lambda) Detect(ctx context.Context) (*resource.Resource, error) {
	if !isAWSLambda() {
		return resource.Empty(), errNotOnLambda
	}

	functionName := awsLambdafuncName()
	awsRegion := awsRegion()
	executionEnvironment := awsLambdaExecutionEnvironment()
	functionVersion := lambdaFunctionVersion()
	functionID := getAWSLambdaARN(ctx, &functionName)
	accountID := getAWSAccountIDFromARN(functionID)

	attributes := []attribute.KeyValue{
		semconv.CloudProviderAWS,
		attribute.String(string(semconv.FaaSNameKey), functionName),
		attribute.String(string(semconv.FaaSInstanceKey), executionEnvironment),
		attribute.String(string(semconv.FaaSVersionKey), functionVersion),
		attribute.String(string(semconv.FaaSIDKey), functionID),
		attribute.String(string(semconv.CloudAccountIDKey), accountID),
		attribute.String(string(semconv.CloudRegionKey), awsRegion),
	}

	return resource.NewSchemaless(attributes...), nil
}

var isAWSLambda = func() bool {
	_, present := os.LookupEnv(functionName)

	return present
}

func awsLambdafuncName() string {
	name, _ := os.LookupEnv(functionName)
	return name
}

func awsRegion() string {
	regionVal, _ := os.LookupEnv(region)
	return regionVal
}

func lambdaFunctionVersion() string {
	version, _ := os.LookupEnv(functionVersion)
	return version
}

func awsLambdaExecutionEnvironment() string {
	executionEnv, _ := os.LookupEnv(executionEnvironmentKey)
	return executionEnv
}

var getAWSLambdaARN = func(ctx context.Context, functionName *string) string {
	lc, ok := lambdacontext.FromContext(ctx)
	if ok {
		return lc.InvokedFunctionArn
	}
	mySession := session.Must(session.NewSession())
	awsLambda := getLambdaClient(mySession)

	input := lambda.GetFunctionInput{
		FunctionName: functionName,
	}

	output, err := awsLambda.GetFunction(&input)
	if err != nil {
		return ""
	}
	return *output.Configuration.FunctionArn
}

func getAWSAccountIDFromARN(arn string) string {
	if arn == "" {
		return ""
	}
	accountID := strings.Split(arn, colonSeperator)[4]
	return accountID
}

//NewResourceDetector will return an implementation for aws lambda resource detector
func NewResourceDetector() resource.Detector {
	return &Lambda{}
}
