package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/migrationhubconfig"
	"go.temporal.io/sdk/workflow"
)

type MigrationHubConfigClient interface {
    CreateHomeRegionControl(ctx workflow.Context, input *migrationhubconfig.CreateHomeRegionControlInput) (*migrationhubconfig.CreateHomeRegionControlOutput, error)
    CreateHomeRegionControlAsync(ctx workflow.Context, input *migrationhubconfig.CreateHomeRegionControlInput) *MigrationhubconfigCreateHomeRegionControlResult

    DescribeHomeRegionControls(ctx workflow.Context, input *migrationhubconfig.DescribeHomeRegionControlsInput) (*migrationhubconfig.DescribeHomeRegionControlsOutput, error)
    DescribeHomeRegionControlsAsync(ctx workflow.Context, input *migrationhubconfig.DescribeHomeRegionControlsInput) *MigrationhubconfigDescribeHomeRegionControlsResult

    GetHomeRegion(ctx workflow.Context, input *migrationhubconfig.GetHomeRegionInput) (*migrationhubconfig.GetHomeRegionOutput, error)
    GetHomeRegionAsync(ctx workflow.Context, input *migrationhubconfig.GetHomeRegionInput) *MigrationhubconfigGetHomeRegionResult
}
type MigrationhubconfigCreateHomeRegionControlResult struct {
	Result workflow.Future
}

func (r *MigrationhubconfigCreateHomeRegionControlResult) Get(ctx workflow.Context) (*migrationhubconfig.CreateHomeRegionControlOutput, error) {
    var output migrationhubconfig.CreateHomeRegionControlOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MigrationhubconfigDescribeHomeRegionControlsResult struct {
	Result workflow.Future
}

func (r *MigrationhubconfigDescribeHomeRegionControlsResult) Get(ctx workflow.Context) (*migrationhubconfig.DescribeHomeRegionControlsOutput, error) {
    var output migrationhubconfig.DescribeHomeRegionControlsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MigrationhubconfigGetHomeRegionResult struct {
	Result workflow.Future
}

func (r *MigrationhubconfigGetHomeRegionResult) Get(ctx workflow.Context) (*migrationhubconfig.GetHomeRegionOutput, error) {
    var output migrationhubconfig.GetHomeRegionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type MigrationHubConfigStub struct {
    activities MigrationHubConfigClient
}

func NewMigrationHubConfigStub() MigrationHubConfigClient {
    return &MigrationHubConfigStub{}
}

func (a *MigrationHubConfigStub) CreateHomeRegionControl(ctx workflow.Context, input *migrationhubconfig.CreateHomeRegionControlInput) (*migrationhubconfig.CreateHomeRegionControlOutput, error) {
    var output migrationhubconfig.CreateHomeRegionControlOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateHomeRegionControl, input).Get(ctx, &output)
    return &output, err
}

func (a *MigrationHubConfigStub) DescribeHomeRegionControls(ctx workflow.Context, input *migrationhubconfig.DescribeHomeRegionControlsInput) (*migrationhubconfig.DescribeHomeRegionControlsOutput, error) {
    var output migrationhubconfig.DescribeHomeRegionControlsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeHomeRegionControls, input).Get(ctx, &output)
    return &output, err
}

func (a *MigrationHubConfigStub) GetHomeRegion(ctx workflow.Context, input *migrationhubconfig.GetHomeRegionInput) (*migrationhubconfig.GetHomeRegionOutput, error) {
    var output migrationhubconfig.GetHomeRegionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetHomeRegion, input).Get(ctx, &output)
    return &output, err
}