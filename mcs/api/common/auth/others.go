package auth

import (
	"go-mcs-sdk/mcs/api/common/constants"
	"go-mcs-sdk/mcs/api/common/web"
	"strconv"

	"github.com/filswan/go-swan-lib/logs"
	libutils "github.com/filswan/go-swan-lib/utils"
)

func (mcsClient *McsClient) CheckLogin() (*string, *string, error) {
	apiUrl := libutils.UrlJoin(mcsClient.BaseUrl, constants.API_URL_USER_CHECK_LOGIN)

	var response struct {
		NetworkName   string `json:"network_name"`
		WalletAddress string `json:"wallet_address"`
	}

	err := web.HttpPost(apiUrl, mcsClient.JwtToken, nil, &response)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, nil, err
	}

	return &response.NetworkName, &response.WalletAddress, nil
}

func (mcsClient *McsClient) GenerateApikey(validDys int) (*string, *string, error) {
	apiUrl := libutils.UrlJoin(mcsClient.BaseUrl, constants.API_URL_USER_GENERATE_APIKEY)
	apiUrl = apiUrl + "?valid_days=" + strconv.Itoa(validDys)

	var response struct {
		Apikey      string `json:"apikey"`
		AccessToken string `json:"access_token"`
	}

	err := web.HttpGet(apiUrl, mcsClient.JwtToken, nil, &response)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, nil, err
	}

	return &response.Apikey, &response.AccessToken, nil
}

func (mcsClient *McsClient) DeleteApikey(apikey string) error {
	apiUrl := libutils.UrlJoin(mcsClient.BaseUrl, constants.API_URL_USER_DELETE_APIKEY)
	apiUrl = apiUrl + "?apikey=" + apikey

	err := web.HttpPut(apiUrl, mcsClient.JwtToken, nil, nil)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	return nil
}

func (mcsClient *McsClient) RegisterEmail(email string) (*string, error) {
	apiUrl := libutils.UrlJoin(mcsClient.BaseUrl, constants.API_URL_USER_REGISTER_EMAIL)
	var params struct {
		Email string `json:"email"`
	}
	params.Email = email

	var response string
	err := web.HttpPost(apiUrl, mcsClient.JwtToken, &params, &response)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return &response, nil
}
