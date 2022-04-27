package utils

import (
	"log"
	"os"
)

var (
	InfoLog    *log.Logger
	WarningLog *log.Logger
	ErrorLog   *log.Logger
)

func init() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Panic(err.Error())
	}
	InfoLog = log.New(file, "[INFO]: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Llongfile)
	WarningLog = log.New(file, "[WARNING]: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Llongfile)
	ErrorLog = log.New(file, "[ERROR]: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Llongfile)
}
