package api

import (
	"encoding/json"
	"fmt"
	"go-mcs-sdk/mcs/common/constants"
	"net/url"
	"strings"

	"github.com/filswan/go-swan-lib/client/web"
	"github.com/filswan/go-swan-lib/logs"
	libutils "github.com/filswan/go-swan-lib/utils"
)

type FileCoinPriceResponse struct {
	Status  string  `json:"status"`
	Data    float64 `json:"data"`
	Message string  `json:"message"`
}

func (mcsCient *MCSClient) GetFileCoinPrice() (*float64, error) {
	apiUrl := libutils.UrlJoin(mcsCient.BaseUrl, constants.API_URL_FILECOIN_PRICE)
	params := url.Values{}
	response, err := web.HttpGet(apiUrl, mcsCient.JwtToken, strings.NewReader(params.Encode()))
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	var fileCoinPriceResponse FileCoinPriceResponse
	err = json.Unmarshal(response, &fileCoinPriceResponse)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	if !strings.EqualFold(fileCoinPriceResponse.Status, constants.HTTP_STATUS_SUCCESS) {
		err := fmt.Errorf("get parameters failed, status:%s,message:%s", fileCoinPriceResponse.Status, fileCoinPriceResponse.Message)
		logs.GetLogger().Error(err)
		return nil, err
	}

	return &fileCoinPriceResponse.Data, nil
}
