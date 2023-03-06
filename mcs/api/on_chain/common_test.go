package api

import (
	"go-mcs-sdk/mcs/api/common"
	"testing"

	"github.com/filswan/go-swan-lib/logs"
)

func GetOnChainClient() (*OnChainClient, error) {
	mcsClient, err := common.GetMcsClient()
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	onChainClient := GetOnChainClientFromMcsClient(*mcsClient)

	return &onChainClient, nil
}

func TestLoginByApikey(t *testing.T) {
	client, err := GetOnChainClient()
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	logs.GetLogger().Info(client)
}

func TestGetSystemParam(t *testing.T) {
	client, err := GetOnChainClient()
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
