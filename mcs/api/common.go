package api

import (
	"go-mcs-sdk/mcs/api/common/constants"
	"go-mcs-sdk/mcs/api/common/utils"
	"net/url"
	"regexp"
	"strconv"
	"strings"

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

	var loginByApikeyResponse struct {
		JwtToken string `json:"jwt_token"`
	}

	err := utils.HttpPost(apiUrl, "", loginByApikeyParams, &loginByApikeyResponse)
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

func (mcsCient *McsClient) GetSystemParam() (*SystemParam, error) {
	apiUrl := libutils.UrlJoin(mcsCient.BaseUrl, constants.API_URL_MCS_GET_PARAMS)
	params := url.Values{}

	var systemParam SystemParam
	err := utils.HttpGet(apiUrl, mcsCient.JwtToken, strings.NewReader(params.Encode()), &systemParam)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return &systemParam, nil
}

func GetHistoricalAveragePriceVerified() (float64, error) {
	apiUrl := constants.API_URL_FIL_PRICE_API

	var storageStats struct {
		AverageCostPushMessage           string `json:"average_cost_push_message"`
		AverageDataCostSealing1TB        string `json:"average_data_cost_sealing_1TB"`
		AverageGasCostSealing1TB         string `json:"average_gas_cost_sealing_1TB"`
		AverageMinPieceSize              string `json:"average_min_piece_size"`
		AveragePricePerGBPerYear         string `json:"average_price_per_GB_per_year"`
		AverageVerifiedPricePerGBPerYear string `json:"average_verified_price_per_GB_per_year"`
		HistoricalAveragePriceVerified   string `json:"historical_average_price_verified"`
	}

	err := utils.HttpGet(apiUrl, "", nil, &storageStats)
	if err != nil {
		logs.GetLogger().Error(err)
		return -1, err
	}

	price := storageStats.HistoricalAveragePriceVerified
	reg := regexp.MustCompile(`\d+\.\d+`)
	result := reg.FindAllStringSubmatch(price, -1)

	priceFloat, err := strconv.ParseFloat(result[0][0], 64)
	if err != nil {
		logs.GetLogger().Error(err)
		return -1, err
	}

	return priceFloat, err
}

// USDC * 1e6
func GetAmount(fizeSizeByte int64, historicalAveragePriceVerified, fileCoinPrice float64, copyNumber int) (int64, error) {
	fileSizeGb := float64(fizeSizeByte) / constants.BYTES_1GB

	amount := historicalAveragePriceVerified * fileSizeGb * float64(constants.DURATION_DAYS_DEFAULT) * float64(copyNumber) * fileCoinPrice

	amount = amount * 1e6
	if amount <= 2 {
		amount = 2
	}

	return int64(amount), nil
}
