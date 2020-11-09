// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package ramstub

import (
	"github.com/aws/aws-sdk-go/service/ram"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	AcceptResourceShareInvitation(ctx workflow.Context, input *ram.AcceptResourceShareInvitationInput) (*ram.AcceptResourceShareInvitationOutput, error)
	AcceptResourceShareInvitationAsync(ctx workflow.Context, input *ram.AcceptResourceShareInvitationInput) *AcceptResourceShareInvitationFuture

	AssociateResourceShare(ctx workflow.Context, input *ram.AssociateResourceShareInput) (*ram.AssociateResourceShareOutput, error)
	AssociateResourceShareAsync(ctx workflow.Context, input *ram.AssociateResourceShareInput) *AssociateResourceShareFuture

	AssociateResourceSharePermission(ctx workflow.Context, input *ram.AssociateResourceSharePermissionInput) (*ram.AssociateResourceSharePermissionOutput, error)
	AssociateResourceSharePermissionAsync(ctx workflow.Context, input *ram.AssociateResourceSharePermissionInput) *AssociateResourceSharePermissionFuture

	CreateResourceShare(ctx workflow.Context, input *ram.CreateResourceShareInput) (*ram.CreateResourceShareOutput, error)
	CreateResourceShareAsync(ctx workflow.Context, input *ram.CreateResourceShareInput) *CreateResourceShareFuture

	DeleteResourceShare(ctx workflow.Context, input *ram.DeleteResourceShareInput) (*ram.DeleteResourceShareOutput, error)
	DeleteResourceShareAsync(ctx workflow.Context, input *ram.DeleteResourceShareInput) *DeleteResourceShareFuture

	DisassociateResourceShare(ctx workflow.Context, input *ram.DisassociateResourceShareInput) (*ram.DisassociateResourceShareOutput, error)
	DisassociateResourceShareAsync(ctx workflow.Context, input *ram.DisassociateResourceShareInput) *DisassociateResourceShareFuture

	DisassociateResourceSharePermission(ctx workflow.Context, input *ram.DisassociateResourceSharePermissionInput) (*ram.DisassociateResourceSharePermissionOutput, error)
	DisassociateResourceSharePermissionAsync(ctx workflow.Context, input *ram.DisassociateResourceSharePermissionInput) *DisassociateResourceSharePermissionFuture

	EnableSharingWithAwsOrganization(ctx workflow.Context, input *ram.EnableSharingWithAwsOrganizationInput) (*ram.EnableSharingWithAwsOrganizationOutput, error)
	EnableSharingWithAwsOrganizationAsync(ctx workflow.Context, input *ram.EnableSharingWithAwsOrganizationInput) *EnableSharingWithAwsOrganizationFuture

	GetPermission(ctx workflow.Context, input *ram.GetPermissionInput) (*ram.GetPermissionOutput, error)
	GetPermissionAsync(ctx workflow.Context, input *ram.GetPermissionInput) *GetPermissionFuture

	GetResourcePolicies(ctx workflow.Context, input *ram.GetResourcePoliciesInput) (*ram.GetResourcePoliciesOutput, error)
	GetResourcePoliciesAsync(ctx workflow.Context, input *ram.GetResourcePoliciesInput) *GetResourcePoliciesFuture

	GetResourceShareAssociations(ctx workflow.Context, input *ram.GetResourceShareAssociationsInput) (*ram.GetResourceShareAssociationsOutput, error)
	GetResourceShareAssociationsAsync(ctx workflow.Context, input *ram.GetResourceShareAssociationsInput) *GetResourceShareAssociationsFuture

	GetResourceShareInvitations(ctx workflow.Context, input *ram.GetResourceShareInvitationsInput) (*ram.GetResourceShareInvitationsOutput, error)
	GetResourceShareInvitationsAsync(ctx workflow.Context, input *ram.GetResourceShareInvitationsInput) *GetResourceShareInvitationsFuture

	GetResourceShares(ctx workflow.Context, input *ram.GetResourceSharesInput) (*ram.GetResourceSharesOutput, error)
	GetResourceSharesAsync(ctx workflow.Context, input *ram.GetResourceSharesInput) *GetResourceSharesFuture

	ListPendingInvitationResources(ctx workflow.Context, input *ram.ListPendingInvitationResourcesInput) (*ram.ListPendingInvitationResourcesOutput, error)
	ListPendingInvitationResourcesAsync(ctx workflow.Context, input *ram.ListPendingInvitationResourcesInput) *ListPendingInvitationResourcesFuture

	ListPermissions(ctx workflow.Context, input *ram.ListPermissionsInput) (*ram.ListPermissionsOutput, error)
	ListPermissionsAsync(ctx workflow.Context, input *ram.ListPermissionsInput) *ListPermissionsFuture

	ListPrincipals(ctx workflow.Context, input *ram.ListPrincipalsInput) (*ram.ListPrincipalsOutput, error)
	ListPrincipalsAsync(ctx workflow.Context, input *ram.ListPrincipalsInput) *ListPrincipalsFuture

	ListResourceSharePermissions(ctx workflow.Context, input *ram.ListResourceSharePermissionsInput) (*ram.ListResourceSharePermissionsOutput, error)
	ListResourceSharePermissionsAsync(ctx workflow.Context, input *ram.ListResourceSharePermissionsInput) *ListResourceSharePermissionsFuture

	ListResourceTypes(ctx workflow.Context, input *ram.ListResourceTypesInput) (*ram.ListResourceTypesOutput, error)
	ListResourceTypesAsync(ctx workflow.Context, input *ram.ListResourceTypesInput) *ListResourceTypesFuture

	ListResources(ctx workflow.Context, input *ram.ListResourcesInput) (*ram.ListResourcesOutput, error)
	ListResourcesAsync(ctx workflow.Context, input *ram.ListResourcesInput) *ListResourcesFuture

	PromoteResourceShareCreatedFromPolicy(ctx workflow.Context, input *ram.PromoteResourceShareCreatedFromPolicyInput) (*ram.PromoteResourceShareCreatedFromPolicyOutput, error)
	PromoteResourceShareCreatedFromPolicyAsync(ctx workflow.Context, input *ram.PromoteResourceShareCreatedFromPolicyInput) *PromoteResourceShareCreatedFromPolicyFuture

	RejectResourceShareInvitation(ctx workflow.Context, input *ram.RejectResourceShareInvitationInput) (*ram.RejectResourceShareInvitationOutput, error)
	RejectResourceShareInvitationAsync(ctx workflow.Context, input *ram.RejectResourceShareInvitationInput) *RejectResourceShareInvitationFuture

	TagResource(ctx workflow.Context, input *ram.TagResourceInput) (*ram.TagResourceOutput, error)
	TagResourceAsync(ctx workflow.Context, input *ram.TagResourceInput) *TagResourceFuture

	UntagResource(ctx workflow.Context, input *ram.UntagResourceInput) (*ram.UntagResourceOutput, error)
	UntagResourceAsync(ctx workflow.Context, input *ram.UntagResourceInput) *UntagResourceFuture

	UpdateResourceShare(ctx workflow.Context, input *ram.UpdateResourceShareInput) (*ram.UpdateResourceShareOutput, error)
	UpdateResourceShareAsync(ctx workflow.Context, input *ram.UpdateResourceShareInput) *UpdateResourceShareFuture
}

func NewClient() Client {
	return &stub{}
}