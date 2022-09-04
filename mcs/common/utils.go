package common

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"regexp"
	"strconv"
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
	resp, _ := http.Get(FilPriceApi)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode == 200 {
		fmt.Println("ok")
	}
	res := new(FilPrice)
	_ = json.Unmarshal(body, res)
	price := res.Data.AveragePricePerGBPerYear
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
