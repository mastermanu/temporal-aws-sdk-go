package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/xray"
	"github.com/aws/aws-sdk-go/service/xray/xrayiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type XRayActivities struct {
	client xrayiface.XRayAPI
}

func NewXRayActivities(session *session.Session, config ...*aws.Config) *XRayActivities {
	client := xray.New(session, config...)
	return &XRayActivities{client: client}
}

func (a *XRayActivities) BatchGetTraces(ctx context.Context, input *xray.BatchGetTracesInput) (*xray.BatchGetTracesOutput, error) {
	return a.client.BatchGetTracesWithContext(ctx, input)
}

func (a *XRayActivities) CreateGroup(ctx context.Context, input *xray.CreateGroupInput) (*xray.CreateGroupOutput, error) {
	return a.client.CreateGroupWithContext(ctx, input)
}

func (a *XRayActivities) CreateSamplingRule(ctx context.Context, input *xray.CreateSamplingRuleInput) (*xray.CreateSamplingRuleOutput, error) {
	return a.client.CreateSamplingRuleWithContext(ctx, input)
}

func (a *XRayActivities) DeleteGroup(ctx context.Context, input *xray.DeleteGroupInput) (*xray.DeleteGroupOutput, error) {
	return a.client.DeleteGroupWithContext(ctx, input)
}

func (a *XRayActivities) DeleteSamplingRule(ctx context.Context, input *xray.DeleteSamplingRuleInput) (*xray.DeleteSamplingRuleOutput, error) {
	return a.client.DeleteSamplingRuleWithContext(ctx, input)
}

func (a *XRayActivities) GetEncryptionConfig(ctx context.Context, input *xray.GetEncryptionConfigInput) (*xray.GetEncryptionConfigOutput, error) {
	return a.client.GetEncryptionConfigWithContext(ctx, input)
}

func (a *XRayActivities) GetGroup(ctx context.Context, input *xray.GetGroupInput) (*xray.GetGroupOutput, error) {
	return a.client.GetGroupWithContext(ctx, input)
}

func (a *XRayActivities) GetGroups(ctx context.Context, input *xray.GetGroupsInput) (*xray.GetGroupsOutput, error) {
	return a.client.GetGroupsWithContext(ctx, input)
}

func (a *XRayActivities) GetSamplingRules(ctx context.Context, input *xray.GetSamplingRulesInput) (*xray.GetSamplingRulesOutput, error) {
	return a.client.GetSamplingRulesWithContext(ctx, input)
}

func (a *XRayActivities) GetSamplingStatisticSummaries(ctx context.Context, input *xray.GetSamplingStatisticSummariesInput) (*xray.GetSamplingStatisticSummariesOutput, error) {
	return a.client.GetSamplingStatisticSummariesWithContext(ctx, input)
}

func (a *XRayActivities) GetSamplingTargets(ctx context.Context, input *xray.GetSamplingTargetsInput) (*xray.GetSamplingTargetsOutput, error) {
	return a.client.GetSamplingTargetsWithContext(ctx, input)
}

func (a *XRayActivities) GetServiceGraph(ctx context.Context, input *xray.GetServiceGraphInput) (*xray.GetServiceGraphOutput, error) {
	return a.client.GetServiceGraphWithContext(ctx, input)
}

func (a *XRayActivities) GetTimeSeriesServiceStatistics(ctx context.Context, input *xray.GetTimeSeriesServiceStatisticsInput) (*xray.GetTimeSeriesServiceStatisticsOutput, error) {
	return a.client.GetTimeSeriesServiceStatisticsWithContext(ctx, input)
}

func (a *XRayActivities) GetTraceGraph(ctx context.Context, input *xray.GetTraceGraphInput) (*xray.GetTraceGraphOutput, error) {
	return a.client.GetTraceGraphWithContext(ctx, input)
}

func (a *XRayActivities) GetTraceSummaries(ctx context.Context, input *xray.GetTraceSummariesInput) (*xray.GetTraceSummariesOutput, error) {
	return a.client.GetTraceSummariesWithContext(ctx, input)
}

func (a *XRayActivities) ListTagsForResource(ctx context.Context, input *xray.ListTagsForResourceInput) (*xray.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *XRayActivities) PutEncryptionConfig(ctx context.Context, input *xray.PutEncryptionConfigInput) (*xray.PutEncryptionConfigOutput, error) {
	return a.client.PutEncryptionConfigWithContext(ctx, input)
}

func (a *XRayActivities) PutTelemetryRecords(ctx context.Context, input *xray.PutTelemetryRecordsInput) (*xray.PutTelemetryRecordsOutput, error) {
	return a.client.PutTelemetryRecordsWithContext(ctx, input)
}

func (a *XRayActivities) PutTraceSegments(ctx context.Context, input *xray.PutTraceSegmentsInput) (*xray.PutTraceSegmentsOutput, error) {
	return a.client.PutTraceSegmentsWithContext(ctx, input)
}

func (a *XRayActivities) TagResource(ctx context.Context, input *xray.TagResourceInput) (*xray.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *XRayActivities) UntagResource(ctx context.Context, input *xray.UntagResourceInput) (*xray.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}

func (a *XRayActivities) UpdateGroup(ctx context.Context, input *xray.UpdateGroupInput) (*xray.UpdateGroupOutput, error) {
	return a.client.UpdateGroupWithContext(ctx, input)
}

func (a *XRayActivities) UpdateSamplingRule(ctx context.Context, input *xray.UpdateSamplingRuleInput) (*xray.UpdateSamplingRuleOutput, error) {
	return a.client.UpdateSamplingRuleWithContext(ctx, input)
}
