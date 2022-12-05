package mcs

import (
	"log"
	"testing"
	"unsafe"
)

const (
	MetaSpaceBackendBaseUrl = "http://192.168.199.61:5212/api/"
	BucketNameForTest       = "zzq-1105"
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
	err = metaClient.GetBucketInfoByBucketName(BucketNameForTest)
	if err != nil {
		log.Println(err)
	}
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
