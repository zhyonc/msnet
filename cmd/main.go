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
		MSRegion:       msnet.GMS,
		MSVersion:      95,
		MSMinorVersion: "1",
	})
	// Set logger
	logFilename := fmt.Sprintf("%s-server-%s.log", msnet.SERVER_TYPE, time.Now().Format("2006-01-02_15-04-05"))
	closeFile := msnet.SetLogger(msnet.LOG_BACKUP_DIR, logFilename, slog.LevelDebug)
	defer closeFile()
	// New Server
	s := server.NewServer(msnet.SERVER_ADDR)
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
