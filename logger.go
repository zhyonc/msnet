package msnet

import (
	"io"
	"log/slog"
	"os"
	"path"
)

func SetLogger(backupDir string, filename string, level slog.Level, done chan bool) {
	var writer io.Writer
	if backupDir != "" {
		err := os.MkdirAll(backupDir, os.ModePerm)
		if err != nil {
			panic("Failed to create dir: " + backupDir)
		}
		filePath := path.Join(backupDir, filename)
		logFile, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			panic(err)
		}
		go func() {
			<-done
			logFile.Close()
		}()
		writer = io.MultiWriter(os.Stderr, logFile)
	} else {
		writer = os.Stderr
	}
	logLevelVar := new(slog.LevelVar)
	logLevelVar.Set(slog.LevelDebug)
	loggerHandler := slog.NewTextHandler(writer, &slog.HandlerOptions{Level: logLevelVar})
	logger := slog.New(loggerHandler)
	slog.SetDefault(logger)
}
