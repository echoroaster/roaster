package logging // import "github.com/echoroaster/roaster/pkg/logging"

import (
	"os"

	"github.com/sirupsen/logrus"
)

var RootLogger = logrus.New()

func init() {
	if os.Getenv("DEBUG") != "" {
		RootLogger.SetLevel(logrus.DebugLevel)
	} else {
		switch os.Getenv("LOG_LEVEL") {
		case "trace":
			RootLogger.SetLevel(logrus.TraceLevel)
		case "debug":
			RootLogger.SetLevel(logrus.DebugLevel)
		case "info":
			RootLogger.SetLevel(logrus.InfoLevel)
		case "warn":
			RootLogger.SetLevel(logrus.WarnLevel)
		case "error":
			RootLogger.SetLevel(logrus.ErrorLevel)
		default:
		}
	}
}
