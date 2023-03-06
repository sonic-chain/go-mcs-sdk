package auth

import (
	"go-mcs-sdk/mcs/config"
	"testing"

	"github.com/filswan/go-swan-lib/logs"
)

func TestLoginByApikey(t *testing.T) {
	apikey := config.GetConfig().Apikey
	accessToken := config.GetConfig().AccessToken
	network := config.GetConfig().Network

	mcsClient, err := LoginByApikey(apikey, accessToken, network)
	if err != nil {
		logs.GetLogger().Fatal(err)
	}

	logs.GetLogger().Info(mcsClient)
}
