package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {

	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {

		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
		}

		go handle(conn)

	}

}

func handle(conn net.Conn) {
	defer conn.Close()
	request(conn)
}

func request(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	i := 0
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)

		if i == 0 {
			mux(conn, ln)
		}
		if ln == "" {
			break
		}
		i++
	}

}

func mux(conn net.Conn, ln string) {

	m := strings.Fields(ln)[0]
	u := strings.Fields(ln)[1]

	fmt.Println("**Method**", m)
	fmt.Println("**URL**", u)

	if m == "GET" && u == "/" {
		riddham(conn)
	}

	if m == "GET" && u == "/age" {
		age(conn)
	}

	if m == "GET" && u == "/weight" {
		weight(conn)
	}

	if m == "GET" && u == "/height" {
		height(conn)
	}

	if m == "GET" && u == "/hobbies" {
		hobbies(conn)
	}

}

func riddham(conn net.Conn) {

	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body>
	<strong>RIDDHAM</strong><br>
	<a href="/">riddham</a><br>
	<a href="/age">age</a><br>
	<a href="/weight">weight</a><br>
	<a href="/height">height</a><br>
	<a href="/hobbies">hobbies</a><br>
	</body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)

}

func age(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body>
	<strong>AGE</strong><br>
	<a href="/">riddham</a><br>
	<a href="/age">age</a><br>
	<a href="/weight">weight</a><br>
	<a href="/height">height</a><br>
	<a href="/hobbies">hobbies</a><br>
	</body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)

}

func weight(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body>
	<strong>WEIGHT</strong><br>
	<a href="/">riddham</a><br>
	<a href="/age">age</a><br>
	<a href="/weight">weight</a><br>
	<a href="/height">height</a><br>
	<a href="/hobbies">hobbies</a><br>
	</body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)

}

func height(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body>
	<strong>HEIGHT</strong><br>
	<a href="/">riddham</a><br>
	<a href="/age">age</a><br>
	<a href="/weight">weight</a><br>
	<a href="/height">height</a><br>
	<a href="/hobbies">hobbies</a><br>
	</body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)

}

func hobbies(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body>
	<strong>HOBBIES</strong><br>
	<a href="/">riddham</a><br>
	<a href="/age">age</a><br>
	<a href="/weight">weight</a><br>
	<a href="/height">height</a><br>
	<a href="/hobbies">hobbies</a><br>
	</body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)

}
