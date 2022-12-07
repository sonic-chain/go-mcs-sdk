package mcs

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"go-mcs-sdk/mcs/common"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"unsafe"
)

type MetaSpaceClient struct {
	MetaSpaceUrl                    string `json:"meta_space_url"`
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
	metaSpaceUrl := os.Getenv("META_SPACE_URL")
	if metaSpaceUrl == "" {
		err = fmt.Errorf("meta space url is null in .env file")
		log.Fatal(err)
		return client
	}
	client.MetaSpaceUrl = metaSpaceUrl
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
	httpRequestUrl := client.MetaSpaceUrl + common.DIRECTORY
	bucketListInfoBytes, err := common.HttpGet(httpRequestUrl, client.JwtToken, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println(*(*string)(unsafe.Pointer(&bucketListInfoBytes)))
	return bucketListInfoBytes, nil
}

func (client *MetaSpaceClient) GetBucketInfoByBucketName(bucketName string) ([]byte, error) {
	httpRequestUrl := client.MetaSpaceUrl + common.DIRECTORY + "/" + bucketName
	bucketListInfoBytes, err := common.HttpGet(httpRequestUrl, client.JwtToken, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println(*(*string)(unsafe.Pointer(&bucketListInfoBytes)))
	return bucketListInfoBytes, nil
}

func (client *MetaSpaceClient) GetBucketIDByBucketName(bucketName string) (string, error) {
	response, err := client.GetBuckets()
	bucketId := ""
	if err != nil {
		log.Println(err)
		return bucketId, err
	}
	//var objectList *ObjectList
	var dict map[string]interface{}
	err = json.Unmarshal(response, &dict)
	if err != nil {
		log.Println(err)
		return bucketId, err
	}
	log.Println(dict)
	dataInReturn := dict["data"].(map[string]interface{})
	objectInData := dataInReturn["objects"].([]interface{})
	for _, v := range objectInData {
		vObject := v.(map[string]interface{})
		bucketNameInReturn := vObject["name"].(string)
		if bucketName == bucketNameInReturn {
			bucketId = vObject["id"].(string)
		}
	}
	return bucketId, nil
}

func (client *MetaSpaceClient) GetFileIDByBucketNameAndFileName(bucketName, fileName string) (string, error) {
	response, err := client.GetBucketInfoByBucketName(bucketName)
	bucketId := ""
	if err != nil {
		log.Println(err)
		return bucketId, err
	}
	//var objectList *ObjectList
	var dict map[string]interface{}
	err = json.Unmarshal(response, &dict)
	if err != nil {
		log.Println(err)
		return bucketId, err
	}
	log.Println(dict)
	dataInReturn := dict["data"].(map[string]interface{})
	objectInData := dataInReturn["objects"].([]interface{})
	for _, v := range objectInData {
		vObject := v.(map[string]interface{})
		fileNameInReturn := vObject["name"].(string)
		if fileName == fileNameInReturn {
			bucketId = vObject["id"].(string)
		}
	}
	return bucketId, nil
}

func (client *MetaSpaceClient) CreateBucket(bucketName string) ([]byte, error) {
	httpRequestUrl := client.MetaSpaceUrl + common.DIRECTORY
	params := make(map[string]string)
	params["path"] = "/" + bucketName
	response, err := common.HttpPut(httpRequestUrl, client.JwtToken, params)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println(*(*string)(unsafe.Pointer(&response)))
	return response, nil
}

func (client *MetaSpaceClient) DeleteBucket(dirs []string) ([]byte, error) {
	httpRequestUrl := client.MetaSpaceUrl + common.DELETE_OBJECT
	params := make(map[string][]string)
	params["item"] = []string{}
	params["dirs"] = dirs
	response, err := common.HttpDelete(httpRequestUrl, client.JwtToken, params)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println(*(*string)(unsafe.Pointer(&response)))
	return response, nil
}

func (client *MetaSpaceClient) CreateUploadSession(bucketName, fileName, filePath string) ([]byte, error) {
	err := CheckSpecialChar(fileName, filePath)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	bucketInfoBytes, err := client.GetBucketInfoByBucketName(bucketName)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var dict map[string]interface{}
	err = json.Unmarshal(bucketInfoBytes, &dict)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	dataInReturn := dict["data"].(map[string]interface{})
	policyId := dataInReturn["policy"].(map[string]interface{})["id"]

	params := make(map[string]interface{})
	_, currentTIme := common.GetCurrentUtcSec()
	fileInfo, err := os.Stat(filePath)
	fileSize := fileInfo.Size()
	params["path"] = fmt.Sprintf("/%s", bucketName)
	params["size"] = fileSize
	params["name"] = fileName
	params["policy_id"] = policyId
	params["last_modified"] = currentTIme
	httpRequestUrl := client.MetaSpaceUrl + common.UPLOAD_SESSION
	response, err := common.HttpPut(httpRequestUrl, client.JwtToken, params)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println(*(*string)(unsafe.Pointer(&response)))
	return response, nil
}

func (client *MetaSpaceClient) UploadToBucket(bucketName, fileName, filePath string) ([]byte, error) {
	err := CheckSpecialChar(fileName, filePath)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if fileInfo.Size() == 0 {
		err = fmt.Errorf("please upload a file larger than 0 byte")
		log.Println(err)
		return nil, err
	}
	uploadSession, err := client.CreateUploadSession(bucketName, fileName, filePath)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var dict map[string]interface{}
	err = json.Unmarshal(uploadSession, &dict)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	sessionId := dict["data"].(map[string]interface{})["sessionID"].(string)
	log.Println(sessionId)
	httpRequestUrl := client.MetaSpaceUrl + common.UPLOAD_SESSION + "/" + sessionId + "/0"
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	params := make(map[string][]byte)
	params["file"] = fileContent
	response, err := common.HttpPost(httpRequestUrl, client.JwtToken, params)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println(*(*string)(unsafe.Pointer(&response)))
	return response, nil
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
