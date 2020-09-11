package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/waf"
	"github.com/aws/aws-sdk-go/service/wafregional"
	"github.com/aws/aws-sdk-go/service/wafregional/wafregionaliface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type WAFRegionalActivities struct {
	client wafregionaliface.WAFRegionalAPI
}

func NewWAFRegionalActivities(session *session.Session, config ...*aws.Config) *WAFRegionalActivities {
	client := wafregional.New(session, config...)
	return &WAFRegionalActivities{client: client}
}

func (a *WAFRegionalActivities) AssociateWebACL(ctx context.Context, input *wafregional.AssociateWebACLInput) (*wafregional.AssociateWebACLOutput, error) {
	return a.client.AssociateWebACLWithContext(ctx, input)
}

func (a *WAFRegionalActivities) CreateByteMatchSet(ctx context.Context, input *waf.CreateByteMatchSetInput) (*waf.CreateByteMatchSetOutput, error) {
	return a.client.CreateByteMatchSetWithContext(ctx, input)
}

func (a *WAFRegionalActivities) CreateGeoMatchSet(ctx context.Context, input *waf.CreateGeoMatchSetInput) (*waf.CreateGeoMatchSetOutput, error) {
	return a.client.CreateGeoMatchSetWithContext(ctx, input)
}

func (a *WAFRegionalActivities) CreateIPSet(ctx context.Context, input *waf.CreateIPSetInput) (*waf.CreateIPSetOutput, error) {
	return a.client.CreateIPSetWithContext(ctx, input)
}

func (a *WAFRegionalActivities) CreateRateBasedRule(ctx context.Context, input *waf.CreateRateBasedRuleInput) (*waf.CreateRateBasedRuleOutput, error) {
	return a.client.CreateRateBasedRuleWithContext(ctx, input)
}

func (a *WAFRegionalActivities) CreateRegexMatchSet(ctx context.Context, input *waf.CreateRegexMatchSetInput) (*waf.CreateRegexMatchSetOutput, error) {
	return a.client.CreateRegexMatchSetWithContext(ctx, input)
}

func (a *WAFRegionalActivities) CreateRegexPatternSet(ctx context.Context, input *waf.CreateRegexPatternSetInput) (*waf.CreateRegexPatternSetOutput, error) {
	return a.client.CreateRegexPatternSetWithContext(ctx, input)
}

func (a *WAFRegionalActivities) CreateRule(ctx context.Context, input *waf.CreateRuleInput) (*waf.CreateRuleOutput, error) {
	return a.client.CreateRuleWithContext(ctx, input)
}

func (a *WAFRegionalActivities) CreateRuleGroup(ctx context.Context, input *waf.CreateRuleGroupInput) (*waf.CreateRuleGroupOutput, error) {
	return a.client.CreateRuleGroupWithContext(ctx, input)
}

func (a *WAFRegionalActivities) CreateSizeConstraintSet(ctx context.Context, input *waf.CreateSizeConstraintSetInput) (*waf.CreateSizeConstraintSetOutput, error) {
	return a.client.CreateSizeConstraintSetWithContext(ctx, input)
}

func (a *WAFRegionalActivities) CreateSqlInjectionMatchSet(ctx context.Context, input *waf.CreateSqlInjectionMatchSetInput) (*waf.CreateSqlInjectionMatchSetOutput, error) {
	return a.client.CreateSqlInjectionMatchSetWithContext(ctx, input)
}

func (a *WAFRegionalActivities) CreateWebACL(ctx context.Context, input *waf.CreateWebACLInput) (*waf.CreateWebACLOutput, error) {
	return a.client.CreateWebACLWithContext(ctx, input)
}

func (a *WAFRegionalActivities) CreateWebACLMigrationStack(ctx context.Context, input *waf.CreateWebACLMigrationStackInput) (*waf.CreateWebACLMigrationStackOutput, error) {
	return a.client.CreateWebACLMigrationStackWithContext(ctx, input)
}

func (a *WAFRegionalActivities) CreateXssMatchSet(ctx context.Context, input *waf.CreateXssMatchSetInput) (*waf.CreateXssMatchSetOutput, error) {
	return a.client.CreateXssMatchSetWithContext(ctx, input)
}

func (a *WAFRegionalActivities) DeleteByteMatchSet(ctx context.Context, input *waf.DeleteByteMatchSetInput) (*waf.DeleteByteMatchSetOutput, error) {
	return a.client.DeleteByteMatchSetWithContext(ctx, input)
}

func (a *WAFRegionalActivities) DeleteGeoMatchSet(ctx context.Context, input *waf.DeleteGeoMatchSetInput) (*waf.DeleteGeoMatchSetOutput, error) {
	return a.client.DeleteGeoMatchSetWithContext(ctx, input)
}

func (a *WAFRegionalActivities) DeleteIPSet(ctx context.Context, input *waf.DeleteIPSetInput) (*waf.DeleteIPSetOutput, error) {
	return a.client.DeleteIPSetWithContext(ctx, input)
}

func (a *WAFRegionalActivities) DeleteLoggingConfiguration(ctx context.Context, input *waf.DeleteLoggingConfigurationInput) (*waf.DeleteLoggingConfigurationOutput, error) {
	return a.client.DeleteLoggingConfigurationWithContext(ctx, input)
}

func (a *WAFRegionalActivities) DeletePermissionPolicy(ctx context.Context, input *waf.DeletePermissionPolicyInput) (*waf.DeletePermissionPolicyOutput, error) {
	return a.client.DeletePermissionPolicyWithContext(ctx, input)
}

func (a *WAFRegionalActivities) DeleteRateBasedRule(ctx context.Context, input *waf.DeleteRateBasedRuleInput) (*waf.DeleteRateBasedRuleOutput, error) {
	return a.client.DeleteRateBasedRuleWithContext(ctx, input)
}

func (a *WAFRegionalActivities) DeleteRegexMatchSet(ctx context.Context, input *waf.DeleteRegexMatchSetInput) (*waf.DeleteRegexMatchSetOutput, error) {
	return a.client.DeleteRegexMatchSetWithContext(ctx, input)
}

func (a *WAFRegionalActivities) DeleteRegexPatternSet(ctx context.Context, input *waf.DeleteRegexPatternSetInput) (*waf.DeleteRegexPatternSetOutput, error) {
	return a.client.DeleteRegexPatternSetWithContext(ctx, input)
}

func (a *WAFRegionalActivities) DeleteRule(ctx context.Context, input *waf.DeleteRuleInput) (*waf.DeleteRuleOutput, error) {
	return a.client.DeleteRuleWithContext(ctx, input)
}

func (a *WAFRegionalActivities) DeleteRuleGroup(ctx context.Context, input *waf.DeleteRuleGroupInput) (*waf.DeleteRuleGroupOutput, error) {
	return a.client.DeleteRuleGroupWithContext(ctx, input)
}

func (a *WAFRegionalActivities) DeleteSizeConstraintSet(ctx context.Context, input *waf.DeleteSizeConstraintSetInput) (*waf.DeleteSizeConstraintSetOutput, error) {
	return a.client.DeleteSizeConstraintSetWithContext(ctx, input)
}

func (a *WAFRegionalActivities) DeleteSqlInjectionMatchSet(ctx context.Context, input *waf.DeleteSqlInjectionMatchSetInput) (*waf.DeleteSqlInjectionMatchSetOutput, error) {
	return a.client.DeleteSqlInjectionMatchSetWithContext(ctx, input)
}

func (a *WAFRegionalActivities) DeleteWebACL(ctx context.Context, input *waf.DeleteWebACLInput) (*waf.DeleteWebACLOutput, error) {
	return a.client.DeleteWebACLWithContext(ctx, input)
}

func (a *WAFRegionalActivities) DeleteXssMatchSet(ctx context.Context, input *waf.DeleteXssMatchSetInput) (*waf.DeleteXssMatchSetOutput, error) {
	return a.client.DeleteXssMatchSetWithContext(ctx, input)
}

func (a *WAFRegionalActivities) DisassociateWebACL(ctx context.Context, input *wafregional.DisassociateWebACLInput) (*wafregional.DisassociateWebACLOutput, error) {
	return a.client.DisassociateWebACLWithContext(ctx, input)
}

func (a *WAFRegionalActivities) GetByteMatchSet(ctx context.Context, input *waf.GetByteMatchSetInput) (*waf.GetByteMatchSetOutput, error) {
	return a.client.GetByteMatchSetWithContext(ctx, input)
}

func (a *WAFRegionalActivities) GetChangeToken(ctx context.Context, input *waf.GetChangeTokenInput) (*waf.GetChangeTokenOutput, error) {
	return a.client.GetChangeTokenWithContext(ctx, input)
}

func (a *WAFRegionalActivities) GetChangeTokenStatus(ctx context.Context, input *waf.GetChangeTokenStatusInput) (*waf.GetChangeTokenStatusOutput, error) {
	return a.client.GetChangeTokenStatusWithContext(ctx, input)
}

func (a *WAFRegionalActivities) GetGeoMatchSet(ctx context.Context, input *waf.GetGeoMatchSetInput) (*waf.GetGeoMatchSetOutput, error) {
	return a.client.GetGeoMatchSetWithContext(ctx, input)
}

func (a *WAFRegionalActivities) GetIPSet(ctx context.Context, input *waf.GetIPSetInput) (*waf.GetIPSetOutput, error) {
	return a.client.GetIPSetWithContext(ctx, input)
}

func (a *WAFRegionalActivities) GetLoggingConfiguration(ctx context.Context, input *waf.GetLoggingConfigurationInput) (*waf.GetLoggingConfigurationOutput, error) {
	return a.client.GetLoggingConfigurationWithContext(ctx, input)
}

func (a *WAFRegionalActivities) GetPermissionPolicy(ctx context.Context, input *waf.GetPermissionPolicyInput) (*waf.GetPermissionPolicyOutput, error) {
	return a.client.GetPermissionPolicyWithContext(ctx, input)
}

func (a *WAFRegionalActivities) GetRateBasedRule(ctx context.Context, input *waf.GetRateBasedRuleInput) (*waf.GetRateBasedRuleOutput, error) {
	return a.client.GetRateBasedRuleWithContext(ctx, input)
}

func (a *WAFRegionalActivities) GetRateBasedRuleManagedKeys(ctx context.Context, input *waf.GetRateBasedRuleManagedKeysInput) (*waf.GetRateBasedRuleManagedKeysOutput, error) {
	return a.client.GetRateBasedRuleManagedKeysWithContext(ctx, input)
}

func (a *WAFRegionalActivities) GetRegexMatchSet(ctx context.Context, input *waf.GetRegexMatchSetInput) (*waf.GetRegexMatchSetOutput, error) {
	return a.client.GetRegexMatchSetWithContext(ctx, input)
}

func (a *WAFRegionalActivities) GetRegexPatternSet(ctx context.Context, input *waf.GetRegexPatternSetInput) (*waf.GetRegexPatternSetOutput, error) {
	return a.client.GetRegexPatternSetWithContext(ctx, input)
}

func (a *WAFRegionalActivities) GetRule(ctx context.Context, input *waf.GetRuleInput) (*waf.GetRuleOutput, error) {
	return a.client.GetRuleWithContext(ctx, input)
}

func (a *WAFRegionalActivities) GetRuleGroup(ctx context.Context, input *waf.GetRuleGroupInput) (*waf.GetRuleGroupOutput, error) {
	return a.client.GetRuleGroupWithContext(ctx, input)
}

func (a *WAFRegionalActivities) GetSampledRequests(ctx context.Context, input *waf.GetSampledRequestsInput) (*waf.GetSampledRequestsOutput, error) {
	return a.client.GetSampledRequestsWithContext(ctx, input)
}

func (a *WAFRegionalActivities) GetSizeConstraintSet(ctx context.Context, input *waf.GetSizeConstraintSetInput) (*waf.GetSizeConstraintSetOutput, error) {
	return a.client.GetSizeConstraintSetWithContext(ctx, input)
}

func (a *WAFRegionalActivities) GetSqlInjectionMatchSet(ctx context.Context, input *waf.GetSqlInjectionMatchSetInput) (*waf.GetSqlInjectionMatchSetOutput, error) {
	return a.client.GetSqlInjectionMatchSetWithContext(ctx, input)
}

func (a *WAFRegionalActivities) GetWebACL(ctx context.Context, input *waf.GetWebACLInput) (*waf.GetWebACLOutput, error) {
	return a.client.GetWebACLWithContext(ctx, input)
}

func (a *WAFRegionalActivities) GetWebACLForResource(ctx context.Context, input *wafregional.GetWebACLForResourceInput) (*wafregional.GetWebACLForResourceOutput, error) {
	return a.client.GetWebACLForResourceWithContext(ctx, input)
}

func (a *WAFRegionalActivities) GetXssMatchSet(ctx context.Context, input *waf.GetXssMatchSetInput) (*waf.GetXssMatchSetOutput, error) {
	return a.client.GetXssMatchSetWithContext(ctx, input)
}

func (a *WAFRegionalActivities) ListActivatedRulesInRuleGroup(ctx context.Context, input *waf.ListActivatedRulesInRuleGroupInput) (*waf.ListActivatedRulesInRuleGroupOutput, error) {
	return a.client.ListActivatedRulesInRuleGroupWithContext(ctx, input)
}

func (a *WAFRegionalActivities) ListByteMatchSets(ctx context.Context, input *waf.ListByteMatchSetsInput) (*waf.ListByteMatchSetsOutput, error) {
	return a.client.ListByteMatchSetsWithContext(ctx, input)
}

func (a *WAFRegionalActivities) ListGeoMatchSets(ctx context.Context, input *waf.ListGeoMatchSetsInput) (*waf.ListGeoMatchSetsOutput, error) {
	return a.client.ListGeoMatchSetsWithContext(ctx, input)
}

func (a *WAFRegionalActivities) ListIPSets(ctx context.Context, input *waf.ListIPSetsInput) (*waf.ListIPSetsOutput, error) {
	return a.client.ListIPSetsWithContext(ctx, input)
}

func (a *WAFRegionalActivities) ListLoggingConfigurations(ctx context.Context, input *waf.ListLoggingConfigurationsInput) (*waf.ListLoggingConfigurationsOutput, error) {
	return a.client.ListLoggingConfigurationsWithContext(ctx, input)
}

func (a *WAFRegionalActivities) ListRateBasedRules(ctx context.Context, input *waf.ListRateBasedRulesInput) (*waf.ListRateBasedRulesOutput, error) {
	return a.client.ListRateBasedRulesWithContext(ctx, input)
}

func (a *WAFRegionalActivities) ListRegexMatchSets(ctx context.Context, input *waf.ListRegexMatchSetsInput) (*waf.ListRegexMatchSetsOutput, error) {
	return a.client.ListRegexMatchSetsWithContext(ctx, input)
}

func (a *WAFRegionalActivities) ListRegexPatternSets(ctx context.Context, input *waf.ListRegexPatternSetsInput) (*waf.ListRegexPatternSetsOutput, error) {
	return a.client.ListRegexPatternSetsWithContext(ctx, input)
}

func (a *WAFRegionalActivities) ListResourcesForWebACL(ctx context.Context, input *wafregional.ListResourcesForWebACLInput) (*wafregional.ListResourcesForWebACLOutput, error) {
	return a.client.ListResourcesForWebACLWithContext(ctx, input)
}

func (a *WAFRegionalActivities) ListRuleGroups(ctx context.Context, input *waf.ListRuleGroupsInput) (*waf.ListRuleGroupsOutput, error) {
	return a.client.ListRuleGroupsWithContext(ctx, input)
}

func (a *WAFRegionalActivities) ListRules(ctx context.Context, input *waf.ListRulesInput) (*waf.ListRulesOutput, error) {
	return a.client.ListRulesWithContext(ctx, input)
}

func (a *WAFRegionalActivities) ListSizeConstraintSets(ctx context.Context, input *waf.ListSizeConstraintSetsInput) (*waf.ListSizeConstraintSetsOutput, error) {
	return a.client.ListSizeConstraintSetsWithContext(ctx, input)
}

func (a *WAFRegionalActivities) ListSqlInjectionMatchSets(ctx context.Context, input *waf.ListSqlInjectionMatchSetsInput) (*waf.ListSqlInjectionMatchSetsOutput, error) {
	return a.client.ListSqlInjectionMatchSetsWithContext(ctx, input)
}

func (a *WAFRegionalActivities) ListSubscribedRuleGroups(ctx context.Context, input *waf.ListSubscribedRuleGroupsInput) (*waf.ListSubscribedRuleGroupsOutput, error) {
	return a.client.ListSubscribedRuleGroupsWithContext(ctx, input)
}

func (a *WAFRegionalActivities) ListTagsForResource(ctx context.Context, input *waf.ListTagsForResourceInput) (*waf.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *WAFRegionalActivities) ListWebACLs(ctx context.Context, input *waf.ListWebACLsInput) (*waf.ListWebACLsOutput, error) {
	return a.client.ListWebACLsWithContext(ctx, input)
}

func (a *WAFRegionalActivities) ListXssMatchSets(ctx context.Context, input *waf.ListXssMatchSetsInput) (*waf.ListXssMatchSetsOutput, error) {
	return a.client.ListXssMatchSetsWithContext(ctx, input)
}

func (a *WAFRegionalActivities) PutLoggingConfiguration(ctx context.Context, input *waf.PutLoggingConfigurationInput) (*waf.PutLoggingConfigurationOutput, error) {
	return a.client.PutLoggingConfigurationWithContext(ctx, input)
}

func (a *WAFRegionalActivities) PutPermissionPolicy(ctx context.Context, input *waf.PutPermissionPolicyInput) (*waf.PutPermissionPolicyOutput, error) {
	return a.client.PutPermissionPolicyWithContext(ctx, input)
}

func (a *WAFRegionalActivities) TagResource(ctx context.Context, input *waf.TagResourceInput) (*waf.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *WAFRegionalActivities) UntagResource(ctx context.Context, input *waf.UntagResourceInput) (*waf.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}

func (a *WAFRegionalActivities) UpdateByteMatchSet(ctx context.Context, input *waf.UpdateByteMatchSetInput) (*waf.UpdateByteMatchSetOutput, error) {
	return a.client.UpdateByteMatchSetWithContext(ctx, input)
}

func (a *WAFRegionalActivities) UpdateGeoMatchSet(ctx context.Context, input *waf.UpdateGeoMatchSetInput) (*waf.UpdateGeoMatchSetOutput, error) {
	return a.client.UpdateGeoMatchSetWithContext(ctx, input)
}

func (a *WAFRegionalActivities) UpdateIPSet(ctx context.Context, input *waf.UpdateIPSetInput) (*waf.UpdateIPSetOutput, error) {
	return a.client.UpdateIPSetWithContext(ctx, input)
}

func (a *WAFRegionalActivities) UpdateRateBasedRule(ctx context.Context, input *waf.UpdateRateBasedRuleInput) (*waf.UpdateRateBasedRuleOutput, error) {
	return a.client.UpdateRateBasedRuleWithContext(ctx, input)
}

func (a *WAFRegionalActivities) UpdateRegexMatchSet(ctx context.Context, input *waf.UpdateRegexMatchSetInput) (*waf.UpdateRegexMatchSetOutput, error) {
	return a.client.UpdateRegexMatchSetWithContext(ctx, input)
}

func (a *WAFRegionalActivities) UpdateRegexPatternSet(ctx context.Context, input *waf.UpdateRegexPatternSetInput) (*waf.UpdateRegexPatternSetOutput, error) {
	return a.client.UpdateRegexPatternSetWithContext(ctx, input)
}

func (a *WAFRegionalActivities) UpdateRule(ctx context.Context, input *waf.UpdateRuleInput) (*waf.UpdateRuleOutput, error) {
	return a.client.UpdateRuleWithContext(ctx, input)
}

func (a *WAFRegionalActivities) UpdateRuleGroup(ctx context.Context, input *waf.UpdateRuleGroupInput) (*waf.UpdateRuleGroupOutput, error) {
	return a.client.UpdateRuleGroupWithContext(ctx, input)
}

func (a *WAFRegionalActivities) UpdateSizeConstraintSet(ctx context.Context, input *waf.UpdateSizeConstraintSetInput) (*waf.UpdateSizeConstraintSetOutput, error) {
	return a.client.UpdateSizeConstraintSetWithContext(ctx, input)
}

func (a *WAFRegionalActivities) UpdateSqlInjectionMatchSet(ctx context.Context, input *waf.UpdateSqlInjectionMatchSetInput) (*waf.UpdateSqlInjectionMatchSetOutput, error) {
	return a.client.UpdateSqlInjectionMatchSetWithContext(ctx, input)
}

func (a *WAFRegionalActivities) UpdateWebACL(ctx context.Context, input *waf.UpdateWebACLInput) (*waf.UpdateWebACLOutput, error) {
	return a.client.UpdateWebACLWithContext(ctx, input)
}

func (a *WAFRegionalActivities) UpdateXssMatchSet(ctx context.Context, input *waf.UpdateXssMatchSetInput) (*waf.UpdateXssMatchSetOutput, error) {
	return a.client.UpdateXssMatchSetWithContext(ctx, input)
}
