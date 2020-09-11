package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/appflow"
	"github.com/aws/aws-sdk-go/service/appflow/appflowiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type AppflowActivities struct {
	client appflowiface.AppflowAPI
}

func NewAppflowActivities(session *session.Session, config ...*aws.Config) *AppflowActivities {
	client := appflow.New(session, config...)
	return &AppflowActivities{client: client}
}

func (a *AppflowActivities) CreateConnectorProfile(ctx context.Context, input *appflow.CreateConnectorProfileInput) (*appflow.CreateConnectorProfileOutput, error) {
	return a.client.CreateConnectorProfileWithContext(ctx, input)
}

func (a *AppflowActivities) CreateFlow(ctx context.Context, input *appflow.CreateFlowInput) (*appflow.CreateFlowOutput, error) {
	return a.client.CreateFlowWithContext(ctx, input)
}

func (a *AppflowActivities) DeleteConnectorProfile(ctx context.Context, input *appflow.DeleteConnectorProfileInput) (*appflow.DeleteConnectorProfileOutput, error) {
	return a.client.DeleteConnectorProfileWithContext(ctx, input)
}

func (a *AppflowActivities) DeleteFlow(ctx context.Context, input *appflow.DeleteFlowInput) (*appflow.DeleteFlowOutput, error) {
	return a.client.DeleteFlowWithContext(ctx, input)
}

func (a *AppflowActivities) DescribeConnectorEntity(ctx context.Context, input *appflow.DescribeConnectorEntityInput) (*appflow.DescribeConnectorEntityOutput, error) {
	return a.client.DescribeConnectorEntityWithContext(ctx, input)
}

func (a *AppflowActivities) DescribeConnectorProfiles(ctx context.Context, input *appflow.DescribeConnectorProfilesInput) (*appflow.DescribeConnectorProfilesOutput, error) {
	return a.client.DescribeConnectorProfilesWithContext(ctx, input)
}

func (a *AppflowActivities) DescribeConnectors(ctx context.Context, input *appflow.DescribeConnectorsInput) (*appflow.DescribeConnectorsOutput, error) {
	return a.client.DescribeConnectorsWithContext(ctx, input)
}

func (a *AppflowActivities) DescribeFlow(ctx context.Context, input *appflow.DescribeFlowInput) (*appflow.DescribeFlowOutput, error) {
	return a.client.DescribeFlowWithContext(ctx, input)
}

func (a *AppflowActivities) DescribeFlowExecutionRecords(ctx context.Context, input *appflow.DescribeFlowExecutionRecordsInput) (*appflow.DescribeFlowExecutionRecordsOutput, error) {
	return a.client.DescribeFlowExecutionRecordsWithContext(ctx, input)
}

func (a *AppflowActivities) ListConnectorEntities(ctx context.Context, input *appflow.ListConnectorEntitiesInput) (*appflow.ListConnectorEntitiesOutput, error) {
	return a.client.ListConnectorEntitiesWithContext(ctx, input)
}

func (a *AppflowActivities) ListFlows(ctx context.Context, input *appflow.ListFlowsInput) (*appflow.ListFlowsOutput, error) {
	return a.client.ListFlowsWithContext(ctx, input)
}

func (a *AppflowActivities) ListTagsForResource(ctx context.Context, input *appflow.ListTagsForResourceInput) (*appflow.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *AppflowActivities) StartFlow(ctx context.Context, input *appflow.StartFlowInput) (*appflow.StartFlowOutput, error) {
	return a.client.StartFlowWithContext(ctx, input)
}

func (a *AppflowActivities) StopFlow(ctx context.Context, input *appflow.StopFlowInput) (*appflow.StopFlowOutput, error) {
	return a.client.StopFlowWithContext(ctx, input)
}

func (a *AppflowActivities) TagResource(ctx context.Context, input *appflow.TagResourceInput) (*appflow.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *AppflowActivities) UntagResource(ctx context.Context, input *appflow.UntagResourceInput) (*appflow.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}

func (a *AppflowActivities) UpdateConnectorProfile(ctx context.Context, input *appflow.UpdateConnectorProfileInput) (*appflow.UpdateConnectorProfileOutput, error) {
	return a.client.UpdateConnectorProfileWithContext(ctx, input)
}

func (a *AppflowActivities) UpdateFlow(ctx context.Context, input *appflow.UpdateFlowInput) (*appflow.UpdateFlowOutput, error) {
	return a.client.UpdateFlowWithContext(ctx, input)
}
