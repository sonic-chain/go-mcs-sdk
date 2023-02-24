package api

import (
	"encoding/json"
	"fmt"
	"go-mcs-sdk/mcs/common/constants"
	"regexp"
	"strconv"
	"strings"

	libutils "github.com/filswan/go-swan-lib/utils"

	"github.com/filswan/go-swan-lib/client/web"
	"github.com/filswan/go-swan-lib/logs"
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

type FilPrice struct {
	Data struct {
		AverageCostPushMessage           string `json:"average_cost_push_message"`
		AverageDataCostSealing1TB        string `json:"average_data_cost_sealing_1TB"`
		AverageGasCostSealing1TB         string `json:"average_gas_cost_sealing_1TB"`
		AverageMinPieceSize              string `json:"average_min_piece_size"`
		AveragePricePerGBPerYear         string `json:"average_price_per_GB_per_year"`
		AverageVerifiedPricePerGBPerYear string `json:"average_verified_price_per_GB_per_year"`
		HistoricalAveragePriceVerified   string `json:"historical_average_price_verified"`
	} `json:"data"`
	Status string `json:"status"`
}

func GetFilPrice() (float64, error) {
	response, err := web.HttpGetNoToken(constants.API_URL_FIL_PRICE_API, nil)
	if err != nil {
		logs.GetLogger().Error(err)
		return -1, err
	}
	filPrice := new(FilPrice)
	err = json.Unmarshal(response, filPrice)
	if err != nil {
		logs.GetLogger().Error(err)
		return -1, err
	}

	price := filPrice.Data.HistoricalAveragePriceVerified
	reg := regexp.MustCompile(`\d+\.\d+`)
	result := reg.FindAllStringSubmatch(price, -1)

	priceFloat, err := strconv.ParseFloat(result[0][0], 64)
	if err != nil {
		logs.GetLogger().Error(err)
		return -1, err
	}

	priceFloat = priceFloat / constants.BYTES_1GB / 1e8
	return priceFloat, err
}

func GetAmount(sizeByte int64, rate float64) (float64, error) {
	price, err := GetFilPrice()
	if err != nil {
		logs.GetLogger().Error(err)
		return -1, err
	}

	amount := float64(sizeByte) * price * rate * constants.DURATION_DAYS_DEFAULT / 365

	if amount == 0 {
		amount = 0.000002
	}

	return amount, nil
}
