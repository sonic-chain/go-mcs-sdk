package bucket

import (
	"testing"

	"github.com/filswan/go-mcs-sdk/mcs/api/user"

	"github.com/filswan/go-mcs-sdk/mcs/api/common/logs"

	"github.com/filswan/go-mcs-sdk/mcs/api/common/constants"

	"github.com/stretchr/testify/assert"
)

var buketClient *BucketClient
var network = constants.PAYMENT_CHAIN_NAME_POLYGON_MUMBAI
var apikey = "9EO9I6rzlfYkcltzOo0ayp"
var accessToken = "hmvYOnAv9JAtXqzi5NWDfuRYMJXY6LDG"
var file2Upload = "/Users/dorachen/work/test2/duration6"
var folder2Upload = "/Users/dorachen/work/test3"

func init() {
	if buketClient != nil {
		return
	}

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
