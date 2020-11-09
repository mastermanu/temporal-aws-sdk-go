// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package ioteventsdatastub

import (
	"github.com/aws/aws-sdk-go/service/ioteventsdata"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	BatchPutMessage(ctx workflow.Context, input *ioteventsdata.BatchPutMessageInput) (*ioteventsdata.BatchPutMessageOutput, error)
	BatchPutMessageAsync(ctx workflow.Context, input *ioteventsdata.BatchPutMessageInput) *BatchPutMessageFuture

	BatchUpdateDetector(ctx workflow.Context, input *ioteventsdata.BatchUpdateDetectorInput) (*ioteventsdata.BatchUpdateDetectorOutput, error)
	BatchUpdateDetectorAsync(ctx workflow.Context, input *ioteventsdata.BatchUpdateDetectorInput) *BatchUpdateDetectorFuture

	DescribeDetector(ctx workflow.Context, input *ioteventsdata.DescribeDetectorInput) (*ioteventsdata.DescribeDetectorOutput, error)
	DescribeDetectorAsync(ctx workflow.Context, input *ioteventsdata.DescribeDetectorInput) *DescribeDetectorFuture

	ListDetectors(ctx workflow.Context, input *ioteventsdata.ListDetectorsInput) (*ioteventsdata.ListDetectorsOutput, error)
	ListDetectorsAsync(ctx workflow.Context, input *ioteventsdata.ListDetectorsInput) *ListDetectorsFuture
}

func NewClient() Client {
	return &stub{}
}