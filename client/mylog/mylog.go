package mylog

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

const (
	LogLevelTrace = iota
	LogLevelInfo
	LogLevelWarning
	LogLevelError
)

var (
	LogLevel int = 3
)

//Init Log
func Init(ll int) {
	LogLevel = ll

}

//Print log message
func Print(ll int, msg ...string) {
	pc, _, line, _ := runtime.Caller(1)
	func_and_line := runtime.FuncForPC(pc).Name() + ":" + strconv.Itoa(line)
	tmestamp := time.Now().Format(time.StampMilli)
	if ll >= LogLevel {
		switch ll {
		case LogLevelError:
			//	Error.Println(out)
			fmt.Println(tmestamp, "[Error]", func_and_line, msg)
		case LogLevelWarning:
			//	Warning.Println(out)
			fmt.Println(tmestamp, "[Warning]", func_and_line, msg)
		case LogLevelInfo:
			//	Info.Println(out)
			fmt.Println(tmestamp, "[Info]", func_and_line, msg)
		case LogLevelTrace:
			//	Trace.Println(out)
			fmt.Println(tmestamp, "[Trace]", func_and_line, msg)
		}

	}

}
