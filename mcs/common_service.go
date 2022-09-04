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
	GasLimit                big.Int `json:"GAS_LIMIT"`
	LockTime                int     `json:"LOCK_TIME"`
	MintContractAddress     string  `json:"MINT_CONTRACT_ADDRESS"`
	PaymentContractAddress  string  `json:"PAYMENT_CONTRACT_ADDRESS"`
	PaymentRecipientAddress string  `json:"PAYMENT_RECIPIENT_ADDRESS"`
	PayMultiplyFactor       float64 `json:"PAY_MULTIPLY_FACTOR"`
	USDCAddress             string  `json:"USDC_ADDRESS"`
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
