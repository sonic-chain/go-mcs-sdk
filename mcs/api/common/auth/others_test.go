package auth

import (
	"go-mcs-sdk/mcs/config"
	"testing"

	"github.com/filswan/go-swan-lib/logs"
)

var mcsClient *McsClient

func init() {
	if mcsClient != nil {
		return
	}

	apikey := config.GetConfig().Apikey
	accessToken := config.GetConfig().AccessToken
	network := config.GetConfig().Network

	var err error
	mcsClient, err = LoginByApikey(apikey, accessToken, network)
	if err != nil {
		logs.GetLogger().Fatal(err)
	}
}

func TestCheckLogin(t *testing.T) {
	mcsClient.JwtToken = "d"
	networkName, walletAddress, err := mcsClient.CheckLogin()
	if err != nil {
		logs.GetLogger().Fatal(err)
	}

	logs.GetLogger().Info(*networkName, ",", *walletAddress)
}

func TestGenerateApikey(t *testing.T) {
	apikey, accessToken, err := mcsClient.GenerateApikey(30)
	if err != nil {
		logs.GetLogger().Fatal(err)
	}

	logs.GetLogger().Info(*apikey, ",", *accessToken)
}

func TestDeleteApikey(t *testing.T) {
	err := mcsClient.DeleteApikey("2dkFLDsWNYDTkZkz6qB6PG")
	if err != nil {
		logs.GetLogger().Fatal(err)
	}
}

func TestRegisterEmail(t *testing.T) {
	response, err := mcsClient.RegisterEmail("fchen@nbai.io")
	if err != nil {
		logs.GetLogger().Fatal(err)
	}

	logs.GetLogger().Info(*response)
}

func TestGetApikeys(t *testing.T) {
	apikeys, err := mcsClient.GetApikeys()
	if err != nil {
		logs.GetLogger().Fatal(err)
	}

	for _, apikey := range apikeys {
		logs.GetLogger().Info(*apikey)
	}
}
