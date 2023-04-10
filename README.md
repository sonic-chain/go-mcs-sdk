# go-mcs-sdk

[![Made by FilSwan](https://img.shields.io/badge/made%20by-FilSwan-green.svg)](https://www.filswan.com/)
[![Chat on discord](https://img.shields.io/badge/join%20-discord-brightgreen.svg)](https://discord.com/invite/KKGhy8ZqzK)

# Table of Contents <!-- omit in toc -->

- [Introduction](#introduction)
    - [Prerequisites](#prerequisites)
- [Usage](#usage)
    - [Installation](#installation)
    - [Getting Started](#getting-started)
- [MCS API](#mcs-api)
- [Documentation](#documentation)
- [Contributing](#contributing)
- [Sponsors](#Sponsors)

# Introduction

A go software development kit for the Multi-Chain Storage (MCS) https://mcs.filswan.com service. It provides a
convenient interface for working with the MCS API. This SDK has the following functionalities:

**Buckets Functions:**
- List all buckets
- Create a bucket
- Delete a bucket
- Get a bucket by bucket name or UID
- Get bucket Uid by its name
- Rename a bucket
- Get total storage size

On-chain files Functions:
- **POST** upload file to Filswan IPFS gate way
- **POST** make payment to swan filecoin storage gate way
- **POST** mint asset as NFT
- **GET** list of files uploaded
- **GET** files by cid
- **GET** status from filecoin

## Prequisites

Polygon Mumbai Testnet Wallet

- [Metamask Tutorial](https://docs.filswan.com/getting-started/beginner-walkthrough/public-testnet/setup-metamask) \
  Polygon Mumbai Testnet RPC - [Signup via Alchemy](https://www.alchemy.com/) \
  You will also need Testnet USDC and MATIC balance to use this
  SDK. [Swan Faucet Tutorial](https://docs.filswan.com/development-resource/swan-token-contract/acquire-testnet-usdc-and-matic-tokens)


# Usage

Instructions for developers working with MCS SDK and API.

## Installation

```
go mod tidy
go mod download
```

## Getting Started
### How to test 
Run the test method in **_test.go and enter the parameters <br>
The parameters are placed at the top of the test script file as constants <br>
List of test scripts:
```
go-mcs-sdk/mcs/bucket_api_test.go
go-mcs-sdk/mcs/mcs_api_test.go
```
#### Run the test method, take the test upload file method as an example:
1. Set the full path of the file you want to upload to the constant at the top of the test file <br>
   FilePathForUpload = <your file full path>
2. Run the upload file test api   <br>
   go test -v -run TestMcsUploadFile

## Documentation

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