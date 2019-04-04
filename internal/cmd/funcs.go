package cmd

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func replace(input, to string, from ...string) (result string) {
	result = input
	for _, i := range from {
		result = strings.ReplaceAll(result, i, to)
	}
	return
}

func def(input interface{}, def interface{}) interface{} {
	switch val := input.(type) {
	case string:
		val = strings.TrimSpace(val)
		if len(val) > 0 {
			return val
		}
	default:
		if val != nil {
			return val
		}
	}
	return def
}

func hash_md5(input string) string {
	m := md5.New()
	m.Write([]byte(input))
	return hex.EncodeToString(m.Sum(nil))
}