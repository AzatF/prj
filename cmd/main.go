package main

import (
	"flag"
	"project/config"
	"project/pkg/logging"
	"project/server"
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

	//test.TestingApp(cfg, logger)

	appServer.HomeServer(cfg.SupportHost, "8888", logger)

}
