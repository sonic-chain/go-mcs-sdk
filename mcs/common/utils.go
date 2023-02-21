package common

import (
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"regexp"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/filswan/go-swan-lib/client/web"
	"github.com/filswan/go-swan-lib/logs"
	"github.com/joho/godotenv"
)

const (
	FilPriceApi = "https://api.filswan.com/stats/storage"
)

type FilPrice struct {
	Data struct {
		AverageCostPushMessage           string `json:"average_cost_push_message"`
		AverageDataCostSealing1TB        string `json:"average_data_cost_sealing_1TB"`
		AverageGasCostSealing1TB         string `json:"average_gas_cost_sealing_1TB"`
		AverageMinPieceSize              string `json:"average_min_piece_size"`
		AveragePricePerGBPerYear         string `json:"average_price_per_GB_per_year"`
		AverageVerifiedPricePerGBPerYear string `json:"average_verified_price_per_GB_per_year"`
	} `json:"data"`
	Status string `json:"status"`
}

func GetFilPrice() (float64, error) {
	response, err := web.HttpGetNoToken(FilPriceApi, nil)
	if err != nil {
		logs.GetLogger().Error(err)
		return -1, err
	}
	filPrice := new(FilPrice)
	err = json.Unmarshal(response, filPrice)
	if err != nil {
		logs.GetLogger().Error(err)
		return -1, err
	}

	price := filPrice.Data.AveragePricePerGBPerYear
	reg := regexp.MustCompile(`\d+\.\d+`)
	result := reg.FindAllStringSubmatch(price, -1)
	return strconv.ParseFloat(result[0][0], 64)
}

func GetAmount(size int64, rate float64) int64 {
	price, _ := GetFilPrice()
	sizeFloat := float64(size)
	amount := sizeFloat * price / 1024 / 1024 / 1024 * 525 / 365 * rate * 1000000000000000000
	return int64(amount)
}

func Bigint2int64(num big.Int) int64 {
	str := num.String()
	numInt, _ := strconv.ParseInt(str, 10, 64)
	return numInt
}

func PersonalSign(message string, privateKey *ecdsa.PrivateKey) (string, error) {
	fullMessage := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(message), message)
	hash := crypto.Keccak256Hash([]byte(fullMessage))
	signatureBytes, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		return "", err
	}
	signatureBytes[64] += 27
	return hexutil.Encode(signatureBytes), nil
}

func LoadEnv() error {
	err := godotenv.Load()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func GetCurrentUtcSec() (string, int64) {
	currentUtcSec := time.Now().UnixNano() / 1e9
	return strconv.FormatInt(currentUtcSec, 10), currentUtcSec
}
