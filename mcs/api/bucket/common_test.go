package bucket

import (
	"go-mcs-sdk/mcs/api/user"
	"go-mcs-sdk/mcs/config"
	"testing"

	"go-mcs-sdk/mcs/api/common/logs"

	"github.com/stretchr/testify/assert"
)

var buketClient *BucketClient

func init() {
	if buketClient != nil {
		return
	}

	apikey := config.GetConfig().Apikey
	accessToken := config.GetConfig().AccessToken
	network := config.GetConfig().Network

	mcsClient, err := user.LoginByApikey(apikey, accessToken, network)
	if err != nil {
		logs.GetLogger().Fatal(err)
	}

	buketClient = GetBucketClient(*mcsClient)
}

func TestGetGateway(t *testing.T) {
	gateway, err := buketClient.GetGateway()
	assert.Nil(t, err)
	assert.NotEmpty(t, gateway)

	logs.GetLogger().Info(*gateway)
}
