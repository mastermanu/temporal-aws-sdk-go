package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/elb"
	"github.com/aws/aws-sdk-go/service/elb/elbiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type ELBActivities struct {
	client elbiface.ELBAPI
}

func NewELBActivities(session *session.Session, config ...*aws.Config) *ELBActivities {
	client := elb.New(session, config...)
	return &ELBActivities{client: client}
}

func (a *ELBActivities) AddTags(ctx context.Context, input *elb.AddTagsInput) (*elb.AddTagsOutput, error) {
	return a.client.AddTagsWithContext(ctx, input)
}

func (a *ELBActivities) ApplySecurityGroupsToLoadBalancer(ctx context.Context, input *elb.ApplySecurityGroupsToLoadBalancerInput) (*elb.ApplySecurityGroupsToLoadBalancerOutput, error) {
	return a.client.ApplySecurityGroupsToLoadBalancerWithContext(ctx, input)
}

func (a *ELBActivities) AttachLoadBalancerToSubnets(ctx context.Context, input *elb.AttachLoadBalancerToSubnetsInput) (*elb.AttachLoadBalancerToSubnetsOutput, error) {
	return a.client.AttachLoadBalancerToSubnetsWithContext(ctx, input)
}

func (a *ELBActivities) ConfigureHealthCheck(ctx context.Context, input *elb.ConfigureHealthCheckInput) (*elb.ConfigureHealthCheckOutput, error) {
	return a.client.ConfigureHealthCheckWithContext(ctx, input)
}

func (a *ELBActivities) CreateAppCookieStickinessPolicy(ctx context.Context, input *elb.CreateAppCookieStickinessPolicyInput) (*elb.CreateAppCookieStickinessPolicyOutput, error) {
	return a.client.CreateAppCookieStickinessPolicyWithContext(ctx, input)
}

func (a *ELBActivities) CreateLBCookieStickinessPolicy(ctx context.Context, input *elb.CreateLBCookieStickinessPolicyInput) (*elb.CreateLBCookieStickinessPolicyOutput, error) {
	return a.client.CreateLBCookieStickinessPolicyWithContext(ctx, input)
}

func (a *ELBActivities) CreateLoadBalancer(ctx context.Context, input *elb.CreateLoadBalancerInput) (*elb.CreateLoadBalancerOutput, error) {
	return a.client.CreateLoadBalancerWithContext(ctx, input)
}

func (a *ELBActivities) CreateLoadBalancerListeners(ctx context.Context, input *elb.CreateLoadBalancerListenersInput) (*elb.CreateLoadBalancerListenersOutput, error) {
	return a.client.CreateLoadBalancerListenersWithContext(ctx, input)
}

func (a *ELBActivities) CreateLoadBalancerPolicy(ctx context.Context, input *elb.CreateLoadBalancerPolicyInput) (*elb.CreateLoadBalancerPolicyOutput, error) {
	return a.client.CreateLoadBalancerPolicyWithContext(ctx, input)
}

func (a *ELBActivities) DeleteLoadBalancer(ctx context.Context, input *elb.DeleteLoadBalancerInput) (*elb.DeleteLoadBalancerOutput, error) {
	return a.client.DeleteLoadBalancerWithContext(ctx, input)
}

func (a *ELBActivities) DeleteLoadBalancerListeners(ctx context.Context, input *elb.DeleteLoadBalancerListenersInput) (*elb.DeleteLoadBalancerListenersOutput, error) {
	return a.client.DeleteLoadBalancerListenersWithContext(ctx, input)
}

func (a *ELBActivities) DeleteLoadBalancerPolicy(ctx context.Context, input *elb.DeleteLoadBalancerPolicyInput) (*elb.DeleteLoadBalancerPolicyOutput, error) {
	return a.client.DeleteLoadBalancerPolicyWithContext(ctx, input)
}

func (a *ELBActivities) DeregisterInstancesFromLoadBalancer(ctx context.Context, input *elb.DeregisterInstancesFromLoadBalancerInput) (*elb.DeregisterInstancesFromLoadBalancerOutput, error) {
	return a.client.DeregisterInstancesFromLoadBalancerWithContext(ctx, input)
}

func (a *ELBActivities) DescribeAccountLimits(ctx context.Context, input *elb.DescribeAccountLimitsInput) (*elb.DescribeAccountLimitsOutput, error) {
	return a.client.DescribeAccountLimitsWithContext(ctx, input)
}

func (a *ELBActivities) DescribeInstanceHealth(ctx context.Context, input *elb.DescribeInstanceHealthInput) (*elb.DescribeInstanceHealthOutput, error) {
	return a.client.DescribeInstanceHealthWithContext(ctx, input)
}

func (a *ELBActivities) DescribeLoadBalancerAttributes(ctx context.Context, input *elb.DescribeLoadBalancerAttributesInput) (*elb.DescribeLoadBalancerAttributesOutput, error) {
	return a.client.DescribeLoadBalancerAttributesWithContext(ctx, input)
}

func (a *ELBActivities) DescribeLoadBalancerPolicies(ctx context.Context, input *elb.DescribeLoadBalancerPoliciesInput) (*elb.DescribeLoadBalancerPoliciesOutput, error) {
	return a.client.DescribeLoadBalancerPoliciesWithContext(ctx, input)
}

func (a *ELBActivities) DescribeLoadBalancerPolicyTypes(ctx context.Context, input *elb.DescribeLoadBalancerPolicyTypesInput) (*elb.DescribeLoadBalancerPolicyTypesOutput, error) {
	return a.client.DescribeLoadBalancerPolicyTypesWithContext(ctx, input)
}

func (a *ELBActivities) DescribeLoadBalancers(ctx context.Context, input *elb.DescribeLoadBalancersInput) (*elb.DescribeLoadBalancersOutput, error) {
	return a.client.DescribeLoadBalancersWithContext(ctx, input)
}

func (a *ELBActivities) DescribeTags(ctx context.Context, input *elb.DescribeTagsInput) (*elb.DescribeTagsOutput, error) {
	return a.client.DescribeTagsWithContext(ctx, input)
}

func (a *ELBActivities) DetachLoadBalancerFromSubnets(ctx context.Context, input *elb.DetachLoadBalancerFromSubnetsInput) (*elb.DetachLoadBalancerFromSubnetsOutput, error) {
	return a.client.DetachLoadBalancerFromSubnetsWithContext(ctx, input)
}

func (a *ELBActivities) DisableAvailabilityZonesForLoadBalancer(ctx context.Context, input *elb.DisableAvailabilityZonesForLoadBalancerInput) (*elb.DisableAvailabilityZonesForLoadBalancerOutput, error) {
	return a.client.DisableAvailabilityZonesForLoadBalancerWithContext(ctx, input)
}

func (a *ELBActivities) EnableAvailabilityZonesForLoadBalancer(ctx context.Context, input *elb.EnableAvailabilityZonesForLoadBalancerInput) (*elb.EnableAvailabilityZonesForLoadBalancerOutput, error) {
	return a.client.EnableAvailabilityZonesForLoadBalancerWithContext(ctx, input)
}

func (a *ELBActivities) ModifyLoadBalancerAttributes(ctx context.Context, input *elb.ModifyLoadBalancerAttributesInput) (*elb.ModifyLoadBalancerAttributesOutput, error) {
	return a.client.ModifyLoadBalancerAttributesWithContext(ctx, input)
}

func (a *ELBActivities) RegisterInstancesWithLoadBalancer(ctx context.Context, input *elb.RegisterInstancesWithLoadBalancerInput) (*elb.RegisterInstancesWithLoadBalancerOutput, error) {
	return a.client.RegisterInstancesWithLoadBalancerWithContext(ctx, input)
}

func (a *ELBActivities) RemoveTags(ctx context.Context, input *elb.RemoveTagsInput) (*elb.RemoveTagsOutput, error) {
	return a.client.RemoveTagsWithContext(ctx, input)
}

func (a *ELBActivities) SetLoadBalancerListenerSSLCertificate(ctx context.Context, input *elb.SetLoadBalancerListenerSSLCertificateInput) (*elb.SetLoadBalancerListenerSSLCertificateOutput, error) {
	return a.client.SetLoadBalancerListenerSSLCertificateWithContext(ctx, input)
}

func (a *ELBActivities) SetLoadBalancerPoliciesForBackendServer(ctx context.Context, input *elb.SetLoadBalancerPoliciesForBackendServerInput) (*elb.SetLoadBalancerPoliciesForBackendServerOutput, error) {
	return a.client.SetLoadBalancerPoliciesForBackendServerWithContext(ctx, input)
}

func (a *ELBActivities) SetLoadBalancerPoliciesOfListener(ctx context.Context, input *elb.SetLoadBalancerPoliciesOfListenerInput) (*elb.SetLoadBalancerPoliciesOfListenerOutput, error) {
	return a.client.SetLoadBalancerPoliciesOfListenerWithContext(ctx, input)
}

func (a *ELBActivities) WaitUntilAnyInstanceInService(ctx context.Context, input *elb.DescribeInstanceHealthInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilAnyInstanceInServiceWithContext(ctx, input, options...)
	})
}

func (a *ELBActivities) WaitUntilInstanceDeregistered(ctx context.Context, input *elb.DescribeInstanceHealthInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilInstanceDeregisteredWithContext(ctx, input, options...)
	})
}

func (a *ELBActivities) WaitUntilInstanceInService(ctx context.Context, input *elb.DescribeInstanceHealthInput) error {
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return a.client.WaitUntilInstanceInServiceWithContext(ctx, input, options...)
	})
}
