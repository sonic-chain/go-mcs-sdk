
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

## BucketClient
```
type BucketClient struct {
	user.McsClient
}
```

## Bucket
```
type Bucket struct {
	BucketUid  string `json:"bucket_uid"`
	Address    string `json:"address"`
	MaxSize    int64  `json:"max_size"`
	Size       int64  `json:"size"`
	IsFree     bool   `json:"is_free"`
	PaymentTx  string `json:"payment_tx"`
	IsActive   bool   `json:"is_active"`
	IsDeleted  bool   `json:"is_deleted"`
	BucketName string `json:"bucket_name"`
	FileNumber int64  `json:"file_number"`
}
```

## OssFile
```
type OssFile struct {
	Name       string `json:"name"`
	Address    string `json:"address"`
	Prefix     string `json:"prefix"`
	BucketUid  string `json:"bucket_uid"`
	FileHash   string `json:"file_hash"`
	Size       int64  `json:"size"`
	PayloadCid string `json:"payload_cid"`
	PinStatus  string `json:"pin_status"`
	IsDeleted  bool   `json:"is_deleted"`
	IsFolder   bool   `json:"is_folder"`
	ObjectName string `json:"object_name"`
	Type       int    `json:"type"`
	gorm.Model
}
```
