package main

import (
	"encoding/json"
	"os"
)

type Server struct {
	// ID will not be outputed.
	ID int `json:"-"`

	// ServerName2 will be converted to JSON type.
	ServerName  string `json:"serverName"`
	ServerName2 string `json:"serverName2,string"`

	// If ServerIP is empty, it will not be outputted.
	ServerIP string `json:"serverIP,omitempty"`
}

func main() {

	s := Server{
		ID:          3,
		ServerName:  `Go "1.0" `,
		ServerName2: `Go "1.0" `,
		ServerIP:    ``,
	}
	b, _ := json.Marshal(s)
	os.Stdout.Write(b)
}
