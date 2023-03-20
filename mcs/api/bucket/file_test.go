package bucket

import (
	"testing"

	"go-mcs-sdk/mcs/api/common/logs"
)

func TestGetFileInfo(t *testing.T) {
	fileInfo, err := buketClient.GetFileInfo(6590)
	if err != nil {
		logs.GetLogger().Fatal(err)
	}

	logs.GetLogger().Info(*fileInfo)
}

func TestDeleteFile(t *testing.T) {
	err := buketClient.DeleteFile(6591)
	if err != nil {
		logs.GetLogger().Fatal(err)
	}
}

func TestCreateFolder(t *testing.T) {
	folderName, err := buketClient.CreateFolder("ddsfads", "", "a7303d2a-acd2-48ac-a062-8454bbf148d2")
	if err != nil {
		logs.GetLogger().Fatal(err)
	}

	logs.GetLogger().Info(*folderName)
}

func TestGetFileInfoByObjectName(t *testing.T) {
	folderName, err := buketClient.GetFileInfoByObjectName("ddsfads/duration7", "a7303d2a-acd2-48ac-a062-8454bbf148d2")
	if err != nil {
		logs.GetLogger().Fatal(err)
	}

	logs.GetLogger().Info(*folderName)
}

func TestCheckFile(t *testing.T) {
	fileInfo, err := buketClient.CheckFile("a7303d2a-acd2-48ac-a062-8454bbf148d2", "ddsfads", "0c3ec30ad80e40d02d66d681a9ba24c4", "duration7")
	if err != nil {
		logs.GetLogger().Fatal(err)
	}

	logs.GetLogger().Info(*fileInfo)
}

func TestUploadFileChunk(t *testing.T) {
	err := buketClient.UploadFile("abc", "ddd/duration21", "/Users/dorachen/work/duration11", true)
	if err != nil {
		logs.GetLogger().Fatal(err)
	}
}

func TestPinFiles2Ipfs(t *testing.T) {
	ossFile, err := buketClient.PinFiles2Ipfs("abc", "eee", "/Users/dorachen/work/test2")
	if err != nil {
		logs.GetLogger().Fatal(err)
	}

	logs.GetLogger().Info(*ossFile)
}
