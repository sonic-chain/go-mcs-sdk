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
