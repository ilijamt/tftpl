package funcs

import (
	"strings"
)

// Default returns a default value in a template when used, if the input is empty, the default value will be returned
//
//   {{ default <input> <default_value>
//
func Default(input interface{}, def interface{}) interface{} {
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
