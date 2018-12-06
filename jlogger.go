package jlogger

import (
	"fmt"
	"log"
	"os"
)

const (
	// TRACE output the logs
	TRACE = "TRACE"
	// DEBUG debug info
	DEBUG = "DEBUG"
	// INFO output info
	INFO = "INFO"
	// ERROR error info
	ERROR = "ERROR"
	// WARN warnning
	WARN = "WARN"
)

// JLogger interface for print log
type JLogger interface {
	Trace(...interface{})
	Debug(...interface{})
	DebugF(formatter string, args ...interface{})
	Info(...interface{})
	InfoF(formatter string, args ...interface{})
	Error(...interface{})
	ErrorF(formatter string, args ...interface{})
	Warn(...interface{})
	WarnF(formatter string, args ...interface{})
}

type jLog struct {
	Logger *log.Logger
}

func merge(v ...interface{}) string {
	var str string
	for _, item := range v {
		str += fmt.Sprintf("%s", item)
	}
	fmt.Println(str)
	return str
}

func (l *jLog) Trace(args ...interface{}) {
	l.Logger.Println(TRACE, merge(args))
}

func (l *jLog) Debug(args ...interface{}) {

	l.Logger.Println(DEBUG, merge(args))
}

func (l *jLog) DebugF(formatter string, args ...interface{}) {
	l.Logger.Printf(DEBUG+formatter, args)
}

func (l *jLog) Info(args ...interface{}) {
	l.Logger.Println(INFO, merge(args))
}

func (l *jLog) InfoF(formatter string, args ...interface{}) {
	l.Logger.Printf(INFO+formatter, args)
}

func (l *jLog) Error(args ...interface{}) {
	l.Logger.Println(ERROR, merge(args))
}

func (l *jLog) ErrorF(formatter string, args ...interface{}) {
	l.Logger.Printf(ERROR+formatter, args)
}

func (l *jLog) Warn(args ...interface{}) {
	l.Logger.Println(WARN, merge(args))
}

func (l *jLog) WarnF(formatter string, args ...interface{}) {
	l.Logger.Printf(WARN+formatter, args)
}

// New instance of Jlog
func New(fileName string) JLogger {

	var l = &jLog{}
	l.Logger = newLoggerFile(fileName, true)
	return l
}

func newLoggerFile(fileName string, debug bool) *log.Logger {

	if debug {
		fileName += ".deubg"
	}
	fileName = fmt.Sprintf("%s.log", fileName)
	osFile, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	logger := log.New(osFile, "", log.Ldate|log.Ltime|log.Lshortfile)

	return logger
}
