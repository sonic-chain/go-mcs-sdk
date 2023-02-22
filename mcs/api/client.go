package api

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
