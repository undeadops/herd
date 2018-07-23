// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package rekognitioniface provides an interface to enable mocking the Amazon Rekognition service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package rekognitioniface

import (
	"github.com/aws/aws-sdk-go-v2/service/rekognition"
)

// RekognitionAPI provides an interface to enable mocking the
// rekognition.Rekognition service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Rekognition.
//    func myFunc(svc rekognitioniface.RekognitionAPI) bool {
//        // Make svc.CompareFaces request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := rekognition.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockRekognitionClient struct {
//        rekognitioniface.RekognitionAPI
//    }
//    func (m *mockRekognitionClient) CompareFaces(input *rekognition.CompareFacesInput) (*rekognition.CompareFacesOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockRekognitionClient{}
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
type RekognitionAPI interface {
	CompareFacesRequest(*rekognition.CompareFacesInput) rekognition.CompareFacesRequest

	CreateCollectionRequest(*rekognition.CreateCollectionInput) rekognition.CreateCollectionRequest

	CreateStreamProcessorRequest(*rekognition.CreateStreamProcessorInput) rekognition.CreateStreamProcessorRequest

	DeleteCollectionRequest(*rekognition.DeleteCollectionInput) rekognition.DeleteCollectionRequest

	DeleteFacesRequest(*rekognition.DeleteFacesInput) rekognition.DeleteFacesRequest

	DeleteStreamProcessorRequest(*rekognition.DeleteStreamProcessorInput) rekognition.DeleteStreamProcessorRequest

	DescribeStreamProcessorRequest(*rekognition.DescribeStreamProcessorInput) rekognition.DescribeStreamProcessorRequest

	DetectFacesRequest(*rekognition.DetectFacesInput) rekognition.DetectFacesRequest

	DetectLabelsRequest(*rekognition.DetectLabelsInput) rekognition.DetectLabelsRequest

	DetectModerationLabelsRequest(*rekognition.DetectModerationLabelsInput) rekognition.DetectModerationLabelsRequest

	DetectTextRequest(*rekognition.DetectTextInput) rekognition.DetectTextRequest

	GetCelebrityInfoRequest(*rekognition.GetCelebrityInfoInput) rekognition.GetCelebrityInfoRequest

	GetCelebrityRecognitionRequest(*rekognition.GetCelebrityRecognitionInput) rekognition.GetCelebrityRecognitionRequest

	GetContentModerationRequest(*rekognition.GetContentModerationInput) rekognition.GetContentModerationRequest

	GetFaceDetectionRequest(*rekognition.GetFaceDetectionInput) rekognition.GetFaceDetectionRequest

	GetFaceSearchRequest(*rekognition.GetFaceSearchInput) rekognition.GetFaceSearchRequest

	GetLabelDetectionRequest(*rekognition.GetLabelDetectionInput) rekognition.GetLabelDetectionRequest

	GetPersonTrackingRequest(*rekognition.GetPersonTrackingInput) rekognition.GetPersonTrackingRequest

	IndexFacesRequest(*rekognition.IndexFacesInput) rekognition.IndexFacesRequest

	ListCollectionsRequest(*rekognition.ListCollectionsInput) rekognition.ListCollectionsRequest

	ListFacesRequest(*rekognition.ListFacesInput) rekognition.ListFacesRequest

	ListStreamProcessorsRequest(*rekognition.ListStreamProcessorsInput) rekognition.ListStreamProcessorsRequest

	RecognizeCelebritiesRequest(*rekognition.RecognizeCelebritiesInput) rekognition.RecognizeCelebritiesRequest

	SearchFacesRequest(*rekognition.SearchFacesInput) rekognition.SearchFacesRequest

	SearchFacesByImageRequest(*rekognition.SearchFacesByImageInput) rekognition.SearchFacesByImageRequest

	StartCelebrityRecognitionRequest(*rekognition.StartCelebrityRecognitionInput) rekognition.StartCelebrityRecognitionRequest

	StartContentModerationRequest(*rekognition.StartContentModerationInput) rekognition.StartContentModerationRequest

	StartFaceDetectionRequest(*rekognition.StartFaceDetectionInput) rekognition.StartFaceDetectionRequest

	StartFaceSearchRequest(*rekognition.StartFaceSearchInput) rekognition.StartFaceSearchRequest

	StartLabelDetectionRequest(*rekognition.StartLabelDetectionInput) rekognition.StartLabelDetectionRequest

	StartPersonTrackingRequest(*rekognition.StartPersonTrackingInput) rekognition.StartPersonTrackingRequest

	StartStreamProcessorRequest(*rekognition.StartStreamProcessorInput) rekognition.StartStreamProcessorRequest

	StopStreamProcessorRequest(*rekognition.StopStreamProcessorInput) rekognition.StopStreamProcessorRequest
}

var _ RekognitionAPI = (*rekognition.Rekognition)(nil)
