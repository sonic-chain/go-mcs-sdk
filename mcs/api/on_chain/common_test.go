package api

import (
	"testing"

	"github.com/filswan/go-mcs-sdk/mcs/api/user"

	"github.com/filswan/go-mcs-sdk/mcs/api/common/constants"
	"github.com/filswan/go-mcs-sdk/mcs/api/common/logs"

	"github.com/stretchr/testify/assert"
)

var onChainClient *OnChainClient
var network = constants.PAYMENT_CHAIN_NAME_POLYGON_MUMBAI

func init() {
	if onChainClient != nil {
		return
	}

	apikey := ""
	accessToken := ""

	mcsClient, err := user.LoginByApikey(apikey, accessToken, network)
	if err != nil {
		logs.GetLogger().Fatal(err)
	}

	onChainClient = GetOnChainClient(*mcsClient)
}

func TestGetSystemParam(t *testing.T) {
	params, err := onChainClient.GetSystemParam()
	assert.Nil(t, err)
	assert.NotEmpty(t, params)

	logs.GetLogger().Info(params)
}

func TestGetFilPrice(t *testing.T) {
	price, err := GetHistoricalAveragePriceVerified()
	assert.Nil(t, err)
	assert.NotNil(t, price)
	assert.GreaterOrEqual(t, price, float64(0))

	logs.GetLogger().Info(price)
}

func TestGetAmount(t *testing.T) {
	amount, err := GetAmount(1, 0.1, 1, 2)
	assert.Nil(t, err)
	assert.NotEmpty(t, amount)
	assert.GreaterOrEqual(t, amount, int64(0))

	logs.GetLogger().Info(amount)
}
