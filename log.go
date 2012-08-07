package gou

import (
	"fmt"
	"log"
	"sync"
)

const (
	NOLOGGING = -1
	FATAL     = 0
	ERROR     = 1
	WARN      = 2
	INFO      = 3
	DEBUG     = 4
)

/*
RED = '\033[0;1;31m'
GREEN = '\033[0;1;32m'
YELLOW = '\033[0;1;33m'
BLUE = '\033[0;1;34m'
MAGENTA = '\033[0;1;35m'
CYAN = '\033[0;1;36m'
WHITE = '\033[0;1;37m'
DARK_MAGENTA = '\033[0;35m'
ANSI_RESET = '\033[0m'
*/

var (
	LogLevel  int = ERROR
	logger    *log.Logger
	logPrefix string = "SMOKR "
	LogColor         = map[int]string{FATAL: "\033[0m\033[37m",
		ERROR: "\033[0m\033[31m",
		WARN:  "\033[0m\033[33m",
		INFO:  "\033[0m\033[35m",
		DEBUG: "\033[0m\033[34m"}
	LogLevelWords map[string]int      = map[string]int{"fatal": 0, "error": 1, "warn": 2, "info": 3, "debug": 4, "none": -1}
	eventHandlers map[string][]func() = make(map[string][]func())
	eventsMu      *sync.Mutex         = new(sync.Mutex)
)

// you can set a logger, and log level,most common usage is:
//
//	smokr.SetLogger(log.New(os.Stderr, "", log.LstdFlags), "debug")
//
//  loglevls:   debug, info, warn, error, fatal
func SetLogger(l *log.Logger, logLevel string) {
	logger = l
	LogLevelSet(logLevel)
}

// you can set a logger, prefix
func SetLoggerPrefix(prefix string) {
	logPrefix = prefix
}

// sets the log level from a string
func LogLevelSet(levelWord string) {
	if lvl, ok := LogLevelWords[levelWord]; ok {
		LogLevel = lvl
	}
}

// Log at debug level
func Debug(v ...interface{}) {
	if logger != nil && LogLevel >= 4 {
		//logger.Output(2, fmt.Sprintln(v...))
		DoLog(3, DEBUG, fmt.Sprint(v...), logger)
	}
}

func Debugf(format string, v ...interface{}) {
	if LogLevel >= 4 {
		DoLog(3, DEBUG, fmt.Sprintf(format, v...), logger)
	}
}

// Log to logger if setup
//    Log(ERROR, "message")
func Log(logLvl int, v ...interface{}) {
	if LogLevel >= logLvl {
		DoLog(3, logLvl, fmt.Sprint(v...), logger)
	}
}

// Log to logger if setup
//    Logf(ERROR, "message %d", 20)
func Logf(logLvl int, format string, v ...interface{}) {
	if LogLevel >= logLvl {
		DoLog(3, logLvl, fmt.Sprintf(format, v...), logger)
	}
}

func DoLog(depth, logLvl int, msg string, lgr *log.Logger) {
	if lgr != nil {
		lgr.Output(depth, logPrefix+LogColor[logLvl]+msg+"\033[0m")
	}
}
