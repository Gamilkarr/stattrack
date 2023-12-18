package main

import (
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"

	"github.com/Gamilkarr/stattrack/internal/endpoints"
	"github.com/Gamilkarr/stattrack/internal/repository"
)

func main() {
	e := &endpoints.Endpoints{
		Repo: &repository.MemStorage{
			Gauge:   make(map[string]float64),
			Counter: make(map[string]int64),
		},
	}

	cfg := newConfig()

	r := chi.NewRouter()
	r.Post("/update/{type}/{name}/{value}", e.UpdateMetrics)
	r.Get("/value/{type}/{name}", e.GetValueMetric)
	r.Get("/", e.GetMetrics)
	log.Fatal(http.ListenAndServe(cfg.address, r))
}
