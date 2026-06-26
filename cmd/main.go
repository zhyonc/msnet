package main

import (
	"fmt"
	"log/slog"

	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/zhyonc/msnet"
	"github.com/zhyonc/msnet/internal/server"
)

func main() {
	// Installation
	msnet.New(&msnet.Setting{
		LocaleRegion:   msnet.GMS,
		MSRegion:       msnet.GMS,
		MSVersion:      95,
		MSMinorVersion: "1",
	})
	// Set logger
	logFilename := fmt.Sprintf(
		"%s-server-%s.log",
		msnet.ServerType,
		time.Now().Format("2006-01-02_15-04-05"),
	)
	closeFile := msnet.SetLogger(msnet.LogBackupDir, logFilename, slog.LevelDebug, false)
	defer closeFile()
	// New Server
	s := server.NewServer(msnet.ServerAddr)
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
				return
			default:
				slog.Info("other signal", "syscall", sig)
			}
		}
	}()
	s.Run()
}
