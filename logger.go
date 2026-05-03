package msnet

import (
	"io"
	"log/slog"
	"os"
	"path"
)

func SetLogger(backupDir string, filename string, level slog.Level) func() {
	var writer io.Writer
	var logFile *os.File
	if backupDir != "" {
		err := os.MkdirAll(backupDir, os.ModePerm)
		if err != nil {
			panic("Failed to create dir: " + backupDir)
		}
		filePath := path.Join(backupDir, filename)
		logFile, err = os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			panic(err)
		}
		writer = io.MultiWriter(os.Stderr, logFile)
	} else {
		writer = os.Stderr
	}
	logLevelVar := new(slog.LevelVar)
	logLevelVar.Set(level)
	loggerHandler := slog.NewTextHandler(writer, &slog.HandlerOptions{Level: logLevelVar})
	logger := slog.New(loggerHandler)
	slog.SetDefault(logger)
	return func() {
		if logFile != nil {
			logFile.Close()
		}
	}
}
