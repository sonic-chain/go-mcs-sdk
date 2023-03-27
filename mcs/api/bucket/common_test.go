package bucket

import (
	"go-mcs-sdk/mcs/api/user"
	"testing"

	"go-mcs-sdk/mcs/api/common/constants"
	"go-mcs-sdk/mcs/api/common/logs"

	"github.com/stretchr/testify/assert"
)

var buketClient *BucketClient
var network = constants.PAYMENT_CHAIN_NAME_POLYGON_MUMBAI

func init() {
	if buketClient != nil {
		return
	}

	apikey := ""
	accessToken := ""

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
