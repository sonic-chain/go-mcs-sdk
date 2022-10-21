package mcs

import (
	"fmt"
	"testing"
)

func TestMcsParams_GetMcsApi(t *testing.T) {
	p := NewMcsParams("polygon.mumbai")
	fmt.Println(p.McsApi)
	fmt.Println(p.USDCAddress)
}
