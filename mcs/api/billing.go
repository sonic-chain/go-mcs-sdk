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
	var fileCoinPriceResponse FileCoinPriceResponse
	err := HttpGet(apiUrl, mcsCient.JwtToken, strings.NewReader(params.Encode()), &fileCoinPriceResponse)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	if !strings.EqualFold(fileCoinPriceResponse.Status, constants.HTTP_STATUS_SUCCESS) {
		err := fmt.Errorf("get parameters failed, status:%s,message:%s", fileCoinPriceResponse.Status, fileCoinPriceResponse.Message)
		logs.GetLogger().Error(err)
		return nil, err
	}

	return &fileCoinPriceResponse.Data, nil
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
	var lockPaymentInfoResponse LockPaymentInfoResponse
	err := HttpGet(apiUrl, mcsCient.JwtToken, strings.NewReader(params.Encode()), &lockPaymentInfoResponse)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	if !strings.EqualFold(lockPaymentInfoResponse.Status, constants.HTTP_STATUS_SUCCESS) {
		err := fmt.Errorf("get parameters failed, status:%s,message:%s", lockPaymentInfoResponse.Status, lockPaymentInfoResponse.Message)
		logs.GetLogger().Error(err)
		return nil, err
	}

	return &lockPaymentInfoResponse.Data, nil
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

type BillingHistoryResponse struct {
	Response
	Data struct {
		Billing          []*BillingHistory `json:"billing"`
		TotalRecordCount int64             `json:"total_record_count"`
	} `json:"data"`
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

	logs.GetLogger().Info(apiUrl)
	var billingHistoryResponse BillingHistoryResponse
	err := HttpGet(apiUrl, mcsCient.JwtToken, nil, &billingHistoryResponse)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, nil, err
	}

	if !strings.EqualFold(billingHistoryResponse.Status, constants.HTTP_STATUS_SUCCESS) {
		err := fmt.Errorf("get parameters failed, status:%s,message:%s", billingHistoryResponse.Status, billingHistoryResponse.Message)
		logs.GetLogger().Error(err)
		return nil, nil, err
	}

	return billingHistoryResponse.Data.Billing, &billingHistoryResponse.Data.TotalRecordCount, nil
}
