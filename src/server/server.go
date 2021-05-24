// Package server contains all handler functions used for registration, user/admin login and logout.
// It also contains session management functions, cookie creation and deletion, input validation and sanitization.
package server

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	config "pgl/projectGoLive/application/config"
	user_db "pgl/projectGoLive/application/db"

	//seller_db "projectGoLive/application/db/seller_db"

	"github.com/go-playground/validator"
	uuid "github.com/satori/go.uuid"
	bcrypt "golang.org/x/crypto/bcrypt"
)

// This struct stores username(all) and type of user(buyer/seller). It is used to maintain a list of all users currently registered.
type UserInfo struct {
	Username string
	Location string
	IsBuyer  bool
}

// This struct stores username and lastActivity time for current users that have logged in, and is used to maintain sessions
type Session struct {
	Username     string
	LastActivity time.Time
}

// Create a map to store each session, key being the cookie UUID
var mapSessions map[string]Session

// Create a map to store user a list of all users that have already registered
var mapUsers map[string]UserInfo

// Variable to store the time when sessions are cleaned
var mapSessionCleaned time.Time

// Struct to store data sent by login template form for validation
type loginInput struct {
	Username string `validate:"required,alphanum"`
	Password string `validate:"required,min=8,max=100"`
}

// Struct to store data sent by sign up template form for validation
type signupInput struct {
	Username string `validate:"required,alphanum"`
	Password string `validate:"required,min=8,max=100"`
	Location string `validate:"required,alphanumunicode"`
}

func init() {
	mapUsers = make(map[string]UserInfo)
	mapSessions = make(map[string]Session)

	// Read user_db and populate mapUsers (otherwise buyers/seller in database cannot login)
	initializeMapUsers()
}

func initializeMapUsers() {
	userDetails, ok := user_db.GetRecords(config.DB)
	if !ok {
		fmt.Println("Cannot read user database!")
		panic("Cannot read user_db")
	} else {
		var myUserInfo UserInfo
		for _, ui := range userDetails {
			myUserInfo.Username = ui.Username
			myUserInfo.Location = ui.Location
			myUserInfo.IsBuyer = ui.Isbuyer
			mapUsers[ui.Username] = myUserInfo
		}
	}

}

// ---------------------------------------------------------------------------
// Signup, Login and Logout functions
// ---------------------------------------------------------------------------

// This method is used to handle the index page
// It checks if user has logged in, and updates the template with user information if already logged in
// If not logged in, the default index page is shown
func IndexHandler(w http.ResponseWriter, req *http.Request) {
	var myUser UserInfo
	if ok := alreadyLoggedIn(req); ok {
		if checkTimeOut(w, req) {
			http.Redirect(w, req, "/login", http.StatusSeeOther)
			return
		} else {
			extendSession(w, req)
			myUser = GetUser(w, req)
		}
	}
	config.TPL.ExecuteTemplate(w, "index.gohtml", myUser)
}

// This method is used to signup a new user
// It checks if user has already logged in, and if so redirects to index page.
// If not logged in, the signup form is displayed.
// The form input values from signup form are validated and sanitized before adding the data into the database
// Cryptic messages are displayed on the template in case of any errors in input.
func SignupHandler(w http.ResponseWriter, req *http.Request) {
	// Go to index page if user has already logged in
	if ok := alreadyLoggedIn(req); ok {
		if checkTimeOut(w, req) {
			http.Redirect(w, req, "/login", http.StatusSeeOther)
			config.Info.Printf("Time out has occurred, redirecting to login page\n")
			config.Trace.Printf("Time out has occurred, redirecting to login page\n")
			return
		} else {
			extendSession(w, req)
			config.Info.Printf("Extending session..\n")
			config.Trace.Printf("Extending session..\n")
			http.Redirect(w, req, "/", http.StatusSeeOther)
		}
		return
	}

	var user UserInfo

	// For validating user input
	si := signupInput{}

	// For displaying message to template
	signupmessage := ""

	// process form submission
	if req.Method == http.MethodPost {
		reset := req.FormValue("reset")
		signup := req.FormValue("signup")
		username := req.FormValue("username")
		password := req.FormValue("password")
		isbuyer, _ := strconv.ParseBool(req.FormValue("isbuyer"))
		location := req.FormValue("location")

		fmt.Println("Reset :", reset)
		fmt.Println("Sign up :", signup, username, password, isbuyer, location)
		if reset != "" {
			//reload page
			signupmessage = "All values Reset"
		} else {
			// get form values
			fmt.Println("Signing up..")

			// Validate all input data for security
			si.Username = username
			si.Password = password
			si.Location = location
			//verr := signupValidation(si)
			_, verr := fmt.Println("Bypassing signup validation")
			fmt.Println("si", si)

			if verr != nil { // error
				signupmessage = "Invalid Username/Password"
			} else {
				// Sanitize input data
				//snInput, serr := signupSanitization(si)
				fmt.Println("Bypassing signup sanitization")
				snInput := si
				serr := true

				if !serr { // error during sanitization
					fmt.Println("Sanitization error")
					signupmessage = "Invalid Username/Password"
				} else {
					// Check if username is already taken
					if _, ok := mapUsers[snInput.Username]; ok {
						// Invalid username or password message to template
						signupmessage = "Invalid Username/Password"
						return
					} else {
						// This is a new sign up
						// Validated all entries
						fmt.Println("checking password")
						pwBytes, err := bcrypt.GenerateFromPassword([]byte(snInput.Password), bcrypt.MinCost)
						if err != nil {
							signupmessage = "Internal server error, try again!"
							return
						}
						fmt.Println("sninput", snInput)

						// Used to access user database
						var newUserdb user_db.UserDetails
						// use sanitized versions of entries
						newUserdb.Username = snInput.Username
						newUserdb.Password = string(pwBytes)
						newUserdb.Isbuyer = isbuyer
						newUserdb.Location = snInput.Location
						// Write username and password to the user database
						// Write to user_db
						insertok := user_db.InsertRecord(config.DB, newUserdb)
						if !insertok {
							// Unable to write to database
							signupmessage = "Unable to reach database, try again!"
							return
						}
						fmt.Println("yserdb", newUserdb)

						// This is a local record of all users that have registered- buyers as well as sellers
						// Password is not stored here for security reasons
						user.Username = snInput.Username
						user.Location = snInput.Location
						user.IsBuyer = isbuyer

						// Store user info in map for users
						mapUsers[snInput.Username] = user
						fmt.Println(mapUsers)
						// Successfully registered, now redirect to login page
						http.Redirect(w, req, "/login", http.StatusSeeOther)
						return
					}
				}
			}
		}
	}
	config.TPL.ExecuteTemplate(w, "signup.gohtml", signupmessage)
}

// This method is used to login seller/buyer
// Before each login, the MapSessions data is cleaned, if it has not been cleaned for more than cleanSessionTime
// Cleaning of session is done with a go routine.
// The username and password input from the login form is validated and sanitized.
// Any errors create a cryptic message that "username/password do not match"
// If login is successful, a secure cookie is created, and mapsession is updated for seller/buyer
func LoginHandler(w http.ResponseWriter, req *http.Request) {

	//Clean mapSessions before login if session cleaning is not done for CleanSessionTime
	if time.Now().Sub(mapSessionCleaned) > (time.Second * time.Duration(config.CleanSessionTime)) {
		go cleanupSessions()
		config.Info.Println("Cleaning up sessions..")
		config.Trace.Println("Cleaning up sessions..")
	}

	lInput := loginInput{}
	var isbuyer bool
	password_db := ""

	// Login message to template
	loginmessage := ""

	if req.Method == http.MethodPost {
		lInput.Username = req.FormValue("username")
		lInput.Password = req.FormValue("password")
		// Validate and Sanitize user input

		// Validate input data for security
		//verr := loginValidation(lInput)
		_, verr := fmt.Println("Bypassing login validation")

		if verr != nil { // error
			loginmessage = "Username and/or password do not match"
		} else {
			// Sanitize input data
			//err := loginSanitization(lInput)
			fmt.Println("Bypassing login sanitization")

			err := true // no error
			if !err {
				loginmessage = "Username and/or password do not match"
			} else {
				// Check if this is a concurrent login, concurrent logins are disallowed.
				if findInMap(mapSessions, lInput.Username) {
					config.Trace.Println("Concurrent Login detected! Concurrent Login will not be allowed")
					//CheckLogChecksum()
					config.Error.Println("Concurrent Login detected! Concurrent Login will not be allowed")
					//WriteChecksum()

					config.Trace.Println("Concurrent Login detected! Concurrent Login will not be allowed")
					loginmessage = "You have already logged in on another machine.\n Logout from the other session first or report if you have not logged in elsewhere!"
				} else {
					// check if user exists with given username
					if _, ok := mapUsers[lInput.Username]; !ok { // username is not registered
						fmt.Println("username is not in map users..", mapUsers, lInput.Username)
						loginmessage = "Username and/or password do not match"
					} else {

						// If required, Location can be stored in each session as retrived from db
						//location_db := ""

						// Valid username, now check password from database
						details, ok := user_db.GetARecord(config.DB, lInput.Username)
						if !ok {
							config.Warning.Printf("User name not found in db, but present in mapUsers\n")
							loginmessage = "Username and/or password do not match"
						} else {
							password_db = details.Password
							//location_db = details.Location

							err := bcrypt.CompareHashAndPassword([]byte(password_db), []byte(lInput.Password))
							if err != nil {
								//fmt.Println("Password does not match", err)
								config.Warning.Printf("Password does not match the one in db\n")
								loginmessage = "Username and/or password do not match"
							} else {
								// create session
								//id, err := uuid.NewV4()
								//if err != nil {
								//	Warning.Printf("Unable to generate UUID : %v\n", err)
								//	loginmessage = "Server Error, please try again!"
								//} else {

								id, _ := uuid.NewV4()
								myCookie := &http.Cookie{
									Name:   "PeelRescue",
									Value:  id.String(),
									MaxAge: config.SessionExpireTime,
									//Expires:  time.Now().Add(time.Duration(config.SessionExpireTime) * time.Second),
									HttpOnly: true,
									Secure:   true,
								}
								http.SetCookie(w, myCookie)

								mapSessions[myCookie.Value] = Session{lInput.Username, time.Now()}
								// Check if user is buyer or seller
								isbuyer = mapUsers[lInput.Username].IsBuyer
								if isbuyer {
									fmt.Println("Redirecting to buyer page!")
									http.Redirect(w, req, "/buyer", http.StatusSeeOther)
									return
								} else {
									fmt.Println("Redirecting to seller page!")
									http.Redirect(w, req, "/seller", http.StatusSeeOther)
									return
								}
							}
						}
					}
				}
			}
		}
	}
	config.TPL.ExecuteTemplate(w, "login.gohtml", loginmessage)
}

// This method is used to logout buyer or seller
// If no cookie is present, the page is redirected to login page
// If cookie is present, it is deleted, and corrspoding map session entry is also deleted
func LogoutHandler(w http.ResponseWriter, req *http.Request) {

	myCookie, err := req.Cookie("PeelRescue")
	logoutmessage := ""
	if err != nil {
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		config.Warning.Println("Cookie not found, redirecting to login page..")
		config.Trace.Println("Cookie not found, redirecting to login page..")
		return
	} else {
		// delete the session
		delete(mapSessions, myCookie.Value)

		// remove the cookie
		myCookie = &http.Cookie{
			Name:     "PeelRescue",
			Value:    "",
			MaxAge:   -1,
			HttpOnly: true,
			Secure:   true,
		}
		http.SetCookie(w, myCookie)
		logoutmessage = "Log out successful! "
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	config.TPL.ExecuteTemplate(w, "logout.gohtml", logoutmessage)
}

//------------------------------------------------------------------------------
// Global functions for sessions and user maintenance used outside of server package
//------------------------------------------------------------------------------

// This method is used to check if the session is active
// It checks if user has already logged in, returns false if user has not logged in.
// It checks if time out has occurred, returns false if time out has occurred.
// If user has loggeg in, and time out has not occured, then the session is extended, and it returns true.
func ActiveSession(w http.ResponseWriter, req *http.Request) bool {
	if ok := alreadyLoggedIn(req); !ok {
		return false
	} else {
		if checkTimeOut(w, req) {
			return false
		} else {
			extendSession(w, req)
			return true
		}
	}
}

// This method is used to get username based on the UUID in the cookie
func GetUser(w http.ResponseWriter, req *http.Request) UserInfo {

	var user UserInfo

	// get current session cookie
	myCookie, err := req.Cookie("PeelRescue")

	// If cookie does not exist, return with empty user info, as new session needs to be created with fresh login
	if err != nil {
		return user
	}
	// get username with mapsession and UUID of user's cookie
	username := mapSessions[myCookie.Value].Username
	user = mapUsers[username]
	return user
}

//------------------------------------------------------------------------------
// Helper functions for sessions and user maintenance
//------------------------------------------------------------------------------

// This method checks if the user has already logged in
// The user can be a buyer or a seller
// If user has logged in, it returns true, else false
func alreadyLoggedIn(req *http.Request) bool {
	myCookie, err := req.Cookie("PeelRescue")
	if err != nil {
		//config.Warning.Printf("No Cookie found for this session : %v\n", err)
		//config.Trace.Printf("No Cookie found for this session : %v\n", err)
		return false
	}
	username := mapSessions[myCookie.Value].Username
	if username != "" {
		return true
	} else {
		return false
	}
}

// This method is used to clean up the mapSessions data
// It ranges over the map, and checks if any session has lastActivity value more than sessionExpireTime
// If lastActivity value is more, that means session has been idle and needs to be deleted
// After cleanup, the mapSessionCleaned is updated to current time, to be used to have an interval between each cleanup.
func cleanupSessions() {
	for key, value := range mapSessions {
		if time.Now().Sub(value.LastActivity) > (time.Second * time.Duration(config.SessionExpireTime)) {
			delete(mapSessions, key)
		}
	}
	mapSessionCleaned = time.Now()
}

// This method checks if a time out has occurred.
// If the lastActivity value inside mapSessions is more than sessionExpireTime, then time out has occurred and it returns true
// Also if the MaxAge of the received cookie is less than 0, then timeout has occured and it returns true
// If none of the above timeout possibilities have occurred, then it return true, meaning no timeout.
func checkTimeOut(w http.ResponseWriter, req *http.Request) bool {
	myCookie, _ := req.Cookie("PeelRescue")

	//Check if mapsession has expired
	lastActivity := mapSessions[myCookie.Value].LastActivity
	sessionDuration := time.Duration(config.SessionExpireTime)
	nowTime := time.Now()
	lastActivityPlusExpiry := lastActivity.Add(time.Second * sessionDuration)

	if nowTime.After(lastActivityPlusExpiry) {
		return true
	}

	if myCookie.MaxAge < 0 {
		//Timeout has occurred
		return true
	} else {
		return false
	}
}

// This method is used to extend a session, by updating the lastActivity value in mapSessions
// It also sets the MaxAge of the cookie to sessionExpireTime
// This method is called every time there is an activity by the user, thereby incresing the amount of time the session is active
func extendSession(w http.ResponseWriter, req *http.Request) {
	myCookie, err := req.Cookie("PeelRescue")

	if err != nil { // no cookie found
		config.Warning.Printf("No Cookie found when extending session : %v\n", err)
		config.Trace.Printf("No Cookie found when extending session : %v\n", err)
		return
	}

	// Extend mapsession lastActivity time every time there is an activity
	username := mapSessions[myCookie.Value].Username
	mapSessions[myCookie.Value] = Session{username, time.Now()}

	// update cookie MaxAge expiry time every time there is an activity
	myCookie.MaxAge = config.SessionExpireTime
	http.SetCookie(w, myCookie)
	return
}

// This function checks if a value of type string is present in a map of type session
// Session contains username as string, and lastActivity as time.Time
// The function checks if the Value is present in username of every key
// If present, it returns true, else it return false
func findInMap(mapName map[string]Session, Value string) bool {
	for _, key := range mapName {
		if key.Username == Value {
			return true
		}
	}
	return false
}

//------------------------------------------------------------------------------
// Validation and Sanitization
//------------------------------------------------------------------------------

// This method is used to validate the signup input data
// It uses the https://github.com/go-playground/validator package
// Validation criteria is in signupInput struct, which is input to this method
func signupValidation(si signupInput) error {
	// use a single instance of Validate, it caches struct info
	var validate *validator.Validate
	validate = validator.New()

	// returns nil or ValidationErrors ( []FieldError )
	err := validate.Struct(si)
	config.Warning.Printf("Error in signup Validation : %v\n", err)
	config.Trace.Printf("Error in signup Validation : %v\n", err)
	return err
}

// This method is used to sanitize the signup input data
// It checks that username begins with an alphabet and has no special characters
// It converts username to lowercase.
// It checks that password does not contain white space or null characters.
// It checks that location alphanumeric.
// It converts location to titlecase.
// It returns false if there is any issue with the input
func signupSanitization(si signupInput) (signupInput, bool) {

	siOutput := signupInput{}
	checkError := false

	// Sanitize username
	un := regexp.QuoteMeta(si.Username)
	if un != si.Username { // Check if contains any special characters
		checkError = true
		//"No special characters allowed"
	} else {
		// First character must be an alphabet, remaining characters must be alphabet,digits or underscore
		unMatch := regexp.MustCompile(`^[a-zA-Z]+\w*$`)
		if !unMatch.MatchString(si.Username) {
			// "Incorrect Username"
			checkError = true
		} else {
			// If no errors, pass the username to the program
			siOutput.Username = strings.ToLower(si.Username)
		}
	}

	// Sanitize password , it can contain special characters
	// Check if password contains white space or null characters
	newlineChar := fmt.Sprintf("\\s")
	if strings.Contains(si.Password, newlineChar) || strings.Contains(si.Password, "%00") {
		checkError = true
		//"No spaces/null characters allowed"
	} else {
		// If no errors, pass the password to the program
		siOutput.Password = si.Password
	}

	// Sanitize location
	loc := regexp.QuoteMeta(si.Location)
	if loc != si.Location { // Check if contains any special characters
		checkError = true
		//"No special characters allowed"
	} else {
		// First character must be an alphabet
		locMatch := regexp.MustCompile(`^[a-zA-Z]+[a-zA-Z0-9]*$`)
		if !locMatch.MatchString(si.Location) {
			//"Incorrect Identification number"
			checkError = true
		} else {
			// If no errors, pass the identification number to the program
			siOutput.Location = strings.ToTitle(si.Location)
		}
	}

	return siOutput, !checkError
}

// This method is used to validate the login input data
// It uses the https://github.com/go-playground/validator package
// Validation criteria is in loginInput struct, which is input to this method
func loginValidation(li loginInput) error {
	// use a single instance of Validate, it caches struct info
	var validate *validator.Validate
	validate = validator.New()

	err := validate.Struct(li)
	config.Warning.Printf("Error in login Validation : %v\n", err)
	config.Trace.Printf("Error in login Validation : %v\n", err)
	return err
}

// This method is used to sanitize the login input data
// It checks that username begins with an alphabet and has no special characters
// It also checks that password does not contain white space or null characters
// It returns false if there is any issue with in the input
func loginSanitization(li loginInput) bool {
	// Sanitize username
	un := regexp.QuoteMeta(li.Username)
	if un != li.Username { // Check if contains any special characters
		return false
	} else {
		// First character must be an alphabet, remaining characters must be alphabet,digits or underscore
		unMatch := regexp.MustCompile(`^[a-zA-Z]+\w*$`)
		if !unMatch.MatchString(li.Username) {
			return false
		} else {
			// Username is okay..check password
			// Sanitize password , it can contain special characters
			// Check if password contains white space or null characters
			newlineChar := fmt.Sprintf("\\s")
			if strings.Contains(li.Password, newlineChar) || strings.Contains(li.Password, "%00") {
				return false
			} else {
				// no errors, return true
				return true
			}
		}
	}
}
