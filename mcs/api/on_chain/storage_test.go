package api

import (
	"testing"

	"github.com/filswan/go-mcs-sdk/mcs/api/common/constants"

	"github.com/filswan/go-mcs-sdk/mcs/api/common/logs"

	"github.com/stretchr/testify/assert"
)

func TestUpload(t *testing.T) {
	uploadFile, err := onChainClient.Upload("", constants.SOURCE_FILE_TYPE_NORMAL)
	assert.Nil(t, err)
	assert.NotEmpty(t, uploadFile)

	logs.GetLogger().Info(uploadFile)
}

func TestGetUserTaskDeals(t *testing.T) {
	pageNumber := 1
	pageSize := 10
	dealsParams := DealsParams{
		PageNumber: &pageNumber,
		PageSize:   &pageSize,
	}
	sourceFileUploads, recCnt, err := onChainClient.GetUserTaskDeals(dealsParams)
	assert.Nil(t, err)
	assert.NotNil(t, sourceFileUploads)
	assert.NotNil(t, recCnt)
	assert.GreaterOrEqual(t, *recCnt, int64(0))

	for _, sourceFileUpload := range sourceFileUploads {
		logs.GetLogger().Info(*sourceFileUpload)
	}

	logs.GetLogger().Info(*recCnt)
}

func TestGetDealDetail(t *testing.T) {
	sourceFileUploadDeal, daoSignatures, daoThreshold, err := onChainClient.GetDealDetail(149717, 198335)
	assert.Nil(t, err)
	assert.NotNil(t, sourceFileUploadDeal)
	assert.NotNil(t, daoSignatures)
	assert.NotNil(t, daoThreshold)
	assert.Greater(t, *daoThreshold, 0)

	logs.GetLogger().Info(*sourceFileUploadDeal)
	for _, daoSignature := range daoSignatures {
		logs.GetLogger().Info(*daoSignature)
	}
	logs.GetLogger().Info(*daoThreshold)
}

func TestGetDealLogs(t *testing.T) {
	offlineDealLogs, err := onChainClient.GetDealLogs(1)
	assert.Nil(t, err)
	assert.NotNil(t, offlineDealLogs)

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
}

func TestGetNftCollections(t *testing.T) {
	nftCollections, err := onChainClient.GetNftCollections()
	assert.Nil(t, err)
	assert.NotNil(t, nftCollections)

	for _, nftCollection := range nftCollections {
		logs.GetLogger().Info(*nftCollection)
	}
}

func TestRecordMintInfo(t *testing.T) {
	name := "abc"
	description := "hello"
	recordMintInfoParams := &RecordMintInfoParams{
		SourceFileUploadId: 151353,
		NftCollectionId:    77,
		TxHash:             "0xesdd",
		TokenId:            5,
		Name:               &name,
		Description:        &description,
	}

	sourceFileMint, err := onChainClient.RecordMintInfo(recordMintInfoParams)
	assert.Nil(t, err)
	assert.NotNil(t, sourceFileMint)

	logs.GetLogger().Info(*sourceFileMint)
}

func TestGetMintInfo(t *testing.T) {
	sourceFileMints, err := onChainClient.GetMintInfo(151353)
	assert.Nil(t, err)
	assert.NotNil(t, sourceFileMints)

	for _, sourceFileMint := range sourceFileMints {
		logs.GetLogger().Info(*sourceFileMint)
	}
}
