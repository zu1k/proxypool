package tool

import (
	"encoding/base64"
)

func Base64DecodeString(src string) (dst string, err error) {
	if src == "" {
		return "", nil
	}

	origin := src
	//if i := len(src) % 4; i != 0 {
	//	for k := 0; k < i; k++ {
	//		src += string(base64.StdPadding)
	//	}
	//}
	var dstbytes []byte
	dstbytes, err = base64.RawURLEncoding.DecodeString(origin)

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
