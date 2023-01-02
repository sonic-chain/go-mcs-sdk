package mcs

import (
	"bytes"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"go-mcs-sdk/mcs/common"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"unsafe"
)

type MetaSpaceClient struct {
	JwtToken                        string `json:"jwt_token"`
	UserWalletAddressForRegisterMcs string `json:"user_wallet_address_for_register_mcs"`
	UserWalletAddressPK             string `json:"user_wallet_address_pk"`
	ChainNameForRegisterOnMcs       string `json:"chain_name_for_register_on_mcs"`
	McsBackendBaseUrl               string `json:"mcs_backend_base_url"`
}

func NewMetaSpaceClient() *MetaSpaceClient {
	metaSpaceClient := MetaSpaceClient{}
	return &metaSpaceClient
}

func (client *MetaSpaceClient) SetJwtToken(jwtToken string) *MetaSpaceClient {
	client.JwtToken = jwtToken
	return client
}

func (client *MetaSpaceClient) GetConfig() *MetaSpaceClient {
	err := common.LoadEnv()
	if err != nil {
		log.Fatal(err)
		return client
	}
	walletAddress := os.Getenv("USER_WALLET_ADDRESS_FOR_REGISTER_MCS")
	if walletAddress == "" {
		err = fmt.Errorf("user wallet address is null in .env file")
		log.Fatal(err)
		return client
	}
	client.UserWalletAddressForRegisterMcs = walletAddress
	walletAddressPK := os.Getenv("USER_WALLET_ADDRESS_PK")
	if walletAddressPK == "" {
		err = fmt.Errorf("user wallet address private key is null in .env file")
		log.Fatal(err)
		return client
	}
	client.UserWalletAddressPK = walletAddressPK
	chainNetworkName := os.Getenv("CHAIN_NAME_FOR_REGISTER_ON_MCS")
	if chainNetworkName == "" {
		err = fmt.Errorf("chain network name is null in .env file")
		log.Fatal(err)
		return client
	}
	client.ChainNameForRegisterOnMcs = chainNetworkName
	mcsBackendBaseUrl := os.Getenv("MCS_BACKEND_BASE_URL")
	if mcsBackendBaseUrl == "" {
		err = fmt.Errorf("mcs backend base url is null in .env file")
		log.Fatal(err)
		return client
	}
	client.McsBackendBaseUrl = mcsBackendBaseUrl
	return client
}

func (client *MetaSpaceClient) GetToken() error {
	mcsClient := NewClient(client.McsBackendBaseUrl)
	user, err := mcsClient.NewUserRegisterService().SetWalletAddress(client.UserWalletAddressForRegisterMcs).Do(context.Background())
	if err != nil {
		log.Println(err)
		return err
	}
	nonce := user.Data.Nonce
	privateKey, _ := crypto.HexToECDSA(client.UserWalletAddressPK)
	signature, _ := common.PersonalSign(nonce, privateKey)
	jwt, err := mcsClient.NewUserLoginService().SetNetwork(client.ChainNameForRegisterOnMcs).SetNonce(nonce).SetWalletAddress(client.UserWalletAddressForRegisterMcs).
		SetSignature(signature).Do(context.Background())
	if err != nil {
		log.Println(err)
		return err
	}
	client.SetJwtToken(jwt.Data.JwtToken)
	return nil
}

func (client *MetaSpaceClient) GetBuckets() ([]byte, error) {
	httpRequestUrl := client.McsBackendBaseUrl + common.BUCKET_LIST
	bucketListInfoBytes, err := common.HttpGet(httpRequestUrl, client.JwtToken, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println(*(*string)(unsafe.Pointer(&bucketListInfoBytes)))
	return bucketListInfoBytes, nil
}

func (client *MetaSpaceClient) CreateBucket(bucketName string) ([]byte, error) {
	httpRequestUrl := client.McsBackendBaseUrl + common.CREATE_BUCKET
	params := make(map[string]string)
	params["bucket_name"] = bucketName
	response, err := common.HttpPost(httpRequestUrl, client.JwtToken, params)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println(*(*string)(unsafe.Pointer(&response)))
	return response, nil
}

func (client *MetaSpaceClient) DeleteBucket(bucketUid string) ([]byte, error) {
	httpRequestUrl := client.McsBackendBaseUrl + common.DELETE_BUCKET + bucketUid
	response, err := common.HttpGet(httpRequestUrl, client.JwtToken, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println(*(*string)(unsafe.Pointer(&response)))
	return response, nil
}

func (client *MetaSpaceClient) GetFileInfo(fileId int) ([]byte, error) {
	httpRequestUrl := client.McsBackendBaseUrl + common.FILE_INFO + strconv.Itoa(fileId)
	params := make(map[string]int)
	params["file_id"] = fileId
	response, err := common.HttpGet(httpRequestUrl, client.JwtToken, params)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println(*(*string)(unsafe.Pointer(&response)))
	return response, nil
}

func (client *MetaSpaceClient) DeleteFile(fileId int) ([]byte, error) {
	httpRequestUrl := client.McsBackendBaseUrl + common.DELETE_FILE + strconv.Itoa(fileId)
	params := make(map[string]int)
	params["file_id"] = fileId
	response, err := common.HttpGet(httpRequestUrl, client.JwtToken, params)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println(*(*string)(unsafe.Pointer(&response)))
	return response, nil
}

func (client *MetaSpaceClient) GetFileList(fileUid, limit, offset string) ([]byte, error) {
	httpRequestUrl := client.McsBackendBaseUrl + common.FILE_LIST + fileUid + "&limit=" + limit + "&offset=" + offset
	response, err := common.HttpGet(httpRequestUrl, client.JwtToken, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println(*(*string)(unsafe.Pointer(&response)))
	return response, nil
}

func (client *MetaSpaceClient) CreateFolder(fileName, prefix, bucketUid string) ([]byte, error) {
	httpRequestUrl := client.McsBackendBaseUrl + common.CREATE_FOLDER
	params := make(map[string]string)
	params["file_name"] = fileName
	params["prefix"] = prefix
	params["bucket_uid"] = bucketUid
	response, err := common.HttpPost(httpRequestUrl, client.JwtToken, params)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println(*(*string)(unsafe.Pointer(&response)))
	return response, nil
}

func (client *MetaSpaceClient) CheckFile(bucketUid, fileHash, fileName, prefix string) ([]byte, error) {
	httpRequestUrl := client.McsBackendBaseUrl + common.CHECK_UPLOAD
	params := make(map[string]string)
	params["bucket_uid"] = bucketUid
	params["file_hash"] = fileHash
	params["file_name"] = fileName
	params["prefix"] = prefix
	response, err := common.HttpPost(httpRequestUrl, client.JwtToken, params)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println(*(*string)(unsafe.Pointer(&response)))
	return response, nil
}

func (client *MetaSpaceClient) UploadChunk(fileHash, uploadFilePath string) ([]byte, error) {
	httpRequestUrl := client.McsBackendBaseUrl + common.UPLOAD_CHUNK
	fileNameWithSuffix := path.Base(uploadFilePath)
	bodyBuffer := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuffer)
	err := bodyWriter.WriteField("hash", fileHash)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	err = bodyWriter.WriteField("file", fileHash)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	writer, err := bodyWriter.CreateFormFile("chunk_form", fileNameWithSuffix)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	chunkFile, err := os.Open(uploadFilePath)
	defer chunkFile.Close()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	_, err = io.Copy(writer, chunkFile)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer bodyWriter.Close()
	request, err := http.NewRequest("POST", httpRequestUrl, bodyBuffer)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", client.JwtToken))
	request.Header.Add("Content-Type", bodyWriter.FormDataContentType())
	response, err := http.DefaultClient.Do(request)
	//response, err := httpClient.Post(httpRequestUrl,bodyWriter.FormDataContentType(),bodyBuffer)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer response.Body.Close()
	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println(*(*string)(unsafe.Pointer(&responseBytes)))
	return responseBytes, nil
}

func (client *MetaSpaceClient) UploadToBucket(bucketName, fileName, filePath string) ([]byte, error) {
	return nil, nil
}

func CheckSpecialChar(fileName, filePath string) error {
	if SpecialChar(fileName) {
		err := fmt.Errorf("file alias has special characters")
		log.Println(err)
		return err
	}
	_, fileNameInPath := filepath.Split(filePath)
	if SpecialChar(fileNameInPath) {
		err := fmt.Errorf("file name in path has special characters")
		log.Println(err)
		return err
	}
	return nil
}

func SpecialChar(line string) bool {
	specialCharacters := "[!@#$%^&*()-+?=,<>/'\" \n\t\v\f\r]"
	reg := regexp.MustCompile(specialCharacters)
	matchStr := reg.FindAllString(line, -1)
	if len(matchStr) > 0 {
		return true
	} else {
		return false
	}
}
