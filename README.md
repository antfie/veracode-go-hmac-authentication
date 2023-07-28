[![Go Report Card](https://goreportcard.com/badge/github.com/antfie/veracode-go-hmac-authentication)](https://goreportcard.com/report/github.com/antfie/veracode-go-hmac-authentication) [![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/antfie/veracode-go-hmac-authentication/blob/master/LICENSE)

# Veracode HMAC Authentication for Go

A Go version of the Veracode HMAC authentication as found in [Veracode Docs](https://docs.veracode.com/r/c_hmac_signing_example). This library has no dependencies.

## Installation

```
go get -u github.com/antfie/veracode-go-hmac-authentication
```

## Example Usage

```go
package main

import (
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/antfie/veracode-go-hmac-authentication/hmac"
)

func main() {
	var apiKeyID = "YOUR_VERACODE_API_KEY_ID"
	var apiKeySecret = "YOUR_VERACODE_API_KEY_SECRET"
	var apiUrl = "https://analysiscenter.veracode.com/api/5.0/getapplist.do"

	response := makeApiRequest(apiKeyID, apiKeySecret, apiUrl, http.MethodGet)
	print(response)
}

func makeApiRequest(apiKeyID, apiKeySecret, apiUrl, httpMethod string) string {
	parsedUrl, err := url.Parse(apiUrl)

	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	req, err := http.NewRequest(httpMethod, parsedUrl.String(), nil)

	if err != nil {
		panic(err)
	}

	authorizationHeader, err := hmac.CalculateAuthorizationHeader(parsedUrl, httpMethod, apiKeyID, apiKeySecret)

	if err != nil {
		panic(err)
	}

	req.Header.Add("Authorization", authorizationHeader)

	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	if resp.StatusCode != http.StatusOK {
		panic("Expected status 200. Status was: " + resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	return string(body[:])
}
```
