package main

import (
	"fmt"
	"html/template"
	"os"
)

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}
	user := struct {
		Name string
	}{
		Name: "Basky come on man",
	}

	fmt.Println("template exp")
	t.Execute(os.Stdout, user)
}
