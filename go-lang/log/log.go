package log

import "os"
import "fmt"

type LogCallback func (s ...interface{})

func SetInfoCallback( f LogCallback)	{ infoCallback	= f }
func SetErrorCallback(f LogCallback)	{ errorCallback = f }
func SetFatalCallback(f LogCallback)	{ fatalCallback = f }

func Info(s ...interface{}) {
	infoCallback(s...)
}
func Error(s ...interface{}) {
	errorCallback(s...)
}
func Fatal(s ...interface{}) {
	fatalCallback(s...)
}

func ExtendError(s interface{}, e error) {
	errorCallback(s, e.Error())
}
func ExtendFatal(s interface{}, e error) {
	fatalCallback(s, e.Error())
}

//private variables

var infoCallback	LogCallback = _info
var errorCallback	LogCallback = _error
var fatalCallback	LogCallback = _fatal


//default callback functions

func _info(s ...interface{}) {
	os.Stdout.Write([]byte(fmt.Sprint(s...)))
}

func _error(s ...interface{}) {
	os.Stderr.Write([]byte(fmt.Sprint(s...)))
}

func _fatal(s ...interface{}) {
	os.Stderr.Write([]byte(fmt.Sprint(s...)))
	panic(fmt.Errorf(fmt.Sprint(s...)))
}
