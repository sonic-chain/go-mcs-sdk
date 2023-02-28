package api

import (
	"encoding/json"
	"fmt"
	"go-mcs-sdk/mcs/common/constants"
	"strings"

	"github.com/filswan/go-swan-lib/client/web"
	"github.com/filswan/go-swan-lib/logs"
	libutils "github.com/filswan/go-swan-lib/utils"
)

type Deal2PreSign struct {
	DealId              int64 `json:"deal_id"`
	SourceFileUploadCnt int   `json:"source_file_upload_cnt"`
	BatchCount          int   `json:"batch_count"`
}

type GetDeals2PreSignResponse struct {
	Response
	Data []*Deal2PreSign `json:"data"`
}

func (mcsCient *MCSClient) GetDeals2PreSign() ([]*Deal2PreSign, error) {
	apiUrl := libutils.UrlJoin(mcsCient.BaseUrl, constants.API_URL_DAO_GET_DEALS_2_PRE_SIGN)
	result, err := web.HttpGet(apiUrl, mcsCient.JwtToken, nil)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	var getDeals2PreSignResponse GetDeals2PreSignResponse
	err = json.Unmarshal(result, &getDeals2PreSignResponse)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	if !strings.EqualFold(getDeals2PreSignResponse.Status, constants.HTTP_STATUS_SUCCESS) {
		err := fmt.Errorf("get parameters failed, status:%s,message:%s", getDeals2PreSignResponse.Status, getDeals2PreSignResponse.Message)
		logs.GetLogger().Error(err)
		return nil, err
	}

	return getDeals2PreSignResponse.Data, nil
}
