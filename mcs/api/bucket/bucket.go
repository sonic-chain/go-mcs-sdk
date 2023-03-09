package bucket

import (
	"go-mcs-sdk/mcs/api/common/auth"
	"go-mcs-sdk/mcs/api/common/constants"
	"go-mcs-sdk/mcs/api/common/web"

	"github.com/filswan/go-swan-lib/logs"
	libutils "github.com/filswan/go-swan-lib/utils"
)

type BucketClient struct {
	auth.McsClient
}

func GetBucketClient(mcsClient auth.McsClient) *BucketClient {
	var bucketClient = &BucketClient{}

	bucketClient.BaseUrl = mcsClient.BaseUrl
	bucketClient.JwtToken = mcsClient.JwtToken

	return bucketClient
}

func (bucketClient *BucketClient) CreateBucket(bucketName string) (*string, error) {
	apiUrl := libutils.UrlJoin(bucketClient.BaseUrl, constants.API_URL_BUCKET_CREATE_BUCKET)

	var bucket struct {
		BucketName string `json:"bucket_name"`
	}
	bucket.BucketName = bucketName

	var bucketUid string
	err := web.HttpPost(apiUrl, bucketClient.JwtToken, &bucket, &bucketUid)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return &bucketUid, nil
}

type Bucket struct {
	BucketUid  string `json:"bucket_uid"`
	Address    string `json:"address"`
	MaxSize    int64  `json:"max_size"`
	Size       int64  `json:"size"`
	IsFree     bool   `json:"is_free"`
	PaymentTx  string `json:"payment_tx"`
	IsActive   bool   `json:"is_active"`
	IsDeleted  bool   `json:"is_deleted"`
	BucketName string `json:"bucket_name"`
	FileNumber int64  `json:"file_number"`
}

func (bucketClient *BucketClient) GetBuckets() ([]*Bucket, error) {
	apiUrl := libutils.UrlJoin(bucketClient.BaseUrl, constants.API_URL_BUCKET_GET_BUCKET_LIST)

	var buckets []*Bucket
	err := web.HttpGet(apiUrl, bucketClient.JwtToken, nil, &buckets)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return buckets, nil
}

func (bucketClient *BucketClient) DeleteBucket(bucketUid string) error {
	apiUrl := libutils.UrlJoin(bucketClient.BaseUrl, constants.API_URL_BUCKET_DELETE_BUCKET)
	apiUrl = apiUrl + "?bucket_uid=" + bucketUid

	err := web.HttpGet(apiUrl, bucketClient.JwtToken, nil, nil)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	return nil
}

func (bucketClient *BucketClient) RenameBucket(newBucketName string, bucketUid string) error {
	apiUrl := libutils.UrlJoin(bucketClient.BaseUrl, constants.API_URL_BUCKET_RENAME_BUCKET)

	var params struct {
		BucketName string `json:"bucket_name"`
		BucketUid  string `json:"bucket_uid"`
	}
	params.BucketName = newBucketName
	params.BucketUid = bucketUid

	var result string
	err := web.HttpPost(apiUrl, bucketClient.JwtToken, &params, &result)
	if err != nil {
		logs.GetLogger().Error(err)
		return err
	}

	return nil
}

func (bucketClient *BucketClient) GetTotalStorageSize() (*int64, error) {
	apiUrl := libutils.UrlJoin(bucketClient.BaseUrl, constants.API_URL_BUCKET_GET_TOTAL_STORAGE_SIZE)

	var totalStorageSize int64

	err := web.HttpGet(apiUrl, bucketClient.JwtToken, nil, &totalStorageSize)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return &totalStorageSize, nil
}
