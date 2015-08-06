package logger

import (
	"fmt"
	"os"

	"github.com/op/go-logging"
)

var format = logging.MustStringFormatter(
	"%{color}%{time:15:04:05.000} %{shortfile} %{shortfunc} %{level:.10s} %{color:reset} %{message}",
)

func SetupLogging(levelStr string, logType string) {
	level, err := logging.LogLevel(levelStr)
	if err != nil {
		fmt.Printf("Unable to understand log level %s\n", levelStr)
		return
	}
	var backend logging.Backend
	switch logType {
	case "syslog":
		backend, err = logging.NewSyslogBackend("go-auth")
		if err != nil {
			fmt.Printf("Unable to create syslog backend: %s\n", err.Error())
			return
		} else {
			fmt.Printf("Logging to syslog\n")
		}
	case "console":
		backend = logging.NewLogBackend(os.Stderr, "", 0)
	default:
		fmt.Printf("Unrecognised log format")
		return
	}

	backendFormatter := logging.NewBackendFormatter(backend, format)
	backendLeveled := logging.AddModuleLevel(backendFormatter)

	backendLeveled.SetLevel(level, "")
	logging.SetBackend(backendLeveled)
}
