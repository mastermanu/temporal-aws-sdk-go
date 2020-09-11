package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/amplify"
	"github.com/aws/aws-sdk-go/service/amplify/amplifyiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type AmplifyActivities struct {
	client amplifyiface.AmplifyAPI
}

func NewAmplifyActivities(session *session.Session, config ...*aws.Config) *AmplifyActivities {
	client := amplify.New(session, config...)
	return &AmplifyActivities{client: client}
}

func (a *AmplifyActivities) CreateApp(ctx context.Context, input *amplify.CreateAppInput) (*amplify.CreateAppOutput, error) {
	return a.client.CreateAppWithContext(ctx, input)
}

func (a *AmplifyActivities) CreateBackendEnvironment(ctx context.Context, input *amplify.CreateBackendEnvironmentInput) (*amplify.CreateBackendEnvironmentOutput, error) {
	return a.client.CreateBackendEnvironmentWithContext(ctx, input)
}

func (a *AmplifyActivities) CreateBranch(ctx context.Context, input *amplify.CreateBranchInput) (*amplify.CreateBranchOutput, error) {
	return a.client.CreateBranchWithContext(ctx, input)
}

func (a *AmplifyActivities) CreateDeployment(ctx context.Context, input *amplify.CreateDeploymentInput) (*amplify.CreateDeploymentOutput, error) {
	return a.client.CreateDeploymentWithContext(ctx, input)
}

func (a *AmplifyActivities) CreateDomainAssociation(ctx context.Context, input *amplify.CreateDomainAssociationInput) (*amplify.CreateDomainAssociationOutput, error) {
	return a.client.CreateDomainAssociationWithContext(ctx, input)
}

func (a *AmplifyActivities) CreateWebhook(ctx context.Context, input *amplify.CreateWebhookInput) (*amplify.CreateWebhookOutput, error) {
	return a.client.CreateWebhookWithContext(ctx, input)
}

func (a *AmplifyActivities) DeleteApp(ctx context.Context, input *amplify.DeleteAppInput) (*amplify.DeleteAppOutput, error) {
	return a.client.DeleteAppWithContext(ctx, input)
}

func (a *AmplifyActivities) DeleteBackendEnvironment(ctx context.Context, input *amplify.DeleteBackendEnvironmentInput) (*amplify.DeleteBackendEnvironmentOutput, error) {
	return a.client.DeleteBackendEnvironmentWithContext(ctx, input)
}

func (a *AmplifyActivities) DeleteBranch(ctx context.Context, input *amplify.DeleteBranchInput) (*amplify.DeleteBranchOutput, error) {
	return a.client.DeleteBranchWithContext(ctx, input)
}

func (a *AmplifyActivities) DeleteDomainAssociation(ctx context.Context, input *amplify.DeleteDomainAssociationInput) (*amplify.DeleteDomainAssociationOutput, error) {
	return a.client.DeleteDomainAssociationWithContext(ctx, input)
}

func (a *AmplifyActivities) DeleteJob(ctx context.Context, input *amplify.DeleteJobInput) (*amplify.DeleteJobOutput, error) {
	return a.client.DeleteJobWithContext(ctx, input)
}

func (a *AmplifyActivities) DeleteWebhook(ctx context.Context, input *amplify.DeleteWebhookInput) (*amplify.DeleteWebhookOutput, error) {
	return a.client.DeleteWebhookWithContext(ctx, input)
}

func (a *AmplifyActivities) GenerateAccessLogs(ctx context.Context, input *amplify.GenerateAccessLogsInput) (*amplify.GenerateAccessLogsOutput, error) {
	return a.client.GenerateAccessLogsWithContext(ctx, input)
}

func (a *AmplifyActivities) GetApp(ctx context.Context, input *amplify.GetAppInput) (*amplify.GetAppOutput, error) {
	return a.client.GetAppWithContext(ctx, input)
}

func (a *AmplifyActivities) GetArtifactUrl(ctx context.Context, input *amplify.GetArtifactUrlInput) (*amplify.GetArtifactUrlOutput, error) {
	return a.client.GetArtifactUrlWithContext(ctx, input)
}

func (a *AmplifyActivities) GetBackendEnvironment(ctx context.Context, input *amplify.GetBackendEnvironmentInput) (*amplify.GetBackendEnvironmentOutput, error) {
	return a.client.GetBackendEnvironmentWithContext(ctx, input)
}

func (a *AmplifyActivities) GetBranch(ctx context.Context, input *amplify.GetBranchInput) (*amplify.GetBranchOutput, error) {
	return a.client.GetBranchWithContext(ctx, input)
}

func (a *AmplifyActivities) GetDomainAssociation(ctx context.Context, input *amplify.GetDomainAssociationInput) (*amplify.GetDomainAssociationOutput, error) {
	return a.client.GetDomainAssociationWithContext(ctx, input)
}

func (a *AmplifyActivities) GetJob(ctx context.Context, input *amplify.GetJobInput) (*amplify.GetJobOutput, error) {
	return a.client.GetJobWithContext(ctx, input)
}

func (a *AmplifyActivities) GetWebhook(ctx context.Context, input *amplify.GetWebhookInput) (*amplify.GetWebhookOutput, error) {
	return a.client.GetWebhookWithContext(ctx, input)
}

func (a *AmplifyActivities) ListApps(ctx context.Context, input *amplify.ListAppsInput) (*amplify.ListAppsOutput, error) {
	return a.client.ListAppsWithContext(ctx, input)
}

func (a *AmplifyActivities) ListArtifacts(ctx context.Context, input *amplify.ListArtifactsInput) (*amplify.ListArtifactsOutput, error) {
	return a.client.ListArtifactsWithContext(ctx, input)
}

func (a *AmplifyActivities) ListBackendEnvironments(ctx context.Context, input *amplify.ListBackendEnvironmentsInput) (*amplify.ListBackendEnvironmentsOutput, error) {
	return a.client.ListBackendEnvironmentsWithContext(ctx, input)
}

func (a *AmplifyActivities) ListBranches(ctx context.Context, input *amplify.ListBranchesInput) (*amplify.ListBranchesOutput, error) {
	return a.client.ListBranchesWithContext(ctx, input)
}

func (a *AmplifyActivities) ListDomainAssociations(ctx context.Context, input *amplify.ListDomainAssociationsInput) (*amplify.ListDomainAssociationsOutput, error) {
	return a.client.ListDomainAssociationsWithContext(ctx, input)
}

func (a *AmplifyActivities) ListJobs(ctx context.Context, input *amplify.ListJobsInput) (*amplify.ListJobsOutput, error) {
	return a.client.ListJobsWithContext(ctx, input)
}

func (a *AmplifyActivities) ListTagsForResource(ctx context.Context, input *amplify.ListTagsForResourceInput) (*amplify.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *AmplifyActivities) ListWebhooks(ctx context.Context, input *amplify.ListWebhooksInput) (*amplify.ListWebhooksOutput, error) {
	return a.client.ListWebhooksWithContext(ctx, input)
}

func (a *AmplifyActivities) StartDeployment(ctx context.Context, input *amplify.StartDeploymentInput) (*amplify.StartDeploymentOutput, error) {
	return a.client.StartDeploymentWithContext(ctx, input)
}

func (a *AmplifyActivities) StartJob(ctx context.Context, input *amplify.StartJobInput) (*amplify.StartJobOutput, error) {
	return a.client.StartJobWithContext(ctx, input)
}

func (a *AmplifyActivities) StopJob(ctx context.Context, input *amplify.StopJobInput) (*amplify.StopJobOutput, error) {
	return a.client.StopJobWithContext(ctx, input)
}

func (a *AmplifyActivities) TagResource(ctx context.Context, input *amplify.TagResourceInput) (*amplify.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *AmplifyActivities) UntagResource(ctx context.Context, input *amplify.UntagResourceInput) (*amplify.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}

func (a *AmplifyActivities) UpdateApp(ctx context.Context, input *amplify.UpdateAppInput) (*amplify.UpdateAppOutput, error) {
	return a.client.UpdateAppWithContext(ctx, input)
}

func (a *AmplifyActivities) UpdateBranch(ctx context.Context, input *amplify.UpdateBranchInput) (*amplify.UpdateBranchOutput, error) {
	return a.client.UpdateBranchWithContext(ctx, input)
}

func (a *AmplifyActivities) UpdateDomainAssociation(ctx context.Context, input *amplify.UpdateDomainAssociationInput) (*amplify.UpdateDomainAssociationOutput, error) {
	return a.client.UpdateDomainAssociationWithContext(ctx, input)
}

func (a *AmplifyActivities) UpdateWebhook(ctx context.Context, input *amplify.UpdateWebhookInput) (*amplify.UpdateWebhookOutput, error) {
	return a.client.UpdateWebhookWithContext(ctx, input)
}
