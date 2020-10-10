// Generated by https://github.com/temporalio/temporal-aws-sdk/cmd/temporal-aws-sdk-gen/main.go
// from https://github.com/aws/aws-sdk-go.
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/mediapackagevod"
	"github.com/aws/aws-sdk-go/service/mediapackagevod/mediapackagevodiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type MediaPackageVodActivities struct {
	client mediapackagevodiface.MediaPackageVodAPI

	sessionFactory SessionFactory
}

func NewMediaPackageVodActivities(sess *session.Session, config ...*aws.Config) *MediaPackageVodActivities {
	client := mediapackagevod.New(sess, config...)
	return &MediaPackageVodActivities{client: client}
}

func NewMediaPackageVodActivitiesWithSessionFactory(sessionFactory SessionFactory) *MediaPackageVodActivities {
	return &MediaPackageVodActivities{sessionFactory: sessionFactory}
}

func (a *MediaPackageVodActivities) getClient(ctx context.Context) (mediapackagevodiface.MediaPackageVodAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return mediapackagevod.New(sess), nil
}

func (a *MediaPackageVodActivities) CreateAsset(ctx context.Context, input *mediapackagevod.CreateAssetInput) (*mediapackagevod.CreateAssetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateAssetWithContext(ctx, input)
}

func (a *MediaPackageVodActivities) CreatePackagingConfiguration(ctx context.Context, input *mediapackagevod.CreatePackagingConfigurationInput) (*mediapackagevod.CreatePackagingConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreatePackagingConfigurationWithContext(ctx, input)
}

func (a *MediaPackageVodActivities) CreatePackagingGroup(ctx context.Context, input *mediapackagevod.CreatePackagingGroupInput) (*mediapackagevod.CreatePackagingGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreatePackagingGroupWithContext(ctx, input)
}

func (a *MediaPackageVodActivities) DeleteAsset(ctx context.Context, input *mediapackagevod.DeleteAssetInput) (*mediapackagevod.DeleteAssetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteAssetWithContext(ctx, input)
}

func (a *MediaPackageVodActivities) DeletePackagingConfiguration(ctx context.Context, input *mediapackagevod.DeletePackagingConfigurationInput) (*mediapackagevod.DeletePackagingConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeletePackagingConfigurationWithContext(ctx, input)
}

func (a *MediaPackageVodActivities) DeletePackagingGroup(ctx context.Context, input *mediapackagevod.DeletePackagingGroupInput) (*mediapackagevod.DeletePackagingGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeletePackagingGroupWithContext(ctx, input)
}

func (a *MediaPackageVodActivities) DescribeAsset(ctx context.Context, input *mediapackagevod.DescribeAssetInput) (*mediapackagevod.DescribeAssetOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeAssetWithContext(ctx, input)
}

func (a *MediaPackageVodActivities) DescribePackagingConfiguration(ctx context.Context, input *mediapackagevod.DescribePackagingConfigurationInput) (*mediapackagevod.DescribePackagingConfigurationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribePackagingConfigurationWithContext(ctx, input)
}

func (a *MediaPackageVodActivities) DescribePackagingGroup(ctx context.Context, input *mediapackagevod.DescribePackagingGroupInput) (*mediapackagevod.DescribePackagingGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribePackagingGroupWithContext(ctx, input)
}

func (a *MediaPackageVodActivities) ListAssets(ctx context.Context, input *mediapackagevod.ListAssetsInput) (*mediapackagevod.ListAssetsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListAssetsWithContext(ctx, input)
}

func (a *MediaPackageVodActivities) ListPackagingConfigurations(ctx context.Context, input *mediapackagevod.ListPackagingConfigurationsInput) (*mediapackagevod.ListPackagingConfigurationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListPackagingConfigurationsWithContext(ctx, input)
}

func (a *MediaPackageVodActivities) ListPackagingGroups(ctx context.Context, input *mediapackagevod.ListPackagingGroupsInput) (*mediapackagevod.ListPackagingGroupsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListPackagingGroupsWithContext(ctx, input)
}

func (a *MediaPackageVodActivities) ListTagsForResource(ctx context.Context, input *mediapackagevod.ListTagsForResourceInput) (*mediapackagevod.ListTagsForResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListTagsForResourceWithContext(ctx, input)
}

func (a *MediaPackageVodActivities) TagResource(ctx context.Context, input *mediapackagevod.TagResourceInput) (*mediapackagevod.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.TagResourceWithContext(ctx, input)
}

func (a *MediaPackageVodActivities) UntagResource(ctx context.Context, input *mediapackagevod.UntagResourceInput) (*mediapackagevod.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UntagResourceWithContext(ctx, input)
}

func (a *MediaPackageVodActivities) UpdatePackagingGroup(ctx context.Context, input *mediapackagevod.UpdatePackagingGroupInput) (*mediapackagevod.UpdatePackagingGroupOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdatePackagingGroupWithContext(ctx, input)
}
