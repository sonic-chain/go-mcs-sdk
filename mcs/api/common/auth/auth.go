package auth

import (
	"go-mcs-sdk/mcs/api/common/constants"
	"go-mcs-sdk/mcs/api/common/web"

	"github.com/filswan/go-swan-lib/logs"
	libutils "github.com/filswan/go-swan-lib/utils"
)

type McsClient struct {
	Network  string `json:"network"`
	BaseUrl  string `json:"base_url"`
	JwtToken string `json:"jwt_token"`
}

type LoginByApikeyParams struct {
	Apikey      string `json:"apikey" binding:"required,min=1,max=100"`
	AccessToken string `json:"access_token" binding:"required,min=1,max=100"`
	Network     string `json:"network" binding:"required,min=1,max=65535"`
}

func LoginByApikey(apikey, accessToken, network string) (*McsClient, error) {
	var params struct {
		Apikey      string `json:"apikey" binding:"required,min=1,max=100"`
		AccessToken string `json:"access_token" binding:"required,min=1,max=100"`
		Network     string `json:"network" binding:"required,min=1,max=65535"`
	}

	params.Apikey = apikey
	params.AccessToken = accessToken
	params.Network = network

	apiUrlBase := ""
	switch network {
	case constants.PAYMENT_CHAIN_NAME_POLYGON_MAINNET:
		apiUrlBase = constants.API_URL_MCS_POLYGON_MAINNET
	case constants.PAYMENT_CHAIN_NAME_POLYGON_MUMBAI:
		apiUrlBase = constants.API_URL_MCS_POLYGON_MUMBAI
	case constants.PAYMENT_CHAIN_NAME_BSC_TESTNET:
		apiUrlBase = constants.API_URL_MCS_BSC_TESTNET
	default:
		apiUrlBase = constants.API_URL_MCS_POLYGON_MAINNET
		network = constants.PAYMENT_CHAIN_NAME_POLYGON_MAINNET
	}

	apiUrl := libutils.UrlJoin(apiUrlBase, constants.LOGIN_BY_APIKEY)

	var loginByApikeyResponse struct {
		JwtToken string `json:"jwt_token"`
	}

	err := web.HttpPost(apiUrl, "", params, &loginByApikeyResponse)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	mcsClient := McsClient{
		Network:  network,
		BaseUrl:  apiUrlBase,
		JwtToken: loginByApikeyResponse.JwtToken,
	}

	return &mcsClient, nil
}
