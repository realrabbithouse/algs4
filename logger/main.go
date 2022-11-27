package main

// From Readme of: https://github.com/sirupsen/logrus
import (
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors:   false,
		DisableSorting:  true,
		ForceColors:     true,
		TimestampFormat: time.RFC822,
		FullTimestamp:   true,
	})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logrus.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	logrus.SetLevel(logrus.TraceLevel)

	// Whether report the calling method.
	logrus.SetReportCaller(false)
}

func main() {
	logrus.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	logrus.WithFields(logrus.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	logrus.WithFields(logrus.Fields{
		"omg":    true,
		"number": 100,
	}).Error("The ice breaks!")

	// A common pattern is to re-use fields between logging statements by re-using
	// the logrus.Entry returned from WithFields()
	contextLogger := logrus.WithFields(logrus.Fields{
		"common": "this is a common field",
		"other":  "I also should be logged always",
	})

	contextLogger.Info("I'll be logged with common and other field")
	contextLogger.Info("Me too")

	log := logrus.StandardLogger() // log := logrus.New()
	// 设置 hook
	hook := MyHook{AppName: "awesome-app"}
	log.AddHook(&hook)

	log.Trace("trace msg")
	log.Debug("debug msg")
	log.Info("info msg")
	log.Warn("warn msg")
	log.Error("error msg")
}

// MyHook user defined hook for logrus.
type MyHook struct {
	AppName string
}

// Levels specifies in what levels this hook will be triggered.
func (h *MyHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// Fire represents the hook action.
func (h *MyHook) Fire(entry *logrus.Entry) error {
	entry.Data["app"] = h.AppName
	return nil
}
