package mcs

import (
	"context"
	"math/big"
)

const (
	McsPolygonMainApi   = "https://api.multichain.storage/api/"
	McsPolygonMumbaiApi = "https://calibration-mcs-api.filswan.com/api/"
)

type McsParams struct {
	ChainName               string
	McsApi                  string
	GasLimit                big.Int `json:"gas_limit"`
	LockTime                int     `json:"lock_time"`
	MintContractAddress     string  `json:"mint_contract_address"`
	PaymentContractAddress  string  `json:"payment_contract_address"`
	PaymentRecipientAddress string  `json:"payment_recipient_address"`
	PayMultiplyFactor       float64 `json:"pay_multiply_factor"`
	USDCAddress             string  `json:"usdc_address"`
}

func NewMcsParams(ChainName string) *McsParams {
	m := &McsParams{}
	if ChainName == "polygon.mainnet" || ChainName == "main" {
		m.McsApi = McsPolygonMainApi
		m.ChainName = "polygon.mainnet"
	} else if ChainName == "polygon.mumbai" {
		m.McsApi = McsPolygonMumbaiApi
	}
	var client = NewClient(m.McsApi)
	resParams, _ := client.NewGetParamsService().Do(context.Background())
	m.LockTime = resParams.Data.LockTime
	m.MintContractAddress = resParams.Data.MintContractAddress
	m.PaymentContractAddress = resParams.Data.PaymentContractAddress
	m.PaymentRecipientAddress = resParams.Data.PaymentRecipientAddress
	m.PayMultiplyFactor = resParams.Data.PayMultiplyFactor
	m.USDCAddress = resParams.Data.USDCAddress
	return m
}
