package core_logger

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	*logrus.Logger

	file *os.File
}

type loggerKey struct{}

var (
	key = loggerKey{}
)

func NewLogger(config Config) (*Logger, error) {
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

	return &Logger{
		Logger: log,
		file: logFile,
	}, nil
}

func ToContext(ctx context.Context, log *logrus.Entry) context.Context {
	return context.WithValue(ctx, key, log)
}

func FromContext(ctx context.Context) *logrus.Entry {
	log, ok := ctx.Value(key).(*logrus.Entry)
	if !ok {
		panic("unable to get logger from context")
	}

	return log
}
