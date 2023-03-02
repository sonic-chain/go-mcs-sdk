package api

import (
	"encoding/json"
	"fmt"
	"go-mcs-sdk/mcs/common/constants"
	"go-mcs-sdk/mcs/common/utils"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/filswan/go-swan-lib/client/web"
	"github.com/filswan/go-swan-lib/logs"
	libutils "github.com/filswan/go-swan-lib/utils"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func HttpPost(uri, tokenString string, params interface{}, result interface{}) error {
	_, _, err := utils.HttpRequest(http.MethodPost, uri, &tokenString, params, nil, result)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	return nil
}

func HttpGet(uri, tokenString string, params interface{}, result interface{}) error {
	_, _, err := utils.HttpRequest(http.MethodGet, uri, &tokenString, params, nil, result)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}
	return nil
}

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

type LoginByApikeyResponse struct {
	Response
	Data struct {
		JwtToken string `json:"jwt_token"`
	} `json:"data"`
}

func LoginByApikey(apikey, accessToken, network string) (*McsClient, error) {
	loginByApikeyParams := LoginByApikeyParams{
		Apikey:      apikey,
		AccessToken: accessToken,
		Network:     network,
	}

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

	mcsClient := McsClient{
		Network:  network,
		BaseUrl:  apiUrlBase,
		JwtToken: loginByApikeyResponse.Data.JwtToken,
	}

	return &mcsClient, nil
}

type SystemParam struct {
	ChainName                   string  `json:"chain_name"`
	PaymentContractAddress      string  `json:"payment_contract_address"`
	PaymentRecipientAddress     string  `json:"payment_recipient_address"`
	DaoContractAddress          string  `json:"dao_contract_address"`
	DexAddress                  string  `json:"dex_address"`
	UsdcWFilPoolContract        string  `json:"usdc_wFil_pool_contract"`
	DefaultNftCollectionAddress string  `json:"default_nft_collection_address"`
	NftCollectionFactoryAddress string  `json:"nft_collection_factory_address"`
	UsdcAddress                 string  `json:"usdc_address"`
	GasLimit                    uint64  `json:"gas_limit"`
	LockTime                    int     `json:"lock_time"`
	PayMultiplyFactor           float32 `json:"pay_multiply_factor"`
	DaoThreshold                int     `json:"dao_threshold"`
	FilecoinPrice               float64 `json:"filecoin_price"`
}

type SystemParamResponse struct {
	Response
	Data SystemParam `json:"data"`
}

func (mcsCient *McsClient) GetSystemParam() (*SystemParam, error) {
	apiUrl := libutils.UrlJoin(mcsCient.BaseUrl, constants.API_URL_MCS_GET_PARAMS)
	params := url.Values{}
	response, err := web.HttpGet(apiUrl, mcsCient.JwtToken, strings.NewReader(params.Encode()))
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	var systemParamResponse SystemParamResponse
	err = json.Unmarshal(response, &systemParamResponse)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	if !strings.EqualFold(systemParamResponse.Status, constants.HTTP_STATUS_SUCCESS) {
		err := fmt.Errorf("get parameters failed, status:%s,message:%s", systemParamResponse.Status, systemParamResponse.Message)
		logs.GetLogger().Error(err)
		return nil, err
	}

	return &systemParamResponse.Data, nil
}

type FilPrice struct {
	Response
	Data struct {
		AverageCostPushMessage           string `json:"average_cost_push_message"`
		AverageDataCostSealing1TB        string `json:"average_data_cost_sealing_1TB"`
		AverageGasCostSealing1TB         string `json:"average_gas_cost_sealing_1TB"`
		AverageMinPieceSize              string `json:"average_min_piece_size"`
		AveragePricePerGBPerYear         string `json:"average_price_per_GB_per_year"`
		AverageVerifiedPricePerGBPerYear string `json:"average_verified_price_per_GB_per_year"`
		HistoricalAveragePriceVerified   string `json:"historical_average_price_verified"`
	} `json:"data"`
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
