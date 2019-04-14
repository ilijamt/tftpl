package cmd

import (
	"github.com/ilijamt/tftpl/internal/funcs"
	"strings"
	"text/template"
)

var funcMap = template.FuncMap{
	"upper":   strings.ToUpper,
	"lower":   strings.ToLower,
	"replace": funcs.Replace,
	"default": funcs.Default,
	"md5":     funcs.HashMD5,
}
