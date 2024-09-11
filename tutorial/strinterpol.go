package main

import (
	"fmt"
	"text/template"
)

func main() {
	template.New("my-templ").Parse("{{.}}")
	x := `
	{{lol}}
	`
	fmt.Print(x)
}
