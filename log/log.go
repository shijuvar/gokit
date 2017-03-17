package log

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

const (
	// UNSPECIFIED logs nothing
	UNSPECIFIED Level = iota // 0 :
	// TRACE logs everything
	TRACE // 1
	// INFO logs Info, Warnings and Errors
	INFO // 2
	// WARNING logs Warning and Errors
	WARNING // 3
	// ERROR just logs Errors
	ERROR // 4
)

// Level holds the log level.
type Level int

// Package level variables, which are pointer to log.Logger.
var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

// initLog initializes log.Logger objects
func initLog(
	traceHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer,
	isFlag bool) {

	// Flags for defines the logging properties, to log.New
	flag := 0
	if isFlag {
		flag = log.Ldate | log.Ltime | log.Lshortfile
	}

	// Create log.Logger objects.
	Trace = log.New(traceHandle, "TRACE: ", flag)
	Info = log.New(infoHandle, "INFO: ", flag)
	Warning = log.New(warningHandle, "WARNING: ", flag)
	Error = log.New(errorHandle, "ERROR: ", flag)

}

// SetLogLevel sets the logging level preference
func SetLogLevel(level Level, logFile string) {

	// Creates os.*File, which has implemented io.Writer intreface
	f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening log file: %s", err.Error())
	}

	// Calls function initLog by specifying log level preference.
	switch level {
	case TRACE:
		initLog(f, f, f, f, true)
		return

	case INFO:
		initLog(ioutil.Discard, f, f, f, true)
		return

	case WARNING:
		initLog(ioutil.Discard, ioutil.Discard, f, f, true)
		return
	case ERROR:
		initLog(ioutil.Discard, ioutil.Discard, ioutil.Discard, f, true)
		return

	default:
		initLog(ioutil.Discard, ioutil.Discard, ioutil.Discard, ioutil.Discard, false)
		f.Close()
		return

	}
}
