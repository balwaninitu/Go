package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB
var tpl *template.Template
var err error
var mapUsers = map[string]UserDetails{}
var mapSessions = map[string]string{}

func init() {
	var err error
	db, err = sql.Open("postgres", "postgres://postgres:password@localhost/dentalclinic1?sslmode=disable")
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("You are connected to your database.")
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))

}

//UserDetails can be exported
type UserDetails struct {
	username string
	Password []byte
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/signup", signup)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, r *http.Request) {
	user := getUserDetails(w, r)
	tpl.ExecuteTemplate(w, "index.gohtml", user)

}

func restricted(w http.ResponseWriter, r *http.Request) {
	user := getUserDetails(w, r)
	if !alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "restricted.gohtml", user)
}

func alreadyLoggedIn(r *http.Request) bool {
	myCookie, err := r.Cookie("myCookie")
	if err != nil {
		return false
	}
	username := mapSessions[myCookie.Value]
	_, ok := mapUsers[username]
	return ok
}

func signup(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	var user UserDetails
	// process form submission
	if r.Method == http.MethodPost {
		// get form values
		username := r.FormValue("username")
		password := r.FormValue("password")

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

			user = UserDetails{username, bPassword}
			mapUsers[username] = user
		}
		// redirect to main index
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return

	}
	tpl.ExecuteTemplate(w, "signup.gohtml", user)
}

func login(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// process form submission
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")
		// check if user exist with username
		user, ok := mapUsers[username]
		if !ok {
			http.Error(w, "username and/or password do not match", http.StatusUnauthorized)
			return
		}
		// Matching of password entered
		err := bcrypt.CompareHashAndPassword(user.Password, []byte(password))
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

	tpl.ExecuteTemplate(w, "login.gohtml", nil)
}

func logout(w http.ResponseWriter, r *http.Request) {
	if !alreadyLoggedIn(r) {
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

func getUserDetails(w http.ResponseWriter, r *http.Request) UserDetails {
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
	var user UserDetails
	if username, ok := mapSessions[myCookie.Value]; ok {
		user = mapUsers[username]
	}

	return user
}
