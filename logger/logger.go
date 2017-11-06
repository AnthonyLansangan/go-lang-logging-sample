//This package utilizes the use of lumberjack
package logger

import (
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"runtime"
)

var (
	trace   *log.Logger
	info    *log.Logger
	warning *log.Logger
	errs    *log.Logger
	fatal   *log.Logger
)

//Setup log configuration
func init() {
	rollingSetings := &lumberjack.Logger{
		Filename: "sample.log",
		MaxSize:  1000, // megabytes
		MaxAge:   30,   //days
		Compress: true,
	}

	trace = log.New(rollingSetings,
		"TRACE: ",
		log.Ldate|log.LUTC)

	info = log.New(rollingSetings,
		"INFO: ",
		log.Ldate|log.Ltime)

	warning = log.New(rollingSetings,
		"WARNING: ",
		log.Ldate|log.Ltime)

	errs = log.New(rollingSetings,
		"ERROR: ",
		log.Ldate|log.Ltime)

	fatal = log.New(rollingSetings,
		"FATAL: ",
		log.Ldate|log.Ltime)

}

//Prints Trace level log
func Trace(v ...interface{}) {
	trace.Println(getCaller(), v)
}

//Prints Info level log
func Info(v ...interface{}) {
	info.Println(getCaller(), v)
}

//Prints Warning level log
func Warning(v ...interface{}) {
	warning.Println(getCaller(), v)
}

//Prints Error level log
func Error(v ...interface{}) {
	errs.Println(getCaller(), v)
}

//Prints Fatal level log
func Fatal(v ...interface{}) {
	fatal.Fatalln(getCaller(), v)
}

func getCaller() string {

	// we get the callers as uintptrs - but we just need 1
	fpcs := make([]uintptr, 1)

	// skip 3 levels to get to the caller of whoever called Caller()
	//	n :=
	runtime.Callers(3, fpcs)

	// get the info of the actual function that's in the pointer
	fun := runtime.FuncForPC(fpcs[0] - 1)

	// return its name
	return fun.Name()
}
