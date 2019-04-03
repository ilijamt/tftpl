package cmd

import (
	"strings"
	"text/template"
)

var funcMap = template.FuncMap{
	"upper":   strings.ToUpper,
	"lower":   strings.ToLower,
	"replace": replace,
	"default": def,
}
