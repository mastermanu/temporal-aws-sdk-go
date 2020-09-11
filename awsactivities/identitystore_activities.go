package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/identitystore"
	"github.com/aws/aws-sdk-go/service/identitystore/identitystoreiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type IdentityStoreActivities struct {
	client identitystoreiface.IdentityStoreAPI
}

func NewIdentityStoreActivities(session *session.Session, config ...*aws.Config) *IdentityStoreActivities {
	client := identitystore.New(session, config...)
	return &IdentityStoreActivities{client: client}
}

func (a *IdentityStoreActivities) DescribeGroup(ctx context.Context, input *identitystore.DescribeGroupInput) (*identitystore.DescribeGroupOutput, error) {
	return a.client.DescribeGroupWithContext(ctx, input)
}

func (a *IdentityStoreActivities) DescribeUser(ctx context.Context, input *identitystore.DescribeUserInput) (*identitystore.DescribeUserOutput, error) {
	return a.client.DescribeUserWithContext(ctx, input)
}

func (a *IdentityStoreActivities) ListGroups(ctx context.Context, input *identitystore.ListGroupsInput) (*identitystore.ListGroupsOutput, error) {
	return a.client.ListGroupsWithContext(ctx, input)
}

func (a *IdentityStoreActivities) ListUsers(ctx context.Context, input *identitystore.ListUsersInput) (*identitystore.ListUsersOutput, error) {
	return a.client.ListUsersWithContext(ctx, input)
}
