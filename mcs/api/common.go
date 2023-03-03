package api

import (
	"encoding/json"
	"fmt"
	"go-mcs-sdk/mcs/common/constants"
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

func (mcsCient *MCSClient) GetSystemParam() (*SystemParam, error) {
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

func GetHistoricalAveragePriceVerified() (float64, error) { //unit:FIL/GiB/Day
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

	return priceFloat, err
}

func GetAmount(fizeSizeByte int64, historicalAveragePriceVerified, fileCoinPrice float64, copyNumber int) (float64, error) {
	fileSizeGb := float64(fizeSizeByte) / constants.BYTES_1GB

	amount := historicalAveragePriceVerified * fileSizeGb * float64(constants.DURATION_DAYS_DEFAULT) * float64(copyNumber) * fileCoinPrice

	if amount <= 0 {
		amount = 0.000002
	}

	return amount, nil
}
