package auth

import (
	"go-mcs-sdk/mcs/api/common/constants"
	"go-mcs-sdk/mcs/api/common/web"

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
