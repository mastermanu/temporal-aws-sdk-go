package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/ebs"
	"go.temporal.io/sdk/workflow"
)

type EBSClient interface {
    CompleteSnapshot(ctx workflow.Context, input *ebs.CompleteSnapshotInput) (*ebs.CompleteSnapshotOutput, error)
    CompleteSnapshotAsync(ctx workflow.Context, input *ebs.CompleteSnapshotInput) *EbsCompleteSnapshotResult

    GetSnapshotBlock(ctx workflow.Context, input *ebs.GetSnapshotBlockInput) (*ebs.GetSnapshotBlockOutput, error)
    GetSnapshotBlockAsync(ctx workflow.Context, input *ebs.GetSnapshotBlockInput) *EbsGetSnapshotBlockResult

    ListChangedBlocks(ctx workflow.Context, input *ebs.ListChangedBlocksInput) (*ebs.ListChangedBlocksOutput, error)
    ListChangedBlocksAsync(ctx workflow.Context, input *ebs.ListChangedBlocksInput) *EbsListChangedBlocksResult

    ListSnapshotBlocks(ctx workflow.Context, input *ebs.ListSnapshotBlocksInput) (*ebs.ListSnapshotBlocksOutput, error)
    ListSnapshotBlocksAsync(ctx workflow.Context, input *ebs.ListSnapshotBlocksInput) *EbsListSnapshotBlocksResult

    PutSnapshotBlock(ctx workflow.Context, input *ebs.PutSnapshotBlockInput) (*ebs.PutSnapshotBlockOutput, error)
    PutSnapshotBlockAsync(ctx workflow.Context, input *ebs.PutSnapshotBlockInput) *EbsPutSnapshotBlockResult

    StartSnapshot(ctx workflow.Context, input *ebs.StartSnapshotInput) (*ebs.StartSnapshotOutput, error)
    StartSnapshotAsync(ctx workflow.Context, input *ebs.StartSnapshotInput) *EbsStartSnapshotResult
}
type EbsCompleteSnapshotResult struct {
	Result workflow.Future
}

func (r *EbsCompleteSnapshotResult) Get(ctx workflow.Context) (*ebs.CompleteSnapshotOutput, error) {
    var output ebs.CompleteSnapshotOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EbsGetSnapshotBlockResult struct {
	Result workflow.Future
}

func (r *EbsGetSnapshotBlockResult) Get(ctx workflow.Context) (*ebs.GetSnapshotBlockOutput, error) {
    var output ebs.GetSnapshotBlockOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EbsListChangedBlocksResult struct {
	Result workflow.Future
}

func (r *EbsListChangedBlocksResult) Get(ctx workflow.Context) (*ebs.ListChangedBlocksOutput, error) {
    var output ebs.ListChangedBlocksOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EbsListSnapshotBlocksResult struct {
	Result workflow.Future
}

func (r *EbsListSnapshotBlocksResult) Get(ctx workflow.Context) (*ebs.ListSnapshotBlocksOutput, error) {
    var output ebs.ListSnapshotBlocksOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EbsPutSnapshotBlockResult struct {
	Result workflow.Future
}

func (r *EbsPutSnapshotBlockResult) Get(ctx workflow.Context) (*ebs.PutSnapshotBlockOutput, error) {
    var output ebs.PutSnapshotBlockOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type EbsStartSnapshotResult struct {
	Result workflow.Future
}

func (r *EbsStartSnapshotResult) Get(ctx workflow.Context) (*ebs.StartSnapshotOutput, error) {
    var output ebs.StartSnapshotOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type EBSStub struct {
    activities EBSClient
}

func NewEBSStub() EBSClient {
    return &EBSStub{}
}

func (a *EBSStub) CompleteSnapshot(ctx workflow.Context, input *ebs.CompleteSnapshotInput) (*ebs.CompleteSnapshotOutput, error) {
    var output ebs.CompleteSnapshotOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CompleteSnapshot, input).Get(ctx, &output)
    return &output, err
}

func (a *EBSStub) GetSnapshotBlock(ctx workflow.Context, input *ebs.GetSnapshotBlockInput) (*ebs.GetSnapshotBlockOutput, error) {
    var output ebs.GetSnapshotBlockOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetSnapshotBlock, input).Get(ctx, &output)
    return &output, err
}

func (a *EBSStub) ListChangedBlocks(ctx workflow.Context, input *ebs.ListChangedBlocksInput) (*ebs.ListChangedBlocksOutput, error) {
    var output ebs.ListChangedBlocksOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListChangedBlocks, input).Get(ctx, &output)
    return &output, err
}

func (a *EBSStub) ListSnapshotBlocks(ctx workflow.Context, input *ebs.ListSnapshotBlocksInput) (*ebs.ListSnapshotBlocksOutput, error) {
    var output ebs.ListSnapshotBlocksOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListSnapshotBlocks, input).Get(ctx, &output)
    return &output, err
}

func (a *EBSStub) PutSnapshotBlock(ctx workflow.Context, input *ebs.PutSnapshotBlockInput) (*ebs.PutSnapshotBlockOutput, error) {
    var output ebs.PutSnapshotBlockOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutSnapshotBlock, input).Get(ctx, &output)
    return &output, err
}

func (a *EBSStub) StartSnapshot(ctx workflow.Context, input *ebs.StartSnapshotInput) (*ebs.StartSnapshotOutput, error) {
    var output ebs.StartSnapshotOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartSnapshot, input).Get(ctx, &output)
    return &output, err
}