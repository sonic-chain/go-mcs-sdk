package mcs

import (
	"bytes"
	"fmt"
	"go-mcs-sdk/mcs/common"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"unsafe"
)

type BucketClient struct {
	McsClient
}

func NewBucketClient() *BucketClient {
	bucketClient := BucketClient{}
	return &bucketClient
}

func (client *BucketClient) GetBuckets() ([]byte, error) {
	httpRequestUrl := client.BaseURL + common.BUCKET_LIST
	bucketListInfoBytes, err := common.HttpGet(httpRequestUrl, client.JwtToken, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println(*(*string)(unsafe.Pointer(&bucketListInfoBytes)))
	return bucketListInfoBytes, nil
}

func (client *BucketClient) CreateBucket(bucketName string) ([]byte, error) {
	httpRequestUrl := client.BaseURL + common.CREATE_BUCKET
	params := make(map[string]string)
	params["bucket_name"] = bucketName
	response, err := common.HttpPost(httpRequestUrl, client.JwtToken, params)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println(*(*string)(unsafe.Pointer(&response)))
	return response, nil
}

func (client *BucketClient) DeleteBucket(bucketUid string) ([]byte, error) {
	httpRequestUrl := client.BaseURL + common.DELETE_BUCKET + bucketUid
	response, err := common.HttpGet(httpRequestUrl, client.JwtToken, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println(*(*string)(unsafe.Pointer(&response)))
	return response, nil
}

func (client *BucketClient) GetFileInfo(fileId int) ([]byte, error) {
	httpRequestUrl := client.BaseURL + common.FILE_INFO + strconv.Itoa(fileId)
	params := make(map[string]int)
	params["file_id"] = fileId
	response, err := common.HttpGet(httpRequestUrl, client.JwtToken, params)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println(*(*string)(unsafe.Pointer(&response)))
	return response, nil
}

func (client *BucketClient) DeleteFile(fileId int) ([]byte, error) {
	httpRequestUrl := client.BaseURL + common.DELETE_FILE + strconv.Itoa(fileId)
	params := make(map[string]int)
	params["file_id"] = fileId
	response, err := common.HttpGet(httpRequestUrl, client.JwtToken, params)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println(*(*string)(unsafe.Pointer(&response)))
	return response, nil
}

func (client *BucketClient) GetFileList(fileUid, limit, offset string) ([]byte, error) {
	httpRequestUrl := client.BaseURL + common.FILE_LIST + fileUid + "&limit=" + limit + "&offset=" + offset
	response, err := common.HttpGet(httpRequestUrl, client.JwtToken, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println(*(*string)(unsafe.Pointer(&response)))
	return response, nil
}

func (client *BucketClient) CreateFolder(fileName, prefix, bucketUid string) ([]byte, error) {
	httpRequestUrl := client.BaseURL + common.CREATE_FOLDER
	params := make(map[string]string)
	params["file_name"] = fileName
	params["prefix"] = prefix
	params["bucket_uid"] = bucketUid
	response, err := common.HttpPost(httpRequestUrl, client.JwtToken, params)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println(*(*string)(unsafe.Pointer(&response)))
	return response, nil
}

func (client *BucketClient) CheckFile(bucketUid, fileHash, fileName, prefix string) ([]byte, error) {
	httpRequestUrl := client.BaseURL + common.CHECK_UPLOAD
	params := make(map[string]string)
	params["bucket_uid"] = bucketUid
	params["file_hash"] = fileHash
	params["file_name"] = fileName
	params["prefix"] = prefix
	response, err := common.HttpPost(httpRequestUrl, client.JwtToken, params)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println(*(*string)(unsafe.Pointer(&response)))
	return response, nil
}

func (client *BucketClient) UploadChunk(fileHash, uploadFilePath string) ([]byte, error) {
	httpRequestUrl := client.BaseURL + common.UPLOAD_CHUNK
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
	request, err := http.NewRequest("POST", httpRequestUrl, payload)
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
	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println(*(*string)(unsafe.Pointer(&responseBytes)))
	return responseBytes, nil
}

func (client *BucketClient) MergeRequest(bucketUid, fileHash, fileName, prefix string) ([]byte, error) {
	httpRequestUrl := client.BaseURL + common.MERGE_FILE
	params := make(map[string]string)
	params["bucket_uid"] = bucketUid
	params["file_hash"] = fileHash
	params["file_name"] = fileName
	params["prefix"] = prefix
	response, err := common.HttpPost(httpRequestUrl, client.JwtToken, params)
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
