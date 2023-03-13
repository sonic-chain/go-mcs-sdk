package bucket

import (
	"go-mcs-sdk/mcs/api/common/constants"
	"go-mcs-sdk/mcs/api/common/web"

	"github.com/filswan/go-swan-lib/logs"
	libutils "github.com/filswan/go-swan-lib/utils"
)

func (bucketClient *BucketClient) GetGateway() ([]string, error) {
	apiUrl := libutils.UrlJoin(bucketClient.BaseUrl, constants.API_URL_BUCKET_GATEWAY_GET_GATEWAY)

	var subDomains []string

	err := web.HttpGet(apiUrl, bucketClient.JwtToken, nil, &subDomains)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return subDomains, nil
}
