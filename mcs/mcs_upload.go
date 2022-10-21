package mcs

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	common2 "go-mcs-sdk/mcs/common"
	"log"
	"math"
	"math/big"
)

type MCSUpload struct {
	ChainName      string
	WalletAddress  string
	RpcEndpoint    string
	PrivateKey     string
	FilePath       string
	UploadIpfsData *UploadIpfsData
	PaymentTxHash  string
	token          string
	client         *Client
	params         *McsParams
	MintInfo       *MintInfo
}

func GetToken(mcs *MCSUpload) (*Client, *McsParams) {
	params := NewMcsParams(mcs.ChainName)
	client := NewClient(params.McsApi)
	user, _ := client.NewUserRegisterService().SetWalletAddress(mcs.WalletAddress).Do(context.Background())
	nonce := user.Data.Nonce
	privateKey, err := crypto.HexToECDSA(mcs.PrivateKey)
	signature, err := common2.PersonalSign(nonce, privateKey)
	if err != nil {
		log.Fatal(err)
	}
	jwt, _ := client.NewUserLoginService().SetNetwork(mcs.ChainName).SetNonce(nonce).SetWalletAddress(mcs.WalletAddress).
		SetSignature(signature).Do(context.Background())
	mcs.token = jwt.Data.JwtToken
	client.SetJwtToken(mcs.token)
	fmt.Println("client.JwtToken:::", client.JwtToken)
	return client, params
}

func NewMCSUpload(mcs *MCSUpload) *MCSUpload {
	client, params := GetToken(mcs)
	mcs.UploadIpfsData, _ = client.NewUploadIpfsService().SetWalletAddress(mcs.WalletAddress).
		SetFilePath(mcs.FilePath).Do(context.Background())
	fmt.Println("file upload:", mcs.UploadIpfsData.Data)
	if mcs.UploadIpfsData.Data.Status != "Free" {
		res, _ := client.NewContractContractApproveUSDCService().SetWalletAddress(mcs.WalletAddress).
			SetUSDCAddress(params.USDCAddress).SetPaymentContractAddress(params.PaymentContractAddress).
			SetRpcEndpoint(mcs.RpcEndpoint).SetPrivateKey(mcs.PrivateKey).SetAmount(big.NewInt(100 * int64(math.Pow(10, 18)))).Do(context.Background())
		fmt.Println("approve usdc tx:", res)
		resPrice, _ := client.NewGetPriceRateService().Do(context.Background())
		resContract, _ := client.NewContractUploadFilePayService().SetWalletAddress(mcs.WalletAddress).
			SetRpcEndpoint(mcs.RpcEndpoint).SetPrivateKey(mcs.PrivateKey).
			SetFileSize(mcs.UploadIpfsData.Data.FileSize).SetWCid(mcs.UploadIpfsData.Data.WCid).
			SetPaymentContractAddress(params.PaymentContractAddress).
			SetPaymentRecipientAddress(params.PaymentRecipientAddress).SetPayMultiplyFactor(params.PayMultiplyFactor).
			SetRate(resPrice.Data).SetLockTime(*big.NewInt(int64(params.LockTime))).Do(context.Background())
		mcs.PaymentTxHash = resContract
		fmt.Println("test upload pay:", mcs.PaymentTxHash)
	}
	mcs.client = client
	mcs.params = params
	return mcs
}

func NewMcsMintNft(mcs *MCSUpload) *MCSUpload {
	uploadData := mcs.UploadIpfsData.Data
	NftMetaUrl, _ := mcs.client.NewUploadNftMetadataService().
		SetImageUrl(uploadData.IpfsURL).SetSize(uploadData.FileSize).SetFileName(mcs.FilePath).
		SetWalletAddress(mcs.WalletAddress).SetTxHash(mcs.PaymentTxHash).Do(context.Background())

	NftContractRes, TokenID, _ := mcs.client.NewContractMintNftService().SetNftMetaUrl(NftMetaUrl.Data.IpfsURL).
		SetRpcEndpoint(mcs.RpcEndpoint).SetMintAddress(mcs.params.MintContractAddress).
		SetWalletAddress(mcs.WalletAddress).SetPrivateKey(mcs.PrivateKey).Do(context.Background())

	mintInfo, _ := mcs.client.NewGetMintInfoService().SetMintAddress(mcs.WalletAddress).
		SetPayloadCid(uploadData.PayloadCid).SetTxHash(NftContractRes).SetTokenId(TokenID.String()).
		SetSourceFileUploadId(uploadData.SourceFileUploadID).Do(context.Background())
	mcs.MintInfo = mintInfo
	return mcs
}
