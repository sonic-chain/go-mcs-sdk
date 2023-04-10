
## McsClient
```
type McsClient struct {
	BaseUrl  string `json:"base_url"`
	JwtToken string `json:"jwt_token"`
}
```

## Apikey
```
type Apikey struct {
	ID          int64  `json:"id"`
	WalletId    int64  `json:"wallet_id"`
	ApiKey      string `json:"api_key"`
	AccessToken string `json:"access_token"`
	ValidDays   int32  `json:"valid_days"`
	CreateAt    int64  `json:"create_at"`
	UpdateAt    int64  `json:"update_at"`
}
```

## Wallet
```
type Wallet struct {
	ID           int64   `json:"id"`
	Address      string  `json:"address"`
	Email        *string `json:"email"`
	EmailStatus  *int    `json:"email_status"`
	EmailPopupAt *int64  `json:"email_popup_at"`
	CreateAt     int64   `json:"create_at"`
	UpdateAt     int64   `json:"update_at"`
}
```

