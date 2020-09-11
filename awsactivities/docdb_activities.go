package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/docdb"
	"github.com/aws/aws-sdk-go/service/docdb/docdbiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type DocDBActivities struct {
	client docdbiface.DocDBAPI
}

func NewDocDBActivities(session *session.Session, config ...*aws.Config) *DocDBActivities {
	client := docdb.New(session, config...)
	return &DocDBActivities{client: client}
}

func (a *DocDBActivities) AddTagsToResource(ctx context.Context, input *docdb.AddTagsToResourceInput) (*docdb.AddTagsToResourceOutput, error) {
	return a.client.AddTagsToResourceWithContext(ctx, input)
}

func (a *DocDBActivities) ApplyPendingMaintenanceAction(ctx context.Context, input *docdb.ApplyPendingMaintenanceActionInput) (*docdb.ApplyPendingMaintenanceActionOutput, error) {
	return a.client.ApplyPendingMaintenanceActionWithContext(ctx, input)
}

func (a *DocDBActivities) CopyDBClusterParameterGroup(ctx context.Context, input *docdb.CopyDBClusterParameterGroupInput) (*docdb.CopyDBClusterParameterGroupOutput, error) {
	return a.client.CopyDBClusterParameterGroupWithContext(ctx, input)
}

func (a *DocDBActivities) CopyDBClusterSnapshot(ctx context.Context, input *docdb.CopyDBClusterSnapshotInput) (*docdb.CopyDBClusterSnapshotOutput, error) {
	return a.client.CopyDBClusterSnapshotWithContext(ctx, input)
}

func (a *DocDBActivities) CreateDBCluster(ctx context.Context, input *docdb.CreateDBClusterInput) (*docdb.CreateDBClusterOutput, error) {
	return a.client.CreateDBClusterWithContext(ctx, input)
}

func (a *DocDBActivities) CreateDBClusterParameterGroup(ctx context.Context, input *docdb.CreateDBClusterParameterGroupInput) (*docdb.CreateDBClusterParameterGroupOutput, error) {
	return a.client.CreateDBClusterParameterGroupWithContext(ctx, input)
}

func (a *DocDBActivities) CreateDBClusterSnapshot(ctx context.Context, input *docdb.CreateDBClusterSnapshotInput) (*docdb.CreateDBClusterSnapshotOutput, error) {
	return a.client.CreateDBClusterSnapshotWithContext(ctx, input)
}

func (a *DocDBActivities) CreateDBInstance(ctx context.Context, input *docdb.CreateDBInstanceInput) (*docdb.CreateDBInstanceOutput, error) {
	return a.client.CreateDBInstanceWithContext(ctx, input)
}

func (a *DocDBActivities) CreateDBSubnetGroup(ctx context.Context, input *docdb.CreateDBSubnetGroupInput) (*docdb.CreateDBSubnetGroupOutput, error) {
	return a.client.CreateDBSubnetGroupWithContext(ctx, input)
}

func (a *DocDBActivities) DeleteDBCluster(ctx context.Context, input *docdb.DeleteDBClusterInput) (*docdb.DeleteDBClusterOutput, error) {
	return a.client.DeleteDBClusterWithContext(ctx, input)
}

func (a *DocDBActivities) DeleteDBClusterParameterGroup(ctx context.Context, input *docdb.DeleteDBClusterParameterGroupInput) (*docdb.DeleteDBClusterParameterGroupOutput, error) {
	return a.client.DeleteDBClusterParameterGroupWithContext(ctx, input)
}

func (a *DocDBActivities) DeleteDBClusterSnapshot(ctx context.Context, input *docdb.DeleteDBClusterSnapshotInput) (*docdb.DeleteDBClusterSnapshotOutput, error) {
	return a.client.DeleteDBClusterSnapshotWithContext(ctx, input)
}

func (a *DocDBActivities) DeleteDBInstance(ctx context.Context, input *docdb.DeleteDBInstanceInput) (*docdb.DeleteDBInstanceOutput, error) {
	return a.client.DeleteDBInstanceWithContext(ctx, input)
}

func (a *DocDBActivities) DeleteDBSubnetGroup(ctx context.Context, input *docdb.DeleteDBSubnetGroupInput) (*docdb.DeleteDBSubnetGroupOutput, error) {
	return a.client.DeleteDBSubnetGroupWithContext(ctx, input)
}

func (a *DocDBActivities) DescribeCertificates(ctx context.Context, input *docdb.DescribeCertificatesInput) (*docdb.DescribeCertificatesOutput, error) {
	return a.client.DescribeCertificatesWithContext(ctx, input)
}

func (a *DocDBActivities) DescribeDBClusterParameterGroups(ctx context.Context, input *docdb.DescribeDBClusterParameterGroupsInput) (*docdb.DescribeDBClusterParameterGroupsOutput, error) {
	return a.client.DescribeDBClusterParameterGroupsWithContext(ctx, input)
}

func (a *DocDBActivities) DescribeDBClusterParameters(ctx context.Context, input *docdb.DescribeDBClusterParametersInput) (*docdb.DescribeDBClusterParametersOutput, error) {
	return a.client.DescribeDBClusterParametersWithContext(ctx, input)
}

func (a *DocDBActivities) DescribeDBClusterSnapshotAttributes(ctx context.Context, input *docdb.DescribeDBClusterSnapshotAttributesInput) (*docdb.DescribeDBClusterSnapshotAttributesOutput, error) {
	return a.client.DescribeDBClusterSnapshotAttributesWithContext(ctx, input)
}

func (a *DocDBActivities) DescribeDBClusterSnapshots(ctx context.Context, input *docdb.DescribeDBClusterSnapshotsInput) (*docdb.DescribeDBClusterSnapshotsOutput, error) {
	return a.client.DescribeDBClusterSnapshotsWithContext(ctx, input)
}

func (a *DocDBActivities) DescribeDBClusters(ctx context.Context, input *docdb.DescribeDBClustersInput) (*docdb.DescribeDBClustersOutput, error) {
	return a.client.DescribeDBClustersWithContext(ctx, input)
}

func (a *DocDBActivities) DescribeDBEngineVersions(ctx context.Context, input *docdb.DescribeDBEngineVersionsInput) (*docdb.DescribeDBEngineVersionsOutput, error) {
	return a.client.DescribeDBEngineVersionsWithContext(ctx, input)
}

func (a *DocDBActivities) DescribeDBInstances(ctx context.Context, input *docdb.DescribeDBInstancesInput) (*docdb.DescribeDBInstancesOutput, error) {
	return a.client.DescribeDBInstancesWithContext(ctx, input)
}

func (a *DocDBActivities) DescribeDBSubnetGroups(ctx context.Context, input *docdb.DescribeDBSubnetGroupsInput) (*docdb.DescribeDBSubnetGroupsOutput, error) {
	return a.client.DescribeDBSubnetGroupsWithContext(ctx, input)
}

func (a *DocDBActivities) DescribeEngineDefaultClusterParameters(ctx context.Context, input *docdb.DescribeEngineDefaultClusterParametersInput) (*docdb.DescribeEngineDefaultClusterParametersOutput, error) {
	return a.client.DescribeEngineDefaultClusterParametersWithContext(ctx, input)
}

func (a *DocDBActivities) DescribeEventCategories(ctx context.Context, input *docdb.DescribeEventCategoriesInput) (*docdb.DescribeEventCategoriesOutput, error) {
	return a.client.DescribeEventCategoriesWithContext(ctx, input)
}

func (a *DocDBActivities) DescribeEvents(ctx context.Context, input *docdb.DescribeEventsInput) (*docdb.DescribeEventsOutput, error) {
	return a.client.DescribeEventsWithContext(ctx, input)
}

func (a *DocDBActivities) DescribeOrderableDBInstanceOptions(ctx context.Context, input *docdb.DescribeOrderableDBInstanceOptionsInput) (*docdb.DescribeOrderableDBInstanceOptionsOutput, error) {
	return a.client.DescribeOrderableDBInstanceOptionsWithContext(ctx, input)
}

func (a *DocDBActivities) DescribePendingMaintenanceActions(ctx context.Context, input *docdb.DescribePendingMaintenanceActionsInput) (*docdb.DescribePendingMaintenanceActionsOutput, error) {
	return a.client.DescribePendingMaintenanceActionsWithContext(ctx, input)
}

func (a *DocDBActivities) FailoverDBCluster(ctx context.Context, input *docdb.FailoverDBClusterInput) (*docdb.FailoverDBClusterOutput, error) {
	return a.client.FailoverDBClusterWithContext(ctx, input)
}

func (a *DocDBActivities) ListTagsForResource(ctx context.Context, input *docdb.ListTagsForResourceInput) (*docdb.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *DocDBActivities) ModifyDBCluster(ctx context.Context, input *docdb.ModifyDBClusterInput) (*docdb.ModifyDBClusterOutput, error) {
	return a.client.ModifyDBClusterWithContext(ctx, input)
}

func (a *DocDBActivities) ModifyDBClusterParameterGroup(ctx context.Context, input *docdb.ModifyDBClusterParameterGroupInput) (*docdb.ModifyDBClusterParameterGroupOutput, error) {
	return a.client.ModifyDBClusterParameterGroupWithContext(ctx, input)
}

func (a *DocDBActivities) ModifyDBClusterSnapshotAttribute(ctx context.Context, input *docdb.ModifyDBClusterSnapshotAttributeInput) (*docdb.ModifyDBClusterSnapshotAttributeOutput, error) {
	return a.client.ModifyDBClusterSnapshotAttributeWithContext(ctx, input)
}

func (a *DocDBActivities) ModifyDBInstance(ctx context.Context, input *docdb.ModifyDBInstanceInput) (*docdb.ModifyDBInstanceOutput, error) {
	return a.client.ModifyDBInstanceWithContext(ctx, input)
}

func (a *DocDBActivities) ModifyDBSubnetGroup(ctx context.Context, input *docdb.ModifyDBSubnetGroupInput) (*docdb.ModifyDBSubnetGroupOutput, error) {
	return a.client.ModifyDBSubnetGroupWithContext(ctx, input)
}

func (a *DocDBActivities) RebootDBInstance(ctx context.Context, input *docdb.RebootDBInstanceInput) (*docdb.RebootDBInstanceOutput, error) {
	return a.client.RebootDBInstanceWithContext(ctx, input)
}

func (a *DocDBActivities) RemoveTagsFromResource(ctx context.Context, input *docdb.RemoveTagsFromResourceInput) (*docdb.RemoveTagsFromResourceOutput, error) {
	return a.client.RemoveTagsFromResourceWithContext(ctx, input)
}

func (a *DocDBActivities) ResetDBClusterParameterGroup(ctx context.Context, input *docdb.ResetDBClusterParameterGroupInput) (*docdb.ResetDBClusterParameterGroupOutput, error) {
	return a.client.ResetDBClusterParameterGroupWithContext(ctx, input)
}

func (a *DocDBActivities) RestoreDBClusterFromSnapshot(ctx context.Context, input *docdb.RestoreDBClusterFromSnapshotInput) (*docdb.RestoreDBClusterFromSnapshotOutput, error) {
	return a.client.RestoreDBClusterFromSnapshotWithContext(ctx, input)
}

func (a *DocDBActivities) RestoreDBClusterToPointInTime(ctx context.Context, input *docdb.RestoreDBClusterToPointInTimeInput) (*docdb.RestoreDBClusterToPointInTimeOutput, error) {
	return a.client.RestoreDBClusterToPointInTimeWithContext(ctx, input)
}

func (a *DocDBActivities) StartDBCluster(ctx context.Context, input *docdb.StartDBClusterInput) (*docdb.StartDBClusterOutput, error) {
	return a.client.StartDBClusterWithContext(ctx, input)
}

func (a *DocDBActivities) StopDBCluster(ctx context.Context, input *docdb.StopDBClusterInput) (*docdb.StopDBClusterOutput, error) {
	return a.client.StopDBClusterWithContext(ctx, input)
}

func (a *DocDBActivities) WaitUntilDBInstanceAvailable(ctx context.Context, input *docdb.DescribeDBInstancesInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilDBInstanceAvailableWithContext(ctx, input, options...)
	})
}

func (a *DocDBActivities) WaitUntilDBInstanceDeleted(ctx context.Context, input *docdb.DescribeDBInstancesInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilDBInstanceDeletedWithContext(ctx, input, options...)
	})
}
