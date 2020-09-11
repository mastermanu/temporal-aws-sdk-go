package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/migrationhubconfig"
	"github.com/aws/aws-sdk-go/service/migrationhubconfig/migrationhubconfigiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type MigrationHubConfigActivities struct {
	client migrationhubconfigiface.MigrationHubConfigAPI
}

func NewMigrationHubConfigActivities(session *session.Session, config ...*aws.Config) *MigrationHubConfigActivities {
	client := migrationhubconfig.New(session, config...)
	return &MigrationHubConfigActivities{client: client}
}

func (a *MigrationHubConfigActivities) CreateHomeRegionControl(ctx context.Context, input *migrationhubconfig.CreateHomeRegionControlInput) (*migrationhubconfig.CreateHomeRegionControlOutput, error) {
	return a.client.CreateHomeRegionControlWithContext(ctx, input)
}

func (a *MigrationHubConfigActivities) DescribeHomeRegionControls(ctx context.Context, input *migrationhubconfig.DescribeHomeRegionControlsInput) (*migrationhubconfig.DescribeHomeRegionControlsOutput, error) {
	return a.client.DescribeHomeRegionControlsWithContext(ctx, input)
}

func (a *MigrationHubConfigActivities) GetHomeRegion(ctx context.Context, input *migrationhubconfig.GetHomeRegionInput) (*migrationhubconfig.GetHomeRegionOutput, error) {
	return a.client.GetHomeRegionWithContext(ctx, input)
}
