package api

import (
	"bytes"
	"fmt"
	"go-mcs-sdk/mcs/api/common/constants"
	"go-mcs-sdk/mcs/common"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path"
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

func (client *McsClient) UploadChunk(fileHash, uploadFilePath string) ([]byte, error) {
	apiUrl := libutils.UrlJoin(client.BaseUrl, constants.UPLOAD_CHUNK)
	fileNameWithSuffix := path.Base(uploadFilePath)
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, err := os.Open(uploadFilePath)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	part1, err := writer.CreateFormFile("file", fileNameWithSuffix)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	_, err = io.Copy(part1, file)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	err = writer.WriteField("hash", fileHash)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	err = writer.Close()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	request, err := http.NewRequest("POST", apiUrl, payload)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", client.JwtToken))
	request.Header.Add("Content-Type", writer.FormDataContentType())
	response, err := http.DefaultClient.Do(request)
	//response, err := httpClient.Post(httpRequestUrl,bodyWriter.FormDataContentType(),bodyBuffer)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer response.Body.Close()
	responseBytes, err := io.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println(*(*string)(unsafe.Pointer(&responseBytes)))
	return responseBytes, nil
}

func (client *McsClient) MergeRequest(bucketUid, fileHash, fileName, prefix string) ([]byte, error) {
	apiUrl := libutils.UrlJoin(client.BaseUrl, constants.MERGE_FILE)
	params := make(map[string]string)
	params["bucket_uid"] = bucketUid
	params["file_hash"] = fileHash
	params["file_name"] = fileName
	params["prefix"] = prefix
	response, err := common.HttpPost(apiUrl, client.JwtToken, params)
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
