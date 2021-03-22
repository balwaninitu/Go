package logger

import (
	"fmt"
	"log"
	"os"
)

var (
	TraceLog   *log.Logger
	InfoLog    *log.Logger
	WarningLog *log.Logger
	ErrorLog   *log.Logger
)

func init() {
	var filepath = "logfile.log"
	openLogfile, err := os.OpenFile(filepath,
		os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Failed to open log file:", err)
		os.Exit(1)
	}
	TraceLog = log.New(openLogfile, "Trace Logger:\t",
		log.Ldate|log.Ltime|log.Lshortfile)

	InfoLog = log.New(openLogfile, "Info Logger:\t",
		log.Ldate|log.Ltime|log.Lshortfile)

	WarningLog = log.New(openLogfile, "Warning Logger:\t",
		log.Ldate|log.Ltime|log.Lshortfile)

	ErrorLog = log.New(openLogfile, "Error Logger:\t",
		log.Ldate|log.Ltime|log.Lshortfile)

}
