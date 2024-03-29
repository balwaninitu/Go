// Package filehandler contains read and write functions for CSV files for admin data, user data and appointment data.
package filehandler

import (
	"encoding/csv"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	bt "vaccineappt/source/btree"
	user "vaccineappt/source/person"
	userll "vaccineappt/source/personll"

	"golang.org/x/crypto/blake2b"
)

const (
	//Time Format constant
	layoutDateTime = "2006-01-02 15:04:05"
)

var (
	mu sync.Mutex

	//Loggers used to log Trace,Info, Warning and Error in file handler package
	Trace   *log.Logger // Just about anything
	Info    *log.Logger // Important information
	Warning *log.Logger // Be concerned
	Error   *log.Logger // Critical problem

	// Paths
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

// This function writes the appointment available CSV file with updated appointment dates
// It takes in the appointment array slice of time.Time
// It issues a wg.Done for concurrent processing
// It uses a mutex lock during the write process
// The updated appointment date/time is written in apptFileName csv file
func WriteApptCSVFile(wg *sync.WaitGroup, apptArray []time.Time) {
	//Handle any panics generated by this function and log into filehandler error log
	defer func() {
		if err := recover(); err != nil {
			Trace.Println(err)
			CheckLogChecksum()
			Error.Println(err)
			WriteChecksum()

		}
	}()

	defer wg.Done()
	mu.Lock()
	{
		csvfile, err := os.Create(filePath + csvPath + apptFileName)
		defer csvfile.Close()

		if err != nil {
			panic(fmt.Sprintf("Unable to create Appointment Available CSV file for writing : %v\n", err))
		}

		csvwriter := csv.NewWriter(csvfile)

		for _, appt := range apptArray {

			timeNow := time.Now()
			if timeNow.Before(appt) {
				d := appt.Format(layoutDateTime)
				dateArray := strings.Split(d, " ")

				if err := csvwriter.Write(dateArray); err != nil {
					Trace.Printf("Unable to write into Appointment Available CSV file : %v\n", err)
					CheckLogChecksum()
					Error.Printf("Unable to write into Appointment Available CSV file : %v\n", err)
					WriteChecksum()

					panic(fmt.Sprintf("Unable to write into Appointment Available CSV file : %v\n", err))
				}
			}
		}
		csvwriter.Flush()
		if err := csvwriter.Error(); err != nil {
			Trace.Printf("Unable to flush Appointment Available CSV file : %v\n", err)
			CheckLogChecksum()
			Error.Printf("Unable to flush Appointment Available CSV file : %v\n", err)
			WriteChecksum()

			return
		}
	}
	mu.Unlock()
}

// This function writes the Person Data CSV file with updated user informtaion
// It takes in the person linked list as an interface
// It issues a wg.Done for concurrent processing
// It uses a mutex lock during the write process
// The updated person data is written in personDataFileName csv file
func WritePersonCSVFile(wg *sync.WaitGroup, p *userll.LinkedList) {
	//Handle any panics generated by this function and log into filehandler error log
	defer func() {
		if err := recover(); err != nil {
			Trace.Println(err)
			CheckLogChecksum()
			Error.Println(err)
			WriteChecksum()

		}
	}()

	defer wg.Done()

	mu.Lock()
	{
		csvfile, err := os.Create(filePath + csvPath + personDataFileName)
		defer csvfile.Close()

		if err != nil {
			Trace.Printf("Unable to create Person Data CSV file for writing : %v\n", err)
			CheckLogChecksum()
			Error.Printf("Unable to create Person Data CSV file for writing : %v\n", err)
			WriteChecksum()

			panic(fmt.Sprintf("Unable to create Person Data CSV file for writing : %v\n", err))
		}

		csvwriter := csv.NewWriter(csvfile)

		personllhead := p.Head
		personllsize := p.Size
		if personllsize != 0 && personllhead != nil {
			for i := 1; i < personllsize+1; i++ {
				oneperson, _ := p.Get(i)

				var line []string

				line = append(line, oneperson.Identification, oneperson.Username, oneperson.Password,
					oneperson.Firstname, oneperson.Lastname, oneperson.Dob, oneperson.Phone,
					oneperson.Address, oneperson.Email, strconv.FormatBool(oneperson.VaccinationQualify),
					strconv.FormatBool(oneperson.FirstVaccineDone), oneperson.FirstVaccineDate,
					oneperson.FirstVaccineTime, strconv.FormatBool(oneperson.SecondVaccineDone),
					oneperson.SecondVaccineDate, oneperson.SecondVaccineTime)

				if err := csvwriter.Write(line); err != nil {
					Trace.Printf("Unable to write into Person Data CSV file : %v\n", err)
					CheckLogChecksum()
					Error.Printf("Unable to write into Person Data CSV file : %v\n", err)
					WriteChecksum()

					panic(fmt.Sprintf("Unable to write into Person Data CSV file : %v\n", err))
				}
			}
			csvwriter.Flush()
			if err := csvwriter.Error(); err != nil {
				Trace.Printf("Unable to flush Person Data CSV file : %v\n", err)
				CheckLogChecksum()
				Error.Printf("Unable to flush Person Data CSV file : %v\n", err)
				WriteChecksum()

				return
			}

		}
	}
	mu.Unlock()
}

// This function reads the admin login information file containing admin username and password
// It returns the username and password in the adminLoginStruct
func ReadAdminCSVFile() (string, string) {

	an := ""
	apw := ""

	csvFile, err := os.Open(filePath + csvPath + adminLoginFile)
	defer csvFile.Close()

	if err != nil {
		Warning.Printf("Unable to open Admin Login CSV file : %v\n", err)
		return an, apw
	}

	csvr := csv.NewReader(csvFile)
	firstline, err := csvr.Read()
	if err != nil {
		Warning.Printf("Unable to read Admin Login CSV file : %v\n", err)
		return an, apw
	}
	an = firstline[0]
	apw = firstline[1]

	return an, apw
}

// This function reads the appointment available CSV file containing list of available appointments
// It returns the appointments in a time.Time slice
func ReadApptCSVFile(wg *sync.WaitGroup) []time.Time {
	var apptArray []time.Time

	csvFile, err := os.Open(filePath + csvPath + apptFileName)
	defer csvFile.Close()

	if err != nil {
		Warning.Printf("Unable to open Appointment Available CSV file : %v\n", err)
		return apptArray
	}

	csvLines, err := csv.NewReader(csvFile).ReadAll()

	if err != nil {
		Warning.Printf("Unable to read Appointment Available CSV file : %v\n", err)
		return apptArray
	}

	var newApptAvailable time.Time
	var newDate, newTime string
	for _, line := range csvLines {
		newDate = line[0]
		newTime = line[1]
		newApptAvailable, err = datePlusTime(newDate, newTime)
		if err != nil {
			Trace.Printf("Unable to Parse time while reading Appointment CSV file: %v\n", err)
			CheckLogChecksum()
			Error.Printf("Unable to Parse time while reading Appointment CSV file: %v\n", err)
			WriteChecksum()

		} else {
			if newApptAvailable.After(time.Now()) {
				apptArray = append(apptArray, newApptAvailable)
			}
		}
	}
	return apptArray
}

// This function reads the Person Data CSV file containing a list of users and their information
// It creates a person linked list to store all the data about the users
// It creates a BST for usernames for quick search operations
// It creates a BST for identification for quick search operations
// It uses goroutines for concurrent creation of these data structures
// It returns the pointers to the linked list, and the two BST structures
func ReadpersonCSVFile(wg *sync.WaitGroup) (*userll.LinkedList, *bt.BST, *bt.BST) {
	readList := &userll.LinkedList{Head: nil, Size: 0}
	bstUserName := &bt.BST{Root: nil}
	bstID := &bt.BST{Root: nil}

	csvFile, err := os.Open(filePath + csvPath + personDataFileName)
	defer csvFile.Close()

	if err != nil {
		Warning.Printf("Unable to open Person Data CSV file: %v\n", err)
		return readList, bstUserName, bstID
	}

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		Warning.Printf("Unable to read Person Data CSV file: %v\n", err)
		return readList, bstUserName, bstID
	}

	for _, line := range csvLines {
		i := 0
		onePerson := user.Person{
			Identification:     line[i],
			Username:           line[i+1],
			Password:           line[i+2],
			Firstname:          line[i+3],
			Lastname:           line[i+4],
			Dob:                line[i+5],
			Phone:              line[i+6],
			Address:            line[i+7],
			Email:              line[i+8],
			VaccinationQualify: convertToBool(line[i+9]),
			FirstVaccineDone:   convertToBool(line[i+10]),
			FirstVaccineDate:   line[i+11],
			FirstVaccineTime:   line[i+12],
			SecondVaccineDone:  convertToBool(line[i+13]),
			SecondVaccineDate:  line[i+14],
			SecondVaccineTime:  line[i+15],
		}

		// Starting goroutines for three functions.
		//Adding nodes to the linked list, and usernames and identification to BST

		wg.Add(3)
		go readList.AddNode(onePerson, wg)
		go bstUserName.Insert(onePerson.Username, wg)
		go bstID.Insert(onePerson.Identification, wg)
		wg.Wait()

	}
	return readList, bstUserName, bstID
}

// This function converts a string to a boolean using strconv.ParseBool function
func convertToBool(stringInput string) bool {
	done, _ := strconv.ParseBool(stringInput)
	return done
}

// This function concatenates date string and time string, and converts into time.Time
// It uses function dateTimeFormat() to convert string to time.Time
// It returns the converted time.Time
func datePlusTime(date, timeOfDay string) (time.Time, error) {
	if date == "" {
		var t time.Time
		return t, nil
	}
	dateWithTime := date + " " + timeOfDay
	return dateTimeFormat(dateWithTime)
}

// This function convert string to time.Time format
// It uses time.Parse function to convert string into time using layoutDateTime (2006-01-02 15:04:05) format
func dateTimeFormat(dateWithTime string) (time.Time, error) {
	dateTemp, err := time.Parse(layoutDateTime, dateWithTime)
	if err != nil {
		Trace.Printf("Unable to Parse time: %v\n", err)
		CheckLogChecksum()
		Error.Printf("Unable to Parse time: %v\n", err)
		WriteChecksum()

		return dateTemp, err
	}
	return dateTemp, nil
}

// This function returns true when the hash value calculated by CalcChecksum() is same as that stored in checksum file
// If hash value is same, then the log file integrity has not been compromised.
// If hash value is different, then the log file has been tampered with, and a false is returned
func CheckLogChecksum() {
	// Get previously stored checksum from checksum file.
	path := filePath + logPath + "fhchecksum"
	logChecksum, err := ioutil.ReadFile(path)
	if err != nil {
		Trace.Printf("Unable to read fhchecksum file: %v\n", err)
		Warning.Printf("Unable to read fhchecksum file: %v\n", err)
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
		Warning.Printf("Checksum for fhchecksum file does not match\n")
	}
}

// This function returns the hash value calculated using Blake2b hash package.
// The file "log" is opened and io.Copy will copy the file content to the hasher in stream fashion
// Hash is calculated over hasher and encoded to a string which is returned.
func CalcChecksum() string {
	// Compute our current log's Blake2b hash
	hasher, _ := blake2b.New256(nil)

	path := filePath + logPath + "fhlog"
	f, err := os.Open(path)
	if err != nil {
		Trace.Printf("Unable to open fhlog file: %v\n", err)
		Warning.Printf("Unable to open fhlog file: %v\n", err)
	}

	defer f.Close()

	if _, err := io.Copy(hasher, f); err != nil {
		Trace.Printf("Unable to io.copy fhlog file: %v\n", err)
		Warning.Printf("Unable to io.copy fhlog file: %v\n", err)
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

	path := filePath + logPath + "fhchecksum"
	filecs, err := os.OpenFile(path, os.O_CREATE, 0666)
	if err != nil {
		Trace.Printf("Failed to open filehandler checksum log file: %v\n", err)
		Warning.Printf("Failed to open filehandler checksum log file: %v\n", err)
	}
	filecs.Write([]byte(hashValue))
	filecs.Close()
}
