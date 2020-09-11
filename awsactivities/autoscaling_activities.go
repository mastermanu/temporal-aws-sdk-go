package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/autoscaling"
	"github.com/aws/aws-sdk-go/service/autoscaling/autoscalingiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type AutoScalingActivities struct {
	client autoscalingiface.AutoScalingAPI
}

func NewAutoScalingActivities(session *session.Session, config ...*aws.Config) *AutoScalingActivities {
	client := autoscaling.New(session, config...)
	return &AutoScalingActivities{client: client}
}

func (a *AutoScalingActivities) AttachInstances(ctx context.Context, input *autoscaling.AttachInstancesInput) (*autoscaling.AttachInstancesOutput, error) {
	return a.client.AttachInstancesWithContext(ctx, input)
}

func (a *AutoScalingActivities) AttachLoadBalancerTargetGroups(ctx context.Context, input *autoscaling.AttachLoadBalancerTargetGroupsInput) (*autoscaling.AttachLoadBalancerTargetGroupsOutput, error) {
	return a.client.AttachLoadBalancerTargetGroupsWithContext(ctx, input)
}

func (a *AutoScalingActivities) AttachLoadBalancers(ctx context.Context, input *autoscaling.AttachLoadBalancersInput) (*autoscaling.AttachLoadBalancersOutput, error) {
	return a.client.AttachLoadBalancersWithContext(ctx, input)
}

func (a *AutoScalingActivities) BatchDeleteScheduledAction(ctx context.Context, input *autoscaling.BatchDeleteScheduledActionInput) (*autoscaling.BatchDeleteScheduledActionOutput, error) {
	return a.client.BatchDeleteScheduledActionWithContext(ctx, input)
}

func (a *AutoScalingActivities) BatchPutScheduledUpdateGroupAction(ctx context.Context, input *autoscaling.BatchPutScheduledUpdateGroupActionInput) (*autoscaling.BatchPutScheduledUpdateGroupActionOutput, error) {
	return a.client.BatchPutScheduledUpdateGroupActionWithContext(ctx, input)
}

func (a *AutoScalingActivities) CancelInstanceRefresh(ctx context.Context, input *autoscaling.CancelInstanceRefreshInput) (*autoscaling.CancelInstanceRefreshOutput, error) {
	return a.client.CancelInstanceRefreshWithContext(ctx, input)
}

func (a *AutoScalingActivities) CompleteLifecycleAction(ctx context.Context, input *autoscaling.CompleteLifecycleActionInput) (*autoscaling.CompleteLifecycleActionOutput, error) {
	return a.client.CompleteLifecycleActionWithContext(ctx, input)
}

func (a *AutoScalingActivities) CreateAutoScalingGroup(ctx context.Context, input *autoscaling.CreateAutoScalingGroupInput) (*autoscaling.CreateAutoScalingGroupOutput, error) {
	return a.client.CreateAutoScalingGroupWithContext(ctx, input)
}

func (a *AutoScalingActivities) CreateLaunchConfiguration(ctx context.Context, input *autoscaling.CreateLaunchConfigurationInput) (*autoscaling.CreateLaunchConfigurationOutput, error) {
	return a.client.CreateLaunchConfigurationWithContext(ctx, input)
}

func (a *AutoScalingActivities) CreateOrUpdateTags(ctx context.Context, input *autoscaling.CreateOrUpdateTagsInput) (*autoscaling.CreateOrUpdateTagsOutput, error) {
	return a.client.CreateOrUpdateTagsWithContext(ctx, input)
}

func (a *AutoScalingActivities) DeleteAutoScalingGroup(ctx context.Context, input *autoscaling.DeleteAutoScalingGroupInput) (*autoscaling.DeleteAutoScalingGroupOutput, error) {
	return a.client.DeleteAutoScalingGroupWithContext(ctx, input)
}

func (a *AutoScalingActivities) DeleteLaunchConfiguration(ctx context.Context, input *autoscaling.DeleteLaunchConfigurationInput) (*autoscaling.DeleteLaunchConfigurationOutput, error) {
	return a.client.DeleteLaunchConfigurationWithContext(ctx, input)
}

func (a *AutoScalingActivities) DeleteLifecycleHook(ctx context.Context, input *autoscaling.DeleteLifecycleHookInput) (*autoscaling.DeleteLifecycleHookOutput, error) {
	return a.client.DeleteLifecycleHookWithContext(ctx, input)
}

func (a *AutoScalingActivities) DeleteNotificationConfiguration(ctx context.Context, input *autoscaling.DeleteNotificationConfigurationInput) (*autoscaling.DeleteNotificationConfigurationOutput, error) {
	return a.client.DeleteNotificationConfigurationWithContext(ctx, input)
}

func (a *AutoScalingActivities) DeletePolicy(ctx context.Context, input *autoscaling.DeletePolicyInput) (*autoscaling.DeletePolicyOutput, error) {
	return a.client.DeletePolicyWithContext(ctx, input)
}

func (a *AutoScalingActivities) DeleteScheduledAction(ctx context.Context, input *autoscaling.DeleteScheduledActionInput) (*autoscaling.DeleteScheduledActionOutput, error) {
	return a.client.DeleteScheduledActionWithContext(ctx, input)
}

func (a *AutoScalingActivities) DeleteTags(ctx context.Context, input *autoscaling.DeleteTagsInput) (*autoscaling.DeleteTagsOutput, error) {
	return a.client.DeleteTagsWithContext(ctx, input)
}

func (a *AutoScalingActivities) DescribeAccountLimits(ctx context.Context, input *autoscaling.DescribeAccountLimitsInput) (*autoscaling.DescribeAccountLimitsOutput, error) {
	return a.client.DescribeAccountLimitsWithContext(ctx, input)
}

func (a *AutoScalingActivities) DescribeAdjustmentTypes(ctx context.Context, input *autoscaling.DescribeAdjustmentTypesInput) (*autoscaling.DescribeAdjustmentTypesOutput, error) {
	return a.client.DescribeAdjustmentTypesWithContext(ctx, input)
}

func (a *AutoScalingActivities) DescribeAutoScalingGroups(ctx context.Context, input *autoscaling.DescribeAutoScalingGroupsInput) (*autoscaling.DescribeAutoScalingGroupsOutput, error) {
	return a.client.DescribeAutoScalingGroupsWithContext(ctx, input)
}

func (a *AutoScalingActivities) DescribeAutoScalingInstances(ctx context.Context, input *autoscaling.DescribeAutoScalingInstancesInput) (*autoscaling.DescribeAutoScalingInstancesOutput, error) {
	return a.client.DescribeAutoScalingInstancesWithContext(ctx, input)
}

func (a *AutoScalingActivities) DescribeAutoScalingNotificationTypes(ctx context.Context, input *autoscaling.DescribeAutoScalingNotificationTypesInput) (*autoscaling.DescribeAutoScalingNotificationTypesOutput, error) {
	return a.client.DescribeAutoScalingNotificationTypesWithContext(ctx, input)
}

func (a *AutoScalingActivities) DescribeInstanceRefreshes(ctx context.Context, input *autoscaling.DescribeInstanceRefreshesInput) (*autoscaling.DescribeInstanceRefreshesOutput, error) {
	return a.client.DescribeInstanceRefreshesWithContext(ctx, input)
}

func (a *AutoScalingActivities) DescribeLaunchConfigurations(ctx context.Context, input *autoscaling.DescribeLaunchConfigurationsInput) (*autoscaling.DescribeLaunchConfigurationsOutput, error) {
	return a.client.DescribeLaunchConfigurationsWithContext(ctx, input)
}

func (a *AutoScalingActivities) DescribeLifecycleHookTypes(ctx context.Context, input *autoscaling.DescribeLifecycleHookTypesInput) (*autoscaling.DescribeLifecycleHookTypesOutput, error) {
	return a.client.DescribeLifecycleHookTypesWithContext(ctx, input)
}

func (a *AutoScalingActivities) DescribeLifecycleHooks(ctx context.Context, input *autoscaling.DescribeLifecycleHooksInput) (*autoscaling.DescribeLifecycleHooksOutput, error) {
	return a.client.DescribeLifecycleHooksWithContext(ctx, input)
}

func (a *AutoScalingActivities) DescribeLoadBalancerTargetGroups(ctx context.Context, input *autoscaling.DescribeLoadBalancerTargetGroupsInput) (*autoscaling.DescribeLoadBalancerTargetGroupsOutput, error) {
	return a.client.DescribeLoadBalancerTargetGroupsWithContext(ctx, input)
}

func (a *AutoScalingActivities) DescribeLoadBalancers(ctx context.Context, input *autoscaling.DescribeLoadBalancersInput) (*autoscaling.DescribeLoadBalancersOutput, error) {
	return a.client.DescribeLoadBalancersWithContext(ctx, input)
}

func (a *AutoScalingActivities) DescribeMetricCollectionTypes(ctx context.Context, input *autoscaling.DescribeMetricCollectionTypesInput) (*autoscaling.DescribeMetricCollectionTypesOutput, error) {
	return a.client.DescribeMetricCollectionTypesWithContext(ctx, input)
}

func (a *AutoScalingActivities) DescribeNotificationConfigurations(ctx context.Context, input *autoscaling.DescribeNotificationConfigurationsInput) (*autoscaling.DescribeNotificationConfigurationsOutput, error) {
	return a.client.DescribeNotificationConfigurationsWithContext(ctx, input)
}

func (a *AutoScalingActivities) DescribePolicies(ctx context.Context, input *autoscaling.DescribePoliciesInput) (*autoscaling.DescribePoliciesOutput, error) {
	return a.client.DescribePoliciesWithContext(ctx, input)
}

func (a *AutoScalingActivities) DescribeScalingActivities(ctx context.Context, input *autoscaling.DescribeScalingActivitiesInput) (*autoscaling.DescribeScalingActivitiesOutput, error) {
	return a.client.DescribeScalingActivitiesWithContext(ctx, input)
}

func (a *AutoScalingActivities) DescribeScalingProcessTypes(ctx context.Context, input *autoscaling.DescribeScalingProcessTypesInput) (*autoscaling.DescribeScalingProcessTypesOutput, error) {
	return a.client.DescribeScalingProcessTypesWithContext(ctx, input)
}

func (a *AutoScalingActivities) DescribeScheduledActions(ctx context.Context, input *autoscaling.DescribeScheduledActionsInput) (*autoscaling.DescribeScheduledActionsOutput, error) {
	return a.client.DescribeScheduledActionsWithContext(ctx, input)
}

func (a *AutoScalingActivities) DescribeTags(ctx context.Context, input *autoscaling.DescribeTagsInput) (*autoscaling.DescribeTagsOutput, error) {
	return a.client.DescribeTagsWithContext(ctx, input)
}

func (a *AutoScalingActivities) DescribeTerminationPolicyTypes(ctx context.Context, input *autoscaling.DescribeTerminationPolicyTypesInput) (*autoscaling.DescribeTerminationPolicyTypesOutput, error) {
	return a.client.DescribeTerminationPolicyTypesWithContext(ctx, input)
}

func (a *AutoScalingActivities) DetachInstances(ctx context.Context, input *autoscaling.DetachInstancesInput) (*autoscaling.DetachInstancesOutput, error) {
	return a.client.DetachInstancesWithContext(ctx, input)
}

func (a *AutoScalingActivities) DetachLoadBalancerTargetGroups(ctx context.Context, input *autoscaling.DetachLoadBalancerTargetGroupsInput) (*autoscaling.DetachLoadBalancerTargetGroupsOutput, error) {
	return a.client.DetachLoadBalancerTargetGroupsWithContext(ctx, input)
}

func (a *AutoScalingActivities) DetachLoadBalancers(ctx context.Context, input *autoscaling.DetachLoadBalancersInput) (*autoscaling.DetachLoadBalancersOutput, error) {
	return a.client.DetachLoadBalancersWithContext(ctx, input)
}

func (a *AutoScalingActivities) DisableMetricsCollection(ctx context.Context, input *autoscaling.DisableMetricsCollectionInput) (*autoscaling.DisableMetricsCollectionOutput, error) {
	return a.client.DisableMetricsCollectionWithContext(ctx, input)
}

func (a *AutoScalingActivities) EnableMetricsCollection(ctx context.Context, input *autoscaling.EnableMetricsCollectionInput) (*autoscaling.EnableMetricsCollectionOutput, error) {
	return a.client.EnableMetricsCollectionWithContext(ctx, input)
}

func (a *AutoScalingActivities) EnterStandby(ctx context.Context, input *autoscaling.EnterStandbyInput) (*autoscaling.EnterStandbyOutput, error) {
	return a.client.EnterStandbyWithContext(ctx, input)
}

func (a *AutoScalingActivities) ExecutePolicy(ctx context.Context, input *autoscaling.ExecutePolicyInput) (*autoscaling.ExecutePolicyOutput, error) {
	return a.client.ExecutePolicyWithContext(ctx, input)
}

func (a *AutoScalingActivities) ExitStandby(ctx context.Context, input *autoscaling.ExitStandbyInput) (*autoscaling.ExitStandbyOutput, error) {
	return a.client.ExitStandbyWithContext(ctx, input)
}

func (a *AutoScalingActivities) PutLifecycleHook(ctx context.Context, input *autoscaling.PutLifecycleHookInput) (*autoscaling.PutLifecycleHookOutput, error) {
	return a.client.PutLifecycleHookWithContext(ctx, input)
}

func (a *AutoScalingActivities) PutNotificationConfiguration(ctx context.Context, input *autoscaling.PutNotificationConfigurationInput) (*autoscaling.PutNotificationConfigurationOutput, error) {
	return a.client.PutNotificationConfigurationWithContext(ctx, input)
}

func (a *AutoScalingActivities) PutScalingPolicy(ctx context.Context, input *autoscaling.PutScalingPolicyInput) (*autoscaling.PutScalingPolicyOutput, error) {
	return a.client.PutScalingPolicyWithContext(ctx, input)
}

func (a *AutoScalingActivities) PutScheduledUpdateGroupAction(ctx context.Context, input *autoscaling.PutScheduledUpdateGroupActionInput) (*autoscaling.PutScheduledUpdateGroupActionOutput, error) {
	return a.client.PutScheduledUpdateGroupActionWithContext(ctx, input)
}

func (a *AutoScalingActivities) RecordLifecycleActionHeartbeat(ctx context.Context, input *autoscaling.RecordLifecycleActionHeartbeatInput) (*autoscaling.RecordLifecycleActionHeartbeatOutput, error) {
	return a.client.RecordLifecycleActionHeartbeatWithContext(ctx, input)
}

func (a *AutoScalingActivities) ResumeProcesses(ctx context.Context, input *autoscaling.ScalingProcessQuery) (*autoscaling.ResumeProcessesOutput, error) {
	return a.client.ResumeProcessesWithContext(ctx, input)
}

func (a *AutoScalingActivities) SetDesiredCapacity(ctx context.Context, input *autoscaling.SetDesiredCapacityInput) (*autoscaling.SetDesiredCapacityOutput, error) {
	return a.client.SetDesiredCapacityWithContext(ctx, input)
}

func (a *AutoScalingActivities) SetInstanceHealth(ctx context.Context, input *autoscaling.SetInstanceHealthInput) (*autoscaling.SetInstanceHealthOutput, error) {
	return a.client.SetInstanceHealthWithContext(ctx, input)
}

func (a *AutoScalingActivities) SetInstanceProtection(ctx context.Context, input *autoscaling.SetInstanceProtectionInput) (*autoscaling.SetInstanceProtectionOutput, error) {
	return a.client.SetInstanceProtectionWithContext(ctx, input)
}

func (a *AutoScalingActivities) StartInstanceRefresh(ctx context.Context, input *autoscaling.StartInstanceRefreshInput) (*autoscaling.StartInstanceRefreshOutput, error) {
	return a.client.StartInstanceRefreshWithContext(ctx, input)
}

func (a *AutoScalingActivities) SuspendProcesses(ctx context.Context, input *autoscaling.ScalingProcessQuery) (*autoscaling.SuspendProcessesOutput, error) {
	return a.client.SuspendProcessesWithContext(ctx, input)
}

func (a *AutoScalingActivities) TerminateInstanceInAutoScalingGroup(ctx context.Context, input *autoscaling.TerminateInstanceInAutoScalingGroupInput) (*autoscaling.TerminateInstanceInAutoScalingGroupOutput, error) {
	return a.client.TerminateInstanceInAutoScalingGroupWithContext(ctx, input)
}

func (a *AutoScalingActivities) UpdateAutoScalingGroup(ctx context.Context, input *autoscaling.UpdateAutoScalingGroupInput) (*autoscaling.UpdateAutoScalingGroupOutput, error) {
	return a.client.UpdateAutoScalingGroupWithContext(ctx, input)
}

func (a *AutoScalingActivities) WaitUntilGroupExists(ctx context.Context, input *autoscaling.DescribeAutoScalingGroupsInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilGroupExistsWithContext(ctx, input, options...)
	})
}

func (a *AutoScalingActivities) WaitUntilGroupInService(ctx context.Context, input *autoscaling.DescribeAutoScalingGroupsInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilGroupInServiceWithContext(ctx, input, options...)
	})
}

func (a *AutoScalingActivities) WaitUntilGroupNotExists(ctx context.Context, input *autoscaling.DescribeAutoScalingGroupsInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilGroupNotExistsWithContext(ctx, input, options...)
	})
}
