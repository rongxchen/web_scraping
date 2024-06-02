package utils

import (
	"encoding/base64"
	"net/url"
)

func UrlEncode(data string) string {
	return url.QueryEscape(data)
}

func Base64Encode(data string) string {
	return base64.StdEncoding.EncodeToString([]byte(data))
}

func Base64Decode(data string) string {
	decoded, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return ""
	}
	return string(decoded)
}
