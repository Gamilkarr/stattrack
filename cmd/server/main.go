package main

import (
	"encoding/json"
	config "github.com/Gamilkarr/stattrack/configs/server"
	"github.com/Gamilkarr/stattrack/internal/compress"
	"github.com/Gamilkarr/stattrack/internal/handlers"
	"github.com/Gamilkarr/stattrack/internal/logger"
	"github.com/Gamilkarr/stattrack/internal/models"
	"github.com/Gamilkarr/stattrack/internal/repository"
	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
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

	if cfg.Restore {
		log.WithField("fatal error", uploading(cfg.FileStoragePath, repo))
	}

	go backUP(cfg.FileStoragePath, cfg.StoreInterval, repo)

	log.WithField("fatal error", http.ListenAndServe(cfg.Address, logger.WithLogging(compress.GzipMiddleware(NewRouter(h))))).Fatal()
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

func backUP(path string, period time.Duration, repo *repository.MemStorage) {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.WithField("backup error", err)
	}
	defer file.Close()
	for {
		time.Sleep(period)
		if err != nil {
			log.WithField("backup error", err)
		}

		enc := json.NewEncoder(file)

		if err = enc.Encode(repo.GetMetrics()); err != nil {
			log.WithField("backup error", err)
		}
	}
}

func uploading(path string, repo *repository.MemStorage) error {
	file, err := os.OpenFile(path, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	dec := json.NewDecoder(file)
	result := make([]models.Metric, 0)
	if err = dec.Decode(&result); err != nil {
		return err
	}
	for _, metric := range result {
		repo.UpdateMetrics(metric)
	}
	return nil
}
