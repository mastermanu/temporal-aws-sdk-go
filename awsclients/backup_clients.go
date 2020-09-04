package awsclients

import (
	"github.com/aws/aws-sdk-go/service/backup"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type BackupClient interface {
    CreateBackupPlan(ctx workflow.Context, input *backup.CreateBackupPlanInput) (*backup.CreateBackupPlanOutput, error)
    CreateBackupPlanAsync(ctx workflow.Context, input *backup.CreateBackupPlanInput) *BackupCreateBackupPlanResult

    CreateBackupSelection(ctx workflow.Context, input *backup.CreateBackupSelectionInput) (*backup.CreateBackupSelectionOutput, error)
    CreateBackupSelectionAsync(ctx workflow.Context, input *backup.CreateBackupSelectionInput) *BackupCreateBackupSelectionResult

    CreateBackupVault(ctx workflow.Context, input *backup.CreateBackupVaultInput) (*backup.CreateBackupVaultOutput, error)
    CreateBackupVaultAsync(ctx workflow.Context, input *backup.CreateBackupVaultInput) *BackupCreateBackupVaultResult

    DeleteBackupPlan(ctx workflow.Context, input *backup.DeleteBackupPlanInput) (*backup.DeleteBackupPlanOutput, error)
    DeleteBackupPlanAsync(ctx workflow.Context, input *backup.DeleteBackupPlanInput) *BackupDeleteBackupPlanResult

    DeleteBackupSelection(ctx workflow.Context, input *backup.DeleteBackupSelectionInput) (*backup.DeleteBackupSelectionOutput, error)
    DeleteBackupSelectionAsync(ctx workflow.Context, input *backup.DeleteBackupSelectionInput) *BackupDeleteBackupSelectionResult

    DeleteBackupVault(ctx workflow.Context, input *backup.DeleteBackupVaultInput) (*backup.DeleteBackupVaultOutput, error)
    DeleteBackupVaultAsync(ctx workflow.Context, input *backup.DeleteBackupVaultInput) *BackupDeleteBackupVaultResult

    DeleteBackupVaultAccessPolicy(ctx workflow.Context, input *backup.DeleteBackupVaultAccessPolicyInput) (*backup.DeleteBackupVaultAccessPolicyOutput, error)
    DeleteBackupVaultAccessPolicyAsync(ctx workflow.Context, input *backup.DeleteBackupVaultAccessPolicyInput) *BackupDeleteBackupVaultAccessPolicyResult

    DeleteBackupVaultNotifications(ctx workflow.Context, input *backup.DeleteBackupVaultNotificationsInput) (*backup.DeleteBackupVaultNotificationsOutput, error)
    DeleteBackupVaultNotificationsAsync(ctx workflow.Context, input *backup.DeleteBackupVaultNotificationsInput) *BackupDeleteBackupVaultNotificationsResult

    DeleteRecoveryPoint(ctx workflow.Context, input *backup.DeleteRecoveryPointInput) (*backup.DeleteRecoveryPointOutput, error)
    DeleteRecoveryPointAsync(ctx workflow.Context, input *backup.DeleteRecoveryPointInput) *BackupDeleteRecoveryPointResult

    DescribeBackupJob(ctx workflow.Context, input *backup.DescribeBackupJobInput) (*backup.DescribeBackupJobOutput, error)
    DescribeBackupJobAsync(ctx workflow.Context, input *backup.DescribeBackupJobInput) *BackupDescribeBackupJobResult

    DescribeBackupVault(ctx workflow.Context, input *backup.DescribeBackupVaultInput) (*backup.DescribeBackupVaultOutput, error)
    DescribeBackupVaultAsync(ctx workflow.Context, input *backup.DescribeBackupVaultInput) *BackupDescribeBackupVaultResult

    DescribeCopyJob(ctx workflow.Context, input *backup.DescribeCopyJobInput) (*backup.DescribeCopyJobOutput, error)
    DescribeCopyJobAsync(ctx workflow.Context, input *backup.DescribeCopyJobInput) *BackupDescribeCopyJobResult

    DescribeProtectedResource(ctx workflow.Context, input *backup.DescribeProtectedResourceInput) (*backup.DescribeProtectedResourceOutput, error)
    DescribeProtectedResourceAsync(ctx workflow.Context, input *backup.DescribeProtectedResourceInput) *BackupDescribeProtectedResourceResult

    DescribeRecoveryPoint(ctx workflow.Context, input *backup.DescribeRecoveryPointInput) (*backup.DescribeRecoveryPointOutput, error)
    DescribeRecoveryPointAsync(ctx workflow.Context, input *backup.DescribeRecoveryPointInput) *BackupDescribeRecoveryPointResult

    DescribeRegionSettings(ctx workflow.Context, input *backup.DescribeRegionSettingsInput) (*backup.DescribeRegionSettingsOutput, error)
    DescribeRegionSettingsAsync(ctx workflow.Context, input *backup.DescribeRegionSettingsInput) *BackupDescribeRegionSettingsResult

    DescribeRestoreJob(ctx workflow.Context, input *backup.DescribeRestoreJobInput) (*backup.DescribeRestoreJobOutput, error)
    DescribeRestoreJobAsync(ctx workflow.Context, input *backup.DescribeRestoreJobInput) *BackupDescribeRestoreJobResult

    ExportBackupPlanTemplate(ctx workflow.Context, input *backup.ExportBackupPlanTemplateInput) (*backup.ExportBackupPlanTemplateOutput, error)
    ExportBackupPlanTemplateAsync(ctx workflow.Context, input *backup.ExportBackupPlanTemplateInput) *BackupExportBackupPlanTemplateResult

    GetBackupPlan(ctx workflow.Context, input *backup.GetBackupPlanInput) (*backup.GetBackupPlanOutput, error)
    GetBackupPlanAsync(ctx workflow.Context, input *backup.GetBackupPlanInput) *BackupGetBackupPlanResult

    GetBackupPlanFromJSON(ctx workflow.Context, input *backup.GetBackupPlanFromJSONInput) (*backup.GetBackupPlanFromJSONOutput, error)
    GetBackupPlanFromJSONAsync(ctx workflow.Context, input *backup.GetBackupPlanFromJSONInput) *BackupGetBackupPlanFromJSONResult

    GetBackupPlanFromTemplate(ctx workflow.Context, input *backup.GetBackupPlanFromTemplateInput) (*backup.GetBackupPlanFromTemplateOutput, error)
    GetBackupPlanFromTemplateAsync(ctx workflow.Context, input *backup.GetBackupPlanFromTemplateInput) *BackupGetBackupPlanFromTemplateResult

    GetBackupSelection(ctx workflow.Context, input *backup.GetBackupSelectionInput) (*backup.GetBackupSelectionOutput, error)
    GetBackupSelectionAsync(ctx workflow.Context, input *backup.GetBackupSelectionInput) *BackupGetBackupSelectionResult

    GetBackupVaultAccessPolicy(ctx workflow.Context, input *backup.GetBackupVaultAccessPolicyInput) (*backup.GetBackupVaultAccessPolicyOutput, error)
    GetBackupVaultAccessPolicyAsync(ctx workflow.Context, input *backup.GetBackupVaultAccessPolicyInput) *BackupGetBackupVaultAccessPolicyResult

    GetBackupVaultNotifications(ctx workflow.Context, input *backup.GetBackupVaultNotificationsInput) (*backup.GetBackupVaultNotificationsOutput, error)
    GetBackupVaultNotificationsAsync(ctx workflow.Context, input *backup.GetBackupVaultNotificationsInput) *BackupGetBackupVaultNotificationsResult

    GetRecoveryPointRestoreMetadata(ctx workflow.Context, input *backup.GetRecoveryPointRestoreMetadataInput) (*backup.GetRecoveryPointRestoreMetadataOutput, error)
    GetRecoveryPointRestoreMetadataAsync(ctx workflow.Context, input *backup.GetRecoveryPointRestoreMetadataInput) *BackupGetRecoveryPointRestoreMetadataResult

    GetSupportedResourceTypes(ctx workflow.Context, input *backup.GetSupportedResourceTypesInput) (*backup.GetSupportedResourceTypesOutput, error)
    GetSupportedResourceTypesAsync(ctx workflow.Context, input *backup.GetSupportedResourceTypesInput) *BackupGetSupportedResourceTypesResult

    ListBackupJobs(ctx workflow.Context, input *backup.ListBackupJobsInput) (*backup.ListBackupJobsOutput, error)
    ListBackupJobsAsync(ctx workflow.Context, input *backup.ListBackupJobsInput) *BackupListBackupJobsResult

    ListBackupPlanTemplates(ctx workflow.Context, input *backup.ListBackupPlanTemplatesInput) (*backup.ListBackupPlanTemplatesOutput, error)
    ListBackupPlanTemplatesAsync(ctx workflow.Context, input *backup.ListBackupPlanTemplatesInput) *BackupListBackupPlanTemplatesResult

    ListBackupPlanVersions(ctx workflow.Context, input *backup.ListBackupPlanVersionsInput) (*backup.ListBackupPlanVersionsOutput, error)
    ListBackupPlanVersionsAsync(ctx workflow.Context, input *backup.ListBackupPlanVersionsInput) *BackupListBackupPlanVersionsResult

    ListBackupPlans(ctx workflow.Context, input *backup.ListBackupPlansInput) (*backup.ListBackupPlansOutput, error)
    ListBackupPlansAsync(ctx workflow.Context, input *backup.ListBackupPlansInput) *BackupListBackupPlansResult

    ListBackupSelections(ctx workflow.Context, input *backup.ListBackupSelectionsInput) (*backup.ListBackupSelectionsOutput, error)
    ListBackupSelectionsAsync(ctx workflow.Context, input *backup.ListBackupSelectionsInput) *BackupListBackupSelectionsResult

    ListBackupVaults(ctx workflow.Context, input *backup.ListBackupVaultsInput) (*backup.ListBackupVaultsOutput, error)
    ListBackupVaultsAsync(ctx workflow.Context, input *backup.ListBackupVaultsInput) *BackupListBackupVaultsResult

    ListCopyJobs(ctx workflow.Context, input *backup.ListCopyJobsInput) (*backup.ListCopyJobsOutput, error)
    ListCopyJobsAsync(ctx workflow.Context, input *backup.ListCopyJobsInput) *BackupListCopyJobsResult

    ListProtectedResources(ctx workflow.Context, input *backup.ListProtectedResourcesInput) (*backup.ListProtectedResourcesOutput, error)
    ListProtectedResourcesAsync(ctx workflow.Context, input *backup.ListProtectedResourcesInput) *BackupListProtectedResourcesResult

    ListRecoveryPointsByBackupVault(ctx workflow.Context, input *backup.ListRecoveryPointsByBackupVaultInput) (*backup.ListRecoveryPointsByBackupVaultOutput, error)
    ListRecoveryPointsByBackupVaultAsync(ctx workflow.Context, input *backup.ListRecoveryPointsByBackupVaultInput) *BackupListRecoveryPointsByBackupVaultResult

    ListRecoveryPointsByResource(ctx workflow.Context, input *backup.ListRecoveryPointsByResourceInput) (*backup.ListRecoveryPointsByResourceOutput, error)
    ListRecoveryPointsByResourceAsync(ctx workflow.Context, input *backup.ListRecoveryPointsByResourceInput) *BackupListRecoveryPointsByResourceResult

    ListRestoreJobs(ctx workflow.Context, input *backup.ListRestoreJobsInput) (*backup.ListRestoreJobsOutput, error)
    ListRestoreJobsAsync(ctx workflow.Context, input *backup.ListRestoreJobsInput) *BackupListRestoreJobsResult

    ListTags(ctx workflow.Context, input *backup.ListTagsInput) (*backup.ListTagsOutput, error)
    ListTagsAsync(ctx workflow.Context, input *backup.ListTagsInput) *BackupListTagsResult

    PutBackupVaultAccessPolicy(ctx workflow.Context, input *backup.PutBackupVaultAccessPolicyInput) (*backup.PutBackupVaultAccessPolicyOutput, error)
    PutBackupVaultAccessPolicyAsync(ctx workflow.Context, input *backup.PutBackupVaultAccessPolicyInput) *BackupPutBackupVaultAccessPolicyResult

    PutBackupVaultNotifications(ctx workflow.Context, input *backup.PutBackupVaultNotificationsInput) (*backup.PutBackupVaultNotificationsOutput, error)
    PutBackupVaultNotificationsAsync(ctx workflow.Context, input *backup.PutBackupVaultNotificationsInput) *BackupPutBackupVaultNotificationsResult

    StartBackupJob(ctx workflow.Context, input *backup.StartBackupJobInput) (*backup.StartBackupJobOutput, error)
    StartBackupJobAsync(ctx workflow.Context, input *backup.StartBackupJobInput) *BackupStartBackupJobResult

    StartCopyJob(ctx workflow.Context, input *backup.StartCopyJobInput) (*backup.StartCopyJobOutput, error)
    StartCopyJobAsync(ctx workflow.Context, input *backup.StartCopyJobInput) *BackupStartCopyJobResult

    StartRestoreJob(ctx workflow.Context, input *backup.StartRestoreJobInput) (*backup.StartRestoreJobOutput, error)
    StartRestoreJobAsync(ctx workflow.Context, input *backup.StartRestoreJobInput) *BackupStartRestoreJobResult

    StopBackupJob(ctx workflow.Context, input *backup.StopBackupJobInput) (*backup.StopBackupJobOutput, error)
    StopBackupJobAsync(ctx workflow.Context, input *backup.StopBackupJobInput) *BackupStopBackupJobResult

    TagResource(ctx workflow.Context, input *backup.TagResourceInput) (*backup.TagResourceOutput, error)
    TagResourceAsync(ctx workflow.Context, input *backup.TagResourceInput) *BackupTagResourceResult

    UntagResource(ctx workflow.Context, input *backup.UntagResourceInput) (*backup.UntagResourceOutput, error)
    UntagResourceAsync(ctx workflow.Context, input *backup.UntagResourceInput) *BackupUntagResourceResult

    UpdateBackupPlan(ctx workflow.Context, input *backup.UpdateBackupPlanInput) (*backup.UpdateBackupPlanOutput, error)
    UpdateBackupPlanAsync(ctx workflow.Context, input *backup.UpdateBackupPlanInput) *BackupUpdateBackupPlanResult

    UpdateRecoveryPointLifecycle(ctx workflow.Context, input *backup.UpdateRecoveryPointLifecycleInput) (*backup.UpdateRecoveryPointLifecycleOutput, error)
    UpdateRecoveryPointLifecycleAsync(ctx workflow.Context, input *backup.UpdateRecoveryPointLifecycleInput) *BackupUpdateRecoveryPointLifecycleResult

    UpdateRegionSettings(ctx workflow.Context, input *backup.UpdateRegionSettingsInput) (*backup.UpdateRegionSettingsOutput, error)
    UpdateRegionSettingsAsync(ctx workflow.Context, input *backup.UpdateRegionSettingsInput) *BackupUpdateRegionSettingsResult
}

type BackupCreateBackupPlanResult struct {
	Result workflow.Future
}

func (r *BackupCreateBackupPlanResult) Get(ctx workflow.Context) (*backup.CreateBackupPlanOutput, error) {
    var output backup.CreateBackupPlanOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BackupCreateBackupSelectionResult struct {
	Result workflow.Future
}

func (r *BackupCreateBackupSelectionResult) Get(ctx workflow.Context) (*backup.CreateBackupSelectionOutput, error) {
    var output backup.CreateBackupSelectionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BackupCreateBackupVaultResult struct {
	Result workflow.Future
}

func (r *BackupCreateBackupVaultResult) Get(ctx workflow.Context) (*backup.CreateBackupVaultOutput, error) {
    var output backup.CreateBackupVaultOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BackupDeleteBackupPlanResult struct {
	Result workflow.Future
}

func (r *BackupDeleteBackupPlanResult) Get(ctx workflow.Context) (*backup.DeleteBackupPlanOutput, error) {
    var output backup.DeleteBackupPlanOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BackupDeleteBackupSelectionResult struct {
	Result workflow.Future
}

func (r *BackupDeleteBackupSelectionResult) Get(ctx workflow.Context) (*backup.DeleteBackupSelectionOutput, error) {
    var output backup.DeleteBackupSelectionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BackupDeleteBackupVaultResult struct {
	Result workflow.Future
}

func (r *BackupDeleteBackupVaultResult) Get(ctx workflow.Context) (*backup.DeleteBackupVaultOutput, error) {
    var output backup.DeleteBackupVaultOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BackupDeleteBackupVaultAccessPolicyResult struct {
	Result workflow.Future
}

func (r *BackupDeleteBackupVaultAccessPolicyResult) Get(ctx workflow.Context) (*backup.DeleteBackupVaultAccessPolicyOutput, error) {
    var output backup.DeleteBackupVaultAccessPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BackupDeleteBackupVaultNotificationsResult struct {
	Result workflow.Future
}

func (r *BackupDeleteBackupVaultNotificationsResult) Get(ctx workflow.Context) (*backup.DeleteBackupVaultNotificationsOutput, error) {
    var output backup.DeleteBackupVaultNotificationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BackupDeleteRecoveryPointResult struct {
	Result workflow.Future
}

func (r *BackupDeleteRecoveryPointResult) Get(ctx workflow.Context) (*backup.DeleteRecoveryPointOutput, error) {
    var output backup.DeleteRecoveryPointOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BackupDescribeBackupJobResult struct {
	Result workflow.Future
}

func (r *BackupDescribeBackupJobResult) Get(ctx workflow.Context) (*backup.DescribeBackupJobOutput, error) {
    var output backup.DescribeBackupJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BackupDescribeBackupVaultResult struct {
	Result workflow.Future
}

func (r *BackupDescribeBackupVaultResult) Get(ctx workflow.Context) (*backup.DescribeBackupVaultOutput, error) {
    var output backup.DescribeBackupVaultOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BackupDescribeCopyJobResult struct {
	Result workflow.Future
}

func (r *BackupDescribeCopyJobResult) Get(ctx workflow.Context) (*backup.DescribeCopyJobOutput, error) {
    var output backup.DescribeCopyJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BackupDescribeProtectedResourceResult struct {
	Result workflow.Future
}

func (r *BackupDescribeProtectedResourceResult) Get(ctx workflow.Context) (*backup.DescribeProtectedResourceOutput, error) {
    var output backup.DescribeProtectedResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BackupDescribeRecoveryPointResult struct {
	Result workflow.Future
}

func (r *BackupDescribeRecoveryPointResult) Get(ctx workflow.Context) (*backup.DescribeRecoveryPointOutput, error) {
    var output backup.DescribeRecoveryPointOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BackupDescribeRegionSettingsResult struct {
	Result workflow.Future
}

func (r *BackupDescribeRegionSettingsResult) Get(ctx workflow.Context) (*backup.DescribeRegionSettingsOutput, error) {
    var output backup.DescribeRegionSettingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BackupDescribeRestoreJobResult struct {
	Result workflow.Future
}

func (r *BackupDescribeRestoreJobResult) Get(ctx workflow.Context) (*backup.DescribeRestoreJobOutput, error) {
    var output backup.DescribeRestoreJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BackupExportBackupPlanTemplateResult struct {
	Result workflow.Future
}

func (r *BackupExportBackupPlanTemplateResult) Get(ctx workflow.Context) (*backup.ExportBackupPlanTemplateOutput, error) {
    var output backup.ExportBackupPlanTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BackupGetBackupPlanResult struct {
	Result workflow.Future
}

func (r *BackupGetBackupPlanResult) Get(ctx workflow.Context) (*backup.GetBackupPlanOutput, error) {
    var output backup.GetBackupPlanOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BackupGetBackupPlanFromJSONResult struct {
	Result workflow.Future
}

func (r *BackupGetBackupPlanFromJSONResult) Get(ctx workflow.Context) (*backup.GetBackupPlanFromJSONOutput, error) {
    var output backup.GetBackupPlanFromJSONOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BackupGetBackupPlanFromTemplateResult struct {
	Result workflow.Future
}

func (r *BackupGetBackupPlanFromTemplateResult) Get(ctx workflow.Context) (*backup.GetBackupPlanFromTemplateOutput, error) {
    var output backup.GetBackupPlanFromTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BackupGetBackupSelectionResult struct {
	Result workflow.Future
}

func (r *BackupGetBackupSelectionResult) Get(ctx workflow.Context) (*backup.GetBackupSelectionOutput, error) {
    var output backup.GetBackupSelectionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BackupGetBackupVaultAccessPolicyResult struct {
	Result workflow.Future
}

func (r *BackupGetBackupVaultAccessPolicyResult) Get(ctx workflow.Context) (*backup.GetBackupVaultAccessPolicyOutput, error) {
    var output backup.GetBackupVaultAccessPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BackupGetBackupVaultNotificationsResult struct {
	Result workflow.Future
}

func (r *BackupGetBackupVaultNotificationsResult) Get(ctx workflow.Context) (*backup.GetBackupVaultNotificationsOutput, error) {
    var output backup.GetBackupVaultNotificationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BackupGetRecoveryPointRestoreMetadataResult struct {
	Result workflow.Future
}

func (r *BackupGetRecoveryPointRestoreMetadataResult) Get(ctx workflow.Context) (*backup.GetRecoveryPointRestoreMetadataOutput, error) {
    var output backup.GetRecoveryPointRestoreMetadataOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BackupGetSupportedResourceTypesResult struct {
	Result workflow.Future
}

func (r *BackupGetSupportedResourceTypesResult) Get(ctx workflow.Context) (*backup.GetSupportedResourceTypesOutput, error) {
    var output backup.GetSupportedResourceTypesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BackupListBackupJobsResult struct {
	Result workflow.Future
}

func (r *BackupListBackupJobsResult) Get(ctx workflow.Context) (*backup.ListBackupJobsOutput, error) {
    var output backup.ListBackupJobsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BackupListBackupPlanTemplatesResult struct {
	Result workflow.Future
}

func (r *BackupListBackupPlanTemplatesResult) Get(ctx workflow.Context) (*backup.ListBackupPlanTemplatesOutput, error) {
    var output backup.ListBackupPlanTemplatesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BackupListBackupPlanVersionsResult struct {
	Result workflow.Future
}

func (r *BackupListBackupPlanVersionsResult) Get(ctx workflow.Context) (*backup.ListBackupPlanVersionsOutput, error) {
    var output backup.ListBackupPlanVersionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BackupListBackupPlansResult struct {
	Result workflow.Future
}

func (r *BackupListBackupPlansResult) Get(ctx workflow.Context) (*backup.ListBackupPlansOutput, error) {
    var output backup.ListBackupPlansOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BackupListBackupSelectionsResult struct {
	Result workflow.Future
}

func (r *BackupListBackupSelectionsResult) Get(ctx workflow.Context) (*backup.ListBackupSelectionsOutput, error) {
    var output backup.ListBackupSelectionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BackupListBackupVaultsResult struct {
	Result workflow.Future
}

func (r *BackupListBackupVaultsResult) Get(ctx workflow.Context) (*backup.ListBackupVaultsOutput, error) {
    var output backup.ListBackupVaultsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BackupListCopyJobsResult struct {
	Result workflow.Future
}

func (r *BackupListCopyJobsResult) Get(ctx workflow.Context) (*backup.ListCopyJobsOutput, error) {
    var output backup.ListCopyJobsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BackupListProtectedResourcesResult struct {
	Result workflow.Future
}

func (r *BackupListProtectedResourcesResult) Get(ctx workflow.Context) (*backup.ListProtectedResourcesOutput, error) {
    var output backup.ListProtectedResourcesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BackupListRecoveryPointsByBackupVaultResult struct {
	Result workflow.Future
}

func (r *BackupListRecoveryPointsByBackupVaultResult) Get(ctx workflow.Context) (*backup.ListRecoveryPointsByBackupVaultOutput, error) {
    var output backup.ListRecoveryPointsByBackupVaultOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BackupListRecoveryPointsByResourceResult struct {
	Result workflow.Future
}

func (r *BackupListRecoveryPointsByResourceResult) Get(ctx workflow.Context) (*backup.ListRecoveryPointsByResourceOutput, error) {
    var output backup.ListRecoveryPointsByResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BackupListRestoreJobsResult struct {
	Result workflow.Future
}

func (r *BackupListRestoreJobsResult) Get(ctx workflow.Context) (*backup.ListRestoreJobsOutput, error) {
    var output backup.ListRestoreJobsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BackupListTagsResult struct {
	Result workflow.Future
}

func (r *BackupListTagsResult) Get(ctx workflow.Context) (*backup.ListTagsOutput, error) {
    var output backup.ListTagsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BackupPutBackupVaultAccessPolicyResult struct {
	Result workflow.Future
}

func (r *BackupPutBackupVaultAccessPolicyResult) Get(ctx workflow.Context) (*backup.PutBackupVaultAccessPolicyOutput, error) {
    var output backup.PutBackupVaultAccessPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BackupPutBackupVaultNotificationsResult struct {
	Result workflow.Future
}

func (r *BackupPutBackupVaultNotificationsResult) Get(ctx workflow.Context) (*backup.PutBackupVaultNotificationsOutput, error) {
    var output backup.PutBackupVaultNotificationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BackupStartBackupJobResult struct {
	Result workflow.Future
}

func (r *BackupStartBackupJobResult) Get(ctx workflow.Context) (*backup.StartBackupJobOutput, error) {
    var output backup.StartBackupJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BackupStartCopyJobResult struct {
	Result workflow.Future
}

func (r *BackupStartCopyJobResult) Get(ctx workflow.Context) (*backup.StartCopyJobOutput, error) {
    var output backup.StartCopyJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BackupStartRestoreJobResult struct {
	Result workflow.Future
}

func (r *BackupStartRestoreJobResult) Get(ctx workflow.Context) (*backup.StartRestoreJobOutput, error) {
    var output backup.StartRestoreJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BackupStopBackupJobResult struct {
	Result workflow.Future
}

func (r *BackupStopBackupJobResult) Get(ctx workflow.Context) (*backup.StopBackupJobOutput, error) {
    var output backup.StopBackupJobOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BackupTagResourceResult struct {
	Result workflow.Future
}

func (r *BackupTagResourceResult) Get(ctx workflow.Context) (*backup.TagResourceOutput, error) {
    var output backup.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BackupUntagResourceResult struct {
	Result workflow.Future
}

func (r *BackupUntagResourceResult) Get(ctx workflow.Context) (*backup.UntagResourceOutput, error) {
    var output backup.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BackupUpdateBackupPlanResult struct {
	Result workflow.Future
}

func (r *BackupUpdateBackupPlanResult) Get(ctx workflow.Context) (*backup.UpdateBackupPlanOutput, error) {
    var output backup.UpdateBackupPlanOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BackupUpdateRecoveryPointLifecycleResult struct {
	Result workflow.Future
}

func (r *BackupUpdateRecoveryPointLifecycleResult) Get(ctx workflow.Context) (*backup.UpdateRecoveryPointLifecycleOutput, error) {
    var output backup.UpdateRecoveryPointLifecycleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BackupUpdateRegionSettingsResult struct {
	Result workflow.Future
}

func (r *BackupUpdateRegionSettingsResult) Get(ctx workflow.Context) (*backup.UpdateRegionSettingsOutput, error) {
    var output backup.UpdateRegionSettingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type BackupStub struct {
    activities awsactivities.BackupActivities
}

func NewBackupStub() BackupClient {
    return &BackupStub{}
}

func (a *BackupStub) CreateBackupPlan(ctx workflow.Context, input *backup.CreateBackupPlanInput) (*backup.CreateBackupPlanOutput, error) {
    var output backup.CreateBackupPlanOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateBackupPlan, input).Get(ctx, &output)
    return &output, err
}

func (a *BackupStub) CreateBackupPlanAsync(ctx workflow.Context, input *backup.CreateBackupPlanInput) *BackupCreateBackupPlanResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateBackupPlan, input)
    return &BackupCreateBackupPlanResult{Result: future}
}

func (a *BackupStub) CreateBackupSelection(ctx workflow.Context, input *backup.CreateBackupSelectionInput) (*backup.CreateBackupSelectionOutput, error) {
    var output backup.CreateBackupSelectionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateBackupSelection, input).Get(ctx, &output)
    return &output, err
}

func (a *BackupStub) CreateBackupSelectionAsync(ctx workflow.Context, input *backup.CreateBackupSelectionInput) *BackupCreateBackupSelectionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateBackupSelection, input)
    return &BackupCreateBackupSelectionResult{Result: future}
}

func (a *BackupStub) CreateBackupVault(ctx workflow.Context, input *backup.CreateBackupVaultInput) (*backup.CreateBackupVaultOutput, error) {
    var output backup.CreateBackupVaultOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateBackupVault, input).Get(ctx, &output)
    return &output, err
}

func (a *BackupStub) CreateBackupVaultAsync(ctx workflow.Context, input *backup.CreateBackupVaultInput) *BackupCreateBackupVaultResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateBackupVault, input)
    return &BackupCreateBackupVaultResult{Result: future}
}

func (a *BackupStub) DeleteBackupPlan(ctx workflow.Context, input *backup.DeleteBackupPlanInput) (*backup.DeleteBackupPlanOutput, error) {
    var output backup.DeleteBackupPlanOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteBackupPlan, input).Get(ctx, &output)
    return &output, err
}

func (a *BackupStub) DeleteBackupPlanAsync(ctx workflow.Context, input *backup.DeleteBackupPlanInput) *BackupDeleteBackupPlanResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteBackupPlan, input)
    return &BackupDeleteBackupPlanResult{Result: future}
}

func (a *BackupStub) DeleteBackupSelection(ctx workflow.Context, input *backup.DeleteBackupSelectionInput) (*backup.DeleteBackupSelectionOutput, error) {
    var output backup.DeleteBackupSelectionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteBackupSelection, input).Get(ctx, &output)
    return &output, err
}

func (a *BackupStub) DeleteBackupSelectionAsync(ctx workflow.Context, input *backup.DeleteBackupSelectionInput) *BackupDeleteBackupSelectionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteBackupSelection, input)
    return &BackupDeleteBackupSelectionResult{Result: future}
}

func (a *BackupStub) DeleteBackupVault(ctx workflow.Context, input *backup.DeleteBackupVaultInput) (*backup.DeleteBackupVaultOutput, error) {
    var output backup.DeleteBackupVaultOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteBackupVault, input).Get(ctx, &output)
    return &output, err
}

func (a *BackupStub) DeleteBackupVaultAsync(ctx workflow.Context, input *backup.DeleteBackupVaultInput) *BackupDeleteBackupVaultResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteBackupVault, input)
    return &BackupDeleteBackupVaultResult{Result: future}
}

func (a *BackupStub) DeleteBackupVaultAccessPolicy(ctx workflow.Context, input *backup.DeleteBackupVaultAccessPolicyInput) (*backup.DeleteBackupVaultAccessPolicyOutput, error) {
    var output backup.DeleteBackupVaultAccessPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteBackupVaultAccessPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *BackupStub) DeleteBackupVaultAccessPolicyAsync(ctx workflow.Context, input *backup.DeleteBackupVaultAccessPolicyInput) *BackupDeleteBackupVaultAccessPolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteBackupVaultAccessPolicy, input)
    return &BackupDeleteBackupVaultAccessPolicyResult{Result: future}
}

func (a *BackupStub) DeleteBackupVaultNotifications(ctx workflow.Context, input *backup.DeleteBackupVaultNotificationsInput) (*backup.DeleteBackupVaultNotificationsOutput, error) {
    var output backup.DeleteBackupVaultNotificationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteBackupVaultNotifications, input).Get(ctx, &output)
    return &output, err
}

func (a *BackupStub) DeleteBackupVaultNotificationsAsync(ctx workflow.Context, input *backup.DeleteBackupVaultNotificationsInput) *BackupDeleteBackupVaultNotificationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteBackupVaultNotifications, input)
    return &BackupDeleteBackupVaultNotificationsResult{Result: future}
}

func (a *BackupStub) DeleteRecoveryPoint(ctx workflow.Context, input *backup.DeleteRecoveryPointInput) (*backup.DeleteRecoveryPointOutput, error) {
    var output backup.DeleteRecoveryPointOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteRecoveryPoint, input).Get(ctx, &output)
    return &output, err
}

func (a *BackupStub) DeleteRecoveryPointAsync(ctx workflow.Context, input *backup.DeleteRecoveryPointInput) *BackupDeleteRecoveryPointResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteRecoveryPoint, input)
    return &BackupDeleteRecoveryPointResult{Result: future}
}

func (a *BackupStub) DescribeBackupJob(ctx workflow.Context, input *backup.DescribeBackupJobInput) (*backup.DescribeBackupJobOutput, error) {
    var output backup.DescribeBackupJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeBackupJob, input).Get(ctx, &output)
    return &output, err
}

func (a *BackupStub) DescribeBackupJobAsync(ctx workflow.Context, input *backup.DescribeBackupJobInput) *BackupDescribeBackupJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeBackupJob, input)
    return &BackupDescribeBackupJobResult{Result: future}
}

func (a *BackupStub) DescribeBackupVault(ctx workflow.Context, input *backup.DescribeBackupVaultInput) (*backup.DescribeBackupVaultOutput, error) {
    var output backup.DescribeBackupVaultOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeBackupVault, input).Get(ctx, &output)
    return &output, err
}

func (a *BackupStub) DescribeBackupVaultAsync(ctx workflow.Context, input *backup.DescribeBackupVaultInput) *BackupDescribeBackupVaultResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeBackupVault, input)
    return &BackupDescribeBackupVaultResult{Result: future}
}

func (a *BackupStub) DescribeCopyJob(ctx workflow.Context, input *backup.DescribeCopyJobInput) (*backup.DescribeCopyJobOutput, error) {
    var output backup.DescribeCopyJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeCopyJob, input).Get(ctx, &output)
    return &output, err
}

func (a *BackupStub) DescribeCopyJobAsync(ctx workflow.Context, input *backup.DescribeCopyJobInput) *BackupDescribeCopyJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeCopyJob, input)
    return &BackupDescribeCopyJobResult{Result: future}
}

func (a *BackupStub) DescribeProtectedResource(ctx workflow.Context, input *backup.DescribeProtectedResourceInput) (*backup.DescribeProtectedResourceOutput, error) {
    var output backup.DescribeProtectedResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeProtectedResource, input).Get(ctx, &output)
    return &output, err
}

func (a *BackupStub) DescribeProtectedResourceAsync(ctx workflow.Context, input *backup.DescribeProtectedResourceInput) *BackupDescribeProtectedResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeProtectedResource, input)
    return &BackupDescribeProtectedResourceResult{Result: future}
}

func (a *BackupStub) DescribeRecoveryPoint(ctx workflow.Context, input *backup.DescribeRecoveryPointInput) (*backup.DescribeRecoveryPointOutput, error) {
    var output backup.DescribeRecoveryPointOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeRecoveryPoint, input).Get(ctx, &output)
    return &output, err
}

func (a *BackupStub) DescribeRecoveryPointAsync(ctx workflow.Context, input *backup.DescribeRecoveryPointInput) *BackupDescribeRecoveryPointResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeRecoveryPoint, input)
    return &BackupDescribeRecoveryPointResult{Result: future}
}

func (a *BackupStub) DescribeRegionSettings(ctx workflow.Context, input *backup.DescribeRegionSettingsInput) (*backup.DescribeRegionSettingsOutput, error) {
    var output backup.DescribeRegionSettingsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeRegionSettings, input).Get(ctx, &output)
    return &output, err
}

func (a *BackupStub) DescribeRegionSettingsAsync(ctx workflow.Context, input *backup.DescribeRegionSettingsInput) *BackupDescribeRegionSettingsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeRegionSettings, input)
    return &BackupDescribeRegionSettingsResult{Result: future}
}

func (a *BackupStub) DescribeRestoreJob(ctx workflow.Context, input *backup.DescribeRestoreJobInput) (*backup.DescribeRestoreJobOutput, error) {
    var output backup.DescribeRestoreJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeRestoreJob, input).Get(ctx, &output)
    return &output, err
}

func (a *BackupStub) DescribeRestoreJobAsync(ctx workflow.Context, input *backup.DescribeRestoreJobInput) *BackupDescribeRestoreJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeRestoreJob, input)
    return &BackupDescribeRestoreJobResult{Result: future}
}

func (a *BackupStub) ExportBackupPlanTemplate(ctx workflow.Context, input *backup.ExportBackupPlanTemplateInput) (*backup.ExportBackupPlanTemplateOutput, error) {
    var output backup.ExportBackupPlanTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ExportBackupPlanTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *BackupStub) ExportBackupPlanTemplateAsync(ctx workflow.Context, input *backup.ExportBackupPlanTemplateInput) *BackupExportBackupPlanTemplateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ExportBackupPlanTemplate, input)
    return &BackupExportBackupPlanTemplateResult{Result: future}
}

func (a *BackupStub) GetBackupPlan(ctx workflow.Context, input *backup.GetBackupPlanInput) (*backup.GetBackupPlanOutput, error) {
    var output backup.GetBackupPlanOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetBackupPlan, input).Get(ctx, &output)
    return &output, err
}

func (a *BackupStub) GetBackupPlanAsync(ctx workflow.Context, input *backup.GetBackupPlanInput) *BackupGetBackupPlanResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetBackupPlan, input)
    return &BackupGetBackupPlanResult{Result: future}
}

func (a *BackupStub) GetBackupPlanFromJSON(ctx workflow.Context, input *backup.GetBackupPlanFromJSONInput) (*backup.GetBackupPlanFromJSONOutput, error) {
    var output backup.GetBackupPlanFromJSONOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetBackupPlanFromJSON, input).Get(ctx, &output)
    return &output, err
}

func (a *BackupStub) GetBackupPlanFromJSONAsync(ctx workflow.Context, input *backup.GetBackupPlanFromJSONInput) *BackupGetBackupPlanFromJSONResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetBackupPlanFromJSON, input)
    return &BackupGetBackupPlanFromJSONResult{Result: future}
}

func (a *BackupStub) GetBackupPlanFromTemplate(ctx workflow.Context, input *backup.GetBackupPlanFromTemplateInput) (*backup.GetBackupPlanFromTemplateOutput, error) {
    var output backup.GetBackupPlanFromTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetBackupPlanFromTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *BackupStub) GetBackupPlanFromTemplateAsync(ctx workflow.Context, input *backup.GetBackupPlanFromTemplateInput) *BackupGetBackupPlanFromTemplateResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetBackupPlanFromTemplate, input)
    return &BackupGetBackupPlanFromTemplateResult{Result: future}
}

func (a *BackupStub) GetBackupSelection(ctx workflow.Context, input *backup.GetBackupSelectionInput) (*backup.GetBackupSelectionOutput, error) {
    var output backup.GetBackupSelectionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetBackupSelection, input).Get(ctx, &output)
    return &output, err
}

func (a *BackupStub) GetBackupSelectionAsync(ctx workflow.Context, input *backup.GetBackupSelectionInput) *BackupGetBackupSelectionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetBackupSelection, input)
    return &BackupGetBackupSelectionResult{Result: future}
}

func (a *BackupStub) GetBackupVaultAccessPolicy(ctx workflow.Context, input *backup.GetBackupVaultAccessPolicyInput) (*backup.GetBackupVaultAccessPolicyOutput, error) {
    var output backup.GetBackupVaultAccessPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetBackupVaultAccessPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *BackupStub) GetBackupVaultAccessPolicyAsync(ctx workflow.Context, input *backup.GetBackupVaultAccessPolicyInput) *BackupGetBackupVaultAccessPolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetBackupVaultAccessPolicy, input)
    return &BackupGetBackupVaultAccessPolicyResult{Result: future}
}

func (a *BackupStub) GetBackupVaultNotifications(ctx workflow.Context, input *backup.GetBackupVaultNotificationsInput) (*backup.GetBackupVaultNotificationsOutput, error) {
    var output backup.GetBackupVaultNotificationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetBackupVaultNotifications, input).Get(ctx, &output)
    return &output, err
}

func (a *BackupStub) GetBackupVaultNotificationsAsync(ctx workflow.Context, input *backup.GetBackupVaultNotificationsInput) *BackupGetBackupVaultNotificationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetBackupVaultNotifications, input)
    return &BackupGetBackupVaultNotificationsResult{Result: future}
}

func (a *BackupStub) GetRecoveryPointRestoreMetadata(ctx workflow.Context, input *backup.GetRecoveryPointRestoreMetadataInput) (*backup.GetRecoveryPointRestoreMetadataOutput, error) {
    var output backup.GetRecoveryPointRestoreMetadataOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetRecoveryPointRestoreMetadata, input).Get(ctx, &output)
    return &output, err
}

func (a *BackupStub) GetRecoveryPointRestoreMetadataAsync(ctx workflow.Context, input *backup.GetRecoveryPointRestoreMetadataInput) *BackupGetRecoveryPointRestoreMetadataResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetRecoveryPointRestoreMetadata, input)
    return &BackupGetRecoveryPointRestoreMetadataResult{Result: future}
}

func (a *BackupStub) GetSupportedResourceTypes(ctx workflow.Context, input *backup.GetSupportedResourceTypesInput) (*backup.GetSupportedResourceTypesOutput, error) {
    var output backup.GetSupportedResourceTypesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetSupportedResourceTypes, input).Get(ctx, &output)
    return &output, err
}

func (a *BackupStub) GetSupportedResourceTypesAsync(ctx workflow.Context, input *backup.GetSupportedResourceTypesInput) *BackupGetSupportedResourceTypesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetSupportedResourceTypes, input)
    return &BackupGetSupportedResourceTypesResult{Result: future}
}

func (a *BackupStub) ListBackupJobs(ctx workflow.Context, input *backup.ListBackupJobsInput) (*backup.ListBackupJobsOutput, error) {
    var output backup.ListBackupJobsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListBackupJobs, input).Get(ctx, &output)
    return &output, err
}

func (a *BackupStub) ListBackupJobsAsync(ctx workflow.Context, input *backup.ListBackupJobsInput) *BackupListBackupJobsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListBackupJobs, input)
    return &BackupListBackupJobsResult{Result: future}
}

func (a *BackupStub) ListBackupPlanTemplates(ctx workflow.Context, input *backup.ListBackupPlanTemplatesInput) (*backup.ListBackupPlanTemplatesOutput, error) {
    var output backup.ListBackupPlanTemplatesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListBackupPlanTemplates, input).Get(ctx, &output)
    return &output, err
}

func (a *BackupStub) ListBackupPlanTemplatesAsync(ctx workflow.Context, input *backup.ListBackupPlanTemplatesInput) *BackupListBackupPlanTemplatesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListBackupPlanTemplates, input)
    return &BackupListBackupPlanTemplatesResult{Result: future}
}

func (a *BackupStub) ListBackupPlanVersions(ctx workflow.Context, input *backup.ListBackupPlanVersionsInput) (*backup.ListBackupPlanVersionsOutput, error) {
    var output backup.ListBackupPlanVersionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListBackupPlanVersions, input).Get(ctx, &output)
    return &output, err
}

func (a *BackupStub) ListBackupPlanVersionsAsync(ctx workflow.Context, input *backup.ListBackupPlanVersionsInput) *BackupListBackupPlanVersionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListBackupPlanVersions, input)
    return &BackupListBackupPlanVersionsResult{Result: future}
}

func (a *BackupStub) ListBackupPlans(ctx workflow.Context, input *backup.ListBackupPlansInput) (*backup.ListBackupPlansOutput, error) {
    var output backup.ListBackupPlansOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListBackupPlans, input).Get(ctx, &output)
    return &output, err
}

func (a *BackupStub) ListBackupPlansAsync(ctx workflow.Context, input *backup.ListBackupPlansInput) *BackupListBackupPlansResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListBackupPlans, input)
    return &BackupListBackupPlansResult{Result: future}
}

func (a *BackupStub) ListBackupSelections(ctx workflow.Context, input *backup.ListBackupSelectionsInput) (*backup.ListBackupSelectionsOutput, error) {
    var output backup.ListBackupSelectionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListBackupSelections, input).Get(ctx, &output)
    return &output, err
}

func (a *BackupStub) ListBackupSelectionsAsync(ctx workflow.Context, input *backup.ListBackupSelectionsInput) *BackupListBackupSelectionsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListBackupSelections, input)
    return &BackupListBackupSelectionsResult{Result: future}
}

func (a *BackupStub) ListBackupVaults(ctx workflow.Context, input *backup.ListBackupVaultsInput) (*backup.ListBackupVaultsOutput, error) {
    var output backup.ListBackupVaultsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListBackupVaults, input).Get(ctx, &output)
    return &output, err
}

func (a *BackupStub) ListBackupVaultsAsync(ctx workflow.Context, input *backup.ListBackupVaultsInput) *BackupListBackupVaultsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListBackupVaults, input)
    return &BackupListBackupVaultsResult{Result: future}
}

func (a *BackupStub) ListCopyJobs(ctx workflow.Context, input *backup.ListCopyJobsInput) (*backup.ListCopyJobsOutput, error) {
    var output backup.ListCopyJobsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListCopyJobs, input).Get(ctx, &output)
    return &output, err
}

func (a *BackupStub) ListCopyJobsAsync(ctx workflow.Context, input *backup.ListCopyJobsInput) *BackupListCopyJobsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListCopyJobs, input)
    return &BackupListCopyJobsResult{Result: future}
}

func (a *BackupStub) ListProtectedResources(ctx workflow.Context, input *backup.ListProtectedResourcesInput) (*backup.ListProtectedResourcesOutput, error) {
    var output backup.ListProtectedResourcesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListProtectedResources, input).Get(ctx, &output)
    return &output, err
}

func (a *BackupStub) ListProtectedResourcesAsync(ctx workflow.Context, input *backup.ListProtectedResourcesInput) *BackupListProtectedResourcesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListProtectedResources, input)
    return &BackupListProtectedResourcesResult{Result: future}
}

func (a *BackupStub) ListRecoveryPointsByBackupVault(ctx workflow.Context, input *backup.ListRecoveryPointsByBackupVaultInput) (*backup.ListRecoveryPointsByBackupVaultOutput, error) {
    var output backup.ListRecoveryPointsByBackupVaultOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListRecoveryPointsByBackupVault, input).Get(ctx, &output)
    return &output, err
}

func (a *BackupStub) ListRecoveryPointsByBackupVaultAsync(ctx workflow.Context, input *backup.ListRecoveryPointsByBackupVaultInput) *BackupListRecoveryPointsByBackupVaultResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListRecoveryPointsByBackupVault, input)
    return &BackupListRecoveryPointsByBackupVaultResult{Result: future}
}

func (a *BackupStub) ListRecoveryPointsByResource(ctx workflow.Context, input *backup.ListRecoveryPointsByResourceInput) (*backup.ListRecoveryPointsByResourceOutput, error) {
    var output backup.ListRecoveryPointsByResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListRecoveryPointsByResource, input).Get(ctx, &output)
    return &output, err
}

func (a *BackupStub) ListRecoveryPointsByResourceAsync(ctx workflow.Context, input *backup.ListRecoveryPointsByResourceInput) *BackupListRecoveryPointsByResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListRecoveryPointsByResource, input)
    return &BackupListRecoveryPointsByResourceResult{Result: future}
}

func (a *BackupStub) ListRestoreJobs(ctx workflow.Context, input *backup.ListRestoreJobsInput) (*backup.ListRestoreJobsOutput, error) {
    var output backup.ListRestoreJobsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListRestoreJobs, input).Get(ctx, &output)
    return &output, err
}

func (a *BackupStub) ListRestoreJobsAsync(ctx workflow.Context, input *backup.ListRestoreJobsInput) *BackupListRestoreJobsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListRestoreJobs, input)
    return &BackupListRestoreJobsResult{Result: future}
}

func (a *BackupStub) ListTags(ctx workflow.Context, input *backup.ListTagsInput) (*backup.ListTagsOutput, error) {
    var output backup.ListTagsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTags, input).Get(ctx, &output)
    return &output, err
}

func (a *BackupStub) ListTagsAsync(ctx workflow.Context, input *backup.ListTagsInput) *BackupListTagsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTags, input)
    return &BackupListTagsResult{Result: future}
}

func (a *BackupStub) PutBackupVaultAccessPolicy(ctx workflow.Context, input *backup.PutBackupVaultAccessPolicyInput) (*backup.PutBackupVaultAccessPolicyOutput, error) {
    var output backup.PutBackupVaultAccessPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutBackupVaultAccessPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *BackupStub) PutBackupVaultAccessPolicyAsync(ctx workflow.Context, input *backup.PutBackupVaultAccessPolicyInput) *BackupPutBackupVaultAccessPolicyResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutBackupVaultAccessPolicy, input)
    return &BackupPutBackupVaultAccessPolicyResult{Result: future}
}

func (a *BackupStub) PutBackupVaultNotifications(ctx workflow.Context, input *backup.PutBackupVaultNotificationsInput) (*backup.PutBackupVaultNotificationsOutput, error) {
    var output backup.PutBackupVaultNotificationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutBackupVaultNotifications, input).Get(ctx, &output)
    return &output, err
}

func (a *BackupStub) PutBackupVaultNotificationsAsync(ctx workflow.Context, input *backup.PutBackupVaultNotificationsInput) *BackupPutBackupVaultNotificationsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PutBackupVaultNotifications, input)
    return &BackupPutBackupVaultNotificationsResult{Result: future}
}

func (a *BackupStub) StartBackupJob(ctx workflow.Context, input *backup.StartBackupJobInput) (*backup.StartBackupJobOutput, error) {
    var output backup.StartBackupJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartBackupJob, input).Get(ctx, &output)
    return &output, err
}

func (a *BackupStub) StartBackupJobAsync(ctx workflow.Context, input *backup.StartBackupJobInput) *BackupStartBackupJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StartBackupJob, input)
    return &BackupStartBackupJobResult{Result: future}
}

func (a *BackupStub) StartCopyJob(ctx workflow.Context, input *backup.StartCopyJobInput) (*backup.StartCopyJobOutput, error) {
    var output backup.StartCopyJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartCopyJob, input).Get(ctx, &output)
    return &output, err
}

func (a *BackupStub) StartCopyJobAsync(ctx workflow.Context, input *backup.StartCopyJobInput) *BackupStartCopyJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StartCopyJob, input)
    return &BackupStartCopyJobResult{Result: future}
}

func (a *BackupStub) StartRestoreJob(ctx workflow.Context, input *backup.StartRestoreJobInput) (*backup.StartRestoreJobOutput, error) {
    var output backup.StartRestoreJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartRestoreJob, input).Get(ctx, &output)
    return &output, err
}

func (a *BackupStub) StartRestoreJobAsync(ctx workflow.Context, input *backup.StartRestoreJobInput) *BackupStartRestoreJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StartRestoreJob, input)
    return &BackupStartRestoreJobResult{Result: future}
}

func (a *BackupStub) StopBackupJob(ctx workflow.Context, input *backup.StopBackupJobInput) (*backup.StopBackupJobOutput, error) {
    var output backup.StopBackupJobOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StopBackupJob, input).Get(ctx, &output)
    return &output, err
}

func (a *BackupStub) StopBackupJobAsync(ctx workflow.Context, input *backup.StopBackupJobInput) *BackupStopBackupJobResult {
    future := workflow.ExecuteActivity(ctx, a.activities.StopBackupJob, input)
    return &BackupStopBackupJobResult{Result: future}
}

func (a *BackupStub) TagResource(ctx workflow.Context, input *backup.TagResourceInput) (*backup.TagResourceOutput, error) {
    var output backup.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *BackupStub) TagResourceAsync(ctx workflow.Context, input *backup.TagResourceInput) *BackupTagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TagResource, input)
    return &BackupTagResourceResult{Result: future}
}

func (a *BackupStub) UntagResource(ctx workflow.Context, input *backup.UntagResourceInput) (*backup.UntagResourceOutput, error) {
    var output backup.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *BackupStub) UntagResourceAsync(ctx workflow.Context, input *backup.UntagResourceInput) *BackupUntagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input)
    return &BackupUntagResourceResult{Result: future}
}

func (a *BackupStub) UpdateBackupPlan(ctx workflow.Context, input *backup.UpdateBackupPlanInput) (*backup.UpdateBackupPlanOutput, error) {
    var output backup.UpdateBackupPlanOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateBackupPlan, input).Get(ctx, &output)
    return &output, err
}

func (a *BackupStub) UpdateBackupPlanAsync(ctx workflow.Context, input *backup.UpdateBackupPlanInput) *BackupUpdateBackupPlanResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateBackupPlan, input)
    return &BackupUpdateBackupPlanResult{Result: future}
}

func (a *BackupStub) UpdateRecoveryPointLifecycle(ctx workflow.Context, input *backup.UpdateRecoveryPointLifecycleInput) (*backup.UpdateRecoveryPointLifecycleOutput, error) {
    var output backup.UpdateRecoveryPointLifecycleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateRecoveryPointLifecycle, input).Get(ctx, &output)
    return &output, err
}

func (a *BackupStub) UpdateRecoveryPointLifecycleAsync(ctx workflow.Context, input *backup.UpdateRecoveryPointLifecycleInput) *BackupUpdateRecoveryPointLifecycleResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateRecoveryPointLifecycle, input)
    return &BackupUpdateRecoveryPointLifecycleResult{Result: future}
}

func (a *BackupStub) UpdateRegionSettings(ctx workflow.Context, input *backup.UpdateRegionSettingsInput) (*backup.UpdateRegionSettingsOutput, error) {
    var output backup.UpdateRegionSettingsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateRegionSettings, input).Get(ctx, &output)
    return &output, err
}

func (a *BackupStub) UpdateRegionSettingsAsync(ctx workflow.Context, input *backup.UpdateRegionSettingsInput) *BackupUpdateRegionSettingsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UpdateRegionSettings, input)
    return &BackupUpdateRegionSettingsResult{Result: future}
}