package main

import (
	"flag"
	"project/config"
	"project/internal/app"
	"project/internal/server"
	"project/pkg/logging"
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

	app.StartServer(cfg.Host, cfg.Port, logger)

}
