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
	"os"
	"path/filepath"
	"strings"
	"time"
)

const HTTP_CONTENT_TYPE_FORM = "application/x-www-form-urlencoded"
const HTTP_CONTENT_TYPE_JSON = "application/json; charset=UTF-8"

func HttpPostNoToken(uri string, params interface{}) ([]byte, error) {
	response, err := HttpRequest(http.MethodPost, uri, "", params, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return response, nil
}

func HttpPost(uri, tokenString string, params interface{}) ([]byte, error) {
	response, err := HttpRequest(http.MethodPost, uri, tokenString, params, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return response, nil
}

func HttpGetNoToken(uri string, params interface{}) ([]byte, error) {
	response, err := HttpRequest(http.MethodGet, uri, "", params, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return response, nil
}

func HttpGetNoTokenTimeout(uri string, params interface{}, timeoutSecond *int) ([]byte, error) {
	response, err := HttpRequest(http.MethodGet, uri, "", params, timeoutSecond)
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

func HttpPut(uri, tokenString string, params interface{}) ([]byte, error) {
	response, err := HttpRequest(http.MethodPut, uri, tokenString, params, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return response, nil
}

func HttpDelete(uri, tokenString string, params interface{}) ([]byte, error) {
	response, err := HttpRequest(http.MethodDelete, uri, tokenString, params, nil)
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

func HttpUploadFileByStream(uri, filefullpath string) ([]byte, error) {
	fileReader, err := os.Open(filefullpath)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	filename := filepath.Base(filefullpath)

	boundary := "MyMultiPartBoundary12345"
	token := "DEPLOY_GATE_TOKEN"
	message := "Uploaded by Nebula"
	releaseNote := "Built by Nebula"
	fieldFormat := "--%s\r\nContent-Disposition: form-data; name=\"%s\"\r\n\r\n%s\r\n"
	tokenPart := fmt.Sprintf(fieldFormat, boundary, "token", token)
	messagePart := fmt.Sprintf(fieldFormat, boundary, "message", message)
	releaseNotePart := fmt.Sprintf(fieldFormat, boundary, "release_note", releaseNote)
	fileName := filename
	fileHeader := "Content-type: application/octet-stream"
	fileFormat := "--%s\r\nContent-Disposition: form-data; name=\"file\"; filename=\"%s\"\r\n%s\r\n\r\n"
	filePart := fmt.Sprintf(fileFormat, boundary, fileName, fileHeader)
	bodyTop := fmt.Sprintf("%s%s%s%s", tokenPart, messagePart, releaseNotePart, filePart)
	bodyBottom := fmt.Sprintf("\r\n--%s--\r\n", boundary)
	body := io.MultiReader(strings.NewReader(bodyTop), fileReader, strings.NewReader(bodyBottom))

	contentType := fmt.Sprintf("multipart/form-data; boundary=%s", boundary)

	response, err := http.Post(uri, contentType, body)
	if err != nil {
		log.Println(err)
		return nil, nil
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		err := fmt.Errorf("http status:%s, code:%d, url:%s", response.Status, response.StatusCode, uri)
		log.Println(err)
		switch response.StatusCode {
		case http.StatusNotFound:
			log.Println("please check your url:", uri)
		}
		return nil, err
	}

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	responseStr := string(responseBody)
	//logs.GetLogger().Info(responseStr)
	filesInfo := strings.Split(responseStr, "\n")
	if len(filesInfo) < 4 {
		err := fmt.Errorf("not enough files info returned, ipfs response:%s", responseStr)
		log.Println(err)
		return nil, err
	}
	responseStr = filesInfo[3]
	return []byte(responseStr), nil
}
