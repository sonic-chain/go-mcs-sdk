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
