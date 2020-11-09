// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package fmsstub

import (
	"github.com/aws/aws-sdk-go/service/fms"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type AssociateAdminAccountFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AssociateAdminAccountFuture) Get(ctx workflow.Context) (*fms.AssociateAdminAccountOutput, error) {
	var output fms.AssociateAdminAccountOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteAppsListFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteAppsListFuture) Get(ctx workflow.Context) (*fms.DeleteAppsListOutput, error) {
	var output fms.DeleteAppsListOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteNotificationChannelFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteNotificationChannelFuture) Get(ctx workflow.Context) (*fms.DeleteNotificationChannelOutput, error) {
	var output fms.DeleteNotificationChannelOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeletePolicyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeletePolicyFuture) Get(ctx workflow.Context) (*fms.DeletePolicyOutput, error) {
	var output fms.DeletePolicyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteProtocolsListFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteProtocolsListFuture) Get(ctx workflow.Context) (*fms.DeleteProtocolsListOutput, error) {
	var output fms.DeleteProtocolsListOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DisassociateAdminAccountFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DisassociateAdminAccountFuture) Get(ctx workflow.Context) (*fms.DisassociateAdminAccountOutput, error) {
	var output fms.DisassociateAdminAccountOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetAdminAccountFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetAdminAccountFuture) Get(ctx workflow.Context) (*fms.GetAdminAccountOutput, error) {
	var output fms.GetAdminAccountOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetAppsListFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetAppsListFuture) Get(ctx workflow.Context) (*fms.GetAppsListOutput, error) {
	var output fms.GetAppsListOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetComplianceDetailFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetComplianceDetailFuture) Get(ctx workflow.Context) (*fms.GetComplianceDetailOutput, error) {
	var output fms.GetComplianceDetailOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetNotificationChannelFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetNotificationChannelFuture) Get(ctx workflow.Context) (*fms.GetNotificationChannelOutput, error) {
	var output fms.GetNotificationChannelOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetPolicyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetPolicyFuture) Get(ctx workflow.Context) (*fms.GetPolicyOutput, error) {
	var output fms.GetPolicyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetProtectionStatusFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetProtectionStatusFuture) Get(ctx workflow.Context) (*fms.GetProtectionStatusOutput, error) {
	var output fms.GetProtectionStatusOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetProtocolsListFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetProtocolsListFuture) Get(ctx workflow.Context) (*fms.GetProtocolsListOutput, error) {
	var output fms.GetProtocolsListOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type GetViolationDetailsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *GetViolationDetailsFuture) Get(ctx workflow.Context) (*fms.GetViolationDetailsOutput, error) {
	var output fms.GetViolationDetailsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListAppsListsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListAppsListsFuture) Get(ctx workflow.Context) (*fms.ListAppsListsOutput, error) {
	var output fms.ListAppsListsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListComplianceStatusFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListComplianceStatusFuture) Get(ctx workflow.Context) (*fms.ListComplianceStatusOutput, error) {
	var output fms.ListComplianceStatusOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListMemberAccountsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListMemberAccountsFuture) Get(ctx workflow.Context) (*fms.ListMemberAccountsOutput, error) {
	var output fms.ListMemberAccountsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListPoliciesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListPoliciesFuture) Get(ctx workflow.Context) (*fms.ListPoliciesOutput, error) {
	var output fms.ListPoliciesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListProtocolsListsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListProtocolsListsFuture) Get(ctx workflow.Context) (*fms.ListProtocolsListsOutput, error) {
	var output fms.ListProtocolsListsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ListTagsForResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ListTagsForResourceFuture) Get(ctx workflow.Context) (*fms.ListTagsForResourceOutput, error) {
	var output fms.ListTagsForResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type PutAppsListFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *PutAppsListFuture) Get(ctx workflow.Context) (*fms.PutAppsListOutput, error) {
	var output fms.PutAppsListOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type PutNotificationChannelFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *PutNotificationChannelFuture) Get(ctx workflow.Context) (*fms.PutNotificationChannelOutput, error) {
	var output fms.PutNotificationChannelOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type PutPolicyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *PutPolicyFuture) Get(ctx workflow.Context) (*fms.PutPolicyOutput, error) {
	var output fms.PutPolicyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type PutProtocolsListFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *PutProtocolsListFuture) Get(ctx workflow.Context) (*fms.PutProtocolsListOutput, error) {
	var output fms.PutProtocolsListOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type TagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *TagResourceFuture) Get(ctx workflow.Context) (*fms.TagResourceOutput, error) {
	var output fms.TagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type UntagResourceFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *UntagResourceFuture) Get(ctx workflow.Context) (*fms.UntagResourceOutput, error) {
	var output fms.UntagResourceOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) AssociateAdminAccount(ctx workflow.Context, input *fms.AssociateAdminAccountInput) (*fms.AssociateAdminAccountOutput, error) {
	var output fms.AssociateAdminAccountOutput
	err := workflow.ExecuteActivity(ctx, "aws.fms.AssociateAdminAccount", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) AssociateAdminAccountAsync(ctx workflow.Context, input *fms.AssociateAdminAccountInput) *AssociateAdminAccountFuture {
	future := workflow.ExecuteActivity(ctx, "aws.fms.AssociateAdminAccount", input)
	return &AssociateAdminAccountFuture{Future: future}
}

func (a *stub) DeleteAppsList(ctx workflow.Context, input *fms.DeleteAppsListInput) (*fms.DeleteAppsListOutput, error) {
	var output fms.DeleteAppsListOutput
	err := workflow.ExecuteActivity(ctx, "aws.fms.DeleteAppsList", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteAppsListAsync(ctx workflow.Context, input *fms.DeleteAppsListInput) *DeleteAppsListFuture {
	future := workflow.ExecuteActivity(ctx, "aws.fms.DeleteAppsList", input)
	return &DeleteAppsListFuture{Future: future}
}

func (a *stub) DeleteNotificationChannel(ctx workflow.Context, input *fms.DeleteNotificationChannelInput) (*fms.DeleteNotificationChannelOutput, error) {
	var output fms.DeleteNotificationChannelOutput
	err := workflow.ExecuteActivity(ctx, "aws.fms.DeleteNotificationChannel", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteNotificationChannelAsync(ctx workflow.Context, input *fms.DeleteNotificationChannelInput) *DeleteNotificationChannelFuture {
	future := workflow.ExecuteActivity(ctx, "aws.fms.DeleteNotificationChannel", input)
	return &DeleteNotificationChannelFuture{Future: future}
}

func (a *stub) DeletePolicy(ctx workflow.Context, input *fms.DeletePolicyInput) (*fms.DeletePolicyOutput, error) {
	var output fms.DeletePolicyOutput
	err := workflow.ExecuteActivity(ctx, "aws.fms.DeletePolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeletePolicyAsync(ctx workflow.Context, input *fms.DeletePolicyInput) *DeletePolicyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.fms.DeletePolicy", input)
	return &DeletePolicyFuture{Future: future}
}

func (a *stub) DeleteProtocolsList(ctx workflow.Context, input *fms.DeleteProtocolsListInput) (*fms.DeleteProtocolsListOutput, error) {
	var output fms.DeleteProtocolsListOutput
	err := workflow.ExecuteActivity(ctx, "aws.fms.DeleteProtocolsList", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteProtocolsListAsync(ctx workflow.Context, input *fms.DeleteProtocolsListInput) *DeleteProtocolsListFuture {
	future := workflow.ExecuteActivity(ctx, "aws.fms.DeleteProtocolsList", input)
	return &DeleteProtocolsListFuture{Future: future}
}

func (a *stub) DisassociateAdminAccount(ctx workflow.Context, input *fms.DisassociateAdminAccountInput) (*fms.DisassociateAdminAccountOutput, error) {
	var output fms.DisassociateAdminAccountOutput
	err := workflow.ExecuteActivity(ctx, "aws.fms.DisassociateAdminAccount", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DisassociateAdminAccountAsync(ctx workflow.Context, input *fms.DisassociateAdminAccountInput) *DisassociateAdminAccountFuture {
	future := workflow.ExecuteActivity(ctx, "aws.fms.DisassociateAdminAccount", input)
	return &DisassociateAdminAccountFuture{Future: future}
}

func (a *stub) GetAdminAccount(ctx workflow.Context, input *fms.GetAdminAccountInput) (*fms.GetAdminAccountOutput, error) {
	var output fms.GetAdminAccountOutput
	err := workflow.ExecuteActivity(ctx, "aws.fms.GetAdminAccount", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetAdminAccountAsync(ctx workflow.Context, input *fms.GetAdminAccountInput) *GetAdminAccountFuture {
	future := workflow.ExecuteActivity(ctx, "aws.fms.GetAdminAccount", input)
	return &GetAdminAccountFuture{Future: future}
}

func (a *stub) GetAppsList(ctx workflow.Context, input *fms.GetAppsListInput) (*fms.GetAppsListOutput, error) {
	var output fms.GetAppsListOutput
	err := workflow.ExecuteActivity(ctx, "aws.fms.GetAppsList", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetAppsListAsync(ctx workflow.Context, input *fms.GetAppsListInput) *GetAppsListFuture {
	future := workflow.ExecuteActivity(ctx, "aws.fms.GetAppsList", input)
	return &GetAppsListFuture{Future: future}
}

func (a *stub) GetComplianceDetail(ctx workflow.Context, input *fms.GetComplianceDetailInput) (*fms.GetComplianceDetailOutput, error) {
	var output fms.GetComplianceDetailOutput
	err := workflow.ExecuteActivity(ctx, "aws.fms.GetComplianceDetail", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetComplianceDetailAsync(ctx workflow.Context, input *fms.GetComplianceDetailInput) *GetComplianceDetailFuture {
	future := workflow.ExecuteActivity(ctx, "aws.fms.GetComplianceDetail", input)
	return &GetComplianceDetailFuture{Future: future}
}

func (a *stub) GetNotificationChannel(ctx workflow.Context, input *fms.GetNotificationChannelInput) (*fms.GetNotificationChannelOutput, error) {
	var output fms.GetNotificationChannelOutput
	err := workflow.ExecuteActivity(ctx, "aws.fms.GetNotificationChannel", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetNotificationChannelAsync(ctx workflow.Context, input *fms.GetNotificationChannelInput) *GetNotificationChannelFuture {
	future := workflow.ExecuteActivity(ctx, "aws.fms.GetNotificationChannel", input)
	return &GetNotificationChannelFuture{Future: future}
}

func (a *stub) GetPolicy(ctx workflow.Context, input *fms.GetPolicyInput) (*fms.GetPolicyOutput, error) {
	var output fms.GetPolicyOutput
	err := workflow.ExecuteActivity(ctx, "aws.fms.GetPolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetPolicyAsync(ctx workflow.Context, input *fms.GetPolicyInput) *GetPolicyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.fms.GetPolicy", input)
	return &GetPolicyFuture{Future: future}
}

func (a *stub) GetProtectionStatus(ctx workflow.Context, input *fms.GetProtectionStatusInput) (*fms.GetProtectionStatusOutput, error) {
	var output fms.GetProtectionStatusOutput
	err := workflow.ExecuteActivity(ctx, "aws.fms.GetProtectionStatus", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetProtectionStatusAsync(ctx workflow.Context, input *fms.GetProtectionStatusInput) *GetProtectionStatusFuture {
	future := workflow.ExecuteActivity(ctx, "aws.fms.GetProtectionStatus", input)
	return &GetProtectionStatusFuture{Future: future}
}

func (a *stub) GetProtocolsList(ctx workflow.Context, input *fms.GetProtocolsListInput) (*fms.GetProtocolsListOutput, error) {
	var output fms.GetProtocolsListOutput
	err := workflow.ExecuteActivity(ctx, "aws.fms.GetProtocolsList", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetProtocolsListAsync(ctx workflow.Context, input *fms.GetProtocolsListInput) *GetProtocolsListFuture {
	future := workflow.ExecuteActivity(ctx, "aws.fms.GetProtocolsList", input)
	return &GetProtocolsListFuture{Future: future}
}

func (a *stub) GetViolationDetails(ctx workflow.Context, input *fms.GetViolationDetailsInput) (*fms.GetViolationDetailsOutput, error) {
	var output fms.GetViolationDetailsOutput
	err := workflow.ExecuteActivity(ctx, "aws.fms.GetViolationDetails", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) GetViolationDetailsAsync(ctx workflow.Context, input *fms.GetViolationDetailsInput) *GetViolationDetailsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.fms.GetViolationDetails", input)
	return &GetViolationDetailsFuture{Future: future}
}

func (a *stub) ListAppsLists(ctx workflow.Context, input *fms.ListAppsListsInput) (*fms.ListAppsListsOutput, error) {
	var output fms.ListAppsListsOutput
	err := workflow.ExecuteActivity(ctx, "aws.fms.ListAppsLists", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListAppsListsAsync(ctx workflow.Context, input *fms.ListAppsListsInput) *ListAppsListsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.fms.ListAppsLists", input)
	return &ListAppsListsFuture{Future: future}
}

func (a *stub) ListComplianceStatus(ctx workflow.Context, input *fms.ListComplianceStatusInput) (*fms.ListComplianceStatusOutput, error) {
	var output fms.ListComplianceStatusOutput
	err := workflow.ExecuteActivity(ctx, "aws.fms.ListComplianceStatus", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListComplianceStatusAsync(ctx workflow.Context, input *fms.ListComplianceStatusInput) *ListComplianceStatusFuture {
	future := workflow.ExecuteActivity(ctx, "aws.fms.ListComplianceStatus", input)
	return &ListComplianceStatusFuture{Future: future}
}

func (a *stub) ListMemberAccounts(ctx workflow.Context, input *fms.ListMemberAccountsInput) (*fms.ListMemberAccountsOutput, error) {
	var output fms.ListMemberAccountsOutput
	err := workflow.ExecuteActivity(ctx, "aws.fms.ListMemberAccounts", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListMemberAccountsAsync(ctx workflow.Context, input *fms.ListMemberAccountsInput) *ListMemberAccountsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.fms.ListMemberAccounts", input)
	return &ListMemberAccountsFuture{Future: future}
}

func (a *stub) ListPolicies(ctx workflow.Context, input *fms.ListPoliciesInput) (*fms.ListPoliciesOutput, error) {
	var output fms.ListPoliciesOutput
	err := workflow.ExecuteActivity(ctx, "aws.fms.ListPolicies", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListPoliciesAsync(ctx workflow.Context, input *fms.ListPoliciesInput) *ListPoliciesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.fms.ListPolicies", input)
	return &ListPoliciesFuture{Future: future}
}

func (a *stub) ListProtocolsLists(ctx workflow.Context, input *fms.ListProtocolsListsInput) (*fms.ListProtocolsListsOutput, error) {
	var output fms.ListProtocolsListsOutput
	err := workflow.ExecuteActivity(ctx, "aws.fms.ListProtocolsLists", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListProtocolsListsAsync(ctx workflow.Context, input *fms.ListProtocolsListsInput) *ListProtocolsListsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.fms.ListProtocolsLists", input)
	return &ListProtocolsListsFuture{Future: future}
}

func (a *stub) ListTagsForResource(ctx workflow.Context, input *fms.ListTagsForResourceInput) (*fms.ListTagsForResourceOutput, error) {
	var output fms.ListTagsForResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.fms.ListTagsForResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ListTagsForResourceAsync(ctx workflow.Context, input *fms.ListTagsForResourceInput) *ListTagsForResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.fms.ListTagsForResource", input)
	return &ListTagsForResourceFuture{Future: future}
}

func (a *stub) PutAppsList(ctx workflow.Context, input *fms.PutAppsListInput) (*fms.PutAppsListOutput, error) {
	var output fms.PutAppsListOutput
	err := workflow.ExecuteActivity(ctx, "aws.fms.PutAppsList", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PutAppsListAsync(ctx workflow.Context, input *fms.PutAppsListInput) *PutAppsListFuture {
	future := workflow.ExecuteActivity(ctx, "aws.fms.PutAppsList", input)
	return &PutAppsListFuture{Future: future}
}

func (a *stub) PutNotificationChannel(ctx workflow.Context, input *fms.PutNotificationChannelInput) (*fms.PutNotificationChannelOutput, error) {
	var output fms.PutNotificationChannelOutput
	err := workflow.ExecuteActivity(ctx, "aws.fms.PutNotificationChannel", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PutNotificationChannelAsync(ctx workflow.Context, input *fms.PutNotificationChannelInput) *PutNotificationChannelFuture {
	future := workflow.ExecuteActivity(ctx, "aws.fms.PutNotificationChannel", input)
	return &PutNotificationChannelFuture{Future: future}
}

func (a *stub) PutPolicy(ctx workflow.Context, input *fms.PutPolicyInput) (*fms.PutPolicyOutput, error) {
	var output fms.PutPolicyOutput
	err := workflow.ExecuteActivity(ctx, "aws.fms.PutPolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PutPolicyAsync(ctx workflow.Context, input *fms.PutPolicyInput) *PutPolicyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.fms.PutPolicy", input)
	return &PutPolicyFuture{Future: future}
}

func (a *stub) PutProtocolsList(ctx workflow.Context, input *fms.PutProtocolsListInput) (*fms.PutProtocolsListOutput, error) {
	var output fms.PutProtocolsListOutput
	err := workflow.ExecuteActivity(ctx, "aws.fms.PutProtocolsList", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) PutProtocolsListAsync(ctx workflow.Context, input *fms.PutProtocolsListInput) *PutProtocolsListFuture {
	future := workflow.ExecuteActivity(ctx, "aws.fms.PutProtocolsList", input)
	return &PutProtocolsListFuture{Future: future}
}

func (a *stub) TagResource(ctx workflow.Context, input *fms.TagResourceInput) (*fms.TagResourceOutput, error) {
	var output fms.TagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.fms.TagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) TagResourceAsync(ctx workflow.Context, input *fms.TagResourceInput) *TagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.fms.TagResource", input)
	return &TagResourceFuture{Future: future}
}

func (a *stub) UntagResource(ctx workflow.Context, input *fms.UntagResourceInput) (*fms.UntagResourceOutput, error) {
	var output fms.UntagResourceOutput
	err := workflow.ExecuteActivity(ctx, "aws.fms.UntagResource", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) UntagResourceAsync(ctx workflow.Context, input *fms.UntagResourceInput) *UntagResourceFuture {
	future := workflow.ExecuteActivity(ctx, "aws.fms.UntagResource", input)
	return &UntagResourceFuture{Future: future}
}