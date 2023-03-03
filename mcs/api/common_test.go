package api

import (
	"testing"

	"github.com/filswan/go-swan-lib/logs"
)

func TestGetSystemParam(t *testing.T) {
	mcsClient, err := GetMcsClient()
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	params, err := mcsClient.GetSystemParam()
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
	amount, err := GetAmount(1, 0.1, 1, 1)
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	logs.GetLogger().Info(amount)
}
