package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/groundstation"
	"go.temporal.io/sdk/workflow"
)

type GroundStationClient interface {
    CancelContact(ctx workflow.Context, input *groundstation.CancelContactInput) (*groundstation.CancelContactOutput, error)
    CancelContactAsync(ctx workflow.Context, input *groundstation.CancelContactInput) *GroundstationCancelContactResult

    CreateConfig(ctx workflow.Context, input *groundstation.CreateConfigInput) (*groundstation.CreateConfigOutput, error)
    CreateConfigAsync(ctx workflow.Context, input *groundstation.CreateConfigInput) *GroundstationCreateConfigResult

    CreateDataflowEndpointGroup(ctx workflow.Context, input *groundstation.CreateDataflowEndpointGroupInput) (*groundstation.CreateDataflowEndpointGroupOutput, error)
    CreateDataflowEndpointGroupAsync(ctx workflow.Context, input *groundstation.CreateDataflowEndpointGroupInput) *GroundstationCreateDataflowEndpointGroupResult

    CreateMissionProfile(ctx workflow.Context, input *groundstation.CreateMissionProfileInput) (*groundstation.CreateMissionProfileOutput, error)
    CreateMissionProfileAsync(ctx workflow.Context, input *groundstation.CreateMissionProfileInput) *GroundstationCreateMissionProfileResult

    DeleteConfig(ctx workflow.Context, input *groundstation.DeleteConfigInput) (*groundstation.DeleteConfigOutput, error)
    DeleteConfigAsync(ctx workflow.Context, input *groundstation.DeleteConfigInput) *GroundstationDeleteConfigResult

    DeleteDataflowEndpointGroup(ctx workflow.Context, input *groundstation.DeleteDataflowEndpointGroupInput) (*groundstation.DeleteDataflowEndpointGroupOutput, error)
    DeleteDataflowEndpointGroupAsync(ctx workflow.Context, input *groundstation.DeleteDataflowEndpointGroupInput) *GroundstationDeleteDataflowEndpointGroupResult

    DeleteMissionProfile(ctx workflow.Context, input *groundstation.DeleteMissionProfileInput) (*groundstation.DeleteMissionProfileOutput, error)
    DeleteMissionProfileAsync(ctx workflow.Context, input *groundstation.DeleteMissionProfileInput) *GroundstationDeleteMissionProfileResult

    DescribeContact(ctx workflow.Context, input *groundstation.DescribeContactInput) (*groundstation.DescribeContactOutput, error)
    DescribeContactAsync(ctx workflow.Context, input *groundstation.DescribeContactInput) *GroundstationDescribeContactResult

    GetConfig(ctx workflow.Context, input *groundstation.GetConfigInput) (*groundstation.GetConfigOutput, error)
    GetConfigAsync(ctx workflow.Context, input *groundstation.GetConfigInput) *GroundstationGetConfigResult

    GetDataflowEndpointGroup(ctx workflow.Context, input *groundstation.GetDataflowEndpointGroupInput) (*groundstation.GetDataflowEndpointGroupOutput, error)
    GetDataflowEndpointGroupAsync(ctx workflow.Context, input *groundstation.GetDataflowEndpointGroupInput) *GroundstationGetDataflowEndpointGroupResult

    GetMinuteUsage(ctx workflow.Context, input *groundstation.GetMinuteUsageInput) (*groundstation.GetMinuteUsageOutput, error)
    GetMinuteUsageAsync(ctx workflow.Context, input *groundstation.GetMinuteUsageInput) *GroundstationGetMinuteUsageResult

    GetMissionProfile(ctx workflow.Context, input *groundstation.GetMissionProfileInput) (*groundstation.GetMissionProfileOutput, error)
    GetMissionProfileAsync(ctx workflow.Context, input *groundstation.GetMissionProfileInput) *GroundstationGetMissionProfileResult

    GetSatellite(ctx workflow.Context, input *groundstation.GetSatelliteInput) (*groundstation.GetSatelliteOutput, error)
    GetSatelliteAsync(ctx workflow.Context, input *groundstation.GetSatelliteInput) *GroundstationGetSatelliteResult

    ListConfigs(ctx workflow.Context, input *groundstation.ListConfigsInput) (*groundstation.ListConfigsOutput, error)
    ListConfigsAsync(ctx workflow.Context, input *groundstation.ListConfigsInput) *GroundstationListConfigsResult

    ListContacts(ctx workflow.Context, input *groundstation.ListContactsInput) (*groundstation.ListContactsOutput, error)
    ListContactsAsync(ctx workflow.Context, input *groundstation.ListContactsInput) *GroundstationListContactsResult

    ListDataflowEndpointGroups(ctx workflow.Context, input *groundstation.ListDataflowEndpointGroupsInput) (*groundstation.ListDataflowEndpointGroupsOutput, error)
    ListDataflowEndpointGroupsAsync(ctx workflow.Context, input *groundstation.ListDataflowEndpointGroupsInput) *GroundstationListDataflowEndpointGroupsResult

    ListGroundStations(ctx workflow.Context, input *groundstation.ListGroundStationsInput) (*groundstation.ListGroundStationsOutput, error)
    ListGroundStationsAsync(ctx workflow.Context, input *groundstation.ListGroundStationsInput) *GroundstationListGroundStationsResult

    ListMissionProfiles(ctx workflow.Context, input *groundstation.ListMissionProfilesInput) (*groundstation.ListMissionProfilesOutput, error)
    ListMissionProfilesAsync(ctx workflow.Context, input *groundstation.ListMissionProfilesInput) *GroundstationListMissionProfilesResult

    ListSatellites(ctx workflow.Context, input *groundstation.ListSatellitesInput) (*groundstation.ListSatellitesOutput, error)
    ListSatellitesAsync(ctx workflow.Context, input *groundstation.ListSatellitesInput) *GroundstationListSatellitesResult

    ListTagsForResource(ctx workflow.Context, input *groundstation.ListTagsForResourceInput) (*groundstation.ListTagsForResourceOutput, error)
    ListTagsForResourceAsync(ctx workflow.Context, input *groundstation.ListTagsForResourceInput) *GroundstationListTagsForResourceResult

    ReserveContact(ctx workflow.Context, input *groundstation.ReserveContactInput) (*groundstation.ReserveContactOutput, error)
    ReserveContactAsync(ctx workflow.Context, input *groundstation.ReserveContactInput) *GroundstationReserveContactResult

    TagResource(ctx workflow.Context, input *groundstation.TagResourceInput) (*groundstation.TagResourceOutput, error)
    TagResourceAsync(ctx workflow.Context, input *groundstation.TagResourceInput) *GroundstationTagResourceResult

    UntagResource(ctx workflow.Context, input *groundstation.UntagResourceInput) (*groundstation.UntagResourceOutput, error)
    UntagResourceAsync(ctx workflow.Context, input *groundstation.UntagResourceInput) *GroundstationUntagResourceResult

    UpdateConfig(ctx workflow.Context, input *groundstation.UpdateConfigInput) (*groundstation.UpdateConfigOutput, error)
    UpdateConfigAsync(ctx workflow.Context, input *groundstation.UpdateConfigInput) *GroundstationUpdateConfigResult

    UpdateMissionProfile(ctx workflow.Context, input *groundstation.UpdateMissionProfileInput) (*groundstation.UpdateMissionProfileOutput, error)
    UpdateMissionProfileAsync(ctx workflow.Context, input *groundstation.UpdateMissionProfileInput) *GroundstationUpdateMissionProfileResult
}
type GroundstationCancelContactResult struct {
	Result workflow.Future
}

func (r *GroundstationCancelContactResult) Get(ctx workflow.Context) (*groundstation.CancelContactOutput, error) {
    var output groundstation.CancelContactOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GroundstationCreateConfigResult struct {
	Result workflow.Future
}

func (r *GroundstationCreateConfigResult) Get(ctx workflow.Context) (*groundstation.CreateConfigOutput, error) {
    var output groundstation.CreateConfigOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GroundstationCreateDataflowEndpointGroupResult struct {
	Result workflow.Future
}

func (r *GroundstationCreateDataflowEndpointGroupResult) Get(ctx workflow.Context) (*groundstation.CreateDataflowEndpointGroupOutput, error) {
    var output groundstation.CreateDataflowEndpointGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GroundstationCreateMissionProfileResult struct {
	Result workflow.Future
}

func (r *GroundstationCreateMissionProfileResult) Get(ctx workflow.Context) (*groundstation.CreateMissionProfileOutput, error) {
    var output groundstation.CreateMissionProfileOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GroundstationDeleteConfigResult struct {
	Result workflow.Future
}

func (r *GroundstationDeleteConfigResult) Get(ctx workflow.Context) (*groundstation.DeleteConfigOutput, error) {
    var output groundstation.DeleteConfigOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GroundstationDeleteDataflowEndpointGroupResult struct {
	Result workflow.Future
}

func (r *GroundstationDeleteDataflowEndpointGroupResult) Get(ctx workflow.Context) (*groundstation.DeleteDataflowEndpointGroupOutput, error) {
    var output groundstation.DeleteDataflowEndpointGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GroundstationDeleteMissionProfileResult struct {
	Result workflow.Future
}

func (r *GroundstationDeleteMissionProfileResult) Get(ctx workflow.Context) (*groundstation.DeleteMissionProfileOutput, error) {
    var output groundstation.DeleteMissionProfileOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GroundstationDescribeContactResult struct {
	Result workflow.Future
}

func (r *GroundstationDescribeContactResult) Get(ctx workflow.Context) (*groundstation.DescribeContactOutput, error) {
    var output groundstation.DescribeContactOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GroundstationGetConfigResult struct {
	Result workflow.Future
}

func (r *GroundstationGetConfigResult) Get(ctx workflow.Context) (*groundstation.GetConfigOutput, error) {
    var output groundstation.GetConfigOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GroundstationGetDataflowEndpointGroupResult struct {
	Result workflow.Future
}

func (r *GroundstationGetDataflowEndpointGroupResult) Get(ctx workflow.Context) (*groundstation.GetDataflowEndpointGroupOutput, error) {
    var output groundstation.GetDataflowEndpointGroupOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GroundstationGetMinuteUsageResult struct {
	Result workflow.Future
}

func (r *GroundstationGetMinuteUsageResult) Get(ctx workflow.Context) (*groundstation.GetMinuteUsageOutput, error) {
    var output groundstation.GetMinuteUsageOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GroundstationGetMissionProfileResult struct {
	Result workflow.Future
}

func (r *GroundstationGetMissionProfileResult) Get(ctx workflow.Context) (*groundstation.GetMissionProfileOutput, error) {
    var output groundstation.GetMissionProfileOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GroundstationGetSatelliteResult struct {
	Result workflow.Future
}

func (r *GroundstationGetSatelliteResult) Get(ctx workflow.Context) (*groundstation.GetSatelliteOutput, error) {
    var output groundstation.GetSatelliteOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GroundstationListConfigsResult struct {
	Result workflow.Future
}

func (r *GroundstationListConfigsResult) Get(ctx workflow.Context) (*groundstation.ListConfigsOutput, error) {
    var output groundstation.ListConfigsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GroundstationListContactsResult struct {
	Result workflow.Future
}

func (r *GroundstationListContactsResult) Get(ctx workflow.Context) (*groundstation.ListContactsOutput, error) {
    var output groundstation.ListContactsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GroundstationListDataflowEndpointGroupsResult struct {
	Result workflow.Future
}

func (r *GroundstationListDataflowEndpointGroupsResult) Get(ctx workflow.Context) (*groundstation.ListDataflowEndpointGroupsOutput, error) {
    var output groundstation.ListDataflowEndpointGroupsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GroundstationListGroundStationsResult struct {
	Result workflow.Future
}

func (r *GroundstationListGroundStationsResult) Get(ctx workflow.Context) (*groundstation.ListGroundStationsOutput, error) {
    var output groundstation.ListGroundStationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GroundstationListMissionProfilesResult struct {
	Result workflow.Future
}

func (r *GroundstationListMissionProfilesResult) Get(ctx workflow.Context) (*groundstation.ListMissionProfilesOutput, error) {
    var output groundstation.ListMissionProfilesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GroundstationListSatellitesResult struct {
	Result workflow.Future
}

func (r *GroundstationListSatellitesResult) Get(ctx workflow.Context) (*groundstation.ListSatellitesOutput, error) {
    var output groundstation.ListSatellitesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GroundstationListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *GroundstationListTagsForResourceResult) Get(ctx workflow.Context) (*groundstation.ListTagsForResourceOutput, error) {
    var output groundstation.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GroundstationReserveContactResult struct {
	Result workflow.Future
}

func (r *GroundstationReserveContactResult) Get(ctx workflow.Context) (*groundstation.ReserveContactOutput, error) {
    var output groundstation.ReserveContactOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GroundstationTagResourceResult struct {
	Result workflow.Future
}

func (r *GroundstationTagResourceResult) Get(ctx workflow.Context) (*groundstation.TagResourceOutput, error) {
    var output groundstation.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GroundstationUntagResourceResult struct {
	Result workflow.Future
}

func (r *GroundstationUntagResourceResult) Get(ctx workflow.Context) (*groundstation.UntagResourceOutput, error) {
    var output groundstation.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GroundstationUpdateConfigResult struct {
	Result workflow.Future
}

func (r *GroundstationUpdateConfigResult) Get(ctx workflow.Context) (*groundstation.UpdateConfigOutput, error) {
    var output groundstation.UpdateConfigOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type GroundstationUpdateMissionProfileResult struct {
	Result workflow.Future
}

func (r *GroundstationUpdateMissionProfileResult) Get(ctx workflow.Context) (*groundstation.UpdateMissionProfileOutput, error) {
    var output groundstation.UpdateMissionProfileOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type GroundStationStub struct {
    activities GroundStationClient
}

func NewGroundStationStub() GroundStationClient {
    return &GroundStationStub{}
}

func (a *GroundStationStub) CancelContact(ctx workflow.Context, input *groundstation.CancelContactInput) (*groundstation.CancelContactOutput, error) {
    var output groundstation.CancelContactOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CancelContact, input).Get(ctx, &output)
    return &output, err
}

func (a *GroundStationStub) CreateConfig(ctx workflow.Context, input *groundstation.CreateConfigInput) (*groundstation.CreateConfigOutput, error) {
    var output groundstation.CreateConfigOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateConfig, input).Get(ctx, &output)
    return &output, err
}

func (a *GroundStationStub) CreateDataflowEndpointGroup(ctx workflow.Context, input *groundstation.CreateDataflowEndpointGroupInput) (*groundstation.CreateDataflowEndpointGroupOutput, error) {
    var output groundstation.CreateDataflowEndpointGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateDataflowEndpointGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *GroundStationStub) CreateMissionProfile(ctx workflow.Context, input *groundstation.CreateMissionProfileInput) (*groundstation.CreateMissionProfileOutput, error) {
    var output groundstation.CreateMissionProfileOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateMissionProfile, input).Get(ctx, &output)
    return &output, err
}

func (a *GroundStationStub) DeleteConfig(ctx workflow.Context, input *groundstation.DeleteConfigInput) (*groundstation.DeleteConfigOutput, error) {
    var output groundstation.DeleteConfigOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteConfig, input).Get(ctx, &output)
    return &output, err
}

func (a *GroundStationStub) DeleteDataflowEndpointGroup(ctx workflow.Context, input *groundstation.DeleteDataflowEndpointGroupInput) (*groundstation.DeleteDataflowEndpointGroupOutput, error) {
    var output groundstation.DeleteDataflowEndpointGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteDataflowEndpointGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *GroundStationStub) DeleteMissionProfile(ctx workflow.Context, input *groundstation.DeleteMissionProfileInput) (*groundstation.DeleteMissionProfileOutput, error) {
    var output groundstation.DeleteMissionProfileOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteMissionProfile, input).Get(ctx, &output)
    return &output, err
}

func (a *GroundStationStub) DescribeContact(ctx workflow.Context, input *groundstation.DescribeContactInput) (*groundstation.DescribeContactOutput, error) {
    var output groundstation.DescribeContactOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeContact, input).Get(ctx, &output)
    return &output, err
}

func (a *GroundStationStub) GetConfig(ctx workflow.Context, input *groundstation.GetConfigInput) (*groundstation.GetConfigOutput, error) {
    var output groundstation.GetConfigOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetConfig, input).Get(ctx, &output)
    return &output, err
}

func (a *GroundStationStub) GetDataflowEndpointGroup(ctx workflow.Context, input *groundstation.GetDataflowEndpointGroupInput) (*groundstation.GetDataflowEndpointGroupOutput, error) {
    var output groundstation.GetDataflowEndpointGroupOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetDataflowEndpointGroup, input).Get(ctx, &output)
    return &output, err
}

func (a *GroundStationStub) GetMinuteUsage(ctx workflow.Context, input *groundstation.GetMinuteUsageInput) (*groundstation.GetMinuteUsageOutput, error) {
    var output groundstation.GetMinuteUsageOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetMinuteUsage, input).Get(ctx, &output)
    return &output, err
}

func (a *GroundStationStub) GetMissionProfile(ctx workflow.Context, input *groundstation.GetMissionProfileInput) (*groundstation.GetMissionProfileOutput, error) {
    var output groundstation.GetMissionProfileOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetMissionProfile, input).Get(ctx, &output)
    return &output, err
}

func (a *GroundStationStub) GetSatellite(ctx workflow.Context, input *groundstation.GetSatelliteInput) (*groundstation.GetSatelliteOutput, error) {
    var output groundstation.GetSatelliteOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetSatellite, input).Get(ctx, &output)
    return &output, err
}

func (a *GroundStationStub) ListConfigs(ctx workflow.Context, input *groundstation.ListConfigsInput) (*groundstation.ListConfigsOutput, error) {
    var output groundstation.ListConfigsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListConfigs, input).Get(ctx, &output)
    return &output, err
}

func (a *GroundStationStub) ListContacts(ctx workflow.Context, input *groundstation.ListContactsInput) (*groundstation.ListContactsOutput, error) {
    var output groundstation.ListContactsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListContacts, input).Get(ctx, &output)
    return &output, err
}

func (a *GroundStationStub) ListDataflowEndpointGroups(ctx workflow.Context, input *groundstation.ListDataflowEndpointGroupsInput) (*groundstation.ListDataflowEndpointGroupsOutput, error) {
    var output groundstation.ListDataflowEndpointGroupsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDataflowEndpointGroups, input).Get(ctx, &output)
    return &output, err
}

func (a *GroundStationStub) ListGroundStations(ctx workflow.Context, input *groundstation.ListGroundStationsInput) (*groundstation.ListGroundStationsOutput, error) {
    var output groundstation.ListGroundStationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListGroundStations, input).Get(ctx, &output)
    return &output, err
}

func (a *GroundStationStub) ListMissionProfiles(ctx workflow.Context, input *groundstation.ListMissionProfilesInput) (*groundstation.ListMissionProfilesOutput, error) {
    var output groundstation.ListMissionProfilesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListMissionProfiles, input).Get(ctx, &output)
    return &output, err
}

func (a *GroundStationStub) ListSatellites(ctx workflow.Context, input *groundstation.ListSatellitesInput) (*groundstation.ListSatellitesOutput, error) {
    var output groundstation.ListSatellitesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListSatellites, input).Get(ctx, &output)
    return &output, err
}

func (a *GroundStationStub) ListTagsForResource(ctx workflow.Context, input *groundstation.ListTagsForResourceInput) (*groundstation.ListTagsForResourceOutput, error) {
    var output groundstation.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *GroundStationStub) ReserveContact(ctx workflow.Context, input *groundstation.ReserveContactInput) (*groundstation.ReserveContactOutput, error) {
    var output groundstation.ReserveContactOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ReserveContact, input).Get(ctx, &output)
    return &output, err
}

func (a *GroundStationStub) TagResource(ctx workflow.Context, input *groundstation.TagResourceInput) (*groundstation.TagResourceOutput, error) {
    var output groundstation.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *GroundStationStub) UntagResource(ctx workflow.Context, input *groundstation.UntagResourceInput) (*groundstation.UntagResourceOutput, error) {
    var output groundstation.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *GroundStationStub) UpdateConfig(ctx workflow.Context, input *groundstation.UpdateConfigInput) (*groundstation.UpdateConfigOutput, error) {
    var output groundstation.UpdateConfigOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateConfig, input).Get(ctx, &output)
    return &output, err
}

func (a *GroundStationStub) UpdateMissionProfile(ctx workflow.Context, input *groundstation.UpdateMissionProfileInput) (*groundstation.UpdateMissionProfileOutput, error) {
    var output groundstation.UpdateMissionProfileOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateMissionProfile, input).Get(ctx, &output)
    return &output, err
}