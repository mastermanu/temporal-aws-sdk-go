package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/appsync"
	"github.com/aws/aws-sdk-go/service/appsync/appsynciface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type AppSyncActivities struct {
	client appsynciface.AppSyncAPI
}

func NewAppSyncActivities(session *session.Session, config ...*aws.Config) *AppSyncActivities {
	client := appsync.New(session, config...)
	return &AppSyncActivities{client: client}
}

func (a *AppSyncActivities) CreateApiCache(ctx context.Context, input *appsync.CreateApiCacheInput) (*appsync.CreateApiCacheOutput, error) {
	return a.client.CreateApiCacheWithContext(ctx, input)
}

func (a *AppSyncActivities) CreateApiKey(ctx context.Context, input *appsync.CreateApiKeyInput) (*appsync.CreateApiKeyOutput, error) {
	return a.client.CreateApiKeyWithContext(ctx, input)
}

func (a *AppSyncActivities) CreateDataSource(ctx context.Context, input *appsync.CreateDataSourceInput) (*appsync.CreateDataSourceOutput, error) {
	return a.client.CreateDataSourceWithContext(ctx, input)
}

func (a *AppSyncActivities) CreateFunction(ctx context.Context, input *appsync.CreateFunctionInput) (*appsync.CreateFunctionOutput, error) {
	return a.client.CreateFunctionWithContext(ctx, input)
}

func (a *AppSyncActivities) CreateGraphqlApi(ctx context.Context, input *appsync.CreateGraphqlApiInput) (*appsync.CreateGraphqlApiOutput, error) {
	return a.client.CreateGraphqlApiWithContext(ctx, input)
}

func (a *AppSyncActivities) CreateResolver(ctx context.Context, input *appsync.CreateResolverInput) (*appsync.CreateResolverOutput, error) {
	return a.client.CreateResolverWithContext(ctx, input)
}

func (a *AppSyncActivities) CreateType(ctx context.Context, input *appsync.CreateTypeInput) (*appsync.CreateTypeOutput, error) {
	return a.client.CreateTypeWithContext(ctx, input)
}

func (a *AppSyncActivities) DeleteApiCache(ctx context.Context, input *appsync.DeleteApiCacheInput) (*appsync.DeleteApiCacheOutput, error) {
	return a.client.DeleteApiCacheWithContext(ctx, input)
}

func (a *AppSyncActivities) DeleteApiKey(ctx context.Context, input *appsync.DeleteApiKeyInput) (*appsync.DeleteApiKeyOutput, error) {
	return a.client.DeleteApiKeyWithContext(ctx, input)
}

func (a *AppSyncActivities) DeleteDataSource(ctx context.Context, input *appsync.DeleteDataSourceInput) (*appsync.DeleteDataSourceOutput, error) {
	return a.client.DeleteDataSourceWithContext(ctx, input)
}

func (a *AppSyncActivities) DeleteFunction(ctx context.Context, input *appsync.DeleteFunctionInput) (*appsync.DeleteFunctionOutput, error) {
	return a.client.DeleteFunctionWithContext(ctx, input)
}

func (a *AppSyncActivities) DeleteGraphqlApi(ctx context.Context, input *appsync.DeleteGraphqlApiInput) (*appsync.DeleteGraphqlApiOutput, error) {
	return a.client.DeleteGraphqlApiWithContext(ctx, input)
}

func (a *AppSyncActivities) DeleteResolver(ctx context.Context, input *appsync.DeleteResolverInput) (*appsync.DeleteResolverOutput, error) {
	return a.client.DeleteResolverWithContext(ctx, input)
}

func (a *AppSyncActivities) DeleteType(ctx context.Context, input *appsync.DeleteTypeInput) (*appsync.DeleteTypeOutput, error) {
	return a.client.DeleteTypeWithContext(ctx, input)
}

func (a *AppSyncActivities) FlushApiCache(ctx context.Context, input *appsync.FlushApiCacheInput) (*appsync.FlushApiCacheOutput, error) {
	return a.client.FlushApiCacheWithContext(ctx, input)
}

func (a *AppSyncActivities) GetApiCache(ctx context.Context, input *appsync.GetApiCacheInput) (*appsync.GetApiCacheOutput, error) {
	return a.client.GetApiCacheWithContext(ctx, input)
}

func (a *AppSyncActivities) GetDataSource(ctx context.Context, input *appsync.GetDataSourceInput) (*appsync.GetDataSourceOutput, error) {
	return a.client.GetDataSourceWithContext(ctx, input)
}

func (a *AppSyncActivities) GetFunction(ctx context.Context, input *appsync.GetFunctionInput) (*appsync.GetFunctionOutput, error) {
	return a.client.GetFunctionWithContext(ctx, input)
}

func (a *AppSyncActivities) GetGraphqlApi(ctx context.Context, input *appsync.GetGraphqlApiInput) (*appsync.GetGraphqlApiOutput, error) {
	return a.client.GetGraphqlApiWithContext(ctx, input)
}

func (a *AppSyncActivities) GetIntrospectionSchema(ctx context.Context, input *appsync.GetIntrospectionSchemaInput) (*appsync.GetIntrospectionSchemaOutput, error) {
	return a.client.GetIntrospectionSchemaWithContext(ctx, input)
}

func (a *AppSyncActivities) GetResolver(ctx context.Context, input *appsync.GetResolverInput) (*appsync.GetResolverOutput, error) {
	return a.client.GetResolverWithContext(ctx, input)
}

func (a *AppSyncActivities) GetSchemaCreationStatus(ctx context.Context, input *appsync.GetSchemaCreationStatusInput) (*appsync.GetSchemaCreationStatusOutput, error) {
	return a.client.GetSchemaCreationStatusWithContext(ctx, input)
}

func (a *AppSyncActivities) GetType(ctx context.Context, input *appsync.GetTypeInput) (*appsync.GetTypeOutput, error) {
	return a.client.GetTypeWithContext(ctx, input)
}

func (a *AppSyncActivities) ListApiKeys(ctx context.Context, input *appsync.ListApiKeysInput) (*appsync.ListApiKeysOutput, error) {
	return a.client.ListApiKeysWithContext(ctx, input)
}

func (a *AppSyncActivities) ListDataSources(ctx context.Context, input *appsync.ListDataSourcesInput) (*appsync.ListDataSourcesOutput, error) {
	return a.client.ListDataSourcesWithContext(ctx, input)
}

func (a *AppSyncActivities) ListFunctions(ctx context.Context, input *appsync.ListFunctionsInput) (*appsync.ListFunctionsOutput, error) {
	return a.client.ListFunctionsWithContext(ctx, input)
}

func (a *AppSyncActivities) ListGraphqlApis(ctx context.Context, input *appsync.ListGraphqlApisInput) (*appsync.ListGraphqlApisOutput, error) {
	return a.client.ListGraphqlApisWithContext(ctx, input)
}

func (a *AppSyncActivities) ListResolvers(ctx context.Context, input *appsync.ListResolversInput) (*appsync.ListResolversOutput, error) {
	return a.client.ListResolversWithContext(ctx, input)
}

func (a *AppSyncActivities) ListResolversByFunction(ctx context.Context, input *appsync.ListResolversByFunctionInput) (*appsync.ListResolversByFunctionOutput, error) {
	return a.client.ListResolversByFunctionWithContext(ctx, input)
}

func (a *AppSyncActivities) ListTagsForResource(ctx context.Context, input *appsync.ListTagsForResourceInput) (*appsync.ListTagsForResourceOutput, error) {
	return a.client.ListTagsForResourceWithContext(ctx, input)
}

func (a *AppSyncActivities) ListTypes(ctx context.Context, input *appsync.ListTypesInput) (*appsync.ListTypesOutput, error) {
	return a.client.ListTypesWithContext(ctx, input)
}

func (a *AppSyncActivities) StartSchemaCreation(ctx context.Context, input *appsync.StartSchemaCreationInput) (*appsync.StartSchemaCreationOutput, error) {
	return a.client.StartSchemaCreationWithContext(ctx, input)
}

func (a *AppSyncActivities) TagResource(ctx context.Context, input *appsync.TagResourceInput) (*appsync.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *AppSyncActivities) UntagResource(ctx context.Context, input *appsync.UntagResourceInput) (*appsync.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}

func (a *AppSyncActivities) UpdateApiCache(ctx context.Context, input *appsync.UpdateApiCacheInput) (*appsync.UpdateApiCacheOutput, error) {
	return a.client.UpdateApiCacheWithContext(ctx, input)
}

func (a *AppSyncActivities) UpdateApiKey(ctx context.Context, input *appsync.UpdateApiKeyInput) (*appsync.UpdateApiKeyOutput, error) {
	return a.client.UpdateApiKeyWithContext(ctx, input)
}

func (a *AppSyncActivities) UpdateDataSource(ctx context.Context, input *appsync.UpdateDataSourceInput) (*appsync.UpdateDataSourceOutput, error) {
	return a.client.UpdateDataSourceWithContext(ctx, input)
}

func (a *AppSyncActivities) UpdateFunction(ctx context.Context, input *appsync.UpdateFunctionInput) (*appsync.UpdateFunctionOutput, error) {
	return a.client.UpdateFunctionWithContext(ctx, input)
}

func (a *AppSyncActivities) UpdateGraphqlApi(ctx context.Context, input *appsync.UpdateGraphqlApiInput) (*appsync.UpdateGraphqlApiOutput, error) {
	return a.client.UpdateGraphqlApiWithContext(ctx, input)
}

func (a *AppSyncActivities) UpdateResolver(ctx context.Context, input *appsync.UpdateResolverInput) (*appsync.UpdateResolverOutput, error) {
	return a.client.UpdateResolverWithContext(ctx, input)
}

func (a *AppSyncActivities) UpdateType(ctx context.Context, input *appsync.UpdateTypeInput) (*appsync.UpdateTypeOutput, error) {
	return a.client.UpdateTypeWithContext(ctx, input)
}
