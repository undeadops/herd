// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecr

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
)

// ECR provides the API operation methods for making requests to
// Amazon EC2 Container Registry. See this package's package overview docs
// for details on the service.
//
// ECR methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type ECR struct {
	*aws.Client
}

// Used for custom client initialization logic
var initClient func(*ECR)

// Used for custom request initialization logic
var initRequest func(*ECR, *aws.Request)

// Service information constants
const (
	ServiceName = "ecr"       // Service endpoint prefix API calls made to.
	EndpointsID = ServiceName // Service ID for Regions and Endpoints metadata.
)

// New creates a new instance of the ECR client with a config.
//
// Example:
//     // Create a ECR client from just a config.
//     svc := ecr.New(myConfig)
func New(config aws.Config) *ECR {
	var signingName string
	signingRegion := config.Region

	svc := &ECR{
		Client: aws.NewClient(
			config,
			aws.Metadata{
				ServiceName:   ServiceName,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				APIVersion:    "2015-09-21",
				JSONVersion:   "1.1",
				TargetPrefix:  "AmazonEC2ContainerRegistry_V20150921",
			},
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(jsonrpc.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(jsonrpc.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(jsonrpc.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(jsonrpc.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc)
	}

	return svc
}

// newRequest creates a new request for a ECR operation and runs any
// custom request initialization.
func (c *ECR) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(c, req)
	}

	return req
}
