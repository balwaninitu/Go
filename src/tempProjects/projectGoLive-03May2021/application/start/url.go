package start

import (
	"net/http"
	"projectGoLive/application/seller"
	"projectGoLive/application/server"
)

func mapUrls() {
	//router.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("/assets/"))))
	// Choose the folder to serve
	resourceDir := "/assets/"
	router.PathPrefix(resourceDir).
		Handler(http.StripPrefix(resourceDir, http.FileServer(http.Dir("."+resourceDir))))

	//staticDir := "/assets/"
	//http.Handle(staticDir, http.StripPrefix(staticDir, http.FileServer(http.Dir("."+staticDir))))

	router.HandleFunc("/", server.IndexHandler)
	router.HandleFunc("/signup", server.SignupHandler)
	router.HandleFunc("/login", server.LoginHandler)
	router.HandleFunc("/logout", server.LogoutHandler)
	//router.HandleFunc("/buyer", buyer.BuyerHandler)
	router.HandleFunc("/seller", seller.SellerHandler)

	// to be updated for handling JSON
}
