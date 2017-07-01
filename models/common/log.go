package common

import (
	"fmt"
	"os"
	"time"

	logging "github.com/iteny/hmgo/go-logging"
)

var Log = logging.MustGetLogger("example")

// Example format string. Everything except the message has a custom color
// which is dependent on the log level. Many fields have a custom output
// formatting too, eg. the time returns the hour down to the milli second.
var format = logging.MustStringFormatter(
	`%{color}%{time:15:04:05.000} %{shortfunc} > %{level:.4s} %{id:03x}%{color:reset} %{message}`,
)

// Password is just an example type implementing the Redactor interface. Any
// time this is logged, the Redacted() function will be called.
type Password string

func (p Password) Redacted() interface{} {
	return logging.Redact(string(p))
}

func init() {
	logFile, err := os.OpenFile("./log/log.txt", os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
	}
	// For demo purposes, create two backend for os.Stderr.
	backend1 := logging.NewLogBackend(logFile, "", 0)
	backend2 := logging.NewLogBackend(os.Stderr, "", 0)
	// For messages written to backend2 we want to add some additional
	// information to the output, including the used log level and the name of
	// the function.
	backend2Formatter := logging.NewBackendFormatter(backend2, format)
	// Only errors and more severe messages should be sent to backend1
	backend1Leveled := logging.AddModuleLevel(backend1)
	// Set the backends to be used.
	backend1Leveled.SetLevel(logging.INFO, "")

	logging.SetBackend(backend1Leveled, backend2Formatter)

}
func Loger(level string, args ...interface{}) {
	switch level {
	case "debug":
		Log.Debug(args)
	case "info":
		Log.Info(args)
	case "notice":
		Log.Notice(args)
	case "warning":
		Log.Warning(args)
	case "error":
		Log.Error(args)
	default:
		Log.Info("first paramter error!")
	}
}
func LogerInsertText(path string, err error) {
	Log.Critical("time:"+time.Now().Format("2006-01-02 15:04:05"), "path:"+path, "info:"+err.Error())
}
