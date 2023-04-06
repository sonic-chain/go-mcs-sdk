
## GetOnChainClient

Definition:

```shell
func GetOnChainClient(mcsClient user.McsClient) *OnChainClient
```

Outputs:

```shell
*OnChainClient  # includes jwt token and other information for use when call the other apis
```
## GetSystemParam

Definition:

```shell
func (onChainClient *OnChainClient) GetSystemParam() (*SystemParam, error)
```

Outputs:

```shell
*SystemParam  # system parameters
error         # error or nil
```

## GetHistoricalAveragePriceVerified

Definition:

```shell
func GetHistoricalAveragePriceVerified() (float64, error)
```

Outputs:

```shell
float64    # historical average verified price
error      # error or nil
```

## GetAmount

Definition:

```shell
func GetAmount(fizeSizeByte int64, historicalAveragePriceVerified, fileCoinPrice float64, copyNumber int) (int64, error)
```

Outputs:

```shell
int64    # amount to pay
error    # error or nil
```

## Upload

Definition:

```shell
func (onChainClient *OnChainClient) Upload(filePath string, fileType int) (*UploadFile, error)
```

Outputs:

```shell
*UploadFile # upload file information
error       # error or nil
```

## GetMintInfo

Definition:

```shell
func (onChainClient *OnChainClient) GetMintInfo(sourceFileUploadId int64) ([]*SourceFileMintOut, error)
```

Outputs:

```shell
*SourceFileMintOut # file mint info
error              # error or nil
```

## GetUserTaskDeals

Definition:

```shell
func (onChainClient *OnChainClient) GetUserTaskDeals(dealsParams DealsParams) ([]*Deal, *int64, error)
```

Outputs:

```shell
[]*Deal # deal list
*int64  # total count
error   # error or nil
```

## GetDealDetail

Definition:

```shell
func (onChainClient *OnChainClient) GetDealDetail(sourceFileUploadId, dealId int64) (*SourceFileUploadDeal, []*DaoSignature, *int, error)
```

Outputs:

```shell
*SourceFileUploadDeal # deal list
[]*DaoSignature       # dao signature list
*int                  # dao threshold
error                 # error or nil
```

## GetDealLogs

Definition:

```shell
func (onChainClient *OnChainClient) GetDealLogs(offlineDealId int64) ([]*OfflineDealLog, error)
```

Outputs:

```shell
[]*OfflineDealLog  # deal logs
error              # error or nil
```

## GetSourceFileUpload

Definition:

```shell
func (onChainClient *OnChainClient) GetSourceFileUpload(sourceFileUploadId int64) (*SourceFileUpload, error)
```

Outputs:

```shell
*SourceFileUpload  # source file upload information
error              # error or nil
```

## UnpinSourceFile

Definition:

```shell
func (onChainClient *OnChainClient) UnpinSourceFile(sourceFileUploadId int64) error
```

Outputs:

```shell
error              # error or nil
```

## WriteNftCollection

Definition:

```shell
func (onChainClient *OnChainClient) WriteNftCollection(nftCollectionParams NftCollectionParams) error
```

Outputs:

```shell
error              # error or nil
```

## GetNftCollections

Definition:

```shell
func (onChainClient *OnChainClient) GetNftCollections() ([]*NftCollection, error)
```

Outputs:

```shell
[]*NftCollection   # NFT collections
error              # error or nil
```

## RecordMintInfo

Definition:

```shell
func (onChainClient *OnChainClient) RecordMintInfo(recordMintInfoParams *RecordMintInfoParams) (*SourceFileMint, error)
```

Outputs:

```shell
*SourceFileMint   # Mint info
error             # error or nil
```












