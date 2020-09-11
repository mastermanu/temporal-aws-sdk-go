package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/networkmanager"
	"github.com/aws/aws-sdk-go/service/networkmanager/networkmanageriface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type NetworkManagerActivities struct {
	client networkmanageriface.NetworkManagerAPI
}

func NewNetworkManagerActivities(session *session.Session, config ...*aws.Config) *NetworkManagerActivities {
	client := networkmanager.New(session, config...)
	return &NetworkManagerActivities{client: client}
}

func (a *NetworkManagerActivities) AssociateCustomerGateway(ctx context.Context, input *networkmanager.AssociateCustomerGatewayInput) (*networkmanager.AssociateCustomerGatewayOutput, error) {
	return a.client.AssociateCustomerGatewayWithContext(ctx, input)
}

func (a *NetworkManagerActivities) AssociateLink(ctx context.Context, input *networkmanager.AssociateLinkInput) (*networkmanager.AssociateLinkOutput, error) {
	return a.client.AssociateLinkWithContext(ctx, input)
}

func (a *NetworkManagerActivities) CreateDevice(ctx context.Context, input *networkmanager.CreateDeviceInput) (*networkmanager.CreateDeviceOutput, error) {
	return a.client.CreateDeviceWithContext(ctx, input)
}

func (a *NetworkManagerActivities) CreateGlobalNetwork(ctx context.Context, input *networkmanager.CreateGlobalNetworkInput) (*networkmanager.CreateGlobalNetworkOutput, error) {
	return a.client.CreateGlobalNetworkWithContext(ctx, input)
}

func (a *NetworkManagerActivities) CreateLink(ctx context.Context, input *networkmanager.CreateLinkInput) (*networkmanager.CreateLinkOutput, error) {
	return a.client.CreateLinkWithContext(ctx, input)
}

func (a *NetworkManagerActivities) CreateSite(ctx context.Context, input *networkmanager.CreateSiteInput) (*networkmanager.CreateSiteOutput, error) {
	return a.client.CreateSiteWithContext(ctx, input)
}

func (a *NetworkManagerActivities) DeleteDevice(ctx context.Context, input *networkmanager.DeleteDeviceInput) (*networkmanager.DeleteDeviceOutput, error) {
	return a.client.DeleteDeviceWithContext(ctx, input)
}

func (a *NetworkManagerActivities) DeleteGlobalNetwork(ctx context.Context, input *networkmanager.DeleteGlobalNetworkInput) (*networkmanager.DeleteGlobalNetworkOutput, error) {
	return a.client.DeleteGlobalNetworkWithContext(ctx, input)
}

func (a *NetworkManagerActivities) DeleteLink(ctx context.Context, input *networkmanager.DeleteLinkInput) (*networkmanager.DeleteLinkOutput, error) {
	return a.client.DeleteLinkWithContext(ctx, input)
}

func (a *NetworkManagerActivities) DeleteSite(ctx context.Context, input *networkmanager.DeleteSiteInput) (*networkmanager.DeleteSiteOutput, error) {
	return a.client.DeleteSiteWithContext(ctx, input)
}

func (a *NetworkManagerActivities) DeregisterTransitGateway(ctx context.Context, input *networkmanager.DeregisterTransitGatewayInput) (*networkmanager.DeregisterTransitGatewayOutput, error) {
	return a.client.DeregisterTransitGatewayWithContext(ctx, input)
}

func (a *NetworkManagerActivities) DescribeGlobalNetworks(ctx context.Context, input *networkmanager.DescribeGlobalNetworksInput) (*networkmanager.DescribeGlobalNetworksOutput, error) {
	return a.client.DescribeGlobalNetworksWithContext(ctx, input)
}

func (a *NetworkManagerActivities) DisassociateCustomerGateway(ctx context.Context, input *networkmanager.DisassociateCustomerGatewayInput) (*networkmanager.DisassociateCustomerGatewayOutput, error) {
	return a.client.DisassociateCustomerGatewayWithContext(ctx, input)
}

func (a *NetworkManagerActivities) DisassociateLink(ctx context.Context, input *networkmanager.DisassociateLinkInput) (*networkmanager.DisassociateLinkOutput, error) {
	return a.client.DisassociateLinkWithContext(ctx, input)
}

func (a *NetworkManagerActivities) GetCustomerGatewayAssociations(ctx context.Context, input *networkmanager.GetCustomerGatewayAssociationsInput) (*networkmanager.GetCustomerGatewayAssociationsOutput, error) {
	return a.client.GetCustomerGatewayAssociationsWithContext(ctx, input)
}

func (a *NetworkManagerActivities) GetDevices(ctx context.Context, input *networkmanager.GetDevicesInput) (*networkmanager.GetDevicesOutput, error) {
	return a.client.GetDevicesWithContext(ctx, input)
}

func (a *NetworkManagerActivities) GetLinkAssociations(ctx context.Context, input *networkmanager.GetLinkAssociationsInput) (*networkmanager.GetLinkAssociationsOutput, error) {
	return a.client.GetLinkAssociationsWithContext(ctx, input)
}

func (a *NetworkManagerActivities) GetLinks(ctx context.Context, input *networkmanager.GetLinksInput) (*networkmanager.GetLinksOutput, error) {
	return a.client.GetLinksWithContext(ctx, input)
}

func (a *NetworkManagerActivities) GetSites(ctx context.Context, input *networkmanager.GetSitesInput) (*networkmanager.GetSitesOutput, error) {
	return a.client.GetSitesWithContext(ctx, input)
}

func (a *NetworkManagerActivities) GetTransitGatewayRegistrations(ctx context.Context, input *networkmanager.GetTransitGatewayRegistrationsInput) (*networkmanager.GetTransitGatewayRegistrationsOutput, error) {
	return a.client.GetTransitGatewayRegistrationsWithContext(ctx, input)
}

func (a *NetworkManagerActivities) ListTagsForResource(ctx context.Context, input *networkmanager.ListTagsForResourceInput) (*networkmanager.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *NetworkManagerActivities) RegisterTransitGateway(ctx context.Context, input *networkmanager.RegisterTransitGatewayInput) (*networkmanager.RegisterTransitGatewayOutput, error) {
	return a.client.RegisterTransitGatewayWithContext(ctx, input)
}

func (a *NetworkManagerActivities) TagResource(ctx context.Context, input *networkmanager.TagResourceInput) (*networkmanager.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *NetworkManagerActivities) UntagResource(ctx context.Context, input *networkmanager.UntagResourceInput) (*networkmanager.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}

func (a *NetworkManagerActivities) UpdateDevice(ctx context.Context, input *networkmanager.UpdateDeviceInput) (*networkmanager.UpdateDeviceOutput, error) {
	return a.client.UpdateDeviceWithContext(ctx, input)
}

func (a *NetworkManagerActivities) UpdateGlobalNetwork(ctx context.Context, input *networkmanager.UpdateGlobalNetworkInput) (*networkmanager.UpdateGlobalNetworkOutput, error) {
	return a.client.UpdateGlobalNetworkWithContext(ctx, input)
}

func (a *NetworkManagerActivities) UpdateLink(ctx context.Context, input *networkmanager.UpdateLinkInput) (*networkmanager.UpdateLinkOutput, error) {
	return a.client.UpdateLinkWithContext(ctx, input)
}

func (a *NetworkManagerActivities) UpdateSite(ctx context.Context, input *networkmanager.UpdateSiteInput) (*networkmanager.UpdateSiteOutput, error) {
	return a.client.UpdateSiteWithContext(ctx, input)
}
