package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/transcribestreamingservice"
	"go.temporal.io/sdk/workflow"
)

type TranscribeStreamingServiceClient interface {
    StartStreamTranscription(ctx workflow.Context, input *transcribestreamingservice.StartStreamTranscriptionInput) (*transcribestreamingservice.StartStreamTranscriptionOutput, error)
    StartStreamTranscriptionAsync(ctx workflow.Context, input *transcribestreamingservice.StartStreamTranscriptionInput) *TranscribestreamingserviceStartStreamTranscriptionResult
}
type TranscribestreamingserviceStartStreamTranscriptionResult struct {
	Result workflow.Future
}

func (r *TranscribestreamingserviceStartStreamTranscriptionResult) Get(ctx workflow.Context) (*transcribestreamingservice.StartStreamTranscriptionOutput, error) {
    var output transcribestreamingservice.StartStreamTranscriptionOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type TranscribeStreamingServiceStub struct {
    activities TranscribeStreamingServiceClient
}

func NewTranscribeStreamingServiceStub() TranscribeStreamingServiceClient {
    return &TranscribeStreamingServiceStub{}
}

func (a *TranscribeStreamingServiceStub) StartStreamTranscription(ctx workflow.Context, input *transcribestreamingservice.StartStreamTranscriptionInput) (*transcribestreamingservice.StartStreamTranscriptionOutput, error) {
    var output transcribestreamingservice.StartStreamTranscriptionOutput
    err := workflow.ExecuteActivity(ctx, a.activities.StartStreamTranscription, input).Get(ctx, &output)
    return &output, err
}