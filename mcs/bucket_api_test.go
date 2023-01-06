package mcs

import (
	"log"
	"testing"
	"unsafe"
)

const (
	BucketNameForTest = "zzq-test"
	FileNameForTest   = "index.jpeg"
	BucketUidForTest  = "cb9063b5-1fbb-4efa-b23f-fcaa7fbecfd4"
	FileIdForTest     = 4064
	FolderNameForTest = "test-folder2"
	Limit             = "10"
	Offset            = "0"
	FileHashForTest   = "c09dbca3794c26051e0fa938fface360"
	FilePathForTest   = "/home/zzq/Pictures/index.jpeg"
	PrefixForTest     = ""
)

func TestBucketApiGetBuckets(t *testing.T) {
	bucketClient := NewBucketClient()
	err := bucketClient.GetConfig().GetJwtToken()
	if err != nil {
		log.Println(err)
		return
	}
	resp, err := bucketClient.GetBuckets()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(*(*string)(unsafe.Pointer(&resp)))
}

func TestBucketApiCreateBucket(t *testing.T) {
	bucketClient := NewBucketClient()
	err := bucketClient.GetConfig().GetJwtToken()
	if err != nil {
		log.Println(err)
		return
	}
	fileId, err := bucketClient.CreateBucket(BucketNameForTest)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(fileId)
}

func TestBucketApiDeleteBucket(t *testing.T) {
	bucketClient := NewBucketClient()
	err := bucketClient.GetConfig().GetJwtToken()
	if err != nil {
		log.Println(err)
		return
	}
	resp, err := bucketClient.DeleteBucket(BucketUidForTest)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(*(*string)(unsafe.Pointer(&resp)))
}

func TestBucketApiGetFileInfo(t *testing.T) {
	bucketClient := NewBucketClient()
	err := bucketClient.GetConfig().GetJwtToken()
	if err != nil {
		log.Println(err)
		return
	}
	resp, err := bucketClient.GetFileInfo(FileIdForTest)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(*(*string)(unsafe.Pointer(&resp)))
}

func TestBucketApiGetFileList(t *testing.T) {
	bucketClient := NewBucketClient()
	err := bucketClient.GetConfig().GetJwtToken()
	if err != nil {
		log.Println(err)
		return
	}
	resp, err := bucketClient.GetFileList("cb9063b5-1fbb-4efa-b23f-fcaa7fbecfd4", Limit, Offset)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(*(*string)(unsafe.Pointer(&resp)))
}

func TestBucketApiDeleteFile(t *testing.T) {
	bucketClient := NewBucketClient()
	err := bucketClient.GetConfig().GetJwtToken()
	if err != nil {
		log.Println(err)
		return
	}
	resp, err := bucketClient.DeleteFile(FileIdForTest)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(*(*string)(unsafe.Pointer(&resp)))
}

func TestBucketApiCreateFolder(t *testing.T) {
	bucketClient := NewBucketClient()
	err := bucketClient.GetConfig().GetJwtToken()
	if err != nil {
		log.Println(err)
		return
	}
	resp, err := bucketClient.CreateFolder(FolderNameForTest, "", BucketUidForTest)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(*(*string)(unsafe.Pointer(&resp)))
}

func TestBucketApiCheckFile(t *testing.T) {
	bucketClient := NewBucketClient()
	err := bucketClient.GetConfig().GetJwtToken()
	if err != nil {
		log.Println(err)
		return
	}
	resp, err := bucketClient.CheckFile(BucketUidForTest, FileHashForTest, FileNameForTest, PrefixForTest)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(*(*string)(unsafe.Pointer(&resp)))
}

func TestBucketApiUploadChunk(t *testing.T) {
	bucketClient := NewBucketClient()
	err := bucketClient.GetConfig().GetJwtToken()
	if err != nil {
		log.Println(err)
		return
	}
	resp, err := bucketClient.UploadChunk(FileHashForTest, FilePathForTest)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(resp)
}

func TestBucketApiMergeRequest(t *testing.T) {
	bucketClient := NewBucketClient()
	err := bucketClient.GetConfig().GetJwtToken()
	if err != nil {
		log.Println(err)
		return
	}
	resp, err := bucketClient.MergeRequest(BucketUidForTest, FileHashForTest, FileNameForTest, PrefixForTest)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(resp)
}
