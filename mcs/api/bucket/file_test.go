package bucket

import (
	"fmt"
	"os"
	"testing"

	"go-mcs-sdk/mcs/api/common/logs"

	"github.com/stretchr/testify/assert"
)

func TestGetFile(t *testing.T) {
	ossFile, err := buketClient.GetFile("aaa", "duration6")
	assert.Nil(t, err)
	assert.NotEmpty(t, ossFile)

	logs.GetLogger().Info(*ossFile)
}

func TestCreateFolder(t *testing.T) {
	folderName, err := buketClient.CreateFolder("aaa", "test", "")
	assert.Nil(t, err)
	assert.NotEmpty(t, folderName)

	logs.GetLogger().Info(*folderName)
}

func TestDeleteFile(t *testing.T) {
	err := buketClient.DeleteFile("aaa", "duration6")
	assert.Nil(t, err)
}

func TestListFiles(t *testing.T) {
	ossFiles, count, err := buketClient.ListFiles("aaa", "", 10, 0)
	assert.Nil(t, err)
	assert.NotNil(t, ossFiles)
	assert.NotNil(t, count)

	for _, ossFile := range ossFiles {
		logs.GetLogger().Info(*ossFile)
	}

	logs.GetLogger().Info(*count)
}

func UploadFile(t *testing.T) {
	err := buketClient.UploadFile("aaa", "test/duration22", file2Upload, true)
	assert.Nil(t, err)
}

func TestGetFileInfo(t *testing.T) {
	fileInfo, err := buketClient.GetFileInfo(6674)
	assert.Nil(t, err)
	assert.NotEmpty(t, fileInfo)

	logs.GetLogger().Info(*fileInfo)
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
