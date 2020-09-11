package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lexmodelbuildingservice"
	"github.com/aws/aws-sdk-go/service/lexmodelbuildingservice/lexmodelbuildingserviceiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type LexModelBuildingServiceActivities struct {
	client lexmodelbuildingserviceiface.LexModelBuildingServiceAPI
}

func NewLexModelBuildingServiceActivities(session *session.Session, config ...*aws.Config) *LexModelBuildingServiceActivities {
	client := lexmodelbuildingservice.New(session, config...)
	return &LexModelBuildingServiceActivities{client: client}
}

func (a *LexModelBuildingServiceActivities) CreateBotVersion(ctx context.Context, input *lexmodelbuildingservice.CreateBotVersionInput) (*lexmodelbuildingservice.CreateBotVersionOutput, error) {
	return a.client.CreateBotVersionWithContext(ctx, input)
}

func (a *LexModelBuildingServiceActivities) CreateIntentVersion(ctx context.Context, input *lexmodelbuildingservice.CreateIntentVersionInput) (*lexmodelbuildingservice.CreateIntentVersionOutput, error) {
	return a.client.CreateIntentVersionWithContext(ctx, input)
}

func (a *LexModelBuildingServiceActivities) CreateSlotTypeVersion(ctx context.Context, input *lexmodelbuildingservice.CreateSlotTypeVersionInput) (*lexmodelbuildingservice.CreateSlotTypeVersionOutput, error) {
	return a.client.CreateSlotTypeVersionWithContext(ctx, input)
}

func (a *LexModelBuildingServiceActivities) DeleteBot(ctx context.Context, input *lexmodelbuildingservice.DeleteBotInput) (*lexmodelbuildingservice.DeleteBotOutput, error) {
	return a.client.DeleteBotWithContext(ctx, input)
}

func (a *LexModelBuildingServiceActivities) DeleteBotAlias(ctx context.Context, input *lexmodelbuildingservice.DeleteBotAliasInput) (*lexmodelbuildingservice.DeleteBotAliasOutput, error) {
	return a.client.DeleteBotAliasWithContext(ctx, input)
}

func (a *LexModelBuildingServiceActivities) DeleteBotChannelAssociation(ctx context.Context, input *lexmodelbuildingservice.DeleteBotChannelAssociationInput) (*lexmodelbuildingservice.DeleteBotChannelAssociationOutput, error) {
	return a.client.DeleteBotChannelAssociationWithContext(ctx, input)
}

func (a *LexModelBuildingServiceActivities) DeleteBotVersion(ctx context.Context, input *lexmodelbuildingservice.DeleteBotVersionInput) (*lexmodelbuildingservice.DeleteBotVersionOutput, error) {
	return a.client.DeleteBotVersionWithContext(ctx, input)
}

func (a *LexModelBuildingServiceActivities) DeleteIntent(ctx context.Context, input *lexmodelbuildingservice.DeleteIntentInput) (*lexmodelbuildingservice.DeleteIntentOutput, error) {
	return a.client.DeleteIntentWithContext(ctx, input)
}

func (a *LexModelBuildingServiceActivities) DeleteIntentVersion(ctx context.Context, input *lexmodelbuildingservice.DeleteIntentVersionInput) (*lexmodelbuildingservice.DeleteIntentVersionOutput, error) {
	return a.client.DeleteIntentVersionWithContext(ctx, input)
}

func (a *LexModelBuildingServiceActivities) DeleteSlotType(ctx context.Context, input *lexmodelbuildingservice.DeleteSlotTypeInput) (*lexmodelbuildingservice.DeleteSlotTypeOutput, error) {
	return a.client.DeleteSlotTypeWithContext(ctx, input)
}

func (a *LexModelBuildingServiceActivities) DeleteSlotTypeVersion(ctx context.Context, input *lexmodelbuildingservice.DeleteSlotTypeVersionInput) (*lexmodelbuildingservice.DeleteSlotTypeVersionOutput, error) {
	return a.client.DeleteSlotTypeVersionWithContext(ctx, input)
}

func (a *LexModelBuildingServiceActivities) DeleteUtterances(ctx context.Context, input *lexmodelbuildingservice.DeleteUtterancesInput) (*lexmodelbuildingservice.DeleteUtterancesOutput, error) {
	return a.client.DeleteUtterancesWithContext(ctx, input)
}

func (a *LexModelBuildingServiceActivities) GetBot(ctx context.Context, input *lexmodelbuildingservice.GetBotInput) (*lexmodelbuildingservice.GetBotOutput, error) {
	return a.client.GetBotWithContext(ctx, input)
}

func (a *LexModelBuildingServiceActivities) GetBotAlias(ctx context.Context, input *lexmodelbuildingservice.GetBotAliasInput) (*lexmodelbuildingservice.GetBotAliasOutput, error) {
	return a.client.GetBotAliasWithContext(ctx, input)
}

func (a *LexModelBuildingServiceActivities) GetBotAliases(ctx context.Context, input *lexmodelbuildingservice.GetBotAliasesInput) (*lexmodelbuildingservice.GetBotAliasesOutput, error) {
	return a.client.GetBotAliasesWithContext(ctx, input)
}

func (a *LexModelBuildingServiceActivities) GetBotChannelAssociation(ctx context.Context, input *lexmodelbuildingservice.GetBotChannelAssociationInput) (*lexmodelbuildingservice.GetBotChannelAssociationOutput, error) {
	return a.client.GetBotChannelAssociationWithContext(ctx, input)
}

func (a *LexModelBuildingServiceActivities) GetBotChannelAssociations(ctx context.Context, input *lexmodelbuildingservice.GetBotChannelAssociationsInput) (*lexmodelbuildingservice.GetBotChannelAssociationsOutput, error) {
	return a.client.GetBotChannelAssociationsWithContext(ctx, input)
}

func (a *LexModelBuildingServiceActivities) GetBotVersions(ctx context.Context, input *lexmodelbuildingservice.GetBotVersionsInput) (*lexmodelbuildingservice.GetBotVersionsOutput, error) {
	return a.client.GetBotVersionsWithContext(ctx, input)
}

func (a *LexModelBuildingServiceActivities) GetBots(ctx context.Context, input *lexmodelbuildingservice.GetBotsInput) (*lexmodelbuildingservice.GetBotsOutput, error) {
	return a.client.GetBotsWithContext(ctx, input)
}

func (a *LexModelBuildingServiceActivities) GetBuiltinIntent(ctx context.Context, input *lexmodelbuildingservice.GetBuiltinIntentInput) (*lexmodelbuildingservice.GetBuiltinIntentOutput, error) {
	return a.client.GetBuiltinIntentWithContext(ctx, input)
}

func (a *LexModelBuildingServiceActivities) GetBuiltinIntents(ctx context.Context, input *lexmodelbuildingservice.GetBuiltinIntentsInput) (*lexmodelbuildingservice.GetBuiltinIntentsOutput, error) {
	return a.client.GetBuiltinIntentsWithContext(ctx, input)
}

func (a *LexModelBuildingServiceActivities) GetBuiltinSlotTypes(ctx context.Context, input *lexmodelbuildingservice.GetBuiltinSlotTypesInput) (*lexmodelbuildingservice.GetBuiltinSlotTypesOutput, error) {
	return a.client.GetBuiltinSlotTypesWithContext(ctx, input)
}

func (a *LexModelBuildingServiceActivities) GetExport(ctx context.Context, input *lexmodelbuildingservice.GetExportInput) (*lexmodelbuildingservice.GetExportOutput, error) {
	return a.client.GetExportWithContext(ctx, input)
}

func (a *LexModelBuildingServiceActivities) GetImport(ctx context.Context, input *lexmodelbuildingservice.GetImportInput) (*lexmodelbuildingservice.GetImportOutput, error) {
	return a.client.GetImportWithContext(ctx, input)
}

func (a *LexModelBuildingServiceActivities) GetIntent(ctx context.Context, input *lexmodelbuildingservice.GetIntentInput) (*lexmodelbuildingservice.GetIntentOutput, error) {
	return a.client.GetIntentWithContext(ctx, input)
}

func (a *LexModelBuildingServiceActivities) GetIntentVersions(ctx context.Context, input *lexmodelbuildingservice.GetIntentVersionsInput) (*lexmodelbuildingservice.GetIntentVersionsOutput, error) {
	return a.client.GetIntentVersionsWithContext(ctx, input)
}

func (a *LexModelBuildingServiceActivities) GetIntents(ctx context.Context, input *lexmodelbuildingservice.GetIntentsInput) (*lexmodelbuildingservice.GetIntentsOutput, error) {
	return a.client.GetIntentsWithContext(ctx, input)
}

func (a *LexModelBuildingServiceActivities) GetSlotType(ctx context.Context, input *lexmodelbuildingservice.GetSlotTypeInput) (*lexmodelbuildingservice.GetSlotTypeOutput, error) {
	return a.client.GetSlotTypeWithContext(ctx, input)
}

func (a *LexModelBuildingServiceActivities) GetSlotTypeVersions(ctx context.Context, input *lexmodelbuildingservice.GetSlotTypeVersionsInput) (*lexmodelbuildingservice.GetSlotTypeVersionsOutput, error) {
	return a.client.GetSlotTypeVersionsWithContext(ctx, input)
}

func (a *LexModelBuildingServiceActivities) GetSlotTypes(ctx context.Context, input *lexmodelbuildingservice.GetSlotTypesInput) (*lexmodelbuildingservice.GetSlotTypesOutput, error) {
	return a.client.GetSlotTypesWithContext(ctx, input)
}

func (a *LexModelBuildingServiceActivities) GetUtterancesView(ctx context.Context, input *lexmodelbuildingservice.GetUtterancesViewInput) (*lexmodelbuildingservice.GetUtterancesViewOutput, error) {
	return a.client.GetUtterancesViewWithContext(ctx, input)
}

func (a *LexModelBuildingServiceActivities) ListTagsForResource(ctx context.Context, input *lexmodelbuildingservice.ListTagsForResourceInput) (*lexmodelbuildingservice.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *LexModelBuildingServiceActivities) PutBot(ctx context.Context, input *lexmodelbuildingservice.PutBotInput) (*lexmodelbuildingservice.PutBotOutput, error) {
	return a.client.PutBotWithContext(ctx, input)
}

func (a *LexModelBuildingServiceActivities) PutBotAlias(ctx context.Context, input *lexmodelbuildingservice.PutBotAliasInput) (*lexmodelbuildingservice.PutBotAliasOutput, error) {
	return a.client.PutBotAliasWithContext(ctx, input)
}

func (a *LexModelBuildingServiceActivities) PutIntent(ctx context.Context, input *lexmodelbuildingservice.PutIntentInput) (*lexmodelbuildingservice.PutIntentOutput, error) {
	return a.client.PutIntentWithContext(ctx, input)
}

func (a *LexModelBuildingServiceActivities) PutSlotType(ctx context.Context, input *lexmodelbuildingservice.PutSlotTypeInput) (*lexmodelbuildingservice.PutSlotTypeOutput, error) {
	return a.client.PutSlotTypeWithContext(ctx, input)
}

func (a *LexModelBuildingServiceActivities) StartImport(ctx context.Context, input *lexmodelbuildingservice.StartImportInput) (*lexmodelbuildingservice.StartImportOutput, error) {
	return a.client.StartImportWithContext(ctx, input)
}

func (a *LexModelBuildingServiceActivities) TagResource(ctx context.Context, input *lexmodelbuildingservice.TagResourceInput) (*lexmodelbuildingservice.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *LexModelBuildingServiceActivities) UntagResource(ctx context.Context, input *lexmodelbuildingservice.UntagResourceInput) (*lexmodelbuildingservice.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}
