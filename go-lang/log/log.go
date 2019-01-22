package log

import "os"
import "fmt"

type LogCallback func (s string)

func SetInfoCallback( f LogCallback)	{ infoCallback	= f }
func SetErrorCallback(f LogCallback)	{ errorCallback = f }
func SetFatalCallback(f LogCallback)	{ fatalCallback = f }

func Info(s string) {
	infoCallback(s)
}
func Error(s string) {
	errorCallback(s)
}
func Fatal(s string) {
	fatalCallback(s)
}

func ExtendError(s string, e error) {
	errorCallback(s + e.Error())
}
func ExtendFatal(s string, e error) {
	fatalCallback(s + e.Error())
}

//private variables

var infoCallback	LogCallback = _info
var errorCallback	LogCallback = _error
var fatalCallback	LogCallback = _fatal


//default callback functions

func _info(s string) {
	os.Stdout.Write([]byte(s))
}
func _error(s string) {
	os.Stderr.Write([]byte(s))
}
func _fatal(s string) {
	os.Stderr.Write([]byte(s))
	panic(fmt.Errorf(s))
}
