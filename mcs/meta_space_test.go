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

func TestMetaSpaceCreateFolder(t *testing.T) {
	metaClient := NewMetaSpaceClient()
	err := metaClient.GetConfig().GetToken()
	if err != nil {
		log.Println(err)
		return
	}
	resp, err := metaClient.CreateFolder(FolderNameForTest, "", BucketUidForTest)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(*(*string)(unsafe.Pointer(&resp)))
}

func TestMetaSpaceCheckFile(t *testing.T) {
	metaClient := NewMetaSpaceClient()
	err := metaClient.GetConfig().GetToken()
	if err != nil {
		log.Println(err)
		return
	}
	resp, err := metaClient.CheckFile(BucketUidForTest, FileHashForTest, FileNameForTest, PrefixForTest)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(*(*string)(unsafe.Pointer(&resp)))
}

func TestMetaSpaceUploadChunk(t *testing.T) {
	metaClient := NewMetaSpaceClient()
	err := metaClient.GetConfig().GetToken()
	if err != nil {
		log.Println(err)
		return
	}
	resp, err := metaClient.UploadChunk(FileHashForTest, FilePathForTest)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(resp)
}

func TestMetaSpaceMergeRequest(t *testing.T) {
	metaClient := NewMetaSpaceClient()
	err := metaClient.GetConfig().GetToken()
	if err != nil {
		log.Println(err)
		return
	}
	resp, err := metaClient.MergeRequest(BucketUidForTest, FileHashForTest, FileNameForTest, PrefixForTest)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(resp)
}
