package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/ses"
	"go.temporal.io/sdk/workflow"
)

type SESClient interface {
    CloneReceiptRuleSet(ctx workflow.Context, input *ses.CloneReceiptRuleSetInput) (*ses.CloneReceiptRuleSetOutput, error)
    CloneReceiptRuleSetAsync(ctx workflow.Context, input *ses.CloneReceiptRuleSetInput) *SesCloneReceiptRuleSetResult

    CreateConfigurationSet(ctx workflow.Context, input *ses.CreateConfigurationSetInput) (*ses.CreateConfigurationSetOutput, error)
    CreateConfigurationSetAsync(ctx workflow.Context, input *ses.CreateConfigurationSetInput) *SesCreateConfigurationSetResult

    CreateConfigurationSetEventDestination(ctx workflow.Context, input *ses.CreateConfigurationSetEventDestinationInput) (*ses.CreateConfigurationSetEventDestinationOutput, error)
    CreateConfigurationSetEventDestinationAsync(ctx workflow.Context, input *ses.CreateConfigurationSetEventDestinationInput) *SesCreateConfigurationSetEventDestinationResult

    CreateConfigurationSetTrackingOptions(ctx workflow.Context, input *ses.CreateConfigurationSetTrackingOptionsInput) (*ses.CreateConfigurationSetTrackingOptionsOutput, error)
    CreateConfigurationSetTrackingOptionsAsync(ctx workflow.Context, input *ses.CreateConfigurationSetTrackingOptionsInput) *SesCreateConfigurationSetTrackingOptionsResult

    CreateCustomVerificationEmailTemplate(ctx workflow.Context, input *ses.CreateCustomVerificationEmailTemplateInput) (*ses.CreateCustomVerificationEmailTemplateOutput, error)
    CreateCustomVerificationEmailTemplateAsync(ctx workflow.Context, input *ses.CreateCustomVerificationEmailTemplateInput) *SesCreateCustomVerificationEmailTemplateResult

    CreateReceiptFilter(ctx workflow.Context, input *ses.CreateReceiptFilterInput) (*ses.CreateReceiptFilterOutput, error)
    CreateReceiptFilterAsync(ctx workflow.Context, input *ses.CreateReceiptFilterInput) *SesCreateReceiptFilterResult

    CreateReceiptRule(ctx workflow.Context, input *ses.CreateReceiptRuleInput) (*ses.CreateReceiptRuleOutput, error)
    CreateReceiptRuleAsync(ctx workflow.Context, input *ses.CreateReceiptRuleInput) *SesCreateReceiptRuleResult

    CreateReceiptRuleSet(ctx workflow.Context, input *ses.CreateReceiptRuleSetInput) (*ses.CreateReceiptRuleSetOutput, error)
    CreateReceiptRuleSetAsync(ctx workflow.Context, input *ses.CreateReceiptRuleSetInput) *SesCreateReceiptRuleSetResult

    CreateTemplate(ctx workflow.Context, input *ses.CreateTemplateInput) (*ses.CreateTemplateOutput, error)
    CreateTemplateAsync(ctx workflow.Context, input *ses.CreateTemplateInput) *SesCreateTemplateResult

    DeleteConfigurationSet(ctx workflow.Context, input *ses.DeleteConfigurationSetInput) (*ses.DeleteConfigurationSetOutput, error)
    DeleteConfigurationSetAsync(ctx workflow.Context, input *ses.DeleteConfigurationSetInput) *SesDeleteConfigurationSetResult

    DeleteConfigurationSetEventDestination(ctx workflow.Context, input *ses.DeleteConfigurationSetEventDestinationInput) (*ses.DeleteConfigurationSetEventDestinationOutput, error)
    DeleteConfigurationSetEventDestinationAsync(ctx workflow.Context, input *ses.DeleteConfigurationSetEventDestinationInput) *SesDeleteConfigurationSetEventDestinationResult

    DeleteConfigurationSetTrackingOptions(ctx workflow.Context, input *ses.DeleteConfigurationSetTrackingOptionsInput) (*ses.DeleteConfigurationSetTrackingOptionsOutput, error)
    DeleteConfigurationSetTrackingOptionsAsync(ctx workflow.Context, input *ses.DeleteConfigurationSetTrackingOptionsInput) *SesDeleteConfigurationSetTrackingOptionsResult

    DeleteCustomVerificationEmailTemplate(ctx workflow.Context, input *ses.DeleteCustomVerificationEmailTemplateInput) (*ses.DeleteCustomVerificationEmailTemplateOutput, error)
    DeleteCustomVerificationEmailTemplateAsync(ctx workflow.Context, input *ses.DeleteCustomVerificationEmailTemplateInput) *SesDeleteCustomVerificationEmailTemplateResult

    DeleteIdentity(ctx workflow.Context, input *ses.DeleteIdentityInput) (*ses.DeleteIdentityOutput, error)
    DeleteIdentityAsync(ctx workflow.Context, input *ses.DeleteIdentityInput) *SesDeleteIdentityResult

    DeleteIdentityPolicy(ctx workflow.Context, input *ses.DeleteIdentityPolicyInput) (*ses.DeleteIdentityPolicyOutput, error)
    DeleteIdentityPolicyAsync(ctx workflow.Context, input *ses.DeleteIdentityPolicyInput) *SesDeleteIdentityPolicyResult

    DeleteReceiptFilter(ctx workflow.Context, input *ses.DeleteReceiptFilterInput) (*ses.DeleteReceiptFilterOutput, error)
    DeleteReceiptFilterAsync(ctx workflow.Context, input *ses.DeleteReceiptFilterInput) *SesDeleteReceiptFilterResult

    DeleteReceiptRule(ctx workflow.Context, input *ses.DeleteReceiptRuleInput) (*ses.DeleteReceiptRuleOutput, error)
    DeleteReceiptRuleAsync(ctx workflow.Context, input *ses.DeleteReceiptRuleInput) *SesDeleteReceiptRuleResult

    DeleteReceiptRuleSet(ctx workflow.Context, input *ses.DeleteReceiptRuleSetInput) (*ses.DeleteReceiptRuleSetOutput, error)
    DeleteReceiptRuleSetAsync(ctx workflow.Context, input *ses.DeleteReceiptRuleSetInput) *SesDeleteReceiptRuleSetResult

    DeleteTemplate(ctx workflow.Context, input *ses.DeleteTemplateInput) (*ses.DeleteTemplateOutput, error)
    DeleteTemplateAsync(ctx workflow.Context, input *ses.DeleteTemplateInput) *SesDeleteTemplateResult

    DeleteVerifiedEmailAddress(ctx workflow.Context, input *ses.DeleteVerifiedEmailAddressInput) (*ses.DeleteVerifiedEmailAddressOutput, error)
    DeleteVerifiedEmailAddressAsync(ctx workflow.Context, input *ses.DeleteVerifiedEmailAddressInput) *SesDeleteVerifiedEmailAddressResult

    DescribeActiveReceiptRuleSet(ctx workflow.Context, input *ses.DescribeActiveReceiptRuleSetInput) (*ses.DescribeActiveReceiptRuleSetOutput, error)
    DescribeActiveReceiptRuleSetAsync(ctx workflow.Context, input *ses.DescribeActiveReceiptRuleSetInput) *SesDescribeActiveReceiptRuleSetResult

    DescribeConfigurationSet(ctx workflow.Context, input *ses.DescribeConfigurationSetInput) (*ses.DescribeConfigurationSetOutput, error)
    DescribeConfigurationSetAsync(ctx workflow.Context, input *ses.DescribeConfigurationSetInput) *SesDescribeConfigurationSetResult

    DescribeReceiptRule(ctx workflow.Context, input *ses.DescribeReceiptRuleInput) (*ses.DescribeReceiptRuleOutput, error)
    DescribeReceiptRuleAsync(ctx workflow.Context, input *ses.DescribeReceiptRuleInput) *SesDescribeReceiptRuleResult

    DescribeReceiptRuleSet(ctx workflow.Context, input *ses.DescribeReceiptRuleSetInput) (*ses.DescribeReceiptRuleSetOutput, error)
    DescribeReceiptRuleSetAsync(ctx workflow.Context, input *ses.DescribeReceiptRuleSetInput) *SesDescribeReceiptRuleSetResult

    GetAccountSendingEnabled(ctx workflow.Context, input *ses.GetAccountSendingEnabledInput) (*ses.GetAccountSendingEnabledOutput, error)
    GetAccountSendingEnabledAsync(ctx workflow.Context, input *ses.GetAccountSendingEnabledInput) *SesGetAccountSendingEnabledResult

    GetCustomVerificationEmailTemplate(ctx workflow.Context, input *ses.GetCustomVerificationEmailTemplateInput) (*ses.GetCustomVerificationEmailTemplateOutput, error)
    GetCustomVerificationEmailTemplateAsync(ctx workflow.Context, input *ses.GetCustomVerificationEmailTemplateInput) *SesGetCustomVerificationEmailTemplateResult

    GetIdentityDkimAttributes(ctx workflow.Context, input *ses.GetIdentityDkimAttributesInput) (*ses.GetIdentityDkimAttributesOutput, error)
    GetIdentityDkimAttributesAsync(ctx workflow.Context, input *ses.GetIdentityDkimAttributesInput) *SesGetIdentityDkimAttributesResult

    GetIdentityMailFromDomainAttributes(ctx workflow.Context, input *ses.GetIdentityMailFromDomainAttributesInput) (*ses.GetIdentityMailFromDomainAttributesOutput, error)
    GetIdentityMailFromDomainAttributesAsync(ctx workflow.Context, input *ses.GetIdentityMailFromDomainAttributesInput) *SesGetIdentityMailFromDomainAttributesResult

    GetIdentityNotificationAttributes(ctx workflow.Context, input *ses.GetIdentityNotificationAttributesInput) (*ses.GetIdentityNotificationAttributesOutput, error)
    GetIdentityNotificationAttributesAsync(ctx workflow.Context, input *ses.GetIdentityNotificationAttributesInput) *SesGetIdentityNotificationAttributesResult

    GetIdentityPolicies(ctx workflow.Context, input *ses.GetIdentityPoliciesInput) (*ses.GetIdentityPoliciesOutput, error)
    GetIdentityPoliciesAsync(ctx workflow.Context, input *ses.GetIdentityPoliciesInput) *SesGetIdentityPoliciesResult

    GetIdentityVerificationAttributes(ctx workflow.Context, input *ses.GetIdentityVerificationAttributesInput) (*ses.GetIdentityVerificationAttributesOutput, error)
    GetIdentityVerificationAttributesAsync(ctx workflow.Context, input *ses.GetIdentityVerificationAttributesInput) *SesGetIdentityVerificationAttributesResult

    GetSendQuota(ctx workflow.Context, input *ses.GetSendQuotaInput) (*ses.GetSendQuotaOutput, error)
    GetSendQuotaAsync(ctx workflow.Context, input *ses.GetSendQuotaInput) *SesGetSendQuotaResult

    GetSendStatistics(ctx workflow.Context, input *ses.GetSendStatisticsInput) (*ses.GetSendStatisticsOutput, error)
    GetSendStatisticsAsync(ctx workflow.Context, input *ses.GetSendStatisticsInput) *SesGetSendStatisticsResult

    GetTemplate(ctx workflow.Context, input *ses.GetTemplateInput) (*ses.GetTemplateOutput, error)
    GetTemplateAsync(ctx workflow.Context, input *ses.GetTemplateInput) *SesGetTemplateResult

    ListConfigurationSets(ctx workflow.Context, input *ses.ListConfigurationSetsInput) (*ses.ListConfigurationSetsOutput, error)
    ListConfigurationSetsAsync(ctx workflow.Context, input *ses.ListConfigurationSetsInput) *SesListConfigurationSetsResult

    ListCustomVerificationEmailTemplates(ctx workflow.Context, input *ses.ListCustomVerificationEmailTemplatesInput) (*ses.ListCustomVerificationEmailTemplatesOutput, error)
    ListCustomVerificationEmailTemplatesAsync(ctx workflow.Context, input *ses.ListCustomVerificationEmailTemplatesInput) *SesListCustomVerificationEmailTemplatesResult

    ListIdentities(ctx workflow.Context, input *ses.ListIdentitiesInput) (*ses.ListIdentitiesOutput, error)
    ListIdentitiesAsync(ctx workflow.Context, input *ses.ListIdentitiesInput) *SesListIdentitiesResult

    ListIdentityPolicies(ctx workflow.Context, input *ses.ListIdentityPoliciesInput) (*ses.ListIdentityPoliciesOutput, error)
    ListIdentityPoliciesAsync(ctx workflow.Context, input *ses.ListIdentityPoliciesInput) *SesListIdentityPoliciesResult

    ListReceiptFilters(ctx workflow.Context, input *ses.ListReceiptFiltersInput) (*ses.ListReceiptFiltersOutput, error)
    ListReceiptFiltersAsync(ctx workflow.Context, input *ses.ListReceiptFiltersInput) *SesListReceiptFiltersResult

    ListReceiptRuleSets(ctx workflow.Context, input *ses.ListReceiptRuleSetsInput) (*ses.ListReceiptRuleSetsOutput, error)
    ListReceiptRuleSetsAsync(ctx workflow.Context, input *ses.ListReceiptRuleSetsInput) *SesListReceiptRuleSetsResult

    ListTemplates(ctx workflow.Context, input *ses.ListTemplatesInput) (*ses.ListTemplatesOutput, error)
    ListTemplatesAsync(ctx workflow.Context, input *ses.ListTemplatesInput) *SesListTemplatesResult

    ListVerifiedEmailAddresses(ctx workflow.Context, input *ses.ListVerifiedEmailAddressesInput) (*ses.ListVerifiedEmailAddressesOutput, error)
    ListVerifiedEmailAddressesAsync(ctx workflow.Context, input *ses.ListVerifiedEmailAddressesInput) *SesListVerifiedEmailAddressesResult

    PutConfigurationSetDeliveryOptions(ctx workflow.Context, input *ses.PutConfigurationSetDeliveryOptionsInput) (*ses.PutConfigurationSetDeliveryOptionsOutput, error)
    PutConfigurationSetDeliveryOptionsAsync(ctx workflow.Context, input *ses.PutConfigurationSetDeliveryOptionsInput) *SesPutConfigurationSetDeliveryOptionsResult

    PutIdentityPolicy(ctx workflow.Context, input *ses.PutIdentityPolicyInput) (*ses.PutIdentityPolicyOutput, error)
    PutIdentityPolicyAsync(ctx workflow.Context, input *ses.PutIdentityPolicyInput) *SesPutIdentityPolicyResult

    ReorderReceiptRuleSet(ctx workflow.Context, input *ses.ReorderReceiptRuleSetInput) (*ses.ReorderReceiptRuleSetOutput, error)
    ReorderReceiptRuleSetAsync(ctx workflow.Context, input *ses.ReorderReceiptRuleSetInput) *SesReorderReceiptRuleSetResult

    SendBounce(ctx workflow.Context, input *ses.SendBounceInput) (*ses.SendBounceOutput, error)
    SendBounceAsync(ctx workflow.Context, input *ses.SendBounceInput) *SesSendBounceResult

    SendBulkTemplatedEmail(ctx workflow.Context, input *ses.SendBulkTemplatedEmailInput) (*ses.SendBulkTemplatedEmailOutput, error)
    SendBulkTemplatedEmailAsync(ctx workflow.Context, input *ses.SendBulkTemplatedEmailInput) *SesSendBulkTemplatedEmailResult

    SendCustomVerificationEmail(ctx workflow.Context, input *ses.SendCustomVerificationEmailInput) (*ses.SendCustomVerificationEmailOutput, error)
    SendCustomVerificationEmailAsync(ctx workflow.Context, input *ses.SendCustomVerificationEmailInput) *SesSendCustomVerificationEmailResult

    SendEmail(ctx workflow.Context, input *ses.SendEmailInput) (*ses.SendEmailOutput, error)
    SendEmailAsync(ctx workflow.Context, input *ses.SendEmailInput) *SesSendEmailResult

    SendRawEmail(ctx workflow.Context, input *ses.SendRawEmailInput) (*ses.SendRawEmailOutput, error)
    SendRawEmailAsync(ctx workflow.Context, input *ses.SendRawEmailInput) *SesSendRawEmailResult

    SendTemplatedEmail(ctx workflow.Context, input *ses.SendTemplatedEmailInput) (*ses.SendTemplatedEmailOutput, error)
    SendTemplatedEmailAsync(ctx workflow.Context, input *ses.SendTemplatedEmailInput) *SesSendTemplatedEmailResult

    SetActiveReceiptRuleSet(ctx workflow.Context, input *ses.SetActiveReceiptRuleSetInput) (*ses.SetActiveReceiptRuleSetOutput, error)
    SetActiveReceiptRuleSetAsync(ctx workflow.Context, input *ses.SetActiveReceiptRuleSetInput) *SesSetActiveReceiptRuleSetResult

    SetIdentityDkimEnabled(ctx workflow.Context, input *ses.SetIdentityDkimEnabledInput) (*ses.SetIdentityDkimEnabledOutput, error)
    SetIdentityDkimEnabledAsync(ctx workflow.Context, input *ses.SetIdentityDkimEnabledInput) *SesSetIdentityDkimEnabledResult

    SetIdentityFeedbackForwardingEnabled(ctx workflow.Context, input *ses.SetIdentityFeedbackForwardingEnabledInput) (*ses.SetIdentityFeedbackForwardingEnabledOutput, error)
    SetIdentityFeedbackForwardingEnabledAsync(ctx workflow.Context, input *ses.SetIdentityFeedbackForwardingEnabledInput) *SesSetIdentityFeedbackForwardingEnabledResult

    SetIdentityHeadersInNotificationsEnabled(ctx workflow.Context, input *ses.SetIdentityHeadersInNotificationsEnabledInput) (*ses.SetIdentityHeadersInNotificationsEnabledOutput, error)
    SetIdentityHeadersInNotificationsEnabledAsync(ctx workflow.Context, input *ses.SetIdentityHeadersInNotificationsEnabledInput) *SesSetIdentityHeadersInNotificationsEnabledResult

    SetIdentityMailFromDomain(ctx workflow.Context, input *ses.SetIdentityMailFromDomainInput) (*ses.SetIdentityMailFromDomainOutput, error)
    SetIdentityMailFromDomainAsync(ctx workflow.Context, input *ses.SetIdentityMailFromDomainInput) *SesSetIdentityMailFromDomainResult

    SetIdentityNotificationTopic(ctx workflow.Context, input *ses.SetIdentityNotificationTopicInput) (*ses.SetIdentityNotificationTopicOutput, error)
    SetIdentityNotificationTopicAsync(ctx workflow.Context, input *ses.SetIdentityNotificationTopicInput) *SesSetIdentityNotificationTopicResult

    SetReceiptRulePosition(ctx workflow.Context, input *ses.SetReceiptRulePositionInput) (*ses.SetReceiptRulePositionOutput, error)
    SetReceiptRulePositionAsync(ctx workflow.Context, input *ses.SetReceiptRulePositionInput) *SesSetReceiptRulePositionResult

    TestRenderTemplate(ctx workflow.Context, input *ses.TestRenderTemplateInput) (*ses.TestRenderTemplateOutput, error)
    TestRenderTemplateAsync(ctx workflow.Context, input *ses.TestRenderTemplateInput) *SesTestRenderTemplateResult

    UpdateAccountSendingEnabled(ctx workflow.Context, input *ses.UpdateAccountSendingEnabledInput) (*ses.UpdateAccountSendingEnabledOutput, error)
    UpdateAccountSendingEnabledAsync(ctx workflow.Context, input *ses.UpdateAccountSendingEnabledInput) *SesUpdateAccountSendingEnabledResult

    UpdateConfigurationSetEventDestination(ctx workflow.Context, input *ses.UpdateConfigurationSetEventDestinationInput) (*ses.UpdateConfigurationSetEventDestinationOutput, error)
    UpdateConfigurationSetEventDestinationAsync(ctx workflow.Context, input *ses.UpdateConfigurationSetEventDestinationInput) *SesUpdateConfigurationSetEventDestinationResult

    UpdateConfigurationSetReputationMetricsEnabled(ctx workflow.Context, input *ses.UpdateConfigurationSetReputationMetricsEnabledInput) (*ses.UpdateConfigurationSetReputationMetricsEnabledOutput, error)
    UpdateConfigurationSetReputationMetricsEnabledAsync(ctx workflow.Context, input *ses.UpdateConfigurationSetReputationMetricsEnabledInput) *SesUpdateConfigurationSetReputationMetricsEnabledResult

    UpdateConfigurationSetSendingEnabled(ctx workflow.Context, input *ses.UpdateConfigurationSetSendingEnabledInput) (*ses.UpdateConfigurationSetSendingEnabledOutput, error)
    UpdateConfigurationSetSendingEnabledAsync(ctx workflow.Context, input *ses.UpdateConfigurationSetSendingEnabledInput) *SesUpdateConfigurationSetSendingEnabledResult

    UpdateConfigurationSetTrackingOptions(ctx workflow.Context, input *ses.UpdateConfigurationSetTrackingOptionsInput) (*ses.UpdateConfigurationSetTrackingOptionsOutput, error)
    UpdateConfigurationSetTrackingOptionsAsync(ctx workflow.Context, input *ses.UpdateConfigurationSetTrackingOptionsInput) *SesUpdateConfigurationSetTrackingOptionsResult

    UpdateCustomVerificationEmailTemplate(ctx workflow.Context, input *ses.UpdateCustomVerificationEmailTemplateInput) (*ses.UpdateCustomVerificationEmailTemplateOutput, error)
    UpdateCustomVerificationEmailTemplateAsync(ctx workflow.Context, input *ses.UpdateCustomVerificationEmailTemplateInput) *SesUpdateCustomVerificationEmailTemplateResult

    UpdateReceiptRule(ctx workflow.Context, input *ses.UpdateReceiptRuleInput) (*ses.UpdateReceiptRuleOutput, error)
    UpdateReceiptRuleAsync(ctx workflow.Context, input *ses.UpdateReceiptRuleInput) *SesUpdateReceiptRuleResult

    UpdateTemplate(ctx workflow.Context, input *ses.UpdateTemplateInput) (*ses.UpdateTemplateOutput, error)
    UpdateTemplateAsync(ctx workflow.Context, input *ses.UpdateTemplateInput) *SesUpdateTemplateResult

    VerifyDomainDkim(ctx workflow.Context, input *ses.VerifyDomainDkimInput) (*ses.VerifyDomainDkimOutput, error)
    VerifyDomainDkimAsync(ctx workflow.Context, input *ses.VerifyDomainDkimInput) *SesVerifyDomainDkimResult

    VerifyDomainIdentity(ctx workflow.Context, input *ses.VerifyDomainIdentityInput) (*ses.VerifyDomainIdentityOutput, error)
    VerifyDomainIdentityAsync(ctx workflow.Context, input *ses.VerifyDomainIdentityInput) *SesVerifyDomainIdentityResult

    VerifyEmailAddress(ctx workflow.Context, input *ses.VerifyEmailAddressInput) (*ses.VerifyEmailAddressOutput, error)
    VerifyEmailAddressAsync(ctx workflow.Context, input *ses.VerifyEmailAddressInput) *SesVerifyEmailAddressResult

    VerifyEmailIdentity(ctx workflow.Context, input *ses.VerifyEmailIdentityInput) (*ses.VerifyEmailIdentityOutput, error)
    VerifyEmailIdentityAsync(ctx workflow.Context, input *ses.VerifyEmailIdentityInput) *SesVerifyEmailIdentityResult

    WaitUntilIdentityExists(ctx workflow.Context, input *ses.GetIdentityVerificationAttributesInput) error}
type SesCloneReceiptRuleSetResult struct {
	Result workflow.Future
}

func (r *SesCloneReceiptRuleSetResult) Get(ctx workflow.Context) (*ses.CloneReceiptRuleSetOutput, error) {
    var output ses.CloneReceiptRuleSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SesCreateConfigurationSetResult struct {
	Result workflow.Future
}

func (r *SesCreateConfigurationSetResult) Get(ctx workflow.Context) (*ses.CreateConfigurationSetOutput, error) {
    var output ses.CreateConfigurationSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SesCreateConfigurationSetEventDestinationResult struct {
	Result workflow.Future
}

func (r *SesCreateConfigurationSetEventDestinationResult) Get(ctx workflow.Context) (*ses.CreateConfigurationSetEventDestinationOutput, error) {
    var output ses.CreateConfigurationSetEventDestinationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SesCreateConfigurationSetTrackingOptionsResult struct {
	Result workflow.Future
}

func (r *SesCreateConfigurationSetTrackingOptionsResult) Get(ctx workflow.Context) (*ses.CreateConfigurationSetTrackingOptionsOutput, error) {
    var output ses.CreateConfigurationSetTrackingOptionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SesCreateCustomVerificationEmailTemplateResult struct {
	Result workflow.Future
}

func (r *SesCreateCustomVerificationEmailTemplateResult) Get(ctx workflow.Context) (*ses.CreateCustomVerificationEmailTemplateOutput, error) {
    var output ses.CreateCustomVerificationEmailTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SesCreateReceiptFilterResult struct {
	Result workflow.Future
}

func (r *SesCreateReceiptFilterResult) Get(ctx workflow.Context) (*ses.CreateReceiptFilterOutput, error) {
    var output ses.CreateReceiptFilterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SesCreateReceiptRuleResult struct {
	Result workflow.Future
}

func (r *SesCreateReceiptRuleResult) Get(ctx workflow.Context) (*ses.CreateReceiptRuleOutput, error) {
    var output ses.CreateReceiptRuleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SesCreateReceiptRuleSetResult struct {
	Result workflow.Future
}

func (r *SesCreateReceiptRuleSetResult) Get(ctx workflow.Context) (*ses.CreateReceiptRuleSetOutput, error) {
    var output ses.CreateReceiptRuleSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SesCreateTemplateResult struct {
	Result workflow.Future
}

func (r *SesCreateTemplateResult) Get(ctx workflow.Context) (*ses.CreateTemplateOutput, error) {
    var output ses.CreateTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SesDeleteConfigurationSetResult struct {
	Result workflow.Future
}

func (r *SesDeleteConfigurationSetResult) Get(ctx workflow.Context) (*ses.DeleteConfigurationSetOutput, error) {
    var output ses.DeleteConfigurationSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SesDeleteConfigurationSetEventDestinationResult struct {
	Result workflow.Future
}

func (r *SesDeleteConfigurationSetEventDestinationResult) Get(ctx workflow.Context) (*ses.DeleteConfigurationSetEventDestinationOutput, error) {
    var output ses.DeleteConfigurationSetEventDestinationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SesDeleteConfigurationSetTrackingOptionsResult struct {
	Result workflow.Future
}

func (r *SesDeleteConfigurationSetTrackingOptionsResult) Get(ctx workflow.Context) (*ses.DeleteConfigurationSetTrackingOptionsOutput, error) {
    var output ses.DeleteConfigurationSetTrackingOptionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SesDeleteCustomVerificationEmailTemplateResult struct {
	Result workflow.Future
}

func (r *SesDeleteCustomVerificationEmailTemplateResult) Get(ctx workflow.Context) (*ses.DeleteCustomVerificationEmailTemplateOutput, error) {
    var output ses.DeleteCustomVerificationEmailTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SesDeleteIdentityResult struct {
	Result workflow.Future
}

func (r *SesDeleteIdentityResult) Get(ctx workflow.Context) (*ses.DeleteIdentityOutput, error) {
    var output ses.DeleteIdentityOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SesDeleteIdentityPolicyResult struct {
	Result workflow.Future
}

func (r *SesDeleteIdentityPolicyResult) Get(ctx workflow.Context) (*ses.DeleteIdentityPolicyOutput, error) {
    var output ses.DeleteIdentityPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SesDeleteReceiptFilterResult struct {
	Result workflow.Future
}

func (r *SesDeleteReceiptFilterResult) Get(ctx workflow.Context) (*ses.DeleteReceiptFilterOutput, error) {
    var output ses.DeleteReceiptFilterOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SesDeleteReceiptRuleResult struct {
	Result workflow.Future
}

func (r *SesDeleteReceiptRuleResult) Get(ctx workflow.Context) (*ses.DeleteReceiptRuleOutput, error) {
    var output ses.DeleteReceiptRuleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SesDeleteReceiptRuleSetResult struct {
	Result workflow.Future
}

func (r *SesDeleteReceiptRuleSetResult) Get(ctx workflow.Context) (*ses.DeleteReceiptRuleSetOutput, error) {
    var output ses.DeleteReceiptRuleSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SesDeleteTemplateResult struct {
	Result workflow.Future
}

func (r *SesDeleteTemplateResult) Get(ctx workflow.Context) (*ses.DeleteTemplateOutput, error) {
    var output ses.DeleteTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SesDeleteVerifiedEmailAddressResult struct {
	Result workflow.Future
}

func (r *SesDeleteVerifiedEmailAddressResult) Get(ctx workflow.Context) (*ses.DeleteVerifiedEmailAddressOutput, error) {
    var output ses.DeleteVerifiedEmailAddressOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SesDescribeActiveReceiptRuleSetResult struct {
	Result workflow.Future
}

func (r *SesDescribeActiveReceiptRuleSetResult) Get(ctx workflow.Context) (*ses.DescribeActiveReceiptRuleSetOutput, error) {
    var output ses.DescribeActiveReceiptRuleSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SesDescribeConfigurationSetResult struct {
	Result workflow.Future
}

func (r *SesDescribeConfigurationSetResult) Get(ctx workflow.Context) (*ses.DescribeConfigurationSetOutput, error) {
    var output ses.DescribeConfigurationSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SesDescribeReceiptRuleResult struct {
	Result workflow.Future
}

func (r *SesDescribeReceiptRuleResult) Get(ctx workflow.Context) (*ses.DescribeReceiptRuleOutput, error) {
    var output ses.DescribeReceiptRuleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SesDescribeReceiptRuleSetResult struct {
	Result workflow.Future
}

func (r *SesDescribeReceiptRuleSetResult) Get(ctx workflow.Context) (*ses.DescribeReceiptRuleSetOutput, error) {
    var output ses.DescribeReceiptRuleSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SesGetAccountSendingEnabledResult struct {
	Result workflow.Future
}

func (r *SesGetAccountSendingEnabledResult) Get(ctx workflow.Context) (*ses.GetAccountSendingEnabledOutput, error) {
    var output ses.GetAccountSendingEnabledOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SesGetCustomVerificationEmailTemplateResult struct {
	Result workflow.Future
}

func (r *SesGetCustomVerificationEmailTemplateResult) Get(ctx workflow.Context) (*ses.GetCustomVerificationEmailTemplateOutput, error) {
    var output ses.GetCustomVerificationEmailTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SesGetIdentityDkimAttributesResult struct {
	Result workflow.Future
}

func (r *SesGetIdentityDkimAttributesResult) Get(ctx workflow.Context) (*ses.GetIdentityDkimAttributesOutput, error) {
    var output ses.GetIdentityDkimAttributesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SesGetIdentityMailFromDomainAttributesResult struct {
	Result workflow.Future
}

func (r *SesGetIdentityMailFromDomainAttributesResult) Get(ctx workflow.Context) (*ses.GetIdentityMailFromDomainAttributesOutput, error) {
    var output ses.GetIdentityMailFromDomainAttributesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SesGetIdentityNotificationAttributesResult struct {
	Result workflow.Future
}

func (r *SesGetIdentityNotificationAttributesResult) Get(ctx workflow.Context) (*ses.GetIdentityNotificationAttributesOutput, error) {
    var output ses.GetIdentityNotificationAttributesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SesGetIdentityPoliciesResult struct {
	Result workflow.Future
}

func (r *SesGetIdentityPoliciesResult) Get(ctx workflow.Context) (*ses.GetIdentityPoliciesOutput, error) {
    var output ses.GetIdentityPoliciesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SesGetIdentityVerificationAttributesResult struct {
	Result workflow.Future
}

func (r *SesGetIdentityVerificationAttributesResult) Get(ctx workflow.Context) (*ses.GetIdentityVerificationAttributesOutput, error) {
    var output ses.GetIdentityVerificationAttributesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SesGetSendQuotaResult struct {
	Result workflow.Future
}

func (r *SesGetSendQuotaResult) Get(ctx workflow.Context) (*ses.GetSendQuotaOutput, error) {
    var output ses.GetSendQuotaOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SesGetSendStatisticsResult struct {
	Result workflow.Future
}

func (r *SesGetSendStatisticsResult) Get(ctx workflow.Context) (*ses.GetSendStatisticsOutput, error) {
    var output ses.GetSendStatisticsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SesGetTemplateResult struct {
	Result workflow.Future
}

func (r *SesGetTemplateResult) Get(ctx workflow.Context) (*ses.GetTemplateOutput, error) {
    var output ses.GetTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SesListConfigurationSetsResult struct {
	Result workflow.Future
}

func (r *SesListConfigurationSetsResult) Get(ctx workflow.Context) (*ses.ListConfigurationSetsOutput, error) {
    var output ses.ListConfigurationSetsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SesListCustomVerificationEmailTemplatesResult struct {
	Result workflow.Future
}

func (r *SesListCustomVerificationEmailTemplatesResult) Get(ctx workflow.Context) (*ses.ListCustomVerificationEmailTemplatesOutput, error) {
    var output ses.ListCustomVerificationEmailTemplatesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SesListIdentitiesResult struct {
	Result workflow.Future
}

func (r *SesListIdentitiesResult) Get(ctx workflow.Context) (*ses.ListIdentitiesOutput, error) {
    var output ses.ListIdentitiesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SesListIdentityPoliciesResult struct {
	Result workflow.Future
}

func (r *SesListIdentityPoliciesResult) Get(ctx workflow.Context) (*ses.ListIdentityPoliciesOutput, error) {
    var output ses.ListIdentityPoliciesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SesListReceiptFiltersResult struct {
	Result workflow.Future
}

func (r *SesListReceiptFiltersResult) Get(ctx workflow.Context) (*ses.ListReceiptFiltersOutput, error) {
    var output ses.ListReceiptFiltersOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SesListReceiptRuleSetsResult struct {
	Result workflow.Future
}

func (r *SesListReceiptRuleSetsResult) Get(ctx workflow.Context) (*ses.ListReceiptRuleSetsOutput, error) {
    var output ses.ListReceiptRuleSetsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SesListTemplatesResult struct {
	Result workflow.Future
}

func (r *SesListTemplatesResult) Get(ctx workflow.Context) (*ses.ListTemplatesOutput, error) {
    var output ses.ListTemplatesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SesListVerifiedEmailAddressesResult struct {
	Result workflow.Future
}

func (r *SesListVerifiedEmailAddressesResult) Get(ctx workflow.Context) (*ses.ListVerifiedEmailAddressesOutput, error) {
    var output ses.ListVerifiedEmailAddressesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SesPutConfigurationSetDeliveryOptionsResult struct {
	Result workflow.Future
}

func (r *SesPutConfigurationSetDeliveryOptionsResult) Get(ctx workflow.Context) (*ses.PutConfigurationSetDeliveryOptionsOutput, error) {
    var output ses.PutConfigurationSetDeliveryOptionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SesPutIdentityPolicyResult struct {
	Result workflow.Future
}

func (r *SesPutIdentityPolicyResult) Get(ctx workflow.Context) (*ses.PutIdentityPolicyOutput, error) {
    var output ses.PutIdentityPolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SesReorderReceiptRuleSetResult struct {
	Result workflow.Future
}

func (r *SesReorderReceiptRuleSetResult) Get(ctx workflow.Context) (*ses.ReorderReceiptRuleSetOutput, error) {
    var output ses.ReorderReceiptRuleSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SesSendBounceResult struct {
	Result workflow.Future
}

func (r *SesSendBounceResult) Get(ctx workflow.Context) (*ses.SendBounceOutput, error) {
    var output ses.SendBounceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SesSendBulkTemplatedEmailResult struct {
	Result workflow.Future
}

func (r *SesSendBulkTemplatedEmailResult) Get(ctx workflow.Context) (*ses.SendBulkTemplatedEmailOutput, error) {
    var output ses.SendBulkTemplatedEmailOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SesSendCustomVerificationEmailResult struct {
	Result workflow.Future
}

func (r *SesSendCustomVerificationEmailResult) Get(ctx workflow.Context) (*ses.SendCustomVerificationEmailOutput, error) {
    var output ses.SendCustomVerificationEmailOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SesSendEmailResult struct {
	Result workflow.Future
}

func (r *SesSendEmailResult) Get(ctx workflow.Context) (*ses.SendEmailOutput, error) {
    var output ses.SendEmailOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SesSendRawEmailResult struct {
	Result workflow.Future
}

func (r *SesSendRawEmailResult) Get(ctx workflow.Context) (*ses.SendRawEmailOutput, error) {
    var output ses.SendRawEmailOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SesSendTemplatedEmailResult struct {
	Result workflow.Future
}

func (r *SesSendTemplatedEmailResult) Get(ctx workflow.Context) (*ses.SendTemplatedEmailOutput, error) {
    var output ses.SendTemplatedEmailOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SesSetActiveReceiptRuleSetResult struct {
	Result workflow.Future
}

func (r *SesSetActiveReceiptRuleSetResult) Get(ctx workflow.Context) (*ses.SetActiveReceiptRuleSetOutput, error) {
    var output ses.SetActiveReceiptRuleSetOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SesSetIdentityDkimEnabledResult struct {
	Result workflow.Future
}

func (r *SesSetIdentityDkimEnabledResult) Get(ctx workflow.Context) (*ses.SetIdentityDkimEnabledOutput, error) {
    var output ses.SetIdentityDkimEnabledOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SesSetIdentityFeedbackForwardingEnabledResult struct {
	Result workflow.Future
}

func (r *SesSetIdentityFeedbackForwardingEnabledResult) Get(ctx workflow.Context) (*ses.SetIdentityFeedbackForwardingEnabledOutput, error) {
    var output ses.SetIdentityFeedbackForwardingEnabledOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SesSetIdentityHeadersInNotificationsEnabledResult struct {
	Result workflow.Future
}

func (r *SesSetIdentityHeadersInNotificationsEnabledResult) Get(ctx workflow.Context) (*ses.SetIdentityHeadersInNotificationsEnabledOutput, error) {
    var output ses.SetIdentityHeadersInNotificationsEnabledOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SesSetIdentityMailFromDomainResult struct {
	Result workflow.Future
}

func (r *SesSetIdentityMailFromDomainResult) Get(ctx workflow.Context) (*ses.SetIdentityMailFromDomainOutput, error) {
    var output ses.SetIdentityMailFromDomainOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SesSetIdentityNotificationTopicResult struct {
	Result workflow.Future
}

func (r *SesSetIdentityNotificationTopicResult) Get(ctx workflow.Context) (*ses.SetIdentityNotificationTopicOutput, error) {
    var output ses.SetIdentityNotificationTopicOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SesSetReceiptRulePositionResult struct {
	Result workflow.Future
}

func (r *SesSetReceiptRulePositionResult) Get(ctx workflow.Context) (*ses.SetReceiptRulePositionOutput, error) {
    var output ses.SetReceiptRulePositionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SesTestRenderTemplateResult struct {
	Result workflow.Future
}

func (r *SesTestRenderTemplateResult) Get(ctx workflow.Context) (*ses.TestRenderTemplateOutput, error) {
    var output ses.TestRenderTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SesUpdateAccountSendingEnabledResult struct {
	Result workflow.Future
}

func (r *SesUpdateAccountSendingEnabledResult) Get(ctx workflow.Context) (*ses.UpdateAccountSendingEnabledOutput, error) {
    var output ses.UpdateAccountSendingEnabledOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SesUpdateConfigurationSetEventDestinationResult struct {
	Result workflow.Future
}

func (r *SesUpdateConfigurationSetEventDestinationResult) Get(ctx workflow.Context) (*ses.UpdateConfigurationSetEventDestinationOutput, error) {
    var output ses.UpdateConfigurationSetEventDestinationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SesUpdateConfigurationSetReputationMetricsEnabledResult struct {
	Result workflow.Future
}

func (r *SesUpdateConfigurationSetReputationMetricsEnabledResult) Get(ctx workflow.Context) (*ses.UpdateConfigurationSetReputationMetricsEnabledOutput, error) {
    var output ses.UpdateConfigurationSetReputationMetricsEnabledOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SesUpdateConfigurationSetSendingEnabledResult struct {
	Result workflow.Future
}

func (r *SesUpdateConfigurationSetSendingEnabledResult) Get(ctx workflow.Context) (*ses.UpdateConfigurationSetSendingEnabledOutput, error) {
    var output ses.UpdateConfigurationSetSendingEnabledOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SesUpdateConfigurationSetTrackingOptionsResult struct {
	Result workflow.Future
}

func (r *SesUpdateConfigurationSetTrackingOptionsResult) Get(ctx workflow.Context) (*ses.UpdateConfigurationSetTrackingOptionsOutput, error) {
    var output ses.UpdateConfigurationSetTrackingOptionsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SesUpdateCustomVerificationEmailTemplateResult struct {
	Result workflow.Future
}

func (r *SesUpdateCustomVerificationEmailTemplateResult) Get(ctx workflow.Context) (*ses.UpdateCustomVerificationEmailTemplateOutput, error) {
    var output ses.UpdateCustomVerificationEmailTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SesUpdateReceiptRuleResult struct {
	Result workflow.Future
}

func (r *SesUpdateReceiptRuleResult) Get(ctx workflow.Context) (*ses.UpdateReceiptRuleOutput, error) {
    var output ses.UpdateReceiptRuleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SesUpdateTemplateResult struct {
	Result workflow.Future
}

func (r *SesUpdateTemplateResult) Get(ctx workflow.Context) (*ses.UpdateTemplateOutput, error) {
    var output ses.UpdateTemplateOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SesVerifyDomainDkimResult struct {
	Result workflow.Future
}

func (r *SesVerifyDomainDkimResult) Get(ctx workflow.Context) (*ses.VerifyDomainDkimOutput, error) {
    var output ses.VerifyDomainDkimOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SesVerifyDomainIdentityResult struct {
	Result workflow.Future
}

func (r *SesVerifyDomainIdentityResult) Get(ctx workflow.Context) (*ses.VerifyDomainIdentityOutput, error) {
    var output ses.VerifyDomainIdentityOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SesVerifyEmailAddressResult struct {
	Result workflow.Future
}

func (r *SesVerifyEmailAddressResult) Get(ctx workflow.Context) (*ses.VerifyEmailAddressOutput, error) {
    var output ses.VerifyEmailAddressOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SesVerifyEmailIdentityResult struct {
	Result workflow.Future
}

func (r *SesVerifyEmailIdentityResult) Get(ctx workflow.Context) (*ses.VerifyEmailIdentityOutput, error) {
    var output ses.VerifyEmailIdentityOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type SESStub struct {
    activities SESClient
}

func NewSESStub() SESClient {
    return &SESStub{}
}

func (a *SESStub) CloneReceiptRuleSet(ctx workflow.Context, input *ses.CloneReceiptRuleSetInput) (*ses.CloneReceiptRuleSetOutput, error) {
    var output ses.CloneReceiptRuleSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CloneReceiptRuleSet, input).Get(ctx, &output)
    return &output, err
}

func (a *SESStub) CreateConfigurationSet(ctx workflow.Context, input *ses.CreateConfigurationSetInput) (*ses.CreateConfigurationSetOutput, error) {
    var output ses.CreateConfigurationSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateConfigurationSet, input).Get(ctx, &output)
    return &output, err
}

func (a *SESStub) CreateConfigurationSetEventDestination(ctx workflow.Context, input *ses.CreateConfigurationSetEventDestinationInput) (*ses.CreateConfigurationSetEventDestinationOutput, error) {
    var output ses.CreateConfigurationSetEventDestinationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateConfigurationSetEventDestination, input).Get(ctx, &output)
    return &output, err
}

func (a *SESStub) CreateConfigurationSetTrackingOptions(ctx workflow.Context, input *ses.CreateConfigurationSetTrackingOptionsInput) (*ses.CreateConfigurationSetTrackingOptionsOutput, error) {
    var output ses.CreateConfigurationSetTrackingOptionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateConfigurationSetTrackingOptions, input).Get(ctx, &output)
    return &output, err
}

func (a *SESStub) CreateCustomVerificationEmailTemplate(ctx workflow.Context, input *ses.CreateCustomVerificationEmailTemplateInput) (*ses.CreateCustomVerificationEmailTemplateOutput, error) {
    var output ses.CreateCustomVerificationEmailTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateCustomVerificationEmailTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *SESStub) CreateReceiptFilter(ctx workflow.Context, input *ses.CreateReceiptFilterInput) (*ses.CreateReceiptFilterOutput, error) {
    var output ses.CreateReceiptFilterOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateReceiptFilter, input).Get(ctx, &output)
    return &output, err
}

func (a *SESStub) CreateReceiptRule(ctx workflow.Context, input *ses.CreateReceiptRuleInput) (*ses.CreateReceiptRuleOutput, error) {
    var output ses.CreateReceiptRuleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateReceiptRule, input).Get(ctx, &output)
    return &output, err
}

func (a *SESStub) CreateReceiptRuleSet(ctx workflow.Context, input *ses.CreateReceiptRuleSetInput) (*ses.CreateReceiptRuleSetOutput, error) {
    var output ses.CreateReceiptRuleSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateReceiptRuleSet, input).Get(ctx, &output)
    return &output, err
}

func (a *SESStub) CreateTemplate(ctx workflow.Context, input *ses.CreateTemplateInput) (*ses.CreateTemplateOutput, error) {
    var output ses.CreateTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *SESStub) DeleteConfigurationSet(ctx workflow.Context, input *ses.DeleteConfigurationSetInput) (*ses.DeleteConfigurationSetOutput, error) {
    var output ses.DeleteConfigurationSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteConfigurationSet, input).Get(ctx, &output)
    return &output, err
}

func (a *SESStub) DeleteConfigurationSetEventDestination(ctx workflow.Context, input *ses.DeleteConfigurationSetEventDestinationInput) (*ses.DeleteConfigurationSetEventDestinationOutput, error) {
    var output ses.DeleteConfigurationSetEventDestinationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteConfigurationSetEventDestination, input).Get(ctx, &output)
    return &output, err
}

func (a *SESStub) DeleteConfigurationSetTrackingOptions(ctx workflow.Context, input *ses.DeleteConfigurationSetTrackingOptionsInput) (*ses.DeleteConfigurationSetTrackingOptionsOutput, error) {
    var output ses.DeleteConfigurationSetTrackingOptionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteConfigurationSetTrackingOptions, input).Get(ctx, &output)
    return &output, err
}

func (a *SESStub) DeleteCustomVerificationEmailTemplate(ctx workflow.Context, input *ses.DeleteCustomVerificationEmailTemplateInput) (*ses.DeleteCustomVerificationEmailTemplateOutput, error) {
    var output ses.DeleteCustomVerificationEmailTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteCustomVerificationEmailTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *SESStub) DeleteIdentity(ctx workflow.Context, input *ses.DeleteIdentityInput) (*ses.DeleteIdentityOutput, error) {
    var output ses.DeleteIdentityOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteIdentity, input).Get(ctx, &output)
    return &output, err
}

func (a *SESStub) DeleteIdentityPolicy(ctx workflow.Context, input *ses.DeleteIdentityPolicyInput) (*ses.DeleteIdentityPolicyOutput, error) {
    var output ses.DeleteIdentityPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteIdentityPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *SESStub) DeleteReceiptFilter(ctx workflow.Context, input *ses.DeleteReceiptFilterInput) (*ses.DeleteReceiptFilterOutput, error) {
    var output ses.DeleteReceiptFilterOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteReceiptFilter, input).Get(ctx, &output)
    return &output, err
}

func (a *SESStub) DeleteReceiptRule(ctx workflow.Context, input *ses.DeleteReceiptRuleInput) (*ses.DeleteReceiptRuleOutput, error) {
    var output ses.DeleteReceiptRuleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteReceiptRule, input).Get(ctx, &output)
    return &output, err
}

func (a *SESStub) DeleteReceiptRuleSet(ctx workflow.Context, input *ses.DeleteReceiptRuleSetInput) (*ses.DeleteReceiptRuleSetOutput, error) {
    var output ses.DeleteReceiptRuleSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteReceiptRuleSet, input).Get(ctx, &output)
    return &output, err
}

func (a *SESStub) DeleteTemplate(ctx workflow.Context, input *ses.DeleteTemplateInput) (*ses.DeleteTemplateOutput, error) {
    var output ses.DeleteTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *SESStub) DeleteVerifiedEmailAddress(ctx workflow.Context, input *ses.DeleteVerifiedEmailAddressInput) (*ses.DeleteVerifiedEmailAddressOutput, error) {
    var output ses.DeleteVerifiedEmailAddressOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteVerifiedEmailAddress, input).Get(ctx, &output)
    return &output, err
}

func (a *SESStub) DescribeActiveReceiptRuleSet(ctx workflow.Context, input *ses.DescribeActiveReceiptRuleSetInput) (*ses.DescribeActiveReceiptRuleSetOutput, error) {
    var output ses.DescribeActiveReceiptRuleSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeActiveReceiptRuleSet, input).Get(ctx, &output)
    return &output, err
}

func (a *SESStub) DescribeConfigurationSet(ctx workflow.Context, input *ses.DescribeConfigurationSetInput) (*ses.DescribeConfigurationSetOutput, error) {
    var output ses.DescribeConfigurationSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeConfigurationSet, input).Get(ctx, &output)
    return &output, err
}

func (a *SESStub) DescribeReceiptRule(ctx workflow.Context, input *ses.DescribeReceiptRuleInput) (*ses.DescribeReceiptRuleOutput, error) {
    var output ses.DescribeReceiptRuleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeReceiptRule, input).Get(ctx, &output)
    return &output, err
}

func (a *SESStub) DescribeReceiptRuleSet(ctx workflow.Context, input *ses.DescribeReceiptRuleSetInput) (*ses.DescribeReceiptRuleSetOutput, error) {
    var output ses.DescribeReceiptRuleSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeReceiptRuleSet, input).Get(ctx, &output)
    return &output, err
}

func (a *SESStub) GetAccountSendingEnabled(ctx workflow.Context, input *ses.GetAccountSendingEnabledInput) (*ses.GetAccountSendingEnabledOutput, error) {
    var output ses.GetAccountSendingEnabledOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetAccountSendingEnabled, input).Get(ctx, &output)
    return &output, err
}

func (a *SESStub) GetCustomVerificationEmailTemplate(ctx workflow.Context, input *ses.GetCustomVerificationEmailTemplateInput) (*ses.GetCustomVerificationEmailTemplateOutput, error) {
    var output ses.GetCustomVerificationEmailTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetCustomVerificationEmailTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *SESStub) GetIdentityDkimAttributes(ctx workflow.Context, input *ses.GetIdentityDkimAttributesInput) (*ses.GetIdentityDkimAttributesOutput, error) {
    var output ses.GetIdentityDkimAttributesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetIdentityDkimAttributes, input).Get(ctx, &output)
    return &output, err
}

func (a *SESStub) GetIdentityMailFromDomainAttributes(ctx workflow.Context, input *ses.GetIdentityMailFromDomainAttributesInput) (*ses.GetIdentityMailFromDomainAttributesOutput, error) {
    var output ses.GetIdentityMailFromDomainAttributesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetIdentityMailFromDomainAttributes, input).Get(ctx, &output)
    return &output, err
}

func (a *SESStub) GetIdentityNotificationAttributes(ctx workflow.Context, input *ses.GetIdentityNotificationAttributesInput) (*ses.GetIdentityNotificationAttributesOutput, error) {
    var output ses.GetIdentityNotificationAttributesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetIdentityNotificationAttributes, input).Get(ctx, &output)
    return &output, err
}

func (a *SESStub) GetIdentityPolicies(ctx workflow.Context, input *ses.GetIdentityPoliciesInput) (*ses.GetIdentityPoliciesOutput, error) {
    var output ses.GetIdentityPoliciesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetIdentityPolicies, input).Get(ctx, &output)
    return &output, err
}

func (a *SESStub) GetIdentityVerificationAttributes(ctx workflow.Context, input *ses.GetIdentityVerificationAttributesInput) (*ses.GetIdentityVerificationAttributesOutput, error) {
    var output ses.GetIdentityVerificationAttributesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetIdentityVerificationAttributes, input).Get(ctx, &output)
    return &output, err
}

func (a *SESStub) GetSendQuota(ctx workflow.Context, input *ses.GetSendQuotaInput) (*ses.GetSendQuotaOutput, error) {
    var output ses.GetSendQuotaOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetSendQuota, input).Get(ctx, &output)
    return &output, err
}

func (a *SESStub) GetSendStatistics(ctx workflow.Context, input *ses.GetSendStatisticsInput) (*ses.GetSendStatisticsOutput, error) {
    var output ses.GetSendStatisticsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetSendStatistics, input).Get(ctx, &output)
    return &output, err
}

func (a *SESStub) GetTemplate(ctx workflow.Context, input *ses.GetTemplateInput) (*ses.GetTemplateOutput, error) {
    var output ses.GetTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *SESStub) ListConfigurationSets(ctx workflow.Context, input *ses.ListConfigurationSetsInput) (*ses.ListConfigurationSetsOutput, error) {
    var output ses.ListConfigurationSetsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListConfigurationSets, input).Get(ctx, &output)
    return &output, err
}

func (a *SESStub) ListCustomVerificationEmailTemplates(ctx workflow.Context, input *ses.ListCustomVerificationEmailTemplatesInput) (*ses.ListCustomVerificationEmailTemplatesOutput, error) {
    var output ses.ListCustomVerificationEmailTemplatesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListCustomVerificationEmailTemplates, input).Get(ctx, &output)
    return &output, err
}

func (a *SESStub) ListIdentities(ctx workflow.Context, input *ses.ListIdentitiesInput) (*ses.ListIdentitiesOutput, error) {
    var output ses.ListIdentitiesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListIdentities, input).Get(ctx, &output)
    return &output, err
}

func (a *SESStub) ListIdentityPolicies(ctx workflow.Context, input *ses.ListIdentityPoliciesInput) (*ses.ListIdentityPoliciesOutput, error) {
    var output ses.ListIdentityPoliciesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListIdentityPolicies, input).Get(ctx, &output)
    return &output, err
}

func (a *SESStub) ListReceiptFilters(ctx workflow.Context, input *ses.ListReceiptFiltersInput) (*ses.ListReceiptFiltersOutput, error) {
    var output ses.ListReceiptFiltersOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListReceiptFilters, input).Get(ctx, &output)
    return &output, err
}

func (a *SESStub) ListReceiptRuleSets(ctx workflow.Context, input *ses.ListReceiptRuleSetsInput) (*ses.ListReceiptRuleSetsOutput, error) {
    var output ses.ListReceiptRuleSetsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListReceiptRuleSets, input).Get(ctx, &output)
    return &output, err
}

func (a *SESStub) ListTemplates(ctx workflow.Context, input *ses.ListTemplatesInput) (*ses.ListTemplatesOutput, error) {
    var output ses.ListTemplatesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTemplates, input).Get(ctx, &output)
    return &output, err
}

func (a *SESStub) ListVerifiedEmailAddresses(ctx workflow.Context, input *ses.ListVerifiedEmailAddressesInput) (*ses.ListVerifiedEmailAddressesOutput, error) {
    var output ses.ListVerifiedEmailAddressesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListVerifiedEmailAddresses, input).Get(ctx, &output)
    return &output, err
}

func (a *SESStub) PutConfigurationSetDeliveryOptions(ctx workflow.Context, input *ses.PutConfigurationSetDeliveryOptionsInput) (*ses.PutConfigurationSetDeliveryOptionsOutput, error) {
    var output ses.PutConfigurationSetDeliveryOptionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutConfigurationSetDeliveryOptions, input).Get(ctx, &output)
    return &output, err
}

func (a *SESStub) PutIdentityPolicy(ctx workflow.Context, input *ses.PutIdentityPolicyInput) (*ses.PutIdentityPolicyOutput, error) {
    var output ses.PutIdentityPolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutIdentityPolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *SESStub) ReorderReceiptRuleSet(ctx workflow.Context, input *ses.ReorderReceiptRuleSetInput) (*ses.ReorderReceiptRuleSetOutput, error) {
    var output ses.ReorderReceiptRuleSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ReorderReceiptRuleSet, input).Get(ctx, &output)
    return &output, err
}

func (a *SESStub) SendBounce(ctx workflow.Context, input *ses.SendBounceInput) (*ses.SendBounceOutput, error) {
    var output ses.SendBounceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SendBounce, input).Get(ctx, &output)
    return &output, err
}

func (a *SESStub) SendBulkTemplatedEmail(ctx workflow.Context, input *ses.SendBulkTemplatedEmailInput) (*ses.SendBulkTemplatedEmailOutput, error) {
    var output ses.SendBulkTemplatedEmailOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SendBulkTemplatedEmail, input).Get(ctx, &output)
    return &output, err
}

func (a *SESStub) SendCustomVerificationEmail(ctx workflow.Context, input *ses.SendCustomVerificationEmailInput) (*ses.SendCustomVerificationEmailOutput, error) {
    var output ses.SendCustomVerificationEmailOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SendCustomVerificationEmail, input).Get(ctx, &output)
    return &output, err
}

func (a *SESStub) SendEmail(ctx workflow.Context, input *ses.SendEmailInput) (*ses.SendEmailOutput, error) {
    var output ses.SendEmailOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SendEmail, input).Get(ctx, &output)
    return &output, err
}

func (a *SESStub) SendRawEmail(ctx workflow.Context, input *ses.SendRawEmailInput) (*ses.SendRawEmailOutput, error) {
    var output ses.SendRawEmailOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SendRawEmail, input).Get(ctx, &output)
    return &output, err
}

func (a *SESStub) SendTemplatedEmail(ctx workflow.Context, input *ses.SendTemplatedEmailInput) (*ses.SendTemplatedEmailOutput, error) {
    var output ses.SendTemplatedEmailOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SendTemplatedEmail, input).Get(ctx, &output)
    return &output, err
}

func (a *SESStub) SetActiveReceiptRuleSet(ctx workflow.Context, input *ses.SetActiveReceiptRuleSetInput) (*ses.SetActiveReceiptRuleSetOutput, error) {
    var output ses.SetActiveReceiptRuleSetOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SetActiveReceiptRuleSet, input).Get(ctx, &output)
    return &output, err
}

func (a *SESStub) SetIdentityDkimEnabled(ctx workflow.Context, input *ses.SetIdentityDkimEnabledInput) (*ses.SetIdentityDkimEnabledOutput, error) {
    var output ses.SetIdentityDkimEnabledOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SetIdentityDkimEnabled, input).Get(ctx, &output)
    return &output, err
}

func (a *SESStub) SetIdentityFeedbackForwardingEnabled(ctx workflow.Context, input *ses.SetIdentityFeedbackForwardingEnabledInput) (*ses.SetIdentityFeedbackForwardingEnabledOutput, error) {
    var output ses.SetIdentityFeedbackForwardingEnabledOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SetIdentityFeedbackForwardingEnabled, input).Get(ctx, &output)
    return &output, err
}

func (a *SESStub) SetIdentityHeadersInNotificationsEnabled(ctx workflow.Context, input *ses.SetIdentityHeadersInNotificationsEnabledInput) (*ses.SetIdentityHeadersInNotificationsEnabledOutput, error) {
    var output ses.SetIdentityHeadersInNotificationsEnabledOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SetIdentityHeadersInNotificationsEnabled, input).Get(ctx, &output)
    return &output, err
}

func (a *SESStub) SetIdentityMailFromDomain(ctx workflow.Context, input *ses.SetIdentityMailFromDomainInput) (*ses.SetIdentityMailFromDomainOutput, error) {
    var output ses.SetIdentityMailFromDomainOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SetIdentityMailFromDomain, input).Get(ctx, &output)
    return &output, err
}

func (a *SESStub) SetIdentityNotificationTopic(ctx workflow.Context, input *ses.SetIdentityNotificationTopicInput) (*ses.SetIdentityNotificationTopicOutput, error) {
    var output ses.SetIdentityNotificationTopicOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SetIdentityNotificationTopic, input).Get(ctx, &output)
    return &output, err
}

func (a *SESStub) SetReceiptRulePosition(ctx workflow.Context, input *ses.SetReceiptRulePositionInput) (*ses.SetReceiptRulePositionOutput, error) {
    var output ses.SetReceiptRulePositionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SetReceiptRulePosition, input).Get(ctx, &output)
    return &output, err
}

func (a *SESStub) TestRenderTemplate(ctx workflow.Context, input *ses.TestRenderTemplateInput) (*ses.TestRenderTemplateOutput, error) {
    var output ses.TestRenderTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TestRenderTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *SESStub) UpdateAccountSendingEnabled(ctx workflow.Context, input *ses.UpdateAccountSendingEnabledInput) (*ses.UpdateAccountSendingEnabledOutput, error) {
    var output ses.UpdateAccountSendingEnabledOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateAccountSendingEnabled, input).Get(ctx, &output)
    return &output, err
}

func (a *SESStub) UpdateConfigurationSetEventDestination(ctx workflow.Context, input *ses.UpdateConfigurationSetEventDestinationInput) (*ses.UpdateConfigurationSetEventDestinationOutput, error) {
    var output ses.UpdateConfigurationSetEventDestinationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateConfigurationSetEventDestination, input).Get(ctx, &output)
    return &output, err
}

func (a *SESStub) UpdateConfigurationSetReputationMetricsEnabled(ctx workflow.Context, input *ses.UpdateConfigurationSetReputationMetricsEnabledInput) (*ses.UpdateConfigurationSetReputationMetricsEnabledOutput, error) {
    var output ses.UpdateConfigurationSetReputationMetricsEnabledOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateConfigurationSetReputationMetricsEnabled, input).Get(ctx, &output)
    return &output, err
}

func (a *SESStub) UpdateConfigurationSetSendingEnabled(ctx workflow.Context, input *ses.UpdateConfigurationSetSendingEnabledInput) (*ses.UpdateConfigurationSetSendingEnabledOutput, error) {
    var output ses.UpdateConfigurationSetSendingEnabledOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateConfigurationSetSendingEnabled, input).Get(ctx, &output)
    return &output, err
}

func (a *SESStub) UpdateConfigurationSetTrackingOptions(ctx workflow.Context, input *ses.UpdateConfigurationSetTrackingOptionsInput) (*ses.UpdateConfigurationSetTrackingOptionsOutput, error) {
    var output ses.UpdateConfigurationSetTrackingOptionsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateConfigurationSetTrackingOptions, input).Get(ctx, &output)
    return &output, err
}

func (a *SESStub) UpdateCustomVerificationEmailTemplate(ctx workflow.Context, input *ses.UpdateCustomVerificationEmailTemplateInput) (*ses.UpdateCustomVerificationEmailTemplateOutput, error) {
    var output ses.UpdateCustomVerificationEmailTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateCustomVerificationEmailTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *SESStub) UpdateReceiptRule(ctx workflow.Context, input *ses.UpdateReceiptRuleInput) (*ses.UpdateReceiptRuleOutput, error) {
    var output ses.UpdateReceiptRuleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateReceiptRule, input).Get(ctx, &output)
    return &output, err
}

func (a *SESStub) UpdateTemplate(ctx workflow.Context, input *ses.UpdateTemplateInput) (*ses.UpdateTemplateOutput, error) {
    var output ses.UpdateTemplateOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateTemplate, input).Get(ctx, &output)
    return &output, err
}

func (a *SESStub) VerifyDomainDkim(ctx workflow.Context, input *ses.VerifyDomainDkimInput) (*ses.VerifyDomainDkimOutput, error) {
    var output ses.VerifyDomainDkimOutput
    err := workflow.ExecuteActivity(ctx, a.activities.VerifyDomainDkim, input).Get(ctx, &output)
    return &output, err
}

func (a *SESStub) VerifyDomainIdentity(ctx workflow.Context, input *ses.VerifyDomainIdentityInput) (*ses.VerifyDomainIdentityOutput, error) {
    var output ses.VerifyDomainIdentityOutput
    err := workflow.ExecuteActivity(ctx, a.activities.VerifyDomainIdentity, input).Get(ctx, &output)
    return &output, err
}

func (a *SESStub) VerifyEmailAddress(ctx workflow.Context, input *ses.VerifyEmailAddressInput) (*ses.VerifyEmailAddressOutput, error) {
    var output ses.VerifyEmailAddressOutput
    err := workflow.ExecuteActivity(ctx, a.activities.VerifyEmailAddress, input).Get(ctx, &output)
    return &output, err
}

func (a *SESStub) VerifyEmailIdentity(ctx workflow.Context, input *ses.VerifyEmailIdentityInput) (*ses.VerifyEmailIdentityOutput, error) {
    var output ses.VerifyEmailIdentityOutput
    err := workflow.ExecuteActivity(ctx, a.activities.VerifyEmailIdentity, input).Get(ctx, &output)
    return &output, err
}


func (a *SESStub) WaitUntilIdentityExists(ctx workflow.Context, input *ses.GetIdentityVerificationAttributesInput) error {
    return a.client.WaitUntilIdentityExists(input)
}