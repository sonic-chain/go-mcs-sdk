# go-mcs-sdk

[![Made by FilSwan](https://img.shields.io/badge/made%20by-FilSwan-green.svg)](https://www.filswan.com/)
[![Chat on discord](https://img.shields.io/badge/join%20-discord-brightgreen.svg)](https://discord.com/invite/KKGhy8ZqzK)

# Table of Contents <!-- omit in toc -->

- [Introduction](#introduction)
    - [Prerequisites](#prerequisites)
- [MCS API](#mcs-api)
- [Usage](#usage)
    - [Installation](#installation)
    - [Getting Started](#getting-started)
    - [Documentation](#documentation)
- [Contributing](#contributing)

# Introduction

A go software development kit for the Multi-Chain Storage (MCS) https://mcs.filswan.com service. It provides a
convenient interface for working with the MCS API. This SDK has the following functionalities:

- **POST** upload file to Filswan IPFS gate way
- **POST** make payment to swan filecoin storage gate way
- **POST** mint asset as NFT
- **GET** list of files uploaded
- **GET** files by cid
- **GET** status from filecoin
  Buckets Functions:
- **GET** all bucket list
- **GET** bucket info by bucket name
- **GET** bucket id by bucket name
- **GET** file id by bucket name and file name
- **PUT** create bucket
- **DELETE** bucket by bucket name
- **PUT** create upload session
- **POST** upload file to bucket

## Prequisites

Polygon Mumbai Testnet Wallet

- [Metamask Tutorial](https://docs.filswan.com/getting-started/beginner-walkthrough/public-testnet/setup-metamask) \
  Polygon Mumbai Testnet RPC - [Signup via Alchemy](https://www.alchemy.com/) \
  You will also need Testnet USDC and MATIC balance to use this
  SDK. [Swan Faucet Tutorial](https://docs.filswan.com/development-resource/swan-token-contract/acquire-testnet-usdc-and-matic-tokens)

# MCS API

For more information about the API usage, check out the MCS API
documentation (https://docs.filswan.com/development-resource/mcp-api).

# Usage

Instructions for developers working with MCS SDK and API.

## Installation

```
go mod tidy
go mod download
```

## Getting Started

### generate .env file

```
cd go-mcs-sdk/mcs
mv .env.example .env
```

### environment variable in .env file
USER_WALLET_ADDRESS_FOR_REGISTER_MCS: The wallet address used to log in to mcs <br>
USER_WALLET_ADDRESS_PK: The private key of the wallet address used to log in to mcs <br> 
CHAIN_NAME_FOR_REGISTER_ON_MCS: polygon network name: polygon.mumbai/polygon.mainnet <br>
MCS_BACKEND_BASE_URL: mcs backend url: http://127.0.0.1:8888/api/ <br>
META_SPACE_URL: meta-space backend url: http://127.0.0.1:9999/api/ <br>
### how to test meta-space
```
cd go-mcs-sdk/mcs/meta_space_test.go
```
modify the constants at the top of this file
- BucketNameForTest: the bucket name that you want to create/delete/get bucket ID
- FileNameForTest: the file name in meta-space bucket
- BucketIdForTest: the bucket ID in meta-space 

### 

## Documentation

# Contributing

## Sponsors

This project is sponsored by Filecoin Foundation

[Flink SDK - A data provider offers Chainlink Oracle service for Filecoin Network ](https://github.com/filecoin-project/devgrants/issues/463)

<img src="https://github.com/filswan/flink/blob/main/filecoin.png" width="200">
