package api

import (
	"go-mcs-sdk/mcs/config"
	"testing"

	"github.com/filswan/go-swan-lib/logs"
)

func TestGetFileCoinPrice(t *testing.T) {
	client, err := GetOnChainClient()
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	filecoinPrice, err := client.GetFileCoinPrice()
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	logs.GetLogger().Info(*filecoinPrice)
}

func TestGetLockPaymentInfo(t *testing.T) {
	client, err := GetOnChainClient()
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	lockPaymentInfo, err := client.GetLockPaymentInfo(2131)
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	logs.GetLogger().Info(*lockPaymentInfo)
}

func TestGetBillingHistory(t *testing.T) {
	client, err := GetOnChainClient()
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
	billings, recCnt, err := client.GetBillingHistory(billingHistoryParams)
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	for _, billing := range billings {
		logs.GetLogger().Info(*billing)
	}

	logs.GetLogger().Info(*recCnt)
}

func TestPayForFile(t *testing.T) {
	client, err := GetOnChainClient()
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	txHash, err := client.PayForFile(1, config.GetConfig().PrivateKey, config.GetConfig().RpcUrl)
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	logs.GetLogger().Info(*txHash)
}
