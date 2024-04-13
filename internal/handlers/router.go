package handlers

import "github.com/go-chi/chi/v5"

func (h *Handler) NewRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/", h.GetMetrics)

	r.Get("/ping", h.Ping)

	r.Post("/update/", h.UpdateJSONMetrics)
	r.Post("/update/{type}/{name}/{value}", h.UpdateMetrics)

	r.Post("/value/", h.GetJSONValueMetric)
	r.Get("/value/{type}/{name}", h.GetValueMetric)

	return r
}
