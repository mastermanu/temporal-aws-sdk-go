package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iotthingsgraph"
	"github.com/aws/aws-sdk-go/service/iotthingsgraph/iotthingsgraphiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type IoTThingsGraphActivities struct {
	client iotthingsgraphiface.IoTThingsGraphAPI
}

func NewIoTThingsGraphActivities(session *session.Session, config ...*aws.Config) *IoTThingsGraphActivities {
	client := iotthingsgraph.New(session, config...)
	return &IoTThingsGraphActivities{client: client}
}

func (a *IoTThingsGraphActivities) AssociateEntityToThing(ctx context.Context, input *iotthingsgraph.AssociateEntityToThingInput) (*iotthingsgraph.AssociateEntityToThingOutput, error) {
	return a.client.AssociateEntityToThingWithContext(ctx, input)
}

func (a *IoTThingsGraphActivities) CreateFlowTemplate(ctx context.Context, input *iotthingsgraph.CreateFlowTemplateInput) (*iotthingsgraph.CreateFlowTemplateOutput, error) {
	return a.client.CreateFlowTemplateWithContext(ctx, input)
}

func (a *IoTThingsGraphActivities) CreateSystemInstance(ctx context.Context, input *iotthingsgraph.CreateSystemInstanceInput) (*iotthingsgraph.CreateSystemInstanceOutput, error) {
	return a.client.CreateSystemInstanceWithContext(ctx, input)
}

func (a *IoTThingsGraphActivities) CreateSystemTemplate(ctx context.Context, input *iotthingsgraph.CreateSystemTemplateInput) (*iotthingsgraph.CreateSystemTemplateOutput, error) {
	return a.client.CreateSystemTemplateWithContext(ctx, input)
}

func (a *IoTThingsGraphActivities) DeleteFlowTemplate(ctx context.Context, input *iotthingsgraph.DeleteFlowTemplateInput) (*iotthingsgraph.DeleteFlowTemplateOutput, error) {
	return a.client.DeleteFlowTemplateWithContext(ctx, input)
}

func (a *IoTThingsGraphActivities) DeleteNamespace(ctx context.Context, input *iotthingsgraph.DeleteNamespaceInput) (*iotthingsgraph.DeleteNamespaceOutput, error) {
	return a.client.DeleteNamespaceWithContext(ctx, input)
}

func (a *IoTThingsGraphActivities) DeleteSystemInstance(ctx context.Context, input *iotthingsgraph.DeleteSystemInstanceInput) (*iotthingsgraph.DeleteSystemInstanceOutput, error) {
	return a.client.DeleteSystemInstanceWithContext(ctx, input)
}

func (a *IoTThingsGraphActivities) DeleteSystemTemplate(ctx context.Context, input *iotthingsgraph.DeleteSystemTemplateInput) (*iotthingsgraph.DeleteSystemTemplateOutput, error) {
	return a.client.DeleteSystemTemplateWithContext(ctx, input)
}

func (a *IoTThingsGraphActivities) DeploySystemInstance(ctx context.Context, input *iotthingsgraph.DeploySystemInstanceInput) (*iotthingsgraph.DeploySystemInstanceOutput, error) {
	return a.client.DeploySystemInstanceWithContext(ctx, input)
}

func (a *IoTThingsGraphActivities) DeprecateFlowTemplate(ctx context.Context, input *iotthingsgraph.DeprecateFlowTemplateInput) (*iotthingsgraph.DeprecateFlowTemplateOutput, error) {
	return a.client.DeprecateFlowTemplateWithContext(ctx, input)
}

func (a *IoTThingsGraphActivities) DeprecateSystemTemplate(ctx context.Context, input *iotthingsgraph.DeprecateSystemTemplateInput) (*iotthingsgraph.DeprecateSystemTemplateOutput, error) {
	return a.client.DeprecateSystemTemplateWithContext(ctx, input)
}

func (a *IoTThingsGraphActivities) DescribeNamespace(ctx context.Context, input *iotthingsgraph.DescribeNamespaceInput) (*iotthingsgraph.DescribeNamespaceOutput, error) {
	return a.client.DescribeNamespaceWithContext(ctx, input)
}

func (a *IoTThingsGraphActivities) DissociateEntityFromThing(ctx context.Context, input *iotthingsgraph.DissociateEntityFromThingInput) (*iotthingsgraph.DissociateEntityFromThingOutput, error) {
	return a.client.DissociateEntityFromThingWithContext(ctx, input)
}

func (a *IoTThingsGraphActivities) GetEntities(ctx context.Context, input *iotthingsgraph.GetEntitiesInput) (*iotthingsgraph.GetEntitiesOutput, error) {
	return a.client.GetEntitiesWithContext(ctx, input)
}

func (a *IoTThingsGraphActivities) GetFlowTemplate(ctx context.Context, input *iotthingsgraph.GetFlowTemplateInput) (*iotthingsgraph.GetFlowTemplateOutput, error) {
	return a.client.GetFlowTemplateWithContext(ctx, input)
}

func (a *IoTThingsGraphActivities) GetFlowTemplateRevisions(ctx context.Context, input *iotthingsgraph.GetFlowTemplateRevisionsInput) (*iotthingsgraph.GetFlowTemplateRevisionsOutput, error) {
	return a.client.GetFlowTemplateRevisionsWithContext(ctx, input)
}

func (a *IoTThingsGraphActivities) GetNamespaceDeletionStatus(ctx context.Context, input *iotthingsgraph.GetNamespaceDeletionStatusInput) (*iotthingsgraph.GetNamespaceDeletionStatusOutput, error) {
	return a.client.GetNamespaceDeletionStatusWithContext(ctx, input)
}

func (a *IoTThingsGraphActivities) GetSystemInstance(ctx context.Context, input *iotthingsgraph.GetSystemInstanceInput) (*iotthingsgraph.GetSystemInstanceOutput, error) {
	return a.client.GetSystemInstanceWithContext(ctx, input)
}

func (a *IoTThingsGraphActivities) GetSystemTemplate(ctx context.Context, input *iotthingsgraph.GetSystemTemplateInput) (*iotthingsgraph.GetSystemTemplateOutput, error) {
	return a.client.GetSystemTemplateWithContext(ctx, input)
}

func (a *IoTThingsGraphActivities) GetSystemTemplateRevisions(ctx context.Context, input *iotthingsgraph.GetSystemTemplateRevisionsInput) (*iotthingsgraph.GetSystemTemplateRevisionsOutput, error) {
	return a.client.GetSystemTemplateRevisionsWithContext(ctx, input)
}

func (a *IoTThingsGraphActivities) GetUploadStatus(ctx context.Context, input *iotthingsgraph.GetUploadStatusInput) (*iotthingsgraph.GetUploadStatusOutput, error) {
	return a.client.GetUploadStatusWithContext(ctx, input)
}

func (a *IoTThingsGraphActivities) ListFlowExecutionMessages(ctx context.Context, input *iotthingsgraph.ListFlowExecutionMessagesInput) (*iotthingsgraph.ListFlowExecutionMessagesOutput, error) {
	return a.client.ListFlowExecutionMessagesWithContext(ctx, input)
}

func (a *IoTThingsGraphActivities) ListTagsForResource(ctx context.Context, input *iotthingsgraph.ListTagsForResourceInput) (*iotthingsgraph.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *IoTThingsGraphActivities) SearchEntities(ctx context.Context, input *iotthingsgraph.SearchEntitiesInput) (*iotthingsgraph.SearchEntitiesOutput, error) {
	return a.client.SearchEntitiesWithContext(ctx, input)
}

func (a *IoTThingsGraphActivities) SearchFlowExecutions(ctx context.Context, input *iotthingsgraph.SearchFlowExecutionsInput) (*iotthingsgraph.SearchFlowExecutionsOutput, error) {
	return a.client.SearchFlowExecutionsWithContext(ctx, input)
}

func (a *IoTThingsGraphActivities) SearchFlowTemplates(ctx context.Context, input *iotthingsgraph.SearchFlowTemplatesInput) (*iotthingsgraph.SearchFlowTemplatesOutput, error) {
	return a.client.SearchFlowTemplatesWithContext(ctx, input)
}

func (a *IoTThingsGraphActivities) SearchSystemInstances(ctx context.Context, input *iotthingsgraph.SearchSystemInstancesInput) (*iotthingsgraph.SearchSystemInstancesOutput, error) {
	return a.client.SearchSystemInstancesWithContext(ctx, input)
}

func (a *IoTThingsGraphActivities) SearchSystemTemplates(ctx context.Context, input *iotthingsgraph.SearchSystemTemplatesInput) (*iotthingsgraph.SearchSystemTemplatesOutput, error) {
	return a.client.SearchSystemTemplatesWithContext(ctx, input)
}

func (a *IoTThingsGraphActivities) SearchThings(ctx context.Context, input *iotthingsgraph.SearchThingsInput) (*iotthingsgraph.SearchThingsOutput, error) {
	return a.client.SearchThingsWithContext(ctx, input)
}

func (a *IoTThingsGraphActivities) TagResource(ctx context.Context, input *iotthingsgraph.TagResourceInput) (*iotthingsgraph.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *IoTThingsGraphActivities) UndeploySystemInstance(ctx context.Context, input *iotthingsgraph.UndeploySystemInstanceInput) (*iotthingsgraph.UndeploySystemInstanceOutput, error) {
	return a.client.UndeploySystemInstanceWithContext(ctx, input)
}

func (a *IoTThingsGraphActivities) UntagResource(ctx context.Context, input *iotthingsgraph.UntagResourceInput) (*iotthingsgraph.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}

func (a *IoTThingsGraphActivities) UpdateFlowTemplate(ctx context.Context, input *iotthingsgraph.UpdateFlowTemplateInput) (*iotthingsgraph.UpdateFlowTemplateOutput, error) {
	return a.client.UpdateFlowTemplateWithContext(ctx, input)
}

func (a *IoTThingsGraphActivities) UpdateSystemTemplate(ctx context.Context, input *iotthingsgraph.UpdateSystemTemplateInput) (*iotthingsgraph.UpdateSystemTemplateOutput, error) {
	return a.client.UpdateSystemTemplateWithContext(ctx, input)
}

func (a *IoTThingsGraphActivities) UploadEntityDefinitions(ctx context.Context, input *iotthingsgraph.UploadEntityDefinitionsInput) (*iotthingsgraph.UploadEntityDefinitionsOutput, error) {
	return a.client.UploadEntityDefinitionsWithContext(ctx, input)
}
