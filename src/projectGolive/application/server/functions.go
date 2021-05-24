/*Functions files includes validation and sanitizations functions
for user input values.*/
package server

import (
	"fmt"
	"net/http"
	config "projectGoLive/application/config"
	user_db "projectGoLive/application/user_db"
	"regexp"
	"strings"
	"time"

	"github.com/go-playground/validator"
)

// This struct stores username(all) and type of user(buyer/seller). It is used to maintain a list of all users currently registered.
type UserInfo struct {
	Username string
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
	Password string `validate:"required,min=4,max=100"`
}

// Struct to store data sent by sign up template form for validation
type signupInput struct {
	Username string `validate:"required"`
	Password string `validate:"required,min=4,max=100"`
	Fullname string `validate:"required"`
	Phone    string `validate:"required,len=8,numeric"`
	Address  string `validate:"required"`
	Email    string `validate:"required,email,contains=@"`
}

// Struct to store data sent to template for displaying user signup errors
type signupErrorMessage struct {
	Mainmessage string
	Username    string
	Password    string
	Fullname    string
	Phone       string
	Address     string
	Email       string
}

func init() {
	mapUsers = make(map[string]UserInfo)
	mapSessions = make(map[string]Session)

	// Read user_db and populate mapUsers (otherwise buyers/seller in database cannot login)
	initializeMapUsers()
}

//checking user credentials available in database
func initializeMapUsers() {
	userDetails, ok := user_db.GetRecords(config.DB)
	if !ok {
		panic("Cannot read user_db")
	} else {
		var myUserInfo UserInfo
		for _, ui := range userDetails {
			myUserInfo.Username = ui.Username
			myUserInfo.IsBuyer = ui.Isbuyer
			mapUsers[ui.Username] = myUserInfo
		}
	}

}

/*This method is used to clean up the mapSessions data
It ranges over the map, and checks if any session has
lastActivity value more than sessionExpireTime If lastActivity value is more,
that means session has been idle and needs to be deleted
After cleanup, the mapSessionCleaned is updated to current time,
to be used to have an interval between each cleanup.*/
func cleanupSessions() {
	for key, value := range mapSessions {
		if time.Now().Sub(value.LastActivity) > (time.Second * time.Duration(config.SessionExpireTime)) {
			delete(mapSessions, key)
		}
	}
	mapSessionCleaned = time.Now()
}

/*This method checks if a time out has occurred.
If the lastActivity value inside mapSessions is more than sessionExpireTime, then time out has occurred and it returns true
Also if the MaxAge of the received cookie is less than 0, then timeout has occured and it returns true
If none of the above timeout possibilities have occurred, then it return true, meaning no timeout.*/
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

/*This method is used to extend a session, by updating the lastActivity value in mapSessions
 It also sets the MaxAge of the cookie to sessionExpireTime
This method is called every time there is an activity by the user, thereby incresing the amount of time the session is active*/
func extendSession(w http.ResponseWriter, req *http.Request) {
	myCookie, err := req.Cookie("PeelRescue")

	if err != nil || myCookie.MaxAge < 0 { // no cookie found
		config.Warning.Printf("No Cookie found when extending session : %v\n", err)
		config.Trace.Printf("No Cookie found when extending session : %v\n", err)
		return
	}
	// Extend mapsession lastActivity time every time there is an activity
	username := mapSessions[myCookie.Value].Username
	if username == "" {
		config.Warning.Printf("Session is invalid")
		return
	}

	mapSessions[myCookie.Value] = Session{username, time.Now()}

	// update cookie MaxAge expiry time every time there is an activity
	myCookie.MaxAge = config.SessionExpireTime
	http.SetCookie(w, myCookie)
	return
}

/*This function checks if a value of type string is present in a map of type session
Session contains username as string, and lastActivity as time.Time
The function checks if the Value is present in username of every key
If present, it returns true, else it return false*/
func findInMap(mapName map[string]Session, Value string) bool {
	for _, key := range mapName {
		if key.Username == Value {
			return true
		}
	}
	return false
}

/*This method is used to validate the signup input data
It uses the https://github.com/go-playground/validator package
Validation criteria is in signupInput struct, which is input to this method*/
func signupValidation(si signupInput) error {
	// use a single instance of Validate, it caches struct info
	var validate *validator.Validate
	validate = validator.New()

	// returns nil or ValidationErrors ( []FieldError )
	err := validate.Struct(si)
	if err != nil {
		config.Warning.Printf("Error in signup Validation : %v\n", err)
		config.Trace.Printf("Error in signup Validation : %v\n", err)
	}
	return err
}

/*This method is used to sanitize the signup input data
It checks that username begins with an alphabet and has no special characters
It converts username to lowercase.
It checks that password does not contain white space or null characters.
It checks that location alphanumeric.
It converts location to titlecase.
It returns false if there is any issue with the input*/
func signupSanitization(si signupInput) (signupInput, bool) {

	siOutput := si
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

	return siOutput, !checkError
}

/*This method is used to validate the login input data
It uses the https://github.com/go-playground/validator package
Validation criteria is in loginInput struct, which is input to this method*/
func loginValidation(li loginInput) error {
	// use a single instance of Validate, it caches struct info
	var validate *validator.Validate
	validate = validator.New()

	err := validate.Struct(li)
	if err != nil {
		config.Warning.Printf("Error in login Validation : %v\n", err)
		config.Trace.Printf("Error in login Validation : %v\n", err)
	}
	return err
}

/*This method is used to sanitize the login input data
It checks that username begins with an alphabet and has no special characters
It also checks that password does not contain white space or null characters
It returns false if there is any issue with in the input*/
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
