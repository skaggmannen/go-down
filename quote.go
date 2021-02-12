package godown

import (
	"bufio"
	"strings"
)

func Quote(items ...string) string {
	val := strings.TrimSpace(strings.Join(items, ""))
	var result []string
	sc := bufio.NewScanner(strings.NewReader(val))
	for sc.Scan() {
		result = append(result, "> "+strings.TrimSpace(sc.Text()))
	}
	return strings.Join(result, "\n") + "\n"
}
