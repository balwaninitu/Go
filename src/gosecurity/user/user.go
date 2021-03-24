package user

import (
	"gosecurity/config"
	"gosecurity/logger"
	"net/http"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	Username string
	Password []byte
}

var mapUsers = map[string]user{}
var mapSessions = map[string]string{}

func init() {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.MinCost)
	mapUsers["admin"] = user{"admin", hashedPassword}
}

/*Index function will execute parsed html template
to route towards the sign up/login page.*/
func Index(w http.ResponseWriter, r *http.Request) {
	myUser := GetUser(w, r)
	config.TPL.ExecuteTemplate(w, "index.gohtml", myUser)
}

//Signup func will let new admin to sign up in the system.
//In case of any error in input it wil display error message.
//Func will map the existence of username and display message accordingly.
//if signup successfully it will set cookie and generate UUID and hashed the password.
func Signup(w http.ResponseWriter, r *http.Request) {
	if AlreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		logger.WarningLog.Println("Unsuccessful signup attempt!")
		return
	}
	var myUser user
	// process form submission
	if r.Method == http.MethodPost {
		// get form values
		username := r.FormValue("username")
		password := r.FormValue("password")

		l := len(r.FormValue("username"))
		if l < 5 || l > 10 {
			http.Error(w, http.StatusText(406)+"- Username must be between 5 and 10 characters long", http.StatusNotAcceptable)
			return
		}
		l = len(r.FormValue("password"))
		if l < 5 || l > 10 {
			http.Error(w, http.StatusText(406)+"- Password must be between 5 and 10 characters long", http.StatusNotAcceptable)
			return
		}
		if username != "" {
			// check if username exist/ taken
			if _, ok := mapUsers[username]; ok {
				http.Error(w, "Username already taken", http.StatusForbidden)
				return
			}
			// create session
			id, _ := uuid.NewV4()
			myCookie := &http.Cookie{
				Name:  "myCookie",
				Value: id.String(),
			}
			http.SetCookie(w, myCookie)
			mapSessions[myCookie.Value] = username

			bPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
			if err != nil {
				http.Error(w, "Internal server error", http.StatusInternalServerError)
				return
			}

			myUser = user{username, bPassword}
			mapUsers[username] = myUser
		}
		// redirect to main index
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return

	}
	config.TPL.ExecuteTemplate(w, "signup.gohtml", myUser)
	logger.TraceLog.Println("Signup successfully!")
}

/*when user login with credential, func first check its existence by matching
and comparing hashed password which get generated during signup */
func Login(w http.ResponseWriter, r *http.Request) {
	if AlreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		logger.WarningLog.Println("Unsuccessful login attempt!")
		return
	}

	// process form submission
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")
		// check if user exist with username
		myUser, ok := mapUsers[username]
		if !ok {
			http.Error(w, "Username and/or password do not match", http.StatusUnauthorized)
			return
		}
		// Matching of password entered
		err := bcrypt.CompareHashAndPassword(myUser.Password, []byte(password))
		if err != nil {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}
		// create session
		id, _ := uuid.NewV4()
		myCookie := &http.Cookie{
			Name:  "myCookie",
			Value: id.String(),
		}
		http.SetCookie(w, myCookie)
		mapSessions[myCookie.Value] = username
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	config.TPL.ExecuteTemplate(w, "login.gohtml", nil)
	logger.TraceLog.Println("Login successfully!")
}

//Logout func when invoked, it will remove the cookie from system and delete the session.
//It will delete cookie by setting MaxAge of cookie to less than zero.
func Logout(w http.ResponseWriter, r *http.Request) {
	if !AlreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	myCookie, _ := r.Cookie("myCookie")
	// delete the session
	delete(mapSessions, myCookie.Value)
	// remove the cookie
	myCookie = &http.Cookie{
		Name:   "myCookie",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, myCookie)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

/*GetUser func will use set cookies func set-Cookie header to the
provided ResponseWriter's headers. The provided cookie must have
a valid Name. Invalid cookies may be silently dropped.*/
func GetUser(w http.ResponseWriter, r *http.Request) user {
	// get current session cookie
	myCookie, err := r.Cookie("myCookie")
	if err != nil {
		id, _ := uuid.NewV4()
		myCookie = &http.Cookie{
			Name:  "myCookie",
			Value: id.String(),
		}

	}
	http.SetCookie(w, myCookie)

	// map user's cookie if exist already
	var myUser user
	if username, ok := mapSessions[myCookie.Value]; ok {
		myUser = mapUsers[username]
	}

	return myUser
}

/*AlreadyLoggedIn func get invoked in login and signup func
to map existence of username*/
func AlreadyLoggedIn(r *http.Request) bool {
	myCookie, err := r.Cookie("myCookie")
	if err != nil {
		return false
	}
	username := mapSessions[myCookie.Value]
	_, ok := mapUsers[username]
	return ok
}
