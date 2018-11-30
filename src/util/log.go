package util

import (
	"log"
	"os"
	"path/filepath"
)

var (
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func LoggerInit() {
	newpath := filepath.Join(".", "log")
	os.MkdirAll(newpath, os.ModePerm)

	infoHandle, _ := os.OpenFile("log/info.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	warningHandle, _ := os.OpenFile("log/warn.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	errorHandle, _ := os.OpenFile("log/error.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)

}
func LogInfo() *log.Logger {
	return Info
}
func LogWarning() *log.Logger {
	return Warning
}
func LogError() *log.Logger {
	return Error
}
