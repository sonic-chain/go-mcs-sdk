package bucket

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-mcs-sdk/mcs/api/common/constants"
	"go-mcs-sdk/mcs/api/common/web"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"

	"github.com/codingsince1985/checksum"
	"github.com/filswan/go-swan-lib/logs"
	libutils "github.com/filswan/go-swan-lib/utils"
	"github.com/jinzhu/gorm"
)

type OssFile struct {
	Name       string `json:"name"`
	Address    string `json:"address"`
	Prefix     string `json:"prefix"`
	BucketUid  string `json:"bucket_uid"`
	FileHash   string `json:"file_hash"`
	Size       int64  `json:"size"`
	PayloadCid string `json:"payload_cid"`
	PinStatus  string `json:"pin_status"`
	IsDeleted  bool   `json:"is_deleted"`
	IsFolder   bool   `json:"is_folder"`
	ObjectName string `json:"object_name"`
	Type       int    `json:"type"`
	gorm.Model
}

func (bucketClient *BucketClient) GetFileInfo(fileId int) (*OssFile, error) {
	apiUrl := libutils.UrlJoin(bucketClient.BaseUrl, constants.API_URL_BUCKET_FILE_GET_FILE_INFO)
	apiUrl = apiUrl + "?file_id=" + strconv.Itoa(fileId)

	var fileInfo OssFile
	err := web.HttpGet(apiUrl, bucketClient.JwtToken, nil, &fileInfo)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return &fileInfo, nil
}

func (bucketClient *BucketClient) DeleteFile(fileId int) error {
	apiUrl := libutils.UrlJoin(bucketClient.BaseUrl, constants.API_URL_BUCKET_FILE_DELETE_FILE)
	apiUrl = apiUrl + "?file_id=" + strconv.Itoa(fileId)

	err := web.HttpGet(apiUrl, bucketClient.JwtToken, nil, nil)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (bucketClient *BucketClient) CreateFolder(fileName, prefix, bucketUid string) (*string, error) {
	apiUrl := libutils.UrlJoin(bucketClient.BaseUrl, constants.API_URL_BUCKET_FILE_CREATE_FOLDER)

	var params struct {
		FileName  string `json:"file_name"`
		Prefix    string `json:"prefix"`
		BucketUid string `json:"bucket_uid"`
	}

	params.FileName = fileName
	params.Prefix = prefix
	params.BucketUid = bucketUid

	var folderName string
	err := web.HttpPost(apiUrl, bucketClient.JwtToken, &params, &folderName)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &folderName, nil
}

func (bucketClient *BucketClient) GetFileInfoByObjectName(objectName, bucketUid string) (*OssFile, error) {
	apiUrl := libutils.UrlJoin(bucketClient.BaseUrl, constants.API_URL_BUCKET_FILE_GET_FILE_INFO_BY_OBJECT_NAME)
	apiUrl = apiUrl + "?bucket_uid=" + bucketUid + "&object_name=" + objectName

	var fileInfo OssFile
	err := web.HttpGet(apiUrl, bucketClient.JwtToken, nil, &fileInfo)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return &fileInfo, nil
}

func getPrefixFileName(objectName string) (string, string) {
	lastIndex := strings.LastIndex(objectName, "/")

	if lastIndex == -1 {
		return "", objectName
	}

	prefix := objectName[0:lastIndex]
	fileName := objectName[lastIndex+1:]

	return prefix, fileName
}

func (bucketClient *BucketClient) getBucketUid(bucketName string) (*string, error) {
	buckets, err := bucketClient.GetBuckets()
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	for _, bucket := range buckets {
		if bucket.BucketName == bucketName {
			return &bucket.BucketUid, nil
		}
	}

	return nil, nil
}

func (bucketClient *BucketClient) UploadFile(bucketName, objectName, filePath string, replace bool) error {
	prefix, fileName := getPrefixFileName(objectName)

	bucketUid, err := bucketClient.getBucketUid(bucketName)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	if bucketUid == nil {
		err := fmt.Errorf("bucket:%s not exists", bucketName)
		logs.GetLogger().Error(err)
		return err
	}

	osFileInfo, err := os.Stat(filePath)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	fileSize := osFileInfo.Size()

	fileHashMd5, err := checksum.MD5sum(filePath)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	ossFileInfo, err := bucketClient.CheckFile(*bucketUid, prefix, fileHashMd5, fileName)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	if ossFileInfo.FileIsExist && replace {
		fileId := ossFileInfo.FileId
		if fileId == 0 {
			ossFile, err := bucketClient.GetFileInfoByObjectName(objectName, *bucketUid)
			if err != nil {
				logs.GetLogger().Error(err)
				return err
			}

			fileId = ossFile.ID
		}
		err = bucketClient.DeleteFile(int(fileId))
		if err != nil {
			logs.GetLogger().Error(err)
			return err
		}
		ossFileInfo, err = bucketClient.CheckFile(*bucketUid, prefix, fileHashMd5, fileName)
		if err != nil {
			logs.GetLogger().Error(err)
			return err
		}
	}

	if !ossFileInfo.FileIsExist {
		if !ossFileInfo.IpfsIsExist {
			file, err := os.Open(filePath)
			if err != nil {
				logs.GetLogger().Error(err)
				return err
			}
			bytesReadTotal := int64(0)
			chunNo := 0

			var wg sync.WaitGroup

			for bytesReadTotal < fileSize {
				var chunkSize int64
				bytesLeft := fileSize - bytesReadTotal
				if bytesLeft >= constants.FILE_CHUNK_SIZE_MAX {
					chunkSize = constants.FILE_CHUNK_SIZE_MAX
				} else {
					chunkSize = bytesLeft
				}
				chunk := make([]byte, chunkSize)
				_, err := file.ReadAt(chunk, bytesReadTotal)
				if err != nil {
					logs.GetLogger().Error(err)
					return err
				}
				bytesReadTotal = bytesReadTotal + chunkSize
				chunNo = chunNo + 1

				partFileName := strconv.Itoa(chunNo) + "_" + fileName

				wg.Add(1)
				go func() {
					logs.GetLogger().Info("file name:", partFileName, ", chunk size:", chunkSize)
					_, err = bucketClient.UploadFileChunk(fileHashMd5, partFileName, chunk)
					if err != nil {
						logs.GetLogger().Error(err)
					}
					wg.Done()
				}()
			}

			wg.Wait()
			_, err = bucketClient.MergeFile(*bucketUid, fileHashMd5, fileName, prefix)
			if err != nil {
				logs.GetLogger().Error(err)
				return err
			}
		}
	}

	return nil
}

type OssFileInfo struct {
	FileId      uint   `form:"file_id" json:"file_id"`
	FileHash    string `form:"file_hash" json:"file_hash"`
	FileIsExist bool   `form:"file_is_exist" json:"file_is_exist"`
	IpfsIsExist bool   `form:"ipfs_is_exist" json:"ipfs_is_exist"`
	Size        int64  `form:"size" json:"size"`
	PayloadCid  string `form:"payload_cid" json:"payload_cid"`
	//IpfsUrl     string `form:"ipfs_url" json:"ipfs_url"`
}

func (bucketClient *BucketClient) CheckFile(bucketUid, prefix, fileHash, fileName string) (*OssFileInfo, error) {
	apiUrl := libutils.UrlJoin(bucketClient.BaseUrl, constants.API_URL_BUCKET_FILE_CHECK_UPLOAD)

	var params struct {
		FileName  string `json:"file_name"`
		FileHash  string `json:"file_hash"`
		Prefix    string `json:"prefix"`
		BucketUid string `json:"bucket_uid"`
	}

	params.FileName = fileName
	params.FileHash = fileHash
	params.Prefix = prefix
	params.BucketUid = bucketUid

	var ossFileInfo OssFileInfo
	err := web.HttpPost(apiUrl, bucketClient.JwtToken, &params, &ossFileInfo)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return &ossFileInfo, nil
}

func (bucketClient *BucketClient) UploadFileChunk(fileHash, fileName string, chunk []byte) ([]string, error) {
	apiUrl := libutils.UrlJoin(bucketClient.BaseUrl, constants.API_URL_BUCKET_FILE_UPLOAD_CHUNK)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", fileName)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	chunkReader := bytes.NewReader(chunk)

	//chunkReader.WriteTo(part)

	_, err = io.Copy(part, chunkReader)
	//n, err := part.Write(chunk)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	err = writer.WriteField("hash", fileHash)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	err = writer.Close()
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	request, err := http.NewRequest("POST", apiUrl, body)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", bucketClient.JwtToken))
	request.Header.Add("Content-Type", writer.FormDataContentType())
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		err := fmt.Errorf("http status: %s, code:%d, url:%s", response.Status, response.StatusCode, apiUrl)
		logs.GetLogger().Error(err)
		switch response.StatusCode {
		case http.StatusNotFound:
			logs.GetLogger().Error("please check your url:", apiUrl)
		case http.StatusUnauthorized:
			logs.GetLogger().Error("Please check your token:", bucketClient.JwtToken)
		}
	}

	responseBytes, err := io.ReadAll(response.Body)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	var mcsResponse struct {
		Status  string   `json:"status"`
		Message string   `json:"message"`
		Data    []string `json:"data"`
	}

	err = json.Unmarshal(responseBytes, &mcsResponse)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	if !strings.EqualFold(mcsResponse.Status, constants.HTTP_STATUS_SUCCESS) {
		err := fmt.Errorf("%s failed, status:%s, message:%s", apiUrl, mcsResponse.Status, mcsResponse.Message)
		logs.GetLogger().Error(err)
		return nil, err
	}

	return mcsResponse.Data, nil
}

func (bucketClient *BucketClient) MergeFile(bucketUid, fileHash, fileName, prefix string) (*OssFileInfo, error) {
	apiUrl := libutils.UrlJoin(bucketClient.BaseUrl, constants.API_URL_BUCKET_FILE_MERGE_FILE)

	var params struct {
		FileName  string `json:"file_name"`
		FileHash  string `json:"file_hash"`
		Prefix    string `json:"prefix"`
		BucketUid string `json:"bucket_uid"`
	}

	params.FileName = fileName
	params.FileHash = fileHash
	params.Prefix = prefix
	params.BucketUid = bucketUid

	var ossFileInfo OssFileInfo
	err := web.HttpPostTimeout(apiUrl, bucketClient.JwtToken, &params, 600, &ossFileInfo)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return &ossFileInfo, nil
}

func (bucketClient *BucketClient) GetFileList(fileUid, limit, offset string) ([]*OssFile, error) {
	apiUrl := libutils.UrlJoin(bucketClient.BaseUrl, constants.API_URL_BUCKET_FILE_GET_FILE_LIST) + fileUid + "&limit=" + limit + "&offset=" + offset

	var files []*OssFile
	err := web.HttpGet(apiUrl, bucketClient.JwtToken, nil, &files)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return files, nil
}

type PinFiles2IpfsResponse struct {
	Status  string  `json:"status"`
	Message string  `json:"message"`
	Data    OssFile `json:"data"`
}

func (bucketClient *BucketClient) PinFiles2Ipfs(bucketName, objectName, folderPath string) (*UploadFile, error) {
	folderName := filepath.Base(objectName)
	if strings.Trim(folderName, " ") == "" {
		folderName = filepath.Base(folderPath)
	}

	prefix := filepath.Dir(objectName)

	bucketUid, err := bucketClient.getBucketUid(bucketName)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)

	err = writer.WriteField("folder_name", folderName)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	err = writer.WriteField("prefix", prefix)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	err = writer.WriteField("bucket_uid", *bucketUid)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	file, err := os.Open(filePath)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	defer file.Close()

	part1, err := writer.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	_, err = io.Copy(part1, file)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	err = writer.Close()
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	apiUrl := libutils.UrlJoin(bucketClient.BaseUrl, constants.API_URL_BUCKET_FILE_PIN_FILES_2_IPFS)
	httpClient := &http.Client{}
	req, err := http.NewRequest("POST", apiUrl, payload)

	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", bucketClient.JwtToken))
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := httpClient.Do(req)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	var pinFiles2IpfsResponse PinFiles2IpfsResponse
	err = json.Unmarshal(body, &pinFiles2IpfsResponse)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	if !strings.EqualFold(pinFiles2IpfsResponse.Status, constants.HTTP_STATUS_SUCCESS) {
		err := fmt.Errorf("get parameters failed, status:%s,message:%s", pinFiles2IpfsResponse.Status, pinFiles2IpfsResponse.Message)
		logs.GetLogger().Error(err)
		return nil, err
	}

	return &pinFiles2IpfsResponse.Data, nil
}
