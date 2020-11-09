// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package mediapackagestub

import (
	"github.com/aws/aws-sdk-go/service/mediapackage"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type ConfigureLogsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ConfigureLogsFuture) Get(ctx workflow.Context) (*mediapackage.ConfigureLogsOutput, error) {
	var output mediapackage.ConfigureLogsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateChannelFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateChannelFuture) Get(ctx workflow.Context) (*mediapackage.CreateChannelOutput, error) {
	var output mediapackage.CreateChannelOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateHarvestJobFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateHarvestJobFuture) Get(ctx workflow.Context) (*mediapackage.CreateHarvestJobOutput, error) {
	var output mediapackage.CreateHarvestJobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateOriginEndpointFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateOriginEndpointFuture) Get(ctx workflow.Context) (*mediapackage.CreateOriginEndpointOutput, error) {
	var output mediapackage.CreateOriginEndpointOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteChannelFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteChannelFuture) Get(ctx workflow.Context) (*mediapackage.DeleteChannelOutput, error) {
	var output mediapackage.DeleteChannelOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteOriginEndpointFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteOriginEndpointFuture) Get(ctx workflow.Context) (*mediapackage.DeleteOriginEndpointOutput, error) {
	var output mediapackage.DeleteOriginEndpointOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeChannelFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeChannelFuture) Get(ctx workflow.Context) (*mediapackage.DescribeChannelOutput, error) {
	var output mediapackage.DescribeChannelOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeHarvestJobFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeHarvestJobFuture) Get(ctx workflow.Context) (*mediapackage.DescribeHarvestJobOutput, error) {
	var output mediapackage.DescribeHarvestJobOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeOriginEndpointFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeOriginEndpointFuture) Get(ctx workflow.Context) (*mediapackage.DescribeOriginEndpointOutput, error) {
	var output mediapackage.DescribeOriginEndpointOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListChannelsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListChannelsFuture) Get(ctx workflow.Context) (*mediapackage.ListChannelsOutput, error) {
	var output mediapackage.ListChannelsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListHarvestJobsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListHarvestJobsFuture) Get(ctx workflow.Context) (*mediapackage.ListHarvestJobsOutput, error) {
	var output mediapackage.ListHarvestJobsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListOriginEndpointsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListOriginEndpointsFuture) Get(ctx workflow.Context) (*mediapackage.ListOriginEndpointsOutput, error) {
	var output mediapackage.ListOriginEndpointsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListTagsForResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListTagsForResourceFuture) Get(ctx workflow.Context) (*mediapackage.ListTagsForResourceOutput, error) {
	var output mediapackage.ListTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type RotateChannelCredentialsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *RotateChannelCredentialsFuture) Get(ctx workflow.Context) (*mediapackage.RotateChannelCredentialsOutput, error) {
	var output mediapackage.RotateChannelCredentialsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type RotateIngestEndpointCredentialsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *RotateIngestEndpointCredentialsFuture) Get(ctx workflow.Context) (*mediapackage.RotateIngestEndpointCredentialsOutput, error) {
	var output mediapackage.RotateIngestEndpointCredentialsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TagResourceFuture) Get(ctx workflow.Context) (*mediapackage.TagResourceOutput, error) {
	var output mediapackage.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UntagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UntagResourceFuture) Get(ctx workflow.Context) (*mediapackage.UntagResourceOutput, error) {
	var output mediapackage.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateChannelFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateChannelFuture) Get(ctx workflow.Context) (*mediapackage.UpdateChannelOutput, error) {
	var output mediapackage.UpdateChannelOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UpdateOriginEndpointFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UpdateOriginEndpointFuture) Get(ctx workflow.Context) (*mediapackage.UpdateOriginEndpointOutput, error) {
	var output mediapackage.UpdateOriginEndpointOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) ConfigureLogs(ctx workflow.Context, input *mediapackage.ConfigureLogsInput) (*mediapackage.ConfigureLogsOutput, error) {
	var output mediapackage.ConfigureLogsOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediapackage.ConfigureLogs", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ConfigureLogsAsync(ctx workflow.Context, input *mediapackage.ConfigureLogsInput) *ConfigureLogsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediapackage.ConfigureLogs", input)
	return &ConfigureLogsFuture{Future: future}
}

func (a *stub) CreateChannel(ctx workflow.Context, input *mediapackage.CreateChannelInput) (*mediapackage.CreateChannelOutput, error) {
	var output mediapackage.CreateChannelOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediapackage.CreateChannel", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateChannelAsync(ctx workflow.Context, input *mediapackage.CreateChannelInput) *CreateChannelFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediapackage.CreateChannel", input)
	return &CreateChannelFuture{Future: future}
}

func (a *stub) CreateHarvestJob(ctx workflow.Context, input *mediapackage.CreateHarvestJobInput) (*mediapackage.CreateHarvestJobOutput, error) {
	var output mediapackage.CreateHarvestJobOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediapackage.CreateHarvestJob", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateHarvestJobAsync(ctx workflow.Context, input *mediapackage.CreateHarvestJobInput) *CreateHarvestJobFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediapackage.CreateHarvestJob", input)
	return &CreateHarvestJobFuture{Future: future}
}

func (a *stub) CreateOriginEndpoint(ctx workflow.Context, input *mediapackage.CreateOriginEndpointInput) (*mediapackage.CreateOriginEndpointOutput, error) {
	var output mediapackage.CreateOriginEndpointOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediapackage.CreateOriginEndpoint", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateOriginEndpointAsync(ctx workflow.Context, input *mediapackage.CreateOriginEndpointInput) *CreateOriginEndpointFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediapackage.CreateOriginEndpoint", input)
	return &CreateOriginEndpointFuture{Future: future}
}

func (a *stub) DeleteChannel(ctx workflow.Context, input *mediapackage.DeleteChannelInput) (*mediapackage.DeleteChannelOutput, error) {
	var output mediapackage.DeleteChannelOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediapackage.DeleteChannel", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteChannelAsync(ctx workflow.Context, input *mediapackage.DeleteChannelInput) *DeleteChannelFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediapackage.DeleteChannel", input)
	return &DeleteChannelFuture{Future: future}
}

func (a *stub) DeleteOriginEndpoint(ctx workflow.Context, input *mediapackage.DeleteOriginEndpointInput) (*mediapackage.DeleteOriginEndpointOutput, error) {
	var output mediapackage.DeleteOriginEndpointOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediapackage.DeleteOriginEndpoint", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteOriginEndpointAsync(ctx workflow.Context, input *mediapackage.DeleteOriginEndpointInput) *DeleteOriginEndpointFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediapackage.DeleteOriginEndpoint", input)
	return &DeleteOriginEndpointFuture{Future: future}
}

func (a *stub) DescribeChannel(ctx workflow.Context, input *mediapackage.DescribeChannelInput) (*mediapackage.DescribeChannelOutput, error) {
	var output mediapackage.DescribeChannelOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediapackage.DescribeChannel", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeChannelAsync(ctx workflow.Context, input *mediapackage.DescribeChannelInput) *DescribeChannelFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediapackage.DescribeChannel", input)
	return &DescribeChannelFuture{Future: future}
}

func (a *stub) DescribeHarvestJob(ctx workflow.Context, input *mediapackage.DescribeHarvestJobInput) (*mediapackage.DescribeHarvestJobOutput, error) {
	var output mediapackage.DescribeHarvestJobOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediapackage.DescribeHarvestJob", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeHarvestJobAsync(ctx workflow.Context, input *mediapackage.DescribeHarvestJobInput) *DescribeHarvestJobFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediapackage.DescribeHarvestJob", input)
	return &DescribeHarvestJobFuture{Future: future}
}

func (a *stub) DescribeOriginEndpoint(ctx workflow.Context, input *mediapackage.DescribeOriginEndpointInput) (*mediapackage.DescribeOriginEndpointOutput, error) {
	var output mediapackage.DescribeOriginEndpointOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediapackage.DescribeOriginEndpoint", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeOriginEndpointAsync(ctx workflow.Context, input *mediapackage.DescribeOriginEndpointInput) *DescribeOriginEndpointFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediapackage.DescribeOriginEndpoint", input)
	return &DescribeOriginEndpointFuture{Future: future}
}

func (a *stub) ListChannels(ctx workflow.Context, input *mediapackage.ListChannelsInput) (*mediapackage.ListChannelsOutput, error) {
	var output mediapackage.ListChannelsOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediapackage.ListChannels", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListChannelsAsync(ctx workflow.Context, input *mediapackage.ListChannelsInput) *ListChannelsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediapackage.ListChannels", input)
	return &ListChannelsFuture{Future: future}
}

func (a *stub) ListHarvestJobs(ctx workflow.Context, input *mediapackage.ListHarvestJobsInput) (*mediapackage.ListHarvestJobsOutput, error) {
	var output mediapackage.ListHarvestJobsOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediapackage.ListHarvestJobs", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListHarvestJobsAsync(ctx workflow.Context, input *mediapackage.ListHarvestJobsInput) *ListHarvestJobsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediapackage.ListHarvestJobs", input)
	return &ListHarvestJobsFuture{Future: future}
}

func (a *stub) ListOriginEndpoints(ctx workflow.Context, input *mediapackage.ListOriginEndpointsInput) (*mediapackage.ListOriginEndpointsOutput, error) {
	var output mediapackage.ListOriginEndpointsOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediapackage.ListOriginEndpoints", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListOriginEndpointsAsync(ctx workflow.Context, input *mediapackage.ListOriginEndpointsInput) *ListOriginEndpointsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediapackage.ListOriginEndpoints", input)
	return &ListOriginEndpointsFuture{Future: future}
}

func (a *stub) ListTagsForResource(ctx workflow.Context, input *mediapackage.ListTagsForResourceInput) (*mediapackage.ListTagsForResourceOutput, error) {
	var output mediapackage.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediapackage.ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTagsForResourceAsync(ctx workflow.Context, input *mediapackage.ListTagsForResourceInput) *ListTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediapackage.ListTagsForResource", input)
	return &ListTagsForResourceFuture{Future: future}
}

func (a *stub) RotateChannelCredentials(ctx workflow.Context, input *mediapackage.RotateChannelCredentialsInput) (*mediapackage.RotateChannelCredentialsOutput, error) {
	var output mediapackage.RotateChannelCredentialsOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediapackage.RotateChannelCredentials", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) RotateChannelCredentialsAsync(ctx workflow.Context, input *mediapackage.RotateChannelCredentialsInput) *RotateChannelCredentialsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediapackage.RotateChannelCredentials", input)
	return &RotateChannelCredentialsFuture{Future: future}
}

func (a *stub) RotateIngestEndpointCredentials(ctx workflow.Context, input *mediapackage.RotateIngestEndpointCredentialsInput) (*mediapackage.RotateIngestEndpointCredentialsOutput, error) {
	var output mediapackage.RotateIngestEndpointCredentialsOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediapackage.RotateIngestEndpointCredentials", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) RotateIngestEndpointCredentialsAsync(ctx workflow.Context, input *mediapackage.RotateIngestEndpointCredentialsInput) *RotateIngestEndpointCredentialsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediapackage.RotateIngestEndpointCredentials", input)
	return &RotateIngestEndpointCredentialsFuture{Future: future}
}

func (a *stub) TagResource(ctx workflow.Context, input *mediapackage.TagResourceInput) (*mediapackage.TagResourceOutput, error) {
	var output mediapackage.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediapackage.TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TagResourceAsync(ctx workflow.Context, input *mediapackage.TagResourceInput) *TagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediapackage.TagResource", input)
	return &TagResourceFuture{Future: future}
}

func (a *stub) UntagResource(ctx workflow.Context, input *mediapackage.UntagResourceInput) (*mediapackage.UntagResourceOutput, error) {
	var output mediapackage.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediapackage.UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UntagResourceAsync(ctx workflow.Context, input *mediapackage.UntagResourceInput) *UntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediapackage.UntagResource", input)
	return &UntagResourceFuture{Future: future}
}

func (a *stub) UpdateChannel(ctx workflow.Context, input *mediapackage.UpdateChannelInput) (*mediapackage.UpdateChannelOutput, error) {
	var output mediapackage.UpdateChannelOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediapackage.UpdateChannel", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateChannelAsync(ctx workflow.Context, input *mediapackage.UpdateChannelInput) *UpdateChannelFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediapackage.UpdateChannel", input)
	return &UpdateChannelFuture{Future: future}
}

func (a *stub) UpdateOriginEndpoint(ctx workflow.Context, input *mediapackage.UpdateOriginEndpointInput) (*mediapackage.UpdateOriginEndpointOutput, error) {
	var output mediapackage.UpdateOriginEndpointOutput
	err := workflow.ExecuteActivity(ctx, "aws.mediapackage.UpdateOriginEndpoint", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UpdateOriginEndpointAsync(ctx workflow.Context, input *mediapackage.UpdateOriginEndpointInput) *UpdateOriginEndpointFuture {
	future := workflow.ExecuteActivity(ctx, "aws.mediapackage.UpdateOriginEndpoint", input)
	return &UpdateOriginEndpointFuture{Future: future}
}