package logging

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

func config() {
	customFormatter := new(logrus.TextFormatter)
	customFormatter.FullTimestamp = true
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	customFormatter.DisableTimestamp = false
	customFormatter.DisableColors = false

	Log.SetFormatter(customFormatter)
	Log.SetOutput(os.Stdout)
	Log.SetLevel(logrus.TraceLevel)
}

func init() {
	config()
}
