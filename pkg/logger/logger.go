package logger

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"runtime/debug"
	"sync"
	"time"

	"github.com/davecgh/go-spew/spew"
)

const (
	levelDebug = `DEBUG`
	levelInfo  = `INFO`
	levelError = `ERROR`
	levelFatal = `FATAL`
)

var logger = NewLogger()

// NewLogger instanciate a new Logger which output to Stdout
func NewLogger() (l Logger) {
	l.SetOutput(os.Stdout)
	return
}

// SetOutput defines the output for the global logger
func SetOutput(o io.Writer) {
	logger.SetOutput(o)
}

// LogFatal logs the error with the global logger then exit if the error is not nil
func LogFatal(err error, args ...interface{}) {
	logger.LogFatal(err, args...)
}

// LogError logs the error with the global logger
func LogError(err error, args ...interface{}) error {
	return logger.LogError(err, args...)
}

// LogInfo logs the message with the global logger
func LogInfo(msg string, args ...interface{}) {
	logger.LogInfo(msg, args...)
}

// LogDebug logs the data with the global logger
func LogDebug(args ...interface{}) {
	logger.LogDebug(args...)
}

// Logger used to write logs to its out stream
type Logger struct {
	out   io.Writer
	mutex sync.Mutex
}

// SetOutput set the output of the logger
func (l *Logger) SetOutput(o io.Writer) {
	l.out = o
}

func (l *Logger) log(t, msg string) {
	msg = fmt.Sprintf("[%s] %s - %s\n", t, time.Now().Format("2006-01-02 15:04:05"), msg)

	l.mutex.Lock()
	defer l.mutex.Unlock()

	l.out.Write([]byte(msg))
}

// LogFatal logs the error then exit if the error is not nil
func (l *Logger) LogFatal(err error, args ...interface{}) {
	if err == nil {
		return
	}
	description := fmt.Sprintf("New runtime error :\n%s\n\n", err.Error())
	if len(args) > 0 {
		for i, arg := range args {
			description += fmt.Sprintf("\n- Arg %d: %s", i, spew.Sdump(arg))
		}
	}
	description += fmt.Sprintf("\nBacktrace:\n%s\n\n", debug.Stack())

	l.log(levelFatal, description)

	os.Exit(1)
}

// LogError logs the error if the error is not nil
func (l *Logger) LogError(err error, args ...interface{}) error {
	if err == nil {
		return nil
	}

	description := fmt.Sprintf("New runtime error :\n%s\n\n", err.Error())
	if len(args) > 0 {
		for i, arg := range args {
			description += fmt.Sprintf("\n- Arg %d: %s", i, spew.Sdump(arg))
		}
	}
	description += fmt.Sprintf("\nBacktrace:\n%s\n\n", debug.Stack())

	l.log(levelError, description)

	return err
}

// LogInfo logs the message
func (l *Logger) LogInfo(msg string, args ...interface{}) {
	if len(args) > 0 {
		msg += fmt.Sprintf("\nArgs: %#v", args)
	}
	l.log(levelInfo, msg)
}

// LogDebug logs the debug data
func (l *Logger) LogDebug(args ...interface{}) {
	var msg string

	for i, arg := range args {
		msg += fmt.Sprintf("\n- Arg %d: %s", i, spew.Sdump(arg))
	}

	pc, fn, line, _ := runtime.Caller(2)

	msg = fmt.Sprintf("%s\n\t in %s (%s-%d)", msg, runtime.FuncForPC(pc).Name(), fn, line)

	l.log(levelDebug, msg)
}
