package godown

import (
	"io/ioutil"
	"testing"
)

func TestDoc(t *testing.T) {
	val := Doc(
		H1("Hello world!"),
		Toc(),
		Intro(),
		Problem(),
		CodeExample(),
	)
	_ = ioutil.WriteFile("./testdata/doc.md", []byte(val), 0644)
}

const example = `
package main

import fmt

func main() {
	fmt.Println("Hello world!")
}
`

func Toc() string {
	return Doc(
		Ul(
			Li(Ref("Introduction")),
			Li(Ref("Problem description")),
		),
	)
}

func CodeExample() string {
	return Doc(
		Code("go", example),
	)
}

func Problem() string {
	return Doc(
		H2("Problem description"),
		P(
			`I have some very important messages for you today. First of all
			 I would like to renounce all these rumours about me being a lizard; it _not_ true!`,
		),
		H3("The blame game"),
		P(
			`The people responsible for spreading these lies are the following:`,
		),
		Ol(
			Li("The Pope."),
			Li("The Queen of England."),
			Li(
				P(`
					My neighbor Keith, who also never returned that VHS of Bloodsports with 
					Jean Claude van Damme that he borrowed back in August of '97. It was very 
					precious to me, Keith, and I'd like it back post haste.
				`),
				P("If you're somehow unable to fulfill my request, consider these options:"),
				Ul(
					Li("Eat shit."),
					Li("Go to hell."),
					Li("Give me back my Bloodsports."),
				),
			),
			Li("My cat."),
			Li(
				P("This stuff:"),
				CodeExample(),
			),
			Li("See ", Link("Introduction", "#introduction"), " for more info."),
		),
		P("That's all folks, have a good night."),
	)
}

func Intro() string {
	return Doc(
		H2("Introduction"),
		P("Please allow me to introduce myself; I'm a man of wealth and taste."),
	)
}
