package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"project/config"
	"project/internal/app"
	"project/internal/server"
	"project/pkg/logging"
	"syscall"
)

var (
	cfgPath string
)

func init() {
	flag.StringVar(&cfgPath, "config", "./etc/.env", "config file path")
}

func main() {

	cfg := config.GetConfig(cfgPath)
	logger := logging.GetLogger(cfg.LogLevel)
	logger.Infof("LogLevel %s", cfg.LogLevel)

	appServer, err := server.NewServer(cfg, logger)
	if err != nil {
		logger.Fatal(err)
	}

	appServer.HomeServer()

	srv := new(app.Server)
	go func() {
		srv.StartServer(cfg.Host, cfg.Port, logger)
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	stopType := <-stop

	logger.Info("STOP SERVER: ", stopType)

	if err := srv.StopServer(context.Background()); err != nil {
		return
	}
}
