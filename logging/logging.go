/*
Simple logging wrapper on top of log package to have leveled log messages
*/
package logging

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"github.com/eBay/fabio/config"
)

// Different log levels with logger type
var (
	tracel   *log.Logger
	debugl   *log.Logger
	infol    *log.Logger
	warningl *log.Logger
	errorl   *log.Logger
)

// log wrapper which creates new logger objects as per the handle defined

func logWrapper(
	traceHandle io.Writer,
	debugHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {

	tracel = log.New(traceHandle,
		"[TRACE] ",
		log.Ldate|log.Ltime|log.Lshortfile)

	debugl = log.New(debugHandle,
		"[DEBUG] ",
		log.Ldate|log.Ltime|log.Lshortfile)

	infol = log.New(infoHandle,
		"[INFO] ",
		log.Ldate|log.Ltime|log.Lshortfile)

	warningl = log.New(warningHandle,
		"[WARNING] ",
		log.Ldate|log.Ltime|log.Lshortfile)

	errorl = log.New(errorHandle,
		"[ERROR] ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

var std = log.New(os.Stderr, "[FATAL] ", log.Ldate|log.Ltime|log.Lshortfile)

func Fatal(v ...interface{}) {
	std.Output(2, fmt.Sprint(v...))
	os.Exit(1)
}

func Trace(args ...interface{}) {
	// default format string used when only a single string is passed
	format := "%s\n"

	length := len(args)

	if length > 1 {
		// first arg is the format string
		format = args[0].(string)
		// delete the first arg from the variadic array
		args = append(args[1:])
	}

	tracel.Printf(format, args...)
}

func Debug(args ...interface{}) {
	// default format string used when only a single string is passed
	format := "%s\n"

	length := len(args)

	if length > 1 {
		// first arg is the format string
		format = args[0].(string)
		// delete the first arg from the variadic array
		args = append(args[1:])
	}

	debugl.Printf(format, args...)
}

func Info(args ...interface{}) {
	// default format string used when only a single string is passed
	format := "%s\n"

	length := len(args)

	if length > 1 {
		// first arg is the format string
		format = args[0].(string)
		// delete the first arg from the variadic array
		args = append(args[1:])
	}

	infol.Printf(format, args...)
}


func Warn(args ...interface{}) {
	// default format string used when only a single string is passed
	format := "%s\n"

	length := len(args)

	if length > 1 {
		// first arg is the format string
		format = args[0].(string)
		// delete the first arg from the variadic array
		args = append(args[1:])
	}

	warningl.Printf(format, args...)
}

func Error(args ...interface{}) {
	// default format string used when only a single string is passed
	format := "%s\n"

	length := len(args)

	if length > 1 {
		// first arg is the format string
		format = args[0].(string)
		// delete the first arg from the variadic array
		args = append(args[1:])
	}

	errorl.Printf(format, args...)
}
// Function that gets called when package is imported. Reads LOGLEVEL from env
// variable else defults to Info.

func init() {
	cfg, err := config.Load(os.Args, os.Environ())
	if err != nil {
		Fatal("[FATAL] %s", err)
	}
	loglevel := strings.ToLower(cfg.Logging.Level)
	levellist := []string{"trace", "debug", "info", "warning", "error"}
	if !isValueInList(loglevel, levellist) {
		loglevel = "info"
	}
	switch loglevel {
	case "error":
		logWrapper(ioutil.Discard, ioutil.Discard, ioutil.Discard, ioutil.Discard, os.Stderr)
	case "warning":
		logWrapper(ioutil.Discard, ioutil.Discard, ioutil.Discard, os.Stderr, os.Stderr)
	case "info":
		logWrapper(ioutil.Discard, ioutil.Discard, os.Stdout, os.Stderr, os.Stderr)
	case "debug":
		logWrapper(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr, os.Stderr)
	case "trace":
		logWrapper(os.Stdout, os.Stdout, os.Stdout, os.Stderr, os.Stderr)
	}
}

func isValueInList(level string, levellist []string) bool {
	for _, lev := range levellist {
		if lev == level {
			return true
		}
	}
	return false
}
