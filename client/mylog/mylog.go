package mylog

import (
	"io"
	"log"
)

const (
	LogLevelTrace = iota
	LogLevelInfo
	LogLevelWarning
	LogLevelError
)

var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger

	LogLevel int = 3
)

//Init Log
func Init(Handle io.Writer, ll int) {
	LogLevel = ll

	Trace = log.New(Handle,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(Handle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(Handle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(Handle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

//Print log message
func Print(ll int, msg ...string) {
	if ll <= LogLevel {
		switch LogLevel {
		case LogLevelError:
			Error.Println(msg)
		case LogLevelWarning:
			Warning.Println(msg)
		case LogLevelInfo:
			Info.Println(msg)
		case LogLevelTrace:
			Trace.Println(msg)
		}

	}

}
