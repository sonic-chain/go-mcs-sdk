package user

import (
	"go-mcs-sdk/mcs/config"
	"testing"

	"github.com/filswan/go-swan-lib/logs"
	"github.com/tj/assert"
)

func TestLoginByApikey(t *testing.T) {
	apikey := config.GetConfig().Apikey
	accessToken := config.GetConfig().AccessToken
	network := config.GetConfig().Network

	mcsClient, err := LoginByApikey(apikey, accessToken, network)
	assert.Nil(t, err, err.Error())

	logs.GetLogger().Info(mcsClient)
}

func TestRegister(t *testing.T) {
	network := config.GetConfig().Network

	nonce, err := Register("0xbE14Eb1ffcA54861D3081560110a45F4A1A9e9c5", network)
	if err != nil {
		logs.GetLogger().Fatal(err)
	}

	logs.GetLogger().Info(*nonce)
}

func TestLoginByPublicKeySignature(t *testing.T) {
	network := config.GetConfig().Network

	mcsClient, err := LoginByPublicKeySignature("548241659096254470622564611792310978072", "0xbE14Eb1ffcA54861D3081560110a45F4A1A9e9c5",
		"0x734e7b4b3c0208c325652f16629c0bea1c6b5ad74068f24dcc9f18c318fba876384bdb5768edc12f2a66ed114c9272b9c9a41e7d76ae91f30c141aa99943ed6c1c", network)
	if err != nil {
		logs.GetLogger().Fatal(err)
	}

	logs.GetLogger().Info(*mcsClient)
}
