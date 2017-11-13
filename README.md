# Dialex SDK for Go
Go Client for HyperLab Dialex API  

[![Build Status](https://travis-ci.org/HyperLab-Solutions-Sdn-Bhd/dialex-sdk-go.svg?branch=master)](https://travis-ci.org/HyperLab-Solutions-Sdn-Bhd/dialex-sdk-go) [![Coverage Status](https://coveralls.io/repos/github/HyperLab-Solutions-Sdn-Bhd/dialex-sdk-go/badge.svg?branch=master)](https://coveralls.io/github/HyperLab-Solutions-Sdn-Bhd/dialex-sdk-go?branch=master)  

## Install  
```bash
go get https://github.com/HyperLab-Solutions-Sdn-Bhd/dialex-sdk-go/dialex
```  

## Quick Start  
```go
import "https://github.com/HyperLab-Solutions-Sdn-Bhd/dialex-sdk-go/dialex"

dial := dialex.New("apikey")
```  

## Developing  

Clone using git:  
`git clone git@github.com:HyperLab-Solutions-Sdn-Bhd/dialex-sdk-go.git`  

Install glide:  
`curl https://glide.sh/get | sh`  

Install dependencies:  
`glide install`  

Run tests:  
`go test -v`  
