package tool

import (
	"encoding/base64"
)

func Base64DecodeString(src string) (dst string, err error) {
	if src == "" {
		return "", nil
	}
	var dstbytes []byte
	dstbytes, err = base64.RawURLEncoding.DecodeString(src)

	if err != nil {
		dstbytes, err = base64.RawStdEncoding.DecodeString(src)
	}
	if err != nil {
		dstbytes, err = base64.StdEncoding.DecodeString(src)
	}
	if err != nil {
		dstbytes, err = base64.URLEncoding.DecodeString(src)
	}
	if err != nil {
		return "", err
	}
	dst = string(dstbytes)
	return
}

func Base64EncodeString(origin string, urlsafe bool) (result string) {
	if urlsafe {
		return base64.URLEncoding.EncodeToString([]byte(origin))
	}
	return base64.StdEncoding.EncodeToString([]byte(origin))
}

func Base64EncodeBytes(origin []byte) (result string) {
	return base64.StdEncoding.EncodeToString([]byte(origin))
}
