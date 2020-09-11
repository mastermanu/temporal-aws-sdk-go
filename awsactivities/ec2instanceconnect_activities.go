package awsactivities

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2instanceconnect"
	"github.com/aws/aws-sdk-go/service/ec2instanceconnect/ec2instanceconnectiface"
	"temporal.io/aws-sdk/internal"
)

// ensure that imports are valid even if not used by the generated code
var _ = internal.SetClientToken

type _ request.Option

type EC2InstanceConnectActivities struct {
	client ec2instanceconnectiface.EC2InstanceConnectAPI
}

func NewEC2InstanceConnectActivities(session *session.Session, config ...*aws.Config) *EC2InstanceConnectActivities {
	client := ec2instanceconnect.New(session, config...)
	return &EC2InstanceConnectActivities{client: client}
}

func (a *EC2InstanceConnectActivities) SendSSHPublicKey(ctx context.Context, input *ec2instanceconnect.SendSSHPublicKeyInput) (*ec2instanceconnect.SendSSHPublicKeyOutput, error) {
	return a.client.SendSSHPublicKeyWithContext(ctx, input)
}
