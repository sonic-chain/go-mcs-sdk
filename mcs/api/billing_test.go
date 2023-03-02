package api

import (
	"testing"

	"github.com/filswan/go-swan-lib/logs"
)

func TestGetFileCoinPrice(t *testing.T) {
	mcsClient, err := GetMcsClient()
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	filecoinPrice, err := mcsClient.GetFileCoinPrice()
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	logs.GetLogger().Info(*filecoinPrice)
}

func TestGetLockPaymentInfo(t *testing.T) {
	mcsClient, err := GetMcsClient()
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	lockPaymentInfo, err := mcsClient.GetLockPaymentInfo(2131)
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	logs.GetLogger().Info(*lockPaymentInfo)
}

func TestGetBillingHistory(t *testing.T) {
	mcsClient, err := GetMcsClient()
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	pageNumber := 1
	pageSize := 10
	billingHistoryParams := BillingHistoryParams{
		PageNumber: &pageNumber,
		PageSize:   &pageSize,
	}
	billings, recCnt, err := mcsClient.GetBillingHistory(billingHistoryParams)
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	for _, billing := range billings {
		logs.GetLogger().Info(*billing)
	}

	logs.GetLogger().Info(*recCnt)
}
