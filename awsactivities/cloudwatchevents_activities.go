package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchevents"
	"github.com/aws/aws-sdk-go/service/cloudwatchevents/cloudwatcheventsiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type CloudWatchEventsActivities struct {
	client cloudwatcheventsiface.CloudWatchEventsAPI
}

func NewCloudWatchEventsActivities(session *session.Session, config ...*aws.Config) *CloudWatchEventsActivities {
	client := cloudwatchevents.New(session, config...)
	return &CloudWatchEventsActivities{client: client}
}

func (a *CloudWatchEventsActivities) ActivateEventSource(ctx context.Context, input *cloudwatchevents.ActivateEventSourceInput) (*cloudwatchevents.ActivateEventSourceOutput, error) {
	return a.client.ActivateEventSourceWithContext(ctx, input)
}

func (a *CloudWatchEventsActivities) CreateEventBus(ctx context.Context, input *cloudwatchevents.CreateEventBusInput) (*cloudwatchevents.CreateEventBusOutput, error) {
	return a.client.CreateEventBusWithContext(ctx, input)
}

func (a *CloudWatchEventsActivities) CreatePartnerEventSource(ctx context.Context, input *cloudwatchevents.CreatePartnerEventSourceInput) (*cloudwatchevents.CreatePartnerEventSourceOutput, error) {
	return a.client.CreatePartnerEventSourceWithContext(ctx, input)
}

func (a *CloudWatchEventsActivities) DeactivateEventSource(ctx context.Context, input *cloudwatchevents.DeactivateEventSourceInput) (*cloudwatchevents.DeactivateEventSourceOutput, error) {
	return a.client.DeactivateEventSourceWithContext(ctx, input)
}

func (a *CloudWatchEventsActivities) DeleteEventBus(ctx context.Context, input *cloudwatchevents.DeleteEventBusInput) (*cloudwatchevents.DeleteEventBusOutput, error) {
	return a.client.DeleteEventBusWithContext(ctx, input)
}

func (a *CloudWatchEventsActivities) DeletePartnerEventSource(ctx context.Context, input *cloudwatchevents.DeletePartnerEventSourceInput) (*cloudwatchevents.DeletePartnerEventSourceOutput, error) {
	return a.client.DeletePartnerEventSourceWithContext(ctx, input)
}

func (a *CloudWatchEventsActivities) DeleteRule(ctx context.Context, input *cloudwatchevents.DeleteRuleInput) (*cloudwatchevents.DeleteRuleOutput, error) {
	return a.client.DeleteRuleWithContext(ctx, input)
}

func (a *CloudWatchEventsActivities) DescribeEventBus(ctx context.Context, input *cloudwatchevents.DescribeEventBusInput) (*cloudwatchevents.DescribeEventBusOutput, error) {
	return a.client.DescribeEventBusWithContext(ctx, input)
}

func (a *CloudWatchEventsActivities) DescribeEventSource(ctx context.Context, input *cloudwatchevents.DescribeEventSourceInput) (*cloudwatchevents.DescribeEventSourceOutput, error) {
	return a.client.DescribeEventSourceWithContext(ctx, input)
}

func (a *CloudWatchEventsActivities) DescribePartnerEventSource(ctx context.Context, input *cloudwatchevents.DescribePartnerEventSourceInput) (*cloudwatchevents.DescribePartnerEventSourceOutput, error) {
	return a.client.DescribePartnerEventSourceWithContext(ctx, input)
}

func (a *CloudWatchEventsActivities) DescribeRule(ctx context.Context, input *cloudwatchevents.DescribeRuleInput) (*cloudwatchevents.DescribeRuleOutput, error) {
	return a.client.DescribeRuleWithContext(ctx, input)
}

func (a *CloudWatchEventsActivities) DisableRule(ctx context.Context, input *cloudwatchevents.DisableRuleInput) (*cloudwatchevents.DisableRuleOutput, error) {
	return a.client.DisableRuleWithContext(ctx, input)
}

func (a *CloudWatchEventsActivities) EnableRule(ctx context.Context, input *cloudwatchevents.EnableRuleInput) (*cloudwatchevents.EnableRuleOutput, error) {
	return a.client.EnableRuleWithContext(ctx, input)
}

func (a *CloudWatchEventsActivities) ListEventBuses(ctx context.Context, input *cloudwatchevents.ListEventBusesInput) (*cloudwatchevents.ListEventBusesOutput, error) {
	return a.client.ListEventBusesWithContext(ctx, input)
}

func (a *CloudWatchEventsActivities) ListEventSources(ctx context.Context, input *cloudwatchevents.ListEventSourcesInput) (*cloudwatchevents.ListEventSourcesOutput, error) {
	return a.client.ListEventSourcesWithContext(ctx, input)
}

func (a *CloudWatchEventsActivities) ListPartnerEventSourceAccounts(ctx context.Context, input *cloudwatchevents.ListPartnerEventSourceAccountsInput) (*cloudwatchevents.ListPartnerEventSourceAccountsOutput, error) {
	return a.client.ListPartnerEventSourceAccountsWithContext(ctx, input)
}

func (a *CloudWatchEventsActivities) ListPartnerEventSources(ctx context.Context, input *cloudwatchevents.ListPartnerEventSourcesInput) (*cloudwatchevents.ListPartnerEventSourcesOutput, error) {
	return a.client.ListPartnerEventSourcesWithContext(ctx, input)
}

func (a *CloudWatchEventsActivities) ListRuleNamesByTarget(ctx context.Context, input *cloudwatchevents.ListRuleNamesByTargetInput) (*cloudwatchevents.ListRuleNamesByTargetOutput, error) {
	return a.client.ListRuleNamesByTargetWithContext(ctx, input)
}

func (a *CloudWatchEventsActivities) ListRules(ctx context.Context, input *cloudwatchevents.ListRulesInput) (*cloudwatchevents.ListRulesOutput, error) {
	return a.client.ListRulesWithContext(ctx, input)
}

func (a *CloudWatchEventsActivities) ListTagsForResource(ctx context.Context, input *cloudwatchevents.ListTagsForResourceInput) (*cloudwatchevents.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *CloudWatchEventsActivities) ListTargetsByRule(ctx context.Context, input *cloudwatchevents.ListTargetsByRuleInput) (*cloudwatchevents.ListTargetsByRuleOutput, error) {
	return a.client.ListTargetsByRuleWithContext(ctx, input)
}

func (a *CloudWatchEventsActivities) PutEvents(ctx context.Context, input *cloudwatchevents.PutEventsInput) (*cloudwatchevents.PutEventsOutput, error) {
	return a.client.PutEventsWithContext(ctx, input)
}

func (a *CloudWatchEventsActivities) PutPartnerEvents(ctx context.Context, input *cloudwatchevents.PutPartnerEventsInput) (*cloudwatchevents.PutPartnerEventsOutput, error) {
	return a.client.PutPartnerEventsWithContext(ctx, input)
}

func (a *CloudWatchEventsActivities) PutPermission(ctx context.Context, input *cloudwatchevents.PutPermissionInput) (*cloudwatchevents.PutPermissionOutput, error) {
	return a.client.PutPermissionWithContext(ctx, input)
}

func (a *CloudWatchEventsActivities) PutRule(ctx context.Context, input *cloudwatchevents.PutRuleInput) (*cloudwatchevents.PutRuleOutput, error) {
	return a.client.PutRuleWithContext(ctx, input)
}

func (a *CloudWatchEventsActivities) PutTargets(ctx context.Context, input *cloudwatchevents.PutTargetsInput) (*cloudwatchevents.PutTargetsOutput, error) {
	return a.client.PutTargetsWithContext(ctx, input)
}

func (a *CloudWatchEventsActivities) RemovePermission(ctx context.Context, input *cloudwatchevents.RemovePermissionInput) (*cloudwatchevents.RemovePermissionOutput, error) {
	return a.client.RemovePermissionWithContext(ctx, input)
}

func (a *CloudWatchEventsActivities) RemoveTargets(ctx context.Context, input *cloudwatchevents.RemoveTargetsInput) (*cloudwatchevents.RemoveTargetsOutput, error) {
	return a.client.RemoveTargetsWithContext(ctx, input)
}

func (a *CloudWatchEventsActivities) TagResource(ctx context.Context, input *cloudwatchevents.TagResourceInput) (*cloudwatchevents.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *CloudWatchEventsActivities) TestEventPattern(ctx context.Context, input *cloudwatchevents.TestEventPatternInput) (*cloudwatchevents.TestEventPatternOutput, error) {
	return a.client.TestEventPatternWithContext(ctx, input)
}

func (a *CloudWatchEventsActivities) UntagResource(ctx context.Context, input *cloudwatchevents.UntagResourceInput) (*cloudwatchevents.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}
