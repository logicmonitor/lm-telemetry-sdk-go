package lambda

import (
	"context"
	"errors"
	"os"

	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
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

	//arnFormat = "arn:aws:lambda:%s:%s:function:%s" //arn:aws:lambda:<region>:<account-id>:function:<name>
)

var (
	errNotOnLambda = errors.New("process is not on Lambda, cannot detect environment variables from lambda")
)

type Client interface {
	GetCallerIdentity(input *sts.GetCallerIdentityInput) (*sts.GetCallerIdentityOutput, error)
}

var getClient = func() Client {
	sess, _ := session.NewSession()
	return sts.New(sess)
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
	functionVersion := lambdaFunctionVersion()
	functionID := getAWSLambdaARN(ctx)
	accountID, _ := getAWSAccountID()

	attributes := []attribute.KeyValue{
		semconv.CloudProviderAWS,
		semconv.CloudPlatformAWSLambda,
	}

	if functionID != "" {
		attributes = append(attributes, attribute.String(string(semconv.FaaSIDKey), functionID))
	} else {
		attributes = append(attributes, []attribute.KeyValue{
			attribute.String(string(semconv.FaaSNameKey), functionName),
			attribute.String(string(semconv.CloudAccountIDKey), accountID),
			attribute.String(string(semconv.CloudRegionKey), awsRegion),
			attribute.String(string(semconv.FaaSVersionKey), functionVersion),
		}...)
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

var getAWSLambdaARN = func(ctx context.Context) string {
	lc, ok := lambdacontext.FromContext(ctx)
	if ok {
		return lc.InvokedFunctionArn
	}
	return ""
}

func getAWSAccountID() (string, error) {
	svc := getClient()
	input := &sts.GetCallerIdentityInput{}
	result, err := svc.GetCallerIdentity(input)
	if err != nil {
		return "", err
	}
	return *result.Account, nil
}

//NewResourceDetector will return an implementation for aws lambda resource detector
func NewResourceDetector() resource.Detector {
	return &Lambda{}
}
