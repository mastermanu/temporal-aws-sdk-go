
package awsactivities

import (
	"github.com/aws/aws-sdk-go/service/pi"
	"github.com/aws/aws-sdk-go/service/pi/piiface"
)

type PIActivities struct {
	client piiface.PIAPI
}

func NewPIActivities(client piiface.PIAPI) *PIActivities {
    return &PIActivities{client: client}
}

func (a *PIActivities) DescribeDimensionKeys(input *pi.DescribeDimensionKeysInput) (*pi.DescribeDimensionKeysOutput, error) {
    return a.client.DescribeDimensionKeys(input)
}

func (a *PIActivities) GetResourceMetrics(input *pi.GetResourceMetricsInput) (*pi.GetResourceMetricsOutput, error) {
    return a.client.GetResourceMetrics(input)
}