package mcs

import (
	"fmt"
	"testing"
)

func TestMCSUpload(t *testing.T) {
	up := MCSUpload{ChainName: "polygon.mumbai", WalletAddress: WalletAddress, PrivateKey: PrivateKey, RpcEndpoint: RpcEndpoint, FilePath: FilePath}
	NewMCSUpload(&up)
	fmt.Println(up.UploadIpfsData)
	fmt.Println(up.token)
	NewMcsMintNft(&up)
	fmt.Println(up.MintInfo)
}
