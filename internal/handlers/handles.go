package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	log "github.com/sirupsen/logrus"

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
	default:
		http.Error(res, "unknown metric type", http.StatusBadRequest)
		return
	}

	result, err := h.Repo.UpdateMetrics(metricForUpdate)
	if err != nil {
		log.WithField("error", err).Error("handler error")
	}

	resJSON, err := json.Marshal(result)
	if err != nil {
		log.WithField("error", err).Error("handler error")
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)

	_, err = res.Write(resJSON)
	if err != nil {
		log.WithField("error", err).Error("handler error")
	}
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

	_, err = res.Write([]byte(value))
	if err != nil {
		log.WithField("error", err).Error("handler error")
	}
}

func (h *Handler) GetMetrics(res http.ResponseWriter, req *http.Request) {
	mes := h.Repo.GetMetrics()
	resJSON, _ := json.Marshal(mes)
	res.Header().Set("Content-Type", "text/html")
	_, err := res.Write(resJSON)
	if err != nil {
		log.WithField("error", err).Error("handler error")
	}
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

	result, err := h.Repo.UpdateMetrics(metricsForUpdate)
	if err != nil {
		log.WithField("error", err).Error("handler error")
	}

	resJSON, err := json.Marshal(result)
	if err != nil {
		log.WithField("error", err).Error("handler error")
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	_, err = res.Write(resJSON)
	if err != nil {
		log.WithField("error", err).Error("handler error")
	}
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

	resJSON, err := json.Marshal(metricValue)
	if err != nil {
		log.WithField("error", err).Error("handler error")
	}
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	_, err = res.Write(resJSON)
	if err != nil {
		log.WithField("error", err).Error("handler error")
	}
}

func (h *Handler) Ping(res http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := h.db.PingContext(ctx); err != nil {
		http.Error(res, "ping error", http.StatusInternalServerError)
		log.WithField("error", err).Error("ping error")
	}
	res.WriteHeader(http.StatusOK)
}
