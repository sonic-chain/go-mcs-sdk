package mcs

import (
	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type Client struct {
	BaseURL  string
	JwtToken string
}

func (c *Client) SetJwtToken(JwtToken string) {
	c.JwtToken = JwtToken
}

func NewClient(mcsApi string) *Client {
	return &Client{
		BaseURL: mcsApi,
	}
}
