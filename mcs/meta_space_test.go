package mcs

import (
	"log"
	"testing"
)

const (
	MetaSpaceBackendBaseUrl = "http://192.168.199.61:5212/api/"
)

func TestMetaSpaceGetBuckets(t *testing.T) {
	metaClient := NewMetaSpaceClient(MetaSpaceBackendBaseUrl)
	err := metaClient.GetToken()
	if err != nil {
		log.Println(err)
	}
	err = metaClient.GetBuckets()
	if err != nil {
		log.Println(err)
	}
}
