package main

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"

	config "github.com/Gamilkarr/stattrack/configs/server"
	"github.com/Gamilkarr/stattrack/internal/handlers"
	"github.com/Gamilkarr/stattrack/internal/logger"
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

	repo, err := repository.NewRepo()
	if err != nil {
		log.WithField("fatal error", err).Fatal()
	}
	e, err := handlers.NewHandler(repo)
	if err != nil {
		log.WithField("fatal error", err).Fatal()
	}

	cfg, err := config.NewConfig()
	if err != nil {
		log.WithField("fatal error", err).Fatal()
	}

	log.WithField("fatal error", http.ListenAndServe(cfg.Address, logger.WithLogging(e.NewRoute()))).Fatal()
}
