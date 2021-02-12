package godown

import (
	"fmt"
	"strings"
)

func H1(val string) string {
	return H(1, val)
}

func H2(val string) string {
	return H(2, val)
}

func H3(val string) string {
	return H(3, val)
}

func H4(val string) string {
	return H(4, val)
}

func H(n int, val string) string {
	return fmt.Sprintf(`%s %s`+"\n\n", strings.Repeat("#", n), val)
}

func anchorName(val string) string {
	val = strings.ReplaceAll(val, " ", "-")
	return strings.ToLower(val)
}
