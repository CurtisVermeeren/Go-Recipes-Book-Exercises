package main

import (
	"errors"
	"flag"
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

// Package level variables, which are pointers to log.Logger.
var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

// initLog initializes log.Logger objects
func initLog(traceHandle io.Writer, infoHandle io.Writer, warningHandle io.Writer, errorHandle io.Writer, isFlag bool) {
	// Flags for defining the logging properties, to log.New
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
func SetLogLevel(level Level) *os.File {
	// Calls function initLog by specifying log level preference.
	// Creates os.*File, which has implemented io.Writer interface
	f, err := os.OpenFile("logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening log file: %s", err.Error())
	}
	switch level {
	case TRACE:
		initLog(f, f, f, f, true)
		return f
	case INFO:
		initLog(ioutil.Discard, f, f, f, true)
		return f
	case WARNING:
		initLog(ioutil.Discard, ioutil.Discard, f, f, true)
		return f
	case ERROR:
		initLog(ioutil.Discard, ioutil.Discard, ioutil.Discard, f, true)
		return f
	default:
		initLog(ioutil.Discard, ioutil.Discard, ioutil.Discard, ioutil.Discard, false)
		f.Close()
		return nil
	}
}
func main() {
	// Parse log level from command line
	logLevel := flag.Int("loglevel", 0, "an integer value (0-4)")
	flag.Parse()

	// Calling the SetLogLevel with the command-line argument
	f := SetLogLevel(Level(*logLevel))
	if f != nil {
		defer f.Close()
	}
	Trace.Println("Main started")
	loop()
	err := errors.New("Sample Error")
	Error.Println(err.Error())
	Trace.Println("Main completed")
}

// A simple function for the logging demo
func loop() {
	Trace.Println("Loop started")
	for i := 0; i < 10; i++ {
		Info.Println("Counter value is:", i)
	}
	Warning.Println("The counter variable is not being used")
	Trace.Println("Loop completed")
}
