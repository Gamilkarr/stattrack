package endpoints

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type Endpoints struct {
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

func (e *Endpoints) UpdateMetrics(res http.ResponseWriter, req *http.Request) {

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
		err = e.Repo.UpdateGaugeMetrics(metricName, value)
		if err != nil {
			http.Error(res, "something wrong", http.StatusInternalServerError)
			return
		}
	case metricCounter:
		value, err := strconv.ParseInt(metricValue, 10, 64)
		if err != nil {
			http.Error(res, "metric counter: invalid value", http.StatusBadRequest)
			return
		}
		err = e.Repo.UpdateCounterMetrics(metricName, value)
		if err != nil {
			http.Error(res, "something wrong", http.StatusInternalServerError)
			return
		}
	default:
		http.Error(res, "unknown metric type", http.StatusBadRequest)
		return
	}
	res.WriteHeader(http.StatusOK)
}

func (e *Endpoints) GetValueMetric(res http.ResponseWriter, req *http.Request) {
	metricType := chi.URLParam(req, "type")
	metricName := chi.URLParam(req, "name")
	var metricValue string

	switch metricType {
	case metricGauge:
		value, ok := e.Repo.GetGaugeMetricValue(metricName)
		if !ok {
			http.Error(res, "metric not found", http.StatusNotFound)
			return
		}
		metricValue = fmt.Sprintf("%g", value)
	case metricCounter:
		value, ok := e.Repo.GetCounterMetricValue(metricName)
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

func (e *Endpoints) GetMetrics(res http.ResponseWriter, req *http.Request) {
	met := Metrics{
		Counter: e.Repo.GetCounterMetrics(),
		Gouge:   e.Repo.GetGaugeMetrics(),
	}

	metJSON, _ := json.Marshal(met)
	res.Write(metJSON)
}
