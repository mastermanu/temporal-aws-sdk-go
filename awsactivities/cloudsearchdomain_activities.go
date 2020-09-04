
package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/cloudsearchdomain"
	"github.com/aws/aws-sdk-go/service/cloudsearchdomain/cloudsearchdomainiface"
)

type CloudSearchDomainActivities struct {
	client cloudsearchdomainiface.CloudSearchDomainAPI
}

func NewCloudSearchDomainActivities(client cloudsearchdomainiface.CloudSearchDomainAPI) *CloudSearchDomainActivities {
    return &CloudSearchDomainActivities{client: client}
}

func (a *CloudSearchDomainActivities) Search(input *cloudsearchdomain.SearchInput) (*cloudsearchdomain.SearchOutput, error) {
    return a.client.Search(input)
}

func (a *CloudSearchDomainActivities) Suggest(input *cloudsearchdomain.SuggestInput) (*cloudsearchdomain.SuggestOutput, error) {
    return a.client.Suggest(input)
}

func (a *CloudSearchDomainActivities) UploadDocuments(input *cloudsearchdomain.UploadDocumentsInput) (*cloudsearchdomain.UploadDocumentsOutput, error) {
    return a.client.UploadDocuments(input)
}