// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package ecsstub

import (
	"github.com/aws/aws-sdk-go/service/ecs"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	CreateCapacityProvider(ctx workflow.Context, input *ecs.CreateCapacityProviderInput) (*ecs.CreateCapacityProviderOutput, error)
	CreateCapacityProviderAsync(ctx workflow.Context, input *ecs.CreateCapacityProviderInput) *CreateCapacityProviderFuture

	CreateCluster(ctx workflow.Context, input *ecs.CreateClusterInput) (*ecs.CreateClusterOutput, error)
	CreateClusterAsync(ctx workflow.Context, input *ecs.CreateClusterInput) *CreateClusterFuture

	CreateService(ctx workflow.Context, input *ecs.CreateServiceInput) (*ecs.CreateServiceOutput, error)
	CreateServiceAsync(ctx workflow.Context, input *ecs.CreateServiceInput) *CreateServiceFuture

	CreateTaskSet(ctx workflow.Context, input *ecs.CreateTaskSetInput) (*ecs.CreateTaskSetOutput, error)
	CreateTaskSetAsync(ctx workflow.Context, input *ecs.CreateTaskSetInput) *CreateTaskSetFuture

	DeleteAccountSetting(ctx workflow.Context, input *ecs.DeleteAccountSettingInput) (*ecs.DeleteAccountSettingOutput, error)
	DeleteAccountSettingAsync(ctx workflow.Context, input *ecs.DeleteAccountSettingInput) *DeleteAccountSettingFuture

	DeleteAttributes(ctx workflow.Context, input *ecs.DeleteAttributesInput) (*ecs.DeleteAttributesOutput, error)
	DeleteAttributesAsync(ctx workflow.Context, input *ecs.DeleteAttributesInput) *DeleteAttributesFuture

	DeleteCapacityProvider(ctx workflow.Context, input *ecs.DeleteCapacityProviderInput) (*ecs.DeleteCapacityProviderOutput, error)
	DeleteCapacityProviderAsync(ctx workflow.Context, input *ecs.DeleteCapacityProviderInput) *DeleteCapacityProviderFuture

	DeleteCluster(ctx workflow.Context, input *ecs.DeleteClusterInput) (*ecs.DeleteClusterOutput, error)
	DeleteClusterAsync(ctx workflow.Context, input *ecs.DeleteClusterInput) *DeleteClusterFuture

	DeleteService(ctx workflow.Context, input *ecs.DeleteServiceInput) (*ecs.DeleteServiceOutput, error)
	DeleteServiceAsync(ctx workflow.Context, input *ecs.DeleteServiceInput) *DeleteServiceFuture

	DeleteTaskSet(ctx workflow.Context, input *ecs.DeleteTaskSetInput) (*ecs.DeleteTaskSetOutput, error)
	DeleteTaskSetAsync(ctx workflow.Context, input *ecs.DeleteTaskSetInput) *DeleteTaskSetFuture

	DeregisterContainerInstance(ctx workflow.Context, input *ecs.DeregisterContainerInstanceInput) (*ecs.DeregisterContainerInstanceOutput, error)
	DeregisterContainerInstanceAsync(ctx workflow.Context, input *ecs.DeregisterContainerInstanceInput) *DeregisterContainerInstanceFuture

	DeregisterTaskDefinition(ctx workflow.Context, input *ecs.DeregisterTaskDefinitionInput) (*ecs.DeregisterTaskDefinitionOutput, error)
	DeregisterTaskDefinitionAsync(ctx workflow.Context, input *ecs.DeregisterTaskDefinitionInput) *DeregisterTaskDefinitionFuture

	DescribeCapacityProviders(ctx workflow.Context, input *ecs.DescribeCapacityProvidersInput) (*ecs.DescribeCapacityProvidersOutput, error)
	DescribeCapacityProvidersAsync(ctx workflow.Context, input *ecs.DescribeCapacityProvidersInput) *DescribeCapacityProvidersFuture

	DescribeClusters(ctx workflow.Context, input *ecs.DescribeClustersInput) (*ecs.DescribeClustersOutput, error)
	DescribeClustersAsync(ctx workflow.Context, input *ecs.DescribeClustersInput) *DescribeClustersFuture

	DescribeContainerInstances(ctx workflow.Context, input *ecs.DescribeContainerInstancesInput) (*ecs.DescribeContainerInstancesOutput, error)
	DescribeContainerInstancesAsync(ctx workflow.Context, input *ecs.DescribeContainerInstancesInput) *DescribeContainerInstancesFuture

	DescribeServices(ctx workflow.Context, input *ecs.DescribeServicesInput) (*ecs.DescribeServicesOutput, error)
	DescribeServicesAsync(ctx workflow.Context, input *ecs.DescribeServicesInput) *DescribeServicesFuture

	DescribeTaskDefinition(ctx workflow.Context, input *ecs.DescribeTaskDefinitionInput) (*ecs.DescribeTaskDefinitionOutput, error)
	DescribeTaskDefinitionAsync(ctx workflow.Context, input *ecs.DescribeTaskDefinitionInput) *DescribeTaskDefinitionFuture

	DescribeTaskSets(ctx workflow.Context, input *ecs.DescribeTaskSetsInput) (*ecs.DescribeTaskSetsOutput, error)
	DescribeTaskSetsAsync(ctx workflow.Context, input *ecs.DescribeTaskSetsInput) *DescribeTaskSetsFuture

	DescribeTasks(ctx workflow.Context, input *ecs.DescribeTasksInput) (*ecs.DescribeTasksOutput, error)
	DescribeTasksAsync(ctx workflow.Context, input *ecs.DescribeTasksInput) *DescribeTasksFuture

	DiscoverPollEndpoint(ctx workflow.Context, input *ecs.DiscoverPollEndpointInput) (*ecs.DiscoverPollEndpointOutput, error)
	DiscoverPollEndpointAsync(ctx workflow.Context, input *ecs.DiscoverPollEndpointInput) *DiscoverPollEndpointFuture

	ListAccountSettings(ctx workflow.Context, input *ecs.ListAccountSettingsInput) (*ecs.ListAccountSettingsOutput, error)
	ListAccountSettingsAsync(ctx workflow.Context, input *ecs.ListAccountSettingsInput) *ListAccountSettingsFuture

	ListAttributes(ctx workflow.Context, input *ecs.ListAttributesInput) (*ecs.ListAttributesOutput, error)
	ListAttributesAsync(ctx workflow.Context, input *ecs.ListAttributesInput) *ListAttributesFuture

	ListClusters(ctx workflow.Context, input *ecs.ListClustersInput) (*ecs.ListClustersOutput, error)
	ListClustersAsync(ctx workflow.Context, input *ecs.ListClustersInput) *ListClustersFuture

	ListContainerInstances(ctx workflow.Context, input *ecs.ListContainerInstancesInput) (*ecs.ListContainerInstancesOutput, error)
	ListContainerInstancesAsync(ctx workflow.Context, input *ecs.ListContainerInstancesInput) *ListContainerInstancesFuture

	ListServices(ctx workflow.Context, input *ecs.ListServicesInput) (*ecs.ListServicesOutput, error)
	ListServicesAsync(ctx workflow.Context, input *ecs.ListServicesInput) *ListServicesFuture

	ListTagsForResource(ctx workflow.Context, input *ecs.ListTagsForResourceInput) (*ecs.ListTagsForResourceOutput, error)
	ListTagsForResourceAsync(ctx workflow.Context, input *ecs.ListTagsForResourceInput) *ListTagsForResourceFuture

	ListTaskDefinitionFamilies(ctx workflow.Context, input *ecs.ListTaskDefinitionFamiliesInput) (*ecs.ListTaskDefinitionFamiliesOutput, error)
	ListTaskDefinitionFamiliesAsync(ctx workflow.Context, input *ecs.ListTaskDefinitionFamiliesInput) *ListTaskDefinitionFamiliesFuture

	ListTaskDefinitions(ctx workflow.Context, input *ecs.ListTaskDefinitionsInput) (*ecs.ListTaskDefinitionsOutput, error)
	ListTaskDefinitionsAsync(ctx workflow.Context, input *ecs.ListTaskDefinitionsInput) *ListTaskDefinitionsFuture

	ListTasks(ctx workflow.Context, input *ecs.ListTasksInput) (*ecs.ListTasksOutput, error)
	ListTasksAsync(ctx workflow.Context, input *ecs.ListTasksInput) *ListTasksFuture

	PutAccountSetting(ctx workflow.Context, input *ecs.PutAccountSettingInput) (*ecs.PutAccountSettingOutput, error)
	PutAccountSettingAsync(ctx workflow.Context, input *ecs.PutAccountSettingInput) *PutAccountSettingFuture

	PutAccountSettingDefault(ctx workflow.Context, input *ecs.PutAccountSettingDefaultInput) (*ecs.PutAccountSettingDefaultOutput, error)
	PutAccountSettingDefaultAsync(ctx workflow.Context, input *ecs.PutAccountSettingDefaultInput) *PutAccountSettingDefaultFuture

	PutAttributes(ctx workflow.Context, input *ecs.PutAttributesInput) (*ecs.PutAttributesOutput, error)
	PutAttributesAsync(ctx workflow.Context, input *ecs.PutAttributesInput) *PutAttributesFuture

	PutClusterCapacityProviders(ctx workflow.Context, input *ecs.PutClusterCapacityProvidersInput) (*ecs.PutClusterCapacityProvidersOutput, error)
	PutClusterCapacityProvidersAsync(ctx workflow.Context, input *ecs.PutClusterCapacityProvidersInput) *PutClusterCapacityProvidersFuture

	RegisterContainerInstance(ctx workflow.Context, input *ecs.RegisterContainerInstanceInput) (*ecs.RegisterContainerInstanceOutput, error)
	RegisterContainerInstanceAsync(ctx workflow.Context, input *ecs.RegisterContainerInstanceInput) *RegisterContainerInstanceFuture

	RegisterTaskDefinition(ctx workflow.Context, input *ecs.RegisterTaskDefinitionInput) (*ecs.RegisterTaskDefinitionOutput, error)
	RegisterTaskDefinitionAsync(ctx workflow.Context, input *ecs.RegisterTaskDefinitionInput) *RegisterTaskDefinitionFuture

	RunTask(ctx workflow.Context, input *ecs.RunTaskInput) (*ecs.RunTaskOutput, error)
	RunTaskAsync(ctx workflow.Context, input *ecs.RunTaskInput) *RunTaskFuture

	StartTask(ctx workflow.Context, input *ecs.StartTaskInput) (*ecs.StartTaskOutput, error)
	StartTaskAsync(ctx workflow.Context, input *ecs.StartTaskInput) *StartTaskFuture

	StopTask(ctx workflow.Context, input *ecs.StopTaskInput) (*ecs.StopTaskOutput, error)
	StopTaskAsync(ctx workflow.Context, input *ecs.StopTaskInput) *StopTaskFuture

	SubmitAttachmentStateChanges(ctx workflow.Context, input *ecs.SubmitAttachmentStateChangesInput) (*ecs.SubmitAttachmentStateChangesOutput, error)
	SubmitAttachmentStateChangesAsync(ctx workflow.Context, input *ecs.SubmitAttachmentStateChangesInput) *SubmitAttachmentStateChangesFuture

	SubmitContainerStateChange(ctx workflow.Context, input *ecs.SubmitContainerStateChangeInput) (*ecs.SubmitContainerStateChangeOutput, error)
	SubmitContainerStateChangeAsync(ctx workflow.Context, input *ecs.SubmitContainerStateChangeInput) *SubmitContainerStateChangeFuture

	SubmitTaskStateChange(ctx workflow.Context, input *ecs.SubmitTaskStateChangeInput) (*ecs.SubmitTaskStateChangeOutput, error)
	SubmitTaskStateChangeAsync(ctx workflow.Context, input *ecs.SubmitTaskStateChangeInput) *SubmitTaskStateChangeFuture

	TagResource(ctx workflow.Context, input *ecs.TagResourceInput) (*ecs.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *ecs.TagResourceInput) *TagResourceFuture

	UntagResource(ctx workflow.Context, input *ecs.UntagResourceInput) (*ecs.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *ecs.UntagResourceInput) *UntagResourceFuture

	UpdateClusterSettings(ctx workflow.Context, input *ecs.UpdateClusterSettingsInput) (*ecs.UpdateClusterSettingsOutput, error)
	UpdateClusterSettingsAsync(ctx workflow.Context, input *ecs.UpdateClusterSettingsInput) *UpdateClusterSettingsFuture

	UpdateContainerAgent(ctx workflow.Context, input *ecs.UpdateContainerAgentInput) (*ecs.UpdateContainerAgentOutput, error)
	UpdateContainerAgentAsync(ctx workflow.Context, input *ecs.UpdateContainerAgentInput) *UpdateContainerAgentFuture

	UpdateContainerInstancesState(ctx workflow.Context, input *ecs.UpdateContainerInstancesStateInput) (*ecs.UpdateContainerInstancesStateOutput, error)
	UpdateContainerInstancesStateAsync(ctx workflow.Context, input *ecs.UpdateContainerInstancesStateInput) *UpdateContainerInstancesStateFuture

	UpdateService(ctx workflow.Context, input *ecs.UpdateServiceInput) (*ecs.UpdateServiceOutput, error)
	UpdateServiceAsync(ctx workflow.Context, input *ecs.UpdateServiceInput) *UpdateServiceFuture

	UpdateServicePrimaryTaskSet(ctx workflow.Context, input *ecs.UpdateServicePrimaryTaskSetInput) (*ecs.UpdateServicePrimaryTaskSetOutput, error)
	UpdateServicePrimaryTaskSetAsync(ctx workflow.Context, input *ecs.UpdateServicePrimaryTaskSetInput) *UpdateServicePrimaryTaskSetFuture

	UpdateTaskSet(ctx workflow.Context, input *ecs.UpdateTaskSetInput) (*ecs.UpdateTaskSetOutput, error)
	UpdateTaskSetAsync(ctx workflow.Context, input *ecs.UpdateTaskSetInput) *UpdateTaskSetFuture

	WaitUntilServicesInactive(ctx workflow.Context, input *ecs.DescribeServicesInput) error
	WaitUntilServicesInactiveAsync(ctx workflow.Context, input *ecs.DescribeServicesInput) *clients.VoidFuture

	WaitUntilServicesStable(ctx workflow.Context, input *ecs.DescribeServicesInput) error
	WaitUntilServicesStableAsync(ctx workflow.Context, input *ecs.DescribeServicesInput) *clients.VoidFuture

	WaitUntilTasksRunning(ctx workflow.Context, input *ecs.DescribeTasksInput) error
	WaitUntilTasksRunningAsync(ctx workflow.Context, input *ecs.DescribeTasksInput) *clients.VoidFuture

	WaitUntilTasksStopped(ctx workflow.Context, input *ecs.DescribeTasksInput) error
	WaitUntilTasksStoppedAsync(ctx workflow.Context, input *ecs.DescribeTasksInput) *clients.VoidFuture
}

func NewClient() Client {
	return &stub{}
}
