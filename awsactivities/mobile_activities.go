package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/mobile"
	"github.com/aws/aws-sdk-go/service/mobile/mobileiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type MobileActivities struct {
	client mobileiface.MobileAPI
}

func NewMobileActivities(session *session.Session, config ...*aws.Config) *MobileActivities {
	client := mobile.New(session, config...)
	return &MobileActivities{client: client}
}

func (a *MobileActivities) CreateProject(ctx context.Context, input *mobile.CreateProjectInput) (*mobile.CreateProjectOutput, error) {
	return a.client.CreateProjectWithContext(ctx, input)
}

func (a *MobileActivities) DeleteProject(ctx context.Context, input *mobile.DeleteProjectInput) (*mobile.DeleteProjectOutput, error) {
	return a.client.DeleteProjectWithContext(ctx, input)
}

func (a *MobileActivities) DescribeBundle(ctx context.Context, input *mobile.DescribeBundleInput) (*mobile.DescribeBundleOutput, error) {
	return a.client.DescribeBundleWithContext(ctx, input)
}

func (a *MobileActivities) DescribeProject(ctx context.Context, input *mobile.DescribeProjectInput) (*mobile.DescribeProjectOutput, error) {
	return a.client.DescribeProjectWithContext(ctx, input)
}

func (a *MobileActivities) ExportBundle(ctx context.Context, input *mobile.ExportBundleInput) (*mobile.ExportBundleOutput, error) {
	return a.client.ExportBundleWithContext(ctx, input)
}

func (a *MobileActivities) ExportProject(ctx context.Context, input *mobile.ExportProjectInput) (*mobile.ExportProjectOutput, error) {
	return a.client.ExportProjectWithContext(ctx, input)
}

func (a *MobileActivities) ListBundles(ctx context.Context, input *mobile.ListBundlesInput) (*mobile.ListBundlesOutput, error) {
	return a.client.ListBundlesWithContext(ctx, input)
}

func (a *MobileActivities) ListProjects(ctx context.Context, input *mobile.ListProjectsInput) (*mobile.ListProjectsOutput, error) {
	return a.client.ListProjectsWithContext(ctx, input)
}

func (a *MobileActivities) UpdateProject(ctx context.Context, input *mobile.UpdateProjectInput) (*mobile.UpdateProjectOutput, error) {
	return a.client.UpdateProjectWithContext(ctx, input)
}
