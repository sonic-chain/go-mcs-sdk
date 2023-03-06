package bucket

import (
	"testing"

	"github.com/filswan/go-swan-lib/logs"
)

func TestCreateBucket(t *testing.T) {
	bucketUid, err := onChainClient.CreateBucket("test1")
	if err != nil {
		logs.GetLogger().Fatal(err)
	}

	logs.GetLogger().Info(*bucketUid)
}
