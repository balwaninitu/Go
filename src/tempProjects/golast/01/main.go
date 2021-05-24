package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

var (
	productCodeSlice = []string{}
	mapUsers         = map[string]customer{}
	mapSessions      = map[string]string{}
	//mapShoppingCart  = map[string]linkedList{}
	tpl  *template.Template
	db   *sql.DB
	err  error
	wg   sync.WaitGroup
	lock = make(chan bool, 1)
)

type customer struct {
	Username string
	Password []byte
	First    string
	Last     string
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	var db *sql.DB
	var err error
	var databaseUser string
	databaseUser = os.Getenv("user")
	var databasePW string
	databasePW = os.Getenv("password")
	databaseString := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/my_db", databaseUser, databasePW)
	db, err = sql.Open("mysql", databaseString)
	//db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/my_db")
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Database opened!")
	}
	defer db.Close()
	router := mux.NewRouter()
	router.HandleFunc("/", start)
	router.HandleFunc("/signup", signup)
	http.ListenAndServe(":5221", router)
}

func start(res http.ResponseWriter, req *http.Request) {
	myUser := getUser(res, req)
	tpl.ExecuteTemplate(res, "index.html", myUser)
}

func getUser(res http.ResponseWriter, req *http.Request) customer {
	myCookie, err := req.Cookie("myCookie")
	if err != nil {
		id, _ := uuid.NewV4()
		//set expire paratmeter to prevent session hijacking
		expireTime := time.Now().Add(30 * time.Minute)
		myCookie = &http.Cookie{
			Name:     "myCookie",
			Value:    id.String(),
			Expires:  expireTime,
			HttpOnly: true,
			Path:     "/",
			Domain:   "127.0.0.1",
			Secure:   true,
		}
		http.SetCookie(res, myCookie)
	}
	log.Println("Cookie session has been created/validated")
	var myUser customer
	if username, ok := mapSessions[myCookie.Value]; ok {
		myUser = mapUsers[username]
		log.Println("user details are now stored in myUser")
	}
	return myUser
}

func signup(res http.ResponseWriter, req *http.Request) {
	if alreadyLoggedIn(req) {
		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	var myUser customer
	if req.Method == http.MethodPost {
		username := req.FormValue("username")
		password := req.FormValue("password")
		firstname := req.FormValue("firstname")
		lastname := req.FormValue("lastname")
		errorMessage := map[string]string{}
		if username == "" || password == "" || firstname == "" || lastname == "" {
			errorMessage["input1"] = "Did you miss out entering any of your particulars?"
			tpl.ExecuteTemplate(res, "signup.html", errorMessage)
			delete(errorMessage, "input1")
			return
		}
		results, err := db.Query("SELECT Username, Password, FirstName, LastName from users")
		if err != nil {
			panic(err.Error())
		}
		checkUsers := map[string]customer{}
		var existingUser customer
		for results.Next() {
			err = results.Scan(&existingUser.Username, &existingUser.Password, &existingUser.First, &existingUser.Last)
			if err != nil {
				panic(err.Error())
			}
			checkUsers[existingUser.Username] = existingUser
		}
		for k := range checkUsers {
			if username == k {
				errorMessage["input0"] = "Username has already been taken"
				tpl.ExecuteTemplate(res, "signup.html", errorMessage)
				delete(errorMessage, "input0")
				return
			}
		}
		delete(checkUsers, existingUser.Username)
		reg, err := regexp.Compile("^[a-zA-Z0-9]*$")
		if err != nil {
			log.Fatal(err)
		}
		if reg.MatchString(username) == false {
			log.Println("username contains non-alphanumeric characters")
			errorMessage["input2"] = "username can only have alphanumeric characters"
			tpl.ExecuteTemplate(res, "signup.html", errorMessage)
			delete(errorMessage, "input2")
			return
		}
		if reg.MatchString(password) == false {
			log.Println("password contains non-alphanumeric characters")
			errorMessage["input3"] = "password can only have alphanumeric characters"
			tpl.ExecuteTemplate(res, "signup.html", errorMessage)
			delete(errorMessage, "input3")
			return
		}
		//username must have at least 5 characters
		if len(username) < 6 || len(username) > 12 {
			log.Println("username contains too few or too many characters")
			errorMessage["input4"] = "username must have at least 5 characters and at most 12 characters"
			tpl.ExecuteTemplate(res, "signup.html", errorMessage)
			delete(errorMessage, "input4")
			return
		}
		//password must have at least 8 characters
		if len(password) < 8 || len(password) > 15 {
			log.Println("password contains too few or too many characters")
			errorMessage["input5"] = "password must have at least 8 characters and at most 15 characters"
			tpl.ExecuteTemplate(res, "signup.html", errorMessage)
			delete(errorMessage, "input5")
			return
		}
		//Considering only alphabetical letters as regular expressions for first and last name
		rege, err := regexp.Compile("^[a-zA-Z]*$")
		if err != nil {
			log.Fatal(err)
		}
		if rege.MatchString(firstname) == false {
			log.Println("First name contains non-alphabetical letters")
			errorMessage["input6"] = "first name can only have alphabetical letters"
			tpl.ExecuteTemplate(res, "signup.html", errorMessage)
			delete(errorMessage, "input6")
			return
		}
		if rege.MatchString(lastname) == false {
			log.Println("Last name contains non-alphabetical letters")
			errorMessage["input7"] = "last name can only have alphabetical letters"
			tpl.ExecuteTemplate(res, "signup.html", errorMessage)
			delete(errorMessage, "input7")
			return
		}
		//first name can have at most 20 characters
		if len(firstname) > 20 {
			log.Println("First name has too many letters")
			errorMessage["input8"] = "first name can have at most 20 characters"
			tpl.ExecuteTemplate(res, "signup.html", errorMessage)
			delete(errorMessage, "input8")
			return
		}
		//last name can have at most 20 characters
		if len(lastname) > 20 {
			log.Println("Last name has too many letters")
			errorMessage["input9"] = "last name can have at most 20 characters"
			tpl.ExecuteTemplate(res, "signup.html", errorMessage)
			delete(errorMessage, "input9")
			return
		}
		//defer function in response to potential panic if password is not encrypted
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered")
			}
		}()
		bPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
		if err != nil {
			http.Error(res, "Internal server error", http.StatusInternalServerError)
			log.Panicln("How could this happen")
		}
		myUser = customer{username, bPassword, firstname, lastname}
		mapUsers[username] = myUser
		query := fmt.Sprintf("INSERT INTO users VALUES ('%s', '%s', '%s', '%s')", myUser.Username, myUser.Password, myUser.First, myUser.Last)
		_, err = db.Query(query)
		if err != nil {
			panic(err.Error())
		}
		id, _ := uuid.NewV4()
		expireTime := time.Now().Add(30 * time.Minute)
		myCookie := &http.Cookie{
			Name:     "myCookie",
			Value:    id.String(),
			Expires:  expireTime,
			HttpOnly: true,
			Path:     "/",
			Domain:   "127.0.0.1",
			Secure:   true,
		}
		http.SetCookie(res, myCookie)
		log.Println("New cookie session created")
		mapSessions[myCookie.Value] = username

		http.Redirect(res, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(res, "signup.html", nil)
}

func alreadyLoggedIn(req *http.Request) bool {
	myCookie, err := req.Cookie("myCookie")
	if err != nil {
		return false
	}
	username := mapSessions[myCookie.Value]
	_, ok := mapUsers[username]
	return ok
}
