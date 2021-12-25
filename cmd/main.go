package main

import (
	"database/sql"
	"flag"
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

	db, err := newDB(config)
	if err != nil {
		logrus.Fatalf("Error while opening db: %v", err)
	}
	defer db.Close()

	repository := repository.NewRepository(config, db)
	services := service.NewService(repository)
	h := handler.NewHandler(services)
	if err != nil {
		logrus.Fatalf("Error parsing config file: %v", err)
	}
	srv := go_booc_exchange.NewServer(config)

	if err := srv.Run(h); err != nil {
		logrus.Fatalf("Error initing server: %s", err.Error())
	}
}

func newDB(config *configs.Config) (*sql.DB, error) {
	db, err := sql.Open(config.DB.DriverName, config.DB.DatabaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
