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
	Response
	Data UploadFile `json:"data"`
}

func (mcsCient *McsClient) UploadFile(filePath string, fileType int) (*UploadFile, error) {
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
	Response
	Data struct {
		Deals            []*Deal `json:"source_file_upload"`
		TotalRecordCount int64   `json:"total_record_count"`
	} `json:"data"`
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

func (mcsCient *McsClient) GetDeals(dealsParams DealsParams) ([]*Deal, *int64, error) {
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
	var dealsResponse DealsResponse
	err := HttpGet(apiUrl, mcsCient.JwtToken, nil, &dealsResponse)
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
	Response
	Data struct {
		SourceFileUploadDeal SourceFileUploadDeal `json:"source_file_upload_deal"`
		DaoThreshold         int                  `json:"dao_threshold"`
		DaoSignatures        []*DaoSignature      `json:"dao_signature"`
	} `json:"data"`
}

func (mcsCient *McsClient) GetDealDetail(sourceFileUploadId, dealId int64) (*SourceFileUploadDeal, []*DaoSignature, *int, error) {
	params := strconv.FormatInt(dealId, 10) + "?source_file_upload_id=" + strconv.FormatInt(sourceFileUploadId, 10)
	apiUrl := libutils.UrlJoin(mcsCient.BaseUrl, constants.API_URL_STORAGE_GET_DEAL_DETAIL, params)
	var sourceFileUploadDealResponse SourceFileUploadDealResponse
	err := HttpGet(apiUrl, mcsCient.JwtToken, nil, &sourceFileUploadDealResponse)
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
	Response
	Data struct {
		OfflineDealLogs []*OfflineDealLog `json:"offline_deal_log"`
	} `json:"data"`
}

func (mcsCient *McsClient) GetDealLogs(offlineDealId int64) ([]*OfflineDealLog, error) {
	apiUrl := libutils.UrlJoin(mcsCient.BaseUrl, constants.API_URL_STORAGE_GET_DEAL_LOG, strconv.FormatInt(offlineDealId, 10))
	var offlineDealLogResponse OfflineDealLogResponse
	err := HttpGet(apiUrl, mcsCient.JwtToken, nil, &offlineDealLogResponse)
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
	Response
	Data struct {
		SourceFileUpload *SourceFileUpload `json:"source_file_upload"`
	} `json:"data"`
}

func (mcsCient *McsClient) GetSourceFileUpload(sourceFileUploadId int64) (*SourceFileUpload, error) {
	apiUrl := libutils.UrlJoin(mcsCient.BaseUrl, constants.API_URL_STORAGE_GET_SOURCE_FILE_UPLOAD, strconv.FormatInt(sourceFileUploadId, 10))

	var sourceFileUploadResponse SourceFileUploadResponse
	err := HttpGet(apiUrl, mcsCient.JwtToken, nil, &sourceFileUploadResponse)
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

func (mcsCient *McsClient) UnpinSourceFile(sourceFileUploadId int64) error {
	apiUrl := libutils.UrlJoin(mcsCient.BaseUrl, constants.API_URL_STORAGE_UNPIN_SOURCE_FILE, strconv.FormatInt(sourceFileUploadId, 10))

	var response Response
	err := HttpPost(apiUrl, mcsCient.JwtToken, nil, &response)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	if !strings.EqualFold(response.Status, constants.HTTP_STATUS_SUCCESS) {
		err := fmt.Errorf("get parameters failed, status:%s,message:%s", response.Status, response.Message)
		logs.GetLogger().Error(err)
		return err
	}

	return nil
}

type NftCollectionParams struct {
	Name            string  `json:"name"`
	Description     *string `json:"description"`
	ImageUrl        *string `json:"image_url"`
	ExternalLink    *string `json:"external_link"`
	SellerFee       *int    `json:"seller_fee"`
	WalletRecipient *string `json:"wallet_recipient"`
	TxHash          string  `json:"tx_hash"`
}

func (mcsCient *McsClient) WriteNftCollection(nftCollectionParams NftCollectionParams) error {
	apiUrl := libutils.UrlJoin(mcsCient.BaseUrl, constants.API_URL_STORAGE_WRITE_NFT_COLLECTION)

	var response Response
	err := HttpPost(apiUrl, mcsCient.JwtToken, nftCollectionParams, &response)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	if !strings.EqualFold(response.Status, constants.HTTP_STATUS_SUCCESS) {
		err := fmt.Errorf("get parameters failed, status:%s,message:%s", response.Status, response.Message)
		logs.GetLogger().Error(err)
		return err
	}

	return nil
}

type NftCollection struct {
	ID                int64   `json:"id"`
	Address           *string `json:"address"`
	WalletId          int64   `json:"wallet_id"`
	Name              string  `json:"name"`
	Description       *string `json:"description"`
	ImageUrl          *string `json:"image_url"`
	ExternalLink      *string `json:"external_link"`
	SellerFee         *int    `json:"seller_fee"`
	WalletIdRecipient *int64  `json:"wallet_id_recipient"`
	TxHash            string  `json:"tx_hash"`
	CreateAt          int64   `json:"create_at"`
	UpdateAt          int64   `json:"update_at"`
	WalletRecipient   string  `json:"wallet_recipient"`
	IsDefault         bool    `json:"is_default"`
}

type GetNftCollectionsResponse struct {
	Response
	Data []*NftCollection `json:"data"`
}

func (mcsCient *McsClient) GetNftCollections() ([]*NftCollection, error) {
	apiUrl := libutils.UrlJoin(mcsCient.BaseUrl, constants.API_URL_STORAGE_GET_NFT_COLLECTIONS)

	var getNftCollectionsResponse GetNftCollectionsResponse
	err := HttpGet(apiUrl, mcsCient.JwtToken, nil, &getNftCollectionsResponse)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	if !strings.EqualFold(getNftCollectionsResponse.Status, constants.HTTP_STATUS_SUCCESS) {
		err := fmt.Errorf("get parameters failed, status:%s,message:%s", getNftCollectionsResponse.Status, getNftCollectionsResponse.Message)
		logs.GetLogger().Error(err)
		return nil, err
	}

	return getNftCollectionsResponse.Data, nil
}

type RecordMintInfoParams struct {
	SourceFileIploadId int64   `json:"source_file_upload_id"`
	NftCollectionId    int64   `json:"nft_collection_id"`
	TxHash             string  `json:"tx_hash"`
	TokenId            int64   `json:"token_id"`
	Name               *string `json:"name"`
	Description        *string `json:"description"`
}

type SourceFileMint struct {
	ID                 int64   `json:"id"`
	SourceFileUploadId int64   `json:"source_file_upload_id"`
	NftTxHash          string  `json:"nft_tx_hash"`
	MintAddress        string  `json:"mint_address"`
	NftCollectionId    int64   `json:"nft_collection_id"`
	TokenId            int64   `json:"token_id"`
	Name               *string `json:"name"`
	Description        *string `json:"description"`
	CreateAt           int64   `json:"create_at"`
	UpdateAt           int64   `json:"update_at"`
}

type RecordMintInfoResponse struct {
	Response
	Data *SourceFileMint `json:"data"`
}

func (mcsCient *McsClient) RecordMintInfo(recordMintInfoParams *RecordMintInfoParams) (*SourceFileMint, error) {
	apiUrl := libutils.UrlJoin(mcsCient.BaseUrl, constants.API_URL_STORAGE_RECORD_MINT_INFO)

	var recordMintInfoResponse RecordMintInfoResponse
	err := HttpPost(apiUrl, mcsCient.JwtToken, nil, &recordMintInfoResponse)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	if !strings.EqualFold(recordMintInfoResponse.Status, constants.HTTP_STATUS_SUCCESS) {
		err := fmt.Errorf("%s failed, status:%s, message:%s", apiUrl, recordMintInfoResponse.Status, recordMintInfoResponse.Message)
		logs.GetLogger().Error(err)
		return nil, err
	}

	return recordMintInfoResponse.Data, nil
}

type SourceFileMintOut struct {
	SourceFileMint
	NftCollectionAddress  string  `json:"nft_collection_address"`
	NftCollectionName     *string `json:"nft_collection_name"`
	NftCollectionImageUrl *string `json:"nft_collection_image_url"`
}

type GetMintInfoResponse struct {
	Response
	Data []*SourceFileMintOut `json:"data"`
}

func (mcsCient *McsClient) GetMintInfo(sourceFileUploadId int64) ([]*SourceFileMintOut, error) {
	apiUrl := libutils.UrlJoin(mcsCient.BaseUrl, constants.API_URL_STORAGE_GET_MINT_INFO, strconv.FormatInt(sourceFileUploadId, 10))

	var getMintInfoResponse GetMintInfoResponse

	err := HttpPost(apiUrl, mcsCient.JwtToken, nil, &getMintInfoResponse)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	if !strings.EqualFold(getMintInfoResponse.Status, constants.HTTP_STATUS_SUCCESS) {
		err := fmt.Errorf("%s failed, status:%s, message:%s", apiUrl, getMintInfoResponse.Status, getMintInfoResponse.Message)
		logs.GetLogger().Error(err)
		return nil, err
	}

	return getMintInfoResponse.Data, nil
}
