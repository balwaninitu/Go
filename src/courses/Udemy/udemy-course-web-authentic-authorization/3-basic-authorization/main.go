package main

import (
	"encoding/base64"
	"fmt"
)

/*type curl -u user:pass -v google.com(-v is verbose asking to see my request)
do this on gitbash which will provide you authorization with word Basic and user:pass
put together in base64. base 64 is nice way to put any binary data in form */
func main() {
	fmt.Println(base64.StdEncoding.EncodeToString([]byte("user:pass")))

	//with above you will get same key(dXNlcjpwYXNz) on your terminal
	/*user can be name or email put colon(:) and then password then encode in base64
	then put basic that is your header authorization ask same with every request & thats basic authentication
	base64 is easily reversible so can not do with unsecure connection always do with https
	dont use basic authentication for endpoint best use is to with login screen and from there can use other routes
	*/

}
