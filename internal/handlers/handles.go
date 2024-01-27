package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"

	"github.com/Gamilkarr/stattrack/internal/models"
)

func (h *Handler) UpdateMetrics(res http.ResponseWriter, req *http.Request) {

	metricType := chi.URLParam(req, "type")
	metricName := chi.URLParam(req, "name")
	metricValue := chi.URLParam(req, "value")

	var metricForUpdate models.Metric
	var result models.Metric

	switch metricType {
	case models.MetricGauge:
		value, err := strconv.ParseFloat(metricValue, 64)
		if err != nil {
			http.Error(res, "metric gauge: invalid value", http.StatusBadRequest)
			return
		}
		metricForUpdate = models.Metric{
			ID:    metricName,
			MType: metricType,
			Value: &value,
		}
		result = h.Repo.UpdateMetrics(metricForUpdate)

	case models.MetricCounter:
		value, err := strconv.ParseInt(metricValue, 10, 64)
		if err != nil {
			http.Error(res, "metric counter: invalid value", http.StatusBadRequest)
			return
		}
		metricForUpdate = models.Metric{
			ID:    metricName,
			MType: metricType,
			Delta: &value,
		}
		result = h.Repo.UpdateMetrics(metricForUpdate)
	default:
		http.Error(res, "unknown metric type", http.StatusBadRequest)
		return
	}
	resJSON, _ := json.Marshal(result)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(resJSON)
}

func (h *Handler) GetValueMetric(res http.ResponseWriter, req *http.Request) {
	metricType := chi.URLParam(req, "type")
	metricName := chi.URLParam(req, "name")
	metric := &models.Metric{
		ID:    metricName,
		MType: metricType,
	}
	var (
		value string
		err   error
	)
	switch metricType {
	case models.MetricGauge:
		if metric, err = h.Repo.GetMetricsValue(*metric); err == nil {
			value = fmt.Sprintf("%g", *metric.Value)
		}
	case models.MetricCounter:
		if metric, err = h.Repo.GetMetricsValue(*metric); err == nil {
			value = fmt.Sprintf("%d", *metric.Delta)
		}
	default:
		http.Error(res, "unknown metric type", http.StatusBadRequest)
		return
	}
	if err != nil {
		http.Error(res, "unknown metric type", http.StatusNotFound)
		return
	}
	res.WriteHeader(http.StatusOK)
	res.Write([]byte(value))
}

func (h *Handler) GetMetrics(res http.ResponseWriter, req *http.Request) {
	mes := h.Repo.GetMetrics()
	resJSON, _ := json.Marshal(mes)
	res.Header().Set("Content-Type", "application/json")
	res.Write(resJSON)
}

func (h *Handler) UpdateJSONMetrics(res http.ResponseWriter, req *http.Request) {
	metricsForUpdate := models.Metric{}
	var buf bytes.Buffer
	_, err := buf.ReadFrom(req.Body)
	if err != nil {
		http.Error(res, "something wrong", http.StatusInternalServerError)
		return
	}
	if err = json.Unmarshal(buf.Bytes(), &metricsForUpdate); err != nil {
		http.Error(res, "something wrong", http.StatusInternalServerError)
		return
	}
	result := h.Repo.UpdateMetrics(metricsForUpdate)

	resJSON, _ := json.Marshal(result)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(resJSON)
}

func (h *Handler) GetJSONValueMetric(res http.ResponseWriter, req *http.Request) {
	metrics := models.Metric{}

	var buf bytes.Buffer
	_, err := buf.ReadFrom(req.Body)
	if err != nil {
		http.Error(res, "something wrong", http.StatusInternalServerError)
		return
	}

	if err = json.Unmarshal(buf.Bytes(), &metrics); err != nil {
		http.Error(res, "something wrong", http.StatusInternalServerError)
		return
	}

	metricValue, err := h.Repo.GetMetricsValue(metrics)
	if err != nil {
		http.Error(res, "metric not found", http.StatusNotFound)
		return
	}

	resJSON, _ := json.Marshal(metricValue)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	res.Write(resJSON)
}
