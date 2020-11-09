// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package autoscalingplansstub

import (
	"github.com/aws/aws-sdk-go/service/autoscalingplans"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type CreateScalingPlanFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateScalingPlanFuture) Get(ctx workflow.Context) (*autoscalingplans.CreateScalingPlanOutput, error) {
	var output autoscalingplans.CreateScalingPlanOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteScalingPlanFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteScalingPlanFuture) Get(ctx workflow.Context) (*autoscalingplans.DeleteScalingPlanOutput, error) {
	var output autoscalingplans.DeleteScalingPlanOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeScalingPlanResourcesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeScalingPlanResourcesFuture) Get(ctx workflow.Context) (*autoscalingplans.DescribeScalingPlanResourcesOutput, error) {
	var output autoscalingplans.DescribeScalingPlanResourcesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeScalingPlansFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeScalingPlansFuture) Get(ctx workflow.Context) (*autoscalingplans.DescribeScalingPlansOutput, error) {
	var output autoscalingplans.DescribeScalingPlansOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetScalingPlanResourceForecastDataFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetScalingPlanResourceForecastDataFuture) Get(ctx workflow.Context) (*autoscalingplans.GetScalingPlanResourceForecastDataOutput, error) {
	var output autoscalingplans.GetScalingPlanResourceForecastDataOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateScalingPlanFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateScalingPlanFuture) Get(ctx workflow.Context) (*autoscalingplans.UpdateScalingPlanOutput, error) {
	var output autoscalingplans.UpdateScalingPlanOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateScalingPlan(ctx workflow.Context, input *autoscalingplans.CreateScalingPlanInput) (*autoscalingplans.CreateScalingPlanOutput, error) {
	var output autoscalingplans.CreateScalingPlanOutput
	err := workflow.ExecuteActivity(ctx, "aws.autoscalingplans.CreateScalingPlan", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateScalingPlanAsync(ctx workflow.Context, input *autoscalingplans.CreateScalingPlanInput) *CreateScalingPlanFuture {
	future := workflow.ExecuteActivity(ctx, "aws.autoscalingplans.CreateScalingPlan", input)
	return &CreateScalingPlanFuture{Future: future}
}

func (a *stub) DeleteScalingPlan(ctx workflow.Context, input *autoscalingplans.DeleteScalingPlanInput) (*autoscalingplans.DeleteScalingPlanOutput, error) {
	var output autoscalingplans.DeleteScalingPlanOutput
	err := workflow.ExecuteActivity(ctx, "aws.autoscalingplans.DeleteScalingPlan", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteScalingPlanAsync(ctx workflow.Context, input *autoscalingplans.DeleteScalingPlanInput) *DeleteScalingPlanFuture {
	future := workflow.ExecuteActivity(ctx, "aws.autoscalingplans.DeleteScalingPlan", input)
	return &DeleteScalingPlanFuture{Future: future}
}

func (a *stub) DescribeScalingPlanResources(ctx workflow.Context, input *autoscalingplans.DescribeScalingPlanResourcesInput) (*autoscalingplans.DescribeScalingPlanResourcesOutput, error) {
	var output autoscalingplans.DescribeScalingPlanResourcesOutput
	err := workflow.ExecuteActivity(ctx, "aws.autoscalingplans.DescribeScalingPlanResources", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeScalingPlanResourcesAsync(ctx workflow.Context, input *autoscalingplans.DescribeScalingPlanResourcesInput) *DescribeScalingPlanResourcesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.autoscalingplans.DescribeScalingPlanResources", input)
	return &DescribeScalingPlanResourcesFuture{Future: future}
}

func (a *stub) DescribeScalingPlans(ctx workflow.Context, input *autoscalingplans.DescribeScalingPlansInput) (*autoscalingplans.DescribeScalingPlansOutput, error) {
	var output autoscalingplans.DescribeScalingPlansOutput
	err := workflow.ExecuteActivity(ctx, "aws.autoscalingplans.DescribeScalingPlans", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeScalingPlansAsync(ctx workflow.Context, input *autoscalingplans.DescribeScalingPlansInput) *DescribeScalingPlansFuture {
	future := workflow.ExecuteActivity(ctx, "aws.autoscalingplans.DescribeScalingPlans", input)
	return &DescribeScalingPlansFuture{Future: future}
}

func (a *stub) GetScalingPlanResourceForecastData(ctx workflow.Context, input *autoscalingplans.GetScalingPlanResourceForecastDataInput) (*autoscalingplans.GetScalingPlanResourceForecastDataOutput, error) {
	var output autoscalingplans.GetScalingPlanResourceForecastDataOutput
	err := workflow.ExecuteActivity(ctx, "aws.autoscalingplans.GetScalingPlanResourceForecastData", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetScalingPlanResourceForecastDataAsync(ctx workflow.Context, input *autoscalingplans.GetScalingPlanResourceForecastDataInput) *GetScalingPlanResourceForecastDataFuture {
	future := workflow.ExecuteActivity(ctx, "aws.autoscalingplans.GetScalingPlanResourceForecastData", input)
	return &GetScalingPlanResourceForecastDataFuture{Future: future}
}

func (a *stub) UpdateScalingPlan(ctx workflow.Context, input *autoscalingplans.UpdateScalingPlanInput) (*autoscalingplans.UpdateScalingPlanOutput, error) {
	var output autoscalingplans.UpdateScalingPlanOutput
	err := workflow.ExecuteActivity(ctx, "aws.autoscalingplans.UpdateScalingPlan", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateScalingPlanAsync(ctx workflow.Context, input *autoscalingplans.UpdateScalingPlanInput) *UpdateScalingPlanFuture {
	future := workflow.ExecuteActivity(ctx, "aws.autoscalingplans.UpdateScalingPlan", input)
	return &UpdateScalingPlanFuture{Future: future}
}