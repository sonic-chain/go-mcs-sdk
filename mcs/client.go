package mcs

import (
	"bytes"
	"context"
	"fmt"
	"github.com/bitly/go-simplejson"
	jsoniter "github.com/json-iterator/go"
	"go-mcs-sdk/mcs/common"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	restApiVersion = "v1"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func newJSON(data []byte) (j *simplejson.Json, err error) {
	j, err = simplejson.NewJson(data)
	if err != nil {
		return nil, err
	}
	return j, nil
}

func (c *Client) debug(format string, v ...interface{}) {
	if c.Debug {
		c.Logger.Printf(format, v...)
	}
}

type doFunc func(req *http.Request) (*http.Response, error)

type Client struct {
	BaseURL    string
	UserAgent  string
	HTTPClient *http.Client
	Debug      bool
	Logger     *log.Logger
	TimeOffset int64
	do         doFunc
	JwtToken   string
}

func (c *Client) SetJwtToken(JwtToken string) {
	c.JwtToken = JwtToken
}

func NewClient(mcsApi string) *Client {
	return &Client{
		BaseURL:    mcsApi,
		UserAgent:  "mcs/go",
		HTTPClient: http.DefaultClient,
		Logger:     log.New(os.Stderr, "mcs-go ", log.LstdFlags),
		Debug:      true,
	}
}

func (c *Client) parseRequest(r *request, opts ...RequestOption) (err error) {
	// set request options from user
	for _, opt := range opts {
		opt(r)
	}
	err = r.validate()
	if err != nil {
		return err
	}

	fullURL := fmt.Sprintf("%s%s", c.BaseURL, r.endpoint)
	queryString := r.query.Encode()
	body := &bytes.Buffer{}

	if r.postBody != nil {
		bodyByte, _ := jsoniter.Marshal(r.postBody)
		body = bytes.NewBuffer(bodyByte)
		c.debug("request post body: %#v", r.postBody)
	}
	header := http.Header{}

	if r.header != nil {
		header = r.header.Clone()
	}
	header.Set("Authorization", "Bearer "+c.JwtToken)

	if queryString != "" {
		fullURL = fmt.Sprintf("%s?%s", fullURL, queryString)
	}
	r.fullURL = fullURL
	r.header = header

	if r.body == nil {
		r.body = body
	}

	return nil
}

func (c *Client) callAPI(ctx context.Context, r *request, opts ...RequestOption) (data []byte, err error) {
	err = c.parseRequest(r, opts...)
	if err != nil {
		return []byte{}, err
	}

	req, err := http.NewRequest(r.method, r.fullURL, r.body)
	if err != nil {
		return []byte{}, err
	}
	req = req.WithContext(ctx)
	req.Header = r.header
	c.debug("request: %#v", req)
	f := c.do
	if f == nil {
		f = c.HTTPClient.Do
	}
	res, err := f(req)
	if err != nil {
		return []byte{}, err
	}
	data, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}
	defer func() {
		cerr := res.Body.Close()
		// Only overwrite the retured error if the original error was nil and an
		// error occurred while closing the body.
		if err == nil && cerr != nil {
			err = cerr
		}
	}()
	c.debug("response: %#v", res)
	c.debug("response body: %s", string(data))
	c.debug("response status code: %d", res.StatusCode)
	//fmt.Println(res.StatusCode)
	if res.StatusCode >= http.StatusBadRequest {
		apiErr := new(common.APIError)
		e := json.Unmarshal(data, apiErr)
		if e != nil {
			c.debug("failed to unmarshal json: %s", e)
		}
		return nil, apiErr
	}
	return data, nil
}

func (c *Client) NewGetParamsService() *GetParamsService {
	return &GetParamsService{c: c}
}

func (c *Client) NewGetPriceRateService() *GetPriceRateService {
	return &GetPriceRateService{c: c}
}

func (c *Client) NewGetUserTaskDealsService() *GetUserTaskDealsService {
	return &GetUserTaskDealsService{c: c}
}

func (c *Client) NewGetPaymentInfoService() *GetPaymentInfoService {
	return &GetPaymentInfoService{c: c}
}

func (c *Client) NewGetMintInfoService() *GetMintInfoService {
	return &GetMintInfoService{c: c}
}

func (c *Client) NewGetDealDetailService() *GetDealDetailService {
	return &GetDealDetailService{c: c}
}

func (c *Client) NewUploadIpfsService() *UploadIpfsService {
	return &UploadIpfsService{c: c}
}

func (c *Client) NewUploadNftMetadataService() *UploadNftMetadataService {
	return &UploadNftMetadataService{c: c}
}

func (c *Client) NewContractContractApproveUSDCService() *ContractApproveUSDCService {
	return &ContractApproveUSDCService{c: c}
}

func (c *Client) NewContractUploadFilePayService() *ContractUploadFilePayService {
	return &ContractUploadFilePayService{c: c}
}

func (c *Client) NewContractMintNftService() *ContractMintNftService {
	return &ContractMintNftService{c: c}
}

func (c *Client) NewUserRegisterService() *UserRegisterService {
	return &UserRegisterService{c: c}
}

func (c *Client) NewUserLoginService() *UserLoginService {
	return &UserLoginService{c: c}
}
