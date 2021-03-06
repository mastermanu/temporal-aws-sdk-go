// Generated by github.com/temporalio/temporal-aws-sdk-generator
// from github.com/aws/aws-sdk-go version 1.35.7
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.

package rekognition

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"

	"go.temporal.io/aws-sdk/internal"
	"github.com/aws/aws-sdk-go/service/rekognition"
	"github.com/aws/aws-sdk-go/service/rekognition/rekognitioniface"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

// SessionFactory returns an aws.Session based on the activity context.
type SessionFactory interface {
	Session(ctx context.Context) (*session.Session, error)
}

type Activities struct {
	client rekognitioniface.RekognitionAPI

	sessionFactory SessionFactory
}

func NewActivities(sess *session.Session, config ...*aws.Config) *Activities {
	client := rekognition.New(sess, config...)
	return &Activities{client: client}
}

func NewActivitiesWithSessionFactory(sessionFactory SessionFactory) *Activities {
	return &Activities{sessionFactory: sessionFactory}
}

func (a *Activities) getClient(ctx context.Context) (rekognitioniface.RekognitionAPI, error) {
	if a.client != nil { // No need to protect with mutex: we know the client never changes
		return a.client, nil
	}

	sess, err := a.sessionFactory.Session(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}

	return rekognition.New(sess), nil
}

func (a *Activities) CompareFaces(ctx context.Context, input *rekognition.CompareFacesInput) (*rekognition.CompareFacesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CompareFacesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateCollection(ctx context.Context, input *rekognition.CreateCollectionInput) (*rekognition.CreateCollectionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateCollectionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateProject(ctx context.Context, input *rekognition.CreateProjectInput) (*rekognition.CreateProjectOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateProjectWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateProjectVersion(ctx context.Context, input *rekognition.CreateProjectVersionInput) (*rekognition.CreateProjectVersionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateProjectVersionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) CreateStreamProcessor(ctx context.Context, input *rekognition.CreateStreamProcessorInput) (*rekognition.CreateStreamProcessorOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.CreateStreamProcessorWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteCollection(ctx context.Context, input *rekognition.DeleteCollectionInput) (*rekognition.DeleteCollectionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteCollectionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteFaces(ctx context.Context, input *rekognition.DeleteFacesInput) (*rekognition.DeleteFacesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteFacesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteProject(ctx context.Context, input *rekognition.DeleteProjectInput) (*rekognition.DeleteProjectOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteProjectWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteProjectVersion(ctx context.Context, input *rekognition.DeleteProjectVersionInput) (*rekognition.DeleteProjectVersionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteProjectVersionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DeleteStreamProcessor(ctx context.Context, input *rekognition.DeleteStreamProcessorInput) (*rekognition.DeleteStreamProcessorOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DeleteStreamProcessorWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeCollection(ctx context.Context, input *rekognition.DescribeCollectionInput) (*rekognition.DescribeCollectionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeCollectionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeProjectVersions(ctx context.Context, input *rekognition.DescribeProjectVersionsInput) (*rekognition.DescribeProjectVersionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeProjectVersionsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeProjects(ctx context.Context, input *rekognition.DescribeProjectsInput) (*rekognition.DescribeProjectsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeProjectsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DescribeStreamProcessor(ctx context.Context, input *rekognition.DescribeStreamProcessorInput) (*rekognition.DescribeStreamProcessorOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DescribeStreamProcessorWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DetectCustomLabels(ctx context.Context, input *rekognition.DetectCustomLabelsInput) (*rekognition.DetectCustomLabelsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DetectCustomLabelsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DetectFaces(ctx context.Context, input *rekognition.DetectFacesInput) (*rekognition.DetectFacesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DetectFacesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DetectLabels(ctx context.Context, input *rekognition.DetectLabelsInput) (*rekognition.DetectLabelsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DetectLabelsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DetectModerationLabels(ctx context.Context, input *rekognition.DetectModerationLabelsInput) (*rekognition.DetectModerationLabelsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DetectModerationLabelsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) DetectText(ctx context.Context, input *rekognition.DetectTextInput) (*rekognition.DetectTextOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.DetectTextWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetCelebrityInfo(ctx context.Context, input *rekognition.GetCelebrityInfoInput) (*rekognition.GetCelebrityInfoOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetCelebrityInfoWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetCelebrityRecognition(ctx context.Context, input *rekognition.GetCelebrityRecognitionInput) (*rekognition.GetCelebrityRecognitionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetCelebrityRecognitionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetContentModeration(ctx context.Context, input *rekognition.GetContentModerationInput) (*rekognition.GetContentModerationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetContentModerationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetFaceDetection(ctx context.Context, input *rekognition.GetFaceDetectionInput) (*rekognition.GetFaceDetectionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetFaceDetectionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetFaceSearch(ctx context.Context, input *rekognition.GetFaceSearchInput) (*rekognition.GetFaceSearchOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetFaceSearchWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetLabelDetection(ctx context.Context, input *rekognition.GetLabelDetectionInput) (*rekognition.GetLabelDetectionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetLabelDetectionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetPersonTracking(ctx context.Context, input *rekognition.GetPersonTrackingInput) (*rekognition.GetPersonTrackingOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetPersonTrackingWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetSegmentDetection(ctx context.Context, input *rekognition.GetSegmentDetectionInput) (*rekognition.GetSegmentDetectionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetSegmentDetectionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) GetTextDetection(ctx context.Context, input *rekognition.GetTextDetectionInput) (*rekognition.GetTextDetectionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.GetTextDetectionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) IndexFaces(ctx context.Context, input *rekognition.IndexFacesInput) (*rekognition.IndexFacesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.IndexFacesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListCollections(ctx context.Context, input *rekognition.ListCollectionsInput) (*rekognition.ListCollectionsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListCollectionsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListFaces(ctx context.Context, input *rekognition.ListFacesInput) (*rekognition.ListFacesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListFacesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) ListStreamProcessors(ctx context.Context, input *rekognition.ListStreamProcessorsInput) (*rekognition.ListStreamProcessorsOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.ListStreamProcessorsWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) RecognizeCelebrities(ctx context.Context, input *rekognition.RecognizeCelebritiesInput) (*rekognition.RecognizeCelebritiesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.RecognizeCelebritiesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) SearchFaces(ctx context.Context, input *rekognition.SearchFacesInput) (*rekognition.SearchFacesOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.SearchFacesWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) SearchFacesByImage(ctx context.Context, input *rekognition.SearchFacesByImageInput) (*rekognition.SearchFacesByImageOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.SearchFacesByImageWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StartCelebrityRecognition(ctx context.Context, input *rekognition.StartCelebrityRecognitionInput) (*rekognition.StartCelebrityRecognitionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StartCelebrityRecognitionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StartContentModeration(ctx context.Context, input *rekognition.StartContentModerationInput) (*rekognition.StartContentModerationOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StartContentModerationWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StartFaceDetection(ctx context.Context, input *rekognition.StartFaceDetectionInput) (*rekognition.StartFaceDetectionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StartFaceDetectionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StartFaceSearch(ctx context.Context, input *rekognition.StartFaceSearchInput) (*rekognition.StartFaceSearchOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StartFaceSearchWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StartLabelDetection(ctx context.Context, input *rekognition.StartLabelDetectionInput) (*rekognition.StartLabelDetectionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StartLabelDetectionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StartPersonTracking(ctx context.Context, input *rekognition.StartPersonTrackingInput) (*rekognition.StartPersonTrackingOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StartPersonTrackingWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StartProjectVersion(ctx context.Context, input *rekognition.StartProjectVersionInput) (*rekognition.StartProjectVersionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StartProjectVersionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StartSegmentDetection(ctx context.Context, input *rekognition.StartSegmentDetectionInput) (*rekognition.StartSegmentDetectionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StartSegmentDetectionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StartStreamProcessor(ctx context.Context, input *rekognition.StartStreamProcessorInput) (*rekognition.StartStreamProcessorOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StartStreamProcessorWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StartTextDetection(ctx context.Context, input *rekognition.StartTextDetectionInput) (*rekognition.StartTextDetectionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StartTextDetectionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StopProjectVersion(ctx context.Context, input *rekognition.StopProjectVersionInput) (*rekognition.StopProjectVersionOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StopProjectVersionWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) StopStreamProcessor(ctx context.Context, input *rekognition.StopStreamProcessorInput) (*rekognition.StopStreamProcessorOutput, error) {
	client, err := a.getClient(ctx)
	if err != nil {
		return nil, internal.EncodeError(err)
	}
	output, err := client.StopStreamProcessorWithContext(ctx, input)

	return output, internal.EncodeError(err)
}

func (a *Activities) WaitUntilProjectVersionRunning(ctx context.Context, input *rekognition.DescribeProjectVersionsInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return internal.EncodeError(err)
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return internal.EncodeError(client.WaitUntilProjectVersionRunningWithContext(ctx, input, options...))
	})
}

func (a *Activities) WaitUntilProjectVersionTrainingCompleted(ctx context.Context, input *rekognition.DescribeProjectVersionsInput) error {
	client, err := a.getClient(ctx)
	if err != nil {
		return internal.EncodeError(err)
	}
	return internal.WaitUntilActivity(ctx, func(ctx context.Context, options ...request.WaiterOption) error {
		return internal.EncodeError(client.WaitUntilProjectVersionTrainingCompletedWithContext(ctx, input, options...))
	})
}
