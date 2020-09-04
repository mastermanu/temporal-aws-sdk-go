package awsclients

import (
	"github.com/aws/aws-sdk-go/service/savingsplans"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type SavingsPlansClient interface {
    CreateSavingsPlan(ctx workflow.Context, input *savingsplans.CreateSavingsPlanInput) (*savingsplans.CreateSavingsPlanOutput, error)
    CreateSavingsPlanAsync(ctx workflow.Context, input *savingsplans.CreateSavingsPlanInput) *SavingsplansCreateSavingsPlanResult

    DescribeSavingsPlanRates(ctx workflow.Context, input *savingsplans.DescribeSavingsPlanRatesInput) (*savingsplans.DescribeSavingsPlanRatesOutput, error)
    DescribeSavingsPlanRatesAsync(ctx workflow.Context, input *savingsplans.DescribeSavingsPlanRatesInput) *SavingsplansDescribeSavingsPlanRatesResult

    DescribeSavingsPlans(ctx workflow.Context, input *savingsplans.DescribeSavingsPlansInput) (*savingsplans.DescribeSavingsPlansOutput, error)
    DescribeSavingsPlansAsync(ctx workflow.Context, input *savingsplans.DescribeSavingsPlansInput) *SavingsplansDescribeSavingsPlansResult

    DescribeSavingsPlansOfferingRates(ctx workflow.Context, input *savingsplans.DescribeSavingsPlansOfferingRatesInput) (*savingsplans.DescribeSavingsPlansOfferingRatesOutput, error)
    DescribeSavingsPlansOfferingRatesAsync(ctx workflow.Context, input *savingsplans.DescribeSavingsPlansOfferingRatesInput) *SavingsplansDescribeSavingsPlansOfferingRatesResult

    DescribeSavingsPlansOfferings(ctx workflow.Context, input *savingsplans.DescribeSavingsPlansOfferingsInput) (*savingsplans.DescribeSavingsPlansOfferingsOutput, error)
    DescribeSavingsPlansOfferingsAsync(ctx workflow.Context, input *savingsplans.DescribeSavingsPlansOfferingsInput) *SavingsplansDescribeSavingsPlansOfferingsResult

    ListTagsForResource(ctx workflow.Context, input *savingsplans.ListTagsForResourceInput) (*savingsplans.ListTagsForResourceOutput, error)
    ListTagsForResourceAsync(ctx workflow.Context, input *savingsplans.ListTagsForResourceInput) *SavingsplansListTagsForResourceResult

    TagResource(ctx workflow.Context, input *savingsplans.TagResourceInput) (*savingsplans.TagResourceOutput, error)
    TagResourceAsync(ctx workflow.Context, input *savingsplans.TagResourceInput) *SavingsplansTagResourceResult

    UntagResource(ctx workflow.Context, input *savingsplans.UntagResourceInput) (*savingsplans.UntagResourceOutput, error)
    UntagResourceAsync(ctx workflow.Context, input *savingsplans.UntagResourceInput) *SavingsplansUntagResourceResult
}

type SavingsplansCreateSavingsPlanResult struct {
	Result workflow.Future
}

func (r *SavingsplansCreateSavingsPlanResult) Get(ctx workflow.Context) (*savingsplans.CreateSavingsPlanOutput, error) {
    var output savingsplans.CreateSavingsPlanOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SavingsplansDescribeSavingsPlanRatesResult struct {
	Result workflow.Future
}

func (r *SavingsplansDescribeSavingsPlanRatesResult) Get(ctx workflow.Context) (*savingsplans.DescribeSavingsPlanRatesOutput, error) {
    var output savingsplans.DescribeSavingsPlanRatesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SavingsplansDescribeSavingsPlansResult struct {
	Result workflow.Future
}

func (r *SavingsplansDescribeSavingsPlansResult) Get(ctx workflow.Context) (*savingsplans.DescribeSavingsPlansOutput, error) {
    var output savingsplans.DescribeSavingsPlansOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SavingsplansDescribeSavingsPlansOfferingRatesResult struct {
	Result workflow.Future
}

func (r *SavingsplansDescribeSavingsPlansOfferingRatesResult) Get(ctx workflow.Context) (*savingsplans.DescribeSavingsPlansOfferingRatesOutput, error) {
    var output savingsplans.DescribeSavingsPlansOfferingRatesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SavingsplansDescribeSavingsPlansOfferingsResult struct {
	Result workflow.Future
}

func (r *SavingsplansDescribeSavingsPlansOfferingsResult) Get(ctx workflow.Context) (*savingsplans.DescribeSavingsPlansOfferingsOutput, error) {
    var output savingsplans.DescribeSavingsPlansOfferingsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SavingsplansListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *SavingsplansListTagsForResourceResult) Get(ctx workflow.Context) (*savingsplans.ListTagsForResourceOutput, error) {
    var output savingsplans.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SavingsplansTagResourceResult struct {
	Result workflow.Future
}

func (r *SavingsplansTagResourceResult) Get(ctx workflow.Context) (*savingsplans.TagResourceOutput, error) {
    var output savingsplans.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SavingsplansUntagResourceResult struct {
	Result workflow.Future
}

func (r *SavingsplansUntagResourceResult) Get(ctx workflow.Context) (*savingsplans.UntagResourceOutput, error) {
    var output savingsplans.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SavingsPlansStub struct {
    activities awsactivities.SavingsPlansActivities
}

func NewSavingsPlansStub() SavingsPlansClient {
    return &SavingsPlansStub{}
}

func (a *SavingsPlansStub) CreateSavingsPlan(ctx workflow.Context, input *savingsplans.CreateSavingsPlanInput) (*savingsplans.CreateSavingsPlanOutput, error) {
    var output savingsplans.CreateSavingsPlanOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateSavingsPlan, input).Get(ctx, &output)
    return &output, err
}

func (a *SavingsPlansStub) CreateSavingsPlanAsync(ctx workflow.Context, input *savingsplans.CreateSavingsPlanInput) *SavingsplansCreateSavingsPlanResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateSavingsPlan, input)
    return &SavingsplansCreateSavingsPlanResult{Result: future}
}

func (a *SavingsPlansStub) DescribeSavingsPlanRates(ctx workflow.Context, input *savingsplans.DescribeSavingsPlanRatesInput) (*savingsplans.DescribeSavingsPlanRatesOutput, error) {
    var output savingsplans.DescribeSavingsPlanRatesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeSavingsPlanRates, input).Get(ctx, &output)
    return &output, err
}

func (a *SavingsPlansStub) DescribeSavingsPlanRatesAsync(ctx workflow.Context, input *savingsplans.DescribeSavingsPlanRatesInput) *SavingsplansDescribeSavingsPlanRatesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeSavingsPlanRates, input)
    return &SavingsplansDescribeSavingsPlanRatesResult{Result: future}
}

func (a *SavingsPlansStub) DescribeSavingsPlans(ctx workflow.Context, input *savingsplans.DescribeSavingsPlansInput) (*savingsplans.DescribeSavingsPlansOutput, error) {
    var output savingsplans.DescribeSavingsPlansOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeSavingsPlans, input).Get(ctx, &output)
    return &output, err
}

func (a *SavingsPlansStub) DescribeSavingsPlansAsync(ctx workflow.Context, input *savingsplans.DescribeSavingsPlansInput) *SavingsplansDescribeSavingsPlansResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeSavingsPlans, input)
    return &SavingsplansDescribeSavingsPlansResult{Result: future}
}

func (a *SavingsPlansStub) DescribeSavingsPlansOfferingRates(ctx workflow.Context, input *savingsplans.DescribeSavingsPlansOfferingRatesInput) (*savingsplans.DescribeSavingsPlansOfferingRatesOutput, error) {
    var output savingsplans.DescribeSavingsPlansOfferingRatesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeSavingsPlansOfferingRates, input).Get(ctx, &output)
    return &output, err
}

func (a *SavingsPlansStub) DescribeSavingsPlansOfferingRatesAsync(ctx workflow.Context, input *savingsplans.DescribeSavingsPlansOfferingRatesInput) *SavingsplansDescribeSavingsPlansOfferingRatesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeSavingsPlansOfferingRates, input)
    return &SavingsplansDescribeSavingsPlansOfferingRatesResult{Result: future}
}

func (a *SavingsPlansStub) DescribeSavingsPlansOfferings(ctx workflow.Context, input *savingsplans.DescribeSavingsPlansOfferingsInput) (*savingsplans.DescribeSavingsPlansOfferingsOutput, error) {
    var output savingsplans.DescribeSavingsPlansOfferingsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeSavingsPlansOfferings, input).Get(ctx, &output)
    return &output, err
}

func (a *SavingsPlansStub) DescribeSavingsPlansOfferingsAsync(ctx workflow.Context, input *savingsplans.DescribeSavingsPlansOfferingsInput) *SavingsplansDescribeSavingsPlansOfferingsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DescribeSavingsPlansOfferings, input)
    return &SavingsplansDescribeSavingsPlansOfferingsResult{Result: future}
}

func (a *SavingsPlansStub) ListTagsForResource(ctx workflow.Context, input *savingsplans.ListTagsForResourceInput) (*savingsplans.ListTagsForResourceOutput, error) {
    var output savingsplans.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *SavingsPlansStub) ListTagsForResourceAsync(ctx workflow.Context, input *savingsplans.ListTagsForResourceInput) *SavingsplansListTagsForResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input)
    return &SavingsplansListTagsForResourceResult{Result: future}
}

func (a *SavingsPlansStub) TagResource(ctx workflow.Context, input *savingsplans.TagResourceInput) (*savingsplans.TagResourceOutput, error) {
    var output savingsplans.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *SavingsPlansStub) TagResourceAsync(ctx workflow.Context, input *savingsplans.TagResourceInput) *SavingsplansTagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TagResource, input)
    return &SavingsplansTagResourceResult{Result: future}
}

func (a *SavingsPlansStub) UntagResource(ctx workflow.Context, input *savingsplans.UntagResourceInput) (*savingsplans.UntagResourceOutput, error) {
    var output savingsplans.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *SavingsPlansStub) UntagResourceAsync(ctx workflow.Context, input *savingsplans.UntagResourceInput) *SavingsplansUntagResourceResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input)
    return &SavingsplansUntagResourceResult{Result: future}
}