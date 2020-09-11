package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/mediaconvert"
	"github.com/aws/aws-sdk-go/service/mediaconvert/mediaconvertiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type MediaConvertActivities struct {
	client mediaconvertiface.MediaConvertAPI
}

func NewMediaConvertActivities(session *session.Session, config ...*aws.Config) *MediaConvertActivities {
	client := mediaconvert.New(session, config...)
	return &MediaConvertActivities{client: client}
}

func (a *MediaConvertActivities) AssociateCertificate(ctx context.Context, input *mediaconvert.AssociateCertificateInput) (*mediaconvert.AssociateCertificateOutput, error) {
	return a.client.AssociateCertificateWithContext(ctx, input)
}

func (a *MediaConvertActivities) CancelJob(ctx context.Context, input *mediaconvert.CancelJobInput) (*mediaconvert.CancelJobOutput, error) {
	return a.client.CancelJobWithContext(ctx, input)
}

func (a *MediaConvertActivities) CreateJob(ctx context.Context, input *mediaconvert.CreateJobInput) (*mediaconvert.CreateJobOutput, error) {
	return a.client.CreateJobWithContext(ctx, input)
}

func (a *MediaConvertActivities) CreateJobTemplate(ctx context.Context, input *mediaconvert.CreateJobTemplateInput) (*mediaconvert.CreateJobTemplateOutput, error) {
	return a.client.CreateJobTemplateWithContext(ctx, input)
}

func (a *MediaConvertActivities) CreatePreset(ctx context.Context, input *mediaconvert.CreatePresetInput) (*mediaconvert.CreatePresetOutput, error) {
	return a.client.CreatePresetWithContext(ctx, input)
}

func (a *MediaConvertActivities) CreateQueue(ctx context.Context, input *mediaconvert.CreateQueueInput) (*mediaconvert.CreateQueueOutput, error) {
	return a.client.CreateQueueWithContext(ctx, input)
}

func (a *MediaConvertActivities) DeleteJobTemplate(ctx context.Context, input *mediaconvert.DeleteJobTemplateInput) (*mediaconvert.DeleteJobTemplateOutput, error) {
	return a.client.DeleteJobTemplateWithContext(ctx, input)
}

func (a *MediaConvertActivities) DeletePreset(ctx context.Context, input *mediaconvert.DeletePresetInput) (*mediaconvert.DeletePresetOutput, error) {
	return a.client.DeletePresetWithContext(ctx, input)
}

func (a *MediaConvertActivities) DeleteQueue(ctx context.Context, input *mediaconvert.DeleteQueueInput) (*mediaconvert.DeleteQueueOutput, error) {
	return a.client.DeleteQueueWithContext(ctx, input)
}

func (a *MediaConvertActivities) DescribeEndpoints(ctx context.Context, input *mediaconvert.DescribeEndpointsInput) (*mediaconvert.DescribeEndpointsOutput, error) {
	return a.client.DescribeEndpointsWithContext(ctx, input)
}

func (a *MediaConvertActivities) DisassociateCertificate(ctx context.Context, input *mediaconvert.DisassociateCertificateInput) (*mediaconvert.DisassociateCertificateOutput, error) {
	return a.client.DisassociateCertificateWithContext(ctx, input)
}

func (a *MediaConvertActivities) GetJob(ctx context.Context, input *mediaconvert.GetJobInput) (*mediaconvert.GetJobOutput, error) {
	return a.client.GetJobWithContext(ctx, input)
}

func (a *MediaConvertActivities) GetJobTemplate(ctx context.Context, input *mediaconvert.GetJobTemplateInput) (*mediaconvert.GetJobTemplateOutput, error) {
	return a.client.GetJobTemplateWithContext(ctx, input)
}

func (a *MediaConvertActivities) GetPreset(ctx context.Context, input *mediaconvert.GetPresetInput) (*mediaconvert.GetPresetOutput, error) {
	return a.client.GetPresetWithContext(ctx, input)
}

func (a *MediaConvertActivities) GetQueue(ctx context.Context, input *mediaconvert.GetQueueInput) (*mediaconvert.GetQueueOutput, error) {
	return a.client.GetQueueWithContext(ctx, input)
}

func (a *MediaConvertActivities) ListJobTemplates(ctx context.Context, input *mediaconvert.ListJobTemplatesInput) (*mediaconvert.ListJobTemplatesOutput, error) {
	return a.client.ListJobTemplatesWithContext(ctx, input)
}

func (a *MediaConvertActivities) ListJobs(ctx context.Context, input *mediaconvert.ListJobsInput) (*mediaconvert.ListJobsOutput, error) {
	return a.client.ListJobsWithContext(ctx, input)
}

func (a *MediaConvertActivities) ListPresets(ctx context.Context, input *mediaconvert.ListPresetsInput) (*mediaconvert.ListPresetsOutput, error) {
	return a.client.ListPresetsWithContext(ctx, input)
}

func (a *MediaConvertActivities) ListQueues(ctx context.Context, input *mediaconvert.ListQueuesInput) (*mediaconvert.ListQueuesOutput, error) {
	return a.client.ListQueuesWithContext(ctx, input)
}

func (a *MediaConvertActivities) ListTagsForResource(ctx context.Context, input *mediaconvert.ListTagsForResourceInput) (*mediaconvert.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *MediaConvertActivities) TagResource(ctx context.Context, input *mediaconvert.TagResourceInput) (*mediaconvert.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *MediaConvertActivities) UntagResource(ctx context.Context, input *mediaconvert.UntagResourceInput) (*mediaconvert.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}

func (a *MediaConvertActivities) UpdateJobTemplate(ctx context.Context, input *mediaconvert.UpdateJobTemplateInput) (*mediaconvert.UpdateJobTemplateOutput, error) {
	return a.client.UpdateJobTemplateWithContext(ctx, input)
}

func (a *MediaConvertActivities) UpdatePreset(ctx context.Context, input *mediaconvert.UpdatePresetInput) (*mediaconvert.UpdatePresetOutput, error) {
	return a.client.UpdatePresetWithContext(ctx, input)
}

func (a *MediaConvertActivities) UpdateQueue(ctx context.Context, input *mediaconvert.UpdateQueueInput) (*mediaconvert.UpdateQueueOutput, error) {
	return a.client.UpdateQueueWithContext(ctx, input)
}
