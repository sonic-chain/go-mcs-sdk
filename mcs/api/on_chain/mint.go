package api

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"go-mcs-sdk/mcs/api/common/logs"
	common2 "go-mcs-sdk/mcs/common"
	"go-mcs-sdk/mcs/contract"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ContractApproveUSDCService struct {
	c                      *Client
	WalletAddress          string
	PrivateKey             string
	Amount                 *big.Int
	RpcEndpoint            string
	USDCAddress            string
	PaymentContractAddress string
}

func (s *ContractApproveUSDCService) SetUSDCAddress(USDCAddress string) *ContractApproveUSDCService {
	s.USDCAddress = USDCAddress
	return s
}

func (s *ContractApproveUSDCService) SetPaymentContractAddress(PaymentContractAddress string) *ContractApproveUSDCService {
	s.PaymentContractAddress = PaymentContractAddress
	return s
}

func (s *ContractApproveUSDCService) SetWalletAddress(WalletAddress string) *ContractApproveUSDCService {
	s.WalletAddress = WalletAddress
	return s
}

func (s *ContractApproveUSDCService) SetPrivateKey(PrivateKey string) *ContractApproveUSDCService {
	s.PrivateKey = PrivateKey
	return s
}

func (s *ContractApproveUSDCService) SetAmount(Amount *big.Int) *ContractApproveUSDCService {
	s.Amount = Amount
	return s
}

func (s *ContractApproveUSDCService) SetRpcEndpoint(RpcEndpoint string) *ContractApproveUSDCService {
	s.RpcEndpoint = RpcEndpoint
	return s
}

func (onChainClient *OnChainClient) Mint(privateKeyStr, rpcUrl string) (TX *string, err error) {
	systemParams, err := onChainClient.GetSystemParam()
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}

	privateKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		logs.GetLogger().Error(err)
		return nil, err
	}
	USDCSpender := common.HexToAddress(systemParams.PaymentContractAddress)

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

	ERC20, err := contract.NewERC20(common.HexToAddress(s.USDCAddress), client)
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

	return tx.Hash().String(), nil
}

type ContractUploadFilePayService struct {
	c                       *Client
	WalletAddress           string
	PrivateKey              string
	RpcEndpoint             string
	FileSize                int64
	WCid                    string
	Rate                    float64
	PaymentRecipientAddress string
	PayMultiplyFactor       float64
	PaymentContractAddress  string
	LockTime                big.Int
}

func (s *ContractUploadFilePayService) SetFileSize(FileSize int64) *ContractUploadFilePayService {
	s.FileSize = FileSize
	return s
}

func (s *ContractUploadFilePayService) SetWCid(WCid string) *ContractUploadFilePayService {
	s.WCid = WCid
	return s
}

func (s *ContractUploadFilePayService) SetRate(Rate float64) *ContractUploadFilePayService {
	s.Rate = Rate
	return s
}

func (s *ContractUploadFilePayService) SetPaymentRecipientAddress(PaymentRecipientAddress string) *ContractUploadFilePayService {
	s.PaymentRecipientAddress = PaymentRecipientAddress
	return s
}

func (s *ContractUploadFilePayService) SetPayMultiplyFactor(PayMultiplyFactor float64) *ContractUploadFilePayService {
	s.PayMultiplyFactor = PayMultiplyFactor
	return s
}
func (s *ContractUploadFilePayService) SetLockTime(LockTime big.Int) *ContractUploadFilePayService {
	s.LockTime = LockTime
	return s
}

func (s *ContractUploadFilePayService) Do(ctx context.Context, opts ...RequestOption) (Tx string, err error) {
	amount := common2.GetAmount(s.FileSize, s.Rate)
	privateKey, err := crypto.HexToECDSA(s.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	client, err := ethclient.Dial(s.RpcEndpoint)
	if err != nil {
		log.Fatal(err)
	}
	ChainId, _ := client.ChainID(context.Background())

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, ChainId)

	SwanPayment, err := contract.NewSwanPayment(common.HexToAddress(s.PaymentContractAddress), client)
	if err != nil {
		log.Fatal("Get SwanPayment contract error", err)
	}
	var paymentParam = contract.IPaymentMinimallockPaymentParam{
		Id:         s.WCid,
		MinPayment: big.NewInt(amount),
		Amount:     big.NewInt(int64(float64(amount) * s.PayMultiplyFactor)),
		LockTime:   big.NewInt(common2.Bigint2int64(s.LockTime) * 86400),
		Recipient:  common.HexToAddress(s.PaymentRecipientAddress),
		Size:       big.NewInt(s.FileSize),
		CopyLimit:  5,
	}
	tx, err := SwanPayment.LockTokenPayment(&bind.TransactOpts{
		From:   auth.From,
		Signer: auth.Signer,
	}, paymentParam)
	if err != nil {
		log.Fatal(err)
	}

	bind.WaitMined(context.Background(), client, tx)
	tx, _, err = client.TransactionByHash(context.Background(), tx.Hash())
	if err != nil {
		log.Fatal(err)
	}
	return tx.Hash().String(), nil
}

type ContractMintNftService struct {
	c             *Client
	WalletAddress string
	PrivateKey    string
	RpcEndpoint   string
	NftMetaUrl    string
	MintAddress   string
}

func (s *ContractMintNftService) SetMintAddress(MintAddress string) *ContractMintNftService {
	s.MintAddress = MintAddress
	return s
}

func (s *ContractMintNftService) SetWalletAddress(WalletAddress string) *ContractMintNftService {
	s.WalletAddress = WalletAddress
	return s
}

func (s *ContractMintNftService) SetPrivateKey(PrivateKey string) *ContractMintNftService {
	s.PrivateKey = PrivateKey
	return s
}

func (s *ContractMintNftService) SetNftMetaUrl(NftMetaUrl string) *ContractMintNftService {
	s.NftMetaUrl = NftMetaUrl
	return s
}

func (s *ContractMintNftService) SetRpcEndpoint(RpcEndpoint string) *ContractMintNftService {
	s.RpcEndpoint = RpcEndpoint
	return s
}

func (s *ContractMintNftService) Do(ctx context.Context, opts ...RequestOption) (TX string, TokenId *big.Int, err error) {
	privateKey, err := crypto.HexToECDSA(s.PrivateKey)
	if err != nil {
		log.Fatal(err)
	}
	client, err := ethclient.Dial(s.RpcEndpoint)
	if err != nil {
		log.Fatal(err)
	}
	ChainId, _ := client.ChainID(context.Background())
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, ChainId)
	Minter, err := contract.NewContract(common.HexToAddress(s.MintAddress), client)
	if err != nil {
		log.Fatal("Get SwanPayment contract error", err)
	}
	tx, err := Minter.MintUnique(&bind.TransactOpts{
		From:   auth.From,
		Signer: auth.Signer,
	}, common.HexToAddress(s.WalletAddress), s.NftMetaUrl)
	if err != nil {
		log.Fatal(err)
	}
	bind.WaitMined(context.Background(), client, tx)
	tx, _, err = client.TransactionByHash(context.Background(), tx.Hash())
	if err != nil {
		log.Fatal(err)
	}

	TokenId, _ = Minter.IdCount(&bind.CallOpts{})
	return tx.Hash().String(), TokenId, nil
}
