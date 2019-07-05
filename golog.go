// very basic colorfoul log on the stderr console
// for more serious needs consider using https://github.com/uber-go/zap
package golog

import (
	"fmt"
	"github.com/lao-tseu-is-alive/goutils"
	"github.com/mgutz/ansi"
	"log"
	"os"
	"time"
)

var loggerInfo = log.New(os.Stderr, "INFO: ", log.Lshortfile)
var loggerTrace = log.New(os.Stderr, "TRACE: ", log.Lshortfile)
var loggerWarning = log.New(os.Stderr, "WARNING: ", log.Lshortfile)
var loggerError = log.New(os.Stderr, "ERROR: ", log.Lshortfile)

func Info(message string, v ...interface{}) {
	blue := ansi.ColorFunc("cyan+")
	filename, line, funcname := goutils.GetCaller(3)
	err := loggerInfo.Output(2,
		blue(fmt.Sprintf(
			"[%s], Function: %s:%d, Message: %s",
			goutils.GetTimeStamp(), funcname, line, fmt.Sprintf(message, v...))))
	if err != nil {
		log.Fatalln(fmt.Sprintf(
			"[%s], Function: %s, File: %s:%d",
			"ERROR trying to output Info(message) to stderr console !", funcname, filename, line))
	}
}

func Trace(message string, v ...interface{}) (string, time.Time) {
	start := time.Now()
	magenta := ansi.ColorFunc("magenta+b")
	filename, line, funcname := goutils.GetCaller(3)
	output := magenta(fmt.Sprintf("Function: %s, Message: %s", funcname, fmt.Sprintf(message, v...)))
	err := loggerTrace.Output(2, fmt.Sprintf("[%s], >ENTERING %s", goutils.GetTimeStamp(), output))
	if err != nil {
		log.Fatalln(fmt.Sprintf(
			"[%s], Function: %s, File: %s:%d",
			"ERROR trying to output Trace(message) to stderr console !", funcname, filename, line))
	}
	return output, start
}

// to be used with Trace like this at the begining of the body of a function
// USAGE : defer golog.Un(golog.Trace("your function message"))
func Un(message string, start time.Time) {
	elapsed := time.Since(start)
	err := loggerTrace.Output(2, fmt.Sprintf("[%s], <EXITING  %s (after %s)", goutils.GetTimeStamp(), message, elapsed))
	if err != nil {
		log.Fatalln("ERROR trying to output UnTrace(message) to stderr console !")
	}
}

func Warn(message string, v ...interface{}) {
	yellow := ansi.ColorFunc("yellow+b")
	filename, line, funcname := goutils.GetCaller(3)
	err := loggerTrace.Output(2,
		yellow(fmt.Sprintf(
			"[%s], Function: %s, Message: %s",
			goutils.GetTimeStamp(), funcname, fmt.Sprintf(message, v...))))
	if err != nil {
		log.Fatalln(fmt.Sprintf(
			"[%s], Function: %s, File: %s:%d",
			"ERROR trying to output Warn(message) to stderr console !", funcname, filename, line))
	}
}

func Err(message string, v ...interface{}) {
	red := ansi.ColorFunc("red+b:white+h")
	filename, line, funcname := goutils.GetCaller(3)
	err := loggerError.Output(2,
		red(fmt.Sprintf(
			"[%s], Function: %s, Message: %s",
			goutils.GetTimeStamp(), funcname, fmt.Sprintf(message, v...))))
	if err != nil {
		log.Fatalln(fmt.Sprintf(
			"[%s], Function: %s, File: %s:%d",
			"ERROR trying to output Err(message) to stderr console !", funcname, filename, line))
	}
}
