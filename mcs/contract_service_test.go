package mcs

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"testing"
)

const (
	WalletAddress = "*"
	RpcEndpoint   = "*"
	PrivateKey    = "*"
	FilePath      = "*"
)

func TestContractApproveUSDCService_Do(t *testing.T) {
	p := NewMcsParams("polygon.mumbai")
	client := NewClient(p.McsApi)
	res, err := client.NewContractContractApproveUSDCService().SetWalletAddress(WalletAddress).
		SetUSDCAddress(p.USDCAddress).SetPaymentContractAddress(p.PaymentContractAddress).
		SetRpcEndpoint(RpcEndpoint).SetPrivateKey(PrivateKey).SetAmount(big.NewInt(100 * int64(math.Pow(10, 18)))).Do(context.Background())
	if err != nil {
		return
	}
	fmt.Println("test approve:", res)

}
func TestUploadFilePayService_Do(t *testing.T) {
	p := NewMcsParams("polygon.mumbai")
	client := NewClient(p.McsApi)
	res, err := client.NewUploadIpfsService().SetWalletAddress(WalletAddress).
		SetFilePath(FilePath).Do(context.Background())
	if err != nil {
		return
	}
	fmt.Println("file upload:", res.Data)
	//resParams, err := client.NewGetParamsService().Do(context.Background())
	//if err != nil {
	//	return
	//}
	resPrice, err := client.NewGetPriceRateService().Do(context.Background())
	if err != nil {
		return
	}
	resContract, err := client.NewContractUploadFilePayService().SetWalletAddress(WalletAddress).
		SetRpcEndpoint(RpcEndpoint).SetPrivateKey(PrivateKey).
		SetFileSize(res.Data.FileSize).SetWCid(res.Data.WCid).
		SetPaymentContractAddress(p.PaymentContractAddress).
		SetPaymentRecipientAddress(p.PaymentRecipientAddress).SetPayMultiplyFactor(p.PayMultiplyFactor).
		SetRate(resPrice.Data).SetLockTime(*big.NewInt(int64(p.LockTime))).Do(context.Background())
	fmt.Println("test upload pay:", resContract)
	if err != nil {
		return
	}
}

func TestContractMintNftService_Do(t *testing.T) {
	p := NewMcsParams("polygon.mumbai")
	client := NewClient(p.McsApi)
	client.SetJwtToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NjY0MzQ5NzksImlhdCI6MTY2NjM0ODU3OSwic3ViIjoiNzEyLHBvbHlnb24ubXVtYmFpIn0.AgnLYTjlxhlm4OvYixdzFk9uFZfXRHFcq8TYQAzVUvQ")
	res, err := client.NewUploadIpfsService().SetWalletAddress(WalletAddress).
		SetFilePath(FilePath).Do(context.Background())
	if err != nil {
		return
	}
	fmt.Println("file upload:", res.Data)
	resParams, err := client.NewGetParamsService().Do(context.Background())
	if err != nil {
		return
	}
	resPrice, err := client.NewGetPriceRateService().Do(context.Background())
	if err != nil {
		return
	}
	PayContractres, err := client.NewContractUploadFilePayService().SetWalletAddress(WalletAddress).
		SetRpcEndpoint(RpcEndpoint).SetPrivateKey(PrivateKey).
		SetFileSize(res.Data.FileSize).SetWCid(res.Data.WCid).
		SetPaymentContractAddress(resParams.Data.PaymentContractAddress).
		SetPaymentRecipientAddress(resParams.Data.PaymentRecipientAddress).SetPayMultiplyFactor(resParams.Data.PayMultiplyFactor).
		SetRate(resPrice.Data).SetLockTime(*big.NewInt(int64(resParams.Data.LockTime))).Do(context.Background())
	fmt.Println("test upload pay:", PayContractres)
	if err != nil {
		return
	}
	FileName := FilePath
	NftMetaUrl, err := client.NewUploadNftMetadataService().
		SetImageUrl(res.Data.IpfsURL).SetSize(res.Data.FileSize).SetFileName(FileName).
		SetWalletAddress(WalletAddress).SetTxHash(PayContractres).Do(context.Background())
	if err != nil {
		return
	}
	NftContractRes, TokenID, err := client.NewContractMintNftService().SetNftMetaUrl(NftMetaUrl.Data.IpfsURL).
		SetRpcEndpoint(RpcEndpoint).SetMintAddress(resParams.Data.MintContractAddress).
		SetWalletAddress(WalletAddress).SetPrivateKey(PrivateKey).Do(context.Background())
	fmt.Println(NftContractRes)

	client.NewGetMintInfoService().SetMintAddress(WalletAddress).
		SetPayloadCid(res.Data.PayloadCid).SetTxHash(NftContractRes).SetTokenId(TokenID.String()).
		SetSourceFileUploadId(res.Data.SourceFileUploadID).Do(context.Background())

}
