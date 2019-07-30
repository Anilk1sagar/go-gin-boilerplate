package utils

import (
	"os"

	"github.com/shiena/ansicolor"
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

// Logger ...
func Logger() *logrus.Logger {

	logger := &logrus.Logger{
		Out:   os.Stderr,
		Level: logrus.DebugLevel,
		Formatter: &prefixed.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp:   true,
			ForceFormatting: true,
			ForceColors:     true,
		},
	}

	logger.Out = ansicolor.NewAnsiColorWriter(os.Stdout)

	return logger
}
