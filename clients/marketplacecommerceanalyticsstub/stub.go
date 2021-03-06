// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package marketplacecommerceanalyticsstub

import (
	"github.com/aws/aws-sdk-go/service/marketplacecommerceanalytics"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type GenerateDataSetFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GenerateDataSetFuture) Get(ctx workflow.Context) (*marketplacecommerceanalytics.GenerateDataSetOutput, error) {
	var output marketplacecommerceanalytics.GenerateDataSetOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type StartSupportDataExportFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *StartSupportDataExportFuture) Get(ctx workflow.Context) (*marketplacecommerceanalytics.StartSupportDataExportOutput, error) {
	var output marketplacecommerceanalytics.StartSupportDataExportOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) GenerateDataSet(ctx workflow.Context, input *marketplacecommerceanalytics.GenerateDataSetInput) (*marketplacecommerceanalytics.GenerateDataSetOutput, error) {
	var output marketplacecommerceanalytics.GenerateDataSetOutput
	err := workflow.ExecuteActivity(ctx, "aws.marketplacecommerceanalytics.GenerateDataSet", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GenerateDataSetAsync(ctx workflow.Context, input *marketplacecommerceanalytics.GenerateDataSetInput) *GenerateDataSetFuture {
	future := workflow.ExecuteActivity(ctx, "aws.marketplacecommerceanalytics.GenerateDataSet", input)
	return &GenerateDataSetFuture{Future: future}
}

func (a *stub) StartSupportDataExport(ctx workflow.Context, input *marketplacecommerceanalytics.StartSupportDataExportInput) (*marketplacecommerceanalytics.StartSupportDataExportOutput, error) {
	var output marketplacecommerceanalytics.StartSupportDataExportOutput
	err := workflow.ExecuteActivity(ctx, "aws.marketplacecommerceanalytics.StartSupportDataExport", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) StartSupportDataExportAsync(ctx workflow.Context, input *marketplacecommerceanalytics.StartSupportDataExportInput) *StartSupportDataExportFuture {
	future := workflow.ExecuteActivity(ctx, "aws.marketplacecommerceanalytics.StartSupportDataExport", input)
	return &StartSupportDataExportFuture{Future: future}
}
