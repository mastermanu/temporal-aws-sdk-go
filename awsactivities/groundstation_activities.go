package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/groundstation"
	"github.com/aws/aws-sdk-go/service/groundstation/groundstationiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type GroundStationActivities struct {
	client groundstationiface.GroundStationAPI
}

func NewGroundStationActivities(session *session.Session, config ...*aws.Config) *GroundStationActivities {
	client := groundstation.New(session, config...)
	return &GroundStationActivities{client: client}
}

func (a *GroundStationActivities) CancelContact(ctx context.Context, input *groundstation.CancelContactInput) (*groundstation.CancelContactOutput, error) {
	return a.client.CancelContactWithContext(ctx, input)
}

func (a *GroundStationActivities) CreateConfig(ctx context.Context, input *groundstation.CreateConfigInput) (*groundstation.CreateConfigOutput, error) {
	return a.client.CreateConfigWithContext(ctx, input)
}

func (a *GroundStationActivities) CreateDataflowEndpointGroup(ctx context.Context, input *groundstation.CreateDataflowEndpointGroupInput) (*groundstation.CreateDataflowEndpointGroupOutput, error) {
	return a.client.CreateDataflowEndpointGroupWithContext(ctx, input)
}

func (a *GroundStationActivities) CreateMissionProfile(ctx context.Context, input *groundstation.CreateMissionProfileInput) (*groundstation.CreateMissionProfileOutput, error) {
	return a.client.CreateMissionProfileWithContext(ctx, input)
}

func (a *GroundStationActivities) DeleteConfig(ctx context.Context, input *groundstation.DeleteConfigInput) (*groundstation.DeleteConfigOutput, error) {
	return a.client.DeleteConfigWithContext(ctx, input)
}

func (a *GroundStationActivities) DeleteDataflowEndpointGroup(ctx context.Context, input *groundstation.DeleteDataflowEndpointGroupInput) (*groundstation.DeleteDataflowEndpointGroupOutput, error) {
	return a.client.DeleteDataflowEndpointGroupWithContext(ctx, input)
}

func (a *GroundStationActivities) DeleteMissionProfile(ctx context.Context, input *groundstation.DeleteMissionProfileInput) (*groundstation.DeleteMissionProfileOutput, error) {
	return a.client.DeleteMissionProfileWithContext(ctx, input)
}

func (a *GroundStationActivities) DescribeContact(ctx context.Context, input *groundstation.DescribeContactInput) (*groundstation.DescribeContactOutput, error) {
	return a.client.DescribeContactWithContext(ctx, input)
}

func (a *GroundStationActivities) GetConfig(ctx context.Context, input *groundstation.GetConfigInput) (*groundstation.GetConfigOutput, error) {
	return a.client.GetConfigWithContext(ctx, input)
}

func (a *GroundStationActivities) GetDataflowEndpointGroup(ctx context.Context, input *groundstation.GetDataflowEndpointGroupInput) (*groundstation.GetDataflowEndpointGroupOutput, error) {
	return a.client.GetDataflowEndpointGroupWithContext(ctx, input)
}

func (a *GroundStationActivities) GetMinuteUsage(ctx context.Context, input *groundstation.GetMinuteUsageInput) (*groundstation.GetMinuteUsageOutput, error) {
	return a.client.GetMinuteUsageWithContext(ctx, input)
}

func (a *GroundStationActivities) GetMissionProfile(ctx context.Context, input *groundstation.GetMissionProfileInput) (*groundstation.GetMissionProfileOutput, error) {
	return a.client.GetMissionProfileWithContext(ctx, input)
}

func (a *GroundStationActivities) GetSatellite(ctx context.Context, input *groundstation.GetSatelliteInput) (*groundstation.GetSatelliteOutput, error) {
	return a.client.GetSatelliteWithContext(ctx, input)
}

func (a *GroundStationActivities) ListConfigs(ctx context.Context, input *groundstation.ListConfigsInput) (*groundstation.ListConfigsOutput, error) {
	return a.client.ListConfigsWithContext(ctx, input)
}

func (a *GroundStationActivities) ListContacts(ctx context.Context, input *groundstation.ListContactsInput) (*groundstation.ListContactsOutput, error) {
	return a.client.ListContactsWithContext(ctx, input)
}

func (a *GroundStationActivities) ListDataflowEndpointGroups(ctx context.Context, input *groundstation.ListDataflowEndpointGroupsInput) (*groundstation.ListDataflowEndpointGroupsOutput, error) {
	return a.client.ListDataflowEndpointGroupsWithContext(ctx, input)
}

func (a *GroundStationActivities) ListGroundStations(ctx context.Context, input *groundstation.ListGroundStationsInput) (*groundstation.ListGroundStationsOutput, error) {
	return a.client.ListGroundStationsWithContext(ctx, input)
}

func (a *GroundStationActivities) ListMissionProfiles(ctx context.Context, input *groundstation.ListMissionProfilesInput) (*groundstation.ListMissionProfilesOutput, error) {
	return a.client.ListMissionProfilesWithContext(ctx, input)
}

func (a *GroundStationActivities) ListSatellites(ctx context.Context, input *groundstation.ListSatellitesInput) (*groundstation.ListSatellitesOutput, error) {
	return a.client.ListSatellitesWithContext(ctx, input)
}

func (a *GroundStationActivities) ListTagsForResource(ctx context.Context, input *groundstation.ListTagsForResourceInput) (*groundstation.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *GroundStationActivities) ReserveContact(ctx context.Context, input *groundstation.ReserveContactInput) (*groundstation.ReserveContactOutput, error) {
	return a.client.ReserveContactWithContext(ctx, input)
}

func (a *GroundStationActivities) TagResource(ctx context.Context, input *groundstation.TagResourceInput) (*groundstation.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *GroundStationActivities) UntagResource(ctx context.Context, input *groundstation.UntagResourceInput) (*groundstation.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}

func (a *GroundStationActivities) UpdateConfig(ctx context.Context, input *groundstation.UpdateConfigInput) (*groundstation.UpdateConfigOutput, error) {
	return a.client.UpdateConfigWithContext(ctx, input)
}

func (a *GroundStationActivities) UpdateMissionProfile(ctx context.Context, input *groundstation.UpdateMissionProfileInput) (*groundstation.UpdateMissionProfileOutput, error) {
	return a.client.UpdateMissionProfileWithContext(ctx, input)
}
