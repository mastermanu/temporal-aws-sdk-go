// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package iotdataplanestub

import (
	"github.com/aws/aws-sdk-go/service/iotdataplane"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type DeleteThingShadowFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteThingShadowFuture) Get(ctx workflow.Context) (*iotdataplane.DeleteThingShadowOutput, error) {
	var output iotdataplane.DeleteThingShadowOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetThingShadowFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetThingShadowFuture) Get(ctx workflow.Context) (*iotdataplane.GetThingShadowOutput, error) {
	var output iotdataplane.GetThingShadowOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListNamedShadowsForThingFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListNamedShadowsForThingFuture) Get(ctx workflow.Context) (*iotdataplane.ListNamedShadowsForThingOutput, error) {
	var output iotdataplane.ListNamedShadowsForThingOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type PublishFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *PublishFuture) Get(ctx workflow.Context) (*iotdataplane.PublishOutput, error) {
	var output iotdataplane.PublishOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateThingShadowFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateThingShadowFuture) Get(ctx workflow.Context) (*iotdataplane.UpdateThingShadowOutput, error) {
	var output iotdataplane.UpdateThingShadowOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteThingShadow(ctx workflow.Context, input *iotdataplane.DeleteThingShadowInput) (*iotdataplane.DeleteThingShadowOutput, error) {
	var output iotdataplane.DeleteThingShadowOutput
	err := workflow.ExecuteActivity(ctx, "aws.iotdataplane.DeleteThingShadow", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteThingShadowAsync(ctx workflow.Context, input *iotdataplane.DeleteThingShadowInput) *DeleteThingShadowFuture {
	future := workflow.ExecuteActivity(ctx, "aws.iotdataplane.DeleteThingShadow", input)
	return &DeleteThingShadowFuture{Future: future}
}

func (a *stub) GetThingShadow(ctx workflow.Context, input *iotdataplane.GetThingShadowInput) (*iotdataplane.GetThingShadowOutput, error) {
	var output iotdataplane.GetThingShadowOutput
	err := workflow.ExecuteActivity(ctx, "aws.iotdataplane.GetThingShadow", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetThingShadowAsync(ctx workflow.Context, input *iotdataplane.GetThingShadowInput) *GetThingShadowFuture {
	future := workflow.ExecuteActivity(ctx, "aws.iotdataplane.GetThingShadow", input)
	return &GetThingShadowFuture{Future: future}
}

func (a *stub) ListNamedShadowsForThing(ctx workflow.Context, input *iotdataplane.ListNamedShadowsForThingInput) (*iotdataplane.ListNamedShadowsForThingOutput, error) {
	var output iotdataplane.ListNamedShadowsForThingOutput
	err := workflow.ExecuteActivity(ctx, "aws.iotdataplane.ListNamedShadowsForThing", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListNamedShadowsForThingAsync(ctx workflow.Context, input *iotdataplane.ListNamedShadowsForThingInput) *ListNamedShadowsForThingFuture {
	future := workflow.ExecuteActivity(ctx, "aws.iotdataplane.ListNamedShadowsForThing", input)
	return &ListNamedShadowsForThingFuture{Future: future}
}

func (a *stub) Publish(ctx workflow.Context, input *iotdataplane.PublishInput) (*iotdataplane.PublishOutput, error) {
	var output iotdataplane.PublishOutput
	err := workflow.ExecuteActivity(ctx, "aws.iotdataplane.Publish", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PublishAsync(ctx workflow.Context, input *iotdataplane.PublishInput) *PublishFuture {
	future := workflow.ExecuteActivity(ctx, "aws.iotdataplane.Publish", input)
	return &PublishFuture{Future: future}
}

func (a *stub) UpdateThingShadow(ctx workflow.Context, input *iotdataplane.UpdateThingShadowInput) (*iotdataplane.UpdateThingShadowOutput, error) {
	var output iotdataplane.UpdateThingShadowOutput
	err := workflow.ExecuteActivity(ctx, "aws.iotdataplane.UpdateThingShadow", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateThingShadowAsync(ctx workflow.Context, input *iotdataplane.UpdateThingShadowInput) *UpdateThingShadowFuture {
	future := workflow.ExecuteActivity(ctx, "aws.iotdataplane.UpdateThingShadow", input)
	return &UpdateThingShadowFuture{Future: future}
}
