package api

type BucketClient struct {
	McsClient
}

func NewBucketClient() *BucketClient {
	bucketClient := BucketClient{}
	return &bucketClient
}
