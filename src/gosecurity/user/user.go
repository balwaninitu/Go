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
	First    string
	Last     string
}

var mapUsers = map[string]user{}
var mapSessions = map[string]string{}

func init() {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.MinCost)
	mapUsers["admin"] = user{"admin", hashedPassword, "admin", "admin"}
}

func Index(res http.ResponseWriter, req *http.Request) {
	myUser := GetUser(res, req)
	config.TPL.ExecuteTemplate(res, "index.gohtml", myUser)
}

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
		firstname := r.FormValue("firstname")
		lastname := r.FormValue("lastname")
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

			myUser = user{username, bPassword, firstname, lastname}
			mapUsers[username] = myUser
		}
		// redirect to main index
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return

	}
	config.TPL.ExecuteTemplate(w, "signup.gohtml", myUser)
	logger.TraceLog.Println("Signup successfully!")
}

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

	// if the user exists already, get user
	var myUser user
	if username, ok := mapSessions[myCookie.Value]; ok {
		myUser = mapUsers[username]
	}

	return myUser
}

func AlreadyLoggedIn(r *http.Request) bool {
	myCookie, err := r.Cookie("myCookie")
	if err != nil {
		return false
	}
	username := mapSessions[myCookie.Value]
	_, ok := mapUsers[username]
	return ok
}

func Restricted(w http.ResponseWriter, r *http.Request) {
	myUser := GetUser(w, r)
	if !AlreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	config.TPL.ExecuteTemplate(w, "restricted.gohtml", myUser)
}
