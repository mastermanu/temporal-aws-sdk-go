package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/applicationdiscoveryservice"
	"go.temporal.io/sdk/workflow"
)

type ApplicationDiscoveryServiceClient interface {
    AssociateConfigurationItemsToApplication(ctx workflow.Context, input *applicationdiscoveryservice.AssociateConfigurationItemsToApplicationInput) (*applicationdiscoveryservice.AssociateConfigurationItemsToApplicationOutput, error)
    AssociateConfigurationItemsToApplicationAsync(ctx workflow.Context, input *applicationdiscoveryservice.AssociateConfigurationItemsToApplicationInput) *ApplicationdiscoveryserviceAssociateConfigurationItemsToApplicationResult

    BatchDeleteImportData(ctx workflow.Context, input *applicationdiscoveryservice.BatchDeleteImportDataInput) (*applicationdiscoveryservice.BatchDeleteImportDataOutput, error)
    BatchDeleteImportDataAsync(ctx workflow.Context, input *applicationdiscoveryservice.BatchDeleteImportDataInput) *ApplicationdiscoveryserviceBatchDeleteImportDataResult

    CreateApplication(ctx workflow.Context, input *applicationdiscoveryservice.CreateApplicationInput) (*applicationdiscoveryservice.CreateApplicationOutput, error)
    CreateApplicationAsync(ctx workflow.Context, input *applicationdiscoveryservice.CreateApplicationInput) *ApplicationdiscoveryserviceCreateApplicationResult

    CreateTags(ctx workflow.Context, input *applicationdiscoveryservice.CreateTagsInput) (*applicationdiscoveryservice.CreateTagsOutput, error)
    CreateTagsAsync(ctx workflow.Context, input *applicationdiscoveryservice.CreateTagsInput) *ApplicationdiscoveryserviceCreateTagsResult

    DeleteApplications(ctx workflow.Context, input *applicationdiscoveryservice.DeleteApplicationsInput) (*applicationdiscoveryservice.DeleteApplicationsOutput, error)
    DeleteApplicationsAsync(ctx workflow.Context, input *applicationdiscoveryservice.DeleteApplicationsInput) *ApplicationdiscoveryserviceDeleteApplicationsResult

    DeleteTags(ctx workflow.Context, input *applicationdiscoveryservice.DeleteTagsInput) (*applicationdiscoveryservice.DeleteTagsOutput, error)
    DeleteTagsAsync(ctx workflow.Context, input *applicationdiscoveryservice.DeleteTagsInput) *ApplicationdiscoveryserviceDeleteTagsResult

    DescribeAgents(ctx workflow.Context, input *applicationdiscoveryservice.DescribeAgentsInput) (*applicationdiscoveryservice.DescribeAgentsOutput, error)
    DescribeAgentsAsync(ctx workflow.Context, input *applicationdiscoveryservice.DescribeAgentsInput) *ApplicationdiscoveryserviceDescribeAgentsResult

    DescribeConfigurations(ctx workflow.Context, input *applicationdiscoveryservice.DescribeConfigurationsInput) (*applicationdiscoveryservice.DescribeConfigurationsOutput, error)
    DescribeConfigurationsAsync(ctx workflow.Context, input *applicationdiscoveryservice.DescribeConfigurationsInput) *ApplicationdiscoveryserviceDescribeConfigurationsResult

    DescribeContinuousExports(ctx workflow.Context, input *applicationdiscoveryservice.DescribeContinuousExportsInput) (*applicationdiscoveryservice.DescribeContinuousExportsOutput, error)
    DescribeContinuousExportsAsync(ctx workflow.Context, input *applicationdiscoveryservice.DescribeContinuousExportsInput) *ApplicationdiscoveryserviceDescribeContinuousExportsResult

    DescribeExportConfigurations(ctx workflow.Context, input *applicationdiscoveryservice.DescribeExportConfigurationsInput) (*applicationdiscoveryservice.DescribeExportConfigurationsOutput, error)
    DescribeExportConfigurationsAsync(ctx workflow.Context, input *applicationdiscoveryservice.DescribeExportConfigurationsInput) *ApplicationdiscoveryserviceDescribeExportConfigurationsResult

    DescribeExportTasks(ctx workflow.Context, input *applicationdiscoveryservice.DescribeExportTasksInput) (*applicationdiscoveryservice.DescribeExportTasksOutput, error)
    DescribeExportTasksAsync(ctx workflow.Context, input *applicationdiscoveryservice.DescribeExportTasksInput) *ApplicationdiscoveryserviceDescribeExportTasksResult

    DescribeImportTasks(ctx workflow.Context, input *applicationdiscoveryservice.DescribeImportTasksInput) (*applicationdiscoveryservice.DescribeImportTasksOutput, error)
    DescribeImportTasksAsync(ctx workflow.Context, input *applicationdiscoveryservice.DescribeImportTasksInput) *ApplicationdiscoveryserviceDescribeImportTasksResult

    DescribeTags(ctx workflow.Context, input *applicationdiscoveryservice.DescribeTagsInput) (*applicationdiscoveryservice.DescribeTagsOutput, error)
    DescribeTagsAsync(ctx workflow.Context, input *applicationdiscoveryservice.DescribeTagsInput) *ApplicationdiscoveryserviceDescribeTagsResult

    DisassociateConfigurationItemsFromApplication(ctx workflow.Context, input *applicationdiscoveryservice.DisassociateConfigurationItemsFromApplicationInput) (*applicationdiscoveryservice.DisassociateConfigurationItemsFromApplicationOutput, error)
    DisassociateConfigurationItemsFromApplicationAsync(ctx workflow.Context, input *applicationdiscoveryservice.DisassociateConfigurationItemsFromApplicationInput) *ApplicationdiscoveryserviceDisassociateConfigurationItemsFromApplicationResult

    ExportConfigurations(ctx workflow.Context, input *applicationdiscoveryservice.ExportConfigurationsInput) (*applicationdiscoveryservice.ExportConfigurationsOutput, error)
    ExportConfigurationsAsync(ctx workflow.Context, input *applicationdiscoveryservice.ExportConfigurationsInput) *ApplicationdiscoveryserviceExportConfigurationsResult

    GetDiscoverySummary(ctx workflow.Context, input *applicationdiscoveryservice.GetDiscoverySummaryInput) (*applicationdiscoveryservice.GetDiscoverySummaryOutput, error)
    GetDiscoverySummaryAsync(ctx workflow.Context, input *applicationdiscoveryservice.GetDiscoverySummaryInput) *ApplicationdiscoveryserviceGetDiscoverySummaryResult

    ListConfigurations(ctx workflow.Context, input *applicationdiscoveryservice.ListConfigurationsInput) (*applicationdiscoveryservice.ListConfigurationsOutput, error)
    ListConfigurationsAsync(ctx workflow.Context, input *applicationdiscoveryservice.ListConfigurationsInput) *ApplicationdiscoveryserviceListConfigurationsResult

    ListServerNeighbors(ctx workflow.Context, input *applicationdiscoveryservice.ListServerNeighborsInput) (*applicationdiscoveryservice.ListServerNeighborsOutput, error)
    ListServerNeighborsAsync(ctx workflow.Context, input *applicationdiscoveryservice.ListServerNeighborsInput) *ApplicationdiscoveryserviceListServerNeighborsResult

    StartContinuousExport(ctx workflow.Context, input *applicationdiscoveryservice.StartContinuousExportInput) (*applicationdiscoveryservice.StartContinuousExportOutput, error)
    StartContinuousExportAsync(ctx workflow.Context, input *applicationdiscoveryservice.StartContinuousExportInput) *ApplicationdiscoveryserviceStartContinuousExportResult

    StartDataCollectionByAgentIds(ctx workflow.Context, input *applicationdiscoveryservice.StartDataCollectionByAgentIdsInput) (*applicationdiscoveryservice.StartDataCollectionByAgentIdsOutput, error)
    StartDataCollectionByAgentIdsAsync(ctx workflow.Context, input *applicationdiscoveryservice.StartDataCollectionByAgentIdsInput) *ApplicationdiscoveryserviceStartDataCollectionByAgentIdsResult

    StartExportTask(ctx workflow.Context, input *applicationdiscoveryservice.StartExportTaskInput) (*applicationdiscoveryservice.StartExportTaskOutput, error)
    StartExportTaskAsync(ctx workflow.Context, input *applicationdiscoveryservice.StartExportTaskInput) *ApplicationdiscoveryserviceStartExportTaskResult

    StartImportTask(ctx workflow.Context, input *applicationdiscoveryservice.StartImportTaskInput) (*applicationdiscoveryservice.StartImportTaskOutput, error)
    StartImportTaskAsync(ctx workflow.Context, input *applicationdiscoveryservice.StartImportTaskInput) *ApplicationdiscoveryserviceStartImportTaskResult

    StopContinuousExport(ctx workflow.Context, input *applicationdiscoveryservice.StopContinuousExportInput) (*applicationdiscoveryservice.StopContinuousExportOutput, error)
    StopContinuousExportAsync(ctx workflow.Context, input *applicationdiscoveryservice.StopContinuousExportInput) *ApplicationdiscoveryserviceStopContinuousExportResult

    StopDataCollectionByAgentIds(ctx workflow.Context, input *applicationdiscoveryservice.StopDataCollectionByAgentIdsInput) (*applicationdiscoveryservice.StopDataCollectionByAgentIdsOutput, error)
    StopDataCollectionByAgentIdsAsync(ctx workflow.Context, input *applicationdiscoveryservice.StopDataCollectionByAgentIdsInput) *ApplicationdiscoveryserviceStopDataCollectionByAgentIdsResult

    UpdateApplication(ctx workflow.Context, input *applicationdiscoveryservice.UpdateApplicationInput) (*applicationdiscoveryservice.UpdateApplicationOutput, error)
    UpdateApplicationAsync(ctx workflow.Context, input *applicationdiscoveryservice.UpdateApplicationInput) *ApplicationdiscoveryserviceUpdateApplicationResult
}
type ApplicationdiscoveryserviceAssociateConfigurationItemsToApplicationResult struct {
	Result workflow.Future
}

func (r *ApplicationdiscoveryserviceAssociateConfigurationItemsToApplicationResult) Get(ctx workflow.Context) (*applicationdiscoveryservice.AssociateConfigurationItemsToApplicationOutput, error) {
    var output applicationdiscoveryservice.AssociateConfigurationItemsToApplicationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ApplicationdiscoveryserviceBatchDeleteImportDataResult struct {
	Result workflow.Future
}

func (r *ApplicationdiscoveryserviceBatchDeleteImportDataResult) Get(ctx workflow.Context) (*applicationdiscoveryservice.BatchDeleteImportDataOutput, error) {
    var output applicationdiscoveryservice.BatchDeleteImportDataOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ApplicationdiscoveryserviceCreateApplicationResult struct {
	Result workflow.Future
}

func (r *ApplicationdiscoveryserviceCreateApplicationResult) Get(ctx workflow.Context) (*applicationdiscoveryservice.CreateApplicationOutput, error) {
    var output applicationdiscoveryservice.CreateApplicationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ApplicationdiscoveryserviceCreateTagsResult struct {
	Result workflow.Future
}

func (r *ApplicationdiscoveryserviceCreateTagsResult) Get(ctx workflow.Context) (*applicationdiscoveryservice.CreateTagsOutput, error) {
    var output applicationdiscoveryservice.CreateTagsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ApplicationdiscoveryserviceDeleteApplicationsResult struct {
	Result workflow.Future
}

func (r *ApplicationdiscoveryserviceDeleteApplicationsResult) Get(ctx workflow.Context) (*applicationdiscoveryservice.DeleteApplicationsOutput, error) {
    var output applicationdiscoveryservice.DeleteApplicationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ApplicationdiscoveryserviceDeleteTagsResult struct {
	Result workflow.Future
}

func (r *ApplicationdiscoveryserviceDeleteTagsResult) Get(ctx workflow.Context) (*applicationdiscoveryservice.DeleteTagsOutput, error) {
    var output applicationdiscoveryservice.DeleteTagsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ApplicationdiscoveryserviceDescribeAgentsResult struct {
	Result workflow.Future
}

func (r *ApplicationdiscoveryserviceDescribeAgentsResult) Get(ctx workflow.Context) (*applicationdiscoveryservice.DescribeAgentsOutput, error) {
    var output applicationdiscoveryservice.DescribeAgentsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ApplicationdiscoveryserviceDescribeConfigurationsResult struct {
	Result workflow.Future
}

func (r *ApplicationdiscoveryserviceDescribeConfigurationsResult) Get(ctx workflow.Context) (*applicationdiscoveryservice.DescribeConfigurationsOutput, error) {
    var output applicationdiscoveryservice.DescribeConfigurationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ApplicationdiscoveryserviceDescribeContinuousExportsResult struct {
	Result workflow.Future
}

func (r *ApplicationdiscoveryserviceDescribeContinuousExportsResult) Get(ctx workflow.Context) (*applicationdiscoveryservice.DescribeContinuousExportsOutput, error) {
    var output applicationdiscoveryservice.DescribeContinuousExportsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ApplicationdiscoveryserviceDescribeExportConfigurationsResult struct {
	Result workflow.Future
}

func (r *ApplicationdiscoveryserviceDescribeExportConfigurationsResult) Get(ctx workflow.Context) (*applicationdiscoveryservice.DescribeExportConfigurationsOutput, error) {
    var output applicationdiscoveryservice.DescribeExportConfigurationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ApplicationdiscoveryserviceDescribeExportTasksResult struct {
	Result workflow.Future
}

func (r *ApplicationdiscoveryserviceDescribeExportTasksResult) Get(ctx workflow.Context) (*applicationdiscoveryservice.DescribeExportTasksOutput, error) {
    var output applicationdiscoveryservice.DescribeExportTasksOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ApplicationdiscoveryserviceDescribeImportTasksResult struct {
	Result workflow.Future
}

func (r *ApplicationdiscoveryserviceDescribeImportTasksResult) Get(ctx workflow.Context) (*applicationdiscoveryservice.DescribeImportTasksOutput, error) {
    var output applicationdiscoveryservice.DescribeImportTasksOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ApplicationdiscoveryserviceDescribeTagsResult struct {
	Result workflow.Future
}

func (r *ApplicationdiscoveryserviceDescribeTagsResult) Get(ctx workflow.Context) (*applicationdiscoveryservice.DescribeTagsOutput, error) {
    var output applicationdiscoveryservice.DescribeTagsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ApplicationdiscoveryserviceDisassociateConfigurationItemsFromApplicationResult struct {
	Result workflow.Future
}

func (r *ApplicationdiscoveryserviceDisassociateConfigurationItemsFromApplicationResult) Get(ctx workflow.Context) (*applicationdiscoveryservice.DisassociateConfigurationItemsFromApplicationOutput, error) {
    var output applicationdiscoveryservice.DisassociateConfigurationItemsFromApplicationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ApplicationdiscoveryserviceExportConfigurationsResult struct {
	Result workflow.Future
}

func (r *ApplicationdiscoveryserviceExportConfigurationsResult) Get(ctx workflow.Context) (*applicationdiscoveryservice.ExportConfigurationsOutput, error) {
    var output applicationdiscoveryservice.ExportConfigurationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ApplicationdiscoveryserviceGetDiscoverySummaryResult struct {
	Result workflow.Future
}

func (r *ApplicationdiscoveryserviceGetDiscoverySummaryResult) Get(ctx workflow.Context) (*applicationdiscoveryservice.GetDiscoverySummaryOutput, error) {
    var output applicationdiscoveryservice.GetDiscoverySummaryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ApplicationdiscoveryserviceListConfigurationsResult struct {
	Result workflow.Future
}

func (r *ApplicationdiscoveryserviceListConfigurationsResult) Get(ctx workflow.Context) (*applicationdiscoveryservice.ListConfigurationsOutput, error) {
    var output applicationdiscoveryservice.ListConfigurationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ApplicationdiscoveryserviceListServerNeighborsResult struct {
	Result workflow.Future
}

func (r *ApplicationdiscoveryserviceListServerNeighborsResult) Get(ctx workflow.Context) (*applicationdiscoveryservice.ListServerNeighborsOutput, error) {
    var output applicationdiscoveryservice.ListServerNeighborsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ApplicationdiscoveryserviceStartContinuousExportResult struct {
	Result workflow.Future
}

func (r *ApplicationdiscoveryserviceStartContinuousExportResult) Get(ctx workflow.Context) (*applicationdiscoveryservice.StartContinuousExportOutput, error) {
    var output applicationdiscoveryservice.StartContinuousExportOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ApplicationdiscoveryserviceStartDataCollectionByAgentIdsResult struct {
	Result workflow.Future
}

func (r *ApplicationdiscoveryserviceStartDataCollectionByAgentIdsResult) Get(ctx workflow.Context) (*applicationdiscoveryservice.StartDataCollectionByAgentIdsOutput, error) {
    var output applicationdiscoveryservice.StartDataCollectionByAgentIdsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ApplicationdiscoveryserviceStartExportTaskResult struct {
	Result workflow.Future
}

func (r *ApplicationdiscoveryserviceStartExportTaskResult) Get(ctx workflow.Context) (*applicationdiscoveryservice.StartExportTaskOutput, error) {
    var output applicationdiscoveryservice.StartExportTaskOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ApplicationdiscoveryserviceStartImportTaskResult struct {
	Result workflow.Future
}

func (r *ApplicationdiscoveryserviceStartImportTaskResult) Get(ctx workflow.Context) (*applicationdiscoveryservice.StartImportTaskOutput, error) {
    var output applicationdiscoveryservice.StartImportTaskOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ApplicationdiscoveryserviceStopContinuousExportResult struct {
	Result workflow.Future
}

func (r *ApplicationdiscoveryserviceStopContinuousExportResult) Get(ctx workflow.Context) (*applicationdiscoveryservice.StopContinuousExportOutput, error) {
    var output applicationdiscoveryservice.StopContinuousExportOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ApplicationdiscoveryserviceStopDataCollectionByAgentIdsResult struct {
	Result workflow.Future
}

func (r *ApplicationdiscoveryserviceStopDataCollectionByAgentIdsResult) Get(ctx workflow.Context) (*applicationdiscoveryservice.StopDataCollectionByAgentIdsOutput, error) {
    var output applicationdiscoveryservice.StopDataCollectionByAgentIdsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type ApplicationdiscoveryserviceUpdateApplicationResult struct {
	Result workflow.Future
}

func (r *ApplicationdiscoveryserviceUpdateApplicationResult) Get(ctx workflow.Context) (*applicationdiscoveryservice.UpdateApplicationOutput, error) {
    var output applicationdiscoveryservice.UpdateApplicationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type ApplicationDiscoveryServiceStub struct {
    activities ApplicationDiscoveryServiceClient
}

func NewApplicationDiscoveryServiceStub() ApplicationDiscoveryServiceClient {
    return &ApplicationDiscoveryServiceStub{}
}

func (a *ApplicationDiscoveryServiceStub) AssociateConfigurationItemsToApplication(ctx workflow.Context, input *applicationdiscoveryservice.AssociateConfigurationItemsToApplicationInput) (*applicationdiscoveryservice.AssociateConfigurationItemsToApplicationOutput, error) {
    var output applicationdiscoveryservice.AssociateConfigurationItemsToApplicationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AssociateConfigurationItemsToApplication, input).Get(ctx, &output)
    return &output, err
}

func (a *ApplicationDiscoveryServiceStub) BatchDeleteImportData(ctx workflow.Context, input *applicationdiscoveryservice.BatchDeleteImportDataInput) (*applicationdiscoveryservice.BatchDeleteImportDataOutput, error) {
    var output applicationdiscoveryservice.BatchDeleteImportDataOutput
    err := workflow.ExecuteActivity(ctx, a.activities.BatchDeleteImportData, input).Get(ctx, &output)
    return &output, err
}

func (a *ApplicationDiscoveryServiceStub) CreateApplication(ctx workflow.Context, input *applicationdiscoveryservice.CreateApplicationInput) (*applicationdiscoveryservice.CreateApplicationOutput, error) {
    var output applicationdiscoveryservice.CreateApplicationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateApplication, input).Get(ctx, &output)
    return &output, err
}

func (a *ApplicationDiscoveryServiceStub) CreateTags(ctx workflow.Context, input *applicationdiscoveryservice.CreateTagsInput) (*applicationdiscoveryservice.CreateTagsOutput, error) {
    var output applicationdiscoveryservice.CreateTagsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateTags, input).Get(ctx, &output)
    return &output, err
}

func (a *ApplicationDiscoveryServiceStub) DeleteApplications(ctx workflow.Context, input *applicationdiscoveryservice.DeleteApplicationsInput) (*applicationdiscoveryservice.DeleteApplicationsOutput, error) {
    var output applicationdiscoveryservice.DeleteApplicationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteApplications, input).Get(ctx, &output)
    return &output, err
}

func (a *ApplicationDiscoveryServiceStub) DeleteTags(ctx workflow.Context, input *applicationdiscoveryservice.DeleteTagsInput) (*applicationdiscoveryservice.DeleteTagsOutput, error) {
    var output applicationdiscoveryservice.DeleteTagsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteTags, input).Get(ctx, &output)
    return &output, err
}

func (a *ApplicationDiscoveryServiceStub) DescribeAgents(ctx workflow.Context, input *applicationdiscoveryservice.DescribeAgentsInput) (*applicationdiscoveryservice.DescribeAgentsOutput, error) {
    var output applicationdiscoveryservice.DescribeAgentsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeAgents, input).Get(ctx, &output)
    return &output, err
}

func (a *ApplicationDiscoveryServiceStub) DescribeConfigurations(ctx workflow.Context, input *applicationdiscoveryservice.DescribeConfigurationsInput) (*applicationdiscoveryservice.DescribeConfigurationsOutput, error) {
    var output applicationdiscoveryservice.DescribeConfigurationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeConfigurations, input).Get(ctx, &output)
    return &output, err
}

func (a *ApplicationDiscoveryServiceStub) DescribeContinuousExports(ctx workflow.Context, input *applicationdiscoveryservice.DescribeContinuousExportsInput) (*applicationdiscoveryservice.DescribeContinuousExportsOutput, error) {
    var output applicationdiscoveryservice.DescribeContinuousExportsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeContinuousExports, input).Get(ctx, &output)
    return &output, err
}

func (a *ApplicationDiscoveryServiceStub) DescribeExportConfigurations(ctx workflow.Context, input *applicationdiscoveryservice.DescribeExportConfigurationsInput) (*applicationdiscoveryservice.DescribeExportConfigurationsOutput, error) {
    var output applicationdiscoveryservice.DescribeExportConfigurationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeExportConfigurations, input).Get(ctx, &output)
    return &output, err
}

func (a *ApplicationDiscoveryServiceStub) DescribeExportTasks(ctx workflow.Context, input *applicationdiscoveryservice.DescribeExportTasksInput) (*applicationdiscoveryservice.DescribeExportTasksOutput, error) {
    var output applicationdiscoveryservice.DescribeExportTasksOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeExportTasks, input).Get(ctx, &output)
    return &output, err
}

func (a *ApplicationDiscoveryServiceStub) DescribeImportTasks(ctx workflow.Context, input *applicationdiscoveryservice.DescribeImportTasksInput) (*applicationdiscoveryservice.DescribeImportTasksOutput, error) {
    var output applicationdiscoveryservice.DescribeImportTasksOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeImportTasks, input).Get(ctx, &output)
    return &output, err
}

func (a *ApplicationDiscoveryServiceStub) DescribeTags(ctx workflow.Context, input *applicationdiscoveryservice.DescribeTagsInput) (*applicationdiscoveryservice.DescribeTagsOutput, error) {
    var output applicationdiscoveryservice.DescribeTagsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeTags, input).Get(ctx, &output)
    return &output, err
}

func (a *ApplicationDiscoveryServiceStub) DisassociateConfigurationItemsFromApplication(ctx workflow.Context, input *applicationdiscoveryservice.DisassociateConfigurationItemsFromApplicationInput) (*applicationdiscoveryservice.DisassociateConfigurationItemsFromApplicationOutput, error) {
    var output applicationdiscoveryservice.DisassociateConfigurationItemsFromApplicationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisassociateConfigurationItemsFromApplication, input).Get(ctx, &output)
    return &output, err
}

func (a *ApplicationDiscoveryServiceStub) ExportConfigurations(ctx workflow.Context, input *applicationdiscoveryservice.ExportConfigurationsInput) (*applicationdiscoveryservice.ExportConfigurationsOutput, error) {
    var output applicationdiscoveryservice.ExportConfigurationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ExportConfigurations, input).Get(ctx, &output)
    return &output, err
}

func (a *ApplicationDiscoveryServiceStub) GetDiscoverySummary(ctx workflow.Context, input *applicationdiscoveryservice.GetDiscoverySummaryInput) (*applicationdiscoveryservice.GetDiscoverySummaryOutput, error) {
    var output applicationdiscoveryservice.GetDiscoverySummaryOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetDiscoverySummary, input).Get(ctx, &output)
    return &output, err
}

func (a *ApplicationDiscoveryServiceStub) ListConfigurations(ctx workflow.Context, input *applicationdiscoveryservice.ListConfigurationsInput) (*applicationdiscoveryservice.ListConfigurationsOutput, error) {
    var output applicationdiscoveryservice.ListConfigurationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListConfigurations, input).Get(ctx, &output)
    return &output, err
}

func (a *ApplicationDiscoveryServiceStub) ListServerNeighbors(ctx workflow.Context, input *applicationdiscoveryservice.ListServerNeighborsInput) (*applicationdiscoveryservice.ListServerNeighborsOutput, error) {
    var output applicationdiscoveryservice.ListServerNeighborsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListServerNeighbors, input).Get(ctx, &output)
    return &output, err
}

func (a *ApplicationDiscoveryServiceStub) StartContinuousExport(ctx workflow.Context, input *applicationdiscoveryservice.StartContinuousExportInput) (*applicationdiscoveryservice.StartContinuousExportOutput, error) {
    var output applicationdiscoveryservice.StartContinuousExportOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartContinuousExport, input).Get(ctx, &output)
    return &output, err
}

func (a *ApplicationDiscoveryServiceStub) StartDataCollectionByAgentIds(ctx workflow.Context, input *applicationdiscoveryservice.StartDataCollectionByAgentIdsInput) (*applicationdiscoveryservice.StartDataCollectionByAgentIdsOutput, error) {
    var output applicationdiscoveryservice.StartDataCollectionByAgentIdsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartDataCollectionByAgentIds, input).Get(ctx, &output)
    return &output, err
}

func (a *ApplicationDiscoveryServiceStub) StartExportTask(ctx workflow.Context, input *applicationdiscoveryservice.StartExportTaskInput) (*applicationdiscoveryservice.StartExportTaskOutput, error) {
    var output applicationdiscoveryservice.StartExportTaskOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartExportTask, input).Get(ctx, &output)
    return &output, err
}

func (a *ApplicationDiscoveryServiceStub) StartImportTask(ctx workflow.Context, input *applicationdiscoveryservice.StartImportTaskInput) (*applicationdiscoveryservice.StartImportTaskOutput, error) {
    var output applicationdiscoveryservice.StartImportTaskOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartImportTask, input).Get(ctx, &output)
    return &output, err
}

func (a *ApplicationDiscoveryServiceStub) StopContinuousExport(ctx workflow.Context, input *applicationdiscoveryservice.StopContinuousExportInput) (*applicationdiscoveryservice.StopContinuousExportOutput, error) {
    var output applicationdiscoveryservice.StopContinuousExportOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StopContinuousExport, input).Get(ctx, &output)
    return &output, err
}

func (a *ApplicationDiscoveryServiceStub) StopDataCollectionByAgentIds(ctx workflow.Context, input *applicationdiscoveryservice.StopDataCollectionByAgentIdsInput) (*applicationdiscoveryservice.StopDataCollectionByAgentIdsOutput, error) {
    var output applicationdiscoveryservice.StopDataCollectionByAgentIdsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StopDataCollectionByAgentIds, input).Get(ctx, &output)
    return &output, err
}

func (a *ApplicationDiscoveryServiceStub) UpdateApplication(ctx workflow.Context, input *applicationdiscoveryservice.UpdateApplicationInput) (*applicationdiscoveryservice.UpdateApplicationOutput, error) {
    var output applicationdiscoveryservice.UpdateApplicationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateApplication, input).Get(ctx, &output)
    return &output, err
}