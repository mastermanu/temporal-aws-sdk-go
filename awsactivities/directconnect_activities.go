// Generated by https://github.com/temporalio/temporal-aws-sdk/cmd/temporal-aws-sdk-gen/main.go
// from https://github.com/aws/aws-sdk-go.
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/directconnect"
	"github.com/aws/aws-sdk-go/service/directconnect/directconnectiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken
type _ request.Option

type DirectConnectActivities struct {
	client directconnectiface.DirectConnectAPI

	sessionFactory SessionFactory
}

func NewDirectConnectActivities(sess *session.Session, config ...*aws.Config) *DirectConnectActivities {
	client := directconnect.New(sess, config...)
	return &DirectConnectActivities{client: client}
}

func NewDirectConnectActivitiesWithSessionFactory(sessionFactory SessionFactory) *DirectConnectActivities {
	return &DirectConnectActivities{sessionFactory: sessionFactory}
}

func (a *DirectConnectActivities) getClient(ctx context.Context) (directconnectiface.DirectConnectAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, err
	}

	return directconnect.New(sess), nil
}

func (a *DirectConnectActivities) AcceptDirectConnectGatewayAssociationProposal(ctx context.Context, input *directconnect.AcceptDirectConnectGatewayAssociationProposalInput) (*directconnect.AcceptDirectConnectGatewayAssociationProposalOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AcceptDirectConnectGatewayAssociationProposalWithContext(ctx, input)
}

func (a *DirectConnectActivities) AllocateConnectionOnInterconnect(ctx context.Context, input *directconnect.AllocateConnectionOnInterconnectInput) (*directconnect.Connection, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AllocateConnectionOnInterconnectWithContext(ctx, input)
}

func (a *DirectConnectActivities) AllocateHostedConnection(ctx context.Context, input *directconnect.AllocateHostedConnectionInput) (*directconnect.Connection, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AllocateHostedConnectionWithContext(ctx, input)
}

func (a *DirectConnectActivities) AllocatePrivateVirtualInterface(ctx context.Context, input *directconnect.AllocatePrivateVirtualInterfaceInput) (*directconnect.VirtualInterface, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AllocatePrivateVirtualInterfaceWithContext(ctx, input)
}

func (a *DirectConnectActivities) AllocatePublicVirtualInterface(ctx context.Context, input *directconnect.AllocatePublicVirtualInterfaceInput) (*directconnect.VirtualInterface, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AllocatePublicVirtualInterfaceWithContext(ctx, input)
}

func (a *DirectConnectActivities) AllocateTransitVirtualInterface(ctx context.Context, input *directconnect.AllocateTransitVirtualInterfaceInput) (*directconnect.AllocateTransitVirtualInterfaceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AllocateTransitVirtualInterfaceWithContext(ctx, input)
}

func (a *DirectConnectActivities) AssociateConnectionWithLag(ctx context.Context, input *directconnect.AssociateConnectionWithLagInput) (*directconnect.Connection, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AssociateConnectionWithLagWithContext(ctx, input)
}

func (a *DirectConnectActivities) AssociateHostedConnection(ctx context.Context, input *directconnect.AssociateHostedConnectionInput) (*directconnect.Connection, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AssociateHostedConnectionWithContext(ctx, input)
}

func (a *DirectConnectActivities) AssociateVirtualInterface(ctx context.Context, input *directconnect.AssociateVirtualInterfaceInput) (*directconnect.VirtualInterface, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.AssociateVirtualInterfaceWithContext(ctx, input)
}

func (a *DirectConnectActivities) ConfirmConnection(ctx context.Context, input *directconnect.ConfirmConnectionInput) (*directconnect.ConfirmConnectionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ConfirmConnectionWithContext(ctx, input)
}

func (a *DirectConnectActivities) ConfirmPrivateVirtualInterface(ctx context.Context, input *directconnect.ConfirmPrivateVirtualInterfaceInput) (*directconnect.ConfirmPrivateVirtualInterfaceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ConfirmPrivateVirtualInterfaceWithContext(ctx, input)
}

func (a *DirectConnectActivities) ConfirmPublicVirtualInterface(ctx context.Context, input *directconnect.ConfirmPublicVirtualInterfaceInput) (*directconnect.ConfirmPublicVirtualInterfaceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ConfirmPublicVirtualInterfaceWithContext(ctx, input)
}

func (a *DirectConnectActivities) ConfirmTransitVirtualInterface(ctx context.Context, input *directconnect.ConfirmTransitVirtualInterfaceInput) (*directconnect.ConfirmTransitVirtualInterfaceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ConfirmTransitVirtualInterfaceWithContext(ctx, input)
}

func (a *DirectConnectActivities) CreateBGPPeer(ctx context.Context, input *directconnect.CreateBGPPeerInput) (*directconnect.CreateBGPPeerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateBGPPeerWithContext(ctx, input)
}

func (a *DirectConnectActivities) CreateConnection(ctx context.Context, input *directconnect.CreateConnectionInput) (*directconnect.Connection, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateConnectionWithContext(ctx, input)
}

func (a *DirectConnectActivities) CreateDirectConnectGateway(ctx context.Context, input *directconnect.CreateDirectConnectGatewayInput) (*directconnect.CreateDirectConnectGatewayOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateDirectConnectGatewayWithContext(ctx, input)
}

func (a *DirectConnectActivities) CreateDirectConnectGatewayAssociation(ctx context.Context, input *directconnect.CreateDirectConnectGatewayAssociationInput) (*directconnect.CreateDirectConnectGatewayAssociationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateDirectConnectGatewayAssociationWithContext(ctx, input)
}

func (a *DirectConnectActivities) CreateDirectConnectGatewayAssociationProposal(ctx context.Context, input *directconnect.CreateDirectConnectGatewayAssociationProposalInput) (*directconnect.CreateDirectConnectGatewayAssociationProposalOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateDirectConnectGatewayAssociationProposalWithContext(ctx, input)
}

func (a *DirectConnectActivities) CreateInterconnect(ctx context.Context, input *directconnect.CreateInterconnectInput) (*directconnect.Interconnect, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateInterconnectWithContext(ctx, input)
}

func (a *DirectConnectActivities) CreateLag(ctx context.Context, input *directconnect.CreateLagInput) (*directconnect.Lag, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateLagWithContext(ctx, input)
}

func (a *DirectConnectActivities) CreatePrivateVirtualInterface(ctx context.Context, input *directconnect.CreatePrivateVirtualInterfaceInput) (*directconnect.VirtualInterface, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreatePrivateVirtualInterfaceWithContext(ctx, input)
}

func (a *DirectConnectActivities) CreatePublicVirtualInterface(ctx context.Context, input *directconnect.CreatePublicVirtualInterfaceInput) (*directconnect.VirtualInterface, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreatePublicVirtualInterfaceWithContext(ctx, input)
}

func (a *DirectConnectActivities) CreateTransitVirtualInterface(ctx context.Context, input *directconnect.CreateTransitVirtualInterfaceInput) (*directconnect.CreateTransitVirtualInterfaceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.CreateTransitVirtualInterfaceWithContext(ctx, input)
}

func (a *DirectConnectActivities) DeleteBGPPeer(ctx context.Context, input *directconnect.DeleteBGPPeerInput) (*directconnect.DeleteBGPPeerOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteBGPPeerWithContext(ctx, input)
}

func (a *DirectConnectActivities) DeleteConnection(ctx context.Context, input *directconnect.DeleteConnectionInput) (*directconnect.Connection, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteConnectionWithContext(ctx, input)
}

func (a *DirectConnectActivities) DeleteDirectConnectGateway(ctx context.Context, input *directconnect.DeleteDirectConnectGatewayInput) (*directconnect.DeleteDirectConnectGatewayOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteDirectConnectGatewayWithContext(ctx, input)
}

func (a *DirectConnectActivities) DeleteDirectConnectGatewayAssociation(ctx context.Context, input *directconnect.DeleteDirectConnectGatewayAssociationInput) (*directconnect.DeleteDirectConnectGatewayAssociationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteDirectConnectGatewayAssociationWithContext(ctx, input)
}

func (a *DirectConnectActivities) DeleteDirectConnectGatewayAssociationProposal(ctx context.Context, input *directconnect.DeleteDirectConnectGatewayAssociationProposalInput) (*directconnect.DeleteDirectConnectGatewayAssociationProposalOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteDirectConnectGatewayAssociationProposalWithContext(ctx, input)
}

func (a *DirectConnectActivities) DeleteInterconnect(ctx context.Context, input *directconnect.DeleteInterconnectInput) (*directconnect.DeleteInterconnectOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteInterconnectWithContext(ctx, input)
}

func (a *DirectConnectActivities) DeleteLag(ctx context.Context, input *directconnect.DeleteLagInput) (*directconnect.Lag, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteLagWithContext(ctx, input)
}

func (a *DirectConnectActivities) DeleteVirtualInterface(ctx context.Context, input *directconnect.DeleteVirtualInterfaceInput) (*directconnect.DeleteVirtualInterfaceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DeleteVirtualInterfaceWithContext(ctx, input)
}

func (a *DirectConnectActivities) DescribeConnectionLoa(ctx context.Context, input *directconnect.DescribeConnectionLoaInput) (*directconnect.DescribeConnectionLoaOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeConnectionLoaWithContext(ctx, input)
}

func (a *DirectConnectActivities) DescribeConnections(ctx context.Context, input *directconnect.DescribeConnectionsInput) (*directconnect.Connections, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeConnectionsWithContext(ctx, input)
}

func (a *DirectConnectActivities) DescribeConnectionsOnInterconnect(ctx context.Context, input *directconnect.DescribeConnectionsOnInterconnectInput) (*directconnect.Connections, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeConnectionsOnInterconnectWithContext(ctx, input)
}

func (a *DirectConnectActivities) DescribeDirectConnectGatewayAssociationProposals(ctx context.Context, input *directconnect.DescribeDirectConnectGatewayAssociationProposalsInput) (*directconnect.DescribeDirectConnectGatewayAssociationProposalsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeDirectConnectGatewayAssociationProposalsWithContext(ctx, input)
}

func (a *DirectConnectActivities) DescribeDirectConnectGatewayAssociations(ctx context.Context, input *directconnect.DescribeDirectConnectGatewayAssociationsInput) (*directconnect.DescribeDirectConnectGatewayAssociationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeDirectConnectGatewayAssociationsWithContext(ctx, input)
}

func (a *DirectConnectActivities) DescribeDirectConnectGatewayAttachments(ctx context.Context, input *directconnect.DescribeDirectConnectGatewayAttachmentsInput) (*directconnect.DescribeDirectConnectGatewayAttachmentsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeDirectConnectGatewayAttachmentsWithContext(ctx, input)
}

func (a *DirectConnectActivities) DescribeDirectConnectGateways(ctx context.Context, input *directconnect.DescribeDirectConnectGatewaysInput) (*directconnect.DescribeDirectConnectGatewaysOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeDirectConnectGatewaysWithContext(ctx, input)
}

func (a *DirectConnectActivities) DescribeHostedConnections(ctx context.Context, input *directconnect.DescribeHostedConnectionsInput) (*directconnect.Connections, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeHostedConnectionsWithContext(ctx, input)
}

func (a *DirectConnectActivities) DescribeInterconnectLoa(ctx context.Context, input *directconnect.DescribeInterconnectLoaInput) (*directconnect.DescribeInterconnectLoaOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeInterconnectLoaWithContext(ctx, input)
}

func (a *DirectConnectActivities) DescribeInterconnects(ctx context.Context, input *directconnect.DescribeInterconnectsInput) (*directconnect.DescribeInterconnectsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeInterconnectsWithContext(ctx, input)
}

func (a *DirectConnectActivities) DescribeLags(ctx context.Context, input *directconnect.DescribeLagsInput) (*directconnect.DescribeLagsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeLagsWithContext(ctx, input)
}

func (a *DirectConnectActivities) DescribeLoa(ctx context.Context, input *directconnect.DescribeLoaInput) (*directconnect.Loa, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeLoaWithContext(ctx, input)
}

func (a *DirectConnectActivities) DescribeLocations(ctx context.Context, input *directconnect.DescribeLocationsInput) (*directconnect.DescribeLocationsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeLocationsWithContext(ctx, input)
}

func (a *DirectConnectActivities) DescribeTags(ctx context.Context, input *directconnect.DescribeTagsInput) (*directconnect.DescribeTagsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeTagsWithContext(ctx, input)
}

func (a *DirectConnectActivities) DescribeVirtualGateways(ctx context.Context, input *directconnect.DescribeVirtualGatewaysInput) (*directconnect.DescribeVirtualGatewaysOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeVirtualGatewaysWithContext(ctx, input)
}

func (a *DirectConnectActivities) DescribeVirtualInterfaces(ctx context.Context, input *directconnect.DescribeVirtualInterfacesInput) (*directconnect.DescribeVirtualInterfacesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DescribeVirtualInterfacesWithContext(ctx, input)
}

func (a *DirectConnectActivities) DisassociateConnectionFromLag(ctx context.Context, input *directconnect.DisassociateConnectionFromLagInput) (*directconnect.Connection, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.DisassociateConnectionFromLagWithContext(ctx, input)
}

func (a *DirectConnectActivities) ListVirtualInterfaceTestHistory(ctx context.Context, input *directconnect.ListVirtualInterfaceTestHistoryInput) (*directconnect.ListVirtualInterfaceTestHistoryOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.ListVirtualInterfaceTestHistoryWithContext(ctx, input)
}

func (a *DirectConnectActivities) StartBgpFailoverTest(ctx context.Context, input *directconnect.StartBgpFailoverTestInput) (*directconnect.StartBgpFailoverTestOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StartBgpFailoverTestWithContext(ctx, input)
}

func (a *DirectConnectActivities) StopBgpFailoverTest(ctx context.Context, input *directconnect.StopBgpFailoverTestInput) (*directconnect.StopBgpFailoverTestOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.StopBgpFailoverTestWithContext(ctx, input)
}

func (a *DirectConnectActivities) TagResource(ctx context.Context, input *directconnect.TagResourceInput) (*directconnect.TagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.TagResourceWithContext(ctx, input)
}

func (a *DirectConnectActivities) UntagResource(ctx context.Context, input *directconnect.UntagResourceInput) (*directconnect.UntagResourceOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UntagResourceWithContext(ctx, input)
}

func (a *DirectConnectActivities) UpdateDirectConnectGatewayAssociation(ctx context.Context, input *directconnect.UpdateDirectConnectGatewayAssociationInput) (*directconnect.UpdateDirectConnectGatewayAssociationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateDirectConnectGatewayAssociationWithContext(ctx, input)
}

func (a *DirectConnectActivities) UpdateLag(ctx context.Context, input *directconnect.UpdateLagInput) (*directconnect.Lag, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateLagWithContext(ctx, input)
}

func (a *DirectConnectActivities) UpdateVirtualInterfaceAttributes(ctx context.Context, input *directconnect.UpdateVirtualInterfaceAttributesInput) (*directconnect.UpdateVirtualInterfaceAttributesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, err
	}
	return client.UpdateVirtualInterfaceAttributesWithContext(ctx, input)
}
