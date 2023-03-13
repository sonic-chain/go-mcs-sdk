package bucket

import (
	"go-mcs-sdk/mcs/api/common/auth"
	"go-mcs-sdk/mcs/config"
	"testing"

	"github.com/filswan/go-swan-lib/logs"
)

var buketClient *BucketClient

func init() {
	if buketClient != nil {
		return
	}

	apikey := config.GetConfig().Apikey
	accessToken := config.GetConfig().AccessToken
	network := config.GetConfig().Network

	mcsClient, err := auth.LoginByApikey(apikey, accessToken, network)
	if err != nil {
		logs.GetLogger().Fatal(err)
	}

	buketClient = GetBucketClient(*mcsClient)
}

func TestGetGateway(t *testing.T) {
	subDomains, err := buketClient.GetGateway()
	if err != nil {
		logs.GetLogger().Fatal(err)
	}

	for _, subDomain := range subDomains {
		logs.GetLogger().Info(subDomain)
	}
}
