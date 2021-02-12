package godown

import "fmt"

func Link(text string, target string) string {
	return fmt.Sprintf("[%s](%s)", text, target)
}

func Ref(text string) string {
	return fmt.Sprintf("[%s](#%s)", text, anchorName(text))
}
