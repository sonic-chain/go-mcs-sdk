package api

import (
	"go-mcs-sdk/mcs/api/common/auth"
	"testing"

	"github.com/filswan/go-swan-lib/logs"
)

func GetOnChainClient4Test() (*OnChainClient, error) {
	mcsClient, err := auth.GetMcsClient()
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	onChainClient := GetOnChainClient(*mcsClient)

	return &onChainClient, nil
}

func TestLoginByApikey(t *testing.T) {
	client, err := GetOnChainClient4Test()
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	logs.GetLogger().Info(client)
}

func TestGetSystemParam(t *testing.T) {
	client, err := GetOnChainClient4Test()
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	params, err := client.GetSystemParam()
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	logs.GetLogger().Info(params)
}

func TestGetFilPrice(t *testing.T) {
	price, err := GetHistoricalAveragePriceVerified()
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	logs.GetLogger().Info(price)
}

func TestGetAmount(t *testing.T) {
	amount, err := GetAmount(1, 0.1, 1, 2)
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	logs.GetLogger().Info(amount)
}
