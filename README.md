# go-down

The is a simple package to generate Markdown using Go code.

## Example

```go
package main

import . "github.com/skaggmannen/go-down"

func main() {
  doc := Doc(
    H1("Hello world!"),
    P("This is earth speaking."),
    Ul(
      Li("An item"),
      Li(Em("Another item"),
    ),
  )
  fmt.Println(doc)
}
```
