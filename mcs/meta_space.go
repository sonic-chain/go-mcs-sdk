package mcs

import (
	"context"
	"github.com/ethereum/go-ethereum/crypto"
	"go-mcs-sdk/mcs/common"
	"log"
)

const (
	MetaSpaceBackendBaseUrl         = ""
	UserWalletAddressForRegisterMcs = ""
	UserWalletAddressPK             = ""
	ChainNameForRegisterOnMcs       = ""
)

type MetaSpaceClient struct {
	MetaSpaceUrl string `json:"meta_space_url"`
	JwtToken     string `json:"jwt_token"`
}

func (client *MetaSpaceClient) NewMetaSpaceClient(metaSpaceUrl string) *MetaSpaceClient {
	metaSpaceClient := MetaSpaceClient{
		MetaSpaceUrl: metaSpaceUrl,
	}
	return &metaSpaceClient
}

func (client *MetaSpaceClient) SetJwtToken(jwtToken string) *MetaSpaceClient {
	client.SetJwtToken(jwtToken)
	return client
}

func (client *MetaSpaceClient) GetToken() error {
	mcsClient := NewClient(MetaSpaceBackendBaseUrl)
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

func (client *MetaSpaceClient) GetBuckets() {
	httpRequestUrl := client.MetaSpaceUrl + common.DIRECTORY
	err := client.NewMetaSpaceClient(MetaSpaceBackendBaseUrl).GetToken()
	if err != nil {
		log.Println(err)
	}
	common.HttpGet(httpRequestUrl, client.JwtToken, nil)
}
