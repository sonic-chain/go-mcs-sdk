package mcs

import (
	"log"
	"testing"
	"unsafe"
)

const (
	BucketNameForTest = "zzq-test"
	FileNameForTest   = "1.jpeg"
	BucketUuidForTest = "c87e2640-a936-4a32-981d-3ea037040e29"
)

func TestMetaSpaceGetBuckets(t *testing.T) {
	metaClient := NewMetaSpaceClient()
	err := metaClient.GetConfig().GetToken()
	if err != nil {
		log.Println(err)
		return
	}
	resp, err := metaClient.GetBuckets()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(*(*string)(unsafe.Pointer(&resp)))
}

func TestMetaSpaceCreateBucket(t *testing.T) {
	metaClient := NewMetaSpaceClient()
	err := metaClient.GetConfig().GetToken()
	if err != nil {
		log.Println(err)
		return
	}
	fileId, err := metaClient.CreateBucket(BucketNameForTest)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(fileId)
}

func TestMetaSpaceDeleteBucket(t *testing.T) {
	metaClient := NewMetaSpaceClient()
	err := metaClient.GetConfig().GetToken()
	if err != nil {
		log.Println(err)
		return
	}
	bucketUid := BucketUuidForTest
	resp, err := metaClient.DeleteBucket(bucketUid)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(*(*string)(unsafe.Pointer(&resp)))
}

func TestMetaSpaceGetBucketID(t *testing.T) {
	metaClient := NewMetaSpaceClient()
	err := metaClient.GetConfig().GetToken()
	if err != nil {
		log.Println(err)
	}
	bucketId, err := metaClient.GetBucketIDByBucketName(BucketNameForTest)
	if err != nil {
		log.Println(err)
	}
	log.Println(bucketId)
}

func TestMetaSpaceUploadToBucket(t *testing.T) {
	metaClient := NewMetaSpaceClient()
	err := metaClient.GetConfig().GetToken()
	if err != nil {
		log.Println(err)
	}
	resp, err := metaClient.UploadToBucket("zzq-test", "4444.jpeg", "/home/zzq/Pictures/4'#.jpeg")
	if err != nil {
		log.Println(err)
	}
	log.Println(*(*string)(unsafe.Pointer(&resp)))
}
