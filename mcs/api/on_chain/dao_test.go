package api

import (
	"testing"

	"go-mcs-sdk/mcs/api/common/logs"

	"github.com/stretchr/testify/assert"
)

func TestGetDeals2PreSign(t *testing.T) {
	deals, err := onChainClient.GetDeals2PreSign()
	assert.Nil(t, err)
	assert.NotNil(t, deals)

	for _, deal := range deals {
		logs.GetLogger().Info(deal)
	}
}

func TestGetDeals2Sign(t *testing.T) {
	deals, err := onChainClient.GetDeals2Sign()
	assert.Nil(t, err)
	assert.NotNil(t, deals)

	for _, deal := range deals {
		logs.GetLogger().Info(*deal)
	}
}

func TestGetDeals2SignHash(t *testing.T) {
	deals, err := onChainClient.GetDeals2SignHash()
	assert.Nil(t, err)
	assert.NotNil(t, deals)

	for _, deal := range deals {
		logs.GetLogger().Info(*deal)
		for _, batchInfo := range deal.BatchInfo {
			logs.GetLogger().Info(*batchInfo)
		}
	}
}
