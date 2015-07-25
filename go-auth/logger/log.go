package logger

import (
	"fmt"
	"os"

	"github.com/op/go-logging"
)

var format = logging.MustStringFormatter(
	"%{color}%{time:15:04:05.000} %{shortfile} %{shortfunc} %{level:.10s} %{color:reset} %{message}",
)

func SetupLogging(levelStr string) {
	level, err := logging.LogLevel(levelStr)
	if err != nil {
		fmt.Printf("Unable to understand log level %s\n", levelStr)
		return
	}

	backend := logging.NewLogBackend(os.Stderr, "", 0)
	backendFormatter := logging.NewBackendFormatter(backend, format)
	backendLeveled := logging.AddModuleLevel(backendFormatter)

	backendLeveled.SetLevel(level, "")
	logging.SetBackend(backendLeveled)
}
