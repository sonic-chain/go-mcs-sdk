package common

import (
	"fmt"
	"testing"

	"github.com/filswan/go-swan-lib/logs"
)

func TestGetFilPrice(t *testing.T) {
	price, err := GetFilPrice()
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}
	fmt.Println(price)
}

func TestGetAmount(t *testing.T) {
	amount, err := GetAmount(123, 5)
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}
	fmt.Println(amount)
}
