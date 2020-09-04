package awsclients

import (
	"github.com/aws/aws-sdk-go/service/amplify"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type AmplifyClient interface {
    CreateApp(ctx workflow.Context, input *amplify.CreateAppInput) (*amplify.CreateAppOutput, error)
    CreateAppAsync(ctx workflow.Context, input *amplify.CreateAppInput) *AmplifyCreateAppResult

    CreateBackendEnvironment(ctx workflow.Context, input *amplify.CreateBackendEnvironmentInput) (*amplify.CreateBackendEnvironmentOutput, error)
    CreateBackendEnvironmentAsync(ctx workflow.Context, input *amplify.CreateBackendEnvironmentInput) *AmplifyCreateBackendEnvironmentResult

    CreateBranch(ctx workflow.Context, input *amplify.CreateBranchInput) (*amplify.CreateBranchOutput, error)
    CreateBranchAsync(ctx workflow.Context, input *amplify.CreateBranchInput) *AmplifyCreateBranchResult

    CreateDeployment(ctx workflow.Context, input *amplify.CreateDeploymentInput) (*amplify.CreateDeploymentOutput, error)
    CreateDeploymentAsync(ctx workflow.Context, input *amplify.CreateDeploymentInput) *AmplifyCreateDeploymentResult

    CreateDomainAssociation(ctx workflow.Context, input *amplify.CreateDomainAssociationInput) (*amplify.CreateDomainAssociationOutput, error)
    CreateDomainAssociationAsync(ctx workflow.Context, input *amplify.CreateDomainAssociationInput) *AmplifyCreateDomainAssociationResult

    CreateWebhook(ctx workflow.Context, input *amplify.CreateWebhookInput) (*amplify.CreateWebhookOutput, error)
    CreateWebhookAsync(ctx workflow.Context, input *amplify.CreateWebhookInput) *AmplifyCreateWebhookResult

    DeleteApp(ctx workflow.Context, input *amplify.DeleteAppInput) (*amplify.DeleteAppOutput, error)
    DeleteAppAsync(ctx workflow.Context, input *amplify.DeleteAppInput) *AmplifyDeleteAppResult

    DeleteBackendEnvironment(ctx workflow.Context, input *amplify.DeleteBackendEnvironmentInput) (*amplify.DeleteBackendEnvironmentOutput, error)
    DeleteBackendEnvironmentAsync(ctx workflow.Context, input *amplify.DeleteBackendEnvironmentInput) *AmplifyDeleteBackendEnvironmentResult

    DeleteBranch(ctx workflow.Context, input *amplify.DeleteBranchInput) (*amplify.DeleteBranchOutput, error)
    DeleteBranchAsync(ctx workflow.Context, input *amplify.DeleteBranchInput) *AmplifyDeleteBranchResult

    DeleteDomainAssociation(ctx workflow.Context, input *amplify.DeleteDomainAssociationInput) (*amplify.DeleteDomainAssociationOutput, error)
    DeleteDomainAssociationAsync(ctx workflow.Context, input *amplify.DeleteDomainAssociationInput) *AmplifyDeleteDomainAssociationResult

    DeleteJob(ctx workflow.Context, input *amplify.DeleteJobInput) (*amplify.DeleteJobOutput, error)
    DeleteJobAsync(ctx workflow.Context, input *amplify.DeleteJobInput) *AmplifyDeleteJobResult

    DeleteWebhook(ctx workflow.Context, input *amplify.DeleteWebhookInput) (*amplify.DeleteWebhookOutput, error)
    DeleteWebhookAsync(ctx workflow.Context, input *amplify.DeleteWebhookInput) *AmplifyDeleteWebhookResult

    GenerateAccessLogs(ctx workflow.Context, input *amplify.GenerateAccessLogsInput) (*amplify.GenerateAccessLogsOutput, error)
    GenerateAccessLogsAsync(ctx workflow.Context, input *amplify.GenerateAccessLogsInput) *AmplifyGenerateAccessLogsResult

    GetApp(ctx workflow.Context, input *amplify.GetAppInput) (*amplify.GetAppOutput, error)
    GetAppAsync(ctx workflow.Context, input *amplify.GetAppInput) *AmplifyGetAppResult

    GetArtifactUrl(ctx workflow.Context, input *amplify.GetArtifactUrlInput) (*amplify.GetArtifactUrlOutput, error)
    GetArtifactUrlAsync(ctx workflow.Context, input *amplify.GetArtifactUrlInput) *AmplifyGetArtifactUrlResult

    GetBackendEnvironment(ctx workflow.Context, input *amplify.GetBackendEnvironmentInput) (*amplify.GetBackendEnvironmentOutput, error)
    GetBackendEnvironmentAsync(ctx workflow.Context, input *amplify.GetBackendEnvironmentInput) *AmplifyGetBackendEnvironmentResult

    GetBranch(ctx workflow.Context, input *amplify.GetBranchInput) (*amplify.GetBranchOutput, error)
    GetBranchAsync(ctx workflow.Context, input *amplify.GetBranchInput) *AmplifyGetBranchResult

    GetDomainAssociation(ctx workflow.Context, input *amplify.GetDomainAssociationInput) (*amplify.GetDomainAssociationOutput, error)
    GetDomainAssociationAsync(ctx workflow.Context, input *amplify.GetDomainAssociationInput) *AmplifyGetDomainAssociationResult

    GetJob(ctx workflow.Context, input *amplify.GetJobInput) (*amplify.GetJobOutput, error)
    GetJobAsync(ctx workflow.Context, input *amplify.GetJobInput) *AmplifyGetJobResult

    GetWebhook(ctx workflow.Context, input *amplify.GetWebhookInput) (*amplify.GetWebhookOutput, error)
    GetWebhookAsync(ctx workflow.Context, input *amplify.GetWebhookInput) *AmplifyGetWebhookResult

    ListApps(ctx workflow.Context, input *amplify.ListAppsInput) (*amplify.ListAppsOutput, error)
    ListAppsAsync(ctx workflow.Context, input *amplify.ListAppsInput) *AmplifyListAppsResult

    ListArtifacts(ctx workflow.Context, input *amplify.ListArtifactsInput) (*amplify.ListArtifactsOutput, error)
    ListArtifactsAsync(ctx workflow.Context, input *amplify.ListArtifactsInput) *AmplifyListArtifactsResult

    ListBackendEnvironments(ctx workflow.Context, input *amplify.ListBackendEnvironmentsInput) (*amplify.ListBackendEnvironmentsOutput, error)
    ListBackendEnvironmentsAsync(ctx workflow.Context, input *amplify.ListBackendEnvironmentsInput) *AmplifyListBackendEnvironmentsResult

    ListBranches(ctx workflow.Context, input *amplify.ListBranchesInput) (*amplify.ListBranchesOutput, error)
    ListBranchesAsync(ctx workflow.Context, input *amplify.ListBranchesInput) *AmplifyListBranchesResult

    ListDomainAssociations(ctx workflow.Context, input *amplify.ListDomainAssociationsInput) (*amplify.ListDomainAssociationsOutput, error)
    ListDomainAssociationsAsync(ctx workflow.Context, input *amplify.ListDomainAssociationsInput) *AmplifyListDomainAssociationsResult

    ListJobs(ctx workflow.Context, input *amplify.ListJobsInput) (*amplify.ListJobsOutput, error)
    ListJobsAsync(ctx workflow.Context, input *amplify.ListJobsInput) *AmplifyListJobsResult

    ListTagsForResource(ctx workflow.Context, input *amplify.ListTagsForResourceInput) (*amplify.ListTagsForResourceOutput, error)
    ListTagsForResourceAsync(ctx workflow.Context, input *amplify.ListTagsForResourceInput) *AmplifyListTagsForResourceResult

    ListWebhooks(ctx workflow.Context, input *amplify.ListWebhooksInput) (*amplify.ListWebhooksOutput, error)
    ListWebhooksAsync(ctx workflow.Context, input *amplify.ListWebhooksInput) *AmplifyListWebhooksResult

    StartDeployment(ctx workflow.Context, input *amplify.StartDeploymentInput) (*amplify.StartDeploymentOutput, error)
    StartDeploymentAsync(ctx workflow.Context, input *amplify.StartDeploymentInput) *AmplifyStartDeploymentResult

    StartJob(ctx workflow.Context, input *amplify.StartJobInput) (*amplify.StartJobOutput, error)
    StartJobAsync(ctx workflow.Context, input *amplify.StartJobInput) *AmplifyStartJobResult

    StopJob(ctx workflow.Context, input *amplify.StopJobInput) (*amplify.StopJobOutput, error)
    StopJobAsync(ctx workflow.Context, input *amplify.StopJobInput) *AmplifyStopJobResult

    TagResource(ctx workflow.Context, input *amplify.TagResourceInput) (*amplify.TagResourceOutput, error)
    TagResourceAsync(ctx workflow.Context, input *amplify.TagResourceInput) *AmplifyTagResourceResult

    UntagResource(ctx workflow.Context, input *amplify.UntagResourceInput) (*amplify.UntagResourceOutput, error)
    UntagResourceAsync(ctx workflow.Context, input *amplify.UntagResourceInput) *AmplifyUntagResourceResult

    UpdateApp(ctx workflow.Context, input *amplify.UpdateAppInput) (*amplify.UpdateAppOutput, error)
    UpdateAppAsync(ctx workflow.Context, input *amplify.UpdateAppInput) *AmplifyUpdateAppResult

    UpdateBranch(ctx workflow.Context, input *amplify.UpdateBranchInput) (*amplify.UpdateBranchOutput, error)
    UpdateBranchAsync(ctx workflow.Context, input *amplify.UpdateBranchInput) *AmplifyUpdateBranchResult

    UpdateDomainAssociation(ctx workflow.Context, input *amplify.UpdateDomainAssociationInput) (*amplify.UpdateDomainAssociationOutput, error)
    UpdateDomainAssociationAsync(ctx workflow.Context, input *amplify.UpdateDomainAssociationInput) *AmplifyUpdateDomainAssociationResult

    UpdateWebhook(ctx workflow.Context, input *amplify.UpdateWebhookInput) (*amplify.UpdateWebhookOutput, error)
    UpdateWebhookAsync(ctx workflow.Context, input *amplify.UpdateWebhookInput) *AmplifyUpdateWebhookResult
}

type AmplifyCreateAppResult struct {
	Result workflow.Future
}

func (r *AmplifyCreateAppResult) Get(ctx workflow.Context) (*amplify.CreateAppOutput, error) {
    var output amplify.CreateAppOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AmplifyCreateBackendEnvironmentResult struct {
	Result workflow.Future
}

func (r *AmplifyCreateBackendEnvironmentResult) Get(ctx workflow.Context) (*amplify.CreateBackendEnvironmentOutput, error) {
    var output amplify.CreateBackendEnvironmentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AmplifyCreateBranchResult struct {
	Result workflow.Future
}

func (r *AmplifyCreateBranchResult) Get(ctx workflow.Context) (*amplify.CreateBranchOutput, error) {
    var output amplify.CreateBranchOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AmplifyCreateDeploymentResult struct {
	Result workflow.Future
}

func (r *AmplifyCreateDeploymentResult) Get(ctx workflow.Context) (*amplify.CreateDeploymentOutput, error) {
    var output amplify.CreateDeploymentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AmplifyCreateDomainAssociationResult struct {
	Result workflow.Future
}

func (r *AmplifyCreateDomainAssociationResult) Get(ctx workflow.Context) (*amplify.CreateDomainAssociationOutput, error) {
    var output amplify.CreateDomainAssociationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AmplifyCreateWebhookResult struct {
	Result workflow.Future
}

func (r *AmplifyCreateWebhookResult) Get(ctx workflow.Context) (*amplify.CreateWebhookOutput, error) {
    var output amplify.CreateWebhookOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AmplifyDeleteAppResult struct {
	Result workflow.Future
}

func (r *AmplifyDeleteAppResult) Get(ctx workflow.Context) (*amplify.DeleteAppOutput, error) {
    var output amplify.DeleteAppOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AmplifyDeleteBackendEnvironmentResult struct {
	Result workflow.Future
}

func (r *AmplifyDeleteBackendEnvironmentResult) Get(ctx workflow.Context) (*amplify.DeleteBackendEnvironmentOutput, error) {
    var output amplify.DeleteBackendEnvironmentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AmplifyDeleteBranchResult struct {
	Result workflow.Future
}

func (r *AmplifyDeleteBranchResult) Get(ctx workflow.Context) (*amplify.DeleteBranchOutput, error) {
    var output amplify.DeleteBranchOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AmplifyDeleteDomainAssociationResult struct {
	Result workflow.Future
}

func (r *AmplifyDeleteDomainAssociationResult) Get(ctx workflow.Context) (*amplify.DeleteDomainAssociationOutput, error) {
    var output amplify.DeleteDomainAssociationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AmplifyDeleteJobResult struct {
	Result workflow.Future
}

func (r *AmplifyDeleteJobResult) Get(ctx workflow.Context) (*amplify.DeleteJobOutput, error) {
    var output amplify.DeleteJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AmplifyDeleteWebhookResult struct {
	Result workflow.Future
}

func (r *AmplifyDeleteWebhookResult) Get(ctx workflow.Context) (*amplify.DeleteWebhookOutput, error) {
    var output amplify.DeleteWebhookOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AmplifyGenerateAccessLogsResult struct {
	Result workflow.Future
}

func (r *AmplifyGenerateAccessLogsResult) Get(ctx workflow.Context) (*amplify.GenerateAccessLogsOutput, error) {
    var output amplify.GenerateAccessLogsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AmplifyGetAppResult struct {
	Result workflow.Future
}

func (r *AmplifyGetAppResult) Get(ctx workflow.Context) (*amplify.GetAppOutput, error) {
    var output amplify.GetAppOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AmplifyGetArtifactUrlResult struct {
	Result workflow.Future
}

func (r *AmplifyGetArtifactUrlResult) Get(ctx workflow.Context) (*amplify.GetArtifactUrlOutput, error) {
    var output amplify.GetArtifactUrlOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AmplifyGetBackendEnvironmentResult struct {
	Result workflow.Future
}

func (r *AmplifyGetBackendEnvironmentResult) Get(ctx workflow.Context) (*amplify.GetBackendEnvironmentOutput, error) {
    var output amplify.GetBackendEnvironmentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AmplifyGetBranchResult struct {
	Result workflow.Future
}

func (r *AmplifyGetBranchResult) Get(ctx workflow.Context) (*amplify.GetBranchOutput, error) {
    var output amplify.GetBranchOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AmplifyGetDomainAssociationResult struct {
	Result workflow.Future
}

func (r *AmplifyGetDomainAssociationResult) Get(ctx workflow.Context) (*amplify.GetDomainAssociationOutput, error) {
    var output amplify.GetDomainAssociationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AmplifyGetJobResult struct {
	Result workflow.Future
}

func (r *AmplifyGetJobResult) Get(ctx workflow.Context) (*amplify.GetJobOutput, error) {
    var output amplify.GetJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AmplifyGetWebhookResult struct {
	Result workflow.Future
}

func (r *AmplifyGetWebhookResult) Get(ctx workflow.Context) (*amplify.GetWebhookOutput, error) {
    var output amplify.GetWebhookOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AmplifyListAppsResult struct {
	Result workflow.Future
}

func (r *AmplifyListAppsResult) Get(ctx workflow.Context) (*amplify.ListAppsOutput, error) {
    var output amplify.ListAppsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AmplifyListArtifactsResult struct {
	Result workflow.Future
}

func (r *AmplifyListArtifactsResult) Get(ctx workflow.Context) (*amplify.ListArtifactsOutput, error) {
    var output amplify.ListArtifactsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AmplifyListBackendEnvironmentsResult struct {
	Result workflow.Future
}

func (r *AmplifyListBackendEnvironmentsResult) Get(ctx workflow.Context) (*amplify.ListBackendEnvironmentsOutput, error) {
    var output amplify.ListBackendEnvironmentsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AmplifyListBranchesResult struct {
	Result workflow.Future
}

func (r *AmplifyListBranchesResult) Get(ctx workflow.Context) (*amplify.ListBranchesOutput, error) {
    var output amplify.ListBranchesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AmplifyListDomainAssociationsResult struct {
	Result workflow.Future
}

func (r *AmplifyListDomainAssociationsResult) Get(ctx workflow.Context) (*amplify.ListDomainAssociationsOutput, error) {
    var output amplify.ListDomainAssociationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AmplifyListJobsResult struct {
	Result workflow.Future
}

func (r *AmplifyListJobsResult) Get(ctx workflow.Context) (*amplify.ListJobsOutput, error) {
    var output amplify.ListJobsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AmplifyListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *AmplifyListTagsForResourceResult) Get(ctx workflow.Context) (*amplify.ListTagsForResourceOutput, error) {
    var output amplify.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AmplifyListWebhooksResult struct {
	Result workflow.Future
}

func (r *AmplifyListWebhooksResult) Get(ctx workflow.Context) (*amplify.ListWebhooksOutput, error) {
    var output amplify.ListWebhooksOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AmplifyStartDeploymentResult struct {
	Result workflow.Future
}

func (r *AmplifyStartDeploymentResult) Get(ctx workflow.Context) (*amplify.StartDeploymentOutput, error) {
    var output amplify.StartDeploymentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AmplifyStartJobResult struct {
	Result workflow.Future
}

func (r *AmplifyStartJobResult) Get(ctx workflow.Context) (*amplify.StartJobOutput, error) {
    var output amplify.StartJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AmplifyStopJobResult struct {
	Result workflow.Future
}

func (r *AmplifyStopJobResult) Get(ctx workflow.Context) (*amplify.StopJobOutput, error) {
    var output amplify.StopJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AmplifyTagResourceResult struct {
	Result workflow.Future
}

func (r *AmplifyTagResourceResult) Get(ctx workflow.Context) (*amplify.TagResourceOutput, error) {
    var output amplify.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AmplifyUntagResourceResult struct {
	Result workflow.Future
}

func (r *AmplifyUntagResourceResult) Get(ctx workflow.Context) (*amplify.UntagResourceOutput, error) {
    var output amplify.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AmplifyUpdateAppResult struct {
	Result workflow.Future
}

func (r *AmplifyUpdateAppResult) Get(ctx workflow.Context) (*amplify.UpdateAppOutput, error) {
    var output amplify.UpdateAppOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AmplifyUpdateBranchResult struct {
	Result workflow.Future
}

func (r *AmplifyUpdateBranchResult) Get(ctx workflow.Context) (*amplify.UpdateBranchOutput, error) {
    var output amplify.UpdateBranchOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AmplifyUpdateDomainAssociationResult struct {
	Result workflow.Future
}

func (r *AmplifyUpdateDomainAssociationResult) Get(ctx workflow.Context) (*amplify.UpdateDomainAssociationOutput, error) {
    var output amplify.UpdateDomainAssociationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AmplifyUpdateWebhookResult struct {
	Result workflow.Future
}

func (r *AmplifyUpdateWebhookResult) Get(ctx workflow.Context) (*amplify.UpdateWebhookOutput, error) {
    var output amplify.UpdateWebhookOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AmplifyStub struct {
    activities awsactivities.AmplifyActivities
}

func NewAmplifyStub() AmplifyClient {
    return &AmplifyStub{}
}

func (a *AmplifyStub) CreateApp(ctx workflow.Context, input *amplify.CreateAppInput) (*amplify.CreateAppOutput, error) {
    var output amplify.CreateAppOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateApp, input).Get(ctx, &output)
    return &output, err
}

func (a *AmplifyStub) CreateAppAsync(ctx workflow.Context, input *amplify.CreateAppInput) *AmplifyCreateAppResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateApp, input)
    return &AmplifyCreateAppResult{Result: future}
}

func (a *AmplifyStub) CreateBackendEnvironment(ctx workflow.Context, input *amplify.CreateBackendEnvironmentInput) (*amplify.CreateBackendEnvironmentOutput, error) {
    var output amplify.CreateBackendEnvironmentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateBackendEnvironment, input).Get(ctx, &output)
    return &output, err
}

func (a *AmplifyStub) CreateBackendEnvironmentAsync(ctx workflow.Context, input *amplify.CreateBackendEnvironmentInput) *AmplifyCreateBackendEnvironmentResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateBackendEnvironment, input)
    return &AmplifyCreateBackendEnvironmentResult{Result: future}
}

func (a *AmplifyStub) CreateBranch(ctx workflow.Context, input *amplify.CreateBranchInput) (*amplify.CreateBranchOutput, error) {
    var output amplify.CreateBranchOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateBranch, input).Get(ctx, &output)
    return &output, err
}

func (a *AmplifyStub) CreateBranchAsync(ctx workflow.Context, input *amplify.CreateBranchInput) *AmplifyCreateBranchResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateBranch, input)
    return &AmplifyCreateBranchResult{Result: future}
}

func (a *AmplifyStub) CreateDeployment(ctx workflow.Context, input *amplify.CreateDeploymentInput) (*amplify.CreateDeploymentOutput, error) {
    var output amplify.CreateDeploymentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateDeployment, input).Get(ctx, &output)
    return &output, err
}

func (a *AmplifyStub) CreateDeploymentAsync(ctx workflow.Context, input *amplify.CreateDeploymentInput) *AmplifyCreateDeploymentResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateDeployment, input)
    return &AmplifyCreateDeploymentResult{Result: future}
}

func (a *AmplifyStub) CreateDomainAssociation(ctx workflow.Context, input *amplify.CreateDomainAssociationInput) (*amplify.CreateDomainAssociationOutput, error) {
    var output amplify.CreateDomainAssociationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateDomainAssociation, input).Get(ctx, &output)
    return &output, err
}

func (a *AmplifyStub) CreateDomainAssociationAsync(ctx workflow.Context, input *amplify.CreateDomainAssociationInput) *AmplifyCreateDomainAssociationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateDomainAssociation, input)
    return &AmplifyCreateDomainAssociationResult{Result: future}
}

func (a *AmplifyStub) CreateWebhook(ctx workflow.Context, input *amplify.CreateWebhookInput) (*amplify.CreateWebhookOutput, error) {
    var output amplify.CreateWebhookOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateWebhook, input).Get(ctx, &output)
    return &output, err
}

func (a *AmplifyStub) CreateWebhookAsync(ctx workflow.Context, input *amplify.CreateWebhookInput) *AmplifyCreateWebhookResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateWebhook, input)
    return &AmplifyCreateWebhookResult{Result: future}
}

func (a *AmplifyStub) DeleteApp(ctx workflow.Context, input *amplify.DeleteAppInput) (*amplify.DeleteAppOutput, error) {
    var output amplify.DeleteAppOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteApp, input).Get(ctx, &output)
    return &output, err
}

func (a *AmplifyStub) DeleteAppAsync(ctx workflow.Context, input *amplify.DeleteAppInput) *AmplifyDeleteAppResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteApp, input)
    return &AmplifyDeleteAppResult{Result: future}
}

func (a *AmplifyStub) DeleteBackendEnvironment(ctx workflow.Context, input *amplify.DeleteBackendEnvironmentInput) (*amplify.DeleteBackendEnvironmentOutput, error) {
    var output amplify.DeleteBackendEnvironmentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteBackendEnvironment, input).Get(ctx, &output)
    return &output, err
}

func (a *AmplifyStub) DeleteBackendEnvironmentAsync(ctx workflow.Context, input *amplify.DeleteBackendEnvironmentInput) *AmplifyDeleteBackendEnvironmentResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteBackendEnvironment, input)
    return &AmplifyDeleteBackendEnvironmentResult{Result: future}
}

func (a *AmplifyStub) DeleteBranch(ctx workflow.Context, input *amplify.DeleteBranchInput) (*amplify.DeleteBranchOutput, error) {
    var output amplify.DeleteBranchOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteBranch, input).Get(ctx, &output)
    return &output, err
}

func (a *AmplifyStub) DeleteBranchAsync(ctx workflow.Context, input *amplify.DeleteBranchInput) *AmplifyDeleteBranchResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteBranch, input)
    return &AmplifyDeleteBranchResult{Result: future}
}

func (a *AmplifyStub) DeleteDomainAssociation(ctx workflow.Context, input *amplify.DeleteDomainAssociationInput) (*amplify.DeleteDomainAssociationOutput, error) {
    var output amplify.DeleteDomainAssociationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteDomainAssociation, input).Get(ctx, &output)
    return &output, err
}

func (a *AmplifyStub) DeleteDomainAssociationAsync(ctx workflow.Context, input *amplify.DeleteDomainAssociationInput) *AmplifyDeleteDomainAssociationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteDomainAssociation, input)
    return &AmplifyDeleteDomainAssociationResult{Result: future}
}

func (a *AmplifyStub) DeleteJob(ctx workflow.Context, input *amplify.DeleteJobInput) (*amplify.DeleteJobOutput, error) {
    var output amplify.DeleteJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteJob, input).Get(ctx, &output)
    return &output, err
}

func (a *AmplifyStub) DeleteJobAsync(ctx workflow.Context, input *amplify.DeleteJobInput) *AmplifyDeleteJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteJob, input)
    return &AmplifyDeleteJobResult{Result: future}
}

func (a *AmplifyStub) DeleteWebhook(ctx workflow.Context, input *amplify.DeleteWebhookInput) (*amplify.DeleteWebhookOutput, error) {
    var output amplify.DeleteWebhookOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteWebhook, input).Get(ctx, &output)
    return &output, err
}

func (a *AmplifyStub) DeleteWebhookAsync(ctx workflow.Context, input *amplify.DeleteWebhookInput) *AmplifyDeleteWebhookResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteWebhook, input)
    return &AmplifyDeleteWebhookResult{Result: future}
}

func (a *AmplifyStub) GenerateAccessLogs(ctx workflow.Context, input *amplify.GenerateAccessLogsInput) (*amplify.GenerateAccessLogsOutput, error) {
    var output amplify.GenerateAccessLogsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GenerateAccessLogs, input).Get(ctx, &output)
    return &output, err
}

func (a *AmplifyStub) GenerateAccessLogsAsync(ctx workflow.Context, input *amplify.GenerateAccessLogsInput) *AmplifyGenerateAccessLogsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GenerateAccessLogs, input)
    return &AmplifyGenerateAccessLogsResult{Result: future}
}

func (a *AmplifyStub) GetApp(ctx workflow.Context, input *amplify.GetAppInput) (*amplify.GetAppOutput, error) {
    var output amplify.GetAppOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetApp, input).Get(ctx, &output)
    return &output, err
}

func (a *AmplifyStub) GetAppAsync(ctx workflow.Context, input *amplify.GetAppInput) *AmplifyGetAppResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetApp, input)
    return &AmplifyGetAppResult{Result: future}
}

func (a *AmplifyStub) GetArtifactUrl(ctx workflow.Context, input *amplify.GetArtifactUrlInput) (*amplify.GetArtifactUrlOutput, error) {
    var output amplify.GetArtifactUrlOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetArtifactUrl, input).Get(ctx, &output)
    return &output, err
}

func (a *AmplifyStub) GetArtifactUrlAsync(ctx workflow.Context, input *amplify.GetArtifactUrlInput) *AmplifyGetArtifactUrlResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetArtifactUrl, input)
    return &AmplifyGetArtifactUrlResult{Result: future}
}

func (a *AmplifyStub) GetBackendEnvironment(ctx workflow.Context, input *amplify.GetBackendEnvironmentInput) (*amplify.GetBackendEnvironmentOutput, error) {
    var output amplify.GetBackendEnvironmentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetBackendEnvironment, input).Get(ctx, &output)
    return &output, err
}

func (a *AmplifyStub) GetBackendEnvironmentAsync(ctx workflow.Context, input *amplify.GetBackendEnvironmentInput) *AmplifyGetBackendEnvironmentResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetBackendEnvironment, input)
    return &AmplifyGetBackendEnvironmentResult{Result: future}
}

func (a *AmplifyStub) GetBranch(ctx workflow.Context, input *amplify.GetBranchInput) (*amplify.GetBranchOutput, error) {
    var output amplify.GetBranchOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetBranch, input).Get(ctx, &output)
    return &output, err
}

func (a *AmplifyStub) GetBranchAsync(ctx workflow.Context, input *amplify.GetBranchInput) *AmplifyGetBranchResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetBranch, input)
    return &AmplifyGetBranchResult{Result: future}
}

func (a *AmplifyStub) GetDomainAssociation(ctx workflow.Context, input *amplify.GetDomainAssociationInput) (*amplify.GetDomainAssociationOutput, error) {
    var output amplify.GetDomainAssociationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetDomainAssociation, input).Get(ctx, &output)
    return &output, err
}

func (a *AmplifyStub) GetDomainAssociationAsync(ctx workflow.Context, input *amplify.GetDomainAssociationInput) *AmplifyGetDomainAssociationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetDomainAssociation, input)
    return &AmplifyGetDomainAssociationResult{Result: future}
}

func (a *AmplifyStub) GetJob(ctx workflow.Context, input *amplify.GetJobInput) (*amplify.GetJobOutput, error) {
    var output amplify.GetJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetJob, input).Get(ctx, &output)
    return &output, err
}

func (a *AmplifyStub) GetJobAsync(ctx workflow.Context, input *amplify.GetJobInput) *AmplifyGetJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetJob, input)
    return &AmplifyGetJobResult{Result: future}
}

func (a *AmplifyStub) GetWebhook(ctx workflow.Context, input *amplify.GetWebhookInput) (*amplify.GetWebhookOutput, error) {
    var output amplify.GetWebhookOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetWebhook, input).Get(ctx, &output)
    return &output, err
}

func (a *AmplifyStub) GetWebhookAsync(ctx workflow.Context, input *amplify.GetWebhookInput) *AmplifyGetWebhookResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetWebhook, input)
    return &AmplifyGetWebhookResult{Result: future}
}

func (a *AmplifyStub) ListApps(ctx workflow.Context, input *amplify.ListAppsInput) (*amplify.ListAppsOutput, error) {
    var output amplify.ListAppsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListApps, input).Get(ctx, &output)
    return &output, err
}

func (a *AmplifyStub) ListAppsAsync(ctx workflow.Context, input *amplify.ListAppsInput) *AmplifyListAppsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListApps, input)
    return &AmplifyListAppsResult{Result: future}
}

func (a *AmplifyStub) ListArtifacts(ctx workflow.Context, input *amplify.ListArtifactsInput) (*amplify.ListArtifactsOutput, error) {
    var output amplify.ListArtifactsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListArtifacts, input).Get(ctx, &output)
    return &output, err
}

func (a *AmplifyStub) ListArtifactsAsync(ctx workflow.Context, input *amplify.ListArtifactsInput) *AmplifyListArtifactsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListArtifacts, input)
    return &AmplifyListArtifactsResult{Result: future}
}

func (a *AmplifyStub) ListBackendEnvironments(ctx workflow.Context, input *amplify.ListBackendEnvironmentsInput) (*amplify.ListBackendEnvironmentsOutput, error) {
    var output amplify.ListBackendEnvironmentsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListBackendEnvironments, input).Get(ctx, &output)
    return &output, err
}

func (a *AmplifyStub) ListBackendEnvironmentsAsync(ctx workflow.Context, input *amplify.ListBackendEnvironmentsInput) *AmplifyListBackendEnvironmentsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListBackendEnvironments, input)
    return &AmplifyListBackendEnvironmentsResult{Result: future}
}

func (a *AmplifyStub) ListBranches(ctx workflow.Context, input *amplify.ListBranchesInput) (*amplify.ListBranchesOutput, error) {
    var output amplify.ListBranchesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListBranches, input).Get(ctx, &output)
    return &output, err
}

func (a *AmplifyStub) ListBranchesAsync(ctx workflow.Context, input *amplify.ListBranchesInput) *AmplifyListBranchesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListBranches, input)
    return &AmplifyListBranchesResult{Result: future}
}

func (a *AmplifyStub) ListDomainAssociations(ctx workflow.Context, input *amplify.ListDomainAssociationsInput) (*amplify.ListDomainAssociationsOutput, error) {
    var output amplify.ListDomainAssociationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDomainAssociations, input).Get(ctx, &output)
    return &output, err
}

func (a *AmplifyStub) ListDomainAssociationsAsync(ctx workflow.Context, input *amplify.ListDomainAssociationsInput) *AmplifyListDomainAssociationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListDomainAssociations, input)
    return &AmplifyListDomainAssociationsResult{Result: future}
}

func (a *AmplifyStub) ListJobs(ctx workflow.Context, input *amplify.ListJobsInput) (*amplify.ListJobsOutput, error) {
    var output amplify.ListJobsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListJobs, input).Get(ctx, &output)
    return &output, err
}

func (a *AmplifyStub) ListJobsAsync(ctx workflow.Context, input *amplify.ListJobsInput) *AmplifyListJobsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListJobs, input)
    return &AmplifyListJobsResult{Result: future}
}

func (a *AmplifyStub) ListTagsForResource(ctx workflow.Context, input *amplify.ListTagsForResourceInput) (*amplify.ListTagsForResourceOutput, error) {
    var output amplify.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *AmplifyStub) ListTagsForResourceAsync(ctx workflow.Context, input *amplify.ListTagsForResourceInput) *AmplifyListTagsForResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input)
    return &AmplifyListTagsForResourceResult{Result: future}
}

func (a *AmplifyStub) ListWebhooks(ctx workflow.Context, input *amplify.ListWebhooksInput) (*amplify.ListWebhooksOutput, error) {
    var output amplify.ListWebhooksOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListWebhooks, input).Get(ctx, &output)
    return &output, err
}

func (a *AmplifyStub) ListWebhooksAsync(ctx workflow.Context, input *amplify.ListWebhooksInput) *AmplifyListWebhooksResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListWebhooks, input)
    return &AmplifyListWebhooksResult{Result: future}
}

func (a *AmplifyStub) StartDeployment(ctx workflow.Context, input *amplify.StartDeploymentInput) (*amplify.StartDeploymentOutput, error) {
    var output amplify.StartDeploymentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartDeployment, input).Get(ctx, &output)
    return &output, err
}

func (a *AmplifyStub) StartDeploymentAsync(ctx workflow.Context, input *amplify.StartDeploymentInput) *AmplifyStartDeploymentResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StartDeployment, input)
    return &AmplifyStartDeploymentResult{Result: future}
}

func (a *AmplifyStub) StartJob(ctx workflow.Context, input *amplify.StartJobInput) (*amplify.StartJobOutput, error) {
    var output amplify.StartJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartJob, input).Get(ctx, &output)
    return &output, err
}

func (a *AmplifyStub) StartJobAsync(ctx workflow.Context, input *amplify.StartJobInput) *AmplifyStartJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StartJob, input)
    return &AmplifyStartJobResult{Result: future}
}

func (a *AmplifyStub) StopJob(ctx workflow.Context, input *amplify.StopJobInput) (*amplify.StopJobOutput, error) {
    var output amplify.StopJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StopJob, input).Get(ctx, &output)
    return &output, err
}

func (a *AmplifyStub) StopJobAsync(ctx workflow.Context, input *amplify.StopJobInput) *AmplifyStopJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StopJob, input)
    return &AmplifyStopJobResult{Result: future}
}

func (a *AmplifyStub) TagResource(ctx workflow.Context, input *amplify.TagResourceInput) (*amplify.TagResourceOutput, error) {
    var output amplify.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *AmplifyStub) TagResourceAsync(ctx workflow.Context, input *amplify.TagResourceInput) *AmplifyTagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TagResource, input)
    return &AmplifyTagResourceResult{Result: future}
}

func (a *AmplifyStub) UntagResource(ctx workflow.Context, input *amplify.UntagResourceInput) (*amplify.UntagResourceOutput, error) {
    var output amplify.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *AmplifyStub) UntagResourceAsync(ctx workflow.Context, input *amplify.UntagResourceInput) *AmplifyUntagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input)
    return &AmplifyUntagResourceResult{Result: future}
}

func (a *AmplifyStub) UpdateApp(ctx workflow.Context, input *amplify.UpdateAppInput) (*amplify.UpdateAppOutput, error) {
    var output amplify.UpdateAppOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateApp, input).Get(ctx, &output)
    return &output, err
}

func (a *AmplifyStub) UpdateAppAsync(ctx workflow.Context, input *amplify.UpdateAppInput) *AmplifyUpdateAppResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateApp, input)
    return &AmplifyUpdateAppResult{Result: future}
}

func (a *AmplifyStub) UpdateBranch(ctx workflow.Context, input *amplify.UpdateBranchInput) (*amplify.UpdateBranchOutput, error) {
    var output amplify.UpdateBranchOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateBranch, input).Get(ctx, &output)
    return &output, err
}

func (a *AmplifyStub) UpdateBranchAsync(ctx workflow.Context, input *amplify.UpdateBranchInput) *AmplifyUpdateBranchResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateBranch, input)
    return &AmplifyUpdateBranchResult{Result: future}
}

func (a *AmplifyStub) UpdateDomainAssociation(ctx workflow.Context, input *amplify.UpdateDomainAssociationInput) (*amplify.UpdateDomainAssociationOutput, error) {
    var output amplify.UpdateDomainAssociationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateDomainAssociation, input).Get(ctx, &output)
    return &output, err
}

func (a *AmplifyStub) UpdateDomainAssociationAsync(ctx workflow.Context, input *amplify.UpdateDomainAssociationInput) *AmplifyUpdateDomainAssociationResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateDomainAssociation, input)
    return &AmplifyUpdateDomainAssociationResult{Result: future}
}

func (a *AmplifyStub) UpdateWebhook(ctx workflow.Context, input *amplify.UpdateWebhookInput) (*amplify.UpdateWebhookOutput, error) {
    var output amplify.UpdateWebhookOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateWebhook, input).Get(ctx, &output)
    return &output, err
}

func (a *AmplifyStub) UpdateWebhookAsync(ctx workflow.Context, input *amplify.UpdateWebhookInput) *AmplifyUpdateWebhookResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateWebhook, input)
    return &AmplifyUpdateWebhookResult{Result: future}
}