package msnet

import (
	"io"
	"log/slog"
	"os"
	"path"
)

func SetLogger(backupDir string, filename string, level slog.Level, addSource bool) func() {
	var writer io.Writer
	var logFile *os.File
	if backupDir != "" {
		err := os.MkdirAll(backupDir, 0750)
		if err != nil {
			panic("Failed to create dir: " + backupDir)
		}
		filePath := path.Join(backupDir, filename)
		logFile, err = os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
		if err != nil {
			panic(err)
		}
		writer = io.MultiWriter(os.Stderr, logFile)
	} else {
		writer = os.Stderr
	}
	logLevelVar := new(slog.LevelVar)
	logLevelVar.Set(level)
	loggerHandler := slog.NewTextHandler(writer, &slog.HandlerOptions{
		Level:     logLevelVar,
		AddSource: addSource,
	})
	logger := slog.New(loggerHandler)
	slog.SetDefault(logger)
	return func() {
		if logFile != nil {
			err := logFile.Close()
			if err != nil {
				slog.Error(err.Error())
			}
		}
	}
}
