package godown

import (
	"bufio"
	"strings"
)

func Indent(text string, prefix string) string {
	var result []string
	sc := bufio.NewScanner(strings.NewReader(text))
	for sc.Scan() {
		result = append(result, sc.Text())
	}
	return strings.Join(result, "\n"+prefix)
}

func ReIndent(val string, prefix string) string {
	var lines []string
	sc := bufio.NewScanner(strings.NewReader(val))
	for sc.Scan() {
		lines = append(lines, strings.TrimSpace(sc.Text()))
	}
	return strings.Join(lines, "\n"+prefix)
}

func fixIndent(val string, prefix string) string {
	var lines []string
	sc := bufio.NewScanner(strings.NewReader(val))
	for sc.Scan() {
		lines = append(lines, strings.TrimSpace(sc.Text()))
	}
	return strings.Join(lines, "\n"+prefix)
}
