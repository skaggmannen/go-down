package godown

import (
	"fmt"
	"strings"
)

func Code(lang string, text string) string {
	return fmt.Sprintf("```%s\n%s\n```", lang, strings.TrimSpace(text))
}
