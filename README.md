# go-down

This is a simple project to generate Markdown using Go code. It's even been used
to generate this readme file!

## Example

```go
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
```