package godown

import (
	"strings"
)

func P(val string) string {
	return fixIndent(strings.TrimSpace(val), "") + "\n\n"
}
