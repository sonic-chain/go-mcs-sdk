package mcs

import (
	"context"
	"net/http"
)

type GetPriceRateService struct {
	c *Client
}
type PriceRate struct {
	Status string  `json:"status"`
	Data   float64 `json:"data"`
}

func (s *GetPriceRateService) Do(ctx context.Context, opts ...RequestOption) (res *PriceRate, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/billing/price/filecoin",
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return &PriceRate{}, err
	}
	res = new(PriceRate)
	err = json.Unmarshal(data, res)

	if err != nil {
		return &PriceRate{}, err
	}

	return res, nil
}

type GetPaymentInfoService struct {
	c                  *Client
	PayloadCid         string
	SourceFileUploadId int64
	WalletAddress      string
}

func (s *GetPaymentInfoService) SetPayloadCid(PayloadCid string) *GetPaymentInfoService {
	s.PayloadCid = PayloadCid
	return s
}

func (s *GetPaymentInfoService) SetSourceFileUploadId(SourceFileUploadId int64) *GetPaymentInfoService {
	s.SourceFileUploadId = SourceFileUploadId
	return s
}

func (s *GetPaymentInfoService) SetWalletAddress(WalletAddress string) *GetPaymentInfoService {
	s.WalletAddress = WalletAddress
	return s
}

type PaymentInfo struct {
	Status string `json:"status"`
	Data   struct {
		WCid         string `json:"w_cid"`
		PayAmount    string `json:"pay_amount"`
		PayTxHash    string `json:"pay_tx_hash"`
		TokenAddress string `json:"token_address"`
	} `json:"data"`
}

func (s *GetPaymentInfoService) Do(ctx context.Context, opts ...RequestOption) (res *PaymentInfo, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/billing/deal/lockpayment/info",
	}
	r.setParam("payload_cid", s.PayloadCid)
	r.setParam("source_file_upload_id", s.SourceFileUploadId)
	r.setParam("wallet_address", s.WalletAddress)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return &PaymentInfo{}, err
	}
	res = new(PaymentInfo)
	err = json.Unmarshal(data, res)

	if err != nil {
		return &PaymentInfo{}, err
	}

	return res, nil
}
