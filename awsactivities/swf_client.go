package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/swf"
	"go.temporal.io/sdk/workflow"
)

type SWFClient interface {
    CountClosedWorkflowExecutions(ctx workflow.Context, input *swf.CountClosedWorkflowExecutionsInput) (*swf.WorkflowExecutionCount, error)
    CountClosedWorkflowExecutionsAsync(ctx workflow.Context, input *swf.CountClosedWorkflowExecutionsInput) *SwfCountClosedWorkflowExecutionsResult

    CountOpenWorkflowExecutions(ctx workflow.Context, input *swf.CountOpenWorkflowExecutionsInput) (*swf.WorkflowExecutionCount, error)
    CountOpenWorkflowExecutionsAsync(ctx workflow.Context, input *swf.CountOpenWorkflowExecutionsInput) *SwfCountOpenWorkflowExecutionsResult

    CountPendingActivityTasks(ctx workflow.Context, input *swf.CountPendingActivityTasksInput) (*swf.PendingTaskCount, error)
    CountPendingActivityTasksAsync(ctx workflow.Context, input *swf.CountPendingActivityTasksInput) *SwfCountPendingActivityTasksResult

    CountPendingDecisionTasks(ctx workflow.Context, input *swf.CountPendingDecisionTasksInput) (*swf.PendingTaskCount, error)
    CountPendingDecisionTasksAsync(ctx workflow.Context, input *swf.CountPendingDecisionTasksInput) *SwfCountPendingDecisionTasksResult

    DeprecateActivityType(ctx workflow.Context, input *swf.DeprecateActivityTypeInput) (*swf.DeprecateActivityTypeOutput, error)
    DeprecateActivityTypeAsync(ctx workflow.Context, input *swf.DeprecateActivityTypeInput) *SwfDeprecateActivityTypeResult

    DeprecateDomain(ctx workflow.Context, input *swf.DeprecateDomainInput) (*swf.DeprecateDomainOutput, error)
    DeprecateDomainAsync(ctx workflow.Context, input *swf.DeprecateDomainInput) *SwfDeprecateDomainResult

    DeprecateWorkflowType(ctx workflow.Context, input *swf.DeprecateWorkflowTypeInput) (*swf.DeprecateWorkflowTypeOutput, error)
    DeprecateWorkflowTypeAsync(ctx workflow.Context, input *swf.DeprecateWorkflowTypeInput) *SwfDeprecateWorkflowTypeResult

    DescribeActivityType(ctx workflow.Context, input *swf.DescribeActivityTypeInput) (*swf.DescribeActivityTypeOutput, error)
    DescribeActivityTypeAsync(ctx workflow.Context, input *swf.DescribeActivityTypeInput) *SwfDescribeActivityTypeResult

    DescribeDomain(ctx workflow.Context, input *swf.DescribeDomainInput) (*swf.DescribeDomainOutput, error)
    DescribeDomainAsync(ctx workflow.Context, input *swf.DescribeDomainInput) *SwfDescribeDomainResult

    DescribeWorkflowExecution(ctx workflow.Context, input *swf.DescribeWorkflowExecutionInput) (*swf.DescribeWorkflowExecutionOutput, error)
    DescribeWorkflowExecutionAsync(ctx workflow.Context, input *swf.DescribeWorkflowExecutionInput) *SwfDescribeWorkflowExecutionResult

    DescribeWorkflowType(ctx workflow.Context, input *swf.DescribeWorkflowTypeInput) (*swf.DescribeWorkflowTypeOutput, error)
    DescribeWorkflowTypeAsync(ctx workflow.Context, input *swf.DescribeWorkflowTypeInput) *SwfDescribeWorkflowTypeResult

    GetWorkflowExecutionHistory(ctx workflow.Context, input *swf.GetWorkflowExecutionHistoryInput) (*swf.GetWorkflowExecutionHistoryOutput, error)
    GetWorkflowExecutionHistoryAsync(ctx workflow.Context, input *swf.GetWorkflowExecutionHistoryInput) *SwfGetWorkflowExecutionHistoryResult

    ListActivityTypes(ctx workflow.Context, input *swf.ListActivityTypesInput) (*swf.ListActivityTypesOutput, error)
    ListActivityTypesAsync(ctx workflow.Context, input *swf.ListActivityTypesInput) *SwfListActivityTypesResult

    ListClosedWorkflowExecutions(ctx workflow.Context, input *swf.ListClosedWorkflowExecutionsInput) (*swf.WorkflowExecutionInfos, error)
    ListClosedWorkflowExecutionsAsync(ctx workflow.Context, input *swf.ListClosedWorkflowExecutionsInput) *SwfListClosedWorkflowExecutionsResult

    ListDomains(ctx workflow.Context, input *swf.ListDomainsInput) (*swf.ListDomainsOutput, error)
    ListDomainsAsync(ctx workflow.Context, input *swf.ListDomainsInput) *SwfListDomainsResult

    ListOpenWorkflowExecutions(ctx workflow.Context, input *swf.ListOpenWorkflowExecutionsInput) (*swf.WorkflowExecutionInfos, error)
    ListOpenWorkflowExecutionsAsync(ctx workflow.Context, input *swf.ListOpenWorkflowExecutionsInput) *SwfListOpenWorkflowExecutionsResult

    ListTagsForResource(ctx workflow.Context, input *swf.ListTagsForResourceInput) (*swf.ListTagsForResourceOutput, error)
    ListTagsForResourceAsync(ctx workflow.Context, input *swf.ListTagsForResourceInput) *SwfListTagsForResourceResult

    ListWorkflowTypes(ctx workflow.Context, input *swf.ListWorkflowTypesInput) (*swf.ListWorkflowTypesOutput, error)
    ListWorkflowTypesAsync(ctx workflow.Context, input *swf.ListWorkflowTypesInput) *SwfListWorkflowTypesResult

    PollForActivityTask(ctx workflow.Context, input *swf.PollForActivityTaskInput) (*swf.PollForActivityTaskOutput, error)
    PollForActivityTaskAsync(ctx workflow.Context, input *swf.PollForActivityTaskInput) *SwfPollForActivityTaskResult

    PollForDecisionTask(ctx workflow.Context, input *swf.PollForDecisionTaskInput) (*swf.PollForDecisionTaskOutput, error)
    PollForDecisionTaskAsync(ctx workflow.Context, input *swf.PollForDecisionTaskInput) *SwfPollForDecisionTaskResult

    RecordActivityTaskHeartbeat(ctx workflow.Context, input *swf.RecordActivityTaskHeartbeatInput) (*swf.RecordActivityTaskHeartbeatOutput, error)
    RecordActivityTaskHeartbeatAsync(ctx workflow.Context, input *swf.RecordActivityTaskHeartbeatInput) *SwfRecordActivityTaskHeartbeatResult

    RegisterActivityType(ctx workflow.Context, input *swf.RegisterActivityTypeInput) (*swf.RegisterActivityTypeOutput, error)
    RegisterActivityTypeAsync(ctx workflow.Context, input *swf.RegisterActivityTypeInput) *SwfRegisterActivityTypeResult

    RegisterDomain(ctx workflow.Context, input *swf.RegisterDomainInput) (*swf.RegisterDomainOutput, error)
    RegisterDomainAsync(ctx workflow.Context, input *swf.RegisterDomainInput) *SwfRegisterDomainResult

    RegisterWorkflowType(ctx workflow.Context, input *swf.RegisterWorkflowTypeInput) (*swf.RegisterWorkflowTypeOutput, error)
    RegisterWorkflowTypeAsync(ctx workflow.Context, input *swf.RegisterWorkflowTypeInput) *SwfRegisterWorkflowTypeResult

    RequestCancelWorkflowExecution(ctx workflow.Context, input *swf.RequestCancelWorkflowExecutionInput) (*swf.RequestCancelWorkflowExecutionOutput, error)
    RequestCancelWorkflowExecutionAsync(ctx workflow.Context, input *swf.RequestCancelWorkflowExecutionInput) *SwfRequestCancelWorkflowExecutionResult

    RespondActivityTaskCanceled(ctx workflow.Context, input *swf.RespondActivityTaskCanceledInput) (*swf.RespondActivityTaskCanceledOutput, error)
    RespondActivityTaskCanceledAsync(ctx workflow.Context, input *swf.RespondActivityTaskCanceledInput) *SwfRespondActivityTaskCanceledResult

    RespondActivityTaskCompleted(ctx workflow.Context, input *swf.RespondActivityTaskCompletedInput) (*swf.RespondActivityTaskCompletedOutput, error)
    RespondActivityTaskCompletedAsync(ctx workflow.Context, input *swf.RespondActivityTaskCompletedInput) *SwfRespondActivityTaskCompletedResult

    RespondActivityTaskFailed(ctx workflow.Context, input *swf.RespondActivityTaskFailedInput) (*swf.RespondActivityTaskFailedOutput, error)
    RespondActivityTaskFailedAsync(ctx workflow.Context, input *swf.RespondActivityTaskFailedInput) *SwfRespondActivityTaskFailedResult

    RespondDecisionTaskCompleted(ctx workflow.Context, input *swf.RespondDecisionTaskCompletedInput) (*swf.RespondDecisionTaskCompletedOutput, error)
    RespondDecisionTaskCompletedAsync(ctx workflow.Context, input *swf.RespondDecisionTaskCompletedInput) *SwfRespondDecisionTaskCompletedResult

    SignalWorkflowExecution(ctx workflow.Context, input *swf.SignalWorkflowExecutionInput) (*swf.SignalWorkflowExecutionOutput, error)
    SignalWorkflowExecutionAsync(ctx workflow.Context, input *swf.SignalWorkflowExecutionInput) *SwfSignalWorkflowExecutionResult

    StartWorkflowExecution(ctx workflow.Context, input *swf.StartWorkflowExecutionInput) (*swf.StartWorkflowExecutionOutput, error)
    StartWorkflowExecutionAsync(ctx workflow.Context, input *swf.StartWorkflowExecutionInput) *SwfStartWorkflowExecutionResult

    TagResource(ctx workflow.Context, input *swf.TagResourceInput) (*swf.TagResourceOutput, error)
    TagResourceAsync(ctx workflow.Context, input *swf.TagResourceInput) *SwfTagResourceResult

    TerminateWorkflowExecution(ctx workflow.Context, input *swf.TerminateWorkflowExecutionInput) (*swf.TerminateWorkflowExecutionOutput, error)
    TerminateWorkflowExecutionAsync(ctx workflow.Context, input *swf.TerminateWorkflowExecutionInput) *SwfTerminateWorkflowExecutionResult

    UndeprecateActivityType(ctx workflow.Context, input *swf.UndeprecateActivityTypeInput) (*swf.UndeprecateActivityTypeOutput, error)
    UndeprecateActivityTypeAsync(ctx workflow.Context, input *swf.UndeprecateActivityTypeInput) *SwfUndeprecateActivityTypeResult

    UndeprecateDomain(ctx workflow.Context, input *swf.UndeprecateDomainInput) (*swf.UndeprecateDomainOutput, error)
    UndeprecateDomainAsync(ctx workflow.Context, input *swf.UndeprecateDomainInput) *SwfUndeprecateDomainResult

    UndeprecateWorkflowType(ctx workflow.Context, input *swf.UndeprecateWorkflowTypeInput) (*swf.UndeprecateWorkflowTypeOutput, error)
    UndeprecateWorkflowTypeAsync(ctx workflow.Context, input *swf.UndeprecateWorkflowTypeInput) *SwfUndeprecateWorkflowTypeResult

    UntagResource(ctx workflow.Context, input *swf.UntagResourceInput) (*swf.UntagResourceOutput, error)
    UntagResourceAsync(ctx workflow.Context, input *swf.UntagResourceInput) *SwfUntagResourceResult
}
type SwfCountClosedWorkflowExecutionsResult struct {
	Result workflow.Future
}

func (r *SwfCountClosedWorkflowExecutionsResult) Get(ctx workflow.Context) (*swf.WorkflowExecutionCount, error) {
    var output swf.WorkflowExecutionCount
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SwfCountOpenWorkflowExecutionsResult struct {
	Result workflow.Future
}

func (r *SwfCountOpenWorkflowExecutionsResult) Get(ctx workflow.Context) (*swf.WorkflowExecutionCount, error) {
    var output swf.WorkflowExecutionCount
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SwfCountPendingActivityTasksResult struct {
	Result workflow.Future
}

func (r *SwfCountPendingActivityTasksResult) Get(ctx workflow.Context) (*swf.PendingTaskCount, error) {
    var output swf.PendingTaskCount
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SwfCountPendingDecisionTasksResult struct {
	Result workflow.Future
}

func (r *SwfCountPendingDecisionTasksResult) Get(ctx workflow.Context) (*swf.PendingTaskCount, error) {
    var output swf.PendingTaskCount
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SwfDeprecateActivityTypeResult struct {
	Result workflow.Future
}

func (r *SwfDeprecateActivityTypeResult) Get(ctx workflow.Context) (*swf.DeprecateActivityTypeOutput, error) {
    var output swf.DeprecateActivityTypeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SwfDeprecateDomainResult struct {
	Result workflow.Future
}

func (r *SwfDeprecateDomainResult) Get(ctx workflow.Context) (*swf.DeprecateDomainOutput, error) {
    var output swf.DeprecateDomainOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SwfDeprecateWorkflowTypeResult struct {
	Result workflow.Future
}

func (r *SwfDeprecateWorkflowTypeResult) Get(ctx workflow.Context) (*swf.DeprecateWorkflowTypeOutput, error) {
    var output swf.DeprecateWorkflowTypeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SwfDescribeActivityTypeResult struct {
	Result workflow.Future
}

func (r *SwfDescribeActivityTypeResult) Get(ctx workflow.Context) (*swf.DescribeActivityTypeOutput, error) {
    var output swf.DescribeActivityTypeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SwfDescribeDomainResult struct {
	Result workflow.Future
}

func (r *SwfDescribeDomainResult) Get(ctx workflow.Context) (*swf.DescribeDomainOutput, error) {
    var output swf.DescribeDomainOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SwfDescribeWorkflowExecutionResult struct {
	Result workflow.Future
}

func (r *SwfDescribeWorkflowExecutionResult) Get(ctx workflow.Context) (*swf.DescribeWorkflowExecutionOutput, error) {
    var output swf.DescribeWorkflowExecutionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SwfDescribeWorkflowTypeResult struct {
	Result workflow.Future
}

func (r *SwfDescribeWorkflowTypeResult) Get(ctx workflow.Context) (*swf.DescribeWorkflowTypeOutput, error) {
    var output swf.DescribeWorkflowTypeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SwfGetWorkflowExecutionHistoryResult struct {
	Result workflow.Future
}

func (r *SwfGetWorkflowExecutionHistoryResult) Get(ctx workflow.Context) (*swf.GetWorkflowExecutionHistoryOutput, error) {
    var output swf.GetWorkflowExecutionHistoryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SwfListActivityTypesResult struct {
	Result workflow.Future
}

func (r *SwfListActivityTypesResult) Get(ctx workflow.Context) (*swf.ListActivityTypesOutput, error) {
    var output swf.ListActivityTypesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SwfListClosedWorkflowExecutionsResult struct {
	Result workflow.Future
}

func (r *SwfListClosedWorkflowExecutionsResult) Get(ctx workflow.Context) (*swf.WorkflowExecutionInfos, error) {
    var output swf.WorkflowExecutionInfos
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SwfListDomainsResult struct {
	Result workflow.Future
}

func (r *SwfListDomainsResult) Get(ctx workflow.Context) (*swf.ListDomainsOutput, error) {
    var output swf.ListDomainsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SwfListOpenWorkflowExecutionsResult struct {
	Result workflow.Future
}

func (r *SwfListOpenWorkflowExecutionsResult) Get(ctx workflow.Context) (*swf.WorkflowExecutionInfos, error) {
    var output swf.WorkflowExecutionInfos
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SwfListTagsForResourceResult struct {
	Result workflow.Future
}

func (r *SwfListTagsForResourceResult) Get(ctx workflow.Context) (*swf.ListTagsForResourceOutput, error) {
    var output swf.ListTagsForResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SwfListWorkflowTypesResult struct {
	Result workflow.Future
}

func (r *SwfListWorkflowTypesResult) Get(ctx workflow.Context) (*swf.ListWorkflowTypesOutput, error) {
    var output swf.ListWorkflowTypesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SwfPollForActivityTaskResult struct {
	Result workflow.Future
}

func (r *SwfPollForActivityTaskResult) Get(ctx workflow.Context) (*swf.PollForActivityTaskOutput, error) {
    var output swf.PollForActivityTaskOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SwfPollForDecisionTaskResult struct {
	Result workflow.Future
}

func (r *SwfPollForDecisionTaskResult) Get(ctx workflow.Context) (*swf.PollForDecisionTaskOutput, error) {
    var output swf.PollForDecisionTaskOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SwfRecordActivityTaskHeartbeatResult struct {
	Result workflow.Future
}

func (r *SwfRecordActivityTaskHeartbeatResult) Get(ctx workflow.Context) (*swf.RecordActivityTaskHeartbeatOutput, error) {
    var output swf.RecordActivityTaskHeartbeatOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SwfRegisterActivityTypeResult struct {
	Result workflow.Future
}

func (r *SwfRegisterActivityTypeResult) Get(ctx workflow.Context) (*swf.RegisterActivityTypeOutput, error) {
    var output swf.RegisterActivityTypeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SwfRegisterDomainResult struct {
	Result workflow.Future
}

func (r *SwfRegisterDomainResult) Get(ctx workflow.Context) (*swf.RegisterDomainOutput, error) {
    var output swf.RegisterDomainOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SwfRegisterWorkflowTypeResult struct {
	Result workflow.Future
}

func (r *SwfRegisterWorkflowTypeResult) Get(ctx workflow.Context) (*swf.RegisterWorkflowTypeOutput, error) {
    var output swf.RegisterWorkflowTypeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SwfRequestCancelWorkflowExecutionResult struct {
	Result workflow.Future
}

func (r *SwfRequestCancelWorkflowExecutionResult) Get(ctx workflow.Context) (*swf.RequestCancelWorkflowExecutionOutput, error) {
    var output swf.RequestCancelWorkflowExecutionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SwfRespondActivityTaskCanceledResult struct {
	Result workflow.Future
}

func (r *SwfRespondActivityTaskCanceledResult) Get(ctx workflow.Context) (*swf.RespondActivityTaskCanceledOutput, error) {
    var output swf.RespondActivityTaskCanceledOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SwfRespondActivityTaskCompletedResult struct {
	Result workflow.Future
}

func (r *SwfRespondActivityTaskCompletedResult) Get(ctx workflow.Context) (*swf.RespondActivityTaskCompletedOutput, error) {
    var output swf.RespondActivityTaskCompletedOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SwfRespondActivityTaskFailedResult struct {
	Result workflow.Future
}

func (r *SwfRespondActivityTaskFailedResult) Get(ctx workflow.Context) (*swf.RespondActivityTaskFailedOutput, error) {
    var output swf.RespondActivityTaskFailedOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SwfRespondDecisionTaskCompletedResult struct {
	Result workflow.Future
}

func (r *SwfRespondDecisionTaskCompletedResult) Get(ctx workflow.Context) (*swf.RespondDecisionTaskCompletedOutput, error) {
    var output swf.RespondDecisionTaskCompletedOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SwfSignalWorkflowExecutionResult struct {
	Result workflow.Future
}

func (r *SwfSignalWorkflowExecutionResult) Get(ctx workflow.Context) (*swf.SignalWorkflowExecutionOutput, error) {
    var output swf.SignalWorkflowExecutionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SwfStartWorkflowExecutionResult struct {
	Result workflow.Future
}

func (r *SwfStartWorkflowExecutionResult) Get(ctx workflow.Context) (*swf.StartWorkflowExecutionOutput, error) {
    var output swf.StartWorkflowExecutionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SwfTagResourceResult struct {
	Result workflow.Future
}

func (r *SwfTagResourceResult) Get(ctx workflow.Context) (*swf.TagResourceOutput, error) {
    var output swf.TagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SwfTerminateWorkflowExecutionResult struct {
	Result workflow.Future
}

func (r *SwfTerminateWorkflowExecutionResult) Get(ctx workflow.Context) (*swf.TerminateWorkflowExecutionOutput, error) {
    var output swf.TerminateWorkflowExecutionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SwfUndeprecateActivityTypeResult struct {
	Result workflow.Future
}

func (r *SwfUndeprecateActivityTypeResult) Get(ctx workflow.Context) (*swf.UndeprecateActivityTypeOutput, error) {
    var output swf.UndeprecateActivityTypeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SwfUndeprecateDomainResult struct {
	Result workflow.Future
}

func (r *SwfUndeprecateDomainResult) Get(ctx workflow.Context) (*swf.UndeprecateDomainOutput, error) {
    var output swf.UndeprecateDomainOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SwfUndeprecateWorkflowTypeResult struct {
	Result workflow.Future
}

func (r *SwfUndeprecateWorkflowTypeResult) Get(ctx workflow.Context) (*swf.UndeprecateWorkflowTypeOutput, error) {
    var output swf.UndeprecateWorkflowTypeOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SwfUntagResourceResult struct {
	Result workflow.Future
}

func (r *SwfUntagResourceResult) Get(ctx workflow.Context) (*swf.UntagResourceOutput, error) {
    var output swf.UntagResourceOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type SWFStub struct {
    activities SWFClient
}

func NewSWFStub() SWFClient {
    return &SWFStub{}
}

func (a *SWFStub) CountClosedWorkflowExecutions(ctx workflow.Context, input *swf.CountClosedWorkflowExecutionsInput) (*swf.WorkflowExecutionCount, error) {
    var output swf.WorkflowExecutionCount
    err := workflow.ExecuteActivity(ctx, a.activities.CountClosedWorkflowExecutions, input).Get(ctx, &output)
    return &output, err
}

func (a *SWFStub) CountOpenWorkflowExecutions(ctx workflow.Context, input *swf.CountOpenWorkflowExecutionsInput) (*swf.WorkflowExecutionCount, error) {
    var output swf.WorkflowExecutionCount
    err := workflow.ExecuteActivity(ctx, a.activities.CountOpenWorkflowExecutions, input).Get(ctx, &output)
    return &output, err
}

func (a *SWFStub) CountPendingActivityTasks(ctx workflow.Context, input *swf.CountPendingActivityTasksInput) (*swf.PendingTaskCount, error) {
    var output swf.PendingTaskCount
    err := workflow.ExecuteActivity(ctx, a.activities.CountPendingActivityTasks, input).Get(ctx, &output)
    return &output, err
}

func (a *SWFStub) CountPendingDecisionTasks(ctx workflow.Context, input *swf.CountPendingDecisionTasksInput) (*swf.PendingTaskCount, error) {
    var output swf.PendingTaskCount
    err := workflow.ExecuteActivity(ctx, a.activities.CountPendingDecisionTasks, input).Get(ctx, &output)
    return &output, err
}

func (a *SWFStub) DeprecateActivityType(ctx workflow.Context, input *swf.DeprecateActivityTypeInput) (*swf.DeprecateActivityTypeOutput, error) {
    var output swf.DeprecateActivityTypeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeprecateActivityType, input).Get(ctx, &output)
    return &output, err
}

func (a *SWFStub) DeprecateDomain(ctx workflow.Context, input *swf.DeprecateDomainInput) (*swf.DeprecateDomainOutput, error) {
    var output swf.DeprecateDomainOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeprecateDomain, input).Get(ctx, &output)
    return &output, err
}

func (a *SWFStub) DeprecateWorkflowType(ctx workflow.Context, input *swf.DeprecateWorkflowTypeInput) (*swf.DeprecateWorkflowTypeOutput, error) {
    var output swf.DeprecateWorkflowTypeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeprecateWorkflowType, input).Get(ctx, &output)
    return &output, err
}

func (a *SWFStub) DescribeActivityType(ctx workflow.Context, input *swf.DescribeActivityTypeInput) (*swf.DescribeActivityTypeOutput, error) {
    var output swf.DescribeActivityTypeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeActivityType, input).Get(ctx, &output)
    return &output, err
}

func (a *SWFStub) DescribeDomain(ctx workflow.Context, input *swf.DescribeDomainInput) (*swf.DescribeDomainOutput, error) {
    var output swf.DescribeDomainOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeDomain, input).Get(ctx, &output)
    return &output, err
}

func (a *SWFStub) DescribeWorkflowExecution(ctx workflow.Context, input *swf.DescribeWorkflowExecutionInput) (*swf.DescribeWorkflowExecutionOutput, error) {
    var output swf.DescribeWorkflowExecutionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeWorkflowExecution, input).Get(ctx, &output)
    return &output, err
}

func (a *SWFStub) DescribeWorkflowType(ctx workflow.Context, input *swf.DescribeWorkflowTypeInput) (*swf.DescribeWorkflowTypeOutput, error) {
    var output swf.DescribeWorkflowTypeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeWorkflowType, input).Get(ctx, &output)
    return &output, err
}

func (a *SWFStub) GetWorkflowExecutionHistory(ctx workflow.Context, input *swf.GetWorkflowExecutionHistoryInput) (*swf.GetWorkflowExecutionHistoryOutput, error) {
    var output swf.GetWorkflowExecutionHistoryOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetWorkflowExecutionHistory, input).Get(ctx, &output)
    return &output, err
}

func (a *SWFStub) ListActivityTypes(ctx workflow.Context, input *swf.ListActivityTypesInput) (*swf.ListActivityTypesOutput, error) {
    var output swf.ListActivityTypesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListActivityTypes, input).Get(ctx, &output)
    return &output, err
}

func (a *SWFStub) ListClosedWorkflowExecutions(ctx workflow.Context, input *swf.ListClosedWorkflowExecutionsInput) (*swf.WorkflowExecutionInfos, error) {
    var output swf.WorkflowExecutionInfos
    err := workflow.ExecuteActivity(ctx, a.activities.ListClosedWorkflowExecutions, input).Get(ctx, &output)
    return &output, err
}

func (a *SWFStub) ListDomains(ctx workflow.Context, input *swf.ListDomainsInput) (*swf.ListDomainsOutput, error) {
    var output swf.ListDomainsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDomains, input).Get(ctx, &output)
    return &output, err
}

func (a *SWFStub) ListOpenWorkflowExecutions(ctx workflow.Context, input *swf.ListOpenWorkflowExecutionsInput) (*swf.WorkflowExecutionInfos, error) {
    var output swf.WorkflowExecutionInfos
    err := workflow.ExecuteActivity(ctx, a.activities.ListOpenWorkflowExecutions, input).Get(ctx, &output)
    return &output, err
}

func (a *SWFStub) ListTagsForResource(ctx workflow.Context, input *swf.ListTagsForResourceInput) (*swf.ListTagsForResourceOutput, error) {
    var output swf.ListTagsForResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListTagsForResource, input).Get(ctx, &output)
    return &output, err
}

func (a *SWFStub) ListWorkflowTypes(ctx workflow.Context, input *swf.ListWorkflowTypesInput) (*swf.ListWorkflowTypesOutput, error) {
    var output swf.ListWorkflowTypesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListWorkflowTypes, input).Get(ctx, &output)
    return &output, err
}

func (a *SWFStub) PollForActivityTask(ctx workflow.Context, input *swf.PollForActivityTaskInput) (*swf.PollForActivityTaskOutput, error) {
    var output swf.PollForActivityTaskOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PollForActivityTask, input).Get(ctx, &output)
    return &output, err
}

func (a *SWFStub) PollForDecisionTask(ctx workflow.Context, input *swf.PollForDecisionTaskInput) (*swf.PollForDecisionTaskOutput, error) {
    var output swf.PollForDecisionTaskOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PollForDecisionTask, input).Get(ctx, &output)
    return &output, err
}

func (a *SWFStub) RecordActivityTaskHeartbeat(ctx workflow.Context, input *swf.RecordActivityTaskHeartbeatInput) (*swf.RecordActivityTaskHeartbeatOutput, error) {
    var output swf.RecordActivityTaskHeartbeatOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RecordActivityTaskHeartbeat, input).Get(ctx, &output)
    return &output, err
}

func (a *SWFStub) RegisterActivityType(ctx workflow.Context, input *swf.RegisterActivityTypeInput) (*swf.RegisterActivityTypeOutput, error) {
    var output swf.RegisterActivityTypeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RegisterActivityType, input).Get(ctx, &output)
    return &output, err
}

func (a *SWFStub) RegisterDomain(ctx workflow.Context, input *swf.RegisterDomainInput) (*swf.RegisterDomainOutput, error) {
    var output swf.RegisterDomainOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RegisterDomain, input).Get(ctx, &output)
    return &output, err
}

func (a *SWFStub) RegisterWorkflowType(ctx workflow.Context, input *swf.RegisterWorkflowTypeInput) (*swf.RegisterWorkflowTypeOutput, error) {
    var output swf.RegisterWorkflowTypeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RegisterWorkflowType, input).Get(ctx, &output)
    return &output, err
}

func (a *SWFStub) RequestCancelWorkflowExecution(ctx workflow.Context, input *swf.RequestCancelWorkflowExecutionInput) (*swf.RequestCancelWorkflowExecutionOutput, error) {
    var output swf.RequestCancelWorkflowExecutionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RequestCancelWorkflowExecution, input).Get(ctx, &output)
    return &output, err
}

func (a *SWFStub) RespondActivityTaskCanceled(ctx workflow.Context, input *swf.RespondActivityTaskCanceledInput) (*swf.RespondActivityTaskCanceledOutput, error) {
    var output swf.RespondActivityTaskCanceledOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RespondActivityTaskCanceled, input).Get(ctx, &output)
    return &output, err
}

func (a *SWFStub) RespondActivityTaskCompleted(ctx workflow.Context, input *swf.RespondActivityTaskCompletedInput) (*swf.RespondActivityTaskCompletedOutput, error) {
    var output swf.RespondActivityTaskCompletedOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RespondActivityTaskCompleted, input).Get(ctx, &output)
    return &output, err
}

func (a *SWFStub) RespondActivityTaskFailed(ctx workflow.Context, input *swf.RespondActivityTaskFailedInput) (*swf.RespondActivityTaskFailedOutput, error) {
    var output swf.RespondActivityTaskFailedOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RespondActivityTaskFailed, input).Get(ctx, &output)
    return &output, err
}

func (a *SWFStub) RespondDecisionTaskCompleted(ctx workflow.Context, input *swf.RespondDecisionTaskCompletedInput) (*swf.RespondDecisionTaskCompletedOutput, error) {
    var output swf.RespondDecisionTaskCompletedOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RespondDecisionTaskCompleted, input).Get(ctx, &output)
    return &output, err
}

func (a *SWFStub) SignalWorkflowExecution(ctx workflow.Context, input *swf.SignalWorkflowExecutionInput) (*swf.SignalWorkflowExecutionOutput, error) {
    var output swf.SignalWorkflowExecutionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SignalWorkflowExecution, input).Get(ctx, &output)
    return &output, err
}

func (a *SWFStub) StartWorkflowExecution(ctx workflow.Context, input *swf.StartWorkflowExecutionInput) (*swf.StartWorkflowExecutionOutput, error) {
    var output swf.StartWorkflowExecutionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartWorkflowExecution, input).Get(ctx, &output)
    return &output, err
}

func (a *SWFStub) TagResource(ctx workflow.Context, input *swf.TagResourceInput) (*swf.TagResourceOutput, error) {
    var output swf.TagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagResource, input).Get(ctx, &output)
    return &output, err
}

func (a *SWFStub) TerminateWorkflowExecution(ctx workflow.Context, input *swf.TerminateWorkflowExecutionInput) (*swf.TerminateWorkflowExecutionOutput, error) {
    var output swf.TerminateWorkflowExecutionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TerminateWorkflowExecution, input).Get(ctx, &output)
    return &output, err
}

func (a *SWFStub) UndeprecateActivityType(ctx workflow.Context, input *swf.UndeprecateActivityTypeInput) (*swf.UndeprecateActivityTypeOutput, error) {
    var output swf.UndeprecateActivityTypeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UndeprecateActivityType, input).Get(ctx, &output)
    return &output, err
}

func (a *SWFStub) UndeprecateDomain(ctx workflow.Context, input *swf.UndeprecateDomainInput) (*swf.UndeprecateDomainOutput, error) {
    var output swf.UndeprecateDomainOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UndeprecateDomain, input).Get(ctx, &output)
    return &output, err
}

func (a *SWFStub) UndeprecateWorkflowType(ctx workflow.Context, input *swf.UndeprecateWorkflowTypeInput) (*swf.UndeprecateWorkflowTypeOutput, error) {
    var output swf.UndeprecateWorkflowTypeOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UndeprecateWorkflowType, input).Get(ctx, &output)
    return &output, err
}

func (a *SWFStub) UntagResource(ctx workflow.Context, input *swf.UntagResourceInput) (*swf.UntagResourceOutput, error) {
    var output swf.UntagResourceOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagResource, input).Get(ctx, &output)
    return &output, err
}