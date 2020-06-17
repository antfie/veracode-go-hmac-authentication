# Veracode HMAC Authentication Example for Go

A Go version of the Veracode HMAC authentication examples as found here: https://help.veracode.com/reader/LMv_dtSHyb7iIxAQznC~9w/CUv4heF9z9tOBnZ1uiB8UA.

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

	authorizationHeader := hmac.CalculateAuthorizationHeader(parsedUrl, httpMethod, apiKeyID, apiKeySecret)
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
