package godown

import (
	"io/ioutil"
	"testing"
)

func TestReadme(t *testing.T) {

	const simpleExample = `
package main

import . "github.com/skaggmannen/go-down"

func main() {
	doc := Doc(
		H1("Top heading"),
		P("This is earth speaking."),
		Ul(
			Li("An item"),
			Li(Strike("Another item")),
		),
		Quote("Build it and they will come."),
		P("How about a link to the ", Ref("Top heading"), "?"),
		P("Or maybe we want the link text to be ", Link("different", "#top-heading"), "."),
	)
	fmt.Println(doc)
}
`
	readme := Doc(
		H1("go-down"),
		P(`
			This is a simple project to generate Markdown using Go code. It's even been used 
			to generate this readme file!
		`),
		H2("Example"),
		Code("go", simpleExample),
	)
	_ = ioutil.WriteFile("README.md", []byte(readme), 0644)
}
