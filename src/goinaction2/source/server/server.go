// Package server contains all handler functions used for registration, user/admin login and logout.
// It also contains session management functions, cookie creation and deletion, input validation and sanitization.
package server

import (
	"encoding/hex"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	age "github.com/bearbin/go-age"
	validator "github.com/go-playground/validator"
	uuid "github.com/satori/go.uuid"
	bcrypt "golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/blake2b"

	//In built packages
	bt "vaccineappt/source/btree"
	fh "vaccineappt/source/filehandler"
	user "vaccineappt/source/person"
	userll "vaccineappt/source/personll"
)

const (
	//Time Format constant
	TimeFormatISO = "2006-01-02"
	Kitchen       = "3:04PM"

	// Expiry time in seconds = 1200 seconds = 20 minutes
	sessionExpireTime int = 1200

	// Time to clean MapSessions 120 seconds = 2 minutes
	cleanSessionTime int = 120

	//Age Criteria for qualifying for vaccination
	ageQualification = 70

	// Paths
	filePath = "./"
	csvPath  = "./csv/"
	tmpPath  = "./tmp/"
	htmlPath = "./html/"
	certPath = "./cert/"
	logPath  = "./log/"
)

var (
	tpl *template.Template
	Wg  *sync.WaitGroup
)

var (
	//Loggers used to log Trace,Info, Warning and Error in server package
	Trace   *log.Logger // Just about anything
	Info    *log.Logger // Important information
	Warning *log.Logger // Be concerned
	Error   *log.Logger // Critical problem
)

// This struct stores username and lastActivity, and is used to maintain sessions
type Session struct {
	Username     string
	LastActivity time.Time
}

// This struct stores the admin username and password as read from csv file
type AdminLoginStruct struct {
	AdminName     string
	AdminPassword string
}

// This struct stores the AppData containing pointers to linked list and BST structures
// It also contains the appointment array, map of sessions, time map of sessions is cleaned, and admin login info
type AppData struct {
	PersonList        *userll.LinkedList
	BstUserName       *bt.BST
	BstID             *bt.BST
	ApptArray         []time.Time
	MapSessions       map[string]Session
	MapSessionCleaned time.Time
	AdminLogin        AdminLoginStruct
}

// Struct to store data sent to template for making/deleting appointments by user
type userapptstruct struct {
	Firstname     string
	Lastname      string
	Message       []string
	Apptinfo      string
	Possibleappts []string
}

// Struct to store data sent to template for displaying user registration errors
type registerError struct {
	Mainmessage    string
	Firstname      string
	Lastname       string
	Identification string
	Username       string
	Password       string
	Dob            string
	Phone          string
	Address        string
	Email          string
}

// Struct to store data sent by registration template form for validation
type registerInput struct {
	Firstname      string `validate:"required,alphanumunicode"`
	Lastname       string `validate:"required,alphanumunicode"`
	Identification string `validate:"required,min=8,max=10,alphanum"`
	Username       string `validate:"required,alphanum"`
	Password       string `validate:"required,min=8,max=100"`
	Dob            string `validate:"required"`
	Phone          string `validate:"required,len=8,numeric"`
	Address        string `validate:"printascii"`
	Email          string `validate:"required,email,contains=@"`
}

// Struct to store data sent by login template form for validation
type loginInput struct {
	Username string `validate:"required,alphanum"`
	Password string `validate:"required,min=8,max=100"`
}

// Struct to store data sent to template for performing admin functions
type adminStruct struct {
	Message    []string
	Users      []string
	Deleteuser string
	ApptAdd    string
}

// Function to initialize the templates for use within server package
func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

// This method is used to clean up the mapSessions data
// It ranges over the map, and checks if any session has lastActivity value more than sessionExpireTime
// If lastActivity value is more, that means session has been idle and needs to be deleted
// After cleanup, the mapSessionCleaned is updated to current time, to be used to have an interval between each cleanup.
func (a *AppData) cleanupSessions() {
	for key, value := range a.MapSessions {
		if time.Now().Sub(value.LastActivity) > (time.Second * time.Duration(sessionExpireTime)) {
			delete(a.MapSessions, key)
		}
	}
	a.MapSessionCleaned = time.Now()
}

// This method is used to check if the session is active
// It checks if user has already logged in, returns false if user has not logged in.
// It checks if time out has occurred, returns false if time out has occurred.
// If user has loggeg in, and time out has not occured, then the session is extended, and it returns true.
func (a *AppData) activeSession(w http.ResponseWriter, req *http.Request) bool {
	if ok := a.alreadyLoggedIn(req); !ok {
		return false
	} else {
		if a.checkTimeOut(w, req) {
			return false
		} else {
			a.extendSession(w, req)
			return true
		}
	}
}

// This method checks if a time out has occurred.
// If the lastActivity value inside MapSessions is more than sessionExpireTime, then time out has occurred and it returns true
// Also if the MaxAge of the received cookie is less than 0, then timeout has occured and it returns true
// If none of the above timeout possibilities have occurred, then it return true, meaning no timeout.
func (a *AppData) checkTimeOut(w http.ResponseWriter, req *http.Request) bool {
	myCookie, _ := req.Cookie("VaccineAppt")

	//Check if mapsession has expired
	lastActivity := a.MapSessions[myCookie.Value].LastActivity
	sessionDuration := time.Duration(sessionExpireTime)
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

// This method is used to extend a session, by updating the lastActivity value in MapSessions
// It also sets the MaxAge of the cookie to sessionExpireTime
// This method is called every time there is an activity by the user, thereby incresing the amount of time the session is active
func (a *AppData) extendSession(w http.ResponseWriter, req *http.Request) {
	myCookie, err := req.Cookie("VaccineAppt")

	if err != nil { // no cookie found
		Warning.Printf("No Cookie found when extending session : %v\n", err)
		Trace.Printf("No Cookie found when extending session : %v\n", err)
		return
	}

	// Extend mapsession lastActivity time every time there is an activity
	username := a.MapSessions[myCookie.Value].Username
	a.MapSessions[myCookie.Value] = Session{username, time.Now()}

	// update cookie MaxAge expiry time every time there is an activity
	myCookie.MaxAge = sessionExpireTime
	http.SetCookie(w, myCookie)
	return
}

// This method checks if the user has already logged in
// The user can be a regular user or an admin
// If user has logged in, it returns true, else false
func (a *AppData) alreadyLoggedIn(req *http.Request) bool {
	myCookie, err := req.Cookie("VaccineAppt")

	if err != nil {
		Warning.Printf("No Cookie found for this session : %v\n", err)
		Trace.Printf("No Cookie found for this session : %v\n", err)
		return false
	}

	username := a.MapSessions[myCookie.Value].Username
	if username == a.AdminLogin.AdminName {
		return true
	} else {
		// User is not admin, check in BST
		usernameNode := a.BstUserName.Search(username)
		if usernameNode == nil { // username is not in BST
			return false
		} else {
			return true
		}
	}
}

// This method is used to handle the index page
// It checks if user has logged in, and updates the template with user information if already logged in
// If not logged in, the default index page is shown
func (a *AppData) IndexHandler(w http.ResponseWriter, req *http.Request) {

	currentPerson := user.Person{}
	if ok := a.alreadyLoggedIn(req); ok {
		if a.checkTimeOut(w, req) {
			http.Redirect(w, req, "/login", http.StatusSeeOther)
			return
		} else {
			a.extendSession(w, req)
		}
		currentPerson = a.getPerson(w, req)
	}
	tpl.ExecuteTemplate(w, "index.gohtml", currentPerson)
}

// This method is used to register a new user
// It checks if user has already logged in, and if so redirects to index page.
// If not logged in, the registration form is displayed.
// The form values from registration form are validated and sanitized before adding the data into the data structures
// Cryptic messages are displayed on the template in case of any errors in input.
func (a *AppData) RegistrationHandler(w http.ResponseWriter, req *http.Request) {
	// Go to index page if user has already logged in
	if ok := a.alreadyLoggedIn(req); ok {
		if a.checkTimeOut(w, req) {
			http.Redirect(w, req, "/login", http.StatusSeeOther)
			Info.Printf("Time out has occurred, redirecting to login page\n")
			Trace.Printf("Time out has occurred, redirecting to login page\n")
			return
		} else {
			a.extendSession(w, req)
			Info.Printf("Extending session..\n")
			Trace.Printf("Extending session..\n")
			http.Redirect(w, req, "/", http.StatusSeeOther)
		}
		return
	}

	regInput := registerInput{}
	regError := registerError{}
	var newPerson user.Person

	// process form submission
	if req.Method == http.MethodPost {
		// get form values
		regInput.Firstname = req.FormValue("firstname")
		regInput.Lastname = req.FormValue("lastname")
		regInput.Identification = req.FormValue("identification")
		regInput.Username = req.FormValue("username")
		regInput.Password = req.FormValue("password")
		regInput.Dob = req.FormValue("dob")
		regInput.Phone = req.FormValue("phone")
		regInput.Address = req.FormValue("address")
		regInput.Email = req.FormValue("email")

		// Validate input data for security
		verr := registerValidation(regInput)

		if verr != nil { // error
			regError.Mainmessage = "Incorrect Input. Please check again!"
		} else {
			// Sanitize input data
			sregInput, sregError, serr := registerSanitization(regInput)
			regError = sregError // Overwrite regError with error from Sanitization, as Validation error will be empty to reach this point

			if !serr { //no error

				newPerson.Identification = sregInput.Identification
				newPerson.Username = sregInput.Username
				newPerson.Firstname = sregInput.Firstname
				newPerson.Lastname = sregInput.Lastname
				newPerson.Phone = sregInput.Phone
				newPerson.Dob = sregInput.Dob
				newPerson.Email = sregInput.Email
				newPerson.Address = sregInput.Address

				if newPerson.Username == a.AdminLogin.AdminName { // Username cannot be same as admin name
					regError.Mainmessage = "Incorrect Identification/Username/Password"
					return
				} else {
					idNode := a.BstID.Search(sregInput.Identification) // Search if ID is unique

					if idNode == nil { // ID is not in BST
						newPerson.Identification = sregInput.Identification
						usernameNode := a.BstUserName.Search(sregInput.Username) // Search if username is unique

						if usernameNode == nil { // username is not in BST

							newPerson.Username = sregInput.Username
							pwBytes, err := bcrypt.GenerateFromPassword([]byte(sregInput.Password), bcrypt.MinCost)
							if err != nil {
								//Password does not match
								regError.Mainmessage = "Incorrect Identification/Username/Password"
								return
							} else {

								newPerson.Password = string(pwBytes)
								// Check if person qualifies for vaccine
								if calculateAge(newPerson.Dob) >= ageQualification {
									newPerson.VaccinationQualify = true
								} else {
									newPerson.VaccinationQualify = false
								}
								// Everything  is OK, add person to linked list and username and ID to the two BSTs
								Wg.Add(3)
								go a.PersonList.AddNode(newPerson, Wg)
								go a.BstUserName.Insert(newPerson.Username, Wg)
								go a.BstID.Insert(newPerson.Identification, Wg)
								Wg.Wait()

								// Write user information to CSV file after successfull registration
								Wg.Add(2)
								fh.WritePersonCSVFile(Wg, a.PersonList)
								fh.WriteApptCSVFile(Wg, a.ApptArray)
								Wg.Wait()
								// redirect to main index
								http.Redirect(w, req, "/login", http.StatusSeeOther)
								return
							}

						} else {
							regError.Mainmessage = "Incorrect Identification/Username/Password"
							return
						}
					} else {
						regError.Mainmessage = "Incorrect Identification/Username/Password"
						return
					}
				}
			}
		}
	}
	tpl.ExecuteTemplate(w, "register.gohtml", regError)
}

// This method is used to login user/admin
// Before each login, the MapSessions data is cleaned, if it has not been cleaned for more than cleanSessionTime
// Cleaning of session is done with a go routine.
// The username and password input from the login form is validated and sanitized.
// Any errors create a cryptic message that "username/password do not match"
// If login is successful, a secure cookie is created, and mapsession is updated for user/admin
func (a *AppData) LoginHandler(w http.ResponseWriter, req *http.Request) {

	//Clean MapSessions before login if session cleaning is not done for cleanSessionTime
	if time.Now().Sub(a.MapSessionCleaned) > (time.Second * time.Duration(cleanSessionTime)) {
		go a.cleanupSessions()
		Info.Println("Cleaning up sessions..")
		Trace.Println("Cleaning up sessions..")
	}
	lInput := loginInput{}
	loginmessage := ""
	var currentPerson user.Person

	// process form submission
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

			if err {
				loginmessage = "Username and/or password do not match"
			} else {
				// Check if this is a concurrent login, concurrent logins are disallowed.
				if findInMap(a.MapSessions, lInput.Username) {
					Trace.Println("Concurrent Login detected! Concurrent Login will not be allowed")
					CheckLogChecksum()
					Error.Println("Concurrent Login detected! Concurrent Login will not be allowed")
					WriteChecksum()

					Trace.Println("Concurrent Login detected! Concurrent Login will not be allowed")
					loginmessage = "You have already logged in on another machine.\n Logout from the other session first or report if you have not logged in elsewhere!"
					// redirect to main index
				} else {
					// Check if admin login

					if lInput.Username == a.AdminLogin.AdminName {
						err := bcrypt.CompareHashAndPassword([]byte(a.AdminLogin.AdminPassword), []byte(lInput.Password))

						if err != nil {
							loginmessage = "Username and/or password do not match"
						} else {

							// create session
							id := uuid.NewV4()

							myCookie := &http.Cookie{
								Name:     "VaccineAppt",
								Value:    id.String(),
								MaxAge:   sessionExpireTime,
								HttpOnly: true,
								Secure:   true,
							}
							http.SetCookie(w, myCookie)

							a.MapSessions[myCookie.Value] = Session{lInput.Username, time.Now()}

							http.Redirect(w, req, "/admin", http.StatusSeeOther)
							return
						}
					} else {
						// check if user exist with username
						usernameNode := a.BstUserName.Search(lInput.Username)
						if usernameNode == nil { // username is not in BST
							loginmessage = "Username and/or password do not match"
						} else {
							var err error
							currentPerson, _, err = a.PersonList.SearchUserName(lInput.Username)

							if err != nil {
								loginmessage = "Username and/or password do not match"
							} else {
								err = bcrypt.CompareHashAndPassword([]byte(currentPerson.Password), []byte(lInput.Password))

								if err != nil {
									loginmessage = "Username and/or password do not match"
								} else {
									// create session
									id := uuid.NewV4()

									myCookie := &http.Cookie{
										Name:     "VaccineAppt",
										Value:    id.String(),
										MaxAge:   sessionExpireTime,
										HttpOnly: true,
										Secure:   true,
									}
									http.SetCookie(w, myCookie)
									a.MapSessions[myCookie.Value] = Session{lInput.Username, time.Now()}

									http.Redirect(w, req, "/", http.StatusSeeOther)
									return
								}
							}
						}
					}
				}
			}
		}
	}
	tpl.ExecuteTemplate(w, "login.gohtml", loginmessage)
}

// This method is used to logout user/admin
// If no cookie is present, the page is redirected to login page
// If cookie is present, it is deleted, and corrspoding map session entry is also deleted
// User data and appointment data is also save to the CSV file
func (a *AppData) LogoutHandler(w http.ResponseWriter, req *http.Request) {

	myCookie, err := req.Cookie("VaccineAppt")

	if err != nil {
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		Warning.Println("Cookie not found, redirecting to login page..")
		Trace.Println("Cookie not found, redirecting to login page..")
		return
	} else {

		// Save data to CSV files
		if _, ok := a.MapSessions[myCookie.Value]; ok {
			Wg.Add(2)
			fh.WritePersonCSVFile(Wg, a.PersonList)
			fh.WriteApptCSVFile(Wg, a.ApptArray)
			Wg.Wait()
		}

		// delete the session
		delete(a.MapSessions, myCookie.Value)

		// remove the cookie
		myCookie = &http.Cookie{
			Name:     "VaccineAppt",
			Value:    "",
			MaxAge:   -1,
			HttpOnly: true,
			Secure:   true,
		}
		http.SetCookie(w, myCookie)

		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
}

// This method is used to view the available appointments by user
// If the person qualifies for vaccination, then first/second/both appointments are displayed
func (a *AppData) ViewapptHandler(w http.ResponseWriter, req *http.Request) {
	if !a.activeSession(w, req) {
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}

	currentPerson := a.getPerson(w, req)
	message := []string{}
	if currentPerson.VaccinationQualify {
		msg, err := currentPerson.PrintAllAppt()
		if err != nil {
			Trace.Printf("Unable to get appointment data : %v\n", err)
			CheckLogChecksum()
			Error.Printf("Unable to get appointment data : %v\n", err)
			WriteChecksum()

			return
		}
		message = msg
	} else {
		message = printNotQualifiedMessage()
	}
	tpl.ExecuteTemplate(w, "viewappt.gohtml", message)
}

// This method is used to make new appointment by user
// It is used when an appointment is not made for first/second/both vaccinations
// It displays the list of available appointments for first/second vaccination
// Entry is done with radio button, which is then updated in appointment array and person data
func (a *AppData) MakeapptHandler(w http.ResponseWriter, req *http.Request) {
	if !a.activeSession(w, req) {
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}

	currentPerson := a.getPerson(w, req)

	makeapptmessage := userapptstruct{}

	if currentPerson.VaccinationQualify {
		err, fn, ln, msg, aptI, posApts := currentPerson.MakeNewAppt(a.ApptArray)
		if err != nil {
			Trace.Printf("Unable to make new appointments %v\n", err)
			CheckLogChecksum()
			Error.Printf("Unable to make new appointments %v\n", err)
			WriteChecksum()

			return
		}
		makeapptmessage.Firstname = fn
		makeapptmessage.Lastname = ln
		makeapptmessage.Message = msg
		makeapptmessage.Apptinfo = aptI
		makeapptmessage.Possibleappts = posApts

	} else {
		makeapptmessage.Message = printNotQualifiedMessage()
	}

	appointmentDate := ""
	// process form submission
	if req.Method == http.MethodPost {
		// get form values
		appointmentDate = req.FormValue("Appttime")
		err, apptA := currentPerson.UpdateNewAppt(appointmentDate, makeapptmessage.Apptinfo, a.ApptArray)
		if err != nil {
			Trace.Printf("Unable to update the appointment %v\n", err)
			CheckLogChecksum()
			Error.Printf("Unable to update the appointment %v\n", err)
			WriteChecksum()

			return
		}
		a.ApptArray = apptA

		//update person
		a.updatePerson(currentPerson, w, req)

		http.Redirect(w, req, "/index", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "makeappt.gohtml", makeapptmessage)
}

// This method is used to delete an appointment by user
// It checks which appointment can be deleted, and asks user to confirm
// Upon deletion, the person data and appointment array are updated
func (a *AppData) DeleteapptHandler(w http.ResponseWriter, req *http.Request) {
	if !a.activeSession(w, req) {
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}

	currentPerson := a.getPerson(w, req)

	deleteapptmessage := userapptstruct{}

	if currentPerson.VaccinationQualify {
		err, fn, ln, msg, aptI := currentPerson.DeleteAppt()
		if err != nil {
			Trace.Printf("Unable to delete the appointment %v\n", err)
			CheckLogChecksum()
			Error.Printf("Unable to delete the appointment %v\n", err)
			WriteChecksum()

			return
		}
		deleteapptmessage.Firstname = fn
		deleteapptmessage.Lastname = ln
		deleteapptmessage.Message = msg
		deleteapptmessage.Apptinfo = aptI
	} else {
		deleteapptmessage.Message = printNotQualifiedMessage()
	}
	// process form submission
	if req.Method == http.MethodPost {
		// get form values
		_ = req.FormValue("submit")
		apptA, err := currentPerson.UpdatedeleteAppt(deleteapptmessage.Apptinfo, a.ApptArray)
		if err != nil {
			Trace.Printf("Unable to update the deleted appointment %v\n", err)
			CheckLogChecksum()
			Error.Printf("Unable to update the deleted appointment %v\n", err)
			WriteChecksum()

			return
		}
		a.ApptArray = apptA

		//update person
		a.updatePerson(currentPerson, w, req)

		http.Redirect(w, req, "/index", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "deleteappt.gohtml", deleteapptmessage)
}

// This method is used to display profile information of the user
func (a *AppData) ProfileHandler(w http.ResponseWriter, req *http.Request) {
	if !a.activeSession(w, req) {
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}
	currentPerson := a.getPerson(w, req)
	tpl.ExecuteTemplate(w, "profile.gohtml", currentPerson)

}

// This method is used to search person data in the person linked list based on the username in the cookie
// This method is only used for user, and not for admin
func (a *AppData) getPerson(w http.ResponseWriter, req *http.Request) user.Person {
	// get current session cookie

	myCookie, _ := req.Cookie("VaccineAppt")

	// update expiry time every time there is an activity

	// if the user exists already, get user
	var currentPerson user.Person
	if s, ok := a.MapSessions[myCookie.Value]; ok {
		if s.Username != a.AdminLogin.AdminName {
			currentPerson, _, _ = a.PersonList.SearchUserName(s.Username)
		}
	}
	return currentPerson

}

// This method is used to overwrite person data in the person linked list based on the username in the cookie
// It takes in value of type person, which is the updated entries of the user data.
// This method is called when appointments are added or deleted by the user
func (a *AppData) updatePerson(p user.Person, w http.ResponseWriter, req *http.Request) {
	// get current session cookie
	myCookie, err := req.Cookie("VaccineAppt")

	if err != nil { //cannot find cookie, return to login
		Warning.Printf("No Cookie found : %v\n", err)
		Trace.Printf("No Cookie found : %v\n", err)
		return
	}

	// if the user exists already, get user
	if s, ok := a.MapSessions[myCookie.Value]; ok {
		if s.Username != a.AdminLogin.AdminName {
			currentPerson, index, _ := a.PersonList.SearchUserName(s.Username)
			if p.Username == currentPerson.Username {
				if a.PersonList.WritePersonData(p, index) != nil {
					Trace.Printf("Error Writing Person Data %v\n", err)
					CheckLogChecksum()
					Error.Printf("Error Writing Person Data %v\n", err)
					WriteChecksum()

				}
			}
		}
	}
}

// This method is used to display the admin template
func (a *AppData) AdminHandler(w http.ResponseWriter, req *http.Request) {

	if !a.activeSession(w, req) {
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "admin.gohtml", nil)
}

// This method is used to list all the current users on the admin template
// It parses through the person linked list and displays the user names
func (a *AppData) ListallusersHandler(w http.ResponseWriter, req *http.Request) {
	if !a.activeSession(w, req) {
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}
	adminToTemplate := adminStruct{}
	adminToTemplate.Message, adminToTemplate.Users = a.PersonList.PrintAllUsers()
	adminToTemplate.Deleteuser = "no"
	tpl.ExecuteTemplate(w, "admin.gohtml", adminToTemplate)

}

// This method is used to list all the current users on the admin template with option to delete
// It parses through the person linked list and displays the user names, along with radio buttons to select which user to delete.
// If admin picks one user to delete, the corresponding person data is deleted from linked list and BST entries
func (a *AppData) DeleteuserHandler(w http.ResponseWriter, req *http.Request) {
	if !a.activeSession(w, req) {
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}
	adminToTemplate := adminStruct{}
	msg, usr := a.PersonList.PrintAllUsers()
	adminToTemplate.Message = msg
	adminToTemplate.Users = usr
	adminToTemplate.Deleteuser = "yes"

	// process form submission
	if req.Method == http.MethodPost {
		// get form values
		userToDeleteIndex := req.FormValue("user")
		// add one to userToDeleteIndex, as index in form is a range starting from 0
		personRemoved, _ := a.PersonList.Remove(convertToInt(userToDeleteIndex) + 1)
		Wg.Add(2)
		go a.BstUserName.Delete(personRemoved.Username, Wg)
		go a.BstID.Delete(personRemoved.Identification, Wg)
		Wg.Wait()

		http.Redirect(w, req, "/admin", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "admin.gohtml", adminToTemplate)

}

// This method is used to view all available appointments by date
// The available times are shown on the same row for each date for compact viewing
func (a *AppData) ViewapptsbydateHandler(w http.ResponseWriter, req *http.Request) {

	if !a.activeSession(w, req) {
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}

	adminToTemplate := listApptByDate(a.ApptArray)
	adminToTemplate.ApptAdd = "no"

	tpl.ExecuteTemplate(w, "admin.gohtml", adminToTemplate)
}

// This method is used to add a date to available appointments, with bulk addition of time
// Appointment time from 9:00 am to 5:45 pm, interval of 15 minutes, are added to the date
// The addition is done in sorted manner, and duplicate entries are deleted
// The Appointment array is modified
func (a *AppData) AddapptsfordateHandler(w http.ResponseWriter, req *http.Request) {

	if !a.activeSession(w, req) {
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}
	adminToTemplate := adminStruct{}
	adminToTemplate.Message = append(adminToTemplate.Message, fmt.Sprintf("Choose a date to add appointments from 9:00am upto 5:45pm"))
	adminToTemplate.ApptAdd = "yes"

	// process form submission
	if req.Method == http.MethodPost {
		// get form values
		addtoThisDate := req.FormValue("appt")
		addThisDate := convertToTime(addtoThisDate)
		a.ApptArray = addApptArray(a.ApptArray, addThisDate)
		http.Redirect(w, req, "/admin", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "admin.gohtml", adminToTemplate)
}

// This method is used to validate the login input data
// It uses the https://github.com/go-playground/validator package
// Validation criteria is in loginInput struct, which is input to this method
func loginValidation(ri loginInput) error {
	// use a single instance of Validate, it caches struct info
	var validate *validator.Validate
	validate = validator.New()

	// returns nil or ValidationErrors ( []FieldError )
	err := validate.Struct(ri)
	Warning.Printf("Error in login Validation : %v\n", err)
	Trace.Printf("Error in login Validation : %v\n", err)
	return err
}

// This method is used to sanitize the login input data
// It checks that username begins with an alphabet and has no special characters
// It also checks that password does not contain white space or null characters
func loginSanitization(ri loginInput) bool {
	// Sanitize username
	un := regexp.QuoteMeta(ri.Username)
	if un != ri.Username { // Check if contains any special characters
		return true
	} else {
		// First character must be an alphabet, remaining characters must be alphabet,digits or underscore
		unMatch := regexp.MustCompile(`^[a-zA-Z]+\w*$`)
		if !unMatch.MatchString(ri.Username) {
			return true
		} else {
			// Username is okay..check password
			// Sanitize password , it can contain special characters
			// Check if password contains white space or null characters
			newlineChar := fmt.Sprintf("\\s")
			if strings.Contains(ri.Password, newlineChar) || strings.Contains(ri.Password, "%00") {
				return true
			} else {
				// no errors, return false
				return false
			}
		}
	}
}

// This method is used to validate the registration input data
// It uses the https://github.com/go-playground/validator package
// Validation criteria is in registerInput struct, which is input to this method
func registerValidation(ri registerInput) error {
	// use a single instance of Validate, it caches struct info
	var validate *validator.Validate
	validate = validator.New()

	// returns nil or ValidationErrors ( []FieldError )
	err := validate.Struct(ri)
	Warning.Printf("Error in registration Validation : %v\n", err)
	Trace.Printf("Error in registration Validation : %v\n", err)
	return err
}

// This method is used to sanitize the registration input data
// It checks that identification begins with an alphabet, is alphanumeric and has no special characters.
// It converts identification to uppercase.
// It checks that username begins with an alphabet and has no special characters
// It converts username to lowercase.
// It checks that password does not contain white space or null characters.
// It checks that firstname and lastname start with an alphabet.
// It converts firstname and lastname to titlecase.
// It checks that DOB is not 140 years before now, or after now.
// It checks that phone number is digits only.
// It checks that address does not contain javascript <script> tag.
// It checks that email begins with an alphabet
// It converts email to lowercase.
func registerSanitization(ri registerInput) (registerInput, registerError, bool) {

	regError := registerError{}
	regOutput := registerInput{}

	checkError := false

	// Sanitize Identification number
	id := regexp.QuoteMeta(ri.Identification)
	if id != ri.Identification { // Check if contains any special characters
		checkError = true
		regError.Identification = "No special characters allowed"
	} else {
		// First character must be an alphabet
		idMatch := regexp.MustCompile(`^[a-zA-Z]+[a-zA-Z0-9]*$`)
		if !idMatch.MatchString(ri.Identification) {
			regError.Identification = "Incorrect Identification number"
			checkError = true
		} else {
			// If no errors, pass the identification number to the program
			regOutput.Identification = strings.ToUpper(ri.Identification)
		}
	}

	// Sanitize username
	un := regexp.QuoteMeta(ri.Username)
	if un != ri.Username { // Check if contains any special characters
		checkError = true
		regError.Username = "No special characters allowed"
	} else {

		// First character must be an alphabet, remaining characters must be alphabet,digits or underscore
		unMatch := regexp.MustCompile(`^[a-zA-Z]+\w*$`)
		if !unMatch.MatchString(ri.Username) {
			regError.Username = "Incorrect Username"
			checkError = true
		} else {
			// If no errors, pass the username to the program
			regOutput.Username = strings.ToLower(ri.Username)
		}
	}

	// Sanitize password , it can contain special characters
	// Check if password contains white space or null characters
	newlineChar := fmt.Sprintf("\\s")
	if strings.Contains(ri.Password, newlineChar) || strings.Contains(ri.Password, "%00") {
		checkError = true
		regError.Password = "No spaces/null characters allowed"
	} else {
		// If no errors, pass the password to the program
		regOutput.Password = ri.Password

	}

	// Sanitize Firstname , it can contain special characters and spaces - world -friendly
	// First character must be an alphabet
	fnMatch := regexp.MustCompile(`^[a-zA-Z]+.*$`)
	if !fnMatch.MatchString(ri.Firstname) {
		regError.Firstname = "Incorrect Firstname"
		checkError = true
	} else {
		// If no errors, pass the Firstname number to the program
		regOutput.Firstname = strings.ToTitle(ri.Firstname)
	}

	// Sanitize Lastname , it can contain special characters and spaces - world -friendly
	// First character must be an alphabet
	lnMatch := regexp.MustCompile(`^[a-zA-Z]+.*$`)
	if !lnMatch.MatchString(ri.Lastname) {
		regError.Lastname = "Incorrect Lastname"
		checkError = true
	} else {
		// If no errors, pass the Lastname number to the program
		regOutput.Lastname = strings.ToTitle(ri.Lastname)
	}

	// Sanitize DOB , it is obtained with date input type from the template
	treg, err := time.Parse(TimeFormatISO, ri.Dob)
	if err != nil {
		regError.Dob = "Incorrect Date"
		checkError = true
	} else {
		// Check that DOB is within range
		// earliest date is now - 140 years
		tNow := time.Now()
		tthen := tNow.AddDate(-140, 0, 0)
		// DOB cannot be before 140 years from now, or after now
		if treg.Before(tthen) || treg.After(tNow) {
			regError.Dob = "Incorrect Date"
			checkError = true
		} else {
			// If no errors, pass the DOB to the program
			regOutput.Dob = ri.Dob
		}
	}

	// Sanitize phone :
	phMatch := regexp.MustCompile(`^[0-9]+$`)
	if !phMatch.MatchString(ri.Phone) {
		regError.Phone = "Incorrect Phone"
		checkError = true
	} else {
		// If no errors, pass the Phone number to the program
		regOutput.Phone = ri.Phone
	}

	// Sanitize address :
	// Check if there is any javascript in address
	adrMatch := regexp.MustCompile(`\<script\>?`)
	if adrMatch.FindStringIndex(ri.Address) != nil { // Found <script> in address
		regError.Address = "Incorrect Address"
		checkError = true
	} else {
		// no javascript inside address
		regOutput.Address = ri.Address
	}

	// Sanitize email :
	// First character must be an alphabet
	emailMatch := regexp.MustCompile(`^[a-zA-Z]+.*$`)
	if !emailMatch.MatchString(ri.Email) {
		regError.Email = "Incorrect Email"
		checkError = true
	} else {
		// If no errors, pass the Email number to the program
		regOutput.Email = strings.ToLower(ri.Email)
	}

	return regOutput, regError, checkError
}

// This function returns a slice of string message
// The message is used when user is not qualified to receive vaccination
func printNotQualifiedMessage() []string {
	message := []string{}
	message = append(message, fmt.Sprintf("You do not qualiy for Vaccination at this point. "))
	message = append(message, fmt.Sprintf("Vaccinations are rolled out only for person(s) aged %d and above.\n", ageQualification))
	message = append(message, fmt.Sprintf("Please login after you meet the age requirements,"))
	message = append(message, fmt.Sprintf("or when MOH contacts you on your phone with updated vaccination information."))
	message = append(message, fmt.Sprintf("Please contact our toll-free numbers for any assistance."))
	return message
}

// This function is used to calculate the age of a person using date of birth in string format
// It uses the convertToTime() function to convert dob into time format
// The Age is then calculated with age.Age function using "github.com/bearbin/go-age" package
// It returns the age as an integer
func calculateAge(dob string) int {
	dobTime := convertToTime(dob)
	age := age.Age(dobTime)
	return age
}

// This function converts a string to time.Time format using time.Parse function
func convertToTime(stringInput string) time.Time {
	t, err := time.Parse(TimeFormatISO, stringInput)

	if err != nil {
		Trace.Printf("Error in Parsing Time, cannot convert string to time.Time : %v\n", err)
		CheckLogChecksum()
		Error.Printf("Error in Parsing Time, cannot convert string to time.Time : %v\n", err)
		WriteChecksum()

	}
	return t
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

// This function converts a string into an integer using strconv.ParseInt function
func convertToInt(stringInput string) int {
	number, _ := strconv.ParseInt(stringInput, 10, 0)
	return int(number)
}

// This function creates adminStruct used for admin template for viewing appointments in concise manner
// It takes in the appointment array slice, and lists all appointment dates
// It lists the time for each date on the same row in kitchen time format ("3:04PM")
// It returns the adminStruct containing message to be displayed
func listApptByDate(apptArray []time.Time) adminStruct {
	adminToTemplate := adminStruct{}

	adminToTemplate.Message = append(adminToTemplate.Message, fmt.Sprintf("Available Appointments by date are: "))

	dateFound := []string{}
	timeNow := time.Now()
	i := 0
	for _, apptItem := range apptArray {
		if timeNow.Before(apptItem) {
			currentDate := apptItem.Format(TimeFormatISO)
			kitchenTime := apptItem.Format(Kitchen)
			if !findDateInArray(dateFound, currentDate) {
				newTime := currentDate + " => " + kitchenTime
				dateFound = append(dateFound, newTime)
				i++
			} else {
				dateFound[i-1] = dateFound[i-1] + " | " + kitchenTime
			}
		}
	}
	for _, item := range dateFound {
		adminToTemplate.Users = append(adminToTemplate.Users, fmt.Sprintf("%s", item))
	}
	if len(dateFound) == 0 {
		adminToTemplate.Message = append(adminToTemplate.Message, fmt.Sprintf("\nNo available appointments found!\n "))
	}
	return adminToTemplate
}

// This function checks if a value of type string is present in a slice of string
// It returns true if arrayValue is present in arrayName, else returns false
// It is used by listApptByDate to sort array using => as a delimiter between each entry.
func findDateInArray(arrayName []string, arrayValue string) bool {
	for _, item := range arrayName {
		var splititem = strings.Split(item, "=>")
		itemInArray := strings.TrimSpace(splititem[0])
		if itemInArray == arrayValue {
			return true
		}
	}
	return false
}

// This function adds appointments in bulk to a particular date
// It takes in an appointment array as a slice of time.Time, and a date for which appointments are to be added
// It adds appointments from 9:00 am upto 5:45 pm for that date, with an interval of 15 minutes.
// It uses the insertApptArray() function to add each appointment in a sorted manner
// The updated appointment array slice is returned.
func addApptArray(apptArray []time.Time, addDate time.Time) []time.Time {

	for i := 9; i < 18; i++ {
		for j := 0; j <= 45; j = j + 15 {
			addThis := addDate
			addThis = addThis.Add(time.Hour * time.Duration(i))
			addThis = addThis.Add(time.Minute * time.Duration(j))
			apptArray = insertApptArray(apptArray, addThis)
		}
	}
	return apptArray
}

// This function takes in appointment array as a slice of time.Time, and an appointment to be added in time.Time
// It inserts the appointment to be added into the slice, and returns the updated appointment array slice
// Insertion is done such that the array is always sorted, and any duplicate entries are removed.
func insertApptArray(apptArray []time.Time, addDate time.Time) []time.Time {

	for i := 0; i < len(apptArray); i++ {
		arrayTime := apptArray[i]
		if i+1 < len(apptArray) { // i+1 is not more than array size
			arrayTimeNext := apptArray[i+1]
			if arrayTime.Before(addDate) && arrayTimeNext.After(addDate) {
				tempArray := append(apptArray[:i+1], addDate)
				tempArray = append(tempArray, apptArray[i+1:]...)
				return tempArray
			} else if arrayTime.Equal(addDate) {
				// Duplicate item, no need to add
				return apptArray
			}
		} else { // last item reached
			return append(apptArray, addDate)
		}
	}
	return apptArray
}

// This function returns true when the hash value calculated by CalcChecksum() is same as that stored in checksum file
// If hash value is same, then the log file integrity has not been compromised.
// If hash value is different, then the log file has been tampered with, and a false is returned
func CheckLogChecksum() {
	// Get previously stored checksum from checksum file.
	path := filePath + logPath + "srvchecksum"
	logChecksum, err := ioutil.ReadFile(path)
	if err != nil {
		Trace.Printf("Unable to read srvchecksum file: %v\n", err)
		Warning.Printf("Unable to read srvchecksum file: %v\n", err)
	}
	str := string(logChecksum) // convert logChecksum in bytes to a 'string'

	// Compute our current log's hash
	hash := CalcChecksum()

	// Compare our calculated hash with our stored hash
	if str == hash {
		Info.Printf("Checksum for fhchecksum file matches\n")
		// Ok the checksums match.
	} else {
		// The file integrity has been compromised...
		Warning.Printf("Checksum for srvchecksum file does not match\n")
	}
}

// This function returns the hash value calculated using Blake2b hash package.
// The file "log" is opened and io.Copy will copy the file content to the hasher in stream fashion
// Hash is calculated over hasher and encoded to a string which is returned.
func CalcChecksum() string {
	// Compute our current log's Blake2b hash
	hasher, _ := blake2b.New256(nil)

	path := filePath + logPath + "srvlog"
	f, err := os.Open(path)
	if err != nil {
		Trace.Printf("Unable to open srvlog file: %v\n", err)
		Warning.Printf("Unable to open srvlog file: %v\n", err)
	}

	defer f.Close()

	if _, err := io.Copy(hasher, f); err != nil {
		Trace.Printf("Unable to io.copy srvlog file: %v\n", err)
		Warning.Printf("Unable to io.copy srvlog file: %v\n", err)
	}

	hash := hasher.Sum(nil)
	encodedHex := hex.EncodeToString(hash[:])
	return encodedHex
}

// This function writes the checksum value to the checksum file.
// It calculates the hash of the log file using CalcChecksum()
// It overwrites/creates the checksum file, writes the hashvalue, and closes the file
func WriteChecksum() {
	hashValue := CalcChecksum()

	path := filePath + logPath + "srvchecksum"
	filecs, err := os.OpenFile(path, os.O_CREATE, 0666)
	if err != nil {
		Trace.Printf("Failed to open filehandler checksum log file: %v\n", err)
		Warning.Printf("Failed to open filehandler checksum log file: %v\n", err)
	}
	filecs.Write([]byte(hashValue))
	filecs.Close()
}
