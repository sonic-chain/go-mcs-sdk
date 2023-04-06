
## GetBucketClient

Definition:

```shell
func GetBucketClient(mcsClient user.McsClient) *BucketClient
```

Outputs:

```shell
*BucketClient  # includes jwt token and other information for use when call the other apis
```

## ListBuckets

Definition:

```shell
func (bucketClient *BucketClient) ListBuckets() ([]*Bucket, error)
```

Outputs:

```shell
[]*Bucket  # the bucket list belong to current user
error      # error or nil
```

## CreateBucket

Definition:

```shell
func (bucketClient *BucketClient) CreateBucket(bucketName string) (*string, error)
```

Outputs:

```shell
*string  # bucket uid
error    # error or nil
```

## DeleteBucket

Definition:

```shell
func (bucketClient *BucketClient) DeleteBucket(bucketName string) error
```

Outputs:

```shell
error    # error or nil
```

## GetBucket

Definition:

```shell
func (bucketClient *BucketClient) GetBucket(bucketName, bucketUid string) (*Bucket, error)
```

Outputs:

```shell
*Bucket  # bucket info whose bucket name or bucket uid is the same as the parameter
error    # error or nil
```

## GetBucketUid

Definition:

```shell
func (bucketClient *BucketClient) GetBucketUid(bucketName string) (*string, error)
```

Outputs:

```shell
*string  # bucket uid whose bucket name is the same as the parameter
error    # error or nil
```

## RenameBucket

Definition:

```shell
func (bucketClient *BucketClient) RenameBucket(newBucketName string, bucketUid string) error
```

Outputs:

```shell
error    # error or nil
```

## GetTotalStorageSize

Definition:

```shell
func (bucketClient *BucketClient) GetTotalStorageSize() (*int64, error)
```

Outputs:

```shell
*int64   # total storage size
error    # error or nil
```





