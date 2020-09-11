package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/databasemigrationservice"
	"github.com/aws/aws-sdk-go/service/databasemigrationservice/databasemigrationserviceiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type DatabaseMigrationServiceActivities struct {
	client databasemigrationserviceiface.DatabaseMigrationServiceAPI
}

func NewDatabaseMigrationServiceActivities(session *session.Session, config ...*aws.Config) *DatabaseMigrationServiceActivities {
	client := databasemigrationservice.New(session, config...)
	return &DatabaseMigrationServiceActivities{client: client}
}

func (a *DatabaseMigrationServiceActivities) AddTagsToResource(ctx context.Context, input *databasemigrationservice.AddTagsToResourceInput) (*databasemigrationservice.AddTagsToResourceOutput, error) {
	return a.client.AddTagsToResourceWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) ApplyPendingMaintenanceAction(ctx context.Context, input *databasemigrationservice.ApplyPendingMaintenanceActionInput) (*databasemigrationservice.ApplyPendingMaintenanceActionOutput, error) {
	return a.client.ApplyPendingMaintenanceActionWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) CancelReplicationTaskAssessmentRun(ctx context.Context, input *databasemigrationservice.CancelReplicationTaskAssessmentRunInput) (*databasemigrationservice.CancelReplicationTaskAssessmentRunOutput, error) {
	return a.client.CancelReplicationTaskAssessmentRunWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) CreateEndpoint(ctx context.Context, input *databasemigrationservice.CreateEndpointInput) (*databasemigrationservice.CreateEndpointOutput, error) {
	return a.client.CreateEndpointWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) CreateEventSubscription(ctx context.Context, input *databasemigrationservice.CreateEventSubscriptionInput) (*databasemigrationservice.CreateEventSubscriptionOutput, error) {
	return a.client.CreateEventSubscriptionWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) CreateReplicationInstance(ctx context.Context, input *databasemigrationservice.CreateReplicationInstanceInput) (*databasemigrationservice.CreateReplicationInstanceOutput, error) {
	return a.client.CreateReplicationInstanceWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) CreateReplicationSubnetGroup(ctx context.Context, input *databasemigrationservice.CreateReplicationSubnetGroupInput) (*databasemigrationservice.CreateReplicationSubnetGroupOutput, error) {
	return a.client.CreateReplicationSubnetGroupWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) CreateReplicationTask(ctx context.Context, input *databasemigrationservice.CreateReplicationTaskInput) (*databasemigrationservice.CreateReplicationTaskOutput, error) {
	return a.client.CreateReplicationTaskWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) DeleteCertificate(ctx context.Context, input *databasemigrationservice.DeleteCertificateInput) (*databasemigrationservice.DeleteCertificateOutput, error) {
	return a.client.DeleteCertificateWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) DeleteConnection(ctx context.Context, input *databasemigrationservice.DeleteConnectionInput) (*databasemigrationservice.DeleteConnectionOutput, error) {
	return a.client.DeleteConnectionWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) DeleteEndpoint(ctx context.Context, input *databasemigrationservice.DeleteEndpointInput) (*databasemigrationservice.DeleteEndpointOutput, error) {
	return a.client.DeleteEndpointWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) DeleteEventSubscription(ctx context.Context, input *databasemigrationservice.DeleteEventSubscriptionInput) (*databasemigrationservice.DeleteEventSubscriptionOutput, error) {
	return a.client.DeleteEventSubscriptionWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) DeleteReplicationInstance(ctx context.Context, input *databasemigrationservice.DeleteReplicationInstanceInput) (*databasemigrationservice.DeleteReplicationInstanceOutput, error) {
	return a.client.DeleteReplicationInstanceWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) DeleteReplicationSubnetGroup(ctx context.Context, input *databasemigrationservice.DeleteReplicationSubnetGroupInput) (*databasemigrationservice.DeleteReplicationSubnetGroupOutput, error) {
	return a.client.DeleteReplicationSubnetGroupWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) DeleteReplicationTask(ctx context.Context, input *databasemigrationservice.DeleteReplicationTaskInput) (*databasemigrationservice.DeleteReplicationTaskOutput, error) {
	return a.client.DeleteReplicationTaskWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) DeleteReplicationTaskAssessmentRun(ctx context.Context, input *databasemigrationservice.DeleteReplicationTaskAssessmentRunInput) (*databasemigrationservice.DeleteReplicationTaskAssessmentRunOutput, error) {
	return a.client.DeleteReplicationTaskAssessmentRunWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) DescribeAccountAttributes(ctx context.Context, input *databasemigrationservice.DescribeAccountAttributesInput) (*databasemigrationservice.DescribeAccountAttributesOutput, error) {
	return a.client.DescribeAccountAttributesWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) DescribeApplicableIndividualAssessments(ctx context.Context, input *databasemigrationservice.DescribeApplicableIndividualAssessmentsInput) (*databasemigrationservice.DescribeApplicableIndividualAssessmentsOutput, error) {
	return a.client.DescribeApplicableIndividualAssessmentsWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) DescribeCertificates(ctx context.Context, input *databasemigrationservice.DescribeCertificatesInput) (*databasemigrationservice.DescribeCertificatesOutput, error) {
	return a.client.DescribeCertificatesWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) DescribeConnections(ctx context.Context, input *databasemigrationservice.DescribeConnectionsInput) (*databasemigrationservice.DescribeConnectionsOutput, error) {
	return a.client.DescribeConnectionsWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) DescribeEndpointTypes(ctx context.Context, input *databasemigrationservice.DescribeEndpointTypesInput) (*databasemigrationservice.DescribeEndpointTypesOutput, error) {
	return a.client.DescribeEndpointTypesWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) DescribeEndpoints(ctx context.Context, input *databasemigrationservice.DescribeEndpointsInput) (*databasemigrationservice.DescribeEndpointsOutput, error) {
	return a.client.DescribeEndpointsWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) DescribeEventCategories(ctx context.Context, input *databasemigrationservice.DescribeEventCategoriesInput) (*databasemigrationservice.DescribeEventCategoriesOutput, error) {
	return a.client.DescribeEventCategoriesWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) DescribeEventSubscriptions(ctx context.Context, input *databasemigrationservice.DescribeEventSubscriptionsInput) (*databasemigrationservice.DescribeEventSubscriptionsOutput, error) {
	return a.client.DescribeEventSubscriptionsWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) DescribeEvents(ctx context.Context, input *databasemigrationservice.DescribeEventsInput) (*databasemigrationservice.DescribeEventsOutput, error) {
	return a.client.DescribeEventsWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) DescribeOrderableReplicationInstances(ctx context.Context, input *databasemigrationservice.DescribeOrderableReplicationInstancesInput) (*databasemigrationservice.DescribeOrderableReplicationInstancesOutput, error) {
	return a.client.DescribeOrderableReplicationInstancesWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) DescribePendingMaintenanceActions(ctx context.Context, input *databasemigrationservice.DescribePendingMaintenanceActionsInput) (*databasemigrationservice.DescribePendingMaintenanceActionsOutput, error) {
	return a.client.DescribePendingMaintenanceActionsWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) DescribeRefreshSchemasStatus(ctx context.Context, input *databasemigrationservice.DescribeRefreshSchemasStatusInput) (*databasemigrationservice.DescribeRefreshSchemasStatusOutput, error) {
	return a.client.DescribeRefreshSchemasStatusWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) DescribeReplicationInstanceTaskLogs(ctx context.Context, input *databasemigrationservice.DescribeReplicationInstanceTaskLogsInput) (*databasemigrationservice.DescribeReplicationInstanceTaskLogsOutput, error) {
	return a.client.DescribeReplicationInstanceTaskLogsWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) DescribeReplicationInstances(ctx context.Context, input *databasemigrationservice.DescribeReplicationInstancesInput) (*databasemigrationservice.DescribeReplicationInstancesOutput, error) {
	return a.client.DescribeReplicationInstancesWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) DescribeReplicationSubnetGroups(ctx context.Context, input *databasemigrationservice.DescribeReplicationSubnetGroupsInput) (*databasemigrationservice.DescribeReplicationSubnetGroupsOutput, error) {
	return a.client.DescribeReplicationSubnetGroupsWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) DescribeReplicationTaskAssessmentResults(ctx context.Context, input *databasemigrationservice.DescribeReplicationTaskAssessmentResultsInput) (*databasemigrationservice.DescribeReplicationTaskAssessmentResultsOutput, error) {
	return a.client.DescribeReplicationTaskAssessmentResultsWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) DescribeReplicationTaskAssessmentRuns(ctx context.Context, input *databasemigrationservice.DescribeReplicationTaskAssessmentRunsInput) (*databasemigrationservice.DescribeReplicationTaskAssessmentRunsOutput, error) {
	return a.client.DescribeReplicationTaskAssessmentRunsWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) DescribeReplicationTaskIndividualAssessments(ctx context.Context, input *databasemigrationservice.DescribeReplicationTaskIndividualAssessmentsInput) (*databasemigrationservice.DescribeReplicationTaskIndividualAssessmentsOutput, error) {
	return a.client.DescribeReplicationTaskIndividualAssessmentsWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) DescribeReplicationTasks(ctx context.Context, input *databasemigrationservice.DescribeReplicationTasksInput) (*databasemigrationservice.DescribeReplicationTasksOutput, error) {
	return a.client.DescribeReplicationTasksWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) DescribeSchemas(ctx context.Context, input *databasemigrationservice.DescribeSchemasInput) (*databasemigrationservice.DescribeSchemasOutput, error) {
	return a.client.DescribeSchemasWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) DescribeTableStatistics(ctx context.Context, input *databasemigrationservice.DescribeTableStatisticsInput) (*databasemigrationservice.DescribeTableStatisticsOutput, error) {
	return a.client.DescribeTableStatisticsWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) ImportCertificate(ctx context.Context, input *databasemigrationservice.ImportCertificateInput) (*databasemigrationservice.ImportCertificateOutput, error) {
	return a.client.ImportCertificateWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) ListTagsForResource(ctx context.Context, input *databasemigrationservice.ListTagsForResourceInput) (*databasemigrationservice.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) ModifyEndpoint(ctx context.Context, input *databasemigrationservice.ModifyEndpointInput) (*databasemigrationservice.ModifyEndpointOutput, error) {
	return a.client.ModifyEndpointWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) ModifyEventSubscription(ctx context.Context, input *databasemigrationservice.ModifyEventSubscriptionInput) (*databasemigrationservice.ModifyEventSubscriptionOutput, error) {
	return a.client.ModifyEventSubscriptionWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) ModifyReplicationInstance(ctx context.Context, input *databasemigrationservice.ModifyReplicationInstanceInput) (*databasemigrationservice.ModifyReplicationInstanceOutput, error) {
	return a.client.ModifyReplicationInstanceWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) ModifyReplicationSubnetGroup(ctx context.Context, input *databasemigrationservice.ModifyReplicationSubnetGroupInput) (*databasemigrationservice.ModifyReplicationSubnetGroupOutput, error) {
	return a.client.ModifyReplicationSubnetGroupWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) ModifyReplicationTask(ctx context.Context, input *databasemigrationservice.ModifyReplicationTaskInput) (*databasemigrationservice.ModifyReplicationTaskOutput, error) {
	return a.client.ModifyReplicationTaskWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) RebootReplicationInstance(ctx context.Context, input *databasemigrationservice.RebootReplicationInstanceInput) (*databasemigrationservice.RebootReplicationInstanceOutput, error) {
	return a.client.RebootReplicationInstanceWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) RefreshSchemas(ctx context.Context, input *databasemigrationservice.RefreshSchemasInput) (*databasemigrationservice.RefreshSchemasOutput, error) {
	return a.client.RefreshSchemasWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) ReloadTables(ctx context.Context, input *databasemigrationservice.ReloadTablesInput) (*databasemigrationservice.ReloadTablesOutput, error) {
	return a.client.ReloadTablesWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) RemoveTagsFromResource(ctx context.Context, input *databasemigrationservice.RemoveTagsFromResourceInput) (*databasemigrationservice.RemoveTagsFromResourceOutput, error) {
	return a.client.RemoveTagsFromResourceWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) StartReplicationTask(ctx context.Context, input *databasemigrationservice.StartReplicationTaskInput) (*databasemigrationservice.StartReplicationTaskOutput, error) {
	return a.client.StartReplicationTaskWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) StartReplicationTaskAssessment(ctx context.Context, input *databasemigrationservice.StartReplicationTaskAssessmentInput) (*databasemigrationservice.StartReplicationTaskAssessmentOutput, error) {
	return a.client.StartReplicationTaskAssessmentWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) StartReplicationTaskAssessmentRun(ctx context.Context, input *databasemigrationservice.StartReplicationTaskAssessmentRunInput) (*databasemigrationservice.StartReplicationTaskAssessmentRunOutput, error) {
	return a.client.StartReplicationTaskAssessmentRunWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) StopReplicationTask(ctx context.Context, input *databasemigrationservice.StopReplicationTaskInput) (*databasemigrationservice.StopReplicationTaskOutput, error) {
	return a.client.StopReplicationTaskWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) TestConnection(ctx context.Context, input *databasemigrationservice.TestConnectionInput) (*databasemigrationservice.TestConnectionOutput, error) {
	return a.client.TestConnectionWithContext(ctx, input)
}

func (a *DatabaseMigrationServiceActivities) WaitUntilEndpointDeleted(ctx context.Context, input *databasemigrationservice.DescribeEndpointsInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilEndpointDeletedWithContext(ctx, input, options...)
	})
}

func (a *DatabaseMigrationServiceActivities) WaitUntilReplicationInstanceAvailable(ctx context.Context, input *databasemigrationservice.DescribeReplicationInstancesInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilReplicationInstanceAvailableWithContext(ctx, input, options...)
	})
}

func (a *DatabaseMigrationServiceActivities) WaitUntilReplicationInstanceDeleted(ctx context.Context, input *databasemigrationservice.DescribeReplicationInstancesInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilReplicationInstanceDeletedWithContext(ctx, input, options...)
	})
}

func (a *DatabaseMigrationServiceActivities) WaitUntilReplicationTaskDeleted(ctx context.Context, input *databasemigrationservice.DescribeReplicationTasksInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilReplicationTaskDeletedWithContext(ctx, input, options...)
	})
}

func (a *DatabaseMigrationServiceActivities) WaitUntilReplicationTaskReady(ctx context.Context, input *databasemigrationservice.DescribeReplicationTasksInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilReplicationTaskReadyWithContext(ctx, input, options...)
	})
}

func (a *DatabaseMigrationServiceActivities) WaitUntilReplicationTaskRunning(ctx context.Context, input *databasemigrationservice.DescribeReplicationTasksInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilReplicationTaskRunningWithContext(ctx, input, options...)
	})
}

func (a *DatabaseMigrationServiceActivities) WaitUntilReplicationTaskStopped(ctx context.Context, input *databasemigrationservice.DescribeReplicationTasksInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilReplicationTaskStoppedWithContext(ctx, input, options...)
	})
}

func (a *DatabaseMigrationServiceActivities) WaitUntilTestConnectionSucceeds(ctx context.Context, input *databasemigrationservice.DescribeConnectionsInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilTestConnectionSucceedsWithContext(ctx, input, options...)
	})
}
