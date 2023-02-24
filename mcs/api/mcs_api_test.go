package api

import (
	"log"
	"testing"
	"unsafe"
)

const (
	Nonce              = ""
	Signature          = ""
	SourceFileUploadId = 2123
	FileName           = "4.jpeg"
	Status             = "Pending"
	TokenId            = 111
	DealId             = 10001
	PayLoadCid         = "ewrew"
	txHash             = "fdgdfgdfg"
	MintAddress        = "gfhfghfghf"
	FilePathForUpload  = "/home/userName/Pictures/5.jpeg"
	Apikey             = ""
	AccessToken        = ""
	ValidDays          = 60
	PageNumber         = 1
	PageSize           = 10
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

func TestMcsGetUserTasksDeals(t *testing.T) {
	mcsClient := NewMcsClient()
	err := mcsClient.GetJwtToken()
	if err != nil {
		log.Println(err)
		return
	}
	resp, err := mcsClient.GetUserTasksDeals(FileName, Status, PageNumber, PageSize)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(*(*string)(unsafe.Pointer(&resp)))
}

func TestMcsGetDealDetail(t *testing.T) {
	mcsClient := NewMcsClient()
	err := mcsClient.GetJwtToken()
	if err != nil {
		log.Println(err)
		return
	}
	resp, err := mcsClient.GetDealDetail(SourceFileUploadId, DealId)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(*(*string)(unsafe.Pointer(&resp)))
}

func TestMcsGetMintInfo(t *testing.T) {
	mcsClient := NewMcsClient()
	err := mcsClient.GetJwtToken()
	if err != nil {
		log.Println(err)
		return
	}
	resp, err := mcsClient.GetMintInfo(SourceFileUploadId, TokenId, PayLoadCid, txHash, MintAddress)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(*(*string)(unsafe.Pointer(&resp)))
}

func TestMcsUploadFile(t *testing.T) {
	mcsClient := NewMcsClient()
	err := mcsClient.GetJwtToken()
	if err != nil {
		log.Println(err)
		return
	}
	resp, err := mcsClient.UploadFile(FilePathForUpload)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(*(*string)(unsafe.Pointer(&resp)))
}

func TestMcsGenerateApikey(t *testing.T) {
	mcsClient := NewMcsClient()
	err := mcsClient.GetJwtToken()
	if err != nil {
		log.Println(err)
		return
	}
	resp, err := mcsClient.GenerateApikey(ValidDays)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(*(*string)(unsafe.Pointer(&resp)))
}
