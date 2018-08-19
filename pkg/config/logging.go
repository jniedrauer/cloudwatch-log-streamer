package config

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

// SetLogLevel sets the log level of a logger object to the configured
// application log level.
func SetLogLevel(logObj *logrus.Logger) error {
	logLevel, err := logrus.ParseLevel(Properties["LogLevel"])
	if err != nil {
		return fmt.Errorf("Error while setting log level: %v", err)
	}

	logObj.SetLevel(logLevel)

	return nil
}
