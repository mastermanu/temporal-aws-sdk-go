package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/codegurureviewer"
	"go.temporal.io/sdk/workflow"
)

type CodeGuruReviewerClient interface {
    AssociateRepository(ctx workflow.Context, input *codegurureviewer.AssociateRepositoryInput) (*codegurureviewer.AssociateRepositoryOutput, error)
    AssociateRepositoryAsync(ctx workflow.Context, input *codegurureviewer.AssociateRepositoryInput) *CodegurureviewerAssociateRepositoryResult

    DescribeCodeReview(ctx workflow.Context, input *codegurureviewer.DescribeCodeReviewInput) (*codegurureviewer.DescribeCodeReviewOutput, error)
    DescribeCodeReviewAsync(ctx workflow.Context, input *codegurureviewer.DescribeCodeReviewInput) *CodegurureviewerDescribeCodeReviewResult

    DescribeRecommendationFeedback(ctx workflow.Context, input *codegurureviewer.DescribeRecommendationFeedbackInput) (*codegurureviewer.DescribeRecommendationFeedbackOutput, error)
    DescribeRecommendationFeedbackAsync(ctx workflow.Context, input *codegurureviewer.DescribeRecommendationFeedbackInput) *CodegurureviewerDescribeRecommendationFeedbackResult

    DescribeRepositoryAssociation(ctx workflow.Context, input *codegurureviewer.DescribeRepositoryAssociationInput) (*codegurureviewer.DescribeRepositoryAssociationOutput, error)
    DescribeRepositoryAssociationAsync(ctx workflow.Context, input *codegurureviewer.DescribeRepositoryAssociationInput) *CodegurureviewerDescribeRepositoryAssociationResult

    DisassociateRepository(ctx workflow.Context, input *codegurureviewer.DisassociateRepositoryInput) (*codegurureviewer.DisassociateRepositoryOutput, error)
    DisassociateRepositoryAsync(ctx workflow.Context, input *codegurureviewer.DisassociateRepositoryInput) *CodegurureviewerDisassociateRepositoryResult

    ListCodeReviews(ctx workflow.Context, input *codegurureviewer.ListCodeReviewsInput) (*codegurureviewer.ListCodeReviewsOutput, error)
    ListCodeReviewsAsync(ctx workflow.Context, input *codegurureviewer.ListCodeReviewsInput) *CodegurureviewerListCodeReviewsResult

    ListRecommendationFeedback(ctx workflow.Context, input *codegurureviewer.ListRecommendationFeedbackInput) (*codegurureviewer.ListRecommendationFeedbackOutput, error)
    ListRecommendationFeedbackAsync(ctx workflow.Context, input *codegurureviewer.ListRecommendationFeedbackInput) *CodegurureviewerListRecommendationFeedbackResult

    ListRecommendations(ctx workflow.Context, input *codegurureviewer.ListRecommendationsInput) (*codegurureviewer.ListRecommendationsOutput, error)
    ListRecommendationsAsync(ctx workflow.Context, input *codegurureviewer.ListRecommendationsInput) *CodegurureviewerListRecommendationsResult

    ListRepositoryAssociations(ctx workflow.Context, input *codegurureviewer.ListRepositoryAssociationsInput) (*codegurureviewer.ListRepositoryAssociationsOutput, error)
    ListRepositoryAssociationsAsync(ctx workflow.Context, input *codegurureviewer.ListRepositoryAssociationsInput) *CodegurureviewerListRepositoryAssociationsResult

    PutRecommendationFeedback(ctx workflow.Context, input *codegurureviewer.PutRecommendationFeedbackInput) (*codegurureviewer.PutRecommendationFeedbackOutput, error)
    PutRecommendationFeedbackAsync(ctx workflow.Context, input *codegurureviewer.PutRecommendationFeedbackInput) *CodegurureviewerPutRecommendationFeedbackResult
}
type CodegurureviewerAssociateRepositoryResult struct {
	Result workflow.Future
}

func (r *CodegurureviewerAssociateRepositoryResult) Get(ctx workflow.Context) (*codegurureviewer.AssociateRepositoryOutput, error) {
    var output codegurureviewer.AssociateRepositoryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodegurureviewerDescribeCodeReviewResult struct {
	Result workflow.Future
}

func (r *CodegurureviewerDescribeCodeReviewResult) Get(ctx workflow.Context) (*codegurureviewer.DescribeCodeReviewOutput, error) {
    var output codegurureviewer.DescribeCodeReviewOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodegurureviewerDescribeRecommendationFeedbackResult struct {
	Result workflow.Future
}

func (r *CodegurureviewerDescribeRecommendationFeedbackResult) Get(ctx workflow.Context) (*codegurureviewer.DescribeRecommendationFeedbackOutput, error) {
    var output codegurureviewer.DescribeRecommendationFeedbackOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodegurureviewerDescribeRepositoryAssociationResult struct {
	Result workflow.Future
}

func (r *CodegurureviewerDescribeRepositoryAssociationResult) Get(ctx workflow.Context) (*codegurureviewer.DescribeRepositoryAssociationOutput, error) {
    var output codegurureviewer.DescribeRepositoryAssociationOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodegurureviewerDisassociateRepositoryResult struct {
	Result workflow.Future
}

func (r *CodegurureviewerDisassociateRepositoryResult) Get(ctx workflow.Context) (*codegurureviewer.DisassociateRepositoryOutput, error) {
    var output codegurureviewer.DisassociateRepositoryOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodegurureviewerListCodeReviewsResult struct {
	Result workflow.Future
}

func (r *CodegurureviewerListCodeReviewsResult) Get(ctx workflow.Context) (*codegurureviewer.ListCodeReviewsOutput, error) {
    var output codegurureviewer.ListCodeReviewsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodegurureviewerListRecommendationFeedbackResult struct {
	Result workflow.Future
}

func (r *CodegurureviewerListRecommendationFeedbackResult) Get(ctx workflow.Context) (*codegurureviewer.ListRecommendationFeedbackOutput, error) {
    var output codegurureviewer.ListRecommendationFeedbackOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodegurureviewerListRecommendationsResult struct {
	Result workflow.Future
}

func (r *CodegurureviewerListRecommendationsResult) Get(ctx workflow.Context) (*codegurureviewer.ListRecommendationsOutput, error) {
    var output codegurureviewer.ListRecommendationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodegurureviewerListRepositoryAssociationsResult struct {
	Result workflow.Future
}

func (r *CodegurureviewerListRepositoryAssociationsResult) Get(ctx workflow.Context) (*codegurureviewer.ListRepositoryAssociationsOutput, error) {
    var output codegurureviewer.ListRepositoryAssociationsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type CodegurureviewerPutRecommendationFeedbackResult struct {
	Result workflow.Future
}

func (r *CodegurureviewerPutRecommendationFeedbackResult) Get(ctx workflow.Context) (*codegurureviewer.PutRecommendationFeedbackOutput, error) {
    var output codegurureviewer.PutRecommendationFeedbackOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type CodeGuruReviewerStub struct {
    activities CodeGuruReviewerClient
}

func NewCodeGuruReviewerStub() CodeGuruReviewerClient {
    return &CodeGuruReviewerStub{}
}

func (a *CodeGuruReviewerStub) AssociateRepository(ctx workflow.Context, input *codegurureviewer.AssociateRepositoryInput) (*codegurureviewer.AssociateRepositoryOutput, error) {
    var output codegurureviewer.AssociateRepositoryOutput
    err := workflow.ExecuteActivity(ctx, a.activities.AssociateRepository, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeGuruReviewerStub) DescribeCodeReview(ctx workflow.Context, input *codegurureviewer.DescribeCodeReviewInput) (*codegurureviewer.DescribeCodeReviewOutput, error) {
    var output codegurureviewer.DescribeCodeReviewOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeCodeReview, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeGuruReviewerStub) DescribeRecommendationFeedback(ctx workflow.Context, input *codegurureviewer.DescribeRecommendationFeedbackInput) (*codegurureviewer.DescribeRecommendationFeedbackOutput, error) {
    var output codegurureviewer.DescribeRecommendationFeedbackOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeRecommendationFeedback, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeGuruReviewerStub) DescribeRepositoryAssociation(ctx workflow.Context, input *codegurureviewer.DescribeRepositoryAssociationInput) (*codegurureviewer.DescribeRepositoryAssociationOutput, error) {
    var output codegurureviewer.DescribeRepositoryAssociationOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeRepositoryAssociation, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeGuruReviewerStub) DisassociateRepository(ctx workflow.Context, input *codegurureviewer.DisassociateRepositoryInput) (*codegurureviewer.DisassociateRepositoryOutput, error) {
    var output codegurureviewer.DisassociateRepositoryOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DisassociateRepository, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeGuruReviewerStub) ListCodeReviews(ctx workflow.Context, input *codegurureviewer.ListCodeReviewsInput) (*codegurureviewer.ListCodeReviewsOutput, error) {
    var output codegurureviewer.ListCodeReviewsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListCodeReviews, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeGuruReviewerStub) ListRecommendationFeedback(ctx workflow.Context, input *codegurureviewer.ListRecommendationFeedbackInput) (*codegurureviewer.ListRecommendationFeedbackOutput, error) {
    var output codegurureviewer.ListRecommendationFeedbackOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListRecommendationFeedback, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeGuruReviewerStub) ListRecommendations(ctx workflow.Context, input *codegurureviewer.ListRecommendationsInput) (*codegurureviewer.ListRecommendationsOutput, error) {
    var output codegurureviewer.ListRecommendationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListRecommendations, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeGuruReviewerStub) ListRepositoryAssociations(ctx workflow.Context, input *codegurureviewer.ListRepositoryAssociationsInput) (*codegurureviewer.ListRepositoryAssociationsOutput, error) {
    var output codegurureviewer.ListRepositoryAssociationsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListRepositoryAssociations, input).Get(ctx, &output)
    return &output, err
}

func (a *CodeGuruReviewerStub) PutRecommendationFeedback(ctx workflow.Context, input *codegurureviewer.PutRecommendationFeedbackInput) (*codegurureviewer.PutRecommendationFeedbackOutput, error) {
    var output codegurureviewer.PutRecommendationFeedbackOutput
    err := workflow.ExecuteActivity(ctx, a.activities.PutRecommendationFeedback, input).Get(ctx, &output)
    return &output, err
}