package bucket

import (
	"testing"

	"github.com/filswan/go-swan-lib/logs"
)

func TestCreateBucket(t *testing.T) {
	bucketUid, err := buketClient.CreateBucket("test23")
	if err != nil {
		logs.GetLogger().Fatal(err)
	}

	logs.GetLogger().Info(*bucketUid)
}

func TestGetBuckets(t *testing.T) {
	buckets, err := buketClient.GetBuckets()
	if err != nil {
		logs.GetLogger().Fatal(err)
	}

	for _, bucket := range buckets {
		logs.GetLogger().Info(*bucket)
	}
}

func TestDeleteBucket(t *testing.T) {
	err := buketClient.DeleteBucket("7bb5d325-e31c-486d-8420-169067dc401b")
	if err != nil {
		logs.GetLogger().Fatal(err)
	}
}

func TestRenameBucket(t *testing.T) {
	err := buketClient.RenameBucket("tests", "a7303d2a-acd2-48ac-a062-8454bbf148d2")
	if err != nil {
		logs.GetLogger().Fatal(err)
	}
}

func TestGetTotalStorageSize(t *testing.T) {
	totalStorageSize, err := buketClient.GetTotalStorageSize()
	if err != nil {
		logs.GetLogger().Fatal(err)
	}

	logs.GetLogger().Info(*totalStorageSize)
}
