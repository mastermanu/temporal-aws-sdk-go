package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
	"github.com/aws/aws-sdk-go/service/kms/kmsiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type KMSActivities struct {
	client kmsiface.KMSAPI
}

func NewKMSActivities(session *session.Session, config ...*aws.Config) *KMSActivities {
	client := kms.New(session, config...)
	return &KMSActivities{client: client}
}

func (a *KMSActivities) CancelKeyDeletion(ctx context.Context, input *kms.CancelKeyDeletionInput) (*kms.CancelKeyDeletionOutput, error) {
	return a.client.CancelKeyDeletionWithContext(ctx, input)
}

func (a *KMSActivities) ConnectCustomKeyStore(ctx context.Context, input *kms.ConnectCustomKeyStoreInput) (*kms.ConnectCustomKeyStoreOutput, error) {
	return a.client.ConnectCustomKeyStoreWithContext(ctx, input)
}

func (a *KMSActivities) CreateAlias(ctx context.Context, input *kms.CreateAliasInput) (*kms.CreateAliasOutput, error) {
	return a.client.CreateAliasWithContext(ctx, input)
}

func (a *KMSActivities) CreateCustomKeyStore(ctx context.Context, input *kms.CreateCustomKeyStoreInput) (*kms.CreateCustomKeyStoreOutput, error) {
	return a.client.CreateCustomKeyStoreWithContext(ctx, input)
}

func (a *KMSActivities) CreateGrant(ctx context.Context, input *kms.CreateGrantInput) (*kms.CreateGrantOutput, error) {
	return a.client.CreateGrantWithContext(ctx, input)
}

func (a *KMSActivities) CreateKey(ctx context.Context, input *kms.CreateKeyInput) (*kms.CreateKeyOutput, error) {
	return a.client.CreateKeyWithContext(ctx, input)
}

func (a *KMSActivities) Decrypt(ctx context.Context, input *kms.DecryptInput) (*kms.DecryptOutput, error) {
	return a.client.DecryptWithContext(ctx, input)
}

func (a *KMSActivities) DeleteAlias(ctx context.Context, input *kms.DeleteAliasInput) (*kms.DeleteAliasOutput, error) {
	return a.client.DeleteAliasWithContext(ctx, input)
}

func (a *KMSActivities) DeleteCustomKeyStore(ctx context.Context, input *kms.DeleteCustomKeyStoreInput) (*kms.DeleteCustomKeyStoreOutput, error) {
	return a.client.DeleteCustomKeyStoreWithContext(ctx, input)
}

func (a *KMSActivities) DeleteImportedKeyMaterial(ctx context.Context, input *kms.DeleteImportedKeyMaterialInput) (*kms.DeleteImportedKeyMaterialOutput, error) {
	return a.client.DeleteImportedKeyMaterialWithContext(ctx, input)
}

func (a *KMSActivities) DescribeCustomKeyStores(ctx context.Context, input *kms.DescribeCustomKeyStoresInput) (*kms.DescribeCustomKeyStoresOutput, error) {
	return a.client.DescribeCustomKeyStoresWithContext(ctx, input)
}

func (a *KMSActivities) DescribeKey(ctx context.Context, input *kms.DescribeKeyInput) (*kms.DescribeKeyOutput, error) {
	return a.client.DescribeKeyWithContext(ctx, input)
}

func (a *KMSActivities) DisableKey(ctx context.Context, input *kms.DisableKeyInput) (*kms.DisableKeyOutput, error) {
	return a.client.DisableKeyWithContext(ctx, input)
}

func (a *KMSActivities) DisableKeyRotation(ctx context.Context, input *kms.DisableKeyRotationInput) (*kms.DisableKeyRotationOutput, error) {
	return a.client.DisableKeyRotationWithContext(ctx, input)
}

func (a *KMSActivities) DisconnectCustomKeyStore(ctx context.Context, input *kms.DisconnectCustomKeyStoreInput) (*kms.DisconnectCustomKeyStoreOutput, error) {
	return a.client.DisconnectCustomKeyStoreWithContext(ctx, input)
}

func (a *KMSActivities) EnableKey(ctx context.Context, input *kms.EnableKeyInput) (*kms.EnableKeyOutput, error) {
	return a.client.EnableKeyWithContext(ctx, input)
}

func (a *KMSActivities) EnableKeyRotation(ctx context.Context, input *kms.EnableKeyRotationInput) (*kms.EnableKeyRotationOutput, error) {
	return a.client.EnableKeyRotationWithContext(ctx, input)
}

func (a *KMSActivities) Encrypt(ctx context.Context, input *kms.EncryptInput) (*kms.EncryptOutput, error) {
	return a.client.EncryptWithContext(ctx, input)
}

func (a *KMSActivities) GenerateDataKey(ctx context.Context, input *kms.GenerateDataKeyInput) (*kms.GenerateDataKeyOutput, error) {
	return a.client.GenerateDataKeyWithContext(ctx, input)
}

func (a *KMSActivities) GenerateDataKeyPair(ctx context.Context, input *kms.GenerateDataKeyPairInput) (*kms.GenerateDataKeyPairOutput, error) {
	return a.client.GenerateDataKeyPairWithContext(ctx, input)
}

func (a *KMSActivities) GenerateDataKeyPairWithoutPlaintext(ctx context.Context, input *kms.GenerateDataKeyPairWithoutPlaintextInput) (*kms.GenerateDataKeyPairWithoutPlaintextOutput, error) {
	return a.client.GenerateDataKeyPairWithoutPlaintextWithContext(ctx, input)
}

func (a *KMSActivities) GenerateDataKeyWithoutPlaintext(ctx context.Context, input *kms.GenerateDataKeyWithoutPlaintextInput) (*kms.GenerateDataKeyWithoutPlaintextOutput, error) {
	return a.client.GenerateDataKeyWithoutPlaintextWithContext(ctx, input)
}

func (a *KMSActivities) GenerateRandom(ctx context.Context, input *kms.GenerateRandomInput) (*kms.GenerateRandomOutput, error) {
	return a.client.GenerateRandomWithContext(ctx, input)
}

func (a *KMSActivities) GetKeyPolicy(ctx context.Context, input *kms.GetKeyPolicyInput) (*kms.GetKeyPolicyOutput, error) {
	return a.client.GetKeyPolicyWithContext(ctx, input)
}

func (a *KMSActivities) GetKeyRotationStatus(ctx context.Context, input *kms.GetKeyRotationStatusInput) (*kms.GetKeyRotationStatusOutput, error) {
	return a.client.GetKeyRotationStatusWithContext(ctx, input)
}

func (a *KMSActivities) GetParametersForImport(ctx context.Context, input *kms.GetParametersForImportInput) (*kms.GetParametersForImportOutput, error) {
	return a.client.GetParametersForImportWithContext(ctx, input)
}

func (a *KMSActivities) GetPublicKey(ctx context.Context, input *kms.GetPublicKeyInput) (*kms.GetPublicKeyOutput, error) {
	return a.client.GetPublicKeyWithContext(ctx, input)
}

func (a *KMSActivities) ImportKeyMaterial(ctx context.Context, input *kms.ImportKeyMaterialInput) (*kms.ImportKeyMaterialOutput, error) {
	return a.client.ImportKeyMaterialWithContext(ctx, input)
}

func (a *KMSActivities) ListAliases(ctx context.Context, input *kms.ListAliasesInput) (*kms.ListAliasesOutput, error) {
	return a.client.ListAliasesWithContext(ctx, input)
}

func (a *KMSActivities) ListGrants(ctx context.Context, input *kms.ListGrantsInput) (*kms.ListGrantsResponse, error) {
	return a.client.ListGrantsWithContext(ctx, input)
}

func (a *KMSActivities) ListKeyPolicies(ctx context.Context, input *kms.ListKeyPoliciesInput) (*kms.ListKeyPoliciesOutput, error) {
	return a.client.ListKeyPoliciesWithContext(ctx, input)
}

func (a *KMSActivities) ListKeys(ctx context.Context, input *kms.ListKeysInput) (*kms.ListKeysOutput, error) {
	return a.client.ListKeysWithContext(ctx, input)
}

func (a *KMSActivities) ListResourceTags(ctx context.Context, input *kms.ListResourceTagsInput) (*kms.ListResourceTagsOutput, error) {
	return a.client.ListResourceTagsWithContext(ctx, input)
}

func (a *KMSActivities) ListRetirableGrants(ctx context.Context, input *kms.ListRetirableGrantsInput) (*kms.ListGrantsResponse, error) {
	return a.client.ListRetirableGrantsWithContext(ctx, input)
}

func (a *KMSActivities) PutKeyPolicy(ctx context.Context, input *kms.PutKeyPolicyInput) (*kms.PutKeyPolicyOutput, error) {
	return a.client.PutKeyPolicyWithContext(ctx, input)
}

func (a *KMSActivities) ReEncrypt(ctx context.Context, input *kms.ReEncryptInput) (*kms.ReEncryptOutput, error) {
	return a.client.ReEncryptWithContext(ctx, input)
}

func (a *KMSActivities) RetireGrant(ctx context.Context, input *kms.RetireGrantInput) (*kms.RetireGrantOutput, error) {
	return a.client.RetireGrantWithContext(ctx, input)
}

func (a *KMSActivities) RevokeGrant(ctx context.Context, input *kms.RevokeGrantInput) (*kms.RevokeGrantOutput, error) {
	return a.client.RevokeGrantWithContext(ctx, input)
}

func (a *KMSActivities) ScheduleKeyDeletion(ctx context.Context, input *kms.ScheduleKeyDeletionInput) (*kms.ScheduleKeyDeletionOutput, error) {
	return a.client.ScheduleKeyDeletionWithContext(ctx, input)
}

func (a *KMSActivities) Sign(ctx context.Context, input *kms.SignInput) (*kms.SignOutput, error) {
	return a.client.SignWithContext(ctx, input)
}

func (a *KMSActivities) TagResource(ctx context.Context, input *kms.TagResourceInput) (*kms.TagResourceOutput, error) {
	return a.client.TagResourceWithContext(ctx, input)
}

func (a *KMSActivities) UntagResource(ctx context.Context, input *kms.UntagResourceInput) (*kms.UntagResourceOutput, error) {
	return a.client.UntagResourceWithContext(ctx, input)
}

func (a *KMSActivities) UpdateAlias(ctx context.Context, input *kms.UpdateAliasInput) (*kms.UpdateAliasOutput, error) {
	return a.client.UpdateAliasWithContext(ctx, input)
}

func (a *KMSActivities) UpdateCustomKeyStore(ctx context.Context, input *kms.UpdateCustomKeyStoreInput) (*kms.UpdateCustomKeyStoreOutput, error) {
	return a.client.UpdateCustomKeyStoreWithContext(ctx, input)
}

func (a *KMSActivities) UpdateKeyDescription(ctx context.Context, input *kms.UpdateKeyDescriptionInput) (*kms.UpdateKeyDescriptionOutput, error) {
	return a.client.UpdateKeyDescriptionWithContext(ctx, input)
}

func (a *KMSActivities) Verify(ctx context.Context, input *kms.VerifyInput) (*kms.VerifyOutput, error) {
	return a.client.VerifyWithContext(ctx, input)
}
