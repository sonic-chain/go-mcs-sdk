package auth

import (
	"encoding/json"
	"fmt"
	"go-mcs-sdk/mcs/common/constants"
	"strings"

	"github.com/filswan/go-swan-lib/client/web"
	"github.com/filswan/go-swan-lib/logs"
	libutils "github.com/filswan/go-swan-lib/utils"
)

type MCSClient struct {
	Network  string `json:"network"`
	BaseUrl  string `json:"base_url"`
	JwtToken string `json:"jwt_token"`
}

type LoginByApikeyParams struct {
	Apikey      string `json:"apikey" binding:"required,min=1,max=100"`
	AccessToken string `json:"access_token" binding:"required,min=1,max=100"`
	Network     string `json:"network" binding:"required,min=1,max=65535"`
}

type LoginByApikeyResponse struct {
	Status string `json:"status"`
	Data   struct {
		JwtToken string `json:"jwt_token"`
	} `json:"data"`
	Message string `json:"message"`
}

func LoginByApikey(apikey, accessToken, network string) (*MCSClient, error) {
	loginByApikeyParams := LoginByApikeyParams{
		Apikey:      apikey,
		AccessToken: accessToken,
		Network:     network,
	}

	apiUrl := ""
	switch network {
	case constants.PAYMENT_CHAIN_NAME_POLYGON_MAINNET:
		apiUrl = constants.API_URL_MCS_POLYGON_MAINNET
	case constants.PAYMENT_CHAIN_NAME_POLYGON_MUMBAI:
		apiUrl = constants.API_URL_MCS_POLYGON_MUMBAI
	case constants.PAYMENT_CHAIN_NAME_BSC_TESTNET:
		apiUrl = constants.API_URL_MCS_BSC_TESTNET
	default:
		apiUrl = constants.API_URL_MCS_POLYGON_MAINNET
		network = constants.PAYMENT_CHAIN_NAME_POLYGON_MAINNET
	}

	apiUrl = libutils.UrlJoin(apiUrl, constants.LOGIN_BY_APIKEY)

	response, err := web.HttpPostNoToken(apiUrl, loginByApikeyParams)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	loginByApikeyResponse := &LoginByApikeyResponse{}
	err = json.Unmarshal(response, loginByApikeyResponse)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	if !strings.EqualFold(loginByApikeyResponse.Status, constants.HTTP_STATUS_SUCCESS) {
		err := fmt.Errorf("login failed,code:%s,message:%s", loginByApikeyResponse.Status, loginByApikeyResponse.Message)
		logs.GetLogger().Error(err)
		return nil, err
	}

	mcsClient := MCSClient{
		Network:  network,
		BaseUrl:  apiUrl,
		JwtToken: loginByApikeyResponse.Data.JwtToken,
	}

	return &mcsClient, nil
}
