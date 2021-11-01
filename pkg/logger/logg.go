package logger

import (
	"log"
	"os"
)

var (
	errLogger  = errLog()
	infoLogger = infoLog()
)

func infoLog() *log.Logger {
	return log.New(os.Stdout, "INFO ", log.Ldate|log.Ltime|log.Llongfile)
}

func errLog() *log.Logger {
	return log.New(os.Stdout, "ERROR ", log.Ldate|log.Ltime|log.Llongfile)
}

func Info(info string) {
	warn := infoLogger.Output(2, info)
	if warn != nil {
		Error(warn)
	}
}

func Error(err error) {
	warn := errLogger.Output(2, err.Error())
	if warn != nil {
		Error(warn)
	}
}
