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

/*OpenFile is the generalized open call; It opens the logfile with specified flag
RDWR for write read and append will add.If successful,methods on the returned File
can be used for I/O. If there is an error,it will be of type *PathError.*/
func init() {
	var filepath = "C:/Projects/Go/src/gosecurity/logger/logfile.log"
	openLogfile, err := os.OpenFile(filepath,
		os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Failed to open log file:", err)
		os.Exit(1)
	} /*Prefix(trace, info, warning, error) log will provide the information, it can be general
	like successfil login etc.New creates a Logger. The out variable sets the destination to which
	log data will be written. The prefix appears at the beginning of each generated log line,
	The flag argument defines the logging properties.*/
	TraceLog = log.New(openLogfile, "Trace Logger:\t",
		log.Ldate|log.Ltime|log.Lshortfile)

	InfoLog = log.New(openLogfile, "Info Logger:\t",
		log.Ldate|log.Ltime|log.Lshortfile)

	WarningLog = log.New(openLogfile, "Warning Logger:\t",
		log.Ldate|log.Ltime|log.Lshortfile)

	ErrorLog = log.New(openLogfile, "Error Logger:\t",
		log.Ldate|log.Ltime|log.Lshortfile)

}
