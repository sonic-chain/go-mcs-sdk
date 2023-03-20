package api

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"go-mcs-sdk/mcs/api/common/constants"
	"go-mcs-sdk/mcs/api/common/utils"
	"go-mcs-sdk/mcs/api/common/web"
	"go-mcs-sdk/mcs/contract"
	"math/big"
	"net/url"
	"strings"

	"go-mcs-sdk/mcs/api/common/logs"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func (onChainClient *OnChainClient) GetFileCoinPrice() (*float64, error) {
	apiUrl := utils.UrlJoin(onChainClient.BaseUrl, constants.API_URL_BILLING_FILECOIN_PRICE)
	params := url.Values{}

	var price float64
	err := web.HttpGet(apiUrl, onChainClient.JwtToken, strings.NewReader(params.Encode()), &price)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return &price, nil
}

type LockPaymentInfo struct {
	WCid         string `json:"w_cid"`
	PayAmount    string `json:"pay_amount"`
	PayTxHash    string `json:"pay_tx_hash"`
	TokenAddress string `json:"token_address"`
}

func (client *OnChainClient) GetLockPaymentInfo(fileUploadId int64) (*LockPaymentInfo, error) {
	apiUrl := utils.UrlJoin(client.BaseUrl, constants.API_URL_BILLING_GET_PAYMENT_INFO)
	apiUrl = apiUrl + "?source_file_upload_id=" + fmt.Sprintf("%d", fileUploadId)
	params := url.Values{}

	var lockPaymentInfo LockPaymentInfo
	err := web.HttpGet(apiUrl, client.JwtToken, strings.NewReader(params.Encode()), &lockPaymentInfo)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	return &lockPaymentInfo, nil
}

type BillingHistory struct {
	PayId        int64  `json:"pay_id"`
	PayTxHash    string `json:"pay_tx_hash"`
	PayAmount    string `json:"pay_amount"`
	UnlockAmount string `json:"unlock_amount"`
	FileName     string `json:"file_name"`
	PayloadCid   string `json:"payload_cid"`
	PayAt        int64  `json:"pay_at"`
	UnlockAt     int64  `json:"unlock_at"`
	Deadline     int64  `json:"deadline"`
	NetworkName  string `json:"network_name"`
	TokenName    string `json:"token_name"`
}

type BillingHistoryParams struct {
	PageNumber *int    `json:"page_number"`
	PageSize   *int    `json:"page_size"`
	FileName   *string `json:"file_name"`
	TxHash     *string `json:"tx_hash"`
	OrderBy    *string `json:"order_by"`
	IsAscend   *string `json:"is_ascend"`
}

func (onChainClient *OnChainClient) GetBillingHistory(billingHistoryParams BillingHistoryParams) ([]*BillingHistory, *int64, error) {
	apiUrl := utils.UrlJoin(onChainClient.BaseUrl, constants.API_URL_BILLING_HISTORY)
	paramItems := []string{}
	if billingHistoryParams.PageNumber != nil {
		paramItems = append(paramItems, "page_number="+fmt.Sprintf("%d", *billingHistoryParams.PageNumber))
	}

	if billingHistoryParams.PageSize != nil {
		paramItems = append(paramItems, "page_size="+fmt.Sprintf("%d", *billingHistoryParams.PageSize))
	}

	if billingHistoryParams.FileName != nil {
		paramItems = append(paramItems, "file_name="+*billingHistoryParams.FileName)
	}

	if billingHistoryParams.TxHash != nil {
		paramItems = append(paramItems, "tx_hash="+*billingHistoryParams.TxHash)
	}

	if billingHistoryParams.OrderBy != nil {
		paramItems = append(paramItems, "order_by="+*billingHistoryParams.OrderBy)
	}

	if billingHistoryParams.IsAscend != nil {
		paramItems = append(paramItems, "is_ascend="+*billingHistoryParams.IsAscend)
	}

	if len(paramItems) > 0 {
		apiUrl = apiUrl + "?"
		for _, paramItem := range paramItems {
			apiUrl = apiUrl + paramItem + "&"
		}

		apiUrl = strings.TrimRight(apiUrl, "&")
	}

	var billings struct {
		Billing          []*BillingHistory `json:"billing"`
		TotalRecordCount int64             `json:"total_record_count"`
	}

	err := web.HttpGet(apiUrl, onChainClient.JwtToken, nil, &billings)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, nil, err
	}

	return billings.Billing, &billings.TotalRecordCount, nil
}

func (onChainClient *OnChainClient) PayForFile(sourceFileUploadId int64, privateKeyStr string, rpcUrl string) (*string, error) {
	sourceFileUpload, err := onChainClient.GetSourceFileUpload(sourceFileUploadId)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	historicalAveragePriceVerified, err := GetHistoricalAveragePriceVerified()
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	systemParams, err := onChainClient.GetSystemParam()
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	amount, err := GetAmount(sourceFileUpload.FileSize, historicalAveragePriceVerified, systemParams.FilecoinPrice, constants.COPY_NUMBER_LIMIT)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	ChainId, err := client.ChainID(context.Background())
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, ChainId)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	SwanPayment, err := contract.NewSwanPayment(common.HexToAddress(systemParams.PaymentContractAddress), client)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	minPayment := big.NewInt(amount)
	amount2Lock := big.NewInt(int64(float64(amount) * float64(systemParams.PayMultiplyFactor)))
	lockTime := int64(constants.DURATION_DAYS_DEFAULT) * constants.SECOND_PER_DAY
	var paymentParam = contract.IPaymentMinimallockPaymentParam{
		Id:         sourceFileUpload.WCid,
		MinPayment: minPayment,
		Amount:     amount2Lock,
		LockTime:   big.NewInt(lockTime),
		Recipient:  common.HexToAddress(systemParams.PaymentRecipientAddress),
		Size:       big.NewInt(sourceFileUpload.FileSize),
		CopyLimit:  constants.COPY_NUMBER_LIMIT,
	}

	txHashApprove, err := Approve(rpcUrl, systemParams, privateKey, amount2Lock)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	logs.GetLogger().Info(*txHashApprove)

	tx, err := SwanPayment.LockTokenPayment(&bind.TransactOpts{
		From:   auth.From,
		Signer: auth.Signer,
	}, paymentParam)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	bind.WaitMined(context.Background(), client, tx)
	tx, _, err = client.TransactionByHash(context.Background(), tx.Hash())
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	txHash := tx.Hash().String()

	return &txHash, nil
}

func Approve(rpcUrl string, systemParams *SystemParam, privateKey *ecdsa.PrivateKey, amount *big.Int) (*string, error) {
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		err := fmt.Errorf("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		logs.GetLogger().Error(err)
		return nil, err
	}

	walletAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	ERC20, err := contract.NewERC20(common.HexToAddress(systemParams.UsdcAddress), client)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	balance, err := ERC20.BalanceOf(&bind.CallOpts{
		Pending:     false,
		From:        common.Address{},
		BlockNumber: nil,
		Context:     nil,
	}, walletAddress)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	logs.GetLogger().Info(walletAddress)
	balanceInt64 := balance.Int64()

	if amount.Int64() > balanceInt64 {
		err := fmt.Errorf("BalanceOf error")
		logs.GetLogger().Error(err)
		return nil, err
	}

	USDCSpender := common.HexToAddress(systemParams.PaymentContractAddress)

	ChainId, err := client.ChainID(context.Background())
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, ChainId)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	logs.GetLogger().Info("USDCSpender:", USDCSpender)
	tx, err := ERC20.Approve(&bind.TransactOpts{
		From:   auth.From,
		Signer: auth.Signer,
	}, USDCSpender, amount)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	bind.WaitMined(context.Background(), client, tx)
	tx, _, err = client.TransactionByHash(context.Background(), tx.Hash())
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	txHash := tx.Hash().String()
	return &txHash, nil
}
