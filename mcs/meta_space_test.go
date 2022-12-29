package mcs

import (
	"log"
	"testing"
	"unsafe"
)

const (
	BucketNameForTest = "zzq-test"
	FileNameForTest   = "1.jpeg"
	BucketUidForTest  = "cb9063b5-1fbb-4efa-b23f-fcaa7fbecfd4"
	FileIdForTest     = 4064
	Limit             = "10"
	Offset            = "0"
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
	resp, err := metaClient.DeleteBucket(BucketUidForTest)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(*(*string)(unsafe.Pointer(&resp)))
}

func TestMetaSpaceGetFileInfo(t *testing.T) {
	metaClient := NewMetaSpaceClient()
	err := metaClient.GetConfig().GetToken()
	if err != nil {
		log.Println(err)
		return
	}
	resp, err := metaClient.GetFileInfo(FileIdForTest)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(*(*string)(unsafe.Pointer(&resp)))
}

func TestMetaSpaceGetFileList(t *testing.T) {
	metaClient := NewMetaSpaceClient()
	err := metaClient.GetConfig().GetToken()
	if err != nil {
		log.Println(err)
		return
	}
	resp, err := metaClient.GetFileList("cb9063b5-1fbb-4efa-b23f-fcaa7fbecfd4", Limit, Offset)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(*(*string)(unsafe.Pointer(&resp)))
}

func TestMetaSpaceDeleteFile(t *testing.T) {
	metaClient := NewMetaSpaceClient()
	err := metaClient.GetConfig().GetToken()
	if err != nil {
		log.Println(err)
		return
	}
	resp, err := metaClient.DeleteFile(FileIdForTest)
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
