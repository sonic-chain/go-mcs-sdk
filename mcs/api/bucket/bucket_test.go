package bucket

import (
	"testing"

	"go-mcs-sdk/mcs/api/common/logs"

	"github.com/stretchr/testify/assert"
)

func TestCreateBucket(t *testing.T) {
	bucketUid, err := buketClient.CreateBucket("test23")
	assert.Nil(t, err)
	assert.NotEmpty(t, bucketUid)

	logs.GetLogger().Info(*bucketUid)
}

func TestGetBuckets(t *testing.T) {
	buckets, err := buketClient.GetBuckets()
	assert.Nil(t, err)
	assert.NotEmpty(t, buckets)

	for _, bucket := range buckets {
		logs.GetLogger().Info(*bucket)
	}
}

func TestDeleteBucket(t *testing.T) {
	err := buketClient.DeleteBucket("a7303d2a-acd2-48ac-a062-8454bbf148d2")
	assert.Nil(t, err)
}

func TestRenameBucket(t *testing.T) {
	err := buketClient.RenameBucket("abc", "0ef9c94d-9bb9-4ce9-b687-7db732a9ce2e")
	assert.Nil(t, err)
}

func TestGetTotalStorageSize(t *testing.T) {
	totalStorageSize, err := buketClient.GetTotalStorageSize()
	assert.Nil(t, err)
	assert.NotNil(t, totalStorageSize)

	logs.GetLogger().Info(*totalStorageSize)
}
