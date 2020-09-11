package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/polly"
	"github.com/aws/aws-sdk-go/service/polly/pollyiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type PollyActivities struct {
	client pollyiface.PollyAPI
}

func NewPollyActivities(session *session.Session, config ...*aws.Config) *PollyActivities {
	client := polly.New(session, config...)
	return &PollyActivities{client: client}
}

func (a *PollyActivities) DeleteLexicon(ctx context.Context, input *polly.DeleteLexiconInput) (*polly.DeleteLexiconOutput, error) {
	return a.client.DeleteLexiconWithContext(ctx, input)
}

func (a *PollyActivities) DescribeVoices(ctx context.Context, input *polly.DescribeVoicesInput) (*polly.DescribeVoicesOutput, error) {
	return a.client.DescribeVoicesWithContext(ctx, input)
}

func (a *PollyActivities) GetLexicon(ctx context.Context, input *polly.GetLexiconInput) (*polly.GetLexiconOutput, error) {
	return a.client.GetLexiconWithContext(ctx, input)
}

func (a *PollyActivities) GetSpeechSynthesisTask(ctx context.Context, input *polly.GetSpeechSynthesisTaskInput) (*polly.GetSpeechSynthesisTaskOutput, error) {
	return a.client.GetSpeechSynthesisTaskWithContext(ctx, input)
}

func (a *PollyActivities) ListLexicons(ctx context.Context, input *polly.ListLexiconsInput) (*polly.ListLexiconsOutput, error) {
	return a.client.ListLexiconsWithContext(ctx, input)
}

func (a *PollyActivities) ListSpeechSynthesisTasks(ctx context.Context, input *polly.ListSpeechSynthesisTasksInput) (*polly.ListSpeechSynthesisTasksOutput, error) {
	return a.client.ListSpeechSynthesisTasksWithContext(ctx, input)
}

func (a *PollyActivities) PutLexicon(ctx context.Context, input *polly.PutLexiconInput) (*polly.PutLexiconOutput, error) {
	return a.client.PutLexiconWithContext(ctx, input)
}

func (a *PollyActivities) StartSpeechSynthesisTask(ctx context.Context, input *polly.StartSpeechSynthesisTaskInput) (*polly.StartSpeechSynthesisTaskOutput, error) {
	return a.client.StartSpeechSynthesisTaskWithContext(ctx, input)
}

func (a *PollyActivities) SynthesizeSpeech(ctx context.Context, input *polly.SynthesizeSpeechInput) (*polly.SynthesizeSpeechOutput, error) {
	return a.client.SynthesizeSpeechWithContext(ctx, input)
}
