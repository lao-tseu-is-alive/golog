// very basic colorfoul log on the stderr console
// for more serious needs consider using https://github.com/uber-go/zap
package golog

import (
	"fmt"
	"github.com/mgutz/ansi"
	"log"
	"os"
	"time"
)

var loggerInfo = log.New(os.Stderr, "INFO: ", log.Lshortfile)
var loggerTrace = log.New(os.Stderr, "TRACE: ", log.Lshortfile)
var loggerWarning = log.New(os.Stderr, "WARNING: ", log.Lshortfile)
var loggerError = log.New(os.Stderr, "ERROR: ", log.Lshortfile)

func addTimeStamp(message string) string {
	t := time.Now()
	timeString := fmt.Sprintf("%d/%02d/%02d %02d:%02d:%02d.%03d",
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second(),
		t.Nanosecond()/100000)
	return fmt.Sprintf("%s %s", timeString, message)
}

func Info(message string, v ...interface{}) {
	blue := ansi.ColorFunc("cyan+")
	err := loggerInfo.Output(2, fmt.Sprintf(blue(addTimeStamp(message)), v...))
	if err != nil {
		log.Fatalln("ERROR trying to output Info(message) to stderr console !")
	}
}

func Trace(message string, v ...interface{}) {
	magenta := ansi.ColorFunc("magenta+b")
	err := loggerTrace.Output(2, fmt.Sprintf(magenta(addTimeStamp(message)), v...))
	if err != nil {
		log.Fatalln("ERROR trying to output Trace(message) to stderr console !")
	}
}

func Warn(message string, v ...interface{}) {
	yellow := ansi.ColorFunc("yellow+b")
	err := loggerWarning.Output(2, fmt.Sprintf(yellow(addTimeStamp(message)), v...))
	if err != nil {
		log.Fatalln("ERROR trying to output Warn(message) to stderr console !")
	}
}

func Err(message string, v ...interface{}) {
	red := ansi.ColorFunc("red+b:white+h")
	err := loggerError.Output(2, fmt.Sprintf(red(addTimeStamp(message)), v...))
	if err != nil {
		log.Fatalln("ERROR trying to output Err(message) to stderr console !")
	}
}
