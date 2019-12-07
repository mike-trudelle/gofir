package hash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"hash/crc32"
	"io"
	"os"
	"strings"
)

// GetHashValue gets the hash of a string using a chosen algorithm
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

	case "CRC32":
		tablePolynomial := crc32.MakeTable(0xedb88320)
		h := crc32.New(tablePolynomial)
		if _, err := io.Copy(h, f); err != nil {
			//some error logging
		}
		hvalue = convByteArrToUcString(h.Sum(nil)[:])

	default:
		h := md5.New()
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
