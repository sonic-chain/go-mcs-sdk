package common

import (
	"fmt"
	"testing"
)

func TestGetFilPrice(t *testing.T) {
	price, err := GetFilPrice()
	if err != nil {
		return
	}
	fmt.Println(price)
}

func TestGetAmount(t *testing.T) {
	amount := GetAmount(123, 5)
	fmt.Println(amount)
}
