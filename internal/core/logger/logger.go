package core_logger

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	*logrus.Logger

	file os.File
}

func NewLogger(config Config) (*logrus.Logger, error) {
	log := logrus.New()

	parsedLevel, err := logrus.ParseLevel(config.Level)
	if err != nil {
		parsedLevel = logrus.DebugLevel
	}

	if err := os.MkdirAll(config.Folder, 0755); err != nil {
		return nil, fmt.Errorf("mkdir log folder: %w", err)
	}

	timestamp := time.Now().UTC().Format("2006-01-02T15-04-05.000000")
	logFilePath := filepath.Join(
		config.Folder,
		fmt.Sprintf("%s.log", timestamp),
	)

	logFile, _ := os.OpenFile(
		logFilePath,
		os.O_CREATE|os.O_WRONLY,
		0755,
	)

	log.SetLevel(parsedLevel)
	log.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339,
	})

	log.SetOutput(io.MultiWriter(os.Stdout, logFile))

	return log, nil
}
