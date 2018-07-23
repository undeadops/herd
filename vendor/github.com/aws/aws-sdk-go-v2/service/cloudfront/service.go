// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudfront

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restxml"
)

// CloudFront provides the API operation methods for making requests to
// Amazon CloudFront. See this package's package overview docs
// for details on the service.
//
// CloudFront methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type CloudFront struct {
	*aws.Client
}

// Used for custom client initialization logic
var initClient func(*CloudFront)

// Used for custom request initialization logic
var initRequest func(*CloudFront, *aws.Request)

// Service information constants
const (
	ServiceName = "cloudfront" // Service endpoint prefix API calls made to.
	EndpointsID = ServiceName  // Service ID for Regions and Endpoints metadata.
)

// New creates a new instance of the CloudFront client with a config.
//
// Example:
//     // Create a CloudFront client from just a config.
//     svc := cloudfront.New(myConfig)
func New(config aws.Config) *CloudFront {
	var signingName string
	signingRegion := config.Region

	svc := &CloudFront{
		Client: aws.NewClient(
			config,
			aws.Metadata{
				ServiceName:   ServiceName,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				APIVersion:    "2017-10-30",
			},
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(restxml.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(restxml.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(restxml.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(restxml.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc)
	}

	return svc
}

// newRequest creates a new request for a CloudFront operation and runs any
// custom request initialization.
func (c *CloudFront) newRequest(op *aws.Operation, params, data interface{}) *aws.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(c, req)
	}

	return req
}
