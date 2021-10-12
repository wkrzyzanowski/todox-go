// Knowledge source about colors and logging with details:
// https://snippets.aktagon.com/snippets/795-logging-in-golang-including-line-numbers-
// https://twin.sh/articles/35/how-to-add-colors-to-your-console-terminal-output-in-go
package tools

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
)

var reset = "\033[0m"
var red = "\033[31m"
var green = "\033[32m"
var yellow = "\033[33m"
var blue = "\033[34m"
var purple = "\033[35m"
var cyan = "\033[36m"
var gray = "\033[37m"
var white = "\033[97m"

var debug = os.Getenv("DEBUG") != ""

var LOGGER Logger = Logger{}

func init() {
	checkOS()
}

type Logger struct {
}

func (logger *Logger) Info(msg string, vars ...interface{}) {
	var fullMessage = buildMessage(msg, "INFO", vars...)
	log.Printf("%v%v%v", blue, fullMessage, reset)
}

func (logger *Logger) Warning(msg string, vars ...interface{}) {
	var fullMessage = buildMessage(msg, "WARNING", vars...)
	log.Printf("%v%v%v", yellow, fullMessage, reset)
}

func (logger *Logger) Error(msg string, vars ...interface{}) {
	var fullMessage = buildMessage(msg, "ERROR", vars...)
	log.Printf("%v%v%v", red, fullMessage, reset)
}

func (logger *Logger) Fatal(msg string, vars ...interface{}) {
	var fullMessage = buildMessage(msg, "FATAL", vars...)
	log.Fatalf("%v%v%v", red, fullMessage, reset)

}

func buildMessage(msg string, logLevel string, vars ...interface{}) string {
	var logMessage = ""
	if debug {
		var baseMessage string
		if len(vars) == 0 {
			baseMessage = fmt.Sprintf("[%v] %v", logLevel, msg)
		} else {
			baseMessage = fmt.Sprintf("[%v] %v %v", logLevel, msg, vars)
		}
		pc, fn, line, _ := runtime.Caller(2)
		debugMessage := fmt.Sprintf("\n[%s:%s:%d]", runtime.FuncForPC(pc).Name(), fn, line)
		logMessage = strings.Join([]string{baseMessage, debugMessage}, " ")
	} else {
		if len(vars) == 0 {
			logMessage = fmt.Sprintf("[%v] %v", logLevel, msg)
		} else {
			logMessage = fmt.Sprintf("[%v] %v %v", logLevel, msg, vars)
		}
	}
	return logMessage
}

func checkOS() {
	log.Println("Setting logger OS to: " + runtime.GOOS)
	if runtime.GOOS == "windows" {
		reset = ""
		red = ""
		green = ""
		yellow = ""
		blue = ""
		purple = ""
		cyan = ""
		gray = ""
		white = ""
	}
}
