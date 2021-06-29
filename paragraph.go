package godown

import (
	"regexp"
	"strings"
)

func P(items ...string) string {
	val := strings.Join(items, "")
	re, _ := regexp.Compile(`\s+`)
	val = string(re.ReplaceAll([]byte(val), []byte(" ")))

	strings.
}

func wrap(text string) string {
	text = strings.ReplaceAll(text, "\n", " ")
}
