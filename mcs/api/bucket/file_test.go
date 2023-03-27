package bucket

import (
	"fmt"
	"os"
	"testing"

	"go-mcs-sdk/mcs/api/common/logs"
	"go-mcs-sdk/mcs/config"

	"github.com/stretchr/testify/assert"
)

func TestGetFile(t *testing.T) {
	ossFile, err := buketClient.GetFile("test/duration22", "0ef9c94d-9bb9-4ce9-b687-7db732a9ce2e")
	assert.Nil(t, err)
	assert.NotEmpty(t, ossFile)

	logs.GetLogger().Info(*ossFile)
}

func TestCreateFolder(t *testing.T) {
	folderName, err := buketClient.CreateFolder("test", "", "0ef9c94d-9bb9-4ce9-b687-7db732a9ce2e")
	assert.Nil(t, err)
	assert.NotEmpty(t, folderName)

	logs.GetLogger().Info(*folderName)
}

func TestUploadFileChunk(t *testing.T) {
	err := buketClient.UploadFile("abc", "test/duration22", config.GetConfig().File2Upload, true)
	assert.Nil(t, err)
}

func TestGetFileInfo(t *testing.T) {
	fileInfo, err := buketClient.GetFileInfo(6674)
	assert.Nil(t, err)
	assert.NotEmpty(t, fileInfo)

	logs.GetLogger().Info(*fileInfo)
}

func TestGetFileByName(t *testing.T) {
	ossFile, err := buketClient.GetFileByName("aaa", "abc")
	assert.Nil(t, err)
	assert.NotEmpty(t, ossFile)

	logs.GetLogger().Info(*ossFile)
}

func TestDeleteFile(t *testing.T) {
	err := buketClient.DeleteFile(6674)
	assert.Nil(t, err)
}

func TestPinFiles2Ipfs(t *testing.T) {
	ossFile, err := buketClient.UploadIpfsFolder("abc", "aaa", "/Users/dorachen/work/test2")
	assert.Nil(t, err)
	assert.NotEmpty(t, ossFile)

	logs.GetLogger().Info(*ossFile)
}

func TestDownloadFile(t *testing.T) {
	path, err := os.Getwd()
	if err != nil {
		logs.GetLogger().Fatal(err)
	}
	fmt.Println(path)
	err = buketClient.DownloadFile("abc", "aaa", path)
	assert.Nil(t, err)
}
