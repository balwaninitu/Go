package main

import (
	"fmt"
	"text/template"
)

func main() {

	t, err := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
	//err = t.ExecuteTemplate(out, "T", "<script>alert('you have been pwned')</script>")
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println(t)

}
