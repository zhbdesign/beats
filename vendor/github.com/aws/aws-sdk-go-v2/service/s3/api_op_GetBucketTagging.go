// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package s3

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/GetBucketTaggingRequest
type GetBucketTaggingInput struct {
	_ struct{} `type:"structure"`

	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`
}

// String returns the string representation
func (s GetBucketTaggingInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetBucketTaggingInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetBucketTaggingInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *GetBucketTaggingInput) getBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetBucketTaggingInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.Bucket != nil {
		v := *s.Bucket

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "Bucket", protocol.StringValue(v), metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/GetBucketTaggingOutput
type GetBucketTaggingOutput struct {
	_ struct{} `type:"structure"`

	// TagSet is a required field
	TagSet []Tag `locationNameList:"Tag" type:"list" required:"true"`
}

// String returns the string representation
func (s GetBucketTaggingOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetBucketTaggingOutput) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.TagSet) > 0 {
		v := s.TagSet

		metadata := protocol.Metadata{ListLocationName: "Tag"}
		ls0 := e.List(protocol.BodyTarget, "TagSet", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	return nil
}

const opGetBucketTagging = "GetBucketTagging"

// GetBucketTaggingRequest returns a request value for making API operation for
// Amazon Simple Storage Service.
//
// Returns the tag set associated with the bucket.
//
//    // Example sending a request using GetBucketTaggingRequest.
//    req := client.GetBucketTaggingRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/s3-2006-03-01/GetBucketTagging
func (c *Client) GetBucketTaggingRequest(input *GetBucketTaggingInput) GetBucketTaggingRequest {
	op := &aws.Operation{
		Name:       opGetBucketTagging,
		HTTPMethod: "GET",
		HTTPPath:   "/{Bucket}?tagging",
	}

	if input == nil {
		input = &GetBucketTaggingInput{}
	}

	req := c.newRequest(op, input, &GetBucketTaggingOutput{})
	return GetBucketTaggingRequest{Request: req, Input: input, Copy: c.GetBucketTaggingRequest}
}

// GetBucketTaggingRequest is the request type for the
// GetBucketTagging API operation.
type GetBucketTaggingRequest struct {
	*aws.Request
	Input *GetBucketTaggingInput
	Copy  func(*GetBucketTaggingInput) GetBucketTaggingRequest
}

// Send marshals and sends the GetBucketTagging API request.
func (r GetBucketTaggingRequest) Send(ctx context.Context) (*GetBucketTaggingResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetBucketTaggingResponse{
		GetBucketTaggingOutput: r.Request.Data.(*GetBucketTaggingOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetBucketTaggingResponse is the response type for the
// GetBucketTagging API operation.
type GetBucketTaggingResponse struct {
	*GetBucketTaggingOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetBucketTagging request.
func (r *GetBucketTaggingResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}