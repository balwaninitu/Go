package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {

	name := os.Args[1]
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])

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

	nf, err := os.Create("index.html") //os.Create it will create new file
	if err != nil {
		log.Fatal("error creating file", err)
	}
	defer nf.Close()

	io.Copy(nf, strings.NewReader(str)) //nf is destination and tstr is source
	//io.copy help copy the function
	//newReader returns new reading from s

}
