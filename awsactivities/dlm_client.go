package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/dlm"
	"go.temporal.io/sdk/workflow"
)

type DLMClient interface {
    CreateLifecyclePolicy(ctx workflow.Context, input *dlm.CreateLifecyclePolicyInput) (*dlm.CreateLifecyclePolicyOutput, error)
    CreateLifecyclePolicyAsync(ctx workflow.Context, input *dlm.CreateLifecyclePolicyInput) *DlmCreateLifecyclePolicyResult

    DeleteLifecyclePolicy(ctx workflow.Context, input *dlm.DeleteLifecyclePolicyInput) (*dlm.DeleteLifecyclePolicyOutput, error)
    DeleteLifecyclePolicyAsync(ctx workflow.Context, input *dlm.DeleteLifecyclePolicyInput) *DlmDeleteLifecyclePolicyResult

    GetLifecyclePolicies(ctx workflow.Context, input *dlm.GetLifecyclePoliciesInput) (*dlm.GetLifecyclePoliciesOutput, error)
    GetLifecyclePoliciesAsync(ctx workflow.Context, input *dlm.GetLifecyclePoliciesInput) *DlmGetLifecyclePoliciesResult

    GetLifecyclePolicy(ctx workflow.Context, input *dlm.GetLifecyclePolicyInput) (*dlm.GetLifecyclePolicyOutput, error)
    GetLifecyclePolicyAsync(ctx workflow.Context, input *dlm.GetLifecyclePolicyInput) *DlmGetLifecyclePolicyResult

    ListTagsForResource(ctx workflow.Context, input *dlm.ListTagsForResourceInput) (*dlm.ListTagsForResourceOutput, error)
    ListTagsForResourceAsync(ctx workflow.Context, input *dlm.ListTagsForResourceInput) *DlmListTagsForResourceResult

    TagResource(ctx workflow.Context, input *dlm.TagResourceInput) (*dlm.TagResourceOutput, error)
    TagResourceAsync(ctx workflow.Context, input *dlm.TagResourceInput) *DlmTagResourceResult

    UntagResource(ctx workflow.Context, input *dlm.UntagResourceInput) (*dlm.UntagResourceOutput, error)
    UntagResourceAsync(ctx workflow.Context, input *dlm.UntagResourceInput) *DlmUntagResourceResult

    UpdateLifecyclePolicy(ctx workflow.Context, input *dlm.UpdateLifecyclePolicyInput) (*dlm.UpdateLifecyclePolicyOutput, error)
    UpdateLifecyclePolicyAsync(ctx workflow.Context, input *dlm.UpdateLifecyclePolicyInput) *DlmUpdateLifecyclePolicyResult
}
type DlmCreateLifecyclePolicyResult struct {
	Result workflow.Future
}

func (r *DlmCreateLifecyclePolicyResult) Get(ctx workflow.Context) (*dlm.CreateLifecyclePolicyOutput, error) {
    var output dlm.CreateLifecyclePolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DlmDeleteLifecyclePolicyResult struct {
	Result workflow.Future
}

func (r *DlmDeleteLifecyclePolicyResult) Get(ctx workflow.Context) (*dlm.DeleteLifecyclePolicyOutput, error) {
    var output dlm.DeleteLifecyclePolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DlmGetLifecyclePoliciesResult struct {
	Result workflow.Future
}

func (r *DlmGetLifecyclePoliciesResult) Get(ctx workflow.Context) (*dlm.GetLifecyclePoliciesOutput, error) {
    var output dlm.GetLifecyclePoliciesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DlmGetLifecyclePolicyResult struct {
	Result workflow.Future
}

func (r *DlmGetLifecyclePolicyResult) Get(ctx workflow.Context) (*dlm.GetLifecyclePolicyOutput, error) {
    var output dlm.GetLifecyclePolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DlmListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *DlmListTagsForResourceResult) Get(ctx workflow.Context) (*dlm.ListTagsForResourceOutput, error) {
    var output dlm.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DlmTagResourceResult struct {
	Result workflow.Future
}

func (r *DlmTagResourceResult) Get(ctx workflow.Context) (*dlm.TagResourceOutput, error) {
    var output dlm.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DlmUntagResourceResult struct {
	Result workflow.Future
}

func (r *DlmUntagResourceResult) Get(ctx workflow.Context) (*dlm.UntagResourceOutput, error) {
    var output dlm.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type DlmUpdateLifecyclePolicyResult struct {
	Result workflow.Future
}

func (r *DlmUpdateLifecyclePolicyResult) Get(ctx workflow.Context) (*dlm.UpdateLifecyclePolicyOutput, error) {
    var output dlm.UpdateLifecyclePolicyOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type DLMStub struct {
    activities DLMClient
}

func NewDLMStub() DLMClient {
    return &DLMStub{}
}

func (a *DLMStub) CreateLifecyclePolicy(ctx workflow.Context, input *dlm.CreateLifecyclePolicyInput) (*dlm.CreateLifecyclePolicyOutput, error) {
    var output dlm.CreateLifecyclePolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateLifecyclePolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *DLMStub) DeleteLifecyclePolicy(ctx workflow.Context, input *dlm.DeleteLifecyclePolicyInput) (*dlm.DeleteLifecyclePolicyOutput, error) {
    var output dlm.DeleteLifecyclePolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteLifecyclePolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *DLMStub) GetLifecyclePolicies(ctx workflow.Context, input *dlm.GetLifecyclePoliciesInput) (*dlm.GetLifecyclePoliciesOutput, error) {
    var output dlm.GetLifecyclePoliciesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetLifecyclePolicies, input).Get(ctx, &output)
    return &output, err
}

func (a *DLMStub) GetLifecyclePolicy(ctx workflow.Context, input *dlm.GetLifecyclePolicyInput) (*dlm.GetLifecyclePolicyOutput, error) {
    var output dlm.GetLifecyclePolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetLifecyclePolicy, input).Get(ctx, &output)
    return &output, err
}

func (a *DLMStub) ListTagsForResource(ctx workflow.Context, input *dlm.ListTagsForResourceInput) (*dlm.ListTagsForResourceOutput, error) {
    var output dlm.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *DLMStub) TagResource(ctx workflow.Context, input *dlm.TagResourceInput) (*dlm.TagResourceOutput, error) {
    var output dlm.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *DLMStub) UntagResource(ctx workflow.Context, input *dlm.UntagResourceInput) (*dlm.UntagResourceOutput, error) {
    var output dlm.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *DLMStub) UpdateLifecyclePolicy(ctx workflow.Context, input *dlm.UpdateLifecyclePolicyInput) (*dlm.UpdateLifecyclePolicyOutput, error) {
    var output dlm.UpdateLifecyclePolicyOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateLifecyclePolicy, input).Get(ctx, &output)
    return &output, err
}