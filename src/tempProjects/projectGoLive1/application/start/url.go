package start

import (
	server "projectGoLive/application/server"
)

func mapUrls() {

	router.HandleFunc("/", server.Index)
	router.HandleFunc("/signup", server.Signup)
	router.HandleFunc("/login", server.Login)
	router.HandleFunc("/logout", server.Logout)

	// to be updated for handling JSON
}
