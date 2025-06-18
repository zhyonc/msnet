package main

import (
	"fmt"
	"log/slog"

	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/zhyonc/msnet"
	"github.com/zhyonc/msnet/internal/enum"
	"github.com/zhyonc/msnet/internal/server"
)

const (
	serverType   string     = "login"
	serverAddr   string     = "127.0.0.1:8484"
	logBackupDir string     = "./log"
	logLevel     slog.Level = slog.LevelDebug
)

func main() {
	// Installation
	msnet.New(&msnet.Setting{
		MSRegion:       enum.GMS,
		MSVersion:      95,
		MSMinorVersion: "1",
	})
	// Set logger
	done := make(chan bool, 1)
	logFilename := fmt.Sprintf("%s-server-%s.log", serverType, time.Now().Format("2006-01-02_15-04-05"))
	msnet.SetLogger(logBackupDir, logFilename, logLevel, done)
	// New Server
	s := server.NewServer(serverAddr)
	// Avoid unexpected exit
	sch := make(chan os.Signal, 1)
	signal.Notify(sch, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		for sig := range sch {
			switch sig {
			case syscall.SIGTERM, syscall.SIGINT:
				slog.Info("Server will shutdown after 3s")
				time.Sleep(3 * time.Second)
				s.Shutdown()
				done <- true
				return
			default:
				slog.Info("other signal", "syscall", sig)
			}
		}
	}()
	s.Run()
}
