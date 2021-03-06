// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package ioteventsstub

import (
	"github.com/aws/aws-sdk-go/service/iotevents"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type CreateDetectorModelFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateDetectorModelFuture) Get(ctx workflow.Context) (*iotevents.CreateDetectorModelOutput, error) {
	var output iotevents.CreateDetectorModelOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateInputFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateInputFuture) Get(ctx workflow.Context) (*iotevents.CreateInputOutput, error) {
	var output iotevents.CreateInputOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteDetectorModelFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteDetectorModelFuture) Get(ctx workflow.Context) (*iotevents.DeleteDetectorModelOutput, error) {
	var output iotevents.DeleteDetectorModelOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteInputFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteInputFuture) Get(ctx workflow.Context) (*iotevents.DeleteInputOutput, error) {
	var output iotevents.DeleteInputOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeDetectorModelFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeDetectorModelFuture) Get(ctx workflow.Context) (*iotevents.DescribeDetectorModelOutput, error) {
	var output iotevents.DescribeDetectorModelOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeInputFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeInputFuture) Get(ctx workflow.Context) (*iotevents.DescribeInputOutput, error) {
	var output iotevents.DescribeInputOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeLoggingOptionsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeLoggingOptionsFuture) Get(ctx workflow.Context) (*iotevents.DescribeLoggingOptionsOutput, error) {
	var output iotevents.DescribeLoggingOptionsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListDetectorModelVersionsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListDetectorModelVersionsFuture) Get(ctx workflow.Context) (*iotevents.ListDetectorModelVersionsOutput, error) {
	var output iotevents.ListDetectorModelVersionsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListDetectorModelsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListDetectorModelsFuture) Get(ctx workflow.Context) (*iotevents.ListDetectorModelsOutput, error) {
	var output iotevents.ListDetectorModelsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListInputsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListInputsFuture) Get(ctx workflow.Context) (*iotevents.ListInputsOutput, error) {
	var output iotevents.ListInputsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListTagsForResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListTagsForResourceFuture) Get(ctx workflow.Context) (*iotevents.ListTagsForResourceOutput, error) {
	var output iotevents.ListTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type PutLoggingOptionsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *PutLoggingOptionsFuture) Get(ctx workflow.Context) (*iotevents.PutLoggingOptionsOutput, error) {
	var output iotevents.PutLoggingOptionsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TagResourceFuture) Get(ctx workflow.Context) (*iotevents.TagResourceOutput, error) {
	var output iotevents.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UntagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UntagResourceFuture) Get(ctx workflow.Context) (*iotevents.UntagResourceOutput, error) {
	var output iotevents.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateDetectorModelFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateDetectorModelFuture) Get(ctx workflow.Context) (*iotevents.UpdateDetectorModelOutput, error) {
	var output iotevents.UpdateDetectorModelOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateInputFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateInputFuture) Get(ctx workflow.Context) (*iotevents.UpdateInputOutput, error) {
	var output iotevents.UpdateInputOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateDetectorModel(ctx workflow.Context, input *iotevents.CreateDetectorModelInput) (*iotevents.CreateDetectorModelOutput, error) {
	var output iotevents.CreateDetectorModelOutput
	err := workflow.ExecuteActivity(ctx, "aws.iotevents.CreateDetectorModel", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateDetectorModelAsync(ctx workflow.Context, input *iotevents.CreateDetectorModelInput) *CreateDetectorModelFuture {
	future := workflow.ExecuteActivity(ctx, "aws.iotevents.CreateDetectorModel", input)
	return &CreateDetectorModelFuture{Future: future}
}

func (a *stub) CreateInput(ctx workflow.Context, input *iotevents.CreateInputInput) (*iotevents.CreateInputOutput, error) {
	var output iotevents.CreateInputOutput
	err := workflow.ExecuteActivity(ctx, "aws.iotevents.CreateInput", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateInputAsync(ctx workflow.Context, input *iotevents.CreateInputInput) *CreateInputFuture {
	future := workflow.ExecuteActivity(ctx, "aws.iotevents.CreateInput", input)
	return &CreateInputFuture{Future: future}
}

func (a *stub) DeleteDetectorModel(ctx workflow.Context, input *iotevents.DeleteDetectorModelInput) (*iotevents.DeleteDetectorModelOutput, error) {
	var output iotevents.DeleteDetectorModelOutput
	err := workflow.ExecuteActivity(ctx, "aws.iotevents.DeleteDetectorModel", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteDetectorModelAsync(ctx workflow.Context, input *iotevents.DeleteDetectorModelInput) *DeleteDetectorModelFuture {
	future := workflow.ExecuteActivity(ctx, "aws.iotevents.DeleteDetectorModel", input)
	return &DeleteDetectorModelFuture{Future: future}
}

func (a *stub) DeleteInput(ctx workflow.Context, input *iotevents.DeleteInputInput) (*iotevents.DeleteInputOutput, error) {
	var output iotevents.DeleteInputOutput
	err := workflow.ExecuteActivity(ctx, "aws.iotevents.DeleteInput", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteInputAsync(ctx workflow.Context, input *iotevents.DeleteInputInput) *DeleteInputFuture {
	future := workflow.ExecuteActivity(ctx, "aws.iotevents.DeleteInput", input)
	return &DeleteInputFuture{Future: future}
}

func (a *stub) DescribeDetectorModel(ctx workflow.Context, input *iotevents.DescribeDetectorModelInput) (*iotevents.DescribeDetectorModelOutput, error) {
	var output iotevents.DescribeDetectorModelOutput
	err := workflow.ExecuteActivity(ctx, "aws.iotevents.DescribeDetectorModel", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeDetectorModelAsync(ctx workflow.Context, input *iotevents.DescribeDetectorModelInput) *DescribeDetectorModelFuture {
	future := workflow.ExecuteActivity(ctx, "aws.iotevents.DescribeDetectorModel", input)
	return &DescribeDetectorModelFuture{Future: future}
}

func (a *stub) DescribeInput(ctx workflow.Context, input *iotevents.DescribeInputInput) (*iotevents.DescribeInputOutput, error) {
	var output iotevents.DescribeInputOutput
	err := workflow.ExecuteActivity(ctx, "aws.iotevents.DescribeInput", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeInputAsync(ctx workflow.Context, input *iotevents.DescribeInputInput) *DescribeInputFuture {
	future := workflow.ExecuteActivity(ctx, "aws.iotevents.DescribeInput", input)
	return &DescribeInputFuture{Future: future}
}

func (a *stub) DescribeLoggingOptions(ctx workflow.Context, input *iotevents.DescribeLoggingOptionsInput) (*iotevents.DescribeLoggingOptionsOutput, error) {
	var output iotevents.DescribeLoggingOptionsOutput
	err := workflow.ExecuteActivity(ctx, "aws.iotevents.DescribeLoggingOptions", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeLoggingOptionsAsync(ctx workflow.Context, input *iotevents.DescribeLoggingOptionsInput) *DescribeLoggingOptionsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.iotevents.DescribeLoggingOptions", input)
	return &DescribeLoggingOptionsFuture{Future: future}
}

func (a *stub) ListDetectorModelVersions(ctx workflow.Context, input *iotevents.ListDetectorModelVersionsInput) (*iotevents.ListDetectorModelVersionsOutput, error) {
	var output iotevents.ListDetectorModelVersionsOutput
	err := workflow.ExecuteActivity(ctx, "aws.iotevents.ListDetectorModelVersions", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListDetectorModelVersionsAsync(ctx workflow.Context, input *iotevents.ListDetectorModelVersionsInput) *ListDetectorModelVersionsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.iotevents.ListDetectorModelVersions", input)
	return &ListDetectorModelVersionsFuture{Future: future}
}

func (a *stub) ListDetectorModels(ctx workflow.Context, input *iotevents.ListDetectorModelsInput) (*iotevents.ListDetectorModelsOutput, error) {
	var output iotevents.ListDetectorModelsOutput
	err := workflow.ExecuteActivity(ctx, "aws.iotevents.ListDetectorModels", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListDetectorModelsAsync(ctx workflow.Context, input *iotevents.ListDetectorModelsInput) *ListDetectorModelsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.iotevents.ListDetectorModels", input)
	return &ListDetectorModelsFuture{Future: future}
}

func (a *stub) ListInputs(ctx workflow.Context, input *iotevents.ListInputsInput) (*iotevents.ListInputsOutput, error) {
	var output iotevents.ListInputsOutput
	err := workflow.ExecuteActivity(ctx, "aws.iotevents.ListInputs", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListInputsAsync(ctx workflow.Context, input *iotevents.ListInputsInput) *ListInputsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.iotevents.ListInputs", input)
	return &ListInputsFuture{Future: future}
}

func (a *stub) ListTagsForResource(ctx workflow.Context, input *iotevents.ListTagsForResourceInput) (*iotevents.ListTagsForResourceOutput, error) {
	var output iotevents.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.iotevents.ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTagsForResourceAsync(ctx workflow.Context, input *iotevents.ListTagsForResourceInput) *ListTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.iotevents.ListTagsForResource", input)
	return &ListTagsForResourceFuture{Future: future}
}

func (a *stub) PutLoggingOptions(ctx workflow.Context, input *iotevents.PutLoggingOptionsInput) (*iotevents.PutLoggingOptionsOutput, error) {
	var output iotevents.PutLoggingOptionsOutput
	err := workflow.ExecuteActivity(ctx, "aws.iotevents.PutLoggingOptions", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PutLoggingOptionsAsync(ctx workflow.Context, input *iotevents.PutLoggingOptionsInput) *PutLoggingOptionsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.iotevents.PutLoggingOptions", input)
	return &PutLoggingOptionsFuture{Future: future}
}

func (a *stub) TagResource(ctx workflow.Context, input *iotevents.TagResourceInput) (*iotevents.TagResourceOutput, error) {
	var output iotevents.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.iotevents.TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TagResourceAsync(ctx workflow.Context, input *iotevents.TagResourceInput) *TagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.iotevents.TagResource", input)
	return &TagResourceFuture{Future: future}
}

func (a *stub) UntagResource(ctx workflow.Context, input *iotevents.UntagResourceInput) (*iotevents.UntagResourceOutput, error) {
	var output iotevents.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.iotevents.UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UntagResourceAsync(ctx workflow.Context, input *iotevents.UntagResourceInput) *UntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.iotevents.UntagResource", input)
	return &UntagResourceFuture{Future: future}
}

func (a *stub) UpdateDetectorModel(ctx workflow.Context, input *iotevents.UpdateDetectorModelInput) (*iotevents.UpdateDetectorModelOutput, error) {
	var output iotevents.UpdateDetectorModelOutput
	err := workflow.ExecuteActivity(ctx, "aws.iotevents.UpdateDetectorModel", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateDetectorModelAsync(ctx workflow.Context, input *iotevents.UpdateDetectorModelInput) *UpdateDetectorModelFuture {
	future := workflow.ExecuteActivity(ctx, "aws.iotevents.UpdateDetectorModel", input)
	return &UpdateDetectorModelFuture{Future: future}
}

func (a *stub) UpdateInput(ctx workflow.Context, input *iotevents.UpdateInputInput) (*iotevents.UpdateInputOutput, error) {
	var output iotevents.UpdateInputOutput
	err := workflow.ExecuteActivity(ctx, "aws.iotevents.UpdateInput", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateInputAsync(ctx workflow.Context, input *iotevents.UpdateInputInput) *UpdateInputFuture {
	future := workflow.ExecuteActivity(ctx, "aws.iotevents.UpdateInput", input)
	return &UpdateInputFuture{Future: future}
}
