package api

import (
	"go-mcs-sdk/mcs/config"
	"testing"

	"github.com/filswan/go-swan-lib/logs"
)

func TestGetMCSSystemParam(t *testing.T) {
	apikey := config.GetConfig().Apikey
	accessToken := config.GetConfig().AccessToken
	network := config.GetConfig().Network

	mcsClient, err := LoginByApikey(apikey, accessToken, network)
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	params, err := mcsClient.GetMCSSystemParam()
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	logs.GetLogger().Info(params)
}

func TestGetFilPrice(t *testing.T) {
	price, err := GetFilPrice()
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	logs.GetLogger().Info(price)
}

func TestGetAmount(t *testing.T) {
	amount, err := GetAmount(1, 0.1)
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	logs.GetLogger().Info(amount)
}
