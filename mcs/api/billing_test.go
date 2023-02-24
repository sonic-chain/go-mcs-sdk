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

	logs.GetLogger().Info(mcsClient)

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

	logs.GetLogger().Info(mcsClient)

	lockPaymentInfo, err := mcsClient.GetLockPaymentInfo(18)
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

	logs.GetLogger().Info(mcsClient)

	pageNumber := 1
	pageSize := 10
	billingHistoryParams := BillingHistoryParams{
		PageNumber: &pageNumber,
		PageSize:   &pageSize,
	}
	billing, recCnt, err := mcsClient.GetBillingHistory(billingHistoryParams)
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	logs.GetLogger().Info(billing)
	logs.GetLogger().Info(*recCnt)
}
