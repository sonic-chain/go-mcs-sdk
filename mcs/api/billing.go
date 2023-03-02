package api

import (
	"fmt"
	"go-mcs-sdk/mcs/common/constants"
	"net/url"
	"strings"

	"github.com/filswan/go-swan-lib/logs"
	libutils "github.com/filswan/go-swan-lib/utils"
)

type FileCoinPriceResponse struct {
	Response
	Data float64 `json:"data"`
}

func (mcsCient *McsClient) GetFileCoinPrice() (*float64, error) {
	apiUrl := libutils.UrlJoin(mcsCient.BaseUrl, constants.API_URL_BILLING_FILECOIN_PRICE)
	params := url.Values{}
	data, err := HttpGet(apiUrl, mcsCient.JwtToken, strings.NewReader(params.Encode()))
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	price, ok := data.(float64)
	if !ok {
		err := fmt.Errorf("invalid data type data:%s", data)
		logs.GetLogger().Error(err)
		return nil, err
	}

	return &price, nil
}

type LockPaymentInfo struct {
	WCid         string `json:"w_cid"`
	PayAmount    string `json:"pay_amount"`
	PayTxHash    string `json:"pay_tx_hash"`
	TokenAddress string `json:"token_address"`
}

type LockPaymentInfoResponse struct {
	Response
	Data LockPaymentInfo `json:"data"`
}

func (mcsCient *McsClient) GetLockPaymentInfo(fileUploadId int64) (*LockPaymentInfo, error) {
	apiUrl := libutils.UrlJoin(mcsCient.BaseUrl, constants.API_URL_BILLING_GET_PAYMENT_INFO)
	apiUrl = apiUrl + "?source_file_upload_id=" + fmt.Sprintf("%d", fileUploadId)
	params := url.Values{}

	data, err := HttpGet(apiUrl, mcsCient.JwtToken, strings.NewReader(params.Encode()))
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	lockPaymentInfo, ok := data.(LockPaymentInfo)
	if !ok {
		err := fmt.Errorf("invalid data type data:%s", data)
		logs.GetLogger().Error(err)
		return nil, err
	}

	return &lockPaymentInfo, nil
}

type BillingHistory struct {
	PayId        int64  `json:"pay_id"`
	PayTxHash    string `json:"pay_tx_hash"`
	PayAmount    string `json:"pay_amount"`
	UnlockAmount string `json:"unlock_amount"`
	FileName     string `json:"file_name"`
	PayloadCid   string `json:"payload_cid"`
	PayAt        int64  `json:"pay_at"`
	UnlockAt     int64  `json:"unlock_at"`
	Deadline     int64  `json:"deadline"`
	NetworkName  string `json:"network_name"`
	TokenName    string `json:"token_name"`
}

type BillingHistoryResponseData struct {
	Billing          []*BillingHistory `json:"billing"`
	TotalRecordCount int64             `json:"total_record_count"`
}

type BillingHistoryParams struct {
	PageNumber *int    `json:"page_number"`
	PageSize   *int    `json:"page_size"`
	FileName   *string `json:"file_name"`
	TxHash     *string `json:"tx_hash"`
	OrderBy    *string `json:"order_by"`
	IsAscend   *string `json:"is_ascend"`
}

func (mcsCient *McsClient) GetBillingHistory(billingHistoryParams BillingHistoryParams) ([]*BillingHistory, *int64, error) {
	apiUrl := libutils.UrlJoin(mcsCient.BaseUrl, constants.API_URL_BILLING_HISTORY)
	paramItems := []string{}
	if billingHistoryParams.PageNumber != nil {
		paramItems = append(paramItems, "page_number="+fmt.Sprintf("%d", *billingHistoryParams.PageNumber))
	}

	if billingHistoryParams.PageSize != nil {
		paramItems = append(paramItems, "page_size="+fmt.Sprintf("%d", *billingHistoryParams.PageSize))
	}

	if billingHistoryParams.FileName != nil {
		paramItems = append(paramItems, "file_name="+*billingHistoryParams.FileName)
	}

	if billingHistoryParams.TxHash != nil {
		paramItems = append(paramItems, "tx_hash="+*billingHistoryParams.TxHash)
	}

	if billingHistoryParams.OrderBy != nil {
		paramItems = append(paramItems, "order_by="+*billingHistoryParams.OrderBy)
	}

	if billingHistoryParams.IsAscend != nil {
		paramItems = append(paramItems, "is_ascend="+*billingHistoryParams.IsAscend)
	}

	if len(paramItems) > 0 {
		apiUrl = apiUrl + "?"
		for _, paramItem := range paramItems {
			apiUrl = apiUrl + paramItem + "&"
		}

		apiUrl = strings.TrimRight(apiUrl, "&")
	}

	data, err := HttpGet(apiUrl, mcsCient.JwtToken, nil)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, nil, err
	}

	billingHistoryResponseData, ok := data.(BillingHistoryResponseData)
	if !ok {
		err := fmt.Errorf("invalid data type data:%s", data)
		logs.GetLogger().Error(err)
		return nil, nil, err
	}

	return billingHistoryResponseData.Billing, &billingHistoryResponseData.TotalRecordCount, nil
}
