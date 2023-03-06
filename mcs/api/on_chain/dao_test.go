package api

import (
	"testing"

	"github.com/filswan/go-swan-lib/logs"
)

func TestGetDeals2PreSign(t *testing.T) {
	deals, err := onChainClient.GetDeals2PreSign()
	if err != nil {
		logs.GetLogger().Fatal(err)
	}

	for _, deal := range deals {
		logs.GetLogger().Info(deal)
	}
}

func TestGetDeals2Sign(t *testing.T) {
	deals, err := onChainClient.GetDeals2Sign()
	if err != nil {
		logs.GetLogger().Fatal(err)
	}

	for _, deal := range deals {
		logs.GetLogger().Info(*deal)
	}
}

func TestGetDeals2SignHash(t *testing.T) {
	deals, err := onChainClient.GetDeals2SignHash()
	if err != nil {
		logs.GetLogger().Fatal(err)
	}

	for _, deal := range deals {
		logs.GetLogger().Info(*deal)
		for _, batchInfo := range deal.BatchInfo {
			logs.GetLogger().Info(*batchInfo)
		}
	}
}
