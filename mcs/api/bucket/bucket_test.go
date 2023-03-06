package bucket

import (
	"go-mcs-sdk/mcs/api/common"

	"github.com/filswan/go-swan-lib/logs"
)

func GetBucketClient() (*BucketClient, error) {
	mcsClient, err := common.GetMcsClient()
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	bucketClient := GetBucketClientFromMcsClient(*mcsClient)

	return &bucketClient, nil
}
