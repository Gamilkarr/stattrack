package handlers

import (
	"github.com/Gamilkarr/stattrack/internal/models"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	Repo repo
}

type repo interface {
	UpdateMetrics(metric models.Metric) models.Metric
	GetMetricsValue(metrics models.Metric) (*models.Metric, error)
	GetMetrics() []models.Metric
}

func NewHandler(repo repo) (*Handler, error) {
	return &Handler{Repo: repo}, nil
}

func (h *Handler) NewRoute() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", h.GetMetrics)

	r.Post("/update/", h.UpdateJSONMetrics)
	r.Post("/update/{type}/{name}/{value}", h.UpdateMetrics)

	r.Post("/value/", h.GetJSONValueMetric)
	r.Get("/value/{type}/{name}", h.GetValueMetric)

	return r
}
