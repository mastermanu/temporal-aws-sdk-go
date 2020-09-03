package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/appconfig"
	"go.temporal.io/sdk/workflow"
)

type AppConfigClient interface {
    CreateApplication(ctx workflow.Context, input *appconfig.CreateApplicationInput) (*appconfig.CreateApplicationOutput, error)
    CreateApplicationAsync(ctx workflow.Context, input *appconfig.CreateApplicationInput) *AppconfigCreateApplicationResult

    CreateConfigurationProfile(ctx workflow.Context, input *appconfig.CreateConfigurationProfileInput) (*appconfig.CreateConfigurationProfileOutput, error)
    CreateConfigurationProfileAsync(ctx workflow.Context, input *appconfig.CreateConfigurationProfileInput) *AppconfigCreateConfigurationProfileResult

    CreateDeploymentStrategy(ctx workflow.Context, input *appconfig.CreateDeploymentStrategyInput) (*appconfig.CreateDeploymentStrategyOutput, error)
    CreateDeploymentStrategyAsync(ctx workflow.Context, input *appconfig.CreateDeploymentStrategyInput) *AppconfigCreateDeploymentStrategyResult

    CreateEnvironment(ctx workflow.Context, input *appconfig.CreateEnvironmentInput) (*appconfig.CreateEnvironmentOutput, error)
    CreateEnvironmentAsync(ctx workflow.Context, input *appconfig.CreateEnvironmentInput) *AppconfigCreateEnvironmentResult

    CreateHostedConfigurationVersion(ctx workflow.Context, input *appconfig.CreateHostedConfigurationVersionInput) (*appconfig.CreateHostedConfigurationVersionOutput, error)
    CreateHostedConfigurationVersionAsync(ctx workflow.Context, input *appconfig.CreateHostedConfigurationVersionInput) *AppconfigCreateHostedConfigurationVersionResult

    DeleteApplication(ctx workflow.Context, input *appconfig.DeleteApplicationInput) (*appconfig.DeleteApplicationOutput, error)
    DeleteApplicationAsync(ctx workflow.Context, input *appconfig.DeleteApplicationInput) *AppconfigDeleteApplicationResult

    DeleteConfigurationProfile(ctx workflow.Context, input *appconfig.DeleteConfigurationProfileInput) (*appconfig.DeleteConfigurationProfileOutput, error)
    DeleteConfigurationProfileAsync(ctx workflow.Context, input *appconfig.DeleteConfigurationProfileInput) *AppconfigDeleteConfigurationProfileResult

    DeleteDeploymentStrategy(ctx workflow.Context, input *appconfig.DeleteDeploymentStrategyInput) (*appconfig.DeleteDeploymentStrategyOutput, error)
    DeleteDeploymentStrategyAsync(ctx workflow.Context, input *appconfig.DeleteDeploymentStrategyInput) *AppconfigDeleteDeploymentStrategyResult

    DeleteEnvironment(ctx workflow.Context, input *appconfig.DeleteEnvironmentInput) (*appconfig.DeleteEnvironmentOutput, error)
    DeleteEnvironmentAsync(ctx workflow.Context, input *appconfig.DeleteEnvironmentInput) *AppconfigDeleteEnvironmentResult

    DeleteHostedConfigurationVersion(ctx workflow.Context, input *appconfig.DeleteHostedConfigurationVersionInput) (*appconfig.DeleteHostedConfigurationVersionOutput, error)
    DeleteHostedConfigurationVersionAsync(ctx workflow.Context, input *appconfig.DeleteHostedConfigurationVersionInput) *AppconfigDeleteHostedConfigurationVersionResult

    GetApplication(ctx workflow.Context, input *appconfig.GetApplicationInput) (*appconfig.GetApplicationOutput, error)
    GetApplicationAsync(ctx workflow.Context, input *appconfig.GetApplicationInput) *AppconfigGetApplicationResult

    GetConfiguration(ctx workflow.Context, input *appconfig.GetConfigurationInput) (*appconfig.GetConfigurationOutput, error)
    GetConfigurationAsync(ctx workflow.Context, input *appconfig.GetConfigurationInput) *AppconfigGetConfigurationResult

    GetConfigurationProfile(ctx workflow.Context, input *appconfig.GetConfigurationProfileInput) (*appconfig.GetConfigurationProfileOutput, error)
    GetConfigurationProfileAsync(ctx workflow.Context, input *appconfig.GetConfigurationProfileInput) *AppconfigGetConfigurationProfileResult

    GetDeployment(ctx workflow.Context, input *appconfig.GetDeploymentInput) (*appconfig.GetDeploymentOutput, error)
    GetDeploymentAsync(ctx workflow.Context, input *appconfig.GetDeploymentInput) *AppconfigGetDeploymentResult

    GetDeploymentStrategy(ctx workflow.Context, input *appconfig.GetDeploymentStrategyInput) (*appconfig.GetDeploymentStrategyOutput, error)
    GetDeploymentStrategyAsync(ctx workflow.Context, input *appconfig.GetDeploymentStrategyInput) *AppconfigGetDeploymentStrategyResult

    GetEnvironment(ctx workflow.Context, input *appconfig.GetEnvironmentInput) (*appconfig.GetEnvironmentOutput, error)
    GetEnvironmentAsync(ctx workflow.Context, input *appconfig.GetEnvironmentInput) *AppconfigGetEnvironmentResult

    GetHostedConfigurationVersion(ctx workflow.Context, input *appconfig.GetHostedConfigurationVersionInput) (*appconfig.GetHostedConfigurationVersionOutput, error)
    GetHostedConfigurationVersionAsync(ctx workflow.Context, input *appconfig.GetHostedConfigurationVersionInput) *AppconfigGetHostedConfigurationVersionResult

    ListApplications(ctx workflow.Context, input *appconfig.ListApplicationsInput) (*appconfig.ListApplicationsOutput, error)
    ListApplicationsAsync(ctx workflow.Context, input *appconfig.ListApplicationsInput) *AppconfigListApplicationsResult

    ListConfigurationProfiles(ctx workflow.Context, input *appconfig.ListConfigurationProfilesInput) (*appconfig.ListConfigurationProfilesOutput, error)
    ListConfigurationProfilesAsync(ctx workflow.Context, input *appconfig.ListConfigurationProfilesInput) *AppconfigListConfigurationProfilesResult

    ListDeploymentStrategies(ctx workflow.Context, input *appconfig.ListDeploymentStrategiesInput) (*appconfig.ListDeploymentStrategiesOutput, error)
    ListDeploymentStrategiesAsync(ctx workflow.Context, input *appconfig.ListDeploymentStrategiesInput) *AppconfigListDeploymentStrategiesResult

    ListDeployments(ctx workflow.Context, input *appconfig.ListDeploymentsInput) (*appconfig.ListDeploymentsOutput, error)
    ListDeploymentsAsync(ctx workflow.Context, input *appconfig.ListDeploymentsInput) *AppconfigListDeploymentsResult

    ListEnvironments(ctx workflow.Context, input *appconfig.ListEnvironmentsInput) (*appconfig.ListEnvironmentsOutput, error)
    ListEnvironmentsAsync(ctx workflow.Context, input *appconfig.ListEnvironmentsInput) *AppconfigListEnvironmentsResult

    ListHostedConfigurationVersions(ctx workflow.Context, input *appconfig.ListHostedConfigurationVersionsInput) (*appconfig.ListHostedConfigurationVersionsOutput, error)
    ListHostedConfigurationVersionsAsync(ctx workflow.Context, input *appconfig.ListHostedConfigurationVersionsInput) *AppconfigListHostedConfigurationVersionsResult

    ListTagsForResource(ctx workflow.Context, input *appconfig.ListTagsForResourceInput) (*appconfig.ListTagsForResourceOutput, error)
    ListTagsForResourceAsync(ctx workflow.Context, input *appconfig.ListTagsForResourceInput) *AppconfigListTagsForResourceResult

    StartDeployment(ctx workflow.Context, input *appconfig.StartDeploymentInput) (*appconfig.StartDeploymentOutput, error)
    StartDeploymentAsync(ctx workflow.Context, input *appconfig.StartDeploymentInput) *AppconfigStartDeploymentResult

    StopDeployment(ctx workflow.Context, input *appconfig.StopDeploymentInput) (*appconfig.StopDeploymentOutput, error)
    StopDeploymentAsync(ctx workflow.Context, input *appconfig.StopDeploymentInput) *AppconfigStopDeploymentResult

    TagResource(ctx workflow.Context, input *appconfig.TagResourceInput) (*appconfig.TagResourceOutput, error)
    TagResourceAsync(ctx workflow.Context, input *appconfig.TagResourceInput) *AppconfigTagResourceResult

    UntagResource(ctx workflow.Context, input *appconfig.UntagResourceInput) (*appconfig.UntagResourceOutput, error)
    UntagResourceAsync(ctx workflow.Context, input *appconfig.UntagResourceInput) *AppconfigUntagResourceResult

    UpdateApplication(ctx workflow.Context, input *appconfig.UpdateApplicationInput) (*appconfig.UpdateApplicationOutput, error)
    UpdateApplicationAsync(ctx workflow.Context, input *appconfig.UpdateApplicationInput) *AppconfigUpdateApplicationResult

    UpdateConfigurationProfile(ctx workflow.Context, input *appconfig.UpdateConfigurationProfileInput) (*appconfig.UpdateConfigurationProfileOutput, error)
    UpdateConfigurationProfileAsync(ctx workflow.Context, input *appconfig.UpdateConfigurationProfileInput) *AppconfigUpdateConfigurationProfileResult

    UpdateDeploymentStrategy(ctx workflow.Context, input *appconfig.UpdateDeploymentStrategyInput) (*appconfig.UpdateDeploymentStrategyOutput, error)
    UpdateDeploymentStrategyAsync(ctx workflow.Context, input *appconfig.UpdateDeploymentStrategyInput) *AppconfigUpdateDeploymentStrategyResult

    UpdateEnvironment(ctx workflow.Context, input *appconfig.UpdateEnvironmentInput) (*appconfig.UpdateEnvironmentOutput, error)
    UpdateEnvironmentAsync(ctx workflow.Context, input *appconfig.UpdateEnvironmentInput) *AppconfigUpdateEnvironmentResult

    ValidateConfiguration(ctx workflow.Context, input *appconfig.ValidateConfigurationInput) (*appconfig.ValidateConfigurationOutput, error)
    ValidateConfigurationAsync(ctx workflow.Context, input *appconfig.ValidateConfigurationInput) *AppconfigValidateConfigurationResult
}
type AppconfigCreateApplicationResult struct {
	Result workflow.Future
}

func (r *AppconfigCreateApplicationResult) Get(ctx workflow.Context) (*appconfig.CreateApplicationOutput, error) {
    var output appconfig.CreateApplicationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppconfigCreateConfigurationProfileResult struct {
	Result workflow.Future
}

func (r *AppconfigCreateConfigurationProfileResult) Get(ctx workflow.Context) (*appconfig.CreateConfigurationProfileOutput, error) {
    var output appconfig.CreateConfigurationProfileOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppconfigCreateDeploymentStrategyResult struct {
	Result workflow.Future
}

func (r *AppconfigCreateDeploymentStrategyResult) Get(ctx workflow.Context) (*appconfig.CreateDeploymentStrategyOutput, error) {
    var output appconfig.CreateDeploymentStrategyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppconfigCreateEnvironmentResult struct {
	Result workflow.Future
}

func (r *AppconfigCreateEnvironmentResult) Get(ctx workflow.Context) (*appconfig.CreateEnvironmentOutput, error) {
    var output appconfig.CreateEnvironmentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppconfigCreateHostedConfigurationVersionResult struct {
	Result workflow.Future
}

func (r *AppconfigCreateHostedConfigurationVersionResult) Get(ctx workflow.Context) (*appconfig.CreateHostedConfigurationVersionOutput, error) {
    var output appconfig.CreateHostedConfigurationVersionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppconfigDeleteApplicationResult struct {
	Result workflow.Future
}

func (r *AppconfigDeleteApplicationResult) Get(ctx workflow.Context) (*appconfig.DeleteApplicationOutput, error) {
    var output appconfig.DeleteApplicationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppconfigDeleteConfigurationProfileResult struct {
	Result workflow.Future
}

func (r *AppconfigDeleteConfigurationProfileResult) Get(ctx workflow.Context) (*appconfig.DeleteConfigurationProfileOutput, error) {
    var output appconfig.DeleteConfigurationProfileOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppconfigDeleteDeploymentStrategyResult struct {
	Result workflow.Future
}

func (r *AppconfigDeleteDeploymentStrategyResult) Get(ctx workflow.Context) (*appconfig.DeleteDeploymentStrategyOutput, error) {
    var output appconfig.DeleteDeploymentStrategyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppconfigDeleteEnvironmentResult struct {
	Result workflow.Future
}

func (r *AppconfigDeleteEnvironmentResult) Get(ctx workflow.Context) (*appconfig.DeleteEnvironmentOutput, error) {
    var output appconfig.DeleteEnvironmentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppconfigDeleteHostedConfigurationVersionResult struct {
	Result workflow.Future
}

func (r *AppconfigDeleteHostedConfigurationVersionResult) Get(ctx workflow.Context) (*appconfig.DeleteHostedConfigurationVersionOutput, error) {
    var output appconfig.DeleteHostedConfigurationVersionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppconfigGetApplicationResult struct {
	Result workflow.Future
}

func (r *AppconfigGetApplicationResult) Get(ctx workflow.Context) (*appconfig.GetApplicationOutput, error) {
    var output appconfig.GetApplicationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppconfigGetConfigurationResult struct {
	Result workflow.Future
}

func (r *AppconfigGetConfigurationResult) Get(ctx workflow.Context) (*appconfig.GetConfigurationOutput, error) {
    var output appconfig.GetConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppconfigGetConfigurationProfileResult struct {
	Result workflow.Future
}

func (r *AppconfigGetConfigurationProfileResult) Get(ctx workflow.Context) (*appconfig.GetConfigurationProfileOutput, error) {
    var output appconfig.GetConfigurationProfileOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppconfigGetDeploymentResult struct {
	Result workflow.Future
}

func (r *AppconfigGetDeploymentResult) Get(ctx workflow.Context) (*appconfig.GetDeploymentOutput, error) {
    var output appconfig.GetDeploymentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppconfigGetDeploymentStrategyResult struct {
	Result workflow.Future
}

func (r *AppconfigGetDeploymentStrategyResult) Get(ctx workflow.Context) (*appconfig.GetDeploymentStrategyOutput, error) {
    var output appconfig.GetDeploymentStrategyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppconfigGetEnvironmentResult struct {
	Result workflow.Future
}

func (r *AppconfigGetEnvironmentResult) Get(ctx workflow.Context) (*appconfig.GetEnvironmentOutput, error) {
    var output appconfig.GetEnvironmentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppconfigGetHostedConfigurationVersionResult struct {
	Result workflow.Future
}

func (r *AppconfigGetHostedConfigurationVersionResult) Get(ctx workflow.Context) (*appconfig.GetHostedConfigurationVersionOutput, error) {
    var output appconfig.GetHostedConfigurationVersionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppconfigListApplicationsResult struct {
	Result workflow.Future
}

func (r *AppconfigListApplicationsResult) Get(ctx workflow.Context) (*appconfig.ListApplicationsOutput, error) {
    var output appconfig.ListApplicationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppconfigListConfigurationProfilesResult struct {
	Result workflow.Future
}

func (r *AppconfigListConfigurationProfilesResult) Get(ctx workflow.Context) (*appconfig.ListConfigurationProfilesOutput, error) {
    var output appconfig.ListConfigurationProfilesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppconfigListDeploymentStrategiesResult struct {
	Result workflow.Future
}

func (r *AppconfigListDeploymentStrategiesResult) Get(ctx workflow.Context) (*appconfig.ListDeploymentStrategiesOutput, error) {
    var output appconfig.ListDeploymentStrategiesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppconfigListDeploymentsResult struct {
	Result workflow.Future
}

func (r *AppconfigListDeploymentsResult) Get(ctx workflow.Context) (*appconfig.ListDeploymentsOutput, error) {
    var output appconfig.ListDeploymentsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppconfigListEnvironmentsResult struct {
	Result workflow.Future
}

func (r *AppconfigListEnvironmentsResult) Get(ctx workflow.Context) (*appconfig.ListEnvironmentsOutput, error) {
    var output appconfig.ListEnvironmentsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppconfigListHostedConfigurationVersionsResult struct {
	Result workflow.Future
}

func (r *AppconfigListHostedConfigurationVersionsResult) Get(ctx workflow.Context) (*appconfig.ListHostedConfigurationVersionsOutput, error) {
    var output appconfig.ListHostedConfigurationVersionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppconfigListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *AppconfigListTagsForResourceResult) Get(ctx workflow.Context) (*appconfig.ListTagsForResourceOutput, error) {
    var output appconfig.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppconfigStartDeploymentResult struct {
	Result workflow.Future
}

func (r *AppconfigStartDeploymentResult) Get(ctx workflow.Context) (*appconfig.StartDeploymentOutput, error) {
    var output appconfig.StartDeploymentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppconfigStopDeploymentResult struct {
	Result workflow.Future
}

func (r *AppconfigStopDeploymentResult) Get(ctx workflow.Context) (*appconfig.StopDeploymentOutput, error) {
    var output appconfig.StopDeploymentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppconfigTagResourceResult struct {
	Result workflow.Future
}

func (r *AppconfigTagResourceResult) Get(ctx workflow.Context) (*appconfig.TagResourceOutput, error) {
    var output appconfig.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppconfigUntagResourceResult struct {
	Result workflow.Future
}

func (r *AppconfigUntagResourceResult) Get(ctx workflow.Context) (*appconfig.UntagResourceOutput, error) {
    var output appconfig.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppconfigUpdateApplicationResult struct {
	Result workflow.Future
}

func (r *AppconfigUpdateApplicationResult) Get(ctx workflow.Context) (*appconfig.UpdateApplicationOutput, error) {
    var output appconfig.UpdateApplicationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppconfigUpdateConfigurationProfileResult struct {
	Result workflow.Future
}

func (r *AppconfigUpdateConfigurationProfileResult) Get(ctx workflow.Context) (*appconfig.UpdateConfigurationProfileOutput, error) {
    var output appconfig.UpdateConfigurationProfileOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppconfigUpdateDeploymentStrategyResult struct {
	Result workflow.Future
}

func (r *AppconfigUpdateDeploymentStrategyResult) Get(ctx workflow.Context) (*appconfig.UpdateDeploymentStrategyOutput, error) {
    var output appconfig.UpdateDeploymentStrategyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppconfigUpdateEnvironmentResult struct {
	Result workflow.Future
}

func (r *AppconfigUpdateEnvironmentResult) Get(ctx workflow.Context) (*appconfig.UpdateEnvironmentOutput, error) {
    var output appconfig.UpdateEnvironmentOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type AppconfigValidateConfigurationResult struct {
	Result workflow.Future
}

func (r *AppconfigValidateConfigurationResult) Get(ctx workflow.Context) (*appconfig.ValidateConfigurationOutput, error) {
    var output appconfig.ValidateConfigurationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type AppConfigStub struct {
    activities AppConfigClient
}

func NewAppConfigStub() AppConfigClient {
    return &AppConfigStub{}
}

func (a *AppConfigStub) CreateApplication(ctx workflow.Context, input *appconfig.CreateApplicationInput) (*appconfig.CreateApplicationOutput, error) {
    var output appconfig.CreateApplicationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateApplication, input).Get(ctx, &output)
    return &output, err
}

func (a *AppConfigStub) CreateConfigurationProfile(ctx workflow.Context, input *appconfig.CreateConfigurationProfileInput) (*appconfig.CreateConfigurationProfileOutput, error) {
    var output appconfig.CreateConfigurationProfileOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateConfigurationProfile, input).Get(ctx, &output)
    return &output, err
}

func (a *AppConfigStub) CreateDeploymentStrategy(ctx workflow.Context, input *appconfig.CreateDeploymentStrategyInput) (*appconfig.CreateDeploymentStrategyOutput, error) {
    var output appconfig.CreateDeploymentStrategyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateDeploymentStrategy, input).Get(ctx, &output)
    return &output, err
}

func (a *AppConfigStub) CreateEnvironment(ctx workflow.Context, input *appconfig.CreateEnvironmentInput) (*appconfig.CreateEnvironmentOutput, error) {
    var output appconfig.CreateEnvironmentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateEnvironment, input).Get(ctx, &output)
    return &output, err
}

func (a *AppConfigStub) CreateHostedConfigurationVersion(ctx workflow.Context, input *appconfig.CreateHostedConfigurationVersionInput) (*appconfig.CreateHostedConfigurationVersionOutput, error) {
    var output appconfig.CreateHostedConfigurationVersionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateHostedConfigurationVersion, input).Get(ctx, &output)
    return &output, err
}

func (a *AppConfigStub) DeleteApplication(ctx workflow.Context, input *appconfig.DeleteApplicationInput) (*appconfig.DeleteApplicationOutput, error) {
    var output appconfig.DeleteApplicationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteApplication, input).Get(ctx, &output)
    return &output, err
}

func (a *AppConfigStub) DeleteConfigurationProfile(ctx workflow.Context, input *appconfig.DeleteConfigurationProfileInput) (*appconfig.DeleteConfigurationProfileOutput, error) {
    var output appconfig.DeleteConfigurationProfileOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteConfigurationProfile, input).Get(ctx, &output)
    return &output, err
}

func (a *AppConfigStub) DeleteDeploymentStrategy(ctx workflow.Context, input *appconfig.DeleteDeploymentStrategyInput) (*appconfig.DeleteDeploymentStrategyOutput, error) {
    var output appconfig.DeleteDeploymentStrategyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteDeploymentStrategy, input).Get(ctx, &output)
    return &output, err
}

func (a *AppConfigStub) DeleteEnvironment(ctx workflow.Context, input *appconfig.DeleteEnvironmentInput) (*appconfig.DeleteEnvironmentOutput, error) {
    var output appconfig.DeleteEnvironmentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteEnvironment, input).Get(ctx, &output)
    return &output, err
}

func (a *AppConfigStub) DeleteHostedConfigurationVersion(ctx workflow.Context, input *appconfig.DeleteHostedConfigurationVersionInput) (*appconfig.DeleteHostedConfigurationVersionOutput, error) {
    var output appconfig.DeleteHostedConfigurationVersionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteHostedConfigurationVersion, input).Get(ctx, &output)
    return &output, err
}

func (a *AppConfigStub) GetApplication(ctx workflow.Context, input *appconfig.GetApplicationInput) (*appconfig.GetApplicationOutput, error) {
    var output appconfig.GetApplicationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetApplication, input).Get(ctx, &output)
    return &output, err
}

func (a *AppConfigStub) GetConfiguration(ctx workflow.Context, input *appconfig.GetConfigurationInput) (*appconfig.GetConfigurationOutput, error) {
    var output appconfig.GetConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetConfiguration, input).Get(ctx, &output)
    return &output, err
}

func (a *AppConfigStub) GetConfigurationProfile(ctx workflow.Context, input *appconfig.GetConfigurationProfileInput) (*appconfig.GetConfigurationProfileOutput, error) {
    var output appconfig.GetConfigurationProfileOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetConfigurationProfile, input).Get(ctx, &output)
    return &output, err
}

func (a *AppConfigStub) GetDeployment(ctx workflow.Context, input *appconfig.GetDeploymentInput) (*appconfig.GetDeploymentOutput, error) {
    var output appconfig.GetDeploymentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetDeployment, input).Get(ctx, &output)
    return &output, err
}

func (a *AppConfigStub) GetDeploymentStrategy(ctx workflow.Context, input *appconfig.GetDeploymentStrategyInput) (*appconfig.GetDeploymentStrategyOutput, error) {
    var output appconfig.GetDeploymentStrategyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetDeploymentStrategy, input).Get(ctx, &output)
    return &output, err
}

func (a *AppConfigStub) GetEnvironment(ctx workflow.Context, input *appconfig.GetEnvironmentInput) (*appconfig.GetEnvironmentOutput, error) {
    var output appconfig.GetEnvironmentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetEnvironment, input).Get(ctx, &output)
    return &output, err
}

func (a *AppConfigStub) GetHostedConfigurationVersion(ctx workflow.Context, input *appconfig.GetHostedConfigurationVersionInput) (*appconfig.GetHostedConfigurationVersionOutput, error) {
    var output appconfig.GetHostedConfigurationVersionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetHostedConfigurationVersion, input).Get(ctx, &output)
    return &output, err
}

func (a *AppConfigStub) ListApplications(ctx workflow.Context, input *appconfig.ListApplicationsInput) (*appconfig.ListApplicationsOutput, error) {
    var output appconfig.ListApplicationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListApplications, input).Get(ctx, &output)
    return &output, err
}

func (a *AppConfigStub) ListConfigurationProfiles(ctx workflow.Context, input *appconfig.ListConfigurationProfilesInput) (*appconfig.ListConfigurationProfilesOutput, error) {
    var output appconfig.ListConfigurationProfilesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListConfigurationProfiles, input).Get(ctx, &output)
    return &output, err
}

func (a *AppConfigStub) ListDeploymentStrategies(ctx workflow.Context, input *appconfig.ListDeploymentStrategiesInput) (*appconfig.ListDeploymentStrategiesOutput, error) {
    var output appconfig.ListDeploymentStrategiesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDeploymentStrategies, input).Get(ctx, &output)
    return &output, err
}

func (a *AppConfigStub) ListDeployments(ctx workflow.Context, input *appconfig.ListDeploymentsInput) (*appconfig.ListDeploymentsOutput, error) {
    var output appconfig.ListDeploymentsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDeployments, input).Get(ctx, &output)
    return &output, err
}

func (a *AppConfigStub) ListEnvironments(ctx workflow.Context, input *appconfig.ListEnvironmentsInput) (*appconfig.ListEnvironmentsOutput, error) {
    var output appconfig.ListEnvironmentsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListEnvironments, input).Get(ctx, &output)
    return &output, err
}

func (a *AppConfigStub) ListHostedConfigurationVersions(ctx workflow.Context, input *appconfig.ListHostedConfigurationVersionsInput) (*appconfig.ListHostedConfigurationVersionsOutput, error) {
    var output appconfig.ListHostedConfigurationVersionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListHostedConfigurationVersions, input).Get(ctx, &output)
    return &output, err
}

func (a *AppConfigStub) ListTagsForResource(ctx workflow.Context, input *appconfig.ListTagsForResourceInput) (*appconfig.ListTagsForResourceOutput, error) {
    var output appconfig.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *AppConfigStub) StartDeployment(ctx workflow.Context, input *appconfig.StartDeploymentInput) (*appconfig.StartDeploymentOutput, error) {
    var output appconfig.StartDeploymentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartDeployment, input).Get(ctx, &output)
    return &output, err
}

func (a *AppConfigStub) StopDeployment(ctx workflow.Context, input *appconfig.StopDeploymentInput) (*appconfig.StopDeploymentOutput, error) {
    var output appconfig.StopDeploymentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StopDeployment, input).Get(ctx, &output)
    return &output, err
}

func (a *AppConfigStub) TagResource(ctx workflow.Context, input *appconfig.TagResourceInput) (*appconfig.TagResourceOutput, error) {
    var output appconfig.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *AppConfigStub) UntagResource(ctx workflow.Context, input *appconfig.UntagResourceInput) (*appconfig.UntagResourceOutput, error) {
    var output appconfig.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *AppConfigStub) UpdateApplication(ctx workflow.Context, input *appconfig.UpdateApplicationInput) (*appconfig.UpdateApplicationOutput, error) {
    var output appconfig.UpdateApplicationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateApplication, input).Get(ctx, &output)
    return &output, err
}

func (a *AppConfigStub) UpdateConfigurationProfile(ctx workflow.Context, input *appconfig.UpdateConfigurationProfileInput) (*appconfig.UpdateConfigurationProfileOutput, error) {
    var output appconfig.UpdateConfigurationProfileOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateConfigurationProfile, input).Get(ctx, &output)
    return &output, err
}

func (a *AppConfigStub) UpdateDeploymentStrategy(ctx workflow.Context, input *appconfig.UpdateDeploymentStrategyInput) (*appconfig.UpdateDeploymentStrategyOutput, error) {
    var output appconfig.UpdateDeploymentStrategyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateDeploymentStrategy, input).Get(ctx, &output)
    return &output, err
}

func (a *AppConfigStub) UpdateEnvironment(ctx workflow.Context, input *appconfig.UpdateEnvironmentInput) (*appconfig.UpdateEnvironmentOutput, error) {
    var output appconfig.UpdateEnvironmentOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateEnvironment, input).Get(ctx, &output)
    return &output, err
}

func (a *AppConfigStub) ValidateConfiguration(ctx workflow.Context, input *appconfig.ValidateConfigurationInput) (*appconfig.ValidateConfigurationOutput, error) {
    var output appconfig.ValidateConfigurationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ValidateConfiguration, input).Get(ctx, &output)
    return &output, err
}