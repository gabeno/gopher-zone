// builtin support for creating dynamic content or showing customized
// output to user. sibling package for html has extra security features
// and is used for generating html

package main

import (
	"os"
	"text/template"
)

func main() {
	// create a new template and parse its body from a string
	// templates comprise of static text and actions enclosed in `{{}}` and used to
	// dynamically insert content
	t1 := template.New("t1")
	t1, err := t1.Parse("Value is {{.}}\n")
	if err != nil {
		panic(err)
	}

	// alternatively use `template.Must` to panic in case `Parse` returns an error
	// used for templates initialized in global scope
	t1 = template.Must(t1.Parse("Value is {{.}}\n"))

	// execute the template
	t1.Execute(os.Stdout, "5")
	t1.Execute(os.Stdout, "some text")
	t1.Execute(os.Stdout, []string{
		"Go",
		"Rust",
		"Python",
		"Typescript",
	})

	// helper function
	Create := func(name, str string) *template.Template {
		return template.Must(template.New(name).Parse(str))
	}

	// if the data is a struct, we can use `{{.FiedName}}` action to access its fields
	// the fields should be exported to be accessible when the template is executing
	// same applies to maps
	t2 := Create("t2", "Name {{.Name}}\n")

	t2.Execute(os.Stdout, struct{ Name string }{"Jane Doe"})
	t2.Execute(os.Stdout, struct{ Name string }{Name: "King Kaka"})
	t2.Execute(os.Stdout, map[string]string{"Name": "Bill Kates"})

	// if/else provides conditional execution for templates
	// `-` trim whitespace
	// a value is considered false if its the default value of a type
	t3 := Create("t3", "{{if . -}} yes {{else -}} no {{end}}\n")

	t3.Execute(os.Stdout, "not empty")
	t3.Execute(os.Stdout, "")

	// range blocks let us loop through slices, arrays, maps or channels
	// inside the range block `{{.}}` is set to the current item of the iteration
	t4 := Create("t4", "Range: {{range .}}{{.}} {{end}}\n")

	t4.Execute(os.Stdout, []string{
		"Go",
		"Rust",
		"Python",
		"Typescript",
	})
}
