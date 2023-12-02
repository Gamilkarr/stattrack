package main

import (
	"net/http"

	"github.com/Gamilkarr/stattrack/internal/endpoints"
	"github.com/Gamilkarr/stattrack/internal/repository"
)

func main() {
	repo := repository.MemStorage{
		Gauge:   make(map[string]float64),
		Counter: make(map[string]int64),
	}
	handlers := endpoints.Endpoints{Repo: &repo}
	http.HandleFunc(`/update/`, handlers.UpdateMetrics)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
