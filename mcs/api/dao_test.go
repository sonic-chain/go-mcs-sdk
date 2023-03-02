package api

import (
	"testing"

	"github.com/filswan/go-swan-lib/logs"
)

func TestGetDeals2PreSign(t *testing.T) {
	mcsClient, err := GetMcsClient()
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	deals, err := mcsClient.GetDeals2PreSign()
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	for _, deal := range deals {
		logs.GetLogger().Info(deal)
	}
}

func TestGetDeals2Sign(t *testing.T) {
	mcsClient, err := GetMcsClient()
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	deals, err := mcsClient.GetDeals2Sign()
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	for _, deal := range deals {
		logs.GetLogger().Info(*deal)
	}
}

func TestGetDeals2SignHash(t *testing.T) {
	mcsClient, err := GetMcsClient()
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	deals, err := mcsClient.GetDeals2SignHash()
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	for _, deal := range deals {
		logs.GetLogger().Info(*deal)
		for _, batchInfo := range deal.BatchInfo {
			logs.GetLogger().Info(*batchInfo)
		}
	}
}
