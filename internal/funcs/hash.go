package funcs

import (
	"crypto/md5"
	"encoding/hex"
)

func HashMD5(input string) string {
	m := md5.New()
	m.Write([]byte(input))
	return hex.EncodeToString(m.Sum(nil))
}
