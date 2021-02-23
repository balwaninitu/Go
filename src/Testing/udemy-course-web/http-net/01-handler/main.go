package main

import (
	"fmt"
	"net/http"
)

type riddham string

var ri riddham

//implicitly riddham is handler because the methode is taking ServeHttp, responsewriter and request which is the functionality of handler
func (ri riddham) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "I am an intelligent girl")
	fmt.Fprintln(w, "I like to play Piano and dance")
}

func main() {
	//var rid riddham
	//listenand serve takes port and handler as input parameter so it is taking ri(riddham)
	http.ListenAndServe(":8080", ri)

}

//data send to body the method should be  post
//data send to url method should be GET
