package main

import (
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"

	config "github.com/Gamilkarr/stattrack/configs/server"
	"github.com/Gamilkarr/stattrack/internal/handlers"
	"github.com/Gamilkarr/stattrack/internal/middleware"
	"github.com/Gamilkarr/stattrack/internal/repository"
)

func main() {
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.WithField("error", err).Error("file closing error")
		}
	}(file)
	log.SetOutput(file)
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.InfoLevel)

	cfg, err := config.NewConfig()
	if err != nil {
		log.WithField("error", err).Fatal("config error")
	}

	repo := repository.NewRepo(cfg.StoreInterval, cfg.FileStoragePath)
	handler := handlers.NewHandler(repo)
	route := handler.NewRouter()

	if cfg.Restore {
		err := repo.Uploading(cfg.FileStoragePath)
		if err != nil {
			log.WithField("error", err).Error("backup data loading error")
		}
	}

	if repo.BackUPPeriod != 0 {
		go repo.RunBackUP()
	}

	if servErr := http.ListenAndServe(cfg.Address, middleware.Logging(middleware.CompressGzip(route))); servErr != nil {
		log.WithField("error", servErr).Fatal("server error")
	}
}
