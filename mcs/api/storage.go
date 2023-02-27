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

type SourceFileUpload struct {
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
		SourceFileUploads []*SourceFileUpload `json:"source_file_upload"`
		TotalRecordCount  int64               `json:"total_record_count"`
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

func (mcsCient *MCSClient) GetDeals(dealsParams DealsParams) ([]*SourceFileUpload, *int64, error) {
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

	return dealsResponse.Data.SourceFileUploads, &dealsResponse.Data.TotalRecordCount, nil
}
