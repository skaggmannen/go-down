package godown

import "strings"

func Em(items ...string) string {
	return "_" + strings.Join(items, "") + "_"
}

func Strong(items ...string) string {
	return "**" + strings.Join(items, "") + "**"
}

func Strike(items ...string) string {
	return "~~" + strings.Join(items, "") + "~~"
}

func Pre(items ...string) string {
	return "`" + strings.Join(items, "") + "`"
}
