package awsclients

import (
	"github.com/aws/aws-sdk-go/service/sqs"
	"go.temporal.io/sdk/workflow"
	"temporal.io/aws-sdk/awsactivities"
)

type SQSClient interface {
    AddPermission(ctx workflow.Context, input *sqs.AddPermissionInput) (*sqs.AddPermissionOutput, error)
    AddPermissionAsync(ctx workflow.Context, input *sqs.AddPermissionInput) *SqsAddPermissionResult

    ChangeMessageVisibility(ctx workflow.Context, input *sqs.ChangeMessageVisibilityInput) (*sqs.ChangeMessageVisibilityOutput, error)
    ChangeMessageVisibilityAsync(ctx workflow.Context, input *sqs.ChangeMessageVisibilityInput) *SqsChangeMessageVisibilityResult

    ChangeMessageVisibilityBatch(ctx workflow.Context, input *sqs.ChangeMessageVisibilityBatchInput) (*sqs.ChangeMessageVisibilityBatchOutput, error)
    ChangeMessageVisibilityBatchAsync(ctx workflow.Context, input *sqs.ChangeMessageVisibilityBatchInput) *SqsChangeMessageVisibilityBatchResult

    CreateQueue(ctx workflow.Context, input *sqs.CreateQueueInput) (*sqs.CreateQueueOutput, error)
    CreateQueueAsync(ctx workflow.Context, input *sqs.CreateQueueInput) *SqsCreateQueueResult

    DeleteMessage(ctx workflow.Context, input *sqs.DeleteMessageInput) (*sqs.DeleteMessageOutput, error)
    DeleteMessageAsync(ctx workflow.Context, input *sqs.DeleteMessageInput) *SqsDeleteMessageResult

    DeleteMessageBatch(ctx workflow.Context, input *sqs.DeleteMessageBatchInput) (*sqs.DeleteMessageBatchOutput, error)
    DeleteMessageBatchAsync(ctx workflow.Context, input *sqs.DeleteMessageBatchInput) *SqsDeleteMessageBatchResult

    DeleteQueue(ctx workflow.Context, input *sqs.DeleteQueueInput) (*sqs.DeleteQueueOutput, error)
    DeleteQueueAsync(ctx workflow.Context, input *sqs.DeleteQueueInput) *SqsDeleteQueueResult

    GetQueueAttributes(ctx workflow.Context, input *sqs.GetQueueAttributesInput) (*sqs.GetQueueAttributesOutput, error)
    GetQueueAttributesAsync(ctx workflow.Context, input *sqs.GetQueueAttributesInput) *SqsGetQueueAttributesResult

    GetQueueUrl(ctx workflow.Context, input *sqs.GetQueueUrlInput) (*sqs.GetQueueUrlOutput, error)
    GetQueueUrlAsync(ctx workflow.Context, input *sqs.GetQueueUrlInput) *SqsGetQueueUrlResult

    ListDeadLetterSourceQueues(ctx workflow.Context, input *sqs.ListDeadLetterSourceQueuesInput) (*sqs.ListDeadLetterSourceQueuesOutput, error)
    ListDeadLetterSourceQueuesAsync(ctx workflow.Context, input *sqs.ListDeadLetterSourceQueuesInput) *SqsListDeadLetterSourceQueuesResult

    ListQueueTags(ctx workflow.Context, input *sqs.ListQueueTagsInput) (*sqs.ListQueueTagsOutput, error)
    ListQueueTagsAsync(ctx workflow.Context, input *sqs.ListQueueTagsInput) *SqsListQueueTagsResult

    ListQueues(ctx workflow.Context, input *sqs.ListQueuesInput) (*sqs.ListQueuesOutput, error)
    ListQueuesAsync(ctx workflow.Context, input *sqs.ListQueuesInput) *SqsListQueuesResult

    PurgeQueue(ctx workflow.Context, input *sqs.PurgeQueueInput) (*sqs.PurgeQueueOutput, error)
    PurgeQueueAsync(ctx workflow.Context, input *sqs.PurgeQueueInput) *SqsPurgeQueueResult

    ReceiveMessage(ctx workflow.Context, input *sqs.ReceiveMessageInput) (*sqs.ReceiveMessageOutput, error)
    ReceiveMessageAsync(ctx workflow.Context, input *sqs.ReceiveMessageInput) *SqsReceiveMessageResult

    RemovePermission(ctx workflow.Context, input *sqs.RemovePermissionInput) (*sqs.RemovePermissionOutput, error)
    RemovePermissionAsync(ctx workflow.Context, input *sqs.RemovePermissionInput) *SqsRemovePermissionResult

    SendMessage(ctx workflow.Context, input *sqs.SendMessageInput) (*sqs.SendMessageOutput, error)
    SendMessageAsync(ctx workflow.Context, input *sqs.SendMessageInput) *SqsSendMessageResult

    SendMessageBatch(ctx workflow.Context, input *sqs.SendMessageBatchInput) (*sqs.SendMessageBatchOutput, error)
    SendMessageBatchAsync(ctx workflow.Context, input *sqs.SendMessageBatchInput) *SqsSendMessageBatchResult

    SetQueueAttributes(ctx workflow.Context, input *sqs.SetQueueAttributesInput) (*sqs.SetQueueAttributesOutput, error)
    SetQueueAttributesAsync(ctx workflow.Context, input *sqs.SetQueueAttributesInput) *SqsSetQueueAttributesResult

    TagQueue(ctx workflow.Context, input *sqs.TagQueueInput) (*sqs.TagQueueOutput, error)
    TagQueueAsync(ctx workflow.Context, input *sqs.TagQueueInput) *SqsTagQueueResult

    UntagQueue(ctx workflow.Context, input *sqs.UntagQueueInput) (*sqs.UntagQueueOutput, error)
    UntagQueueAsync(ctx workflow.Context, input *sqs.UntagQueueInput) *SqsUntagQueueResult
}

type SqsAddPermissionResult struct {
	Result workflow.Future
}

func (r *SqsAddPermissionResult) Get(ctx workflow.Context) (*sqs.AddPermissionOutput, error) {
    var output sqs.AddPermissionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SqsChangeMessageVisibilityResult struct {
	Result workflow.Future
}

func (r *SqsChangeMessageVisibilityResult) Get(ctx workflow.Context) (*sqs.ChangeMessageVisibilityOutput, error) {
    var output sqs.ChangeMessageVisibilityOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SqsChangeMessageVisibilityBatchResult struct {
	Result workflow.Future
}

func (r *SqsChangeMessageVisibilityBatchResult) Get(ctx workflow.Context) (*sqs.ChangeMessageVisibilityBatchOutput, error) {
    var output sqs.ChangeMessageVisibilityBatchOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SqsCreateQueueResult struct {
	Result workflow.Future
}

func (r *SqsCreateQueueResult) Get(ctx workflow.Context) (*sqs.CreateQueueOutput, error) {
    var output sqs.CreateQueueOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SqsDeleteMessageResult struct {
	Result workflow.Future
}

func (r *SqsDeleteMessageResult) Get(ctx workflow.Context) (*sqs.DeleteMessageOutput, error) {
    var output sqs.DeleteMessageOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SqsDeleteMessageBatchResult struct {
	Result workflow.Future
}

func (r *SqsDeleteMessageBatchResult) Get(ctx workflow.Context) (*sqs.DeleteMessageBatchOutput, error) {
    var output sqs.DeleteMessageBatchOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SqsDeleteQueueResult struct {
	Result workflow.Future
}

func (r *SqsDeleteQueueResult) Get(ctx workflow.Context) (*sqs.DeleteQueueOutput, error) {
    var output sqs.DeleteQueueOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SqsGetQueueAttributesResult struct {
	Result workflow.Future
}

func (r *SqsGetQueueAttributesResult) Get(ctx workflow.Context) (*sqs.GetQueueAttributesOutput, error) {
    var output sqs.GetQueueAttributesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SqsGetQueueUrlResult struct {
	Result workflow.Future
}

func (r *SqsGetQueueUrlResult) Get(ctx workflow.Context) (*sqs.GetQueueUrlOutput, error) {
    var output sqs.GetQueueUrlOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SqsListDeadLetterSourceQueuesResult struct {
	Result workflow.Future
}

func (r *SqsListDeadLetterSourceQueuesResult) Get(ctx workflow.Context) (*sqs.ListDeadLetterSourceQueuesOutput, error) {
    var output sqs.ListDeadLetterSourceQueuesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SqsListQueueTagsResult struct {
	Result workflow.Future
}

func (r *SqsListQueueTagsResult) Get(ctx workflow.Context) (*sqs.ListQueueTagsOutput, error) {
    var output sqs.ListQueueTagsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SqsListQueuesResult struct {
	Result workflow.Future
}

func (r *SqsListQueuesResult) Get(ctx workflow.Context) (*sqs.ListQueuesOutput, error) {
    var output sqs.ListQueuesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SqsPurgeQueueResult struct {
	Result workflow.Future
}

func (r *SqsPurgeQueueResult) Get(ctx workflow.Context) (*sqs.PurgeQueueOutput, error) {
    var output sqs.PurgeQueueOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SqsReceiveMessageResult struct {
	Result workflow.Future
}

func (r *SqsReceiveMessageResult) Get(ctx workflow.Context) (*sqs.ReceiveMessageOutput, error) {
    var output sqs.ReceiveMessageOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SqsRemovePermissionResult struct {
	Result workflow.Future
}

func (r *SqsRemovePermissionResult) Get(ctx workflow.Context) (*sqs.RemovePermissionOutput, error) {
    var output sqs.RemovePermissionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SqsSendMessageResult struct {
	Result workflow.Future
}

func (r *SqsSendMessageResult) Get(ctx workflow.Context) (*sqs.SendMessageOutput, error) {
    var output sqs.SendMessageOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SqsSendMessageBatchResult struct {
	Result workflow.Future
}

func (r *SqsSendMessageBatchResult) Get(ctx workflow.Context) (*sqs.SendMessageBatchOutput, error) {
    var output sqs.SendMessageBatchOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SqsSetQueueAttributesResult struct {
	Result workflow.Future
}

func (r *SqsSetQueueAttributesResult) Get(ctx workflow.Context) (*sqs.SetQueueAttributesOutput, error) {
    var output sqs.SetQueueAttributesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SqsTagQueueResult struct {
	Result workflow.Future
}

func (r *SqsTagQueueResult) Get(ctx workflow.Context) (*sqs.TagQueueOutput, error) {
    var output sqs.TagQueueOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SqsUntagQueueResult struct {
	Result workflow.Future
}

func (r *SqsUntagQueueResult) Get(ctx workflow.Context) (*sqs.UntagQueueOutput, error) {
    var output sqs.UntagQueueOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type SQSStub struct {
    activities awsactivities.SQSActivities
}

func NewSQSStub() SQSClient {
    return &SQSStub{}
}

func (a *SQSStub) AddPermission(ctx workflow.Context, input *sqs.AddPermissionInput) (*sqs.AddPermissionOutput, error) {
    var output sqs.AddPermissionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AddPermission, input).Get(ctx, &output)
    return &output, err
}

func (a *SQSStub) AddPermissionAsync(ctx workflow.Context, input *sqs.AddPermissionInput) *SqsAddPermissionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.AddPermission, input)
    return &SqsAddPermissionResult{Result: future}
}

func (a *SQSStub) ChangeMessageVisibility(ctx workflow.Context, input *sqs.ChangeMessageVisibilityInput) (*sqs.ChangeMessageVisibilityOutput, error) {
    var output sqs.ChangeMessageVisibilityOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ChangeMessageVisibility, input).Get(ctx, &output)
    return &output, err
}

func (a *SQSStub) ChangeMessageVisibilityAsync(ctx workflow.Context, input *sqs.ChangeMessageVisibilityInput) *SqsChangeMessageVisibilityResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ChangeMessageVisibility, input)
    return &SqsChangeMessageVisibilityResult{Result: future}
}

func (a *SQSStub) ChangeMessageVisibilityBatch(ctx workflow.Context, input *sqs.ChangeMessageVisibilityBatchInput) (*sqs.ChangeMessageVisibilityBatchOutput, error) {
    var output sqs.ChangeMessageVisibilityBatchOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ChangeMessageVisibilityBatch, input).Get(ctx, &output)
    return &output, err
}

func (a *SQSStub) ChangeMessageVisibilityBatchAsync(ctx workflow.Context, input *sqs.ChangeMessageVisibilityBatchInput) *SqsChangeMessageVisibilityBatchResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ChangeMessageVisibilityBatch, input)
    return &SqsChangeMessageVisibilityBatchResult{Result: future}
}

func (a *SQSStub) CreateQueue(ctx workflow.Context, input *sqs.CreateQueueInput) (*sqs.CreateQueueOutput, error) {
    var output sqs.CreateQueueOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateQueue, input).Get(ctx, &output)
    return &output, err
}

func (a *SQSStub) CreateQueueAsync(ctx workflow.Context, input *sqs.CreateQueueInput) *SqsCreateQueueResult {
    future := workflow.ExecuteActivity(ctx, a.activities.CreateQueue, input)
    return &SqsCreateQueueResult{Result: future}
}

func (a *SQSStub) DeleteMessage(ctx workflow.Context, input *sqs.DeleteMessageInput) (*sqs.DeleteMessageOutput, error) {
    var output sqs.DeleteMessageOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteMessage, input).Get(ctx, &output)
    return &output, err
}

func (a *SQSStub) DeleteMessageAsync(ctx workflow.Context, input *sqs.DeleteMessageInput) *SqsDeleteMessageResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteMessage, input)
    return &SqsDeleteMessageResult{Result: future}
}

func (a *SQSStub) DeleteMessageBatch(ctx workflow.Context, input *sqs.DeleteMessageBatchInput) (*sqs.DeleteMessageBatchOutput, error) {
    var output sqs.DeleteMessageBatchOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteMessageBatch, input).Get(ctx, &output)
    return &output, err
}

func (a *SQSStub) DeleteMessageBatchAsync(ctx workflow.Context, input *sqs.DeleteMessageBatchInput) *SqsDeleteMessageBatchResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteMessageBatch, input)
    return &SqsDeleteMessageBatchResult{Result: future}
}

func (a *SQSStub) DeleteQueue(ctx workflow.Context, input *sqs.DeleteQueueInput) (*sqs.DeleteQueueOutput, error) {
    var output sqs.DeleteQueueOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteQueue, input).Get(ctx, &output)
    return &output, err
}

func (a *SQSStub) DeleteQueueAsync(ctx workflow.Context, input *sqs.DeleteQueueInput) *SqsDeleteQueueResult {
    future := workflow.ExecuteActivity(ctx, a.activities.DeleteQueue, input)
    return &SqsDeleteQueueResult{Result: future}
}

func (a *SQSStub) GetQueueAttributes(ctx workflow.Context, input *sqs.GetQueueAttributesInput) (*sqs.GetQueueAttributesOutput, error) {
    var output sqs.GetQueueAttributesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetQueueAttributes, input).Get(ctx, &output)
    return &output, err
}

func (a *SQSStub) GetQueueAttributesAsync(ctx workflow.Context, input *sqs.GetQueueAttributesInput) *SqsGetQueueAttributesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetQueueAttributes, input)
    return &SqsGetQueueAttributesResult{Result: future}
}

func (a *SQSStub) GetQueueUrl(ctx workflow.Context, input *sqs.GetQueueUrlInput) (*sqs.GetQueueUrlOutput, error) {
    var output sqs.GetQueueUrlOutput
    err := workflow.ExecuteActivity(ctx, a.activities.GetQueueUrl, input).Get(ctx, &output)
    return &output, err
}

func (a *SQSStub) GetQueueUrlAsync(ctx workflow.Context, input *sqs.GetQueueUrlInput) *SqsGetQueueUrlResult {
    future := workflow.ExecuteActivity(ctx, a.activities.GetQueueUrl, input)
    return &SqsGetQueueUrlResult{Result: future}
}

func (a *SQSStub) ListDeadLetterSourceQueues(ctx workflow.Context, input *sqs.ListDeadLetterSourceQueuesInput) (*sqs.ListDeadLetterSourceQueuesOutput, error) {
    var output sqs.ListDeadLetterSourceQueuesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListDeadLetterSourceQueues, input).Get(ctx, &output)
    return &output, err
}

func (a *SQSStub) ListDeadLetterSourceQueuesAsync(ctx workflow.Context, input *sqs.ListDeadLetterSourceQueuesInput) *SqsListDeadLetterSourceQueuesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListDeadLetterSourceQueues, input)
    return &SqsListDeadLetterSourceQueuesResult{Result: future}
}

func (a *SQSStub) ListQueueTags(ctx workflow.Context, input *sqs.ListQueueTagsInput) (*sqs.ListQueueTagsOutput, error) {
    var output sqs.ListQueueTagsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListQueueTags, input).Get(ctx, &output)
    return &output, err
}

func (a *SQSStub) ListQueueTagsAsync(ctx workflow.Context, input *sqs.ListQueueTagsInput) *SqsListQueueTagsResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListQueueTags, input)
    return &SqsListQueueTagsResult{Result: future}
}

func (a *SQSStub) ListQueues(ctx workflow.Context, input *sqs.ListQueuesInput) (*sqs.ListQueuesOutput, error) {
    var output sqs.ListQueuesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListQueues, input).Get(ctx, &output)
    return &output, err
}

func (a *SQSStub) ListQueuesAsync(ctx workflow.Context, input *sqs.ListQueuesInput) *SqsListQueuesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ListQueues, input)
    return &SqsListQueuesResult{Result: future}
}

func (a *SQSStub) PurgeQueue(ctx workflow.Context, input *sqs.PurgeQueueInput) (*sqs.PurgeQueueOutput, error) {
    var output sqs.PurgeQueueOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PurgeQueue, input).Get(ctx, &output)
    return &output, err
}

func (a *SQSStub) PurgeQueueAsync(ctx workflow.Context, input *sqs.PurgeQueueInput) *SqsPurgeQueueResult {
    future := workflow.ExecuteActivity(ctx, a.activities.PurgeQueue, input)
    return &SqsPurgeQueueResult{Result: future}
}

func (a *SQSStub) ReceiveMessage(ctx workflow.Context, input *sqs.ReceiveMessageInput) (*sqs.ReceiveMessageOutput, error) {
    var output sqs.ReceiveMessageOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ReceiveMessage, input).Get(ctx, &output)
    return &output, err
}

func (a *SQSStub) ReceiveMessageAsync(ctx workflow.Context, input *sqs.ReceiveMessageInput) *SqsReceiveMessageResult {
    future := workflow.ExecuteActivity(ctx, a.activities.ReceiveMessage, input)
    return &SqsReceiveMessageResult{Result: future}
}

func (a *SQSStub) RemovePermission(ctx workflow.Context, input *sqs.RemovePermissionInput) (*sqs.RemovePermissionOutput, error) {
    var output sqs.RemovePermissionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.RemovePermission, input).Get(ctx, &output)
    return &output, err
}

func (a *SQSStub) RemovePermissionAsync(ctx workflow.Context, input *sqs.RemovePermissionInput) *SqsRemovePermissionResult {
    future := workflow.ExecuteActivity(ctx, a.activities.RemovePermission, input)
    return &SqsRemovePermissionResult{Result: future}
}

func (a *SQSStub) SendMessage(ctx workflow.Context, input *sqs.SendMessageInput) (*sqs.SendMessageOutput, error) {
    var output sqs.SendMessageOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SendMessage, input).Get(ctx, &output)
    return &output, err
}

func (a *SQSStub) SendMessageAsync(ctx workflow.Context, input *sqs.SendMessageInput) *SqsSendMessageResult {
    future := workflow.ExecuteActivity(ctx, a.activities.SendMessage, input)
    return &SqsSendMessageResult{Result: future}
}

func (a *SQSStub) SendMessageBatch(ctx workflow.Context, input *sqs.SendMessageBatchInput) (*sqs.SendMessageBatchOutput, error) {
    var output sqs.SendMessageBatchOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SendMessageBatch, input).Get(ctx, &output)
    return &output, err
}

func (a *SQSStub) SendMessageBatchAsync(ctx workflow.Context, input *sqs.SendMessageBatchInput) *SqsSendMessageBatchResult {
    future := workflow.ExecuteActivity(ctx, a.activities.SendMessageBatch, input)
    return &SqsSendMessageBatchResult{Result: future}
}

func (a *SQSStub) SetQueueAttributes(ctx workflow.Context, input *sqs.SetQueueAttributesInput) (*sqs.SetQueueAttributesOutput, error) {
    var output sqs.SetQueueAttributesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.SetQueueAttributes, input).Get(ctx, &output)
    return &output, err
}

func (a *SQSStub) SetQueueAttributesAsync(ctx workflow.Context, input *sqs.SetQueueAttributesInput) *SqsSetQueueAttributesResult {
    future := workflow.ExecuteActivity(ctx, a.activities.SetQueueAttributes, input)
    return &SqsSetQueueAttributesResult{Result: future}
}

func (a *SQSStub) TagQueue(ctx workflow.Context, input *sqs.TagQueueInput) (*sqs.TagQueueOutput, error) {
    var output sqs.TagQueueOutput
    err := workflow.ExecuteActivity(ctx, a.activities.TagQueue, input).Get(ctx, &output)
    return &output, err
}

func (a *SQSStub) TagQueueAsync(ctx workflow.Context, input *sqs.TagQueueInput) *SqsTagQueueResult {
    future := workflow.ExecuteActivity(ctx, a.activities.TagQueue, input)
    return &SqsTagQueueResult{Result: future}
}

func (a *SQSStub) UntagQueue(ctx workflow.Context, input *sqs.UntagQueueInput) (*sqs.UntagQueueOutput, error) {
    var output sqs.UntagQueueOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UntagQueue, input).Get(ctx, &output)
    return &output, err
}

func (a *SQSStub) UntagQueueAsync(ctx workflow.Context, input *sqs.UntagQueueInput) *SqsUntagQueueResult {
    future := workflow.ExecuteActivity(ctx, a.activities.UntagQueue, input)
    return &SqsUntagQueueResult{Result: future}
}