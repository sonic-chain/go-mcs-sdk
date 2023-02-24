package common

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

const HTTP_CONTENT_TYPE_FORM = "application/x-www-form-urlencoded"
const HTTP_CONTENT_TYPE_JSON = "application/json; charset=UTF-8"

func HttpPost(uri, tokenString string, params interface{}) ([]byte, error) {
	response, err := HttpRequest(http.MethodPost, uri, tokenString, params, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return response, nil
}

func HttpGet(uri, tokenString string, params interface{}) ([]byte, error) {
	response, err := HttpRequest(http.MethodGet, uri, tokenString, params, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return response, nil
}

func HttpRequest(httpMethod, uri, tokenString string, params interface{}, timeoutSecond *int) ([]byte, error) {
	var request *http.Request
	var err error

	switch params := params.(type) {
	case io.Reader:
		request, err = http.NewRequest(httpMethod, uri, params)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		request.Header.Set("Content-Type", HTTP_CONTENT_TYPE_FORM)
	default:
		jsonReq, errJson := json.Marshal(params)
		if errJson != nil {
			log.Println(errJson)
			return nil, errJson
		}

		request, err = http.NewRequest(httpMethod, uri, bytes.NewBuffer(jsonReq))
		if err != nil {
			log.Println(err)
			return nil, err
		}
		request.Header.Set("Content-Type", HTTP_CONTENT_TYPE_JSON)
	}

	if len(strings.Trim(tokenString, " ")) > 0 {
		request.Header.Set("Authorization", "Bearer "+tokenString)
	}

	customTransport := http.DefaultTransport.(*http.Transport).Clone()
	customTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	client := &http.Client{Transport: customTransport}
	if timeoutSecond != nil {
		client.Timeout = time.Duration(*timeoutSecond) * time.Second
	}

	response, err := client.Do(request)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		err := fmt.Errorf("http status: %s, code:%d, url:%s", response.Status, response.StatusCode, uri)
		log.Println(err)
		switch response.StatusCode {
		case http.StatusNotFound:
			log.Println("please check your url:", uri)
		case http.StatusUnauthorized:
			log.Println("Please check your token:", tokenString)
		}
		return nil, err
	}

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return responseBody, nil
}
