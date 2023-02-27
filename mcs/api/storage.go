package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-mcs-sdk/mcs/common/constants"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/filswan/go-swan-lib/logs"
	libutils "github.com/filswan/go-swan-lib/utils"
)

type UploadFile struct {
	SourceFileUploadId int64  `json:"source_file_upload_id"`
	PayloadCid         string `json:"payload_cid"`
	IpfsUrl            string `json:"ipfs_url"`
	FileSize           int64  `json:"file_size"`
	WCid               string `json:"w_cid"`
	Status             string `json:"status"`
}

type UploadFileResponse struct {
	Status  string     `json:"status"`
	Data    UploadFile `json:"data"`
	Message string     `json:"message"`
}

func (mcsCient *MCSClient) UploadFile(filePath string, fileType int) (*UploadFile, error) {
	httpRequestUrl := libutils.UrlJoin(mcsCient.BaseUrl, constants.API_URL_STORAGE_UPLOAD_FILE)
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, err := os.Open(filePath)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	defer file.Close()

	part1, err := writer.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	_, err = io.Copy(part1, file)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	err = writer.WriteField("duration", strconv.Itoa(constants.DURATION_DAYS_DEFAULT))
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	err = writer.WriteField("file_type", strconv.Itoa(fileType))
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	err = writer.Close()
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	httpClient := &http.Client{}
	req, err := http.NewRequest("POST", httpRequestUrl, payload)

	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", mcsCient.JwtToken))
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := httpClient.Do(req)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	var uploadFileResponse UploadFileResponse
	err = json.Unmarshal(body, &uploadFileResponse)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	if !strings.EqualFold(uploadFileResponse.Status, constants.HTTP_STATUS_SUCCESS) {
		err := fmt.Errorf("get parameters failed, status:%s,message:%s", uploadFileResponse.Status, uploadFileResponse.Message)
		logs.GetLogger().Error(err)
		return nil, err
	}

	return &uploadFileResponse.Data, nil
}
