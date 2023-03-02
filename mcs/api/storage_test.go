package api

import (
	"go-mcs-sdk/mcs/common/constants"
	"go-mcs-sdk/mcs/config"
	"log"
	"testing"

	"github.com/filswan/go-swan-lib/logs"
)

func TestMcsUploadFile(t *testing.T) {
	mcsClient, err := GetMcsClient()
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	uploadFile, err := mcsClient.UploadFile(config.GetConfig().File2Upload, constants.SOURCE_FILE_TYPE_NORMAL)
	if err != nil {
		log.Println(err)
		return
	}

	logs.GetLogger().Info(uploadFile)
}

func TestGetDeals(t *testing.T) {
	mcsClient, err := GetMcsClient()
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	pageNumber := 1
	pageSize := 10
	dealsParams := DealsParams{
		PageNumber: &pageNumber,
		PageSize:   &pageSize,
	}
	sourceFileUploads, recCnt, err := mcsClient.GetDeals(dealsParams)
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	for _, sourceFileUpload := range sourceFileUploads {
		logs.GetLogger().Info(*sourceFileUpload)
	}

	logs.GetLogger().Info(*recCnt)
}

func TestMcsGetDealDetail(t *testing.T) {
	mcsClient, err := GetMcsClient()
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	sourceFileUploadDeal, daoSignatures, daoThreshold, err := mcsClient.GetDealDetail(149717, 198335)
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	logs.GetLogger().Info(*sourceFileUploadDeal)
	for _, daoSignature := range daoSignatures {
		logs.GetLogger().Info(*daoSignature)
	}
	logs.GetLogger().Info(*daoThreshold)
}

func TestMcsGetDealLog(t *testing.T) {
	mcsClient, err := GetMcsClient()
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	offlineDealLogs, err := mcsClient.GetDealLogs(1)
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	for _, offlineDealLog := range offlineDealLogs {
		logs.GetLogger().Info(*offlineDealLog)
	}
}

func TestMcsGetSourceFileUpload(t *testing.T) {
	mcsClient, err := GetMcsClient()
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	sourceFileUpload, err := mcsClient.GetSourceFileUpload(148234)
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	logs.GetLogger().Info(*sourceFileUpload)
}

func TestUnpinSourceFile(t *testing.T) {
	mcsClient, err := GetMcsClient()
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	err = mcsClient.UnpinSourceFile(148234)
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}
}

func TestWriteNftCollection(t *testing.T) {
	mcsClient, err := GetMcsClient()
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	nftCollectionParams := NftCollectionParams{
		Name:   "aaadd",
		TxHash: "0x68c28a439efcb9bbebec7992e0e7bac5d84bd6a06890bf35678f4fdf2ac2e519",
	}

	err = mcsClient.WriteNftCollection(nftCollectionParams)
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}
}

func TestGetNftCollections(t *testing.T) {
	mcsClient, err := GetMcsClient()
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	nftCollections, err := mcsClient.GetNftCollections()
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	for _, nftCollection := range nftCollections {
		logs.GetLogger().Info(*nftCollection)
	}
}

func TestRecordMintInfo(t *testing.T) {
	mcsClient, err := GetMcsClient()
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	name := "abc"
	description := "hello"
	recordMintInfoParams := &RecordMintInfoParams{
		SourceFileIploadId: 1,
		NftCollectionId:    52,
		TxHash:             "0xesdd",
		TokenId:            5,
		Name:               &name,
		Description:        &description,
	}

	sourceFileMint, err := mcsClient.RecordMintInfo(recordMintInfoParams)
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	logs.GetLogger().Info(*sourceFileMint)
}

func TestGetMintInfo(t *testing.T) {
	mcsClient, err := GetMcsClient()
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	sourceFileMints, err := mcsClient.GetMintInfo(1)
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	for _, sourceFileMint := range sourceFileMints {
		logs.GetLogger().Info(*sourceFileMint)
	}
}
