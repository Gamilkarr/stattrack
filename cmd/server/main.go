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
	defer file.Close()
	log.SetOutput(file)
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.InfoLevel)

	cfg, err := config.NewConfig()
	if err != nil {
		log.WithField("fatal error", err).Fatal()
	}

	repo, err := repository.NewRepo(cfg.StoreInterval, cfg.FileStoragePath)
	if err != nil {
		log.WithField("fatal error", err).Fatal()
	}
	h, err := handlers.NewHandler(repo)
	if err != nil {
		log.WithField("fatal error", err).Fatal()
	}

	if cfg.Restore {
		log.WithField("fatal error", repo.Uploading(cfg.FileStoragePath))
	}

	if repo.BackUPPeriod != 0 {
		go repo.RunBackUP()
	}

	route := h.NewRouter()
	if servErr := http.ListenAndServe(cfg.Address, middleware.Logging(middleware.CompressGzip(route))); servErr != nil {
		log.WithField("error", servErr).Fatal("server error")
	}
}
