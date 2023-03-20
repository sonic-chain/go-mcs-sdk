package api

import (
	"go-mcs-sdk/mcs/config"
	"testing"

	"go-mcs-sdk/mcs/api/common/logs"

	"github.com/stretchr/testify/assert"
)

func TestGetFileCoinPrice(t *testing.T) {
	filecoinPrice, err := onChainClient.GetFileCoinPrice()
	assert.Nil(t, err)
	assert.NotEmpty(t, filecoinPrice)
	assert.GreaterOrEqual(t, *filecoinPrice, 0)

	logs.GetLogger().Info(*filecoinPrice)
}

func TestGetLockPaymentInfo(t *testing.T) {
	lockPaymentInfo, err := onChainClient.GetLockPaymentInfo(2131)
	assert.Nil(t, err)
	assert.NotEmpty(t, lockPaymentInfo)

	logs.GetLogger().Info(*lockPaymentInfo)
}

func TestGetBillingHistory(t *testing.T) {
	pageNumber := 1
	pageSize := 10
	billingHistoryParams := BillingHistoryParams{
		PageNumber: &pageNumber,
		PageSize:   &pageSize,
	}
	billings, recCnt, err := onChainClient.GetBillingHistory(billingHistoryParams)
	if err != nil {
		logs.GetLogger().Fatal(err)
	}
	assert.Nil(t, err)
	assert.NotEmpty(t, billings)
	assert.NotEmpty(t, recCnt)

	for _, billing := range billings {
		logs.GetLogger().Info(*billing)
	}

	logs.GetLogger().Info(*recCnt)
}

func TestPayForFile(t *testing.T) {
	txHash, err := onChainClient.PayForFile(1, config.GetConfig().PrivateKey, config.GetConfig().RpcUrl)
	if err != nil {
		logs.GetLogger().Fatal(err)
	}
	assert.Nil(t, err)
	assert.NotEmpty(t, txHash)

	logs.GetLogger().Info(*txHash)
}
