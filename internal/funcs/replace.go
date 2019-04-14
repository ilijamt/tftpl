package funcs

import (
	"strings"
)

func Replace(input, to string, from ...string) (result string) {
	result = input
	for _, i := range from {
		result = strings.ReplaceAll(result, i, to)
	}
	return
}
