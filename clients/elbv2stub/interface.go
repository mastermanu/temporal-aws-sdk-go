// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package elbv2stub

import (
	"github.com/aws/aws-sdk-go/service/elbv2"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type Client interface {
	AddListenerCertificates(ctx workflow.Context, input *elbv2.AddListenerCertificatesInput) (*elbv2.AddListenerCertificatesOutput, error)
	AddListenerCertificatesAsync(ctx workflow.Context, input *elbv2.AddListenerCertificatesInput) *AddListenerCertificatesFuture

	AddTags(ctx workflow.Context, input *elbv2.AddTagsInput) (*elbv2.AddTagsOutput, error)
	AddTagsAsync(ctx workflow.Context, input *elbv2.AddTagsInput) *AddTagsFuture

	CreateListener(ctx workflow.Context, input *elbv2.CreateListenerInput) (*elbv2.CreateListenerOutput, error)
	CreateListenerAsync(ctx workflow.Context, input *elbv2.CreateListenerInput) *CreateListenerFuture

	CreateLoadBalancer(ctx workflow.Context, input *elbv2.CreateLoadBalancerInput) (*elbv2.CreateLoadBalancerOutput, error)
	CreateLoadBalancerAsync(ctx workflow.Context, input *elbv2.CreateLoadBalancerInput) *CreateLoadBalancerFuture

	CreateRule(ctx workflow.Context, input *elbv2.CreateRuleInput) (*elbv2.CreateRuleOutput, error)
	CreateRuleAsync(ctx workflow.Context, input *elbv2.CreateRuleInput) *CreateRuleFuture

	CreateTargetGroup(ctx workflow.Context, input *elbv2.CreateTargetGroupInput) (*elbv2.CreateTargetGroupOutput, error)
	CreateTargetGroupAsync(ctx workflow.Context, input *elbv2.CreateTargetGroupInput) *CreateTargetGroupFuture

	DeleteListener(ctx workflow.Context, input *elbv2.DeleteListenerInput) (*elbv2.DeleteListenerOutput, error)
	DeleteListenerAsync(ctx workflow.Context, input *elbv2.DeleteListenerInput) *DeleteListenerFuture

	DeleteLoadBalancer(ctx workflow.Context, input *elbv2.DeleteLoadBalancerInput) (*elbv2.DeleteLoadBalancerOutput, error)
	DeleteLoadBalancerAsync(ctx workflow.Context, input *elbv2.DeleteLoadBalancerInput) *DeleteLoadBalancerFuture

	DeleteRule(ctx workflow.Context, input *elbv2.DeleteRuleInput) (*elbv2.DeleteRuleOutput, error)
	DeleteRuleAsync(ctx workflow.Context, input *elbv2.DeleteRuleInput) *DeleteRuleFuture

	DeleteTargetGroup(ctx workflow.Context, input *elbv2.DeleteTargetGroupInput) (*elbv2.DeleteTargetGroupOutput, error)
	DeleteTargetGroupAsync(ctx workflow.Context, input *elbv2.DeleteTargetGroupInput) *DeleteTargetGroupFuture

	DeregisterTargets(ctx workflow.Context, input *elbv2.DeregisterTargetsInput) (*elbv2.DeregisterTargetsOutput, error)
	DeregisterTargetsAsync(ctx workflow.Context, input *elbv2.DeregisterTargetsInput) *DeregisterTargetsFuture

	DescribeAccountLimits(ctx workflow.Context, input *elbv2.DescribeAccountLimitsInput) (*elbv2.DescribeAccountLimitsOutput, error)
	DescribeAccountLimitsAsync(ctx workflow.Context, input *elbv2.DescribeAccountLimitsInput) *DescribeAccountLimitsFuture

	DescribeListenerCertificates(ctx workflow.Context, input *elbv2.DescribeListenerCertificatesInput) (*elbv2.DescribeListenerCertificatesOutput, error)
	DescribeListenerCertificatesAsync(ctx workflow.Context, input *elbv2.DescribeListenerCertificatesInput) *DescribeListenerCertificatesFuture

	DescribeListeners(ctx workflow.Context, input *elbv2.DescribeListenersInput) (*elbv2.DescribeListenersOutput, error)
	DescribeListenersAsync(ctx workflow.Context, input *elbv2.DescribeListenersInput) *DescribeListenersFuture

	DescribeLoadBalancerAttributes(ctx workflow.Context, input *elbv2.DescribeLoadBalancerAttributesInput) (*elbv2.DescribeLoadBalancerAttributesOutput, error)
	DescribeLoadBalancerAttributesAsync(ctx workflow.Context, input *elbv2.DescribeLoadBalancerAttributesInput) *DescribeLoadBalancerAttributesFuture

	DescribeLoadBalancers(ctx workflow.Context, input *elbv2.DescribeLoadBalancersInput) (*elbv2.DescribeLoadBalancersOutput, error)
	DescribeLoadBalancersAsync(ctx workflow.Context, input *elbv2.DescribeLoadBalancersInput) *DescribeLoadBalancersFuture

	DescribeRules(ctx workflow.Context, input *elbv2.DescribeRulesInput) (*elbv2.DescribeRulesOutput, error)
	DescribeRulesAsync(ctx workflow.Context, input *elbv2.DescribeRulesInput) *DescribeRulesFuture

	DescribeSSLPolicies(ctx workflow.Context, input *elbv2.DescribeSSLPoliciesInput) (*elbv2.DescribeSSLPoliciesOutput, error)
	DescribeSSLPoliciesAsync(ctx workflow.Context, input *elbv2.DescribeSSLPoliciesInput) *DescribeSSLPoliciesFuture

	DescribeTags(ctx workflow.Context, input *elbv2.DescribeTagsInput) (*elbv2.DescribeTagsOutput, error)
	DescribeTagsAsync(ctx workflow.Context, input *elbv2.DescribeTagsInput) *DescribeTagsFuture

	DescribeTargetGroupAttributes(ctx workflow.Context, input *elbv2.DescribeTargetGroupAttributesInput) (*elbv2.DescribeTargetGroupAttributesOutput, error)
	DescribeTargetGroupAttributesAsync(ctx workflow.Context, input *elbv2.DescribeTargetGroupAttributesInput) *DescribeTargetGroupAttributesFuture

	DescribeTargetGroups(ctx workflow.Context, input *elbv2.DescribeTargetGroupsInput) (*elbv2.DescribeTargetGroupsOutput, error)
	DescribeTargetGroupsAsync(ctx workflow.Context, input *elbv2.DescribeTargetGroupsInput) *DescribeTargetGroupsFuture

	DescribeTargetHealth(ctx workflow.Context, input *elbv2.DescribeTargetHealthInput) (*elbv2.DescribeTargetHealthOutput, error)
	DescribeTargetHealthAsync(ctx workflow.Context, input *elbv2.DescribeTargetHealthInput) *DescribeTargetHealthFuture

	ModifyListener(ctx workflow.Context, input *elbv2.ModifyListenerInput) (*elbv2.ModifyListenerOutput, error)
	ModifyListenerAsync(ctx workflow.Context, input *elbv2.ModifyListenerInput) *ModifyListenerFuture

	ModifyLoadBalancerAttributes(ctx workflow.Context, input *elbv2.ModifyLoadBalancerAttributesInput) (*elbv2.ModifyLoadBalancerAttributesOutput, error)
	ModifyLoadBalancerAttributesAsync(ctx workflow.Context, input *elbv2.ModifyLoadBalancerAttributesInput) *ModifyLoadBalancerAttributesFuture

	ModifyRule(ctx workflow.Context, input *elbv2.ModifyRuleInput) (*elbv2.ModifyRuleOutput, error)
	ModifyRuleAsync(ctx workflow.Context, input *elbv2.ModifyRuleInput) *ModifyRuleFuture

	ModifyTargetGroup(ctx workflow.Context, input *elbv2.ModifyTargetGroupInput) (*elbv2.ModifyTargetGroupOutput, error)
	ModifyTargetGroupAsync(ctx workflow.Context, input *elbv2.ModifyTargetGroupInput) *ModifyTargetGroupFuture

	ModifyTargetGroupAttributes(ctx workflow.Context, input *elbv2.ModifyTargetGroupAttributesInput) (*elbv2.ModifyTargetGroupAttributesOutput, error)
	ModifyTargetGroupAttributesAsync(ctx workflow.Context, input *elbv2.ModifyTargetGroupAttributesInput) *ModifyTargetGroupAttributesFuture

	RegisterTargets(ctx workflow.Context, input *elbv2.RegisterTargetsInput) (*elbv2.RegisterTargetsOutput, error)
	RegisterTargetsAsync(ctx workflow.Context, input *elbv2.RegisterTargetsInput) *RegisterTargetsFuture

	RemoveListenerCertificates(ctx workflow.Context, input *elbv2.RemoveListenerCertificatesInput) (*elbv2.RemoveListenerCertificatesOutput, error)
	RemoveListenerCertificatesAsync(ctx workflow.Context, input *elbv2.RemoveListenerCertificatesInput) *RemoveListenerCertificatesFuture

	RemoveTags(ctx workflow.Context, input *elbv2.RemoveTagsInput) (*elbv2.RemoveTagsOutput, error)
	RemoveTagsAsync(ctx workflow.Context, input *elbv2.RemoveTagsInput) *RemoveTagsFuture

	SetIpAddressType(ctx workflow.Context, input *elbv2.SetIpAddressTypeInput) (*elbv2.SetIpAddressTypeOutput, error)
	SetIpAddressTypeAsync(ctx workflow.Context, input *elbv2.SetIpAddressTypeInput) *SetIpAddressTypeFuture

	SetRulePriorities(ctx workflow.Context, input *elbv2.SetRulePrioritiesInput) (*elbv2.SetRulePrioritiesOutput, error)
	SetRulePrioritiesAsync(ctx workflow.Context, input *elbv2.SetRulePrioritiesInput) *SetRulePrioritiesFuture

	SetSecurityGroups(ctx workflow.Context, input *elbv2.SetSecurityGroupsInput) (*elbv2.SetSecurityGroupsOutput, error)
	SetSecurityGroupsAsync(ctx workflow.Context, input *elbv2.SetSecurityGroupsInput) *SetSecurityGroupsFuture

	SetSubnets(ctx workflow.Context, input *elbv2.SetSubnetsInput) (*elbv2.SetSubnetsOutput, error)
	SetSubnetsAsync(ctx workflow.Context, input *elbv2.SetSubnetsInput) *SetSubnetsFuture

	WaitUntilLoadBalancerAvailable(ctx workflow.Context, input *elbv2.DescribeLoadBalancersInput) error
	WaitUntilLoadBalancerAvailableAsync(ctx workflow.Context, input *elbv2.DescribeLoadBalancersInput) *clients.VoidFuture

	WaitUntilLoadBalancerExists(ctx workflow.Context, input *elbv2.DescribeLoadBalancersInput) error
	WaitUntilLoadBalancerExistsAsync(ctx workflow.Context, input *elbv2.DescribeLoadBalancersInput) *clients.VoidFuture

	WaitUntilLoadBalancersDeleted(ctx workflow.Context, input *elbv2.DescribeLoadBalancersInput) error
	WaitUntilLoadBalancersDeletedAsync(ctx workflow.Context, input *elbv2.DescribeLoadBalancersInput) *clients.VoidFuture

	WaitUntilTargetDeregistered(ctx workflow.Context, input *elbv2.DescribeTargetHealthInput) error
	WaitUntilTargetDeregisteredAsync(ctx workflow.Context, input *elbv2.DescribeTargetHealthInput) *clients.VoidFuture

	WaitUntilTargetInService(ctx workflow.Context, input *elbv2.DescribeTargetHealthInput) error
	WaitUntilTargetInServiceAsync(ctx workflow.Context, input *elbv2.DescribeTargetHealthInput) *clients.VoidFuture
}

func NewClient() Client {
	return &stub{}
}