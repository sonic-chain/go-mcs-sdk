package bucket

import (
	"go-mcs-sdk/mcs/api"
	"go-mcs-sdk/mcs/api/common/constants"
	"go-mcs-sdk/mcs/api/common/utils"

	"github.com/filswan/go-swan-lib/logs"
	libutils "github.com/filswan/go-swan-lib/utils"
)

type BucketClient struct {
	api.McsClient
}

func (bucketClient *BucketClient) CreateBucket(bucketName string) error {
	apiUrl := libutils.UrlJoin(bucketClient.BaseUrl, constants.API_URL_BUCKET_CREATE_BUCKET)

	var bucket struct {
		BucketName string `json:"bucket_name"`
	}
	bucket.BucketName = bucketName

	err := utils.HttpPost(apiUrl, bucketClient.JwtToken, &bucket, nil)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	return nil
}
