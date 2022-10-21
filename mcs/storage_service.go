package mcs

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"path/filepath"
)

type GetUserTaskDealsService struct {
	c             *Client
	WalletAddress string
}
type UserTaskDeals struct {
	Status string            `json:"status"`
	Data   UserTaskDealsData `json:"data"`
}
type OfflineDeal struct {
	ID             int64       `json:"id"`
	CarFileID      int64       `json:"car_file_id"`
	DealCid        string      `json:"deal_cid"`
	MinerID        int64       `json:"miner_id"`
	Verified       bool        `json:"verified"`
	StartEpoch     int64       `json:"start_epoch"`
	SenderWalletID int64       `json:"sender_wallet_id"`
	Status         string      `json:"status"`
	DealID         interface{} `json:"deal_id"`
	OnChainStatus  string      `json:"on_chain_status"`
	UnlockTxHash   interface{} `json:"unlock_tx_hash"`
	UnlockAt       interface{} `json:"unlock_at"`
	CreateAt       int64       `json:"create_at"`
	UpdateAt       int64       `json:"update_at"`
	MinerFid       string      `json:"miner_fid"`
}
type SourceFileUpload struct {
	SourceFileUploadID int64         `json:"source_file_upload_id"`
	CarFileID          int64         `json:"car_file_id"`
	FileName           string        `json:"file_name"`
	FileSize           int64         `json:"file_size"`
	UploadAt           int64         `json:"upload_at"`
	Duration           int64         `json:"duration"`
	IpfsURL            string        `json:"ipfs_url"`
	PinStatus          string        `json:"pin_status"`
	PayloadCid         string        `json:"payload_cid"`
	WCid               string        `json:"w_cid"`
	Status             string        `json:"status"`
	DealSuccess        bool          `json:"deal_success"`
	IsMinted           bool          `json:"is_minted"`
	TokenID            string        `json:"token_id"`
	MintAddress        string        `json:"mint_address"`
	NftTxHash          string        `json:"nft_tx_hash"`
	OfflineDeal        []OfflineDeal `json:"offline_deal"`
}
type UserTaskDealsData struct {
	SourceFileUpload []SourceFileUpload `json:"source_file_upload"`
	TotalRecordCount int64              `json:"total_record_count"`
}

func (s *GetUserTaskDealsService) SetAddress(WalletAddress string) *GetUserTaskDealsService {
	s.WalletAddress = WalletAddress
	return s
}
func (s *GetUserTaskDealsService) Do(ctx context.Context, opts ...RequestOption) (res *UserTaskDeals, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/storage/tasks/deals",
	}
	r.setParam("wallet_address", s.WalletAddress)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return &UserTaskDeals{}, err
	}
	res = new(UserTaskDeals)
	err = json.Unmarshal(data, res)

	if err != nil {
		return &UserTaskDeals{}, err
	}

	return res, nil
}

type GetMintInfoService struct {
	c                  *Client
	PayloadCid         string
	SourceFileUploadId int64
	TxHash             string
	TokenId            string
	MintAddress        string
}

func (s *GetMintInfoService) SetPayloadCid(PayloadCid string) *GetMintInfoService {
	s.PayloadCid = PayloadCid
	return s
}

func (s *GetMintInfoService) SetSourceFileUploadId(SourceFileUploadId int64) *GetMintInfoService {
	s.SourceFileUploadId = SourceFileUploadId
	return s
}

func (s *GetMintInfoService) SetTxHash(TxHash string) *GetMintInfoService {
	s.TxHash = TxHash
	return s
}

func (s *GetMintInfoService) SetTokenId(TokenId string) *GetMintInfoService {
	s.TokenId = TokenId
	return s
}

func (s *GetMintInfoService) SetMintAddress(MintAddress string) *GetMintInfoService {
	s.MintAddress = MintAddress
	return s
}

type MintInfo struct {
	Status string `json:"status"`
	Data   struct {
		ID                 int64  `json:"id"`
		SourceFileUploadID int64  `json:"source_file_upload_id"`
		NftTxHash          string `json:"nft_tx_hash"`
		MintAddress        string `json:"mint_address"`
		TokenID            string `json:"token_id"`
		CreateAt           int64  `json:"create_at"`
		UpdateAt           int64  `json:"update_at"`
	} `json:"data"`
}

func (s *GetMintInfoService) Do(ctx context.Context, opts ...RequestOption) (res *MintInfo, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/storage/mint/info",
	}
	r.postBody = params{
		"source_file_upload_id": s.SourceFileUploadId,
		"payload_cid":           s.PayloadCid,
		"tx_hash":               s.TxHash,
		"token_id":              s.TokenId,
		"mint_address":          s.MintAddress,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return &MintInfo{}, err
	}
	res = new(MintInfo)
	err = json.Unmarshal(data, res)

	if err != nil {
		return &MintInfo{}, err
	}

	return res, nil
}

type GetDealDetailService struct {
	c                  *Client
	SourceFileUploadId int64
	WalletAddress      string
}

func (s *GetDealDetailService) SetSourceFileUploadId(SourceFileUploadId int64) *GetDealDetailService {
	s.SourceFileUploadId = SourceFileUploadId
	return s
}

func (s *GetDealDetailService) SetWalletAddress(WalletAddress string) *GetDealDetailService {
	s.WalletAddress = WalletAddress
	return s
}

type DealDetail struct {
	Status string `json:"status"`
	Data   struct {
		DaoSignature []struct {
			WalletSigner string      `json:"wallet_signer"`
			TxHash       interface{} `json:"tx_hash"`
			Status       interface{} `json:"status"`
			CreateAt     interface{} `json:"create_at"`
		} `json:"dao_signature"`
		DaoThreshold         int64 `json:"dao_threshold"`
		SourceFileUploadDeal struct {
			DealID                   interface{} `json:"deal_id"`
			DealCid                  interface{} `json:"deal_cid"`
			MessageCid               interface{} `json:"message_cid"`
			Height                   interface{} `json:"height"`
			PieceCid                 interface{} `json:"piece_cid"`
			VerifiedDeal             interface{} `json:"verified_deal"`
			StoragePricePerEpoch     interface{} `json:"storage_price_per_epoch"`
			Signature                interface{} `json:"signature"`
			SignatureType            interface{} `json:"signature_type"`
			CreatedAt                interface{} `json:"created_at"`
			PieceSizeFormat          interface{} `json:"piece_size_format"`
			StartHeight              interface{} `json:"start_height"`
			EndHeight                interface{} `json:"end_height"`
			Client                   interface{} `json:"client"`
			ClientCollateralFormat   interface{} `json:"client_collateral_format"`
			Provider                 interface{} `json:"provider"`
			ProviderTag              interface{} `json:"provider_tag"`
			VerifiedProvider         interface{} `json:"verified_provider"`
			ProviderCollateralFormat interface{} `json:"provider_collateral_format"`
			Status                   interface{} `json:"status"`
			NetworkName              interface{} `json:"network_name"`
			StoragePrice             interface{} `json:"storage_price"`
			IpfsURL                  string      `json:"ipfs_url"`
			FileName                 string      `json:"file_name"`
			WCid                     string      `json:"w_cid"`
			CarFilePayloadCid        string      `json:"car_file_payload_cid"`
			LockedAt                 int64       `json:"locked_at"`
			LockedFee                string      `json:"locked_fee"`
			Unlocked                 bool        `json:"unlocked"`
		} `json:"source_file_upload_deal"`
	} `json:"data"`
}

func (s *GetDealDetailService) Do(ctx context.Context, opts ...RequestOption) (res *DealDetail, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/storage/deal/detail/0",
	}

	r.setParam("wallet_address", s.WalletAddress)
	r.setParam("source_file_upload_id", s.SourceFileUploadId)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return &DealDetail{}, err
	}
	res = new(DealDetail)
	err = json.Unmarshal(data, res)

	if err != nil {
		return &DealDetail{}, err
	}

	return res, nil
}

type UploadIpfsService struct {
	c             *Client
	FilePath      string
	WalletAddress string
}

func (s *UploadIpfsService) SetFilePath(FilePath string) *UploadIpfsService {
	s.FilePath = FilePath
	return s
}

func (s *UploadIpfsService) SetWalletAddress(WalletAddress string) *UploadIpfsService {
	s.WalletAddress = WalletAddress
	return s
}

type UploadIpfsData struct {
	Status string `json:"status"`
	Data   struct {
		SourceFileUploadID int64  `json:"source_file_upload_id"`
		PayloadCid         string `json:"payload_cid"`
		IpfsURL            string `json:"ipfs_url"`
		FileSize           int64  `json:"file_size"`
		WCid               string `json:"w_cid"`
		Status             string `json:"status"`
	} `json:"data"`
}

func (s *UploadIpfsService) Do(ctx context.Context, opts ...RequestOption) (res *UploadIpfsData, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/storage/ipfs/upload",
	}

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, errFile1 := os.Open(s.FilePath)
	defer file.Close()
	part1, errFile1 := writer.CreateFormFile("file", filepath.Base(s.FilePath))
	_, errFile1 = io.Copy(part1, file)
	if errFile1 != nil {
		fmt.Println(errFile1)
		return
	}
	_ = writer.WriteField("duration", "525")
	_ = writer.WriteField("storage_copy", "5")
	_ = writer.WriteField("wallet_address", s.WalletAddress)
	errWriter := writer.Close()
	if err != nil {
		fmt.Println(errWriter)
		return
	}
	r.body = payload

	header := http.Header{}
	header.Add("Content-Type", writer.FormDataContentType())
	r.header = header
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return &UploadIpfsData{}, err
	}
	res = new(UploadIpfsData)
	err = json.Unmarshal(data, res)

	if err != nil {
		return &UploadIpfsData{}, err
	}

	return res, nil
}

type UploadNftMetadataService struct {
	c             *Client
	FileName      string
	WalletAddress string
	ImageUrl      string
	TxHash        string
	size          int64
}

func (s *UploadNftMetadataService) SetFileName(FileName string) *UploadNftMetadataService {
	s.FileName = FileName
	return s
}

func (s *UploadNftMetadataService) SetWalletAddress(WalletAddress string) *UploadNftMetadataService {
	s.WalletAddress = WalletAddress
	return s
}

func (s *UploadNftMetadataService) SetImageUrl(ImageUrl string) *UploadNftMetadataService {
	s.ImageUrl = ImageUrl
	return s
}

func (s *UploadNftMetadataService) SetTxHash(TxHash string) *UploadNftMetadataService {
	s.TxHash = TxHash
	return s
}

func (s *UploadNftMetadataService) SetSize(size int64) *UploadNftMetadataService {
	s.size = size
	return s
}

type NftMetadata struct {
	Status string `json:"status"`
	Data   struct {
		SourceFileUploadID int64  `json:"source_file_upload_id"`
		PayloadCid         string `json:"payload_cid"`
		IpfsURL            string `json:"ipfs_url"`
		FileSize           int64  `json:"file_size"`
		WCid               string `json:"w_cid"`
	} `json:"data"`
}
type Attributes struct {
	TraitType string `json:"trait_type"`
	Value     int64  `json:"value"`
}
type FileUrl struct {
	Name        string       `json:"name"`
	Image       string       `json:"image"`
	Description string       `json:"description"`
	TxHash      string       `json:"tx_hash"`
	Attributes  []Attributes `json:"attributes"`
	ExternalURL string       `json:"external_url"`
}

func (s *UploadNftMetadataService) Do(ctx context.Context, opts ...RequestOption) (res *NftMetadata, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/storage/ipfs/upload",
	}
	attributes := Attributes{TraitType: "Size", Value: s.size}
	var attributesList []Attributes
	attributesList = append(attributesList, attributes)
	file := &FileUrl{Name: s.FileName, Image: s.ImageUrl, TxHash: s.TxHash, ExternalURL: s.ImageUrl, Attributes: attributesList}
	filString, _ := json.Marshal(file)
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition",
		fmt.Sprintf(`form-data; name="file"; filename="%s"`, s.FileName))
	h.Set("Content-Type", "application/json")
	p, _ := writer.CreatePart(h)
	p.Write(filString)
	_ = writer.WriteField("duration", "525")
	_ = writer.WriteField("file_type", "1")
	_ = writer.WriteField("wallet_address", s.WalletAddress)
	errWriter := writer.Close()
	if err != nil {
		fmt.Println(errWriter)
		return
	}

	r.body = payload

	header := http.Header{}
	header.Add("Content-Type", writer.FormDataContentType())
	r.header = header
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return &NftMetadata{}, err
	}
	res = new(NftMetadata)
	err = json.Unmarshal(data, res)

	if err != nil {
		return &NftMetadata{}, err
	}

	return res, nil
}
