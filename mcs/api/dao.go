package api

import (
	"go-mcs-sdk/mcs/common/constants"

	"github.com/filswan/go-swan-lib/logs"
	libutils "github.com/filswan/go-swan-lib/utils"
)

type Deal2PreSign struct {
	DealId              int64 `json:"deal_id"`
	SourceFileUploadCnt int   `json:"source_file_upload_cnt"`
	BatchCount          int   `json:"batch_count"`
}

func (mcsCient *McsClient) GetDeals2PreSign() ([]*Deal2PreSign, error) {
	apiUrl := libutils.UrlJoin(mcsCient.BaseUrl, constants.API_URL_DAO_GET_DEALS_2_PRE_SIGN)

	var deals []*Deal2PreSign
	err := HttpGet(apiUrl, mcsCient.JwtToken, nil, &deals)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return deals, nil
}

type Deal2SignBatchInfo struct {
	BatchNo int      `json:"batch_no"`
	WCid    []string `json:"w_cid"`
}

type Deal2Sign struct {
	OfflineDealId int64                 `json:"offline_deal_id"`
	CarFileId     int64                 `json:"car_file_id"`
	DealId        int64                 `json:"deal_id"`
	BatchCount    int                   `json:"batch_count"`
	BatchSizeMax  int                   `json:"batch_size_max"`
	BatchInfo     []*Deal2SignBatchInfo `json:"batch_info"`
}

func (mcsCient *McsClient) GetDeals2Sign() ([]*Deal2Sign, error) {
	apiUrl := libutils.UrlJoin(mcsCient.BaseUrl, constants.API_URL_DAO_GET_DEALS_2_SIGN)

	var deals []*Deal2Sign
	err := HttpGet(apiUrl, mcsCient.JwtToken, nil, &deals)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return deals, nil
}

func (mcsCient *McsClient) GetDeals2SignHash() ([]*Deal2Sign, error) {
	apiUrl := libutils.UrlJoin(mcsCient.BaseUrl, constants.API_URL_DAO_GET_DEALS_2_SIGN_HASH)

	var deals []*Deal2Sign
	err := HttpGet(apiUrl, mcsCient.JwtToken, nil, &deals)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return deals, nil
}
