/*Package server contains all handlers use for sign up,
seller/buyer login and logout.It also contains functions file for
input validation and sanitization.*/
package server

import (
	"net/http"
	"strconv"
	"time"

	config "projectGoLive/application/config"
	user_db "projectGoLive/application/user_db"

	uuid "github.com/satori/go.uuid"
	bcrypt "golang.org/x/crypto/bcrypt"
)

/*This method is used to handle the index page
It checks if user has logged in, and updates the template with user information if already logged in
If not logged in, the default index page is shown*/
func IndexHandler(w http.ResponseWriter, req *http.Request) {
	myUser := GetUser(w, req)
	if ok := alreadyLoggedIn(req); ok {
		if myUser.Username == "admin" {
			//redirecting to admin page
			http.Redirect(w, req, "/admin", http.StatusSeeOther)
			return
		} else if myUser.IsBuyer {
			//redirecting to buyer page
			http.Redirect(w, req, "/buyer", http.StatusSeeOther)
			return
		} else {
			//redirecting to seller page
			http.Redirect(w, req, "/seller", http.StatusSeeOther)
			return
		}
	}
	config.TPL.ExecuteTemplate(w, "index.gohtml", myUser)
}

/*This method is used to signup a new user
It checks if user has already logged in, and if so redirects to index page. If not logged in, the signup form is displayed.
The form input values from signup form are validated and sanitized before adding the data into the database
 Cryptic messages are displayed on the template in case of any errors in input.*/

func SignupHandler(w http.ResponseWriter, req *http.Request) {
	// Go to landing page of every user, if user has already logged in
	myUser := GetUser(w, req)
	if ok := alreadyLoggedIn(req); ok {
		if myUser.Username == "admin" {
			//redirecting to admin page
			http.Redirect(w, req, "/admin", http.StatusSeeOther)
			return
		} else if myUser.IsBuyer {
			//redirecting to buyer page
			http.Redirect(w, req, "/buyer", http.StatusSeeOther)
			return
		} else {
			//redirecting to seller page
			http.Redirect(w, req, "/seller", http.StatusSeeOther)
			return
		}
	}
	var user UserInfo
	// For validating user input
	si := signupInput{}

	// For displaying message to template
	sErrMsg := signupErrorMessage{}

	// process form submission
	if req.Method == http.MethodPost {
		reset := req.FormValue("reset")
		signup := req.FormValue("signup")

		username := req.FormValue("username")
		password := req.FormValue("password")
		isbuyer, _ := strconv.ParseBool(req.FormValue("isbuyer"))
		address := req.FormValue("address")
		fullname := req.FormValue("fullname")
		phone := req.FormValue("phone")
		email := req.FormValue("email")

		if reset != "" {
			//reload page
			sErrMsg.Mainmessage = "All values Reset"
		} else if signup != "" {
			// get form values

			if username == config.AdminName { // Username cannot be same as admin name
				sErrMsg.Mainmessage = "Invalid Username/Password"
				return
			} else {
				// Validate all input data for security
				si.Username = username
				si.Password = password
				si.Fullname = fullname
				si.Address = address
				si.Phone = phone
				si.Email = email

				verr := signupValidation(si)
				if verr != nil { // error
					sErrMsg.Mainmessage = "Invalid Username/Password"
				} else {
					// Sanitize input data
					snInput, serr := signupSanitization(si)

					if !serr { // error during sanitization
						config.Error.Println("Sanitization error")
						sErrMsg.Mainmessage = "Invalid Username/Password"
					} else {
						// Check if username is already taken
						if _, ok := mapUsers[snInput.Username]; ok {
							// Invalid username or password message to template
							sErrMsg.Mainmessage = "Invalid Username/Password"
							return
						} else {
							// This is a new sign up
							// Validated all entries
							pwBytes, err := bcrypt.GenerateFromPassword([]byte(snInput.Password), bcrypt.MinCost)
							if err != nil {
								sErrMsg.Mainmessage = "Internal server error, try again!"
								return
							}

							// Used to access user database
							var newUserdb user_db.UserDetails
							// use sanitized versions of entries
							newUserdb.Username = snInput.Username
							newUserdb.Password = string(pwBytes)
							newUserdb.Isbuyer = isbuyer
							newUserdb.Fullname = snInput.Fullname
							newUserdb.Address = snInput.Address
							newUserdb.Phone = snInput.Phone
							newUserdb.Email = snInput.Email

							// Write username and password to the user database
							// Write to user_db
							insertok := user_db.InsertRecord(config.DB, newUserdb)
							if !insertok {
								// Unable to write to database
								sErrMsg.Mainmessage = "Unable to reach database, try again!"
								return
							}

							// This is a local record of all users that have registered- buyers as well as sellers
							// Password is not stored here for security reasons
							user.Username = snInput.Username
							user.IsBuyer = isbuyer

							// Store user info in map for users
							mapUsers[snInput.Username] = user
							// Successfully registered, now redirect to login page
							http.Redirect(w, req, "/login", http.StatusSeeOther)
							return
						}
					}
				}
			}
		}
	}
	config.TPL.ExecuteTemplate(w, "signup.gohtml", sErrMsg)
}

/*This method is used to login seller/buyer
Before each login, the MapSessions data is cleaned, if it has not been cleaned for more than cleanSessionTime
The username and password input from the login form is validated and sanitized.
Any errors create a cryptic message that "username/password do not match"
If login is successful, a secure cookie is created, and mapsession is updated for seller/buyer*/
func LoginHandler(w http.ResponseWriter, req *http.Request) {

	//Clean mapSessions before login if session cleaning is not done for CleanSessionTime
	if time.Now().Sub(mapSessionCleaned) > (time.Second * time.Duration(config.CleanSessionTime)) {
		cleanupSessions()
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
		verr := loginValidation(lInput)

		if verr != nil { // error
			loginmessage = "Username and/or password do not match"
		} else {
			// Sanitize input data
			err := loginSanitization(lInput)

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
					// Check if admin login
					if lInput.Username == config.AdminName {
						//err := bcrypt.CompareHashAndPassword([]byte(config.AdminPw), []byte(lInput.Password))
						//fmt.Println(err)
						//if err != nil {
						if config.AdminPw != lInput.Password {
							loginmessage = "Username and/or password do not match"
						} else {

							// create session
							//id, _ := uuid.NewV4()
							id, err := uuid.NewV4()
							if err != nil {
								config.Warning.Printf("Unable to generate UUID : %v\n", err)
								loginmessage = "Server Error, please try again!"
							} else {

								myCookie := &http.Cookie{
									Name:     "PeelRescue",
									Value:    id.String(),
									MaxAge:   config.SessionExpireTime,
									HttpOnly: true,
									Secure:   true,
								}
								http.SetCookie(w, myCookie)

								mapSessions[myCookie.Value] = Session{lInput.Username, time.Now()}
								http.Redirect(w, req, "/admin", http.StatusSeeOther)
								return
							}
						}
					} else {
						// check if user exists with given username
						if _, ok := mapUsers[lInput.Username]; !ok { // username is not registered
							loginmessage = "Username and/or password do not match"
						} else {
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
									config.Warning.Printf("Password does not match the one in db\n")
									loginmessage = "Username and/or password do not match"
								} else {
									// create session
									//id := uuid.NewV4()

									id, err := uuid.NewV4()
									if err != nil {
										config.Warning.Printf("Unable to generate UUID : %v\n", err)
										loginmessage = "Server Error, please try again!"
									} else {
										//fmt.Println("Login is correct:")
										myCookie := &http.Cookie{
											Name:     "PeelRescue",
											Value:    id.String(),
											MaxAge:   config.SessionExpireTime,
											HttpOnly: true,
											Secure:   true,
										}
										http.SetCookie(w, myCookie)
										//fmt.Println("new mycookie : ", myCookie.Value, myCookie.MaxAge)

										mapSessions[myCookie.Value] = Session{lInput.Username, time.Now()}
										// Check if user is buyer or seller
										//fmt.Println("map session created")
										//for key, value := range mapSessions {
										//	fmt.Println(key, value)
										//}
										isbuyer = mapUsers[lInput.Username].IsBuyer
										//fmt.Println("Is buyer value ", isbuyer)
										if isbuyer {
											//fmt.Println("redirecion to buyer page")

											http.Redirect(w, req, "/buyer", http.StatusSeeOther)
											return
										} else {
											//fmt.Println("redirecion to seller page")
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
		}
	}
	config.TPL.ExecuteTemplate(w, "login.gohtml", loginmessage)
}

/*This method is used to logout buyer or seller
If no cookie is present, the page is redirected to login page
If cookie is present, it is deleted, and corrspoding map session entry is also deleted*/
func LogoutHandler(w http.ResponseWriter, req *http.Request) {
	logoutmessage := "Log out successful! \nThank you for using Peel Rescue!\nKudos! You are saving Earth one Peel at a Time!"

	myCookie, err := req.Cookie("PeelRescue")
	//fmt.Println(myCookie.Value, myCookie.MaxAge)
	if err != nil {
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		config.Warning.Println("Cookie not found, redirecting to login page..")
		config.Trace.Println("Cookie not found, redirecting to login page..")
		return
	} else {
		config.Info.Println("Deleting cookie and clearing session..")

		// delete the session
		delete(mapSessions, myCookie.Value)
		myCookie = &http.Cookie{
			Name:     "PeelRescue",
			Value:    "",
			MaxAge:   -1,
			HttpOnly: true,
			Secure:   true,
		}
		http.SetCookie(w, myCookie)
	}
	config.TPL.ExecuteTemplate(w, "logout.gohtml", logoutmessage)
}

/*This method is used to check if the session is active
It checks if user has already logged in, returns false if user has not logged in.
It checks if time out has occurred, returns false if time out has occurred.
If user has loggeg in, and time out has not occured, then the session is extended, and it returns true.*/
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

/*This method checks if the user has already logged in
The user can be a buyer or a seller
If user has logged in, it returns true, else false*/
func alreadyLoggedIn(req *http.Request) bool {
	myCookie, err := req.Cookie("PeelRescue")
	if err != nil || myCookie.MaxAge < 0 {
		return false
	}
	username := mapSessions[myCookie.Value].Username
	if username != "" {
		return true
	} else {
		return false
	}
}
