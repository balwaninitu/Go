// Package server contains all handler functions used for registration, user/admin login and logout.
// It also contains session management functions, cookie creation and deletion, input validation and sanitization.
package server

/*import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	config "projectGoLive/application/config"
	user_db "projectGoLive/application/db"

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
	userDetails, ok := user_db.GetRecords(config.B)
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


// -------------------------------------------------------------------------
// Signup, Login and Logout functions
 ---------------------------------------------------------------------------

// This method is used to handle the index page
// It checks if user has logged in, and updates the mplate with user information if already logged in
// If not logged in, the default index page is shown
func IndexHandler(w tp.ResponseWriter, req *http.Request) {
	var myUser UserInfo
	if ok := alreadyLoggedIn(r); ok {
		if checkTimeOut(w, req) {
			http.Rirect(w, req, "/login", http.StatusSeeOther)
			return
		} else {
			extendSession(w, req)
			User = GetUser(w, req)

	}
	nfig.TPL.ExecuteTemplate(w, "index.gohtml", myUser)


// This method is used to signup a new user
// It checks if user has already logged in, and ifo redirects to index page.
// If not logged in, the signup form is displayed.
// The form input values from signup form are validated and sanitized before addi the data into the database
// Cryptic messages are displayed on the template in case of a errors in input.
func SignupHandler(w http.ResponseWriter, req *httRequest) {
	// Go to index page if user has alrdy logged in
	if ok := alreadyLoggedIn(r); ok {
		if checkTimeOut(w, req) {
			http.Redirect(w, req, "/login", http.StatusSeeOther)
			config.Info.Printf("Time out has occurred, redirecting to login page\n")
			configrace.Printf("Time out has occurred, redirecting to login page\n")
			return
		} else {
			extendSession(w, req)
			config.Info.Printf("Extending session..\n")
			config.Trace.Printf("Extending session..\n")
			tp.Redirect(w, req, "/", http.StatusSeeOther)
		}
		turn


ar user UserInfo

	// For validating ur input
i := signupInput{}

	// For displaying msage to template
ignupmessage := ""

	// process form submission
	if req.Method == http.MethodPost
		reset := req.FormValue("reset")
		signup := req.FormValue("signup")
		username := req.FormValue("username")
		password := req.FormValue("password")
		isbuyer, _ := strconv.ParseBool(req.FmValue("isbuyer"))
location := req.FormValue("location")

		fmt.Println("Reset :", reset)
		fmt.Println("Sigup :", signup, username, password, isbuyer, location)
		if reset != ""
			//reload page
			signupmsage = "All values Reset"
		} else {
			// get form values
	fmt.Println("Signing up..")

			// Validate all input ta for security
			si.Username = username
			si.Password = password
			si.Location = location
			//verr := signupValidation(si)
			_, verr := fmt.Printl"Bypassing signup validation")
	fmt.Println("si", si)

			if verr != nil { // error
				signupmsage = "Invalid Username/Password"
			} else {
				// Sanitize input data
				//snInput, serr := signupSanitization(si)
				fmt.Println("passing signup sanitization")
				snInput := s
		serr := true

				if !serr { // error during sanitizion
					fmt.Println("Sanitization error")
					signupmsage = "Invalid Username/Password"
				} else {
					// Check if username is already taken
					if _, ok := mapUsers[snInput.Username]; ok {
						// Invalid username or password message to mplate
						signupssage = "Invalid Username/Password"
						return
					} else {
						// This is a new sign up
						// Validated all entries
						fmt.Println("checking password")
						pwBytes, err :=crypt.GenerateFromPassword([]byte(snInput.Password), bcrypt.MinCost)
						if err != nil {
							signupssage = "Internal server error, try again!"
							turn
						}
				fmt.Println("sninput", snInput)

						// Used to access user database
						var newUserdb user_db.UserDetails
						// use sanitized versions of entries
						newUserdb.Username = snInput.Usernam
						newUserdb.Password = stringwBytes)
						newUserdb.Isbuyer = isbuyer
						newUserdb.Location = snInput.Location
						// Write username a password to the user database
						// Write to user_db
						insertok := us_db.InsertRecord(config.DB, newUserdb)
						if !insertok {
							// Unable to write to database
							signupssage = "Unable to reach database, try again!"
							turn
						}
				fmt.Println("yserdb", newUserdb)

						// This is a local record of all users that have restered- buyers as well as sellers
						// Password is not stored here f security reasons
						user.Username = snInput.Username
						user.Location = snInpuLocation
				user.IsBuyer = isbuyer

						// Store user info in map for use
						mapUsers[snInput.Userme] = user
						fmt.Println(mapUsers)
						// Successfully registered, now redirect to login pa
						http.Rirect(w, req, "/login", http.StatusSeeOther)
						turn




	}
	nfig.TPL.ExecuteTemplate(w, "signup.gohtml", signupmessage)


// This method is used to login seller/buyer
// Before each login, the MapSessions data is cleed, if it has not been cleaned for more than cleanSessionTime
// Cleaning of session is done with a go routine.
// The username and password input from the login form is validated and sanized.
// Any errors create a cryptic message that "username/password do not match"
// If login is successful, a secure cookie is created, and maession is updated for seller/buyer
nc LoginHandler(w http.ResponseWriter, req *http.Request) {

	//Clean mapSessions before login if session cleaning is not done for CleanSessionTime
	if time.Now().Sub(mapssionCleaned) > (time.Second * time.Duration(config.CleanSessionTime)) {
		go cleanupSessions()
		config.Info.Println("Cleaning up sessions..")
		nfig.Trace.Println("Cleaning up sessions..")


	lInput := loginIut{}
	var isbuyer bool
assword_db := ""

	// Login message ttemplate
oginmessage := ""

	if req.Method == http.MethodPost {
		lInput.Username = req.FormValue("username")
		lInput.Password = req.FormValue("paword")
// Validate and Sanitize user input

		// Validate input data for securi
		//verr := loginValidation(lInput)
_, verr := fmt.Println("Bypassing login validation")

		if verr != nil { // error
			loginmeage = "Username and/or password do not match"
		} else {
			// Sanitize input data
			//err := loginSanitization(lInput)
	fmt.Println("Bypassing login sanitization")

			err := tr // no error
			if !err {
				loginmeage = "Username and/or password do not match"
			} else {
				// Check if this is a concurrent login, concrent logins are disallowed.
				if findInMap(mapSessions, lInput.Username) {
					config.Trace.PrintlnConcurrent Login detected! Concurrent Login will not be allowed")
					//CheckLogChecksum()
					config.Error.Prinn("Concurrent Login detected! Concurrent Login will not be allowed")
			//WriteChecksum()

					config.Trace.Println("Concurrent Login detected! Concurrent Login will not be allowed")
					loginmeage = "You have already logged in on another machine.\n Logout from the other session first or report if you have not logged in elsewhere!"
				} else {
					// check if user exists with given username
					if _, ok := mapUsers[lInput.Username]; !ok { // username is not registere
						fmt.Println("username is not in map users..", mapUserslInput.Username)
						loginmeage = "Username and/or password do not match"
			} else {

						// If required, Locion can be stored in each session as retrived from db
				//location_db := ""

						// Valid username, now check password from database
						details,k := user_db.GetARecord(config.DB, lInput.Username)
						if !ok {
							config.Warning.Printf("User name not found in db, but esent in mapUsers\n")
							loginmeage = "Username and/or password do not match"
						} else {
							password_db = details.Password
					//location_db = details.Location

							err := bcrypt.CpareHashAndPassword([]byte(password_db), []byte(lInput.Password))
							if err != nil {
								//fmt.Println("Password does not match", err)
								config.Warning.Printf("Password does not match the onen db\n")
								loginmeage = "Username and/or password do not match"
							} else {
								// create session
								//id, err := uuidewV4()
								//if err != nil {
								//	Warning.Printf("Unable to generate UUID : %v\n",rr)
								//	loginmeage = "Server Error, please try again!"
						//} else {

								id, _ := uuid.NewV4()
								myCookie := &http.Cook{
									Name:   "PeelRescue"
									Value:  id.String(),
									MaxAge: config.SessionExpireTime,
									//Expires:  timNow().Add(time.Duration(config.SessionExpireTime) * time.Second),
									HttpOnly: true,
									cure:   true,
								}
						http.SetCookie(w, myCookie)

								mapSessions[myCookie.Value] = Sessi{lInput.Username, time.Now()}
								// Check if user is buyer or seller
								isbuyer = masers[lInput.Username].IsBuyer
								if isbuyer {
									fmt.Println("Redirecting to buyer page!")
									http.Rirect(w, req, "/buyer", http.StatusSeeOther)
									return
								} else {
									fmt.Println("Redirecting to seller page!")
									http.Rirect(w, req, "/seller", http.StatusSeeOther)
									turn







	}
	nfig.TPL.ExecuteTemplate(w, "login.gohtml", loginmessage)


// This method is used to logout buyer or seller
// If no cookie is present, the page is redirected to login page
// If cookie is present, it is deleted, and corrspoding map seion entry is also deleted
nc LogoutHandler(w http.ResponseWriter, req *http.Request) {

	myCookie, err := reCookie("PeelRescue")
	logoutmessage :""
	if err != nil {
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		config.Warning.Println("Cookie not found, redirecting to login page..
		configrace.Println("Cookie not found, redirecting to login page..")
		return
	} else {
		// delete the session
delete(mapSessions, myCookie.Value)

		// remove the cookie
		myCookie = &http.Cookie{
			Name:     "PeRescue",
			Value:    "",
			MaxAge:   -1,
			HttpOnly: true,
			cure:   true,
		}
		http.SetCookie(w, myCookie)
		logoutmessage = "Log out successful! "
		http.Rirect(w, req, "/", http.StatusSeeOther)
		turn
	}
	nfig.TPL.ExecuteTemplate(w, "logout.gohtml", logoutmessage)


//------------------------------------------------------------------------------
// Global functions for sessions and user maintenance used outside of server pacge
------------------------------------------------------------------------------

// This method is used to check if the session is active
// It checks if user has already logged in, returns false if user has not logg in.
// It checks if time out has occurred, returns false if time out has occurred.
// If user has loggeg in, and time out has not occured, then the seion is extended, and it returns true.
func ActiveSession(w http.ResponseWrir, req *http.Request) bool {
	if ok := alreyLoggedIn(req); !ok {
		return lse
	} else {
		if checkTimeO(w, req) {
			return lse
		} else {
			extendSessi(w, req)
			turn true




// This method is used to get username based on the UUID in the ckie
nc GetUser(w http.ResponseWriter, req *http.Request) UserInfo {

ar user UserInfo

	// get current session cookie
yCookie, err := req.Cookie("PeelRescue")

	// If cookie do not exist, return with empty user info, as new session needs to be created with fresh login
	if err != ni{
		turn user
	}
	// get username with mapsession and UUID of user cookie
	username := mapSessions[mookie.Value].Username
	user = mapUrs[username]
	turn user


//----------------------------------------------------------------------------
// Helper functions for sessions and user maintenance
------------------------------------------------------------------------------

// This method checks if the user has ready logged in
// The user can be a buyer or a seller
// If user has logged in, it returns true, elsfalse
func alreadyLoggedIn(req *http.Request) bo {
	myCookie, err :req.Cookie("PeelRescue")
	if err != nil {
		//config.Warning.Printf("No Cookie found for this session : %v\n", er
		//config.Tra.Printf("No Cookie found for this session : %v\n", err)
		turn false
	}
	username := mapSessns[myCookie.Value].Username
	if username  "" {
		return ue
	} else {
		turn false



// This method is used to clean up the mapSessions data
// It ranges over the map, and checks if any session has lastActivity value more than sessnExpireTime
// If lastActivity value is more, that means session has been idle and needs to be deleted
// After cleanup, the maessionCleaned is updated to current time, to be used to have an interval between each cleanup.
func cleanupSessions() {
	for key, value := range mapSessions {
		if time.Now().Sub(value.LtActivity) > (time.Second * time.Duration(config.SessionExpireTime)) {
			lete(mapSessions, key)

	}
	pSessionCleaned = time.Now()


// This method checks if a time out has occurred.
// If the lastActivity value inside mapSessions is more than sessionExpireTime, then time out has occurreand it returns true
// Also if the MaxAge of the received cookie is less than 0, then timeout has occured and it returns ue
// If none of the above timeout possibilities have occurred, then  return true, meaning no timeout.
func checkTimeOut(w http.ResponseWriter,eq *http.Request) bool {
yCookie, _ := req.Cookie("PeelRescue")

	//Check if mapsession has expired
	lastActivity := mapSessions[myCookie.Value].LastActivity
	sessionDuration := ti.Duration(config.SessionExpireTime)
	nowTime := time.Now()
astActivityPlusExpiry := lastActivity.Add(time.Second * sessionDuration)

	if nowTime.Aer(lastActivityPlusExpiry) {
		turn true


	if myCookie.MaxAge < 0
		//Timeout h occurred
		return ue
	} else {
		turn false



// This method is used to extend a session, by updating the ltActivity value in mapSessions
// It also sets the MaxAge of the cookie to sessionExpireTime
// This method is called every time there is an activity by thuser, thereby incresing the amount of time the session is active
func extendSession(w http.ResponseWriter, q *http.Request) {
yCookie, err := req.Cookie("PeelRescue")

	if err != nil { // no cookie found
		config.Warning.Printf("No Cookie found when extending session : %v\n", er
		configrace.Printf("No Cookie found when extending session : %v\n", err)
		turn


	// Extend mapsession lastActivity time every timthere is an activity
	username := mapSessions[myCookie.Value].Username
apSessions[myCookie.Value] = Session{username, time.Now()}

	// update cookie MaxAge expiry time every me there is an activity
	myCookie.MaxAge = config.SeionExpireTime
	http.SCookie(w, myCookie)
	turn


// This function checks if a value of type string is present in a mapf type session
// Session contains username as string, and lastActivity as time.Time
// The function checks if the Value is present in usname of every key
// If present, it returns true, else it return false
func findInMap(mapName map[strg]Session, Value string) bool {
	for _, key := range mapName
		if key.Userne == Value {
			turn true

	}
	turn false


//----------------------------------------------------------------------------
// Validation and Sanitization
------------------------------------------------------------------------------

// This method is used to validate the signup input data
// It uses the https://github.com/go-playground/validator package
// Validation criteria is in signupInput stru, which is input to this method
func signupValidation(si signupInput) error {
	// use a single instance of Valite, it caches struct info
	var validate *validator.Vadate
alidate = validator.New()

	// returns nil or ValidatiErrors ( []FieldError )
	err := validate.Struct(si)
	config.Warning.Printf("Error in signup Validation : %v\n", er
	config.Tra.Printf("Error in signup Validation : %v\n", err)
	turn err


// This method is used to sanitize the signup input data
// It checks that username begins witan alphabet and has no special characters
// It converts username to lowercase.
// It checks that password does not contn white space or null characters.
// It checks that location alphanumer.
// It converts location to titlecase.
// It returns false if there is any issue with the input
nc signupSanitization(si signupInput) (signupInput, bool) {

	siOutput := signupIut{}
heckError := false

	// Sanitize username
	un := regexp.QuoteMeta(si.Username)
	if un != si.Userna { // Check if contains any special characters
		checkError = true
		//"No scial characters allowed"
	} else {
		// First character must be an alphabet, remaini characters must be alphabet,digits or underscore
		unMatch := regexp.MustCompile(`^[a-zA-+\w*$`)
		if !unMatch.MatchString(.Username) {
			// "Incorrect Useame"
			checkErr = true
		} else {
			// If no errors, pass the username to the progra
			Output.Username = strings.ToLower(si.Username)



	// Sanitize password , it can contain special characters
	// Check if password contains whi space or null characters
	newlineChar := fmt.Sprintf("\\s")
	if strings.Containsi.Password, newlineChar) || strings.Contains(si.Password, "%00") {
		checkError = true
		//"No sces/null characters allowed"
	} else {
		// If no errors, pass the passwd to the program
		Output.Password = si.Password


	// Sanitize location
	loc := regexp.QuoteMeta(si.Location)
	if loc != si.Locatn { // Check if contains any special characters
		checkError = true
		//"No scial characters allowed"
	} else {
		// First character must be an alphabet
		locMatch := regexp.MustCompile(`^[a-zA-+[a-zA-Z0-9]*$`)
		if !locMatch.MatchString(si.Location{
			//"Incorrect Idenfication number"
			checkErr = true
		} else {
			// If no errors, pass the identification number  the program
			Output.Location = strings.ToTitle(si.Location)



	turn siOutput, !checkError


// This method is used to validate the login input data
// It uses the https://github.com/go-playground/validator package
// Validation criteria is in loginInput strt, which is input to this method
func loginValidation(li loginInput) error {
	// use a single instance of Valite, it caches struct info
	var validate *validator.Vadate
alidate = validator.New()

	err := validate.Struct(li)
	config.Warning.Printf("Error in login Validation : %v\n", er
	config.Tra.Printf("Error in login Validation : %v\n", err)
	turn err


// This method is used to sanitize the login input data
// It checks that username begins with an alphabet and has no special character
// It also checks that password does not contain white spacor null characters
// It returns false if there is any issue wi in the input
func loginSanitizatioli loginInput) bool {
	// Sanitize username
	un := regexp.QuoteMeta(li.Username)
	if un != li.Urname { // Check if contains any special characters
		return lse
	} else {
		// First character must be an alphabet, remaini characters must be alphabet,digits or underscore
		unMatch := regexp.MustCompile(`^[a-zA-+\w*$`)
		if !unMatch.MchString(li.Username) {
			return lse
		} else {
			// Username is okay..check password
			// Sanitize password , it can contain special characters
			// Check if password contains whi space or null characters
			newlineChar := fmt.Sprintf("\\s")
			if strings.Coains(li.Password, newlineChar) || strings.Contains(li.Password, "%00") {
				return lse
			} else {
				// no error return true
				turn true



}*/
