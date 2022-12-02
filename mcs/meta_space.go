package mcs

import (
	"context"
	"github.com/ethereum/go-ethereum/crypto"
	"go-mcs-sdk/mcs/common"
	"log"
)

const (
	McsBackendBaseUrl               = "http://192.168.199.61:8889/api/"
	UserWalletAddressForRegisterMcs = "0x7d2C017e20Ee3D86047727197094fCD197656168"
	UserWalletAddressPK             = "9197b7d31cb4548aa4bba82d3a15bdf9f35814d130e9077b4b0ed8a7235addbe"
	ChainNameForRegisterOnMcs       = "polygon.mumbai"
)

type MetaSpaceClient struct {
	MetaSpaceUrl string `json:"meta_space_url"`
	JwtToken     string `json:"jwt_token"`
}

func NewMetaSpaceClient(metaSpaceUrl string) *MetaSpaceClient {
	metaSpaceClient := MetaSpaceClient{
		MetaSpaceUrl: metaSpaceUrl,
	}
	return &metaSpaceClient
}

func (client *MetaSpaceClient) SetJwtToken(jwtToken string) *MetaSpaceClient {
	client.JwtToken = jwtToken
	return client
}

func (client *MetaSpaceClient) GetToken() error {
	mcsClient := NewClient(McsBackendBaseUrl)
	user, err := mcsClient.NewUserRegisterService().SetWalletAddress(UserWalletAddressForRegisterMcs).Do(context.Background())
	if err != nil {
		log.Println(err)
		return err
	}
	nonce := user.Data.Nonce
	privateKey, _ := crypto.HexToECDSA(UserWalletAddressPK)
	signature, _ := common.PersonalSign(nonce, privateKey)
	jwt, err := mcsClient.NewUserLoginService().SetNetwork(ChainNameForRegisterOnMcs).SetNonce(nonce).SetWalletAddress(UserWalletAddressForRegisterMcs).
		SetSignature(signature).Do(context.Background())
	if err != nil {
		log.Println(err)
		return err
	}
	client.SetJwtToken(jwt.Data.JwtToken)
	return nil
}

func (client *MetaSpaceClient) GetBuckets() error {
	httpRequestUrl := client.MetaSpaceUrl + common.DIRECTORY
	bytes, err := common.HttpGet(httpRequestUrl, client.JwtToken, nil)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println(bytes)
	return err
}
