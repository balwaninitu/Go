package main

// import (
// 	"fmt"
// 	"net/http"
// 	"time"
// )

// func a(res http.ResponseWriter, req *http.Request) {
// 	fmt.Println("message out")
// }

// func main() {
// 	s := &http.Server{
// 		Addr:        ":8080",
// 		ReadTimeout: 10 * time.Second,
// 	}

// 	s.Handler()
// 	err := s.ListenAndServe()
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }