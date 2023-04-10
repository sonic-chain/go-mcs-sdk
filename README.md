# go-mcs-sdk

[![Made by FilSwan](https://img.shields.io/badge/made%20by-FilSwan-green.svg)](https://www.filswan.com/)
[![Chat on discord](https://img.shields.io/badge/join%20-discord-brightgreen.svg)](https://discord.com/invite/KKGhy8ZqzK)

# Table of Contents <!-- omit in toc -->

- [Introduction](#introduction)
    - [Functions](#Functions)
    - [Data Structures](#Data-Structures)
    - [Constants](#Constants)
- [Prerequisites](#Prerequisites)
- [Usage](#usage)
    - [Download SDK](#Download-SDK)
    - [Call SDK](#Call-SDK)
    - [Documentation](#documentation)
- [MCS API](#mcs-api)
- [Contributing](#contributing)
- [Sponsors](#Sponsors)

## Introduction

A Golang software development kit for the [Multi-Chain Storage (MCS) Service](https://mcs.filswan.com) . It provides a
convenient interface for working with the MCS API. 

### Functions:

- [User Functions](https://github.com/filswan/go-mcs-sdk/blob/dev/mcs/api/docs/user.md)
- [Bucket Functions](https://github.com/filswan/go-mcs-sdk/blob/dev/mcs/api/docs/bucket.md)
- [On-chain Functions](https://github.com/filswan/go-mcs-sdk/blob/dev/mcs/api/docs/on-chain.md)

### Data Structures:
- [Struct](https://github.com/filswan/go-mcs-sdk/blob/dev/mcs/api/docs/struct.md)

### Constants:
- [Constants](https://github.com/filswan/go-mcs-sdk/blob/dev/mcs/api/common/constants/constants.go)

## Prerequisites
- [Metamask Wallet](https://docs.filswan.com/getting-started/beginner-walkthrough/public-testnet/setup-metamask)
- [Polygon Mumbai Testnet RPC](https://www.alchemy.com/)
- [Testnet USDC and MATIC balance](https://docs.filswan.com/development-resource/swan-token-contract/acquire-testnet-usdc-and-matic-tokens)
- [Optional: apikey](https://calibration-mcs.filswan.com/) -> Setting -> Create API Key

## Usage

### Download SDK
```
go get go-mcs-sdk
```


### Call SDK
1. Login using either of the below ways:
```
mcsClient, err := LoginByApikey(apikey, accessToken, network)
apikey: your apikey
accessToken: the access token for your apikey
network: defined in constants

mcsClient: result including the information to access the other API(s)
err: when err generated while accessing this api, the error info will store in err
```
```
step 1.
nonce, err := Register(publicKeyAddress, network)
publicKeyAddress: your wallet public key address
network: defined in constants

nonce: MCS generated nonce for the related parameters
err: when err generated while accessing this api, the error info will store in err

step 2.
mcsClient, err := LoginByPublicKeySignature(nonce, publicKeyAddress, publicKeySignature, network)

nonce: MCS generated nonce from last step
publicKeyAddress: your wallet public key address
publicKeySignature: public key signature generated from meta mask wallet
network: defined in constants

mcsClient: result including the information to access the other API(s)
err: when err generated while accessing this api, the error info will store in err
```
- See [Constants](#Constants) to show optional network
- You can get the public key signature from [Public Key Signature](https://ibuxj.csb.app/)

2. Call `user` related api(s) using `mcsClient` got from last step, such as:
```
wallet, err := mcsClient.GetWallet()
wallet: the wallet that the apikey belong to
err: when err generated while accessing this api, the error info will store in err
```
3. If you want to call `bucket` related api(s), you need change `McsClient` to `BucketClient` first:
```
buketClient := GetBucketClient(*mcsClient)
```
then call `bucket` related api(s) using `buketClient` got from above, such as:
```
buckets, err := buketClient.ListBuckets()
buckets: bucket list
err: when err generated while accessing this api, the error info will store in err
```
4. If you want to call `on-chain` related api(s), you need change `McsClient` to `OnChainClient` first:
```
onChainClient = GetOnChainClient(*mcsClient)
```
then call `on-chain` related api(s) using `onChainClient` got from above, such as:
```
filecoinPrice, err := onChainClient.GetFileCoinPrice()
filecoinPrice: filecoin price
err: when err generated while accessing this api, the error info will store in err
```

### Documentation

For more examples please see the [SDK documentation](https://docs.filswan.com/multi-chain-storage/developer-quickstart/sdk)

## MCS API

For more information about the API usage, check out the MCS API
documentation (https://docs.filswan.com/development-resource/mcp-api).

## Contributing

Feel free to join in and discuss. Suggestions are welcome! [Open an issue](https://github.com/filswan/python-mcs-sdk/issues) or [Join the Discord](https://discord.com/invite/KKGhy8ZqzK)!

## Sponsors

This project is sponsored by Filecoin Foundation

[Flink SDK - A data provider offers Chainlink Oracle service for Filecoin Network ](https://github.com/filecoin-project/devgrants/issues/463)

<img src="https://github.com/filswan/flink/blob/main/filecoin.png" width="200">