package handlers

import "github.com/go-chi/chi/v5"

type Handler struct {
	Repo repo
}

type repo interface {
	UpdateGaugeMetrics(name string, val float64) error
	UpdateCounterMetrics(name string, val int64) error

	GetGaugeMetricValue(name string) (float64, bool)
	GetCounterMetricValue(name string) (int64, bool)

	GetCounterMetrics() map[string]int64
	GetGaugeMetrics() map[string]float64
}

const (
	metricGauge   = "gauge"
	metricCounter = "counter"
)

func NewHandler(repo repo) (*Handler, error) {
	return &Handler{Repo: repo}, nil
}

func (h *Handler) NewRoute() *chi.Mux {
	r := chi.NewRouter()
	r.Post("/update/{type}/{name}/{value}", h.UpdateMetrics)
	r.Get("/value/{type}/{name}", h.GetValueMetric)
	r.Get("/", h.GetMetrics)
	return r
}
