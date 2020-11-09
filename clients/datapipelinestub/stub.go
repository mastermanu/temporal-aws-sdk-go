// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package datapipelinestub

import (
	"github.com/aws/aws-sdk-go/service/datapipeline"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type ActivatePipelineFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ActivatePipelineFuture) Get(ctx workflow.Context) (*datapipeline.ActivatePipelineOutput, error) {
	var output datapipeline.ActivatePipelineOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AddTagsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AddTagsFuture) Get(ctx workflow.Context) (*datapipeline.AddTagsOutput, error) {
	var output datapipeline.AddTagsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreatePipelineFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreatePipelineFuture) Get(ctx workflow.Context) (*datapipeline.CreatePipelineOutput, error) {
	var output datapipeline.CreatePipelineOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeactivatePipelineFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeactivatePipelineFuture) Get(ctx workflow.Context) (*datapipeline.DeactivatePipelineOutput, error) {
	var output datapipeline.DeactivatePipelineOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeletePipelineFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeletePipelineFuture) Get(ctx workflow.Context) (*datapipeline.DeletePipelineOutput, error) {
	var output datapipeline.DeletePipelineOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeObjectsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeObjectsFuture) Get(ctx workflow.Context) (*datapipeline.DescribeObjectsOutput, error) {
	var output datapipeline.DescribeObjectsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribePipelinesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribePipelinesFuture) Get(ctx workflow.Context) (*datapipeline.DescribePipelinesOutput, error) {
	var output datapipeline.DescribePipelinesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type EvaluateExpressionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *EvaluateExpressionFuture) Get(ctx workflow.Context) (*datapipeline.EvaluateExpressionOutput, error) {
	var output datapipeline.EvaluateExpressionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetPipelineDefinitionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetPipelineDefinitionFuture) Get(ctx workflow.Context) (*datapipeline.GetPipelineDefinitionOutput, error) {
	var output datapipeline.GetPipelineDefinitionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListPipelinesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListPipelinesFuture) Get(ctx workflow.Context) (*datapipeline.ListPipelinesOutput, error) {
	var output datapipeline.ListPipelinesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type PollForTaskFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *PollForTaskFuture) Get(ctx workflow.Context) (*datapipeline.PollForTaskOutput, error) {
	var output datapipeline.PollForTaskOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type PutPipelineDefinitionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *PutPipelineDefinitionFuture) Get(ctx workflow.Context) (*datapipeline.PutPipelineDefinitionOutput, error) {
	var output datapipeline.PutPipelineDefinitionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type QueryObjectsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *QueryObjectsFuture) Get(ctx workflow.Context) (*datapipeline.QueryObjectsOutput, error) {
	var output datapipeline.QueryObjectsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type RemoveTagsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *RemoveTagsFuture) Get(ctx workflow.Context) (*datapipeline.RemoveTagsOutput, error) {
	var output datapipeline.RemoveTagsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ReportTaskProgressFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ReportTaskProgressFuture) Get(ctx workflow.Context) (*datapipeline.ReportTaskProgressOutput, error) {
	var output datapipeline.ReportTaskProgressOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ReportTaskRunnerHeartbeatFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ReportTaskRunnerHeartbeatFuture) Get(ctx workflow.Context) (*datapipeline.ReportTaskRunnerHeartbeatOutput, error) {
	var output datapipeline.ReportTaskRunnerHeartbeatOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SetStatusFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SetStatusFuture) Get(ctx workflow.Context) (*datapipeline.SetStatusOutput, error) {
	var output datapipeline.SetStatusOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SetTaskStatusFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SetTaskStatusFuture) Get(ctx workflow.Context) (*datapipeline.SetTaskStatusOutput, error) {
	var output datapipeline.SetTaskStatusOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ValidatePipelineDefinitionFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ValidatePipelineDefinitionFuture) Get(ctx workflow.Context) (*datapipeline.ValidatePipelineDefinitionOutput, error) {
	var output datapipeline.ValidatePipelineDefinitionOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) ActivatePipeline(ctx workflow.Context, input *datapipeline.ActivatePipelineInput) (*datapipeline.ActivatePipelineOutput, error) {
	var output datapipeline.ActivatePipelineOutput
	err := workflow.ExecuteActivity(ctx, "aws.datapipeline.ActivatePipeline", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ActivatePipelineAsync(ctx workflow.Context, input *datapipeline.ActivatePipelineInput) *ActivatePipelineFuture {
	future := workflow.ExecuteActivity(ctx, "aws.datapipeline.ActivatePipeline", input)
	return &ActivatePipelineFuture{Future: future}
}

func (a *stub) AddTags(ctx workflow.Context, input *datapipeline.AddTagsInput) (*datapipeline.AddTagsOutput, error) {
	var output datapipeline.AddTagsOutput
	err := workflow.ExecuteActivity(ctx, "aws.datapipeline.AddTags", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) AddTagsAsync(ctx workflow.Context, input *datapipeline.AddTagsInput) *AddTagsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.datapipeline.AddTags", input)
	return &AddTagsFuture{Future: future}
}

func (a *stub) CreatePipeline(ctx workflow.Context, input *datapipeline.CreatePipelineInput) (*datapipeline.CreatePipelineOutput, error) {
	var output datapipeline.CreatePipelineOutput
	err := workflow.ExecuteActivity(ctx, "aws.datapipeline.CreatePipeline", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreatePipelineAsync(ctx workflow.Context, input *datapipeline.CreatePipelineInput) *CreatePipelineFuture {
	future := workflow.ExecuteActivity(ctx, "aws.datapipeline.CreatePipeline", input)
	return &CreatePipelineFuture{Future: future}
}

func (a *stub) DeactivatePipeline(ctx workflow.Context, input *datapipeline.DeactivatePipelineInput) (*datapipeline.DeactivatePipelineOutput, error) {
	var output datapipeline.DeactivatePipelineOutput
	err := workflow.ExecuteActivity(ctx, "aws.datapipeline.DeactivatePipeline", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeactivatePipelineAsync(ctx workflow.Context, input *datapipeline.DeactivatePipelineInput) *DeactivatePipelineFuture {
	future := workflow.ExecuteActivity(ctx, "aws.datapipeline.DeactivatePipeline", input)
	return &DeactivatePipelineFuture{Future: future}
}

func (a *stub) DeletePipeline(ctx workflow.Context, input *datapipeline.DeletePipelineInput) (*datapipeline.DeletePipelineOutput, error) {
	var output datapipeline.DeletePipelineOutput
	err := workflow.ExecuteActivity(ctx, "aws.datapipeline.DeletePipeline", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeletePipelineAsync(ctx workflow.Context, input *datapipeline.DeletePipelineInput) *DeletePipelineFuture {
	future := workflow.ExecuteActivity(ctx, "aws.datapipeline.DeletePipeline", input)
	return &DeletePipelineFuture{Future: future}
}

func (a *stub) DescribeObjects(ctx workflow.Context, input *datapipeline.DescribeObjectsInput) (*datapipeline.DescribeObjectsOutput, error) {
	var output datapipeline.DescribeObjectsOutput
	err := workflow.ExecuteActivity(ctx, "aws.datapipeline.DescribeObjects", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeObjectsAsync(ctx workflow.Context, input *datapipeline.DescribeObjectsInput) *DescribeObjectsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.datapipeline.DescribeObjects", input)
	return &DescribeObjectsFuture{Future: future}
}

func (a *stub) DescribePipelines(ctx workflow.Context, input *datapipeline.DescribePipelinesInput) (*datapipeline.DescribePipelinesOutput, error) {
	var output datapipeline.DescribePipelinesOutput
	err := workflow.ExecuteActivity(ctx, "aws.datapipeline.DescribePipelines", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribePipelinesAsync(ctx workflow.Context, input *datapipeline.DescribePipelinesInput) *DescribePipelinesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.datapipeline.DescribePipelines", input)
	return &DescribePipelinesFuture{Future: future}
}

func (a *stub) EvaluateExpression(ctx workflow.Context, input *datapipeline.EvaluateExpressionInput) (*datapipeline.EvaluateExpressionOutput, error) {
	var output datapipeline.EvaluateExpressionOutput
	err := workflow.ExecuteActivity(ctx, "aws.datapipeline.EvaluateExpression", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) EvaluateExpressionAsync(ctx workflow.Context, input *datapipeline.EvaluateExpressionInput) *EvaluateExpressionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.datapipeline.EvaluateExpression", input)
	return &EvaluateExpressionFuture{Future: future}
}

func (a *stub) GetPipelineDefinition(ctx workflow.Context, input *datapipeline.GetPipelineDefinitionInput) (*datapipeline.GetPipelineDefinitionOutput, error) {
	var output datapipeline.GetPipelineDefinitionOutput
	err := workflow.ExecuteActivity(ctx, "aws.datapipeline.GetPipelineDefinition", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetPipelineDefinitionAsync(ctx workflow.Context, input *datapipeline.GetPipelineDefinitionInput) *GetPipelineDefinitionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.datapipeline.GetPipelineDefinition", input)
	return &GetPipelineDefinitionFuture{Future: future}
}

func (a *stub) ListPipelines(ctx workflow.Context, input *datapipeline.ListPipelinesInput) (*datapipeline.ListPipelinesOutput, error) {
	var output datapipeline.ListPipelinesOutput
	err := workflow.ExecuteActivity(ctx, "aws.datapipeline.ListPipelines", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListPipelinesAsync(ctx workflow.Context, input *datapipeline.ListPipelinesInput) *ListPipelinesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.datapipeline.ListPipelines", input)
	return &ListPipelinesFuture{Future: future}
}

func (a *stub) PollForTask(ctx workflow.Context, input *datapipeline.PollForTaskInput) (*datapipeline.PollForTaskOutput, error) {
	var output datapipeline.PollForTaskOutput
	err := workflow.ExecuteActivity(ctx, "aws.datapipeline.PollForTask", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PollForTaskAsync(ctx workflow.Context, input *datapipeline.PollForTaskInput) *PollForTaskFuture {
	future := workflow.ExecuteActivity(ctx, "aws.datapipeline.PollForTask", input)
	return &PollForTaskFuture{Future: future}
}

func (a *stub) PutPipelineDefinition(ctx workflow.Context, input *datapipeline.PutPipelineDefinitionInput) (*datapipeline.PutPipelineDefinitionOutput, error) {
	var output datapipeline.PutPipelineDefinitionOutput
	err := workflow.ExecuteActivity(ctx, "aws.datapipeline.PutPipelineDefinition", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PutPipelineDefinitionAsync(ctx workflow.Context, input *datapipeline.PutPipelineDefinitionInput) *PutPipelineDefinitionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.datapipeline.PutPipelineDefinition", input)
	return &PutPipelineDefinitionFuture{Future: future}
}

func (a *stub) QueryObjects(ctx workflow.Context, input *datapipeline.QueryObjectsInput) (*datapipeline.QueryObjectsOutput, error) {
	var output datapipeline.QueryObjectsOutput
	err := workflow.ExecuteActivity(ctx, "aws.datapipeline.QueryObjects", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) QueryObjectsAsync(ctx workflow.Context, input *datapipeline.QueryObjectsInput) *QueryObjectsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.datapipeline.QueryObjects", input)
	return &QueryObjectsFuture{Future: future}
}

func (a *stub) RemoveTags(ctx workflow.Context, input *datapipeline.RemoveTagsInput) (*datapipeline.RemoveTagsOutput, error) {
	var output datapipeline.RemoveTagsOutput
	err := workflow.ExecuteActivity(ctx, "aws.datapipeline.RemoveTags", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) RemoveTagsAsync(ctx workflow.Context, input *datapipeline.RemoveTagsInput) *RemoveTagsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.datapipeline.RemoveTags", input)
	return &RemoveTagsFuture{Future: future}
}

func (a *stub) ReportTaskProgress(ctx workflow.Context, input *datapipeline.ReportTaskProgressInput) (*datapipeline.ReportTaskProgressOutput, error) {
	var output datapipeline.ReportTaskProgressOutput
	err := workflow.ExecuteActivity(ctx, "aws.datapipeline.ReportTaskProgress", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ReportTaskProgressAsync(ctx workflow.Context, input *datapipeline.ReportTaskProgressInput) *ReportTaskProgressFuture {
	future := workflow.ExecuteActivity(ctx, "aws.datapipeline.ReportTaskProgress", input)
	return &ReportTaskProgressFuture{Future: future}
}

func (a *stub) ReportTaskRunnerHeartbeat(ctx workflow.Context, input *datapipeline.ReportTaskRunnerHeartbeatInput) (*datapipeline.ReportTaskRunnerHeartbeatOutput, error) {
	var output datapipeline.ReportTaskRunnerHeartbeatOutput
	err := workflow.ExecuteActivity(ctx, "aws.datapipeline.ReportTaskRunnerHeartbeat", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ReportTaskRunnerHeartbeatAsync(ctx workflow.Context, input *datapipeline.ReportTaskRunnerHeartbeatInput) *ReportTaskRunnerHeartbeatFuture {
	future := workflow.ExecuteActivity(ctx, "aws.datapipeline.ReportTaskRunnerHeartbeat", input)
	return &ReportTaskRunnerHeartbeatFuture{Future: future}
}

func (a *stub) SetStatus(ctx workflow.Context, input *datapipeline.SetStatusInput) (*datapipeline.SetStatusOutput, error) {
	var output datapipeline.SetStatusOutput
	err := workflow.ExecuteActivity(ctx, "aws.datapipeline.SetStatus", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) SetStatusAsync(ctx workflow.Context, input *datapipeline.SetStatusInput) *SetStatusFuture {
	future := workflow.ExecuteActivity(ctx, "aws.datapipeline.SetStatus", input)
	return &SetStatusFuture{Future: future}
}

func (a *stub) SetTaskStatus(ctx workflow.Context, input *datapipeline.SetTaskStatusInput) (*datapipeline.SetTaskStatusOutput, error) {
	var output datapipeline.SetTaskStatusOutput
	err := workflow.ExecuteActivity(ctx, "aws.datapipeline.SetTaskStatus", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) SetTaskStatusAsync(ctx workflow.Context, input *datapipeline.SetTaskStatusInput) *SetTaskStatusFuture {
	future := workflow.ExecuteActivity(ctx, "aws.datapipeline.SetTaskStatus", input)
	return &SetTaskStatusFuture{Future: future}
}

func (a *stub) ValidatePipelineDefinition(ctx workflow.Context, input *datapipeline.ValidatePipelineDefinitionInput) (*datapipeline.ValidatePipelineDefinitionOutput, error) {
	var output datapipeline.ValidatePipelineDefinitionOutput
	err := workflow.ExecuteActivity(ctx, "aws.datapipeline.ValidatePipelineDefinition", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ValidatePipelineDefinitionAsync(ctx workflow.Context, input *datapipeline.ValidatePipelineDefinitionInput) *ValidatePipelineDefinitionFuture {
	future := workflow.ExecuteActivity(ctx, "aws.datapipeline.ValidatePipelineDefinition", input)
	return &ValidatePipelineDefinitionFuture{Future: future}
}