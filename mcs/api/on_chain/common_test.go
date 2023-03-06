package api

import (
	"go-mcs-sdk/mcs/api/common/auth"
	"go-mcs-sdk/mcs/config"
	"testing"

	"github.com/filswan/go-swan-lib/logs"
)

var onChainClient *OnChainClient

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

	onChainClient = GetOnChainClient(*mcsClient)
}

func TestGetSystemParam(t *testing.T) {
	params, err := onChainClient.GetSystemParam()
	if err != nil {
		logs.GetLogger().Fatal(err)
	}

	logs.GetLogger().Info(params)
}

func TestGetFilPrice(t *testing.T) {
	price, err := GetHistoricalAveragePriceVerified()
	if err != nil {
		logs.GetLogger().Fatal(err)
	}

	logs.GetLogger().Info(price)
}

func TestGetAmount(t *testing.T) {
	amount, err := GetAmount(1, 0.1, 1, 2)
	if err != nil {
		logs.GetLogger().Fatal(err)
	}

	logs.GetLogger().Info(amount)
}
