// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package elbstub

import (
	"github.com/aws/aws-sdk-go/service/elb"
	"go.temporal.io/sdk/workflow"

	"go.temporal.io/aws-sdk/clients"
)

// ensure that imports are valid even if not used by the generated code
var _ clients.VoidFuture

type stub struct{}

type AddTagsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AddTagsFuture) Get(ctx workflow.Context) (*elb.AddTagsOutput, error) {
	var output elb.AddTagsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ApplySecurityGroupsToLoadBalancerFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ApplySecurityGroupsToLoadBalancerFuture) Get(ctx workflow.Context) (*elb.ApplySecurityGroupsToLoadBalancerOutput, error) {
	var output elb.ApplySecurityGroupsToLoadBalancerOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type AttachLoadBalancerToSubnetsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *AttachLoadBalancerToSubnetsFuture) Get(ctx workflow.Context) (*elb.AttachLoadBalancerToSubnetsOutput, error) {
	var output elb.AttachLoadBalancerToSubnetsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ConfigureHealthCheckFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ConfigureHealthCheckFuture) Get(ctx workflow.Context) (*elb.ConfigureHealthCheckOutput, error) {
	var output elb.ConfigureHealthCheckOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateAppCookieStickinessPolicyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateAppCookieStickinessPolicyFuture) Get(ctx workflow.Context) (*elb.CreateAppCookieStickinessPolicyOutput, error) {
	var output elb.CreateAppCookieStickinessPolicyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateLBCookieStickinessPolicyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateLBCookieStickinessPolicyFuture) Get(ctx workflow.Context) (*elb.CreateLBCookieStickinessPolicyOutput, error) {
	var output elb.CreateLBCookieStickinessPolicyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateLoadBalancerFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateLoadBalancerFuture) Get(ctx workflow.Context) (*elb.CreateLoadBalancerOutput, error) {
	var output elb.CreateLoadBalancerOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateLoadBalancerListenersFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateLoadBalancerListenersFuture) Get(ctx workflow.Context) (*elb.CreateLoadBalancerListenersOutput, error) {
	var output elb.CreateLoadBalancerListenersOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type CreateLoadBalancerPolicyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *CreateLoadBalancerPolicyFuture) Get(ctx workflow.Context) (*elb.CreateLoadBalancerPolicyOutput, error) {
	var output elb.CreateLoadBalancerPolicyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteLoadBalancerFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteLoadBalancerFuture) Get(ctx workflow.Context) (*elb.DeleteLoadBalancerOutput, error) {
	var output elb.DeleteLoadBalancerOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteLoadBalancerListenersFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteLoadBalancerListenersFuture) Get(ctx workflow.Context) (*elb.DeleteLoadBalancerListenersOutput, error) {
	var output elb.DeleteLoadBalancerListenersOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeleteLoadBalancerPolicyFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeleteLoadBalancerPolicyFuture) Get(ctx workflow.Context) (*elb.DeleteLoadBalancerPolicyOutput, error) {
	var output elb.DeleteLoadBalancerPolicyOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DeregisterInstancesFromLoadBalancerFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DeregisterInstancesFromLoadBalancerFuture) Get(ctx workflow.Context) (*elb.DeregisterInstancesFromLoadBalancerOutput, error) {
	var output elb.DeregisterInstancesFromLoadBalancerOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeAccountLimitsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeAccountLimitsFuture) Get(ctx workflow.Context) (*elb.DescribeAccountLimitsOutput, error) {
	var output elb.DescribeAccountLimitsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeInstanceHealthFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeInstanceHealthFuture) Get(ctx workflow.Context) (*elb.DescribeInstanceHealthOutput, error) {
	var output elb.DescribeInstanceHealthOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeLoadBalancerAttributesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeLoadBalancerAttributesFuture) Get(ctx workflow.Context) (*elb.DescribeLoadBalancerAttributesOutput, error) {
	var output elb.DescribeLoadBalancerAttributesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeLoadBalancerPoliciesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeLoadBalancerPoliciesFuture) Get(ctx workflow.Context) (*elb.DescribeLoadBalancerPoliciesOutput, error) {
	var output elb.DescribeLoadBalancerPoliciesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeLoadBalancerPolicyTypesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeLoadBalancerPolicyTypesFuture) Get(ctx workflow.Context) (*elb.DescribeLoadBalancerPolicyTypesOutput, error) {
	var output elb.DescribeLoadBalancerPolicyTypesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeLoadBalancersFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeLoadBalancersFuture) Get(ctx workflow.Context) (*elb.DescribeLoadBalancersOutput, error) {
	var output elb.DescribeLoadBalancersOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DescribeTagsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DescribeTagsFuture) Get(ctx workflow.Context) (*elb.DescribeTagsOutput, error) {
	var output elb.DescribeTagsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DetachLoadBalancerFromSubnetsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DetachLoadBalancerFromSubnetsFuture) Get(ctx workflow.Context) (*elb.DetachLoadBalancerFromSubnetsOutput, error) {
	var output elb.DetachLoadBalancerFromSubnetsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type DisableAvailabilityZonesForLoadBalancerFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *DisableAvailabilityZonesForLoadBalancerFuture) Get(ctx workflow.Context) (*elb.DisableAvailabilityZonesForLoadBalancerOutput, error) {
	var output elb.DisableAvailabilityZonesForLoadBalancerOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type EnableAvailabilityZonesForLoadBalancerFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *EnableAvailabilityZonesForLoadBalancerFuture) Get(ctx workflow.Context) (*elb.EnableAvailabilityZonesForLoadBalancerOutput, error) {
	var output elb.EnableAvailabilityZonesForLoadBalancerOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type ModifyLoadBalancerAttributesFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *ModifyLoadBalancerAttributesFuture) Get(ctx workflow.Context) (*elb.ModifyLoadBalancerAttributesOutput, error) {
	var output elb.ModifyLoadBalancerAttributesOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type RegisterInstancesWithLoadBalancerFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *RegisterInstancesWithLoadBalancerFuture) Get(ctx workflow.Context) (*elb.RegisterInstancesWithLoadBalancerOutput, error) {
	var output elb.RegisterInstancesWithLoadBalancerOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type RemoveTagsFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *RemoveTagsFuture) Get(ctx workflow.Context) (*elb.RemoveTagsOutput, error) {
	var output elb.RemoveTagsOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SetLoadBalancerListenerSSLCertificateFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SetLoadBalancerListenerSSLCertificateFuture) Get(ctx workflow.Context) (*elb.SetLoadBalancerListenerSSLCertificateOutput, error) {
	var output elb.SetLoadBalancerListenerSSLCertificateOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SetLoadBalancerPoliciesForBackendServerFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SetLoadBalancerPoliciesForBackendServerFuture) Get(ctx workflow.Context) (*elb.SetLoadBalancerPoliciesForBackendServerOutput, error) {
	var output elb.SetLoadBalancerPoliciesForBackendServerOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

type SetLoadBalancerPoliciesOfListenerFuture struct {
	// public to support Selector.addFuture
	Future workflow.Future
}

func (r *SetLoadBalancerPoliciesOfListenerFuture) Get(ctx workflow.Context) (*elb.SetLoadBalancerPoliciesOfListenerOutput, error) {
	var output elb.SetLoadBalancerPoliciesOfListenerOutput
	err := r.Future.Get(ctx, &output)
	return &output, err
}

func (a *stub) AddTags(ctx workflow.Context, input *elb.AddTagsInput) (*elb.AddTagsOutput, error) {
	var output elb.AddTagsOutput
	err := workflow.ExecuteActivity(ctx, "aws.elb.AddTags", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) AddTagsAsync(ctx workflow.Context, input *elb.AddTagsInput) *AddTagsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.elb.AddTags", input)
	return &AddTagsFuture{Future: future}
}

func (a *stub) ApplySecurityGroupsToLoadBalancer(ctx workflow.Context, input *elb.ApplySecurityGroupsToLoadBalancerInput) (*elb.ApplySecurityGroupsToLoadBalancerOutput, error) {
	var output elb.ApplySecurityGroupsToLoadBalancerOutput
	err := workflow.ExecuteActivity(ctx, "aws.elb.ApplySecurityGroupsToLoadBalancer", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ApplySecurityGroupsToLoadBalancerAsync(ctx workflow.Context, input *elb.ApplySecurityGroupsToLoadBalancerInput) *ApplySecurityGroupsToLoadBalancerFuture {
	future := workflow.ExecuteActivity(ctx, "aws.elb.ApplySecurityGroupsToLoadBalancer", input)
	return &ApplySecurityGroupsToLoadBalancerFuture{Future: future}
}

func (a *stub) AttachLoadBalancerToSubnets(ctx workflow.Context, input *elb.AttachLoadBalancerToSubnetsInput) (*elb.AttachLoadBalancerToSubnetsOutput, error) {
	var output elb.AttachLoadBalancerToSubnetsOutput
	err := workflow.ExecuteActivity(ctx, "aws.elb.AttachLoadBalancerToSubnets", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) AttachLoadBalancerToSubnetsAsync(ctx workflow.Context, input *elb.AttachLoadBalancerToSubnetsInput) *AttachLoadBalancerToSubnetsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.elb.AttachLoadBalancerToSubnets", input)
	return &AttachLoadBalancerToSubnetsFuture{Future: future}
}

func (a *stub) ConfigureHealthCheck(ctx workflow.Context, input *elb.ConfigureHealthCheckInput) (*elb.ConfigureHealthCheckOutput, error) {
	var output elb.ConfigureHealthCheckOutput
	err := workflow.ExecuteActivity(ctx, "aws.elb.ConfigureHealthCheck", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ConfigureHealthCheckAsync(ctx workflow.Context, input *elb.ConfigureHealthCheckInput) *ConfigureHealthCheckFuture {
	future := workflow.ExecuteActivity(ctx, "aws.elb.ConfigureHealthCheck", input)
	return &ConfigureHealthCheckFuture{Future: future}
}

func (a *stub) CreateAppCookieStickinessPolicy(ctx workflow.Context, input *elb.CreateAppCookieStickinessPolicyInput) (*elb.CreateAppCookieStickinessPolicyOutput, error) {
	var output elb.CreateAppCookieStickinessPolicyOutput
	err := workflow.ExecuteActivity(ctx, "aws.elb.CreateAppCookieStickinessPolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateAppCookieStickinessPolicyAsync(ctx workflow.Context, input *elb.CreateAppCookieStickinessPolicyInput) *CreateAppCookieStickinessPolicyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.elb.CreateAppCookieStickinessPolicy", input)
	return &CreateAppCookieStickinessPolicyFuture{Future: future}
}

func (a *stub) CreateLBCookieStickinessPolicy(ctx workflow.Context, input *elb.CreateLBCookieStickinessPolicyInput) (*elb.CreateLBCookieStickinessPolicyOutput, error) {
	var output elb.CreateLBCookieStickinessPolicyOutput
	err := workflow.ExecuteActivity(ctx, "aws.elb.CreateLBCookieStickinessPolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateLBCookieStickinessPolicyAsync(ctx workflow.Context, input *elb.CreateLBCookieStickinessPolicyInput) *CreateLBCookieStickinessPolicyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.elb.CreateLBCookieStickinessPolicy", input)
	return &CreateLBCookieStickinessPolicyFuture{Future: future}
}

func (a *stub) CreateLoadBalancer(ctx workflow.Context, input *elb.CreateLoadBalancerInput) (*elb.CreateLoadBalancerOutput, error) {
	var output elb.CreateLoadBalancerOutput
	err := workflow.ExecuteActivity(ctx, "aws.elb.CreateLoadBalancer", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateLoadBalancerAsync(ctx workflow.Context, input *elb.CreateLoadBalancerInput) *CreateLoadBalancerFuture {
	future := workflow.ExecuteActivity(ctx, "aws.elb.CreateLoadBalancer", input)
	return &CreateLoadBalancerFuture{Future: future}
}

func (a *stub) CreateLoadBalancerListeners(ctx workflow.Context, input *elb.CreateLoadBalancerListenersInput) (*elb.CreateLoadBalancerListenersOutput, error) {
	var output elb.CreateLoadBalancerListenersOutput
	err := workflow.ExecuteActivity(ctx, "aws.elb.CreateLoadBalancerListeners", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateLoadBalancerListenersAsync(ctx workflow.Context, input *elb.CreateLoadBalancerListenersInput) *CreateLoadBalancerListenersFuture {
	future := workflow.ExecuteActivity(ctx, "aws.elb.CreateLoadBalancerListeners", input)
	return &CreateLoadBalancerListenersFuture{Future: future}
}

func (a *stub) CreateLoadBalancerPolicy(ctx workflow.Context, input *elb.CreateLoadBalancerPolicyInput) (*elb.CreateLoadBalancerPolicyOutput, error) {
	var output elb.CreateLoadBalancerPolicyOutput
	err := workflow.ExecuteActivity(ctx, "aws.elb.CreateLoadBalancerPolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) CreateLoadBalancerPolicyAsync(ctx workflow.Context, input *elb.CreateLoadBalancerPolicyInput) *CreateLoadBalancerPolicyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.elb.CreateLoadBalancerPolicy", input)
	return &CreateLoadBalancerPolicyFuture{Future: future}
}

func (a *stub) DeleteLoadBalancer(ctx workflow.Context, input *elb.DeleteLoadBalancerInput) (*elb.DeleteLoadBalancerOutput, error) {
	var output elb.DeleteLoadBalancerOutput
	err := workflow.ExecuteActivity(ctx, "aws.elb.DeleteLoadBalancer", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteLoadBalancerAsync(ctx workflow.Context, input *elb.DeleteLoadBalancerInput) *DeleteLoadBalancerFuture {
	future := workflow.ExecuteActivity(ctx, "aws.elb.DeleteLoadBalancer", input)
	return &DeleteLoadBalancerFuture{Future: future}
}

func (a *stub) DeleteLoadBalancerListeners(ctx workflow.Context, input *elb.DeleteLoadBalancerListenersInput) (*elb.DeleteLoadBalancerListenersOutput, error) {
	var output elb.DeleteLoadBalancerListenersOutput
	err := workflow.ExecuteActivity(ctx, "aws.elb.DeleteLoadBalancerListeners", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteLoadBalancerListenersAsync(ctx workflow.Context, input *elb.DeleteLoadBalancerListenersInput) *DeleteLoadBalancerListenersFuture {
	future := workflow.ExecuteActivity(ctx, "aws.elb.DeleteLoadBalancerListeners", input)
	return &DeleteLoadBalancerListenersFuture{Future: future}
}

func (a *stub) DeleteLoadBalancerPolicy(ctx workflow.Context, input *elb.DeleteLoadBalancerPolicyInput) (*elb.DeleteLoadBalancerPolicyOutput, error) {
	var output elb.DeleteLoadBalancerPolicyOutput
	err := workflow.ExecuteActivity(ctx, "aws.elb.DeleteLoadBalancerPolicy", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeleteLoadBalancerPolicyAsync(ctx workflow.Context, input *elb.DeleteLoadBalancerPolicyInput) *DeleteLoadBalancerPolicyFuture {
	future := workflow.ExecuteActivity(ctx, "aws.elb.DeleteLoadBalancerPolicy", input)
	return &DeleteLoadBalancerPolicyFuture{Future: future}
}

func (a *stub) DeregisterInstancesFromLoadBalancer(ctx workflow.Context, input *elb.DeregisterInstancesFromLoadBalancerInput) (*elb.DeregisterInstancesFromLoadBalancerOutput, error) {
	var output elb.DeregisterInstancesFromLoadBalancerOutput
	err := workflow.ExecuteActivity(ctx, "aws.elb.DeregisterInstancesFromLoadBalancer", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DeregisterInstancesFromLoadBalancerAsync(ctx workflow.Context, input *elb.DeregisterInstancesFromLoadBalancerInput) *DeregisterInstancesFromLoadBalancerFuture {
	future := workflow.ExecuteActivity(ctx, "aws.elb.DeregisterInstancesFromLoadBalancer", input)
	return &DeregisterInstancesFromLoadBalancerFuture{Future: future}
}

func (a *stub) DescribeAccountLimits(ctx workflow.Context, input *elb.DescribeAccountLimitsInput) (*elb.DescribeAccountLimitsOutput, error) {
	var output elb.DescribeAccountLimitsOutput
	err := workflow.ExecuteActivity(ctx, "aws.elb.DescribeAccountLimits", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeAccountLimitsAsync(ctx workflow.Context, input *elb.DescribeAccountLimitsInput) *DescribeAccountLimitsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.elb.DescribeAccountLimits", input)
	return &DescribeAccountLimitsFuture{Future: future}
}

func (a *stub) DescribeInstanceHealth(ctx workflow.Context, input *elb.DescribeInstanceHealthInput) (*elb.DescribeInstanceHealthOutput, error) {
	var output elb.DescribeInstanceHealthOutput
	err := workflow.ExecuteActivity(ctx, "aws.elb.DescribeInstanceHealth", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeInstanceHealthAsync(ctx workflow.Context, input *elb.DescribeInstanceHealthInput) *DescribeInstanceHealthFuture {
	future := workflow.ExecuteActivity(ctx, "aws.elb.DescribeInstanceHealth", input)
	return &DescribeInstanceHealthFuture{Future: future}
}

func (a *stub) DescribeLoadBalancerAttributes(ctx workflow.Context, input *elb.DescribeLoadBalancerAttributesInput) (*elb.DescribeLoadBalancerAttributesOutput, error) {
	var output elb.DescribeLoadBalancerAttributesOutput
	err := workflow.ExecuteActivity(ctx, "aws.elb.DescribeLoadBalancerAttributes", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeLoadBalancerAttributesAsync(ctx workflow.Context, input *elb.DescribeLoadBalancerAttributesInput) *DescribeLoadBalancerAttributesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.elb.DescribeLoadBalancerAttributes", input)
	return &DescribeLoadBalancerAttributesFuture{Future: future}
}

func (a *stub) DescribeLoadBalancerPolicies(ctx workflow.Context, input *elb.DescribeLoadBalancerPoliciesInput) (*elb.DescribeLoadBalancerPoliciesOutput, error) {
	var output elb.DescribeLoadBalancerPoliciesOutput
	err := workflow.ExecuteActivity(ctx, "aws.elb.DescribeLoadBalancerPolicies", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeLoadBalancerPoliciesAsync(ctx workflow.Context, input *elb.DescribeLoadBalancerPoliciesInput) *DescribeLoadBalancerPoliciesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.elb.DescribeLoadBalancerPolicies", input)
	return &DescribeLoadBalancerPoliciesFuture{Future: future}
}

func (a *stub) DescribeLoadBalancerPolicyTypes(ctx workflow.Context, input *elb.DescribeLoadBalancerPolicyTypesInput) (*elb.DescribeLoadBalancerPolicyTypesOutput, error) {
	var output elb.DescribeLoadBalancerPolicyTypesOutput
	err := workflow.ExecuteActivity(ctx, "aws.elb.DescribeLoadBalancerPolicyTypes", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeLoadBalancerPolicyTypesAsync(ctx workflow.Context, input *elb.DescribeLoadBalancerPolicyTypesInput) *DescribeLoadBalancerPolicyTypesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.elb.DescribeLoadBalancerPolicyTypes", input)
	return &DescribeLoadBalancerPolicyTypesFuture{Future: future}
}

func (a *stub) DescribeLoadBalancers(ctx workflow.Context, input *elb.DescribeLoadBalancersInput) (*elb.DescribeLoadBalancersOutput, error) {
	var output elb.DescribeLoadBalancersOutput
	err := workflow.ExecuteActivity(ctx, "aws.elb.DescribeLoadBalancers", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeLoadBalancersAsync(ctx workflow.Context, input *elb.DescribeLoadBalancersInput) *DescribeLoadBalancersFuture {
	future := workflow.ExecuteActivity(ctx, "aws.elb.DescribeLoadBalancers", input)
	return &DescribeLoadBalancersFuture{Future: future}
}

func (a *stub) DescribeTags(ctx workflow.Context, input *elb.DescribeTagsInput) (*elb.DescribeTagsOutput, error) {
	var output elb.DescribeTagsOutput
	err := workflow.ExecuteActivity(ctx, "aws.elb.DescribeTags", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DescribeTagsAsync(ctx workflow.Context, input *elb.DescribeTagsInput) *DescribeTagsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.elb.DescribeTags", input)
	return &DescribeTagsFuture{Future: future}
}

func (a *stub) DetachLoadBalancerFromSubnets(ctx workflow.Context, input *elb.DetachLoadBalancerFromSubnetsInput) (*elb.DetachLoadBalancerFromSubnetsOutput, error) {
	var output elb.DetachLoadBalancerFromSubnetsOutput
	err := workflow.ExecuteActivity(ctx, "aws.elb.DetachLoadBalancerFromSubnets", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DetachLoadBalancerFromSubnetsAsync(ctx workflow.Context, input *elb.DetachLoadBalancerFromSubnetsInput) *DetachLoadBalancerFromSubnetsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.elb.DetachLoadBalancerFromSubnets", input)
	return &DetachLoadBalancerFromSubnetsFuture{Future: future}
}

func (a *stub) DisableAvailabilityZonesForLoadBalancer(ctx workflow.Context, input *elb.DisableAvailabilityZonesForLoadBalancerInput) (*elb.DisableAvailabilityZonesForLoadBalancerOutput, error) {
	var output elb.DisableAvailabilityZonesForLoadBalancerOutput
	err := workflow.ExecuteActivity(ctx, "aws.elb.DisableAvailabilityZonesForLoadBalancer", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) DisableAvailabilityZonesForLoadBalancerAsync(ctx workflow.Context, input *elb.DisableAvailabilityZonesForLoadBalancerInput) *DisableAvailabilityZonesForLoadBalancerFuture {
	future := workflow.ExecuteActivity(ctx, "aws.elb.DisableAvailabilityZonesForLoadBalancer", input)
	return &DisableAvailabilityZonesForLoadBalancerFuture{Future: future}
}

func (a *stub) EnableAvailabilityZonesForLoadBalancer(ctx workflow.Context, input *elb.EnableAvailabilityZonesForLoadBalancerInput) (*elb.EnableAvailabilityZonesForLoadBalancerOutput, error) {
	var output elb.EnableAvailabilityZonesForLoadBalancerOutput
	err := workflow.ExecuteActivity(ctx, "aws.elb.EnableAvailabilityZonesForLoadBalancer", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) EnableAvailabilityZonesForLoadBalancerAsync(ctx workflow.Context, input *elb.EnableAvailabilityZonesForLoadBalancerInput) *EnableAvailabilityZonesForLoadBalancerFuture {
	future := workflow.ExecuteActivity(ctx, "aws.elb.EnableAvailabilityZonesForLoadBalancer", input)
	return &EnableAvailabilityZonesForLoadBalancerFuture{Future: future}
}

func (a *stub) ModifyLoadBalancerAttributes(ctx workflow.Context, input *elb.ModifyLoadBalancerAttributesInput) (*elb.ModifyLoadBalancerAttributesOutput, error) {
	var output elb.ModifyLoadBalancerAttributesOutput
	err := workflow.ExecuteActivity(ctx, "aws.elb.ModifyLoadBalancerAttributes", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) ModifyLoadBalancerAttributesAsync(ctx workflow.Context, input *elb.ModifyLoadBalancerAttributesInput) *ModifyLoadBalancerAttributesFuture {
	future := workflow.ExecuteActivity(ctx, "aws.elb.ModifyLoadBalancerAttributes", input)
	return &ModifyLoadBalancerAttributesFuture{Future: future}
}

func (a *stub) RegisterInstancesWithLoadBalancer(ctx workflow.Context, input *elb.RegisterInstancesWithLoadBalancerInput) (*elb.RegisterInstancesWithLoadBalancerOutput, error) {
	var output elb.RegisterInstancesWithLoadBalancerOutput
	err := workflow.ExecuteActivity(ctx, "aws.elb.RegisterInstancesWithLoadBalancer", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) RegisterInstancesWithLoadBalancerAsync(ctx workflow.Context, input *elb.RegisterInstancesWithLoadBalancerInput) *RegisterInstancesWithLoadBalancerFuture {
	future := workflow.ExecuteActivity(ctx, "aws.elb.RegisterInstancesWithLoadBalancer", input)
	return &RegisterInstancesWithLoadBalancerFuture{Future: future}
}

func (a *stub) RemoveTags(ctx workflow.Context, input *elb.RemoveTagsInput) (*elb.RemoveTagsOutput, error) {
	var output elb.RemoveTagsOutput
	err := workflow.ExecuteActivity(ctx, "aws.elb.RemoveTags", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) RemoveTagsAsync(ctx workflow.Context, input *elb.RemoveTagsInput) *RemoveTagsFuture {
	future := workflow.ExecuteActivity(ctx, "aws.elb.RemoveTags", input)
	return &RemoveTagsFuture{Future: future}
}

func (a *stub) SetLoadBalancerListenerSSLCertificate(ctx workflow.Context, input *elb.SetLoadBalancerListenerSSLCertificateInput) (*elb.SetLoadBalancerListenerSSLCertificateOutput, error) {
	var output elb.SetLoadBalancerListenerSSLCertificateOutput
	err := workflow.ExecuteActivity(ctx, "aws.elb.SetLoadBalancerListenerSSLCertificate", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) SetLoadBalancerListenerSSLCertificateAsync(ctx workflow.Context, input *elb.SetLoadBalancerListenerSSLCertificateInput) *SetLoadBalancerListenerSSLCertificateFuture {
	future := workflow.ExecuteActivity(ctx, "aws.elb.SetLoadBalancerListenerSSLCertificate", input)
	return &SetLoadBalancerListenerSSLCertificateFuture{Future: future}
}

func (a *stub) SetLoadBalancerPoliciesForBackendServer(ctx workflow.Context, input *elb.SetLoadBalancerPoliciesForBackendServerInput) (*elb.SetLoadBalancerPoliciesForBackendServerOutput, error) {
	var output elb.SetLoadBalancerPoliciesForBackendServerOutput
	err := workflow.ExecuteActivity(ctx, "aws.elb.SetLoadBalancerPoliciesForBackendServer", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) SetLoadBalancerPoliciesForBackendServerAsync(ctx workflow.Context, input *elb.SetLoadBalancerPoliciesForBackendServerInput) *SetLoadBalancerPoliciesForBackendServerFuture {
	future := workflow.ExecuteActivity(ctx, "aws.elb.SetLoadBalancerPoliciesForBackendServer", input)
	return &SetLoadBalancerPoliciesForBackendServerFuture{Future: future}
}

func (a *stub) SetLoadBalancerPoliciesOfListener(ctx workflow.Context, input *elb.SetLoadBalancerPoliciesOfListenerInput) (*elb.SetLoadBalancerPoliciesOfListenerOutput, error) {
	var output elb.SetLoadBalancerPoliciesOfListenerOutput
	err := workflow.ExecuteActivity(ctx, "aws.elb.SetLoadBalancerPoliciesOfListener", input).Get(ctx, &output)
	return &output, err
}

func (a *stub) SetLoadBalancerPoliciesOfListenerAsync(ctx workflow.Context, input *elb.SetLoadBalancerPoliciesOfListenerInput) *SetLoadBalancerPoliciesOfListenerFuture {
	future := workflow.ExecuteActivity(ctx, "aws.elb.SetLoadBalancerPoliciesOfListener", input)
	return &SetLoadBalancerPoliciesOfListenerFuture{Future: future}
}

func (a *stub) WaitUntilAnyInstanceInService(ctx workflow.Context, input *elb.DescribeInstanceHealthInput) error {
	return workflow.ExecuteActivity(ctx, "aws.elb.WaitUntilAnyInstanceInService", input).Get(ctx, nil)
}

func (a *stub) WaitUntilAnyInstanceInServiceAsync(ctx workflow.Context, input *elb.DescribeInstanceHealthInput) *clients.VoidFuture {
	future := workflow.ExecuteActivity(ctx, "aws.elb.WaitUntilAnyInstanceInService", input)
	return clients.NewVoidFuture(future)
}

func (a *stub) WaitUntilInstanceDeregistered(ctx workflow.Context, input *elb.DescribeInstanceHealthInput) error {
	return workflow.ExecuteActivity(ctx, "aws.elb.WaitUntilInstanceDeregistered", input).Get(ctx, nil)
}

func (a *stub) WaitUntilInstanceDeregisteredAsync(ctx workflow.Context, input *elb.DescribeInstanceHealthInput) *clients.VoidFuture {
	future := workflow.ExecuteActivity(ctx, "aws.elb.WaitUntilInstanceDeregistered", input)
	return clients.NewVoidFuture(future)
}

func (a *stub) WaitUntilInstanceInService(ctx workflow.Context, input *elb.DescribeInstanceHealthInput) error {
	return workflow.ExecuteActivity(ctx, "aws.elb.WaitUntilInstanceInService", input).Get(ctx, nil)
}

func (a *stub) WaitUntilInstanceInServiceAsync(ctx workflow.Context, input *elb.DescribeInstanceHealthInput) *clients.VoidFuture {
	future := workflow.ExecuteActivity(ctx, "aws.elb.WaitUntilInstanceInService", input)
	return clients.NewVoidFuture(future)
}