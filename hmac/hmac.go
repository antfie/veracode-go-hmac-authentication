package hmac

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/url"
	"strconv"
	"time"
)

// Included in the signature to inform Veracode of the signature version.
const veracodeRequestVersionString = "vcode_request_version_1"

// Expected format for the unencrypted data string.
const dataFormat = "id=%s&host=%s&url=%s&method=%s"

// Expected format for the Authorization header.
const headerFormat = "%s id=%s,ts=%s,nonce=%X,sig=%X"

// Expect prefix to the Authorization header.
const veracodeHMACSHA256 = "VERACODE-HMAC-SHA-256"

// CalculateAuthorizationHeader produces the value to be used with the Authorization HTTP Request header when making Veracode API calls.
func CalculateAuthorizationHeader(url *url.URL, httpMethod, apiKeyID, apiKeySecret string) string {
	nonce := createNonce(16)
	timestampMilliseconds := strconv.FormatInt(time.Now().UnixNano()/int64(1000000), 10)
	data := fmt.Sprintf(dataFormat, apiKeyID, url.Hostname(), url.RequestURI(), httpMethod)
	dataSignature := calculateSignature(fromHexString(apiKeySecret), nonce, []byte(timestampMilliseconds), []byte(data))
	return fmt.Sprintf(headerFormat, veracodeHMACSHA256, apiKeyID, timestampMilliseconds, nonce, dataSignature)
}

func createNonce(size int) []byte {
	nonce := make([]byte, size)

	_, err := rand.Read(nonce)

	if err != nil {
		panic(err)
	}

	return nonce
}

func fromHexString(input string) []byte {
	result, err := hex.DecodeString(input)

	if err != nil {
		panic(err)
	}

	return result
}

func calculateSignature(key, nonce, timestamp, data []byte) []byte {
	encryptedNonce := hmac256(nonce, key)
	encryptedTimestampMilliseconds := hmac256(timestamp, encryptedNonce)
	signingKey := hmac256([]byte(veracodeRequestVersionString), encryptedTimestampMilliseconds)
	return hmac256(data, signingKey)
}

func hmac256(message, key []byte) []byte {
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	return mac.Sum(nil)
}
