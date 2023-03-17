package logModule

import (
	"log"
	"os"
	"time"
)

var Logger *log.Logger
var Logfile *os.File

func InitLogs() (*os.File, *log.Logger) {
	date := time.Now().Format("02-01-2006")
	file := date + "_logfile.log"
	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("Can't init logger: ", err)
	}
	logcreated := log.New(f, "", log.LstdFlags)
	return f, logcreated
}
