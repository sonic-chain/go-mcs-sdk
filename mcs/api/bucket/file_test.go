package bucket

import (
	"testing"

	"github.com/filswan/go-swan-lib/logs"
)

func TestGetFileInfo(t *testing.T) {
	fileInfo, err := onChainClient.GetFileInfo(6590)
	if err != nil {
		logs.GetLogger().Fatal(err)
	}

	logs.GetLogger().Info(*fileInfo)
}

func TestDeleteFile(t *testing.T) {
	err := onChainClient.DeleteFile(6591)
	if err != nil {
		logs.GetLogger().Fatal(err)
	}
}

func TestCreateFolder(t *testing.T) {
	folderName, err := onChainClient.CreateFolder("ddsfads", "", "a7303d2a-acd2-48ac-a062-8454bbf148d2")
	if err != nil {
		logs.GetLogger().Fatal(err)
	}

	logs.GetLogger().Info(*folderName)
}

func TestGetFileInfoByObjectName(t *testing.T) {
	folderName, err := onChainClient.GetFileInfoByObjectName("ddsfads/duration7", "a7303d2a-acd2-48ac-a062-8454bbf148d2")
	if err != nil {
		logs.GetLogger().Fatal(err)
	}

	logs.GetLogger().Info(*folderName)
}

func TestCheckFile(t *testing.T) {
	fileInfo, err := onChainClient.CheckFile("a7303d2a-acd2-48ac-a062-8454bbf148d2", "ddsfads", "0c3ec30ad80e40d02d66d681a9ba24c4", "duration7")
	if err != nil {
		logs.GetLogger().Fatal(err)
	}

	logs.GetLogger().Info(*fileInfo)
}

func TestUploadFileChunk(t *testing.T) {
	err := onChainClient.UploadFile("abc", "ddd/test65", "/Users/dorachen/work/test4", true)
	if err != nil {
		logs.GetLogger().Fatal(err)
	}
}
