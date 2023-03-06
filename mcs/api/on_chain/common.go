package api

import (
	"go-mcs-sdk/mcs/api/common"
	"go-mcs-sdk/mcs/api/common/constants"
	"go-mcs-sdk/mcs/api/common/utils"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/filswan/go-swan-lib/logs"
	libutils "github.com/filswan/go-swan-lib/utils"
)

type OnChainClient struct {
	BaseUrl  string `json:"base_url"`
	JwtToken string `json:"jwt_token"`
}

func GetOnChainClientFromMcsClient(mcsClient common.McsClient) OnChainClient {
	var onChainClient = OnChainClient{}

	onChainClient.BaseUrl = mcsClient.BaseUrl
	onChainClient.JwtToken = mcsClient.JwtToken

	return onChainClient
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

func (onChainClient *OnChainClient) GetSystemParam() (*SystemParam, error) {
	apiUrl := libutils.UrlJoin(onChainClient.BaseUrl, constants.API_URL_MCS_GET_PARAMS)
	params := url.Values{}

	var systemParam SystemParam
	err := utils.HttpGet(apiUrl, onChainClient.JwtToken, strings.NewReader(params.Encode()), &systemParam)
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
