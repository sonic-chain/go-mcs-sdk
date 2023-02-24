package common

import (
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/joho/godotenv"
)

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
