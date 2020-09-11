package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/medialive"
	"github.com/aws/aws-sdk-go/service/medialive/medialiveiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type MediaLiveActivities struct {
	client medialiveiface.MediaLiveAPI
}

func NewMediaLiveActivities(session *session.Session, config ...*aws.Config) *MediaLiveActivities {
	client := medialive.New(session, config...)
	return &MediaLiveActivities{client: client}
}

func (a *MediaLiveActivities) BatchUpdateSchedule(ctx context.Context, input *medialive.BatchUpdateScheduleInput) (*medialive.BatchUpdateScheduleOutput, error) {
	return a.client.BatchUpdateScheduleWithContext(ctx, input)
}

func (a *MediaLiveActivities) CreateChannel(ctx context.Context, input *medialive.CreateChannelInput) (*medialive.CreateChannelOutput, error) {
	return a.client.CreateChannelWithContext(ctx, input)
}

func (a *MediaLiveActivities) CreateInput(ctx context.Context, input *medialive.CreateInputInput) (*medialive.CreateInputOutput, error) {
	return a.client.CreateInputWithContext(ctx, input)
}

func (a *MediaLiveActivities) CreateInputSecurityGroup(ctx context.Context, input *medialive.CreateInputSecurityGroupInput) (*medialive.CreateInputSecurityGroupOutput, error) {
	return a.client.CreateInputSecurityGroupWithContext(ctx, input)
}

func (a *MediaLiveActivities) CreateMultiplex(ctx context.Context, input *medialive.CreateMultiplexInput) (*medialive.CreateMultiplexOutput, error) {
	return a.client.CreateMultiplexWithContext(ctx, input)
}

func (a *MediaLiveActivities) CreateMultiplexProgram(ctx context.Context, input *medialive.CreateMultiplexProgramInput) (*medialive.CreateMultiplexProgramOutput, error) {
	return a.client.CreateMultiplexProgramWithContext(ctx, input)
}

func (a *MediaLiveActivities) CreateTags(ctx context.Context, input *medialive.CreateTagsInput) (*medialive.CreateTagsOutput, error) {
	return a.client.CreateTagsWithContext(ctx, input)
}

func (a *MediaLiveActivities) DeleteChannel(ctx context.Context, input *medialive.DeleteChannelInput) (*medialive.DeleteChannelOutput, error) {
	return a.client.DeleteChannelWithContext(ctx, input)
}

func (a *MediaLiveActivities) DeleteInput(ctx context.Context, input *medialive.DeleteInputInput) (*medialive.DeleteInputOutput, error) {
	return a.client.DeleteInputWithContext(ctx, input)
}

func (a *MediaLiveActivities) DeleteInputSecurityGroup(ctx context.Context, input *medialive.DeleteInputSecurityGroupInput) (*medialive.DeleteInputSecurityGroupOutput, error) {
	return a.client.DeleteInputSecurityGroupWithContext(ctx, input)
}

func (a *MediaLiveActivities) DeleteMultiplex(ctx context.Context, input *medialive.DeleteMultiplexInput) (*medialive.DeleteMultiplexOutput, error) {
	return a.client.DeleteMultiplexWithContext(ctx, input)
}

func (a *MediaLiveActivities) DeleteMultiplexProgram(ctx context.Context, input *medialive.DeleteMultiplexProgramInput) (*medialive.DeleteMultiplexProgramOutput, error) {
	return a.client.DeleteMultiplexProgramWithContext(ctx, input)
}

func (a *MediaLiveActivities) DeleteReservation(ctx context.Context, input *medialive.DeleteReservationInput) (*medialive.DeleteReservationOutput, error) {
	return a.client.DeleteReservationWithContext(ctx, input)
}

func (a *MediaLiveActivities) DeleteSchedule(ctx context.Context, input *medialive.DeleteScheduleInput) (*medialive.DeleteScheduleOutput, error) {
	return a.client.DeleteScheduleWithContext(ctx, input)
}

func (a *MediaLiveActivities) DeleteTags(ctx context.Context, input *medialive.DeleteTagsInput) (*medialive.DeleteTagsOutput, error) {
	return a.client.DeleteTagsWithContext(ctx, input)
}

func (a *MediaLiveActivities) DescribeChannel(ctx context.Context, input *medialive.DescribeChannelInput) (*medialive.DescribeChannelOutput, error) {
	return a.client.DescribeChannelWithContext(ctx, input)
}

func (a *MediaLiveActivities) DescribeInput(ctx context.Context, input *medialive.DescribeInputInput) (*medialive.DescribeInputOutput, error) {
	return a.client.DescribeInputWithContext(ctx, input)
}

func (a *MediaLiveActivities) DescribeInputDevice(ctx context.Context, input *medialive.DescribeInputDeviceInput) (*medialive.DescribeInputDeviceOutput, error) {
	return a.client.DescribeInputDeviceWithContext(ctx, input)
}

func (a *MediaLiveActivities) DescribeInputDeviceThumbnail(ctx context.Context, input *medialive.DescribeInputDeviceThumbnailInput) (*medialive.DescribeInputDeviceThumbnailOutput, error) {
	return a.client.DescribeInputDeviceThumbnailWithContext(ctx, input)
}

func (a *MediaLiveActivities) DescribeInputSecurityGroup(ctx context.Context, input *medialive.DescribeInputSecurityGroupInput) (*medialive.DescribeInputSecurityGroupOutput, error) {
	return a.client.DescribeInputSecurityGroupWithContext(ctx, input)
}

func (a *MediaLiveActivities) DescribeMultiplex(ctx context.Context, input *medialive.DescribeMultiplexInput) (*medialive.DescribeMultiplexOutput, error) {
	return a.client.DescribeMultiplexWithContext(ctx, input)
}

func (a *MediaLiveActivities) DescribeMultiplexProgram(ctx context.Context, input *medialive.DescribeMultiplexProgramInput) (*medialive.DescribeMultiplexProgramOutput, error) {
	return a.client.DescribeMultiplexProgramWithContext(ctx, input)
}

func (a *MediaLiveActivities) DescribeOffering(ctx context.Context, input *medialive.DescribeOfferingInput) (*medialive.DescribeOfferingOutput, error) {
	return a.client.DescribeOfferingWithContext(ctx, input)
}

func (a *MediaLiveActivities) DescribeReservation(ctx context.Context, input *medialive.DescribeReservationInput) (*medialive.DescribeReservationOutput, error) {
	return a.client.DescribeReservationWithContext(ctx, input)
}

func (a *MediaLiveActivities) DescribeSchedule(ctx context.Context, input *medialive.DescribeScheduleInput) (*medialive.DescribeScheduleOutput, error) {
	return a.client.DescribeScheduleWithContext(ctx, input)
}

func (a *MediaLiveActivities) ListChannels(ctx context.Context, input *medialive.ListChannelsInput) (*medialive.ListChannelsOutput, error) {
	return a.client.ListChannelsWithContext(ctx, input)
}

func (a *MediaLiveActivities) ListInputDevices(ctx context.Context, input *medialive.ListInputDevicesInput) (*medialive.ListInputDevicesOutput, error) {
	return a.client.ListInputDevicesWithContext(ctx, input)
}

func (a *MediaLiveActivities) ListInputSecurityGroups(ctx context.Context, input *medialive.ListInputSecurityGroupsInput) (*medialive.ListInputSecurityGroupsOutput, error) {
	return a.client.ListInputSecurityGroupsWithContext(ctx, input)
}

func (a *MediaLiveActivities) ListInputs(ctx context.Context, input *medialive.ListInputsInput) (*medialive.ListInputsOutput, error) {
	return a.client.ListInputsWithContext(ctx, input)
}

func (a *MediaLiveActivities) ListMultiplexPrograms(ctx context.Context, input *medialive.ListMultiplexProgramsInput) (*medialive.ListMultiplexProgramsOutput, error) {
	return a.client.ListMultiplexProgramsWithContext(ctx, input)
}

func (a *MediaLiveActivities) ListMultiplexes(ctx context.Context, input *medialive.ListMultiplexesInput) (*medialive.ListMultiplexesOutput, error) {
	return a.client.ListMultiplexesWithContext(ctx, input)
}

func (a *MediaLiveActivities) ListOfferings(ctx context.Context, input *medialive.ListOfferingsInput) (*medialive.ListOfferingsOutput, error) {
	return a.client.ListOfferingsWithContext(ctx, input)
}

func (a *MediaLiveActivities) ListReservations(ctx context.Context, input *medialive.ListReservationsInput) (*medialive.ListReservationsOutput, error) {
	return a.client.ListReservationsWithContext(ctx, input)
}

func (a *MediaLiveActivities) ListTagsForResource(ctx context.Context, input *medialive.ListTagsForResourceInput) (*medialive.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *MediaLiveActivities) PurchaseOffering(ctx context.Context, input *medialive.PurchaseOfferingInput) (*medialive.PurchaseOfferingOutput, error) {
	return a.client.PurchaseOfferingWithContext(ctx, input)
}

func (a *MediaLiveActivities) StartChannel(ctx context.Context, input *medialive.StartChannelInput) (*medialive.StartChannelOutput, error) {
	return a.client.StartChannelWithContext(ctx, input)
}

func (a *MediaLiveActivities) StartMultiplex(ctx context.Context, input *medialive.StartMultiplexInput) (*medialive.StartMultiplexOutput, error) {
	return a.client.StartMultiplexWithContext(ctx, input)
}

func (a *MediaLiveActivities) StopChannel(ctx context.Context, input *medialive.StopChannelInput) (*medialive.StopChannelOutput, error) {
	return a.client.StopChannelWithContext(ctx, input)
}

func (a *MediaLiveActivities) StopMultiplex(ctx context.Context, input *medialive.StopMultiplexInput) (*medialive.StopMultiplexOutput, error) {
	return a.client.StopMultiplexWithContext(ctx, input)
}

func (a *MediaLiveActivities) UpdateChannel(ctx context.Context, input *medialive.UpdateChannelInput) (*medialive.UpdateChannelOutput, error) {
	return a.client.UpdateChannelWithContext(ctx, input)
}

func (a *MediaLiveActivities) UpdateChannelClass(ctx context.Context, input *medialive.UpdateChannelClassInput) (*medialive.UpdateChannelClassOutput, error) {
	return a.client.UpdateChannelClassWithContext(ctx, input)
}

func (a *MediaLiveActivities) UpdateInput(ctx context.Context, input *medialive.UpdateInputInput) (*medialive.UpdateInputOutput, error) {
	return a.client.UpdateInputWithContext(ctx, input)
}

func (a *MediaLiveActivities) UpdateInputDevice(ctx context.Context, input *medialive.UpdateInputDeviceInput) (*medialive.UpdateInputDeviceOutput, error) {
	return a.client.UpdateInputDeviceWithContext(ctx, input)
}

func (a *MediaLiveActivities) UpdateInputSecurityGroup(ctx context.Context, input *medialive.UpdateInputSecurityGroupInput) (*medialive.UpdateInputSecurityGroupOutput, error) {
	return a.client.UpdateInputSecurityGroupWithContext(ctx, input)
}

func (a *MediaLiveActivities) UpdateMultiplex(ctx context.Context, input *medialive.UpdateMultiplexInput) (*medialive.UpdateMultiplexOutput, error) {
	return a.client.UpdateMultiplexWithContext(ctx, input)
}

func (a *MediaLiveActivities) UpdateMultiplexProgram(ctx context.Context, input *medialive.UpdateMultiplexProgramInput) (*medialive.UpdateMultiplexProgramOutput, error) {
	return a.client.UpdateMultiplexProgramWithContext(ctx, input)
}

func (a *MediaLiveActivities) UpdateReservation(ctx context.Context, input *medialive.UpdateReservationInput) (*medialive.UpdateReservationOutput, error) {
	return a.client.UpdateReservationWithContext(ctx, input)
}

func (a *MediaLiveActivities) WaitUntilChannelCreated(ctx context.Context, input *medialive.DescribeChannelInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilChannelCreatedWithContext(ctx, input, options...)
	})
}

func (a *MediaLiveActivities) WaitUntilChannelDeleted(ctx context.Context, input *medialive.DescribeChannelInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilChannelDeletedWithContext(ctx, input, options...)
	})
}

func (a *MediaLiveActivities) WaitUntilChannelRunning(ctx context.Context, input *medialive.DescribeChannelInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilChannelRunningWithContext(ctx, input, options...)
	})
}

func (a *MediaLiveActivities) WaitUntilChannelStopped(ctx context.Context, input *medialive.DescribeChannelInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilChannelStoppedWithContext(ctx, input, options...)
	})
}

func (a *MediaLiveActivities) WaitUntilInputAttached(ctx context.Context, input *medialive.DescribeInputInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilInputAttachedWithContext(ctx, input, options...)
	})
}

func (a *MediaLiveActivities) WaitUntilInputDeleted(ctx context.Context, input *medialive.DescribeInputInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilInputDeletedWithContext(ctx, input, options...)
	})
}

func (a *MediaLiveActivities) WaitUntilInputDetached(ctx context.Context, input *medialive.DescribeInputInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilInputDetachedWithContext(ctx, input, options...)
	})
}

func (a *MediaLiveActivities) WaitUntilMultiplexCreated(ctx context.Context, input *medialive.DescribeMultiplexInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilMultiplexCreatedWithContext(ctx, input, options...)
	})
}

func (a *MediaLiveActivities) WaitUntilMultiplexDeleted(ctx context.Context, input *medialive.DescribeMultiplexInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilMultiplexDeletedWithContext(ctx, input, options...)
	})
}

func (a *MediaLiveActivities) WaitUntilMultiplexRunning(ctx context.Context, input *medialive.DescribeMultiplexInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilMultiplexRunningWithContext(ctx, input, options...)
	})
}

func (a *MediaLiveActivities) WaitUntilMultiplexStopped(ctx context.Context, input *medialive.DescribeMultiplexInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilMultiplexStoppedWithContext(ctx, input, options...)
	})
}
