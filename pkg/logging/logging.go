package logging // import "github.com/echoroaster/roaster/pkg/logging"

import (
	"os"

	"github.com/sirupsen/logrus"
)

var RootLogger = logrus.New()

func init() {
	if os.Getenv("DEBUG") != "" {
		RootLogger.SetLevel(logrus.DebugLevel)
	}
}
