package mcs

import (
	"context"
	"net/http"
)

type UserRegisterService struct {
	c             *Client
	WalletAddress string
}

type UserRegisterData struct {
	Status string `json:"status"`
	Data   struct {
		Nonce string `json:"nonce"`
	} `json:"data"`
}

func (s *UserRegisterService) SetWalletAddress(WalletAddress string) *UserRegisterService {
	s.WalletAddress = WalletAddress
	return s
}
func (s *UserRegisterService) Do(ctx context.Context, opts ...RequestOption) (res *UserRegisterData, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/user/register",
	}
	r.postBody = params{
		"public_key_address": s.WalletAddress,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return &UserRegisterData{}, err
	}
	res = new(UserRegisterData)
	err = json.Unmarshal(data, res)
	if err != nil {
		return &UserRegisterData{}, err
	}

	return res, nil
}

type UserLoginService struct {
	c             *Client
	WalletAddress string
	Signature     string
	Nonce         string
	Network       string
}

type UserLoginData struct {
	Status string `json:"status"`
	Data   struct {
		JwtToken string `json:"jwt_token"`
	} `json:"data"`
}

func (s *UserLoginService) SetWalletAddress(WalletAddress string) *UserLoginService {
	s.WalletAddress = WalletAddress
	return s
}

func (s *UserLoginService) SetSignature(Signature string) *UserLoginService {
	s.Signature = Signature
	return s
}

func (s *UserLoginService) SetNonce(Nonce string) *UserLoginService {
	s.Nonce = Nonce
	return s
}

func (s *UserLoginService) SetNetwork(Network string) *UserLoginService {
	s.Network = Network
	return s
}

func (s *UserLoginService) Do(ctx context.Context, opts ...RequestOption) (res *UserLoginData, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: "/user/login_by_metamask_signature",
	}
	r.postBody = params{
		"public_key_address": s.WalletAddress,
		"nonce":              s.Nonce,
		"signature":          s.Signature,
		"network":            s.Network,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return &UserLoginData{}, err
	}
	res = new(UserLoginData)
	err = json.Unmarshal(data, res)
	if err != nil {
		return &UserLoginData{}, err
	}

	return res, nil
}
