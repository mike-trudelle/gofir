package hash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
	"strings"
)

func GetHashValue(s string, a string) string {

	f, err := os.Open(s)

	if err != nil {
		//some error logging
	}

	defer f.Close()

	var hvalue string

	switch a {
	case "MD5":
		h := md5.New()
		if _, err := io.Copy(h, f); err != nil {
			//some error logging
		}
		hvalue = convByteArrToUcString(h.Sum(nil)[:])

	case "SHA1":
		h := sha1.New()
		if _, err := io.Copy(h, f); err != nil {
			//some error logging
		}
		hvalue = convByteArrToUcString(h.Sum(nil)[:])

	case "SHA256":
		h := sha256.New()
		if _, err := io.Copy(h, f); err != nil {
			//some error logging
		}
		hvalue = convByteArrToUcString(h.Sum(nil)[:])
	}

	return hvalue
}

func convByteArrToUcString(b []byte) string {
	return strings.ToUpper(hex.EncodeToString(b))
}
