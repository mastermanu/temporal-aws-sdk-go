// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package comprehendmedicalstub

import (
	"github.com/aws/aws-sdk-go/service/comprehendmedical"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	DescribeEntitiesDetectionV2Job(ctx workflow.Context, input *comprehendmedical.DescribeEntitiesDetectionV2JobInput) (*comprehendmedical.DescribeEntitiesDetectionV2JobOutput, error)
	DescribeEntitiesDetectionV2JobAsync(ctx workflow.Context, input *comprehendmedical.DescribeEntitiesDetectionV2JobInput) *DescribeEntitiesDetectionV2JobFuture

	DescribeICD10CMInferenceJob(ctx workflow.Context, input *comprehendmedical.DescribeICD10CMInferenceJobInput) (*comprehendmedical.DescribeICD10CMInferenceJobOutput, error)
	DescribeICD10CMInferenceJobAsync(ctx workflow.Context, input *comprehendmedical.DescribeICD10CMInferenceJobInput) *DescribeICD10CMInferenceJobFuture

	DescribePHIDetectionJob(ctx workflow.Context, input *comprehendmedical.DescribePHIDetectionJobInput) (*comprehendmedical.DescribePHIDetectionJobOutput, error)
	DescribePHIDetectionJobAsync(ctx workflow.Context, input *comprehendmedical.DescribePHIDetectionJobInput) *DescribePHIDetectionJobFuture

	DescribeRxNormInferenceJob(ctx workflow.Context, input *comprehendmedical.DescribeRxNormInferenceJobInput) (*comprehendmedical.DescribeRxNormInferenceJobOutput, error)
	DescribeRxNormInferenceJobAsync(ctx workflow.Context, input *comprehendmedical.DescribeRxNormInferenceJobInput) *DescribeRxNormInferenceJobFuture

	DetectEntities(ctx workflow.Context, input *comprehendmedical.DetectEntitiesInput) (*comprehendmedical.DetectEntitiesOutput, error)
	DetectEntitiesAsync(ctx workflow.Context, input *comprehendmedical.DetectEntitiesInput) *DetectEntitiesFuture

	DetectEntitiesV2(ctx workflow.Context, input *comprehendmedical.DetectEntitiesV2Input) (*comprehendmedical.DetectEntitiesV2Output, error)
	DetectEntitiesV2Async(ctx workflow.Context, input *comprehendmedical.DetectEntitiesV2Input) *DetectEntitiesV2Future

	DetectPHI(ctx workflow.Context, input *comprehendmedical.DetectPHIInput) (*comprehendmedical.DetectPHIOutput, error)
	DetectPHIAsync(ctx workflow.Context, input *comprehendmedical.DetectPHIInput) *DetectPHIFuture

	InferICD10CM(ctx workflow.Context, input *comprehendmedical.InferICD10CMInput) (*comprehendmedical.InferICD10CMOutput, error)
	InferICD10CMAsync(ctx workflow.Context, input *comprehendmedical.InferICD10CMInput) *InferICD10CMFuture

	InferRxNorm(ctx workflow.Context, input *comprehendmedical.InferRxNormInput) (*comprehendmedical.InferRxNormOutput, error)
	InferRxNormAsync(ctx workflow.Context, input *comprehendmedical.InferRxNormInput) *InferRxNormFuture

	ListEntitiesDetectionV2Jobs(ctx workflow.Context, input *comprehendmedical.ListEntitiesDetectionV2JobsInput) (*comprehendmedical.ListEntitiesDetectionV2JobsOutput, error)
	ListEntitiesDetectionV2JobsAsync(ctx workflow.Context, input *comprehendmedical.ListEntitiesDetectionV2JobsInput) *ListEntitiesDetectionV2JobsFuture

	ListICD10CMInferenceJobs(ctx workflow.Context, input *comprehendmedical.ListICD10CMInferenceJobsInput) (*comprehendmedical.ListICD10CMInferenceJobsOutput, error)
	ListICD10CMInferenceJobsAsync(ctx workflow.Context, input *comprehendmedical.ListICD10CMInferenceJobsInput) *ListICD10CMInferenceJobsFuture

	ListPHIDetectionJobs(ctx workflow.Context, input *comprehendmedical.ListPHIDetectionJobsInput) (*comprehendmedical.ListPHIDetectionJobsOutput, error)
	ListPHIDetectionJobsAsync(ctx workflow.Context, input *comprehendmedical.ListPHIDetectionJobsInput) *ListPHIDetectionJobsFuture

	ListRxNormInferenceJobs(ctx workflow.Context, input *comprehendmedical.ListRxNormInferenceJobsInput) (*comprehendmedical.ListRxNormInferenceJobsOutput, error)
	ListRxNormInferenceJobsAsync(ctx workflow.Context, input *comprehendmedical.ListRxNormInferenceJobsInput) *ListRxNormInferenceJobsFuture

	StartEntitiesDetectionV2Job(ctx workflow.Context, input *comprehendmedical.StartEntitiesDetectionV2JobInput) (*comprehendmedical.StartEntitiesDetectionV2JobOutput, error)
	StartEntitiesDetectionV2JobAsync(ctx workflow.Context, input *comprehendmedical.StartEntitiesDetectionV2JobInput) *StartEntitiesDetectionV2JobFuture

	StartICD10CMInferenceJob(ctx workflow.Context, input *comprehendmedical.StartICD10CMInferenceJobInput) (*comprehendmedical.StartICD10CMInferenceJobOutput, error)
	StartICD10CMInferenceJobAsync(ctx workflow.Context, input *comprehendmedical.StartICD10CMInferenceJobInput) *StartICD10CMInferenceJobFuture

	StartPHIDetectionJob(ctx workflow.Context, input *comprehendmedical.StartPHIDetectionJobInput) (*comprehendmedical.StartPHIDetectionJobOutput, error)
	StartPHIDetectionJobAsync(ctx workflow.Context, input *comprehendmedical.StartPHIDetectionJobInput) *StartPHIDetectionJobFuture

	StartRxNormInferenceJob(ctx workflow.Context, input *comprehendmedical.StartRxNormInferenceJobInput) (*comprehendmedical.StartRxNormInferenceJobOutput, error)
	StartRxNormInferenceJobAsync(ctx workflow.Context, input *comprehendmedical.StartRxNormInferenceJobInput) *StartRxNormInferenceJobFuture

	StopEntitiesDetectionV2Job(ctx workflow.Context, input *comprehendmedical.StopEntitiesDetectionV2JobInput) (*comprehendmedical.StopEntitiesDetectionV2JobOutput, error)
	StopEntitiesDetectionV2JobAsync(ctx workflow.Context, input *comprehendmedical.StopEntitiesDetectionV2JobInput) *StopEntitiesDetectionV2JobFuture

	StopICD10CMInferenceJob(ctx workflow.Context, input *comprehendmedical.StopICD10CMInferenceJobInput) (*comprehendmedical.StopICD10CMInferenceJobOutput, error)
	StopICD10CMInferenceJobAsync(ctx workflow.Context, input *comprehendmedical.StopICD10CMInferenceJobInput) *StopICD10CMInferenceJobFuture

	StopPHIDetectionJob(ctx workflow.Context, input *comprehendmedical.StopPHIDetectionJobInput) (*comprehendmedical.StopPHIDetectionJobOutput, error)
	StopPHIDetectionJobAsync(ctx workflow.Context, input *comprehendmedical.StopPHIDetectionJobInput) *StopPHIDetectionJobFuture

	StopRxNormInferenceJob(ctx workflow.Context, input *comprehendmedical.StopRxNormInferenceJobInput) (*comprehendmedical.StopRxNormInferenceJobOutput, error)
	StopRxNormInferenceJobAsync(ctx workflow.Context, input *comprehendmedical.StopRxNormInferenceJobInput) *StopRxNormInferenceJobFuture
}

func NewClient() Client {
	return &stub{}
}