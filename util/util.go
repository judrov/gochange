// Package util contains utility functions for creating HTTP request.
package util

// Imports required packages.
import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// Hash computes hash over parameters
func Hash(dataStr string) string {
	data := []byte(dataStr)
	hash := sha256.New()
	hash.Write(data)
	md5 := hash.Sum(nil)
	md5Str := hex.EncodeToString(md5)
	return md5Str
}

// GetTimeNow returns current time using UTC in ISO 8601 format
func GetTimeNow() string {
	const ISO8601_UTC = "2006-01-02T15:04:05Z"
	timestamp := time.Now().UTC()
	timestampStr := timestamp.Format(ISO8601_UTC)
	return timestampStr
}

// Post creates and sends a new HTTP post request
func Post(url, params string, bodyRes interface{}) error {
	req, err := http.NewRequest("POST", url, strings.NewReader(params))
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if err := Unmarshal(res, bodyRes); err != nil {
		return err
	}
	return nil
}

// Unmarshal parses JSON-encoded data and stores the result to Response.
func Unmarshal(res *http.Response, bodyRes interface{}) error {
	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, &bodyRes); err != nil {
		return err
	}
	return nil
}
