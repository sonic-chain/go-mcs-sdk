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

	deals2PreSign, err := mcsClient.GetDeals2PreSign()
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	for _, deal2Sign := range deals2PreSign {
		logs.GetLogger().Info(deal2Sign)
	}
}

func TestGetDeals2Sign(t *testing.T) {
	mcsClient, err := GetMcsClient()
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	deals2PreSign, err := mcsClient.GetDeals2Sign()
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	for _, deal2Sign := range deals2PreSign {
		logs.GetLogger().Info(*deal2Sign)
	}
}

func TestGetDeals2SignHash(t *testing.T) {
	mcsClient, err := GetMcsClient()
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	deals2PreSign, err := mcsClient.GetDeals2SignHash()
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	for _, deal2Sign := range deals2PreSign {
		logs.GetLogger().Info(*deal2Sign)
		for _, batchInfo := range deal2Sign.BatchInfo {
			logs.GetLogger().Info(*batchInfo)
		}
	}
}
