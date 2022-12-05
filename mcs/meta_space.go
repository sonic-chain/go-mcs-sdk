package mcs

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"go-mcs-sdk/mcs/common"
	"log"
	"os"
	"time"
	"unsafe"
)

const (
	McsBackendBaseUrl   = "http://192.168.199.61:8889/api/"
	BucketNameForCreate = "zzq-test"
)

type MetaSpaceClient struct {
	MetaSpaceUrl                    string `json:"meta_space_url"`
	JwtToken                        string `json:"jwt_token"`
	UserWalletAddressForRegisterMcs string `json:"user_wallet_address_for_register_mcs"`
	UserWalletAddressPK             string `json:"user_wallet_address_pk"`
	ChainNameForRegisterOnMcs       string `json:"chain_name_for_register_on_mcs"`
}

func NewMetaSpaceClient(metaSpaceUrl string) *MetaSpaceClient {
	metaSpaceClient := MetaSpaceClient{
		MetaSpaceUrl: metaSpaceUrl,
	}
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
	return client
}

func (client *MetaSpaceClient) GetToken() error {
	mcsClient := NewClient(McsBackendBaseUrl)
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

type ObjectList struct {
	Parent        string         `json:"parent,omitempty"`
	Objects       []*Object      `json:"objects"`
	Policy        *PolicySummary `json:"policy,omitempty"`
	FolderCnt     int64          `json:"folder_cnt"`
	FreeFolderCnt int64          `json:"free_folder_cnt"`
	FileCnt       int64          `json:"file_cnt"`
}

// Object 文件或者目录
type Object struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	Path          string    `json:"path"`
	Pic           string    `json:"pic"`
	Size          uint64    `json:"size"`
	Type          string    `json:"type"`
	PayloadCid    string    `json:"payload_cid"`
	IpfsUrl       string    `json:"ipfs_url"`
	PinStatus     string    `json:"pin_status"`
	Date          time.Time `json:"date"`
	FilesCount    int       `json:"files_count"`
	CreateDate    time.Time `json:"create_date"`
	UpdateDate    time.Time `json:"update_date"`
	Key           string    `json:"key,omitempty"`
	SourceEnabled bool      `json:"source_enabled"`
}
type PolicySummary struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Type     string   `json:"type"`
	MaxSize  uint64   `json:"max_size"`
	FileType []string `json:"file_type"`
}
