// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package apigatewayiface provides an interface to enable mocking the Amazon API Gateway service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package apigatewayiface

import (
	"github.com/aws/aws-sdk-go-v2/service/apigateway"
)

// APIGatewayAPI provides an interface to enable mocking the
// apigateway.APIGateway service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon API Gateway.
//    func myFunc(svc apigatewayiface.APIGatewayAPI) bool {
//        // Make svc.CreateApiKey request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := apigateway.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockAPIGatewayClient struct {
//        apigatewayiface.APIGatewayAPI
//    }
//    func (m *mockAPIGatewayClient) CreateApiKey(input *apigateway.CreateApiKeyInput) (*apigateway.UpdateApiKeyOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockAPIGatewayClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type APIGatewayAPI interface {
	CreateApiKeyRequest(*apigateway.CreateApiKeyInput) apigateway.CreateApiKeyRequest

	CreateAuthorizerRequest(*apigateway.CreateAuthorizerInput) apigateway.CreateAuthorizerRequest

	CreateBasePathMappingRequest(*apigateway.CreateBasePathMappingInput) apigateway.CreateBasePathMappingRequest

	CreateDeploymentRequest(*apigateway.CreateDeploymentInput) apigateway.CreateDeploymentRequest

	CreateDocumentationPartRequest(*apigateway.CreateDocumentationPartInput) apigateway.CreateDocumentationPartRequest

	CreateDocumentationVersionRequest(*apigateway.CreateDocumentationVersionInput) apigateway.CreateDocumentationVersionRequest

	CreateDomainNameRequest(*apigateway.CreateDomainNameInput) apigateway.CreateDomainNameRequest

	CreateModelRequest(*apigateway.CreateModelInput) apigateway.CreateModelRequest

	CreateRequestValidatorRequest(*apigateway.CreateRequestValidatorInput) apigateway.CreateRequestValidatorRequest

	CreateResourceRequest(*apigateway.CreateResourceInput) apigateway.CreateResourceRequest

	CreateRestApiRequest(*apigateway.CreateRestApiInput) apigateway.CreateRestApiRequest

	CreateStageRequest(*apigateway.CreateStageInput) apigateway.CreateStageRequest

	CreateUsagePlanRequest(*apigateway.CreateUsagePlanInput) apigateway.CreateUsagePlanRequest

	CreateUsagePlanKeyRequest(*apigateway.CreateUsagePlanKeyInput) apigateway.CreateUsagePlanKeyRequest

	CreateVpcLinkRequest(*apigateway.CreateVpcLinkInput) apigateway.CreateVpcLinkRequest

	DeleteApiKeyRequest(*apigateway.DeleteApiKeyInput) apigateway.DeleteApiKeyRequest

	DeleteAuthorizerRequest(*apigateway.DeleteAuthorizerInput) apigateway.DeleteAuthorizerRequest

	DeleteBasePathMappingRequest(*apigateway.DeleteBasePathMappingInput) apigateway.DeleteBasePathMappingRequest

	DeleteClientCertificateRequest(*apigateway.DeleteClientCertificateInput) apigateway.DeleteClientCertificateRequest

	DeleteDeploymentRequest(*apigateway.DeleteDeploymentInput) apigateway.DeleteDeploymentRequest

	DeleteDocumentationPartRequest(*apigateway.DeleteDocumentationPartInput) apigateway.DeleteDocumentationPartRequest

	DeleteDocumentationVersionRequest(*apigateway.DeleteDocumentationVersionInput) apigateway.DeleteDocumentationVersionRequest

	DeleteDomainNameRequest(*apigateway.DeleteDomainNameInput) apigateway.DeleteDomainNameRequest

	DeleteGatewayResponseRequest(*apigateway.DeleteGatewayResponseInput) apigateway.DeleteGatewayResponseRequest

	DeleteIntegrationRequest(*apigateway.DeleteIntegrationInput) apigateway.DeleteIntegrationRequest

	DeleteIntegrationResponseRequest(*apigateway.DeleteIntegrationResponseInput) apigateway.DeleteIntegrationResponseRequest

	DeleteMethodRequest(*apigateway.DeleteMethodInput) apigateway.DeleteMethodRequest

	DeleteMethodResponseRequest(*apigateway.DeleteMethodResponseInput) apigateway.DeleteMethodResponseRequest

	DeleteModelRequest(*apigateway.DeleteModelInput) apigateway.DeleteModelRequest

	DeleteRequestValidatorRequest(*apigateway.DeleteRequestValidatorInput) apigateway.DeleteRequestValidatorRequest

	DeleteResourceRequest(*apigateway.DeleteResourceInput) apigateway.DeleteResourceRequest

	DeleteRestApiRequest(*apigateway.DeleteRestApiInput) apigateway.DeleteRestApiRequest

	DeleteStageRequest(*apigateway.DeleteStageInput) apigateway.DeleteStageRequest

	DeleteUsagePlanRequest(*apigateway.DeleteUsagePlanInput) apigateway.DeleteUsagePlanRequest

	DeleteUsagePlanKeyRequest(*apigateway.DeleteUsagePlanKeyInput) apigateway.DeleteUsagePlanKeyRequest

	DeleteVpcLinkRequest(*apigateway.DeleteVpcLinkInput) apigateway.DeleteVpcLinkRequest

	FlushStageAuthorizersCacheRequest(*apigateway.FlushStageAuthorizersCacheInput) apigateway.FlushStageAuthorizersCacheRequest

	FlushStageCacheRequest(*apigateway.FlushStageCacheInput) apigateway.FlushStageCacheRequest

	GenerateClientCertificateRequest(*apigateway.GenerateClientCertificateInput) apigateway.GenerateClientCertificateRequest

	GetAccountRequest(*apigateway.GetAccountInput) apigateway.GetAccountRequest

	GetApiKeyRequest(*apigateway.GetApiKeyInput) apigateway.GetApiKeyRequest

	GetApiKeysRequest(*apigateway.GetApiKeysInput) apigateway.GetApiKeysRequest

	GetAuthorizerRequest(*apigateway.GetAuthorizerInput) apigateway.GetAuthorizerRequest

	GetAuthorizersRequest(*apigateway.GetAuthorizersInput) apigateway.GetAuthorizersRequest

	GetBasePathMappingRequest(*apigateway.GetBasePathMappingInput) apigateway.GetBasePathMappingRequest

	GetBasePathMappingsRequest(*apigateway.GetBasePathMappingsInput) apigateway.GetBasePathMappingsRequest

	GetClientCertificateRequest(*apigateway.GetClientCertificateInput) apigateway.GetClientCertificateRequest

	GetClientCertificatesRequest(*apigateway.GetClientCertificatesInput) apigateway.GetClientCertificatesRequest

	GetDeploymentRequest(*apigateway.GetDeploymentInput) apigateway.GetDeploymentRequest

	GetDeploymentsRequest(*apigateway.GetDeploymentsInput) apigateway.GetDeploymentsRequest

	GetDocumentationPartRequest(*apigateway.GetDocumentationPartInput) apigateway.GetDocumentationPartRequest

	GetDocumentationPartsRequest(*apigateway.GetDocumentationPartsInput) apigateway.GetDocumentationPartsRequest

	GetDocumentationVersionRequest(*apigateway.GetDocumentationVersionInput) apigateway.GetDocumentationVersionRequest

	GetDocumentationVersionsRequest(*apigateway.GetDocumentationVersionsInput) apigateway.GetDocumentationVersionsRequest

	GetDomainNameRequest(*apigateway.GetDomainNameInput) apigateway.GetDomainNameRequest

	GetDomainNamesRequest(*apigateway.GetDomainNamesInput) apigateway.GetDomainNamesRequest

	GetExportRequest(*apigateway.GetExportInput) apigateway.GetExportRequest

	GetGatewayResponseRequest(*apigateway.GetGatewayResponseInput) apigateway.GetGatewayResponseRequest

	GetGatewayResponsesRequest(*apigateway.GetGatewayResponsesInput) apigateway.GetGatewayResponsesRequest

	GetIntegrationRequest(*apigateway.GetIntegrationInput) apigateway.GetIntegrationRequest

	GetIntegrationResponseRequest(*apigateway.GetIntegrationResponseInput) apigateway.GetIntegrationResponseRequest

	GetMethodRequest(*apigateway.GetMethodInput) apigateway.GetMethodRequest

	GetMethodResponseRequest(*apigateway.GetMethodResponseInput) apigateway.GetMethodResponseRequest

	GetModelRequest(*apigateway.GetModelInput) apigateway.GetModelRequest

	GetModelTemplateRequest(*apigateway.GetModelTemplateInput) apigateway.GetModelTemplateRequest

	GetModelsRequest(*apigateway.GetModelsInput) apigateway.GetModelsRequest

	GetRequestValidatorRequest(*apigateway.GetRequestValidatorInput) apigateway.GetRequestValidatorRequest

	GetRequestValidatorsRequest(*apigateway.GetRequestValidatorsInput) apigateway.GetRequestValidatorsRequest

	GetResourceRequest(*apigateway.GetResourceInput) apigateway.GetResourceRequest

	GetResourcesRequest(*apigateway.GetResourcesInput) apigateway.GetResourcesRequest

	GetRestApiRequest(*apigateway.GetRestApiInput) apigateway.GetRestApiRequest

	GetRestApisRequest(*apigateway.GetRestApisInput) apigateway.GetRestApisRequest

	GetSdkRequest(*apigateway.GetSdkInput) apigateway.GetSdkRequest

	GetSdkTypeRequest(*apigateway.GetSdkTypeInput) apigateway.GetSdkTypeRequest

	GetSdkTypesRequest(*apigateway.GetSdkTypesInput) apigateway.GetSdkTypesRequest

	GetStageRequest(*apigateway.GetStageInput) apigateway.GetStageRequest

	GetStagesRequest(*apigateway.GetStagesInput) apigateway.GetStagesRequest

	GetTagsRequest(*apigateway.GetTagsInput) apigateway.GetTagsRequest

	GetUsageRequest(*apigateway.GetUsageInput) apigateway.GetUsageRequest

	GetUsagePlanRequest(*apigateway.GetUsagePlanInput) apigateway.GetUsagePlanRequest

	GetUsagePlanKeyRequest(*apigateway.GetUsagePlanKeyInput) apigateway.GetUsagePlanKeyRequest

	GetUsagePlanKeysRequest(*apigateway.GetUsagePlanKeysInput) apigateway.GetUsagePlanKeysRequest

	GetUsagePlansRequest(*apigateway.GetUsagePlansInput) apigateway.GetUsagePlansRequest

	GetVpcLinkRequest(*apigateway.GetVpcLinkInput) apigateway.GetVpcLinkRequest

	GetVpcLinksRequest(*apigateway.GetVpcLinksInput) apigateway.GetVpcLinksRequest

	ImportApiKeysRequest(*apigateway.ImportApiKeysInput) apigateway.ImportApiKeysRequest

	ImportDocumentationPartsRequest(*apigateway.ImportDocumentationPartsInput) apigateway.ImportDocumentationPartsRequest

	ImportRestApiRequest(*apigateway.ImportRestApiInput) apigateway.ImportRestApiRequest

	PutGatewayResponseRequest(*apigateway.PutGatewayResponseInput) apigateway.PutGatewayResponseRequest

	PutIntegrationRequest(*apigateway.PutIntegrationInput) apigateway.PutIntegrationRequest

	PutIntegrationResponseRequest(*apigateway.PutIntegrationResponseInput) apigateway.PutIntegrationResponseRequest

	PutMethodRequest(*apigateway.PutMethodInput) apigateway.PutMethodRequest

	PutMethodResponseRequest(*apigateway.PutMethodResponseInput) apigateway.PutMethodResponseRequest

	PutRestApiRequest(*apigateway.PutRestApiInput) apigateway.PutRestApiRequest

	TagResourceRequest(*apigateway.TagResourceInput) apigateway.TagResourceRequest

	TestInvokeAuthorizerRequest(*apigateway.TestInvokeAuthorizerInput) apigateway.TestInvokeAuthorizerRequest

	TestInvokeMethodRequest(*apigateway.TestInvokeMethodInput) apigateway.TestInvokeMethodRequest

	UntagResourceRequest(*apigateway.UntagResourceInput) apigateway.UntagResourceRequest

	UpdateAccountRequest(*apigateway.UpdateAccountInput) apigateway.UpdateAccountRequest

	UpdateApiKeyRequest(*apigateway.UpdateApiKeyInput) apigateway.UpdateApiKeyRequest

	UpdateAuthorizerRequest(*apigateway.UpdateAuthorizerInput) apigateway.UpdateAuthorizerRequest

	UpdateBasePathMappingRequest(*apigateway.UpdateBasePathMappingInput) apigateway.UpdateBasePathMappingRequest

	UpdateClientCertificateRequest(*apigateway.UpdateClientCertificateInput) apigateway.UpdateClientCertificateRequest

	UpdateDeploymentRequest(*apigateway.UpdateDeploymentInput) apigateway.UpdateDeploymentRequest

	UpdateDocumentationPartRequest(*apigateway.UpdateDocumentationPartInput) apigateway.UpdateDocumentationPartRequest

	UpdateDocumentationVersionRequest(*apigateway.UpdateDocumentationVersionInput) apigateway.UpdateDocumentationVersionRequest

	UpdateDomainNameRequest(*apigateway.UpdateDomainNameInput) apigateway.UpdateDomainNameRequest

	UpdateGatewayResponseRequest(*apigateway.UpdateGatewayResponseInput) apigateway.UpdateGatewayResponseRequest

	UpdateIntegrationRequest(*apigateway.UpdateIntegrationInput) apigateway.UpdateIntegrationRequest

	UpdateIntegrationResponseRequest(*apigateway.UpdateIntegrationResponseInput) apigateway.UpdateIntegrationResponseRequest

	UpdateMethodRequest(*apigateway.UpdateMethodInput) apigateway.UpdateMethodRequest

	UpdateMethodResponseRequest(*apigateway.UpdateMethodResponseInput) apigateway.UpdateMethodResponseRequest

	UpdateModelRequest(*apigateway.UpdateModelInput) apigateway.UpdateModelRequest

	UpdateRequestValidatorRequest(*apigateway.UpdateRequestValidatorInput) apigateway.UpdateRequestValidatorRequest

	UpdateResourceRequest(*apigateway.UpdateResourceInput) apigateway.UpdateResourceRequest

	UpdateRestApiRequest(*apigateway.UpdateRestApiInput) apigateway.UpdateRestApiRequest

	UpdateStageRequest(*apigateway.UpdateStageInput) apigateway.UpdateStageRequest

	UpdateUsageRequest(*apigateway.UpdateUsageInput) apigateway.UpdateUsageRequest

	UpdateUsagePlanRequest(*apigateway.UpdateUsagePlanInput) apigateway.UpdateUsagePlanRequest

	UpdateVpcLinkRequest(*apigateway.UpdateVpcLinkInput) apigateway.UpdateVpcLinkRequest
}

var _ APIGatewayAPI = (*apigateway.APIGateway)(nil)