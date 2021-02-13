package godown

import (
	"io/ioutil"
	"testing"
)

func TestReadme(t *testing.T) {
	simpleExample, _ := ioutil.ReadFile("readme_test.go")

	readme := Doc(
		H1("go-down"),
		P(`
			This is a simple project to generate Markdown using Go code. It's even been used 
			to generate this readme file!
		`),
		H2("Example"),
		Code("go", string(simpleExample)),
	)
	_ = ioutil.WriteFile("README.md", []byte(readme), 0644)
}
