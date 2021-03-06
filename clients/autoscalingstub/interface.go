// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package autoscalingstub

import (
	"github.com/aws/aws-sdk-go/service/autoscaling"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	AttachInstances(ctx workflow.Context, input *autoscaling.AttachInstancesInput) (*autoscaling.AttachInstancesOutput, error)
	AttachInstancesAsync(ctx workflow.Context, input *autoscaling.AttachInstancesInput) *AttachInstancesFuture

	AttachLoadBalancerTargetGroups(ctx workflow.Context, input *autoscaling.AttachLoadBalancerTargetGroupsInput) (*autoscaling.AttachLoadBalancerTargetGroupsOutput, error)
	AttachLoadBalancerTargetGroupsAsync(ctx workflow.Context, input *autoscaling.AttachLoadBalancerTargetGroupsInput) *AttachLoadBalancerTargetGroupsFuture

	AttachLoadBalancers(ctx workflow.Context, input *autoscaling.AttachLoadBalancersInput) (*autoscaling.AttachLoadBalancersOutput, error)
	AttachLoadBalancersAsync(ctx workflow.Context, input *autoscaling.AttachLoadBalancersInput) *AttachLoadBalancersFuture

	BatchDeleteScheduledAction(ctx workflow.Context, input *autoscaling.BatchDeleteScheduledActionInput) (*autoscaling.BatchDeleteScheduledActionOutput, error)
	BatchDeleteScheduledActionAsync(ctx workflow.Context, input *autoscaling.BatchDeleteScheduledActionInput) *BatchDeleteScheduledActionFuture

	BatchPutScheduledUpdateGroupAction(ctx workflow.Context, input *autoscaling.BatchPutScheduledUpdateGroupActionInput) (*autoscaling.BatchPutScheduledUpdateGroupActionOutput, error)
	BatchPutScheduledUpdateGroupActionAsync(ctx workflow.Context, input *autoscaling.BatchPutScheduledUpdateGroupActionInput) *BatchPutScheduledUpdateGroupActionFuture

	CancelInstanceRefresh(ctx workflow.Context, input *autoscaling.CancelInstanceRefreshInput) (*autoscaling.CancelInstanceRefreshOutput, error)
	CancelInstanceRefreshAsync(ctx workflow.Context, input *autoscaling.CancelInstanceRefreshInput) *CancelInstanceRefreshFuture

	CompleteLifecycleAction(ctx workflow.Context, input *autoscaling.CompleteLifecycleActionInput) (*autoscaling.CompleteLifecycleActionOutput, error)
	CompleteLifecycleActionAsync(ctx workflow.Context, input *autoscaling.CompleteLifecycleActionInput) *CompleteLifecycleActionFuture

	CreateAutoScalingGroup(ctx workflow.Context, input *autoscaling.CreateAutoScalingGroupInput) (*autoscaling.CreateAutoScalingGroupOutput, error)
	CreateAutoScalingGroupAsync(ctx workflow.Context, input *autoscaling.CreateAutoScalingGroupInput) *CreateAutoScalingGroupFuture

	CreateLaunchConfiguration(ctx workflow.Context, input *autoscaling.CreateLaunchConfigurationInput) (*autoscaling.CreateLaunchConfigurationOutput, error)
	CreateLaunchConfigurationAsync(ctx workflow.Context, input *autoscaling.CreateLaunchConfigurationInput) *CreateLaunchConfigurationFuture

	CreateOrUpdateTags(ctx workflow.Context, input *autoscaling.CreateOrUpdateTagsInput) (*autoscaling.CreateOrUpdateTagsOutput, error)
	CreateOrUpdateTagsAsync(ctx workflow.Context, input *autoscaling.CreateOrUpdateTagsInput) *CreateOrUpdateTagsFuture

	DeleteAutoScalingGroup(ctx workflow.Context, input *autoscaling.DeleteAutoScalingGroupInput) (*autoscaling.DeleteAutoScalingGroupOutput, error)
	DeleteAutoScalingGroupAsync(ctx workflow.Context, input *autoscaling.DeleteAutoScalingGroupInput) *DeleteAutoScalingGroupFuture

	DeleteLaunchConfiguration(ctx workflow.Context, input *autoscaling.DeleteLaunchConfigurationInput) (*autoscaling.DeleteLaunchConfigurationOutput, error)
	DeleteLaunchConfigurationAsync(ctx workflow.Context, input *autoscaling.DeleteLaunchConfigurationInput) *DeleteLaunchConfigurationFuture

	DeleteLifecycleHook(ctx workflow.Context, input *autoscaling.DeleteLifecycleHookInput) (*autoscaling.DeleteLifecycleHookOutput, error)
	DeleteLifecycleHookAsync(ctx workflow.Context, input *autoscaling.DeleteLifecycleHookInput) *DeleteLifecycleHookFuture

	DeleteNotificationConfiguration(ctx workflow.Context, input *autoscaling.DeleteNotificationConfigurationInput) (*autoscaling.DeleteNotificationConfigurationOutput, error)
	DeleteNotificationConfigurationAsync(ctx workflow.Context, input *autoscaling.DeleteNotificationConfigurationInput) *DeleteNotificationConfigurationFuture

	DeletePolicy(ctx workflow.Context, input *autoscaling.DeletePolicyInput) (*autoscaling.DeletePolicyOutput, error)
	DeletePolicyAsync(ctx workflow.Context, input *autoscaling.DeletePolicyInput) *DeletePolicyFuture

	DeleteScheduledAction(ctx workflow.Context, input *autoscaling.DeleteScheduledActionInput) (*autoscaling.DeleteScheduledActionOutput, error)
	DeleteScheduledActionAsync(ctx workflow.Context, input *autoscaling.DeleteScheduledActionInput) *DeleteScheduledActionFuture

	DeleteTags(ctx workflow.Context, input *autoscaling.DeleteTagsInput) (*autoscaling.DeleteTagsOutput, error)
	DeleteTagsAsync(ctx workflow.Context, input *autoscaling.DeleteTagsInput) *DeleteTagsFuture

	DescribeAccountLimits(ctx workflow.Context, input *autoscaling.DescribeAccountLimitsInput) (*autoscaling.DescribeAccountLimitsOutput, error)
	DescribeAccountLimitsAsync(ctx workflow.Context, input *autoscaling.DescribeAccountLimitsInput) *DescribeAccountLimitsFuture

	DescribeAdjustmentTypes(ctx workflow.Context, input *autoscaling.DescribeAdjustmentTypesInput) (*autoscaling.DescribeAdjustmentTypesOutput, error)
	DescribeAdjustmentTypesAsync(ctx workflow.Context, input *autoscaling.DescribeAdjustmentTypesInput) *DescribeAdjustmentTypesFuture

	DescribeAutoScalingGroups(ctx workflow.Context, input *autoscaling.DescribeAutoScalingGroupsInput) (*autoscaling.DescribeAutoScalingGroupsOutput, error)
	DescribeAutoScalingGroupsAsync(ctx workflow.Context, input *autoscaling.DescribeAutoScalingGroupsInput) *DescribeAutoScalingGroupsFuture

	DescribeAutoScalingInstances(ctx workflow.Context, input *autoscaling.DescribeAutoScalingInstancesInput) (*autoscaling.DescribeAutoScalingInstancesOutput, error)
	DescribeAutoScalingInstancesAsync(ctx workflow.Context, input *autoscaling.DescribeAutoScalingInstancesInput) *DescribeAutoScalingInstancesFuture

	DescribeAutoScalingNotificationTypes(ctx workflow.Context, input *autoscaling.DescribeAutoScalingNotificationTypesInput) (*autoscaling.DescribeAutoScalingNotificationTypesOutput, error)
	DescribeAutoScalingNotificationTypesAsync(ctx workflow.Context, input *autoscaling.DescribeAutoScalingNotificationTypesInput) *DescribeAutoScalingNotificationTypesFuture

	DescribeInstanceRefreshes(ctx workflow.Context, input *autoscaling.DescribeInstanceRefreshesInput) (*autoscaling.DescribeInstanceRefreshesOutput, error)
	DescribeInstanceRefreshesAsync(ctx workflow.Context, input *autoscaling.DescribeInstanceRefreshesInput) *DescribeInstanceRefreshesFuture

	DescribeLaunchConfigurations(ctx workflow.Context, input *autoscaling.DescribeLaunchConfigurationsInput) (*autoscaling.DescribeLaunchConfigurationsOutput, error)
	DescribeLaunchConfigurationsAsync(ctx workflow.Context, input *autoscaling.DescribeLaunchConfigurationsInput) *DescribeLaunchConfigurationsFuture

	DescribeLifecycleHookTypes(ctx workflow.Context, input *autoscaling.DescribeLifecycleHookTypesInput) (*autoscaling.DescribeLifecycleHookTypesOutput, error)
	DescribeLifecycleHookTypesAsync(ctx workflow.Context, input *autoscaling.DescribeLifecycleHookTypesInput) *DescribeLifecycleHookTypesFuture

	DescribeLifecycleHooks(ctx workflow.Context, input *autoscaling.DescribeLifecycleHooksInput) (*autoscaling.DescribeLifecycleHooksOutput, error)
	DescribeLifecycleHooksAsync(ctx workflow.Context, input *autoscaling.DescribeLifecycleHooksInput) *DescribeLifecycleHooksFuture

	DescribeLoadBalancerTargetGroups(ctx workflow.Context, input *autoscaling.DescribeLoadBalancerTargetGroupsInput) (*autoscaling.DescribeLoadBalancerTargetGroupsOutput, error)
	DescribeLoadBalancerTargetGroupsAsync(ctx workflow.Context, input *autoscaling.DescribeLoadBalancerTargetGroupsInput) *DescribeLoadBalancerTargetGroupsFuture

	DescribeLoadBalancers(ctx workflow.Context, input *autoscaling.DescribeLoadBalancersInput) (*autoscaling.DescribeLoadBalancersOutput, error)
	DescribeLoadBalancersAsync(ctx workflow.Context, input *autoscaling.DescribeLoadBalancersInput) *DescribeLoadBalancersFuture

	DescribeMetricCollectionTypes(ctx workflow.Context, input *autoscaling.DescribeMetricCollectionTypesInput) (*autoscaling.DescribeMetricCollectionTypesOutput, error)
	DescribeMetricCollectionTypesAsync(ctx workflow.Context, input *autoscaling.DescribeMetricCollectionTypesInput) *DescribeMetricCollectionTypesFuture

	DescribeNotificationConfigurations(ctx workflow.Context, input *autoscaling.DescribeNotificationConfigurationsInput) (*autoscaling.DescribeNotificationConfigurationsOutput, error)
	DescribeNotificationConfigurationsAsync(ctx workflow.Context, input *autoscaling.DescribeNotificationConfigurationsInput) *DescribeNotificationConfigurationsFuture

	DescribePolicies(ctx workflow.Context, input *autoscaling.DescribePoliciesInput) (*autoscaling.DescribePoliciesOutput, error)
	DescribePoliciesAsync(ctx workflow.Context, input *autoscaling.DescribePoliciesInput) *DescribePoliciesFuture

	DescribeScalingActivities(ctx workflow.Context, input *autoscaling.DescribeScalingActivitiesInput) (*autoscaling.DescribeScalingActivitiesOutput, error)
	DescribeScalingActivitiesAsync(ctx workflow.Context, input *autoscaling.DescribeScalingActivitiesInput) *DescribeScalingActivitiesFuture

	DescribeScalingProcessTypes(ctx workflow.Context, input *autoscaling.DescribeScalingProcessTypesInput) (*autoscaling.DescribeScalingProcessTypesOutput, error)
	DescribeScalingProcessTypesAsync(ctx workflow.Context, input *autoscaling.DescribeScalingProcessTypesInput) *DescribeScalingProcessTypesFuture

	DescribeScheduledActions(ctx workflow.Context, input *autoscaling.DescribeScheduledActionsInput) (*autoscaling.DescribeScheduledActionsOutput, error)
	DescribeScheduledActionsAsync(ctx workflow.Context, input *autoscaling.DescribeScheduledActionsInput) *DescribeScheduledActionsFuture

	DescribeTags(ctx workflow.Context, input *autoscaling.DescribeTagsInput) (*autoscaling.DescribeTagsOutput, error)
	DescribeTagsAsync(ctx workflow.Context, input *autoscaling.DescribeTagsInput) *DescribeTagsFuture

	DescribeTerminationPolicyTypes(ctx workflow.Context, input *autoscaling.DescribeTerminationPolicyTypesInput) (*autoscaling.DescribeTerminationPolicyTypesOutput, error)
	DescribeTerminationPolicyTypesAsync(ctx workflow.Context, input *autoscaling.DescribeTerminationPolicyTypesInput) *DescribeTerminationPolicyTypesFuture

	DetachInstances(ctx workflow.Context, input *autoscaling.DetachInstancesInput) (*autoscaling.DetachInstancesOutput, error)
	DetachInstancesAsync(ctx workflow.Context, input *autoscaling.DetachInstancesInput) *DetachInstancesFuture

	DetachLoadBalancerTargetGroups(ctx workflow.Context, input *autoscaling.DetachLoadBalancerTargetGroupsInput) (*autoscaling.DetachLoadBalancerTargetGroupsOutput, error)
	DetachLoadBalancerTargetGroupsAsync(ctx workflow.Context, input *autoscaling.DetachLoadBalancerTargetGroupsInput) *DetachLoadBalancerTargetGroupsFuture

	DetachLoadBalancers(ctx workflow.Context, input *autoscaling.DetachLoadBalancersInput) (*autoscaling.DetachLoadBalancersOutput, error)
	DetachLoadBalancersAsync(ctx workflow.Context, input *autoscaling.DetachLoadBalancersInput) *DetachLoadBalancersFuture

	DisableMetricsCollection(ctx workflow.Context, input *autoscaling.DisableMetricsCollectionInput) (*autoscaling.DisableMetricsCollectionOutput, error)
	DisableMetricsCollectionAsync(ctx workflow.Context, input *autoscaling.DisableMetricsCollectionInput) *DisableMetricsCollectionFuture

	EnableMetricsCollection(ctx workflow.Context, input *autoscaling.EnableMetricsCollectionInput) (*autoscaling.EnableMetricsCollectionOutput, error)
	EnableMetricsCollectionAsync(ctx workflow.Context, input *autoscaling.EnableMetricsCollectionInput) *EnableMetricsCollectionFuture

	EnterStandby(ctx workflow.Context, input *autoscaling.EnterStandbyInput) (*autoscaling.EnterStandbyOutput, error)
	EnterStandbyAsync(ctx workflow.Context, input *autoscaling.EnterStandbyInput) *EnterStandbyFuture

	ExecutePolicy(ctx workflow.Context, input *autoscaling.ExecutePolicyInput) (*autoscaling.ExecutePolicyOutput, error)
	ExecutePolicyAsync(ctx workflow.Context, input *autoscaling.ExecutePolicyInput) *ExecutePolicyFuture

	ExitStandby(ctx workflow.Context, input *autoscaling.ExitStandbyInput) (*autoscaling.ExitStandbyOutput, error)
	ExitStandbyAsync(ctx workflow.Context, input *autoscaling.ExitStandbyInput) *ExitStandbyFuture

	PutLifecycleHook(ctx workflow.Context, input *autoscaling.PutLifecycleHookInput) (*autoscaling.PutLifecycleHookOutput, error)
	PutLifecycleHookAsync(ctx workflow.Context, input *autoscaling.PutLifecycleHookInput) *PutLifecycleHookFuture

	PutNotificationConfiguration(ctx workflow.Context, input *autoscaling.PutNotificationConfigurationInput) (*autoscaling.PutNotificationConfigurationOutput, error)
	PutNotificationConfigurationAsync(ctx workflow.Context, input *autoscaling.PutNotificationConfigurationInput) *PutNotificationConfigurationFuture

	PutScalingPolicy(ctx workflow.Context, input *autoscaling.PutScalingPolicyInput) (*autoscaling.PutScalingPolicyOutput, error)
	PutScalingPolicyAsync(ctx workflow.Context, input *autoscaling.PutScalingPolicyInput) *PutScalingPolicyFuture

	PutScheduledUpdateGroupAction(ctx workflow.Context, input *autoscaling.PutScheduledUpdateGroupActionInput) (*autoscaling.PutScheduledUpdateGroupActionOutput, error)
	PutScheduledUpdateGroupActionAsync(ctx workflow.Context, input *autoscaling.PutScheduledUpdateGroupActionInput) *PutScheduledUpdateGroupActionFuture

	RecordLifecycleActionHeartbeat(ctx workflow.Context, input *autoscaling.RecordLifecycleActionHeartbeatInput) (*autoscaling.RecordLifecycleActionHeartbeatOutput, error)
	RecordLifecycleActionHeartbeatAsync(ctx workflow.Context, input *autoscaling.RecordLifecycleActionHeartbeatInput) *RecordLifecycleActionHeartbeatFuture

	ResumeProcesses(ctx workflow.Context, input *autoscaling.ScalingProcessQuery) (*autoscaling.ResumeProcessesOutput, error)
	ResumeProcessesAsync(ctx workflow.Context, input *autoscaling.ScalingProcessQuery) *ResumeProcessesFuture

	SetDesiredCapacity(ctx workflow.Context, input *autoscaling.SetDesiredCapacityInput) (*autoscaling.SetDesiredCapacityOutput, error)
	SetDesiredCapacityAsync(ctx workflow.Context, input *autoscaling.SetDesiredCapacityInput) *SetDesiredCapacityFuture

	SetInstanceHealth(ctx workflow.Context, input *autoscaling.SetInstanceHealthInput) (*autoscaling.SetInstanceHealthOutput, error)
	SetInstanceHealthAsync(ctx workflow.Context, input *autoscaling.SetInstanceHealthInput) *SetInstanceHealthFuture

	SetInstanceProtection(ctx workflow.Context, input *autoscaling.SetInstanceProtectionInput) (*autoscaling.SetInstanceProtectionOutput, error)
	SetInstanceProtectionAsync(ctx workflow.Context, input *autoscaling.SetInstanceProtectionInput) *SetInstanceProtectionFuture

	StartInstanceRefresh(ctx workflow.Context, input *autoscaling.StartInstanceRefreshInput) (*autoscaling.StartInstanceRefreshOutput, error)
	StartInstanceRefreshAsync(ctx workflow.Context, input *autoscaling.StartInstanceRefreshInput) *StartInstanceRefreshFuture

	SuspendProcesses(ctx workflow.Context, input *autoscaling.ScalingProcessQuery) (*autoscaling.SuspendProcessesOutput, error)
	SuspendProcessesAsync(ctx workflow.Context, input *autoscaling.ScalingProcessQuery) *SuspendProcessesFuture

	TerminateInstanceInAutoScalingGroup(ctx workflow.Context, input *autoscaling.TerminateInstanceInAutoScalingGroupInput) (*autoscaling.TerminateInstanceInAutoScalingGroupOutput, error)
	TerminateInstanceInAutoScalingGroupAsync(ctx workflow.Context, input *autoscaling.TerminateInstanceInAutoScalingGroupInput) *TerminateInstanceInAutoScalingGroupFuture

	UpdateAutoScalingGroup(ctx workflow.Context, input *autoscaling.UpdateAutoScalingGroupInput) (*autoscaling.UpdateAutoScalingGroupOutput, error)
	UpdateAutoScalingGroupAsync(ctx workflow.Context, input *autoscaling.UpdateAutoScalingGroupInput) *UpdateAutoScalingGroupFuture

	WaitUntilGroupExists(ctx workflow.Context, input *autoscaling.DescribeAutoScalingGroupsInput) error
	WaitUntilGroupExistsAsync(ctx workflow.Context, input *autoscaling.DescribeAutoScalingGroupsInput) *clients.VoidFuture

	WaitUntilGroupInService(ctx workflow.Context, input *autoscaling.DescribeAutoScalingGroupsInput) error
	WaitUntilGroupInServiceAsync(ctx workflow.Context, input *autoscaling.DescribeAutoScalingGroupsInput) *clients.VoidFuture

	WaitUntilGroupNotExists(ctx workflow.Context, input *autoscaling.DescribeAutoScalingGroupsInput) error
	WaitUntilGroupNotExistsAsync(ctx workflow.Context, input *autoscaling.DescribeAutoScalingGroupsInput) *clients.VoidFuture
}

func NewClient() Client {
	return &stub{}
}
