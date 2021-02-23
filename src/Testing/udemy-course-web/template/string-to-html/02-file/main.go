package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {

	name := "Nitu B"
	//sprint convert html to string
	str := fmt.Sprint(`
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Hello World!</title>
	</head>
	<body>
	<h1>` + name + `</h1>
	</body>
	</html>
	`)

	fmt.Println(str)

	nf, err := os.Create("index.html")
	if err != nil {
		log.Fatal("error creating file", err)
	}
	defer nf.Close()

	io.Copy(nf, strings.NewReader(str))

}
