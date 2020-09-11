package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/fsx"
	"github.com/aws/aws-sdk-go/service/fsx/fsxiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type FSxActivities struct {
	client fsxiface.FSxAPI
}

func NewFSxActivities(session *session.Session, config ...*aws.Config) *FSxActivities {
	client := fsx.New(session, config...)
	return &FSxActivities{client: client}
}

func (a *FSxActivities) CancelDataRepositoryTask(ctx context.Context, input *fsx.CancelDataRepositoryTaskInput) (*fsx.CancelDataRepositoryTaskOutput, error) {
	return a.client.CancelDataRepositoryTaskWithContext(ctx, input)
}

func (a *FSxActivities) CreateBackup(ctx context.Context, input *fsx.CreateBackupInput) (*fsx.CreateBackupOutput, error) {
	return a.client.CreateBackupWithContext(ctx, input)
}

func (a *FSxActivities) CreateDataRepositoryTask(ctx context.Context, input *fsx.CreateDataRepositoryTaskInput) (*fsx.CreateDataRepositoryTaskOutput, error) {
	return a.client.CreateDataRepositoryTaskWithContext(ctx, input)
}

func (a *FSxActivities) CreateFileSystem(ctx context.Context, input *fsx.CreateFileSystemInput) (*fsx.CreateFileSystemOutput, error) {
	return a.client.CreateFileSystemWithContext(ctx, input)
}

func (a *FSxActivities) CreateFileSystemFromBackup(ctx context.Context, input *fsx.CreateFileSystemFromBackupInput) (*fsx.CreateFileSystemFromBackupOutput, error) {
	return a.client.CreateFileSystemFromBackupWithContext(ctx, input)
}

func (a *FSxActivities) DeleteBackup(ctx context.Context, input *fsx.DeleteBackupInput) (*fsx.DeleteBackupOutput, error) {
	return a.client.DeleteBackupWithContext(ctx, input)
}

func (a *FSxActivities) DeleteFileSystem(ctx context.Context, input *fsx.DeleteFileSystemInput) (*fsx.DeleteFileSystemOutput, error) {
	return a.client.DeleteFileSystemWithContext(ctx, input)
}

func (a *FSxActivities) DescribeBackups(ctx context.Context, input *fsx.DescribeBackupsInput) (*fsx.DescribeBackupsOutput, error) {
	return a.client.DescribeBackupsWithContext(ctx, input)
}

func (a *FSxActivities) DescribeDataRepositoryTasks(ctx context.Context, input *fsx.DescribeDataRepositoryTasksInput) (*fsx.DescribeDataRepositoryTasksOutput, error) {
	return a.client.DescribeDataRepositoryTasksWithContext(ctx, input)
}

func (a *FSxActivities) DescribeFileSystems(ctx context.Context, input *fsx.DescribeFileSystemsInput) (*fsx.DescribeFileSystemsOutput, error) {
	return a.client.DescribeFileSystemsWithContext(ctx, input)
}

func (a *FSxActivities) ListTagsForResource(ctx context.Context, input *fsx.ListTagsForResourceInput) (*fsx.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *FSxActivities) TagResource(ctx context.Context, input *fsx.TagResourceInput) (*fsx.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *FSxActivities) UntagResource(ctx context.Context, input *fsx.UntagResourceInput) (*fsx.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}

func (a *FSxActivities) UpdateFileSystem(ctx context.Context, input *fsx.UpdateFileSystemInput) (*fsx.UpdateFileSystemOutput, error) {
	return a.client.UpdateFileSystemWithContext(ctx, input)
}
