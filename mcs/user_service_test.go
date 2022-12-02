package mcs

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	common2 "go-mcs-sdk/mcs/common"
	"log"
	"testing"
)

func TestUserRegisterService_Do(t *testing.T) {
	p := NewMcsParams("polygon.mumbai")
	client := NewClient(p.McsApi)
	user, _ := client.NewUserRegisterService().SetWalletAddress(WalletAddress).Do(context.Background())
	fmt.Println(user.Data)
}

func TestUserLoginService_Do(t *testing.T) {
	p := NewMcsParams("polygon.mainnet")
	client := NewClient(p.McsApi)
	user, _ := client.NewUserRegisterService().SetWalletAddress(WalletAddress).Do(context.Background())
	nonce := user.Data.Nonce
	fmt.Println(nonce)
	privateKey, err := crypto.HexToECDSA(PrivateKey)
	signature, err := common2.PersonalSign(nonce, privateKey)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(signature)
	jwt, _ := client.NewUserLoginService().SetNetwork(p.ChainName).SetNonce(nonce).SetWalletAddress(WalletAddress).SetSignature(signature).Do(context.Background())
	fmt.Println(jwt.Data.JwtToken)
}
