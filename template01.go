package main

import (
	"os"
	"text/template"
)

// create a struct that matches the "blanks" in our template string
type Todo struct {
	Name string
	Description string
}

func main() {
	td := Todo{"Test templates", "Let's test a template to see that magic."}

	// put the pieces together
    	t, err := template.New("todos").Parse("You have a task named \"{{ .Name}}\" with description: \"{{ .Description}}\"")
	

    err = t.Execute(os.Stdout, td)
    // error checking
    if err != nil {
	    panic(err)
    }

        // it is possible to reuse the same template, without
    // needing to create or parse it again by providing the struct you want
    // to use to the "Execute" function again
    rzfNew := Todo{"Make breakfast", "Grind and brew coffee, and find a croissant"}

    err = t.Execute(os.Stdout, rzfNew) // resuse our same template
    if err != nil {
	    panic(err)
    }
}
