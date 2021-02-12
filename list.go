package godown

import (
	"bytes"
	"fmt"
	"strings"
)

func Ol(items ...string) string {
	var buf bytes.Buffer
	prefix := strings.Repeat(" ", len(fmt.Sprintf("%d. ", len(items))))
	for i, item := range items {
		_, _ = fmt.Fprintln(&buf, fmt.Sprintf("%d. %s", i+1, Indent(item, prefix)))
	}
	return buf.String() + "\n"
}

// ===================================================================================

func Ul(items ...string) string {
	var buf bytes.Buffer
	for _, item := range items {
		_, _ = fmt.Fprintln(&buf, fmt.Sprintf("* %s", Indent(item, "  ")))
	}
	return buf.String() + "\n"
}

// ===================================================================================

func Li(items ...string) string {
	return strings.Join(items, "")
}
