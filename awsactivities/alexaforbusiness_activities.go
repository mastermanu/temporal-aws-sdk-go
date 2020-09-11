package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/alexaforbusiness"
	"github.com/aws/aws-sdk-go/service/alexaforbusiness/alexaforbusinessiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type AlexaForBusinessActivities struct {
	client alexaforbusinessiface.AlexaForBusinessAPI
}

func NewAlexaForBusinessActivities(session *session.Session, config ...*aws.Config) *AlexaForBusinessActivities {
	client := alexaforbusiness.New(session, config...)
	return &AlexaForBusinessActivities{client: client}
}

func (a *AlexaForBusinessActivities) ApproveSkill(ctx context.Context, input *alexaforbusiness.ApproveSkillInput) (*alexaforbusiness.ApproveSkillOutput, error) {
	return a.client.ApproveSkillWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) AssociateContactWithAddressBook(ctx context.Context, input *alexaforbusiness.AssociateContactWithAddressBookInput) (*alexaforbusiness.AssociateContactWithAddressBookOutput, error) {
	return a.client.AssociateContactWithAddressBookWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) AssociateDeviceWithNetworkProfile(ctx context.Context, input *alexaforbusiness.AssociateDeviceWithNetworkProfileInput) (*alexaforbusiness.AssociateDeviceWithNetworkProfileOutput, error) {
	return a.client.AssociateDeviceWithNetworkProfileWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) AssociateDeviceWithRoom(ctx context.Context, input *alexaforbusiness.AssociateDeviceWithRoomInput) (*alexaforbusiness.AssociateDeviceWithRoomOutput, error) {
	return a.client.AssociateDeviceWithRoomWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) AssociateSkillGroupWithRoom(ctx context.Context, input *alexaforbusiness.AssociateSkillGroupWithRoomInput) (*alexaforbusiness.AssociateSkillGroupWithRoomOutput, error) {
	return a.client.AssociateSkillGroupWithRoomWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) AssociateSkillWithSkillGroup(ctx context.Context, input *alexaforbusiness.AssociateSkillWithSkillGroupInput) (*alexaforbusiness.AssociateSkillWithSkillGroupOutput, error) {
	return a.client.AssociateSkillWithSkillGroupWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) AssociateSkillWithUsers(ctx context.Context, input *alexaforbusiness.AssociateSkillWithUsersInput) (*alexaforbusiness.AssociateSkillWithUsersOutput, error) {
	return a.client.AssociateSkillWithUsersWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) CreateAddressBook(ctx context.Context, input *alexaforbusiness.CreateAddressBookInput) (*alexaforbusiness.CreateAddressBookOutput, error) {
	return a.client.CreateAddressBookWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) CreateBusinessReportSchedule(ctx context.Context, input *alexaforbusiness.CreateBusinessReportScheduleInput) (*alexaforbusiness.CreateBusinessReportScheduleOutput, error) {
	return a.client.CreateBusinessReportScheduleWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) CreateConferenceProvider(ctx context.Context, input *alexaforbusiness.CreateConferenceProviderInput) (*alexaforbusiness.CreateConferenceProviderOutput, error) {
	return a.client.CreateConferenceProviderWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) CreateContact(ctx context.Context, input *alexaforbusiness.CreateContactInput) (*alexaforbusiness.CreateContactOutput, error) {
	return a.client.CreateContactWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) CreateGatewayGroup(ctx context.Context, input *alexaforbusiness.CreateGatewayGroupInput) (*alexaforbusiness.CreateGatewayGroupOutput, error) {
	return a.client.CreateGatewayGroupWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) CreateNetworkProfile(ctx context.Context, input *alexaforbusiness.CreateNetworkProfileInput) (*alexaforbusiness.CreateNetworkProfileOutput, error) {
	return a.client.CreateNetworkProfileWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) CreateProfile(ctx context.Context, input *alexaforbusiness.CreateProfileInput) (*alexaforbusiness.CreateProfileOutput, error) {
	return a.client.CreateProfileWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) CreateRoom(ctx context.Context, input *alexaforbusiness.CreateRoomInput) (*alexaforbusiness.CreateRoomOutput, error) {
	return a.client.CreateRoomWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) CreateSkillGroup(ctx context.Context, input *alexaforbusiness.CreateSkillGroupInput) (*alexaforbusiness.CreateSkillGroupOutput, error) {
	return a.client.CreateSkillGroupWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) CreateUser(ctx context.Context, input *alexaforbusiness.CreateUserInput) (*alexaforbusiness.CreateUserOutput, error) {
	return a.client.CreateUserWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) DeleteAddressBook(ctx context.Context, input *alexaforbusiness.DeleteAddressBookInput) (*alexaforbusiness.DeleteAddressBookOutput, error) {
	return a.client.DeleteAddressBookWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) DeleteBusinessReportSchedule(ctx context.Context, input *alexaforbusiness.DeleteBusinessReportScheduleInput) (*alexaforbusiness.DeleteBusinessReportScheduleOutput, error) {
	return a.client.DeleteBusinessReportScheduleWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) DeleteConferenceProvider(ctx context.Context, input *alexaforbusiness.DeleteConferenceProviderInput) (*alexaforbusiness.DeleteConferenceProviderOutput, error) {
	return a.client.DeleteConferenceProviderWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) DeleteContact(ctx context.Context, input *alexaforbusiness.DeleteContactInput) (*alexaforbusiness.DeleteContactOutput, error) {
	return a.client.DeleteContactWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) DeleteDevice(ctx context.Context, input *alexaforbusiness.DeleteDeviceInput) (*alexaforbusiness.DeleteDeviceOutput, error) {
	return a.client.DeleteDeviceWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) DeleteDeviceUsageData(ctx context.Context, input *alexaforbusiness.DeleteDeviceUsageDataInput) (*alexaforbusiness.DeleteDeviceUsageDataOutput, error) {
	return a.client.DeleteDeviceUsageDataWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) DeleteGatewayGroup(ctx context.Context, input *alexaforbusiness.DeleteGatewayGroupInput) (*alexaforbusiness.DeleteGatewayGroupOutput, error) {
	return a.client.DeleteGatewayGroupWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) DeleteNetworkProfile(ctx context.Context, input *alexaforbusiness.DeleteNetworkProfileInput) (*alexaforbusiness.DeleteNetworkProfileOutput, error) {
	return a.client.DeleteNetworkProfileWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) DeleteProfile(ctx context.Context, input *alexaforbusiness.DeleteProfileInput) (*alexaforbusiness.DeleteProfileOutput, error) {
	return a.client.DeleteProfileWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) DeleteRoom(ctx context.Context, input *alexaforbusiness.DeleteRoomInput) (*alexaforbusiness.DeleteRoomOutput, error) {
	return a.client.DeleteRoomWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) DeleteRoomSkillParameter(ctx context.Context, input *alexaforbusiness.DeleteRoomSkillParameterInput) (*alexaforbusiness.DeleteRoomSkillParameterOutput, error) {
	return a.client.DeleteRoomSkillParameterWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) DeleteSkillAuthorization(ctx context.Context, input *alexaforbusiness.DeleteSkillAuthorizationInput) (*alexaforbusiness.DeleteSkillAuthorizationOutput, error) {
	return a.client.DeleteSkillAuthorizationWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) DeleteSkillGroup(ctx context.Context, input *alexaforbusiness.DeleteSkillGroupInput) (*alexaforbusiness.DeleteSkillGroupOutput, error) {
	return a.client.DeleteSkillGroupWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) DeleteUser(ctx context.Context, input *alexaforbusiness.DeleteUserInput) (*alexaforbusiness.DeleteUserOutput, error) {
	return a.client.DeleteUserWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) DisassociateContactFromAddressBook(ctx context.Context, input *alexaforbusiness.DisassociateContactFromAddressBookInput) (*alexaforbusiness.DisassociateContactFromAddressBookOutput, error) {
	return a.client.DisassociateContactFromAddressBookWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) DisassociateDeviceFromRoom(ctx context.Context, input *alexaforbusiness.DisassociateDeviceFromRoomInput) (*alexaforbusiness.DisassociateDeviceFromRoomOutput, error) {
	return a.client.DisassociateDeviceFromRoomWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) DisassociateSkillFromSkillGroup(ctx context.Context, input *alexaforbusiness.DisassociateSkillFromSkillGroupInput) (*alexaforbusiness.DisassociateSkillFromSkillGroupOutput, error) {
	return a.client.DisassociateSkillFromSkillGroupWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) DisassociateSkillFromUsers(ctx context.Context, input *alexaforbusiness.DisassociateSkillFromUsersInput) (*alexaforbusiness.DisassociateSkillFromUsersOutput, error) {
	return a.client.DisassociateSkillFromUsersWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) DisassociateSkillGroupFromRoom(ctx context.Context, input *alexaforbusiness.DisassociateSkillGroupFromRoomInput) (*alexaforbusiness.DisassociateSkillGroupFromRoomOutput, error) {
	return a.client.DisassociateSkillGroupFromRoomWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) ForgetSmartHomeAppliances(ctx context.Context, input *alexaforbusiness.ForgetSmartHomeAppliancesInput) (*alexaforbusiness.ForgetSmartHomeAppliancesOutput, error) {
	return a.client.ForgetSmartHomeAppliancesWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) GetAddressBook(ctx context.Context, input *alexaforbusiness.GetAddressBookInput) (*alexaforbusiness.GetAddressBookOutput, error) {
	return a.client.GetAddressBookWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) GetConferencePreference(ctx context.Context, input *alexaforbusiness.GetConferencePreferenceInput) (*alexaforbusiness.GetConferencePreferenceOutput, error) {
	return a.client.GetConferencePreferenceWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) GetConferenceProvider(ctx context.Context, input *alexaforbusiness.GetConferenceProviderInput) (*alexaforbusiness.GetConferenceProviderOutput, error) {
	return a.client.GetConferenceProviderWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) GetContact(ctx context.Context, input *alexaforbusiness.GetContactInput) (*alexaforbusiness.GetContactOutput, error) {
	return a.client.GetContactWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) GetDevice(ctx context.Context, input *alexaforbusiness.GetDeviceInput) (*alexaforbusiness.GetDeviceOutput, error) {
	return a.client.GetDeviceWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) GetGateway(ctx context.Context, input *alexaforbusiness.GetGatewayInput) (*alexaforbusiness.GetGatewayOutput, error) {
	return a.client.GetGatewayWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) GetGatewayGroup(ctx context.Context, input *alexaforbusiness.GetGatewayGroupInput) (*alexaforbusiness.GetGatewayGroupOutput, error) {
	return a.client.GetGatewayGroupWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) GetInvitationConfiguration(ctx context.Context, input *alexaforbusiness.GetInvitationConfigurationInput) (*alexaforbusiness.GetInvitationConfigurationOutput, error) {
	return a.client.GetInvitationConfigurationWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) GetNetworkProfile(ctx context.Context, input *alexaforbusiness.GetNetworkProfileInput) (*alexaforbusiness.GetNetworkProfileOutput, error) {
	return a.client.GetNetworkProfileWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) GetProfile(ctx context.Context, input *alexaforbusiness.GetProfileInput) (*alexaforbusiness.GetProfileOutput, error) {
	return a.client.GetProfileWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) GetRoom(ctx context.Context, input *alexaforbusiness.GetRoomInput) (*alexaforbusiness.GetRoomOutput, error) {
	return a.client.GetRoomWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) GetRoomSkillParameter(ctx context.Context, input *alexaforbusiness.GetRoomSkillParameterInput) (*alexaforbusiness.GetRoomSkillParameterOutput, error) {
	return a.client.GetRoomSkillParameterWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) GetSkillGroup(ctx context.Context, input *alexaforbusiness.GetSkillGroupInput) (*alexaforbusiness.GetSkillGroupOutput, error) {
	return a.client.GetSkillGroupWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) ListBusinessReportSchedules(ctx context.Context, input *alexaforbusiness.ListBusinessReportSchedulesInput) (*alexaforbusiness.ListBusinessReportSchedulesOutput, error) {
	return a.client.ListBusinessReportSchedulesWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) ListConferenceProviders(ctx context.Context, input *alexaforbusiness.ListConferenceProvidersInput) (*alexaforbusiness.ListConferenceProvidersOutput, error) {
	return a.client.ListConferenceProvidersWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) ListDeviceEvents(ctx context.Context, input *alexaforbusiness.ListDeviceEventsInput) (*alexaforbusiness.ListDeviceEventsOutput, error) {
	return a.client.ListDeviceEventsWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) ListGatewayGroups(ctx context.Context, input *alexaforbusiness.ListGatewayGroupsInput) (*alexaforbusiness.ListGatewayGroupsOutput, error) {
	return a.client.ListGatewayGroupsWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) ListGateways(ctx context.Context, input *alexaforbusiness.ListGatewaysInput) (*alexaforbusiness.ListGatewaysOutput, error) {
	return a.client.ListGatewaysWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) ListSkills(ctx context.Context, input *alexaforbusiness.ListSkillsInput) (*alexaforbusiness.ListSkillsOutput, error) {
	return a.client.ListSkillsWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) ListSkillsStoreCategories(ctx context.Context, input *alexaforbusiness.ListSkillsStoreCategoriesInput) (*alexaforbusiness.ListSkillsStoreCategoriesOutput, error) {
	return a.client.ListSkillsStoreCategoriesWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) ListSkillsStoreSkillsByCategory(ctx context.Context, input *alexaforbusiness.ListSkillsStoreSkillsByCategoryInput) (*alexaforbusiness.ListSkillsStoreSkillsByCategoryOutput, error) {
	return a.client.ListSkillsStoreSkillsByCategoryWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) ListSmartHomeAppliances(ctx context.Context, input *alexaforbusiness.ListSmartHomeAppliancesInput) (*alexaforbusiness.ListSmartHomeAppliancesOutput, error) {
	return a.client.ListSmartHomeAppliancesWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) ListTags(ctx context.Context, input *alexaforbusiness.ListTagsInput) (*alexaforbusiness.ListTagsOutput, error) {
	return a.client.ListTagsWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) PutConferencePreference(ctx context.Context, input *alexaforbusiness.PutConferencePreferenceInput) (*alexaforbusiness.PutConferencePreferenceOutput, error) {
	return a.client.PutConferencePreferenceWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) PutInvitationConfiguration(ctx context.Context, input *alexaforbusiness.PutInvitationConfigurationInput) (*alexaforbusiness.PutInvitationConfigurationOutput, error) {
	return a.client.PutInvitationConfigurationWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) PutRoomSkillParameter(ctx context.Context, input *alexaforbusiness.PutRoomSkillParameterInput) (*alexaforbusiness.PutRoomSkillParameterOutput, error) {
	return a.client.PutRoomSkillParameterWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) PutSkillAuthorization(ctx context.Context, input *alexaforbusiness.PutSkillAuthorizationInput) (*alexaforbusiness.PutSkillAuthorizationOutput, error) {
	return a.client.PutSkillAuthorizationWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) RegisterAVSDevice(ctx context.Context, input *alexaforbusiness.RegisterAVSDeviceInput) (*alexaforbusiness.RegisterAVSDeviceOutput, error) {
	return a.client.RegisterAVSDeviceWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) RejectSkill(ctx context.Context, input *alexaforbusiness.RejectSkillInput) (*alexaforbusiness.RejectSkillOutput, error) {
	return a.client.RejectSkillWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) ResolveRoom(ctx context.Context, input *alexaforbusiness.ResolveRoomInput) (*alexaforbusiness.ResolveRoomOutput, error) {
	return a.client.ResolveRoomWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) RevokeInvitation(ctx context.Context, input *alexaforbusiness.RevokeInvitationInput) (*alexaforbusiness.RevokeInvitationOutput, error) {
	return a.client.RevokeInvitationWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) SearchAddressBooks(ctx context.Context, input *alexaforbusiness.SearchAddressBooksInput) (*alexaforbusiness.SearchAddressBooksOutput, error) {
	return a.client.SearchAddressBooksWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) SearchContacts(ctx context.Context, input *alexaforbusiness.SearchContactsInput) (*alexaforbusiness.SearchContactsOutput, error) {
	return a.client.SearchContactsWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) SearchDevices(ctx context.Context, input *alexaforbusiness.SearchDevicesInput) (*alexaforbusiness.SearchDevicesOutput, error) {
	return a.client.SearchDevicesWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) SearchNetworkProfiles(ctx context.Context, input *alexaforbusiness.SearchNetworkProfilesInput) (*alexaforbusiness.SearchNetworkProfilesOutput, error) {
	return a.client.SearchNetworkProfilesWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) SearchProfiles(ctx context.Context, input *alexaforbusiness.SearchProfilesInput) (*alexaforbusiness.SearchProfilesOutput, error) {
	return a.client.SearchProfilesWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) SearchRooms(ctx context.Context, input *alexaforbusiness.SearchRoomsInput) (*alexaforbusiness.SearchRoomsOutput, error) {
	return a.client.SearchRoomsWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) SearchSkillGroups(ctx context.Context, input *alexaforbusiness.SearchSkillGroupsInput) (*alexaforbusiness.SearchSkillGroupsOutput, error) {
	return a.client.SearchSkillGroupsWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) SearchUsers(ctx context.Context, input *alexaforbusiness.SearchUsersInput) (*alexaforbusiness.SearchUsersOutput, error) {
	return a.client.SearchUsersWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) SendAnnouncement(ctx context.Context, input *alexaforbusiness.SendAnnouncementInput) (*alexaforbusiness.SendAnnouncementOutput, error) {
	return a.client.SendAnnouncementWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) SendInvitation(ctx context.Context, input *alexaforbusiness.SendInvitationInput) (*alexaforbusiness.SendInvitationOutput, error) {
	return a.client.SendInvitationWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) StartDeviceSync(ctx context.Context, input *alexaforbusiness.StartDeviceSyncInput) (*alexaforbusiness.StartDeviceSyncOutput, error) {
	return a.client.StartDeviceSyncWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) StartSmartHomeApplianceDiscovery(ctx context.Context, input *alexaforbusiness.StartSmartHomeApplianceDiscoveryInput) (*alexaforbusiness.StartSmartHomeApplianceDiscoveryOutput, error) {
	return a.client.StartSmartHomeApplianceDiscoveryWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) TagResource(ctx context.Context, input *alexaforbusiness.TagResourceInput) (*alexaforbusiness.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) UntagResource(ctx context.Context, input *alexaforbusiness.UntagResourceInput) (*alexaforbusiness.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) UpdateAddressBook(ctx context.Context, input *alexaforbusiness.UpdateAddressBookInput) (*alexaforbusiness.UpdateAddressBookOutput, error) {
	return a.client.UpdateAddressBookWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) UpdateBusinessReportSchedule(ctx context.Context, input *alexaforbusiness.UpdateBusinessReportScheduleInput) (*alexaforbusiness.UpdateBusinessReportScheduleOutput, error) {
	return a.client.UpdateBusinessReportScheduleWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) UpdateConferenceProvider(ctx context.Context, input *alexaforbusiness.UpdateConferenceProviderInput) (*alexaforbusiness.UpdateConferenceProviderOutput, error) {
	return a.client.UpdateConferenceProviderWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) UpdateContact(ctx context.Context, input *alexaforbusiness.UpdateContactInput) (*alexaforbusiness.UpdateContactOutput, error) {
	return a.client.UpdateContactWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) UpdateDevice(ctx context.Context, input *alexaforbusiness.UpdateDeviceInput) (*alexaforbusiness.UpdateDeviceOutput, error) {
	return a.client.UpdateDeviceWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) UpdateGateway(ctx context.Context, input *alexaforbusiness.UpdateGatewayInput) (*alexaforbusiness.UpdateGatewayOutput, error) {
	return a.client.UpdateGatewayWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) UpdateGatewayGroup(ctx context.Context, input *alexaforbusiness.UpdateGatewayGroupInput) (*alexaforbusiness.UpdateGatewayGroupOutput, error) {
	return a.client.UpdateGatewayGroupWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) UpdateNetworkProfile(ctx context.Context, input *alexaforbusiness.UpdateNetworkProfileInput) (*alexaforbusiness.UpdateNetworkProfileOutput, error) {
	return a.client.UpdateNetworkProfileWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) UpdateProfile(ctx context.Context, input *alexaforbusiness.UpdateProfileInput) (*alexaforbusiness.UpdateProfileOutput, error) {
	return a.client.UpdateProfileWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) UpdateRoom(ctx context.Context, input *alexaforbusiness.UpdateRoomInput) (*alexaforbusiness.UpdateRoomOutput, error) {
	return a.client.UpdateRoomWithContext(ctx, input)
}

func (a *AlexaForBusinessActivities) UpdateSkillGroup(ctx context.Context, input *alexaforbusiness.UpdateSkillGroupInput) (*alexaforbusiness.UpdateSkillGroupOutput, error) {
	return a.client.UpdateSkillGroupWithContext(ctx, input)
}
