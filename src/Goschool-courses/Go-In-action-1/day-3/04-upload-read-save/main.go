package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/", form)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func form(w http.ResponseWriter, req *http.Request) {

	var stringToPrint string

	if req.Method == http.MethodPost {
		file, fileInfo, err := req.FormFile("filename")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer file.Close()
		fmt.Println("\nFile:", file, "\nFileHeaderProperties", fileInfo, "\nerr:", err)
		readData, err := ioutil.ReadAll(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		stringToPrint = string(readData)

		destFile, err := os.Create(fileInfo.Filename)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer destFile.Close()
		_, err = destFile.Write(readData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "text/html;charset=utf-8")
	io.WriteString(w, `
	<formmethod="POST"enctype="multipart/form-data">
	<inputtype="file"name="filename">
	<inputtype="submit"></form><br>`+stringToPrint)

}
