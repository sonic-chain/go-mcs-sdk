package mcs

import (
	"log"
	"testing"
	"unsafe"
)

const (
	Nonce              = "12955819538690153468153899560298852982"
	Signature          = "0xa77366b42c8d7691d2ec69455897cf8caf502adc319e8d0a2aae587d1e746ba27e29055ca770b8c5d40094165fdfd178cc380cae12536fc7df9182f2ff00133d1b"
	SourceFileUploadId = 2123
)

func TestMcsGetJwtToken(t *testing.T) {
	mcsClient := NewMcsClient()
	resp, err := mcsClient.UserLogin(mcsClient.UserWalletAddressForRegisterMcs, Signature, Nonce, mcsClient.ChainNameForRegisterOnMcs)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(*(*string)(unsafe.Pointer(&resp)))
}

func TestMcsUserRegister(t *testing.T) {
	mcsClient := NewMcsClient()
	nonce, err := mcsClient.UserRegister(mcsClient.UserWalletAddressForRegisterMcs)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(*nonce)
}

func TestMcsGetJwtToken2(t *testing.T) {
	mcsClient := NewMcsClient()
	err := mcsClient.GetJwtToken()
	if err != nil {
		log.Println(err)
		return
	}
}

func TestMcsGetParams(t *testing.T) {
	mcsClient := NewMcsClient()
	resp, err := mcsClient.GetParams()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(*(*string)(unsafe.Pointer(&resp)))
}

func TestMcsGetPriceRate(t *testing.T) {
	mcsClient := NewMcsClient()
	err := mcsClient.GetJwtToken()
	if err != nil {
		log.Println(err)
		return
	}
	resp, err := mcsClient.GetPriceRate()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(*(*string)(unsafe.Pointer(&resp)))
}

func TestMcsGetPaymentInfo(t *testing.T) {
	mcsClient := NewMcsClient()
	err := mcsClient.GetJwtToken()
	if err != nil {
		log.Println(err)
		return
	}
	resp, err := mcsClient.GetPaymentInfo(SourceFileUploadId)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(*(*string)(unsafe.Pointer(&resp)))
}
