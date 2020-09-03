package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/mobile"
	"go.temporal.io/sdk/workflow"
)

type MobileClient interface {
    CreateProject(ctx workflow.Context, input *mobile.CreateProjectInput) (*mobile.CreateProjectOutput, error)
    CreateProjectAsync(ctx workflow.Context, input *mobile.CreateProjectInput) *MobileCreateProjectResult

    DeleteProject(ctx workflow.Context, input *mobile.DeleteProjectInput) (*mobile.DeleteProjectOutput, error)
    DeleteProjectAsync(ctx workflow.Context, input *mobile.DeleteProjectInput) *MobileDeleteProjectResult

    DescribeBundle(ctx workflow.Context, input *mobile.DescribeBundleInput) (*mobile.DescribeBundleOutput, error)
    DescribeBundleAsync(ctx workflow.Context, input *mobile.DescribeBundleInput) *MobileDescribeBundleResult

    DescribeProject(ctx workflow.Context, input *mobile.DescribeProjectInput) (*mobile.DescribeProjectOutput, error)
    DescribeProjectAsync(ctx workflow.Context, input *mobile.DescribeProjectInput) *MobileDescribeProjectResult

    ExportBundle(ctx workflow.Context, input *mobile.ExportBundleInput) (*mobile.ExportBundleOutput, error)
    ExportBundleAsync(ctx workflow.Context, input *mobile.ExportBundleInput) *MobileExportBundleResult

    ExportProject(ctx workflow.Context, input *mobile.ExportProjectInput) (*mobile.ExportProjectOutput, error)
    ExportProjectAsync(ctx workflow.Context, input *mobile.ExportProjectInput) *MobileExportProjectResult

    ListBundles(ctx workflow.Context, input *mobile.ListBundlesInput) (*mobile.ListBundlesOutput, error)
    ListBundlesAsync(ctx workflow.Context, input *mobile.ListBundlesInput) *MobileListBundlesResult

    ListProjects(ctx workflow.Context, input *mobile.ListProjectsInput) (*mobile.ListProjectsOutput, error)
    ListProjectsAsync(ctx workflow.Context, input *mobile.ListProjectsInput) *MobileListProjectsResult

    UpdateProject(ctx workflow.Context, input *mobile.UpdateProjectInput) (*mobile.UpdateProjectOutput, error)
    UpdateProjectAsync(ctx workflow.Context, input *mobile.UpdateProjectInput) *MobileUpdateProjectResult
}
type MobileCreateProjectResult struct {
	Result workflow.Future
}

func (r *MobileCreateProjectResult) Get(ctx workflow.Context) (*mobile.CreateProjectOutput, error) {
    var output mobile.CreateProjectOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MobileDeleteProjectResult struct {
	Result workflow.Future
}

func (r *MobileDeleteProjectResult) Get(ctx workflow.Context) (*mobile.DeleteProjectOutput, error) {
    var output mobile.DeleteProjectOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MobileDescribeBundleResult struct {
	Result workflow.Future
}

func (r *MobileDescribeBundleResult) Get(ctx workflow.Context) (*mobile.DescribeBundleOutput, error) {
    var output mobile.DescribeBundleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MobileDescribeProjectResult struct {
	Result workflow.Future
}

func (r *MobileDescribeProjectResult) Get(ctx workflow.Context) (*mobile.DescribeProjectOutput, error) {
    var output mobile.DescribeProjectOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MobileExportBundleResult struct {
	Result workflow.Future
}

func (r *MobileExportBundleResult) Get(ctx workflow.Context) (*mobile.ExportBundleOutput, error) {
    var output mobile.ExportBundleOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MobileExportProjectResult struct {
	Result workflow.Future
}

func (r *MobileExportProjectResult) Get(ctx workflow.Context) (*mobile.ExportProjectOutput, error) {
    var output mobile.ExportProjectOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MobileListBundlesResult struct {
	Result workflow.Future
}

func (r *MobileListBundlesResult) Get(ctx workflow.Context) (*mobile.ListBundlesOutput, error) {
    var output mobile.ListBundlesOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MobileListProjectsResult struct {
	Result workflow.Future
}

func (r *MobileListProjectsResult) Get(ctx workflow.Context) (*mobile.ListProjectsOutput, error) {
    var output mobile.ListProjectsOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}

type MobileUpdateProjectResult struct {
	Result workflow.Future
}

func (r *MobileUpdateProjectResult) Get(ctx workflow.Context) (*mobile.UpdateProjectOutput, error) {
    var output mobile.UpdateProjectOutput
    err := r.Result.Get(ctx, &output)
    return &output, err
}


type MobileStub struct {
    activities MobileClient
}

func NewMobileStub() MobileClient {
    return &MobileStub{}
}

func (a *MobileStub) CreateProject(ctx workflow.Context, input *mobile.CreateProjectInput) (*mobile.CreateProjectOutput, error) {
    var output mobile.CreateProjectOutput
    err := workflow.ExecuteActivity(ctx, a.activities.CreateProject, input).Get(ctx, &output)
    return &output, err
}

func (a *MobileStub) DeleteProject(ctx workflow.Context, input *mobile.DeleteProjectInput) (*mobile.DeleteProjectOutput, error) {
    var output mobile.DeleteProjectOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DeleteProject, input).Get(ctx, &output)
    return &output, err
}

func (a *MobileStub) DescribeBundle(ctx workflow.Context, input *mobile.DescribeBundleInput) (*mobile.DescribeBundleOutput, error) {
    var output mobile.DescribeBundleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeBundle, input).Get(ctx, &output)
    return &output, err
}

func (a *MobileStub) DescribeProject(ctx workflow.Context, input *mobile.DescribeProjectInput) (*mobile.DescribeProjectOutput, error) {
    var output mobile.DescribeProjectOutput
    err := workflow.ExecuteActivity(ctx, a.activities.DescribeProject, input).Get(ctx, &output)
    return &output, err
}

func (a *MobileStub) ExportBundle(ctx workflow.Context, input *mobile.ExportBundleInput) (*mobile.ExportBundleOutput, error) {
    var output mobile.ExportBundleOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ExportBundle, input).Get(ctx, &output)
    return &output, err
}

func (a *MobileStub) ExportProject(ctx workflow.Context, input *mobile.ExportProjectInput) (*mobile.ExportProjectOutput, error) {
    var output mobile.ExportProjectOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ExportProject, input).Get(ctx, &output)
    return &output, err
}

func (a *MobileStub) ListBundles(ctx workflow.Context, input *mobile.ListBundlesInput) (*mobile.ListBundlesOutput, error) {
    var output mobile.ListBundlesOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListBundles, input).Get(ctx, &output)
    return &output, err
}

func (a *MobileStub) ListProjects(ctx workflow.Context, input *mobile.ListProjectsInput) (*mobile.ListProjectsOutput, error) {
    var output mobile.ListProjectsOutput
    err := workflow.ExecuteActivity(ctx, a.activities.ListProjects, input).Get(ctx, &output)
    return &output, err
}

func (a *MobileStub) UpdateProject(ctx workflow.Context, input *mobile.UpdateProjectInput) (*mobile.UpdateProjectOutput, error) {
    var output mobile.UpdateProjectOutput
    err := workflow.ExecuteActivity(ctx, a.activities.UpdateProject, input).Get(ctx, &output)
    return &output, err
}