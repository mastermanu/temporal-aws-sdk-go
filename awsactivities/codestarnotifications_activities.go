package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/codestarnotifications"
	"github.com/aws/aws-sdk-go/service/codestarnotifications/codestarnotificationsiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type CodeStarNotificationsActivities struct {
	client codestarnotificationsiface.CodeStarNotificationsAPI
}

func NewCodeStarNotificationsActivities(session *session.Session, config ...*aws.Config) *CodeStarNotificationsActivities {
	client := codestarnotifications.New(session, config...)
	return &CodeStarNotificationsActivities{client: client}
}

func (a *CodeStarNotificationsActivities) CreateNotificationRule(ctx context.Context, input *codestarnotifications.CreateNotificationRuleInput) (*codestarnotifications.CreateNotificationRuleOutput, error) {
	return a.client.CreateNotificationRuleWithContext(ctx, input)
}

func (a *CodeStarNotificationsActivities) DeleteNotificationRule(ctx context.Context, input *codestarnotifications.DeleteNotificationRuleInput) (*codestarnotifications.DeleteNotificationRuleOutput, error) {
	return a.client.DeleteNotificationRuleWithContext(ctx, input)
}

func (a *CodeStarNotificationsActivities) DeleteTarget(ctx context.Context, input *codestarnotifications.DeleteTargetInput) (*codestarnotifications.DeleteTargetOutput, error) {
	return a.client.DeleteTargetWithContext(ctx, input)
}

func (a *CodeStarNotificationsActivities) DescribeNotificationRule(ctx context.Context, input *codestarnotifications.DescribeNotificationRuleInput) (*codestarnotifications.DescribeNotificationRuleOutput, error) {
	return a.client.DescribeNotificationRuleWithContext(ctx, input)
}

func (a *CodeStarNotificationsActivities) ListEventTypes(ctx context.Context, input *codestarnotifications.ListEventTypesInput) (*codestarnotifications.ListEventTypesOutput, error) {
	return a.client.ListEventTypesWithContext(ctx, input)
}

func (a *CodeStarNotificationsActivities) ListNotificationRules(ctx context.Context, input *codestarnotifications.ListNotificationRulesInput) (*codestarnotifications.ListNotificationRulesOutput, error) {
	return a.client.ListNotificationRulesWithContext(ctx, input)
}

func (a *CodeStarNotificationsActivities) ListTagsForResource(ctx context.Context, input *codestarnotifications.ListTagsForResourceInput) (*codestarnotifications.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *CodeStarNotificationsActivities) ListTargets(ctx context.Context, input *codestarnotifications.ListTargetsInput) (*codestarnotifications.ListTargetsOutput, error) {
	return a.client.ListTargetsWithContext(ctx, input)
}

func (a *CodeStarNotificationsActivities) Subscribe(ctx context.Context, input *codestarnotifications.SubscribeInput) (*codestarnotifications.SubscribeOutput, error) {
	return a.client.SubscribeWithContext(ctx, input)
}

func (a *CodeStarNotificationsActivities) TagResource(ctx context.Context, input *codestarnotifications.TagResourceInput) (*codestarnotifications.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *CodeStarNotificationsActivities) Unsubscribe(ctx context.Context, input *codestarnotifications.UnsubscribeInput) (*codestarnotifications.UnsubscribeOutput, error) {
	return a.client.UnsubscribeWithContext(ctx, input)
}

func (a *CodeStarNotificationsActivities) UntagResource(ctx context.Context, input *codestarnotifications.UntagResourceInput) (*codestarnotifications.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}

func (a *CodeStarNotificationsActivities) UpdateNotificationRule(ctx context.Context, input *codestarnotifications.UpdateNotificationRuleInput) (*codestarnotifications.UpdateNotificationRuleOutput, error) {
	return a.client.UpdateNotificationRuleWithContext(ctx, input)
}
