//Package main is the main function for Vaccine Appointment System
// It contains constant declarations and main function calls
// It imports the server and filehandler packages
// It perfoms basic checks on CSV Read/Write and exits if there is failure writing temporary CSV files
// Developed by Pallavi Limaye - 28/03/2021
package main

import (
	"fmt"
	"goInAction-2-p/filehandler"
	"goInAction-2-p/server"
	"net/http"
	"sync"
	"time"

	"encoding/csv"
	"io"
	"io/ioutil"
	"log"
	"os"
	//Internal packages of vaccineappt
	//fh "vaccineappt/source/filehandler"
	//server "vaccineappt/source/server"
)

const (
	// File Paths used
	filePath = "./"
	csvPath  = "./csv/"
	tmpPath  = "./tmp/"
	htmlPath = "./html/"
	certPath = "./cert/"
	logPath  = "./log/"

	// CSV File names
	apptFileName       = "freeappointments.csv"
	personDataFileName = "personData.csv"
	adminLoginFile     = "adminLogin.csv"
)

var (
	// Global variables
	wg sync.WaitGroup
)

// The initialization function prepares for templates parsing, and opening of CSV files and Log files and calculating checksum for log file.
// It also creates the loggers for Trace, Info, Warning and Error messages
// Any errors with writing temporary files will result in fatal error and the application will exit.
func init() {
	server.Wg = &wg

	// Create log file for writing Errors
	sfile, err := os.OpenFile(logPath+"srvlog", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open server error log file:", err)
	}

	// Create checksum file, and write hash of log file into it
	server.WriteChecksum()

	// Create loggers for logging Trace, Info, Warning and Error messages
	server.Trace = log.New(ioutil.Discard,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	server.Info = log.New(os.Stdout,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	server.Warning = log.New(os.Stdout,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	server.Error = log.New(io.MultiWriter(sfile, os.Stderr),
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	// Create log file for writing Errors
	fhfile, err := os.OpenFile(logPath+"fhlog", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open filehandler error log file:", err)
	}

	// Create checksum file, and write hash of log file into it
	filehandler.WriteChecksum()

	// Create loggers for logging Trace, Info, Warning and Error messages
	filehandler.Trace = log.New(ioutil.Discard,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	filehandler.Info = log.New(os.Stdout,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	filehandler.Warning = log.New(os.Stdout,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	filehandler.Error = log.New(io.MultiWriter(fhfile, os.Stderr),
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	// Perform preliminary checks on CSV files, to check if they can be read and written into
	// Check if CSV Appointment file can be opened
	csvapptfile, err := os.Open(filePath + csvPath + apptFileName)
	if err != nil {
		filehandler.Warning.Printf("Unable to open Appointment Available CSV file: %v\n", err)
	}

	// Check if CSV Appointment file can be read
	_, err = csv.NewReader(csvapptfile).ReadAll()
	if err != nil {
		filehandler.Warning.Printf("Unable to read Appointment Available CSV file: %v\n", err)
	}
	csvapptfile.Close()

	// Check if a temporary CSV Appointment file can be created and written into /tmp folder. Deleted after test
	csvtmp := filePath + tmpPath + "tmp_" + apptFileName
	csvappttmpfile, err := os.Create(csvtmp)
	if err != nil {
		log.Fatalf("Failed creating temporary Appointment Available CSV file: %v\n", err)
	}
	csvappttmpfile.Close()

	// Remove temporary CSV Appointment file from /tmp folder
	err = os.Remove(csvtmp)
	if err != nil {
		filehandler.Warning.Printf("Unable to remove temporary Appointment CSV file: %v\n", err)
	}

	// Check if Person Data CSV file can be opened
	csvPerFile, err := os.Open(filePath + csvPath + personDataFileName)
	if err != nil {
		filehandler.Warning.Printf("Unable to open Person Data CSV file: %v\n", err)
	}

	// Check if Person Data CSV file can be read
	_, err = csv.NewReader(csvPerFile).ReadAll()
	if err != nil {
		filehandler.Warning.Println("Unable to read Person Data CSV file")
	}
	csvPerFile.Close()

	// Check if a temporary CSV Person Data file can be created and written into /tmp folder. Deleted after test
	csvpertmp := filePath + tmpPath + "tmp_" + personDataFileName
	csvpertmpfile, err := os.Create(csvpertmp)
	if err != nil {
		log.Fatalf("Failed creating temporary Person Data CSV file: %v\n", err)
	}
	csvpertmpfile.Close()

	// Remove temporary Person Data CSV file from /tmp folder
	err = os.Remove(csvpertmp)
	if err != nil {
		filehandler.Warning.Printf("Unable to remove temporary Person CSV file: %v\n", err)
	}

}

// The main function for the Vaccine Appointment System.
// It initializes all data structures, reading of CSV files and handling panic
// It contains all handle functions for all templates.
// It calls the ListenAndServeTLS function using SSL certificate and private key generated with OpenSSL
func main() {
	//Handle any panics generated by server and log into server error log
	defer func() {
		if err := recover(); err != nil {
			server.Error.Println(err)
		}
	}()

	// Initialize data structures, map sessions and appointment array
	appDS := server.AppData{}
	appDS.MapSessions = make(map[string]server.Session)
	appDS.MapSessionCleaned = time.Now()
	appDS.ApptArray = filehandler.ReadApptCSVFile(&wg)

	// Read all person data from CSV file
	appDS.PersonList, appDS.BstUserName, appDS.BstID = filehandler.ReadpersonCSVFile(&wg)

	// Read Admin Login and Password information
	appDS.AdminLogin.AdminName, appDS.AdminLogin.AdminPassword = filehandler.ReadAdminCSVFile()

	// Handle function for all server functions
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./Assets"))))

	http.HandleFunc("/", appDS.IndexHandler)
	http.HandleFunc("/register", appDS.RegistrationHandler)
	http.HandleFunc("/login", appDS.LoginHandler)
	http.HandleFunc("/logout", appDS.LogoutHandler)
	http.HandleFunc("/viewappt", appDS.ViewapptHandler)
	http.HandleFunc("/makeappt", appDS.MakeapptHandler)
	http.HandleFunc("/deleteappt", appDS.DeleteapptHandler)
	http.HandleFunc("/profile", appDS.ProfileHandler)
	http.HandleFunc("/admin", appDS.AdminHandler)
	http.HandleFunc("/listallusers", appDS.ListallusersHandler)
	http.HandleFunc("/deleteuser", appDS.DeleteuserHandler)
	http.HandleFunc("/viewapptsbydate", appDS.ViewapptsbydateHandler)
	http.HandleFunc("/addapptsfordate", appDS.AddapptsfordateHandler)

	// HTTPs server : Listen and serve using TLS(SSL certificate and Private key)
	fmt.Println("Listening....")
	log.Fatal(http.ListenAndServeTLS(":5221", certPath+"cert.pem", certPath+"key.pem", nil))
}
