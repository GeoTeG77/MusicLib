package logger

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func InitLogger() (*log.Logger, *os.File, error) {
	logger := log.New()
	logger.SetOutput(os.Stdout)
	logger.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	logger.SetLevel(log.DebugLevel)

	return logger, nil, nil
}
