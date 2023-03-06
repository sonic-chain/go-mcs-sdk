package bucket

import (
	"go-mcs-sdk/mcs/api/common/constants"
	"go-mcs-sdk/mcs/api/common/utils"
	"log"
	"strconv"

	"github.com/filswan/go-swan-lib/logs"
	libutils "github.com/filswan/go-swan-lib/utils"
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
}

func (bucketClient *BucketClient) GetFileInfo(fileId int) (*OssFile, error) {
	apiUrl := libutils.UrlJoin(bucketClient.BaseUrl, constants.API_URL_BUCKET_FILE_GET_FILE_INFO)

	var fileInfo OssFile
	err := utils.HttpGet(apiUrl, bucketClient.JwtToken, &fileId, &fileInfo)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return &fileInfo, nil
}

func (bucketClient *BucketClient) DeleteFile(fileId int) error {
	apiUrl := libutils.UrlJoin(bucketClient.BaseUrl, constants.API_URL_BUCKET_FILE_DELETE_FILE, strconv.Itoa(fileId))

	err := utils.HttpGet(apiUrl, bucketClient.JwtToken, &fileId, nil)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (bucketClient *BucketClient) CreateFolder(fileName, prefix, bucketUid string) error {
	apiUrl := libutils.UrlJoin(bucketClient.BaseUrl, constants.API_URL_BUCKET_FILE_CREATE_FOLDER)

	var params struct {
		FileName  string `json:"file_name"`
		Prefix    string `json:"prefix"`
		BucketUid string `json:"bucket_uid"`
	}

	params.FileName = fileName
	params.Prefix = prefix
	params.BucketUid = bucketUid

	err := utils.HttpPost(apiUrl, bucketClient.JwtToken, &params, nil)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (bucketClient *BucketClient) GetFileInfoByObjectName(objectName, bucketUid string) (*OssFile, error) {
	apiUrl := libutils.UrlJoin(bucketClient.BaseUrl, constants.API_URL_BUCKET_FILE_GET_FILE_INFO)
	apiUrl = apiUrl + "?bucket_uid=" + bucketUid + "&objectName=" + objectName

	var fileInfo OssFile
	err := utils.HttpGet(apiUrl, bucketClient.JwtToken, nil, &fileInfo)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return &fileInfo, nil
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

func (bucketClient *BucketClient) CheckFile(bucketUid, fileHash, fileName, prefix string) (*OssFileInfo, error) {
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
	err := utils.HttpPost(apiUrl, bucketClient.JwtToken, &params, &ossFileInfo)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &ossFileInfo, nil
}
