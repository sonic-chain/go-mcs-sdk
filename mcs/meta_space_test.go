package mcs

import (
	"log"
	"testing"
	"unsafe"
)

const (
	MetaSpaceBackendBaseUrl = "http://192.168.199.61:5212/api/"
	BucketNameForTest       = "zzq-1105"
	FileNameForTest         = "1.jpeg"
)

func TestMetaSpaceGetBuckets(t *testing.T) {
	metaClient := NewMetaSpaceClient(MetaSpaceBackendBaseUrl)
	err := metaClient.GetToken()
	if err != nil {
		log.Println(err)
	}
	resp, err := metaClient.GetBuckets()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(*(*string)(unsafe.Pointer(&resp)))
}

func TestMetaSpaceGetBucketInfo(t *testing.T) {
	metaClient := NewMetaSpaceClient(MetaSpaceBackendBaseUrl)
	err := metaClient.GetToken()
	if err != nil {
		log.Println(err)
	}
	resp, err := metaClient.GetBucketInfoByBucketName(BucketNameForTest)
	if err != nil {
		log.Println(err)
	}
	log.Println(*(*string)(unsafe.Pointer(&resp)))
}

func TestMetaSpaceGetBucketID(t *testing.T) {
	metaClient := NewMetaSpaceClient(MetaSpaceBackendBaseUrl)
	err := metaClient.GetToken()
	if err != nil {
		log.Println(err)
	}
	bucketId, err := metaClient.GetBucketIDByBucketName(BucketNameForTest)
	if err != nil {
		log.Println(err)
	}
	log.Println(bucketId)
}

func TestMetaSpaceGetFileID(t *testing.T) {
	metaClient := NewMetaSpaceClient(MetaSpaceBackendBaseUrl)
	err := metaClient.GetToken()
	if err != nil {
		log.Println(err)
	}
	fileId, err := metaClient.GetFileIDByBucketNameAndFileName(BucketNameForTest, FileNameForTest)
	if err != nil {
		log.Println(err)
	}
	log.Println(fileId)
}

func TestMetaSpaceCreateBucket(t *testing.T) {
	metaClient := NewMetaSpaceClient(MetaSpaceBackendBaseUrl)
	err := metaClient.GetToken()
	if err != nil {
		log.Println(err)
	}
	fileId, err := metaClient.CreateBucket(BucketNameForCreate)
	if err != nil {
		log.Println(err)
	}
	log.Println(fileId)
}
