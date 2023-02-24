package api

import (
	"testing"

	"github.com/filswan/go-swan-lib/logs"
)

func TestGetFilPrice(t *testing.T) {
	price, err := GetFilPrice()
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	logs.GetLogger().Info(price)
}

func TestGetAmount(t *testing.T) {
	amount, err := GetAmount(1, 0.1)
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	logs.GetLogger().Info(amount)
}
