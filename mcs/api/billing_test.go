package api

import (
	"go-mcs-sdk/mcs/config"
	"testing"

	"github.com/filswan/go-swan-lib/logs"
)

func TestGetFileCoinPrice(t *testing.T) {
	apikey := config.GetConfig().Apikey
	accessToken := config.GetConfig().AccessToken
	network := config.GetConfig().Network

	mcsClient, err := LoginByApikey(apikey, accessToken, network)
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	logs.GetLogger().Info(mcsClient)

	filecoinPrice, err := mcsClient.GetFileCoinPrice()
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	logs.GetLogger().Info(*filecoinPrice)
}
