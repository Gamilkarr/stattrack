package main

import (
	"github.com/caarlos0/env/v6"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"

	"github.com/Gamilkarr/stattrack/internal/endpoints"
	"github.com/Gamilkarr/stattrack/internal/repository"
)

type Config struct {
	Address string `env:"ADDRESS"`
}

func main() {
	e := &endpoints.Endpoints{
		Repo: &repository.MemStorage{
			Gauge:   make(map[string]float64),
			Counter: make(map[string]int64),
		},
	}
	var cfg Config

	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal(err)
	}
	parseFlags()
	if cfg.Address != "" {
		flagRunAddr = cfg.Address
	}
	r := chi.NewRouter()
	r.Post("/update/{type}/{name}/{value}", e.UpdateMetrics)
	r.Get("/value/{type}/{name}", e.GetValueMetric)
	r.Get("/", e.GetMetrics)
	log.Fatal(http.ListenAndServe(flagRunAddr, r))
}
