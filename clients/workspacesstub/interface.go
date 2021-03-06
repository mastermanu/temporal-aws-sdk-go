// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package workspacesstub

import (
	"github.com/aws/aws-sdk-go/service/workspaces"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	AssociateConnectionAlias(ctx workflow.Context, input *workspaces.AssociateConnectionAliasInput) (*workspaces.AssociateConnectionAliasOutput, error)
	AssociateConnectionAliasAsync(ctx workflow.Context, input *workspaces.AssociateConnectionAliasInput) *AssociateConnectionAliasFuture

	AssociateIpGroups(ctx workflow.Context, input *workspaces.AssociateIpGroupsInput) (*workspaces.AssociateIpGroupsOutput, error)
	AssociateIpGroupsAsync(ctx workflow.Context, input *workspaces.AssociateIpGroupsInput) *AssociateIpGroupsFuture

	AuthorizeIpRules(ctx workflow.Context, input *workspaces.AuthorizeIpRulesInput) (*workspaces.AuthorizeIpRulesOutput, error)
	AuthorizeIpRulesAsync(ctx workflow.Context, input *workspaces.AuthorizeIpRulesInput) *AuthorizeIpRulesFuture

	CopyWorkspaceImage(ctx workflow.Context, input *workspaces.CopyWorkspaceImageInput) (*workspaces.CopyWorkspaceImageOutput, error)
	CopyWorkspaceImageAsync(ctx workflow.Context, input *workspaces.CopyWorkspaceImageInput) *CopyWorkspaceImageFuture

	CreateConnectionAlias(ctx workflow.Context, input *workspaces.CreateConnectionAliasInput) (*workspaces.CreateConnectionAliasOutput, error)
	CreateConnectionAliasAsync(ctx workflow.Context, input *workspaces.CreateConnectionAliasInput) *CreateConnectionAliasFuture

	CreateIpGroup(ctx workflow.Context, input *workspaces.CreateIpGroupInput) (*workspaces.CreateIpGroupOutput, error)
	CreateIpGroupAsync(ctx workflow.Context, input *workspaces.CreateIpGroupInput) *CreateIpGroupFuture

	CreateTags(ctx workflow.Context, input *workspaces.CreateTagsInput) (*workspaces.CreateTagsOutput, error)
	CreateTagsAsync(ctx workflow.Context, input *workspaces.CreateTagsInput) *CreateTagsFuture

	CreateWorkspaces(ctx workflow.Context, input *workspaces.CreateWorkspacesInput) (*workspaces.CreateWorkspacesOutput, error)
	CreateWorkspacesAsync(ctx workflow.Context, input *workspaces.CreateWorkspacesInput) *CreateWorkspacesFuture

	DeleteConnectionAlias(ctx workflow.Context, input *workspaces.DeleteConnectionAliasInput) (*workspaces.DeleteConnectionAliasOutput, error)
	DeleteConnectionAliasAsync(ctx workflow.Context, input *workspaces.DeleteConnectionAliasInput) *DeleteConnectionAliasFuture

	DeleteIpGroup(ctx workflow.Context, input *workspaces.DeleteIpGroupInput) (*workspaces.DeleteIpGroupOutput, error)
	DeleteIpGroupAsync(ctx workflow.Context, input *workspaces.DeleteIpGroupInput) *DeleteIpGroupFuture

	DeleteTags(ctx workflow.Context, input *workspaces.DeleteTagsInput) (*workspaces.DeleteTagsOutput, error)
	DeleteTagsAsync(ctx workflow.Context, input *workspaces.DeleteTagsInput) *DeleteTagsFuture

	DeleteWorkspaceImage(ctx workflow.Context, input *workspaces.DeleteWorkspaceImageInput) (*workspaces.DeleteWorkspaceImageOutput, error)
	DeleteWorkspaceImageAsync(ctx workflow.Context, input *workspaces.DeleteWorkspaceImageInput) *DeleteWorkspaceImageFuture

	DeregisterWorkspaceDirectory(ctx workflow.Context, input *workspaces.DeregisterWorkspaceDirectoryInput) (*workspaces.DeregisterWorkspaceDirectoryOutput, error)
	DeregisterWorkspaceDirectoryAsync(ctx workflow.Context, input *workspaces.DeregisterWorkspaceDirectoryInput) *DeregisterWorkspaceDirectoryFuture

	DescribeAccount(ctx workflow.Context, input *workspaces.DescribeAccountInput) (*workspaces.DescribeAccountOutput, error)
	DescribeAccountAsync(ctx workflow.Context, input *workspaces.DescribeAccountInput) *DescribeAccountFuture

	DescribeAccountModifications(ctx workflow.Context, input *workspaces.DescribeAccountModificationsInput) (*workspaces.DescribeAccountModificationsOutput, error)
	DescribeAccountModificationsAsync(ctx workflow.Context, input *workspaces.DescribeAccountModificationsInput) *DescribeAccountModificationsFuture

	DescribeClientProperties(ctx workflow.Context, input *workspaces.DescribeClientPropertiesInput) (*workspaces.DescribeClientPropertiesOutput, error)
	DescribeClientPropertiesAsync(ctx workflow.Context, input *workspaces.DescribeClientPropertiesInput) *DescribeClientPropertiesFuture

	DescribeConnectionAliasPermissions(ctx workflow.Context, input *workspaces.DescribeConnectionAliasPermissionsInput) (*workspaces.DescribeConnectionAliasPermissionsOutput, error)
	DescribeConnectionAliasPermissionsAsync(ctx workflow.Context, input *workspaces.DescribeConnectionAliasPermissionsInput) *DescribeConnectionAliasPermissionsFuture

	DescribeConnectionAliases(ctx workflow.Context, input *workspaces.DescribeConnectionAliasesInput) (*workspaces.DescribeConnectionAliasesOutput, error)
	DescribeConnectionAliasesAsync(ctx workflow.Context, input *workspaces.DescribeConnectionAliasesInput) *DescribeConnectionAliasesFuture

	DescribeIpGroups(ctx workflow.Context, input *workspaces.DescribeIpGroupsInput) (*workspaces.DescribeIpGroupsOutput, error)
	DescribeIpGroupsAsync(ctx workflow.Context, input *workspaces.DescribeIpGroupsInput) *DescribeIpGroupsFuture

	DescribeTags(ctx workflow.Context, input *workspaces.DescribeTagsInput) (*workspaces.DescribeTagsOutput, error)
	DescribeTagsAsync(ctx workflow.Context, input *workspaces.DescribeTagsInput) *DescribeTagsFuture

	DescribeWorkspaceBundles(ctx workflow.Context, input *workspaces.DescribeWorkspaceBundlesInput) (*workspaces.DescribeWorkspaceBundlesOutput, error)
	DescribeWorkspaceBundlesAsync(ctx workflow.Context, input *workspaces.DescribeWorkspaceBundlesInput) *DescribeWorkspaceBundlesFuture

	DescribeWorkspaceDirectories(ctx workflow.Context, input *workspaces.DescribeWorkspaceDirectoriesInput) (*workspaces.DescribeWorkspaceDirectoriesOutput, error)
	DescribeWorkspaceDirectoriesAsync(ctx workflow.Context, input *workspaces.DescribeWorkspaceDirectoriesInput) *DescribeWorkspaceDirectoriesFuture

	DescribeWorkspaceImagePermissions(ctx workflow.Context, input *workspaces.DescribeWorkspaceImagePermissionsInput) (*workspaces.DescribeWorkspaceImagePermissionsOutput, error)
	DescribeWorkspaceImagePermissionsAsync(ctx workflow.Context, input *workspaces.DescribeWorkspaceImagePermissionsInput) *DescribeWorkspaceImagePermissionsFuture

	DescribeWorkspaceImages(ctx workflow.Context, input *workspaces.DescribeWorkspaceImagesInput) (*workspaces.DescribeWorkspaceImagesOutput, error)
	DescribeWorkspaceImagesAsync(ctx workflow.Context, input *workspaces.DescribeWorkspaceImagesInput) *DescribeWorkspaceImagesFuture

	DescribeWorkspaceSnapshots(ctx workflow.Context, input *workspaces.DescribeWorkspaceSnapshotsInput) (*workspaces.DescribeWorkspaceSnapshotsOutput, error)
	DescribeWorkspaceSnapshotsAsync(ctx workflow.Context, input *workspaces.DescribeWorkspaceSnapshotsInput) *DescribeWorkspaceSnapshotsFuture

	DescribeWorkspaces(ctx workflow.Context, input *workspaces.DescribeWorkspacesInput) (*workspaces.DescribeWorkspacesOutput, error)
	DescribeWorkspacesAsync(ctx workflow.Context, input *workspaces.DescribeWorkspacesInput) *DescribeWorkspacesFuture

	DescribeWorkspacesConnectionStatus(ctx workflow.Context, input *workspaces.DescribeWorkspacesConnectionStatusInput) (*workspaces.DescribeWorkspacesConnectionStatusOutput, error)
	DescribeWorkspacesConnectionStatusAsync(ctx workflow.Context, input *workspaces.DescribeWorkspacesConnectionStatusInput) *DescribeWorkspacesConnectionStatusFuture

	DisassociateConnectionAlias(ctx workflow.Context, input *workspaces.DisassociateConnectionAliasInput) (*workspaces.DisassociateConnectionAliasOutput, error)
	DisassociateConnectionAliasAsync(ctx workflow.Context, input *workspaces.DisassociateConnectionAliasInput) *DisassociateConnectionAliasFuture

	DisassociateIpGroups(ctx workflow.Context, input *workspaces.DisassociateIpGroupsInput) (*workspaces.DisassociateIpGroupsOutput, error)
	DisassociateIpGroupsAsync(ctx workflow.Context, input *workspaces.DisassociateIpGroupsInput) *DisassociateIpGroupsFuture

	ImportWorkspaceImage(ctx workflow.Context, input *workspaces.ImportWorkspaceImageInput) (*workspaces.ImportWorkspaceImageOutput, error)
	ImportWorkspaceImageAsync(ctx workflow.Context, input *workspaces.ImportWorkspaceImageInput) *ImportWorkspaceImageFuture

	ListAvailableManagementCidrRanges(ctx workflow.Context, input *workspaces.ListAvailableManagementCidrRangesInput) (*workspaces.ListAvailableManagementCidrRangesOutput, error)
	ListAvailableManagementCidrRangesAsync(ctx workflow.Context, input *workspaces.ListAvailableManagementCidrRangesInput) *ListAvailableManagementCidrRangesFuture

	MigrateWorkspace(ctx workflow.Context, input *workspaces.MigrateWorkspaceInput) (*workspaces.MigrateWorkspaceOutput, error)
	MigrateWorkspaceAsync(ctx workflow.Context, input *workspaces.MigrateWorkspaceInput) *MigrateWorkspaceFuture

	ModifyAccount(ctx workflow.Context, input *workspaces.ModifyAccountInput) (*workspaces.ModifyAccountOutput, error)
	ModifyAccountAsync(ctx workflow.Context, input *workspaces.ModifyAccountInput) *ModifyAccountFuture

	ModifyClientProperties(ctx workflow.Context, input *workspaces.ModifyClientPropertiesInput) (*workspaces.ModifyClientPropertiesOutput, error)
	ModifyClientPropertiesAsync(ctx workflow.Context, input *workspaces.ModifyClientPropertiesInput) *ModifyClientPropertiesFuture

	ModifySelfservicePermissions(ctx workflow.Context, input *workspaces.ModifySelfservicePermissionsInput) (*workspaces.ModifySelfservicePermissionsOutput, error)
	ModifySelfservicePermissionsAsync(ctx workflow.Context, input *workspaces.ModifySelfservicePermissionsInput) *ModifySelfservicePermissionsFuture

	ModifyWorkspaceAccessProperties(ctx workflow.Context, input *workspaces.ModifyWorkspaceAccessPropertiesInput) (*workspaces.ModifyWorkspaceAccessPropertiesOutput, error)
	ModifyWorkspaceAccessPropertiesAsync(ctx workflow.Context, input *workspaces.ModifyWorkspaceAccessPropertiesInput) *ModifyWorkspaceAccessPropertiesFuture

	ModifyWorkspaceCreationProperties(ctx workflow.Context, input *workspaces.ModifyWorkspaceCreationPropertiesInput) (*workspaces.ModifyWorkspaceCreationPropertiesOutput, error)
	ModifyWorkspaceCreationPropertiesAsync(ctx workflow.Context, input *workspaces.ModifyWorkspaceCreationPropertiesInput) *ModifyWorkspaceCreationPropertiesFuture

	ModifyWorkspaceProperties(ctx workflow.Context, input *workspaces.ModifyWorkspacePropertiesInput) (*workspaces.ModifyWorkspacePropertiesOutput, error)
	ModifyWorkspacePropertiesAsync(ctx workflow.Context, input *workspaces.ModifyWorkspacePropertiesInput) *ModifyWorkspacePropertiesFuture

	ModifyWorkspaceState(ctx workflow.Context, input *workspaces.ModifyWorkspaceStateInput) (*workspaces.ModifyWorkspaceStateOutput, error)
	ModifyWorkspaceStateAsync(ctx workflow.Context, input *workspaces.ModifyWorkspaceStateInput) *ModifyWorkspaceStateFuture

	RebootWorkspaces(ctx workflow.Context, input *workspaces.RebootWorkspacesInput) (*workspaces.RebootWorkspacesOutput, error)
	RebootWorkspacesAsync(ctx workflow.Context, input *workspaces.RebootWorkspacesInput) *RebootWorkspacesFuture

	RebuildWorkspaces(ctx workflow.Context, input *workspaces.RebuildWorkspacesInput) (*workspaces.RebuildWorkspacesOutput, error)
	RebuildWorkspacesAsync(ctx workflow.Context, input *workspaces.RebuildWorkspacesInput) *RebuildWorkspacesFuture

	RegisterWorkspaceDirectory(ctx workflow.Context, input *workspaces.RegisterWorkspaceDirectoryInput) (*workspaces.RegisterWorkspaceDirectoryOutput, error)
	RegisterWorkspaceDirectoryAsync(ctx workflow.Context, input *workspaces.RegisterWorkspaceDirectoryInput) *RegisterWorkspaceDirectoryFuture

	RestoreWorkspace(ctx workflow.Context, input *workspaces.RestoreWorkspaceInput) (*workspaces.RestoreWorkspaceOutput, error)
	RestoreWorkspaceAsync(ctx workflow.Context, input *workspaces.RestoreWorkspaceInput) *RestoreWorkspaceFuture

	RevokeIpRules(ctx workflow.Context, input *workspaces.RevokeIpRulesInput) (*workspaces.RevokeIpRulesOutput, error)
	RevokeIpRulesAsync(ctx workflow.Context, input *workspaces.RevokeIpRulesInput) *RevokeIpRulesFuture

	StartWorkspaces(ctx workflow.Context, input *workspaces.StartWorkspacesInput) (*workspaces.StartWorkspacesOutput, error)
	StartWorkspacesAsync(ctx workflow.Context, input *workspaces.StartWorkspacesInput) *StartWorkspacesFuture

	StopWorkspaces(ctx workflow.Context, input *workspaces.StopWorkspacesInput) (*workspaces.StopWorkspacesOutput, error)
	StopWorkspacesAsync(ctx workflow.Context, input *workspaces.StopWorkspacesInput) *StopWorkspacesFuture

	TerminateWorkspaces(ctx workflow.Context, input *workspaces.TerminateWorkspacesInput) (*workspaces.TerminateWorkspacesOutput, error)
	TerminateWorkspacesAsync(ctx workflow.Context, input *workspaces.TerminateWorkspacesInput) *TerminateWorkspacesFuture

	UpdateConnectionAliasPermission(ctx workflow.Context, input *workspaces.UpdateConnectionAliasPermissionInput) (*workspaces.UpdateConnectionAliasPermissionOutput, error)
	UpdateConnectionAliasPermissionAsync(ctx workflow.Context, input *workspaces.UpdateConnectionAliasPermissionInput) *UpdateConnectionAliasPermissionFuture

	UpdateRulesOfIpGroup(ctx workflow.Context, input *workspaces.UpdateRulesOfIpGroupInput) (*workspaces.UpdateRulesOfIpGroupOutput, error)
	UpdateRulesOfIpGroupAsync(ctx workflow.Context, input *workspaces.UpdateRulesOfIpGroupInput) *UpdateRulesOfIpGroupFuture

	UpdateWorkspaceImagePermission(ctx workflow.Context, input *workspaces.UpdateWorkspaceImagePermissionInput) (*workspaces.UpdateWorkspaceImagePermissionOutput, error)
	UpdateWorkspaceImagePermissionAsync(ctx workflow.Context, input *workspaces.UpdateWorkspaceImagePermissionInput) *UpdateWorkspaceImagePermissionFuture
}

func NewClient() Client {
	return &stub{}
}
