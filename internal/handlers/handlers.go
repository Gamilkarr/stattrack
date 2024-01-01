package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func (h *Handler) UpdateMetrics(res http.ResponseWriter, req *http.Request) {

	metricType := chi.URLParam(req, "type")
	metricName := chi.URLParam(req, "name")
	metricValue := chi.URLParam(req, "value")

	switch metricType {
	case metricGauge:
		value, err := strconv.ParseFloat(metricValue, 64)
		if err != nil {
			http.Error(res, "metric gauge: invalid value", http.StatusBadRequest)
			return
		}
		if err = h.Repo.UpdateGaugeMetrics(metricName, value); err != nil {
			http.Error(res, "something wrong", http.StatusInternalServerError)
			return
		}
	case metricCounter:
		value, err := strconv.ParseInt(metricValue, 10, 64)
		if err != nil {
			http.Error(res, "metric counter: invalid value", http.StatusBadRequest)
			return
		}
		if err = h.Repo.UpdateCounterMetrics(metricName, value); err != nil {
			http.Error(res, "something wrong", http.StatusInternalServerError)
			return
		}
	default:
		http.Error(res, "unknown metric type", http.StatusBadRequest)
		return
	}
	res.WriteHeader(http.StatusOK)
}

func (h *Handler) GetValueMetric(res http.ResponseWriter, req *http.Request) {
	metricType := chi.URLParam(req, "type")
	metricName := chi.URLParam(req, "name")
	var metricValue string

	switch metricType {
	case metricGauge:
		value, ok := h.Repo.GetGaugeMetricValue(metricName)
		if !ok {
			http.Error(res, "metric not found", http.StatusNotFound)
			return
		}
		metricValue = fmt.Sprintf("%g", value)
	case metricCounter:
		value, ok := h.Repo.GetCounterMetricValue(metricName)
		if !ok {
			http.Error(res, "metric not found", http.StatusNotFound)
			return
		}
		metricValue = fmt.Sprintf("%d", value)
	default:
		http.Error(res, "unknown metric type", http.StatusBadRequest)
		return
	}
	res.WriteHeader(http.StatusOK)
	res.Write([]byte(metricValue))
}

type Metrics struct {
	Counter map[string]int64   `json:"counter"`
	Gouge   map[string]float64 `json:"gouge"`
}

func (h *Handler) GetMetrics(res http.ResponseWriter, req *http.Request) {
	met := Metrics{
		Counter: h.Repo.GetCounterMetrics(),
		Gouge:   h.Repo.GetGaugeMetrics(),
	}

	metJSON, _ := json.Marshal(met)
	res.Write(metJSON)
}
