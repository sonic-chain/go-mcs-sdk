package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-mcs-sdk/mcs/common/constants"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/filswan/go-swan-lib/client/web"
	"github.com/filswan/go-swan-lib/logs"
	libutils "github.com/filswan/go-swan-lib/utils"
)

type UploadFile struct {
	SourceFileUploadId int64  `json:"source_file_upload_id"`
	PayloadCid         string `json:"payload_cid"`
	IpfsUrl            string `json:"ipfs_url"`
	FileSize           int64  `json:"file_size"`
	WCid               string `json:"w_cid"`
	Status             string `json:"status"`
}

type UploadFileResponse struct {
	Status  string     `json:"status"`
	Data    UploadFile `json:"data"`
	Message string     `json:"message"`
}

func (mcsCient *MCSClient) UploadFile(filePath string, fileType int) (*UploadFile, error) {
	httpRequestUrl := libutils.UrlJoin(mcsCient.BaseUrl, constants.API_URL_STORAGE_UPLOAD_FILE)
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, err := os.Open(filePath)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	defer file.Close()

	part1, err := writer.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	_, err = io.Copy(part1, file)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	err = writer.WriteField("duration", strconv.Itoa(constants.DURATION_DAYS_DEFAULT))
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	err = writer.WriteField("file_type", strconv.Itoa(fileType))
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	err = writer.Close()
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	httpClient := &http.Client{}
	req, err := http.NewRequest("POST", httpRequestUrl, payload)

	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", mcsCient.JwtToken))
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := httpClient.Do(req)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	var uploadFileResponse UploadFileResponse
	err = json.Unmarshal(body, &uploadFileResponse)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	if !strings.EqualFold(uploadFileResponse.Status, constants.HTTP_STATUS_SUCCESS) {
		err := fmt.Errorf("get parameters failed, status:%s,message:%s", uploadFileResponse.Status, uploadFileResponse.Message)
		logs.GetLogger().Error(err)
		return nil, err
	}

	return &uploadFileResponse.Data, nil
}

type OfflineDeal struct {
	Id             int64   `json:"id"`
	CarFileId      int64   `json:"car_file_id"`
	DealCid        string  `json:"deal_cid"`
	MinerId        int64   `json:"miner_id"`
	Verified       bool    `json:"verified"`
	StartEpoch     int     `json:"start_epoch"`
	SenderWalletId int64   `json:"sender_wallet_id"`
	Status         string  `json:"status"`
	DealId         *int64  `json:"deal_id"`
	OnChainStatus  *string `json:"on_chain_status"`
	UnlockTxHash   *string `json:"unlock_tx_hash"`
	UnlockAt       *int64  `json:"unlock_at"`
	Note           *string `json:"note"`
	NetworkId      int64   `json:"network_id"`
	MinerFid       string  `json:"miner_fid"`
	CreateAt       int64   `json:"create_at"`
	UpdateAt       int64   `json:"update_at"`
}

type Deal struct {
	SourceFileUploadId int64          `json:"source_file_upload_id"`
	FileName           string         `json:"file_name"`
	FileSize           int64          `json:"file_size"`
	UploadAt           int64          `json:"upload_at"`
	Duration           int            `json:"duration"`
	IpfsUrl            string         `json:"ipfs_url"`
	PinStatus          string         `json:"pin_status"`
	PayAmount          string         `json:"pay_amount"`
	Status             string         `json:"status"`
	Note               string         `json:"note"`
	IsFree             bool           `json:"is_free"`
	IsMinted           bool           `json:"is_minted"`
	RefundedBySelf     bool           `json:"refunded_by_self"`
	OfflineDeals       []*OfflineDeal `json:"offline_deal"`
}

type DealsResponse struct {
	Status string `json:"status"`
	Data   struct {
		Deals            []*Deal `json:"source_file_upload"`
		TotalRecordCount int64   `json:"total_record_count"`
	} `json:"data"`
	Message string `json:"message"`
}

type DealsParams struct {
	PageNumber *int    `json:"page_number"`
	PageSize   *int    `json:"page_size"`
	FileName   *string `json:"file_name"`
	Status     *string `json:"status"`
	IsMinted   *string `json:"is_minted"`
	OrderBy    *string `json:"order_by"`
	IsAscend   *string `json:"is_ascend"`
}

func (mcsCient *MCSClient) GetDeals(dealsParams DealsParams) ([]*Deal, *int64, error) {
	apiUrl := libutils.UrlJoin(mcsCient.BaseUrl, constants.API_URL_STORAGE_GET_DEALS)
	paramItems := []string{}
	if dealsParams.PageNumber != nil {
		paramItems = append(paramItems, "page_number="+fmt.Sprintf("%d", *dealsParams.PageNumber))
	}

	if dealsParams.PageSize != nil {
		paramItems = append(paramItems, "page_size="+fmt.Sprintf("%d", *dealsParams.PageSize))
	}

	if dealsParams.FileName != nil {
		paramItems = append(paramItems, "file_name="+*dealsParams.FileName)
	}

	if dealsParams.Status != nil {
		paramItems = append(paramItems, "status="+*dealsParams.Status)
	}

	if dealsParams.IsMinted != nil {
		paramItems = append(paramItems, "is_minted="+*dealsParams.IsMinted)
	}

	if dealsParams.OrderBy != nil {
		paramItems = append(paramItems, "order_by="+*dealsParams.OrderBy)
	}

	if dealsParams.IsAscend != nil {
		paramItems = append(paramItems, "is_ascend="+*dealsParams.IsAscend)
	}

	if len(paramItems) > 0 {
		apiUrl = apiUrl + "?"
		for _, paramItem := range paramItems {
			apiUrl = apiUrl + paramItem + "&"
		}

		apiUrl = strings.TrimRight(apiUrl, "&")
	}

	logs.GetLogger().Info(apiUrl)
	response, err := web.HttpGet(apiUrl, mcsCient.JwtToken, nil)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, nil, err
	}

	var dealsResponse DealsResponse
	err = json.Unmarshal(response, &dealsResponse)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, nil, err
	}

	if !strings.EqualFold(dealsResponse.Status, constants.HTTP_STATUS_SUCCESS) {
		err := fmt.Errorf("get parameters failed, status:%s,message:%s", dealsResponse.Status, dealsResponse.Message)
		logs.GetLogger().Error(err)
		return nil, nil, err
	}

	return dealsResponse.Data.Deals, &dealsResponse.Data.TotalRecordCount, nil
}

type SourceFileUploadDeal struct {
	DealID                   *int    `json:"deal_id"`
	DealCid                  *string `json:"deal_cid"`
	MessageCid               *string `json:"message_cid"`
	Height                   *int    `json:"height"`
	PieceCid                 *string `json:"piece_cid"`
	VerifiedDeal             *bool   `json:"verified_deal"`
	StoragePricePerEpoch     *int    `json:"storage_price_per_epoch"`
	Signature                *string `json:"signature"`
	SignatureType            *string `json:"signature_type"`
	CreatedAt                *int    `json:"created_at"`
	PieceSizeFormat          *string `json:"piece_size_format"`
	StartHeight              *int    `json:"start_height"`
	EndHeight                *int    `json:"end_height"`
	Client                   *string `json:"client"`
	ClientCollateralFormat   *string `json:"client_collateral_format"`
	Provider                 *string `json:"provider"`
	ProviderTag              *string `json:"provider_tag"`
	VerifiedProvider         *int    `json:"verified_provider"`
	ProviderCollateralFormat *string `json:"provider_collateral_format"`
	Status                   *int    `json:"status"`
	NetworkName              *string `json:"network_name"`
	StoragePrice             *int    `json:"storage_price"`
	IpfsUrl                  string  `json:"ipfs_url"`
	FileName                 string  `json:"file_name"`
	WCid                     string  `json:"w_cid"`
	CarFilePayloadCid        string  `json:"car_file_payload_cid"`
	LockedAt                 int64   `json:"locked_at"`
	LockedFee                string  `json:"locked_fee"`
	Unlocked                 bool    `json:"unlocked"`
}

type DaoSignature struct {
	WalletSigner string  `json:"wallet_signer"`
	TxHash       *string `json:"tx_hash"`
	Status       *string `json:"status"`
	CreateAt     *int64  `json:"create_at"`
}

type SourceFileUploadDealResponse struct {
	Status string `json:"status"`
	Data   struct {
		SourceFileUploadDeal SourceFileUploadDeal `json:"source_file_upload_deal"`
		DaoThreshold         int                  `json:"dao_threshold"`
		DaoSignatures        []*DaoSignature      `json:"dao_signature"`
	} `json:"data"`
	Message string `json:"message"`
}

func (mcsCient *MCSClient) GetDealDetail(sourceFileUploadId, dealId int64) (*SourceFileUploadDeal, []*DaoSignature, *int, error) {
	params := strconv.FormatInt(dealId, 10) + "?source_file_upload_id=" + strconv.FormatInt(sourceFileUploadId, 10)
	apiUrl := libutils.UrlJoin(mcsCient.BaseUrl, constants.API_URL_STORAGE_GET_DEAL_DETAIL, params)
	response, err := web.HttpGet(apiUrl, mcsCient.JwtToken, nil)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, nil, nil, err
	}

	var sourceFileUploadDealResponse SourceFileUploadDealResponse
	err = json.Unmarshal(response, &sourceFileUploadDealResponse)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, nil, nil, err
	}

	if !strings.EqualFold(sourceFileUploadDealResponse.Status, constants.HTTP_STATUS_SUCCESS) {
		err := fmt.Errorf("get parameters failed, status:%s,message:%s", sourceFileUploadDealResponse.Status, sourceFileUploadDealResponse.Message)
		logs.GetLogger().Error(err)
		return nil, nil, nil, err
	}

	sourceFileUploadDeal := &sourceFileUploadDealResponse.Data.SourceFileUploadDeal
	daoSignatures := sourceFileUploadDealResponse.Data.DaoSignatures
	daoThreshold := &sourceFileUploadDealResponse.Data.DaoThreshold

	return sourceFileUploadDeal, daoSignatures, daoThreshold, nil
}

type OfflineDealLog struct {
	Id             int64  `json:"id"`
	OfflineDealId  int64  `json:"offline_deal_id"`
	OnChainStatus  string `json:"on_chain_status"`
	OnChainMessage string `json:"on_chain_message"`
	CreateAt       int64  `json:"create_at"`
}

type OfflineDealLogResponse struct {
	Status string `json:"status"`
	Data   struct {
		OfflineDealLogs []*OfflineDealLog `json:"offline_deal_log"`
	} `json:"data"`
	Message string `json:"message"`
}

func (mcsCient *MCSClient) GetDealLogs(offlineDealId int64) ([]*OfflineDealLog, error) {
	apiUrl := libutils.UrlJoin(mcsCient.BaseUrl, constants.API_URL_STORAGE_GET_DEAL_LOG, strconv.FormatInt(offlineDealId, 10))
	response, err := web.HttpGet(apiUrl, mcsCient.JwtToken, nil)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	var offlineDealLogResponse OfflineDealLogResponse
	err = json.Unmarshal(response, &offlineDealLogResponse)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	if !strings.EqualFold(offlineDealLogResponse.Status, constants.HTTP_STATUS_SUCCESS) {
		err := fmt.Errorf("get parameters failed, status:%s,message:%s", offlineDealLogResponse.Status, offlineDealLogResponse.Message)
		logs.GetLogger().Error(err)
		return nil, err
	}

	offlineDealLogs := offlineDealLogResponse.Data.OfflineDealLogs

	return offlineDealLogs, nil
}

type SourceFileUpload struct {
	WCid   string `json:"w_cid"`
	Status string `json:"status"`
	IsFree bool   `json:"is_free"`
}

type SourceFileUploadResponse struct {
	Status string `json:"status"`
	Data   struct {
		SourceFileUpload *SourceFileUpload `json:"source_file_upload"`
	} `json:"data"`
	Message string `json:"message"`
}

func (mcsCient *MCSClient) GetSourceFileUpload(sourceFileUploadId int64) (*SourceFileUpload, error) {
	apiUrl := libutils.UrlJoin(mcsCient.BaseUrl, constants.API_URL_STORAGE_GET_SOURCE_FILE_UPLOAD, strconv.FormatInt(sourceFileUploadId, 10))
	response, err := web.HttpGet(apiUrl, mcsCient.JwtToken, nil)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	var sourceFileUploadResponse SourceFileUploadResponse
	err = json.Unmarshal(response, &sourceFileUploadResponse)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	if !strings.EqualFold(sourceFileUploadResponse.Status, constants.HTTP_STATUS_SUCCESS) {
		err := fmt.Errorf("get parameters failed, status:%s,message:%s", sourceFileUploadResponse.Status, sourceFileUploadResponse.Message)
		logs.GetLogger().Error(err)
		return nil, err
	}

	sourceFileUpload := sourceFileUploadResponse.Data.SourceFileUpload

	return sourceFileUpload, nil
}
