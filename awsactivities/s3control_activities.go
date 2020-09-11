package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3control"
	"github.com/aws/aws-sdk-go/service/s3control/s3controliface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type S3ControlActivities struct {
	client s3controliface.S3ControlAPI
}

func NewS3ControlActivities(session *session.Session, config ...*aws.Config) *S3ControlActivities {
	client := s3control.New(session, config...)
	return &S3ControlActivities{client: client}
}

func (a *S3ControlActivities) CreateAccessPoint(ctx context.Context, input *s3control.CreateAccessPointInput) (*s3control.CreateAccessPointOutput, error) {
	return a.client.CreateAccessPointWithContext(ctx, input)
}

func (a *S3ControlActivities) CreateJob(ctx context.Context, input *s3control.CreateJobInput) (*s3control.CreateJobOutput, error) {
	return a.client.CreateJobWithContext(ctx, input)
}

func (a *S3ControlActivities) DeleteAccessPoint(ctx context.Context, input *s3control.DeleteAccessPointInput) (*s3control.DeleteAccessPointOutput, error) {
	return a.client.DeleteAccessPointWithContext(ctx, input)
}

func (a *S3ControlActivities) DeleteAccessPointPolicy(ctx context.Context, input *s3control.DeleteAccessPointPolicyInput) (*s3control.DeleteAccessPointPolicyOutput, error) {
	return a.client.DeleteAccessPointPolicyWithContext(ctx, input)
}

func (a *S3ControlActivities) DeleteJobTagging(ctx context.Context, input *s3control.DeleteJobTaggingInput) (*s3control.DeleteJobTaggingOutput, error) {
	return a.client.DeleteJobTaggingWithContext(ctx, input)
}

func (a *S3ControlActivities) DeletePublicAccessBlock(ctx context.Context, input *s3control.DeletePublicAccessBlockInput) (*s3control.DeletePublicAccessBlockOutput, error) {
	return a.client.DeletePublicAccessBlockWithContext(ctx, input)
}

func (a *S3ControlActivities) DescribeJob(ctx context.Context, input *s3control.DescribeJobInput) (*s3control.DescribeJobOutput, error) {
	return a.client.DescribeJobWithContext(ctx, input)
}

func (a *S3ControlActivities) GetAccessPoint(ctx context.Context, input *s3control.GetAccessPointInput) (*s3control.GetAccessPointOutput, error) {
	return a.client.GetAccessPointWithContext(ctx, input)
}

func (a *S3ControlActivities) GetAccessPointPolicy(ctx context.Context, input *s3control.GetAccessPointPolicyInput) (*s3control.GetAccessPointPolicyOutput, error) {
	return a.client.GetAccessPointPolicyWithContext(ctx, input)
}

func (a *S3ControlActivities) GetAccessPointPolicyStatus(ctx context.Context, input *s3control.GetAccessPointPolicyStatusInput) (*s3control.GetAccessPointPolicyStatusOutput, error) {
	return a.client.GetAccessPointPolicyStatusWithContext(ctx, input)
}

func (a *S3ControlActivities) GetJobTagging(ctx context.Context, input *s3control.GetJobTaggingInput) (*s3control.GetJobTaggingOutput, error) {
	return a.client.GetJobTaggingWithContext(ctx, input)
}

func (a *S3ControlActivities) GetPublicAccessBlock(ctx context.Context, input *s3control.GetPublicAccessBlockInput) (*s3control.GetPublicAccessBlockOutput, error) {
	return a.client.GetPublicAccessBlockWithContext(ctx, input)
}

func (a *S3ControlActivities) ListAccessPoints(ctx context.Context, input *s3control.ListAccessPointsInput) (*s3control.ListAccessPointsOutput, error) {
	return a.client.ListAccessPointsWithContext(ctx, input)
}

func (a *S3ControlActivities) ListJobs(ctx context.Context, input *s3control.ListJobsInput) (*s3control.ListJobsOutput, error) {
	return a.client.ListJobsWithContext(ctx, input)
}

func (a *S3ControlActivities) PutAccessPointPolicy(ctx context.Context, input *s3control.PutAccessPointPolicyInput) (*s3control.PutAccessPointPolicyOutput, error) {
	return a.client.PutAccessPointPolicyWithContext(ctx, input)
}

func (a *S3ControlActivities) PutJobTagging(ctx context.Context, input *s3control.PutJobTaggingInput) (*s3control.PutJobTaggingOutput, error) {
	return a.client.PutJobTaggingWithContext(ctx, input)
}

func (a *S3ControlActivities) PutPublicAccessBlock(ctx context.Context, input *s3control.PutPublicAccessBlockInput) (*s3control.PutPublicAccessBlockOutput, error) {
	return a.client.PutPublicAccessBlockWithContext(ctx, input)
}

func (a *S3ControlActivities) UpdateJobPriority(ctx context.Context, input *s3control.UpdateJobPriorityInput) (*s3control.UpdateJobPriorityOutput, error) {
	return a.client.UpdateJobPriorityWithContext(ctx, input)
}

func (a *S3ControlActivities) UpdateJobStatus(ctx context.Context, input *s3control.UpdateJobStatusInput) (*s3control.UpdateJobStatusOutput, error) {
	return a.client.UpdateJobStatusWithContext(ctx, input)
}
