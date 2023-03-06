package bucket

import (
	"go-mcs-sdk/mcs/api/common/auth"
	"go-mcs-sdk/mcs/config"

	"github.com/filswan/go-swan-lib/logs"
)

var onChainClient *BucketClient

func init() {
	if onChainClient != nil {
		return
	}

	apikey := config.GetConfig().Apikey
	accessToken := config.GetConfig().AccessToken
	network := config.GetConfig().Network

	mcsClient, err := auth.LoginByApikey(apikey, accessToken, network)
	if err != nil {
		logs.GetLogger().Fatal(err)
	}

	onChainClient = GetBucketClient(*mcsClient)
}
