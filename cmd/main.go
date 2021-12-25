package main

import (
	"flag"
	"fmt"
	"github.com/pArtour/book-exchange"
	"github.com/pArtour/book-exchange/configs"
	"github.com/pArtour/book-exchange/pkg/handler"
	"github.com/pArtour/book-exchange/pkg/repository"
	"github.com/pArtour/book-exchange/pkg/service"
	"github.com/sirupsen/logrus"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/config.yml", "path to config file")
}

func main() {
	flag.Parse()
	config, err := configs.InitConfig(configPath)
	repository := repository.NewRepository()
	services := service.NewService(repository)
	h := handler.NewHandler(services)
	if err != nil {
		logrus.Fatalf("Error parsing config file: %v", err)
	}
	srv := go_booc_exchange.NewServer(config)
	fmt.Printf("Server running on port %s", config.Server.BindAddr)

	if err := srv.Run(h); err != nil {
		logrus.Fatalf("Error initing server: %s", err.Error())
	}
}
