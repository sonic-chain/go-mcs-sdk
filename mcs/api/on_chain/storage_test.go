package api

import (
	"go-mcs-sdk/mcs/api/common/constants"
	"go-mcs-sdk/mcs/config"
	"testing"

	"go-mcs-sdk/mcs/api/common/logs"

	"github.com/stretchr/testify/assert"
)

func TestUploadFile(t *testing.T) {
	uploadFile, err := onChainClient.UploadFile(config.GetConfig().File2Upload, constants.SOURCE_FILE_TYPE_NORMAL)
	assert.Nil(t, err)
	assert.NotEmpty(t, uploadFile)

	logs.GetLogger().Info(uploadFile)
}

func TestGetDeals(t *testing.T) {
	pageNumber := 1
	pageSize := 10
	dealsParams := DealsParams{
		PageNumber: &pageNumber,
		PageSize:   &pageSize,
	}
	sourceFileUploads, recCnt, err := onChainClient.GetDeals(dealsParams)
	assert.Nil(t, err)
	assert.NotEmpty(t, sourceFileUploads)

	for _, sourceFileUpload := range sourceFileUploads {
		logs.GetLogger().Info(*sourceFileUpload)
	}

	logs.GetLogger().Info(*recCnt)
}

func TestGetDealDetail(t *testing.T) {
	sourceFileUploadDeal, daoSignatures, daoThreshold, err := onChainClient.GetDealDetail(149717, 198335)
	assert.Nil(t, err)
	assert.NotEmpty(t, sourceFileUploadDeal)

	logs.GetLogger().Info(*sourceFileUploadDeal)
	for _, daoSignature := range daoSignatures {
		logs.GetLogger().Info(*daoSignature)
	}
	logs.GetLogger().Info(*daoThreshold)
}

func TestGetDealLogs(t *testing.T) {
	offlineDealLogs, err := onChainClient.GetDealLogs(1)
	assert.Nil(t, err)
	assert.NotEmpty(t, offlineDealLogs)

	for _, offlineDealLog := range offlineDealLogs {
		logs.GetLogger().Info(*offlineDealLog)
	}
}

func TestGetSourceFileUpload(t *testing.T) {
	sourceFileUpload, err := onChainClient.GetSourceFileUpload(148234)
	assert.Nil(t, err)
	assert.NotEmpty(t, sourceFileUpload)

	logs.GetLogger().Info(*sourceFileUpload)
}

func TestUnpinSourceFile(t *testing.T) {
	err := onChainClient.UnpinSourceFile(148234)
	assert.Nil(t, err)
}

func TestWriteNftCollection(t *testing.T) {
	nftCollectionParams := NftCollectionParams{
		Name:   "aaadd",
		TxHash: "0x68c28a439efcb9bbebec7992e0e7bac5d84bd6a06890bf35678f4fdf2ac2e519",
	}

	err := onChainClient.WriteNftCollection(nftCollectionParams)
	assert.Nil(t, err)
	assert.NotEmpty(t, nftCollectionParams)
}

func TestGetNftCollections(t *testing.T) {
	nftCollections, err := onChainClient.GetNftCollections()
	assert.Nil(t, err)
	assert.NotEmpty(t, nftCollections)

	for _, nftCollection := range nftCollections {
		logs.GetLogger().Info(*nftCollection)
	}
}

func TestRecordMintInfo(t *testing.T) {
	name := "abc"
	description := "hello"
	recordMintInfoParams := &RecordMintInfoParams{
		SourceFileUploadId: 148234,
		NftCollectionId:    77,
		TxHash:             "0xesdd",
		TokenId:            5,
		Name:               &name,
		Description:        &description,
	}

	sourceFileMint, err := onChainClient.RecordMintInfo(recordMintInfoParams)
	assert.Nil(t, err)
	assert.NotEmpty(t, sourceFileMint)

	logs.GetLogger().Info(*sourceFileMint)
}

func TestGetMintInfo(t *testing.T) {
	sourceFileMints, err := onChainClient.GetMintInfo(1)
	assert.Nil(t, err)
	assert.NotEmpty(t, sourceFileMints)

	for _, sourceFileMint := range sourceFileMints {
		logs.GetLogger().Info(*sourceFileMint)
	}
}
