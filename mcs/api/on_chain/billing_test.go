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
	assert.GreaterOrEqual(t, *filecoinPrice, float64(0))

	logs.GetLogger().Info(*filecoinPrice)
}

func TestGetLockPaymentInfo(t *testing.T) {
	lockPaymentInfo, err := onChainClient.GetLockPaymentInfo(148234)
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
	assert.Nil(t, err)
	assert.NotEmpty(t, billings)
	assert.NotEmpty(t, recCnt)

	for _, billing := range billings {
		logs.GetLogger().Info(*billing)
	}

	logs.GetLogger().Info(*recCnt)
}

func TestPayForFile(t *testing.T) {
	txHash, err := onChainClient.PayForFile(148234, config.GetConfig().PrivateKey, config.GetConfig().RpcUrl)
	assert.Nil(t, err)
	assert.NotEmpty(t, txHash)

	logs.GetLogger().Info(*txHash)
}
