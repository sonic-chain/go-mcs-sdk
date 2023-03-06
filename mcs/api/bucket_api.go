package api

import (
	"fmt"
	"go-mcs-sdk/mcs/api/common/constants"
	"go-mcs-sdk/mcs/common"
	"log"
	"path/filepath"
	"regexp"
	"unsafe"

	libutils "github.com/filswan/go-swan-lib/utils"
)

type BucketClient struct {
	McsClient
}

func NewBucketClient() *BucketClient {
	bucketClient := BucketClient{}
	return &bucketClient
}

func (client *McsClient) GetFileList(fileUid, limit, offset string) ([]byte, error) {
	apiUrl := libutils.UrlJoin(client.BaseUrl, constants.FILE_LIST) + fileUid + "&limit=" + limit + "&offset=" + offset
	response, err := common.HttpGet(apiUrl, client.JwtToken, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println(*(*string)(unsafe.Pointer(&response)))
	return response, nil
}

func CheckSpecialChar(fileName, filePath string) error {
	if SpecialChar(fileName) {
		err := fmt.Errorf("file alias has special characters")
		log.Println(err)
		return err
	}
	_, fileNameInPath := filepath.Split(filePath)
	if SpecialChar(fileNameInPath) {
		err := fmt.Errorf("file name in path has special characters")
		log.Println(err)
		return err
	}
	return nil
}

func SpecialChar(line string) bool {
	specialCharacters := "[!@#$%^&*()-+?=,<>/'\" \n\t\v\f\r]"
	reg := regexp.MustCompile(specialCharacters)
	matchStr := reg.FindAllString(line, -1)
	if len(matchStr) > 0 {
		return true
	} else {
		return false
	}
}
