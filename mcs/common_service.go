package mcs

import (
	"context"
	"math/big"
	"net/http"
)

type GetParamsService struct {
	c *Client
}

type ParamsData struct {
	ChainName               string  `json:"chain_name"`
	GasLimit                big.Int `json:"gas_limit"`
	LockTime                int     `json:"lock_time"`
	MintContractAddress     string  `json:"mint_contract_address"`
	PaymentContractAddress  string  `json:"payment_contract_address"`
	PaymentRecipientAddress string  `json:"payment_recipient_address"`
	PayMultiplyFactor       float64 `json:"pay_multiply_factor"`
	USDCAddress             string  `json:"usdc_address"`
}

type Params struct {
	Status string      `json:"status"`
	Data   *ParamsData `json:"data"`
}

func (s *GetParamsService) Do(ctx context.Context, opts ...RequestOption) (res *Params, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/common/system/params",
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return &Params{}, err
	}
	res = new(Params)
	err = json.Unmarshal(data, res)

	if err != nil {
		return &Params{}, err
	}

	return res, nil
}
