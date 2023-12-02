package endpoints

import (
	"net/http"
	"strconv"
	"strings"
)

type Endpoints struct {
	Repo repo
}

type repo interface {
	UpdateGaugeMetrics(name string, val float64) error
	UpdateCounterMetrics(name string, val int64) error
}

const (
	metricGauge   = "gauge"
	metricCounter = "counter"
)

func (e *Endpoints) UpdateMetrics(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(res, "only POST requests are allowed", http.StatusMethodNotAllowed)
		return
	}
	args := strings.Split(strings.TrimPrefix(req.URL.Path, "/update/"), "/")
	if len(args) < 3 {
		http.Error(res, "empty name", http.StatusNotFound)
		return
	}
	var (
		metricType  = args[0]
		metricName  = args[1]
		metricValue = args[2]
	)
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
