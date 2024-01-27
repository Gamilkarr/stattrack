package main

import (
	"github.com/go-chi/chi/v5"
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
	h, err := handlers.NewHandler(repo)
	if err != nil {
		log.WithField("fatal error", err).Fatal()
	}

	cfg, err := config.NewConfig()
	if err != nil {
		log.WithField("fatal error", err).Fatal()
	}

	log.WithField("fatal error", http.ListenAndServe(cfg.Address, logger.WithLogging(NewRouter(h)))).Fatal()
}

func NewRouter(h *handlers.Handler) *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", h.GetMetrics)

	r.Post("/update/", h.UpdateJSONMetrics)
	r.Post("/update/{type}/{name}/{value}", h.UpdateMetrics)

	r.Post("/value/", h.GetJSONValueMetric)
	r.Get("/value/{type}/{name}", h.GetValueMetric)

	return r
}
