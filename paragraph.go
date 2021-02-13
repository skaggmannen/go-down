package godown

import (
	"strings"
)

func P(items ...string) string {
	val := strings.Join(items, "")
	return fixIndent(strings.TrimSpace(val), "") + "\n\n"
}
