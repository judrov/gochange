package changeorg

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"strings"
	"time"
)

// Post creates and sends a new HTTP post request
func Post(url, params string, bodyRes *Response) error {
	req, err := http.NewRequest("POST", url, strings.NewReader(params))
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if err := unmarshal(res, bodyRes); err != nil {
		return err
	}
	return nil
}

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
	fmt.Println("Current time: " + timestampStr)
	return timestampStr
}
