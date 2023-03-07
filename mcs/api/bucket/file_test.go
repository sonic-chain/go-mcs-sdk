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
