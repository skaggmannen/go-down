package godown

import (
	"strings"
)

func Doc(items ...string) string {
	return strings.Join(items, "")
}
