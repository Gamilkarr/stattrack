package repository

import (
	"errors"
	"github.com/Gamilkarr/stattrack/internal/models"
)

type MemStorage struct {
	Gauge   map[string]float64
	Counter map[string]int64
}

func NewRepo() (*MemStorage, error) {
	return &MemStorage{
		Gauge:   make(map[string]float64),
		Counter: make(map[string]int64),
	}, nil
}

func (m *MemStorage) UpdateMetrics(metric models.Metric) models.Metric {
	var result models.Metric
	switch metric.MType {
	case models.MetricGauge:
		m.Gauge[metric.ID] = *metric.Value
		val := m.Gauge[metric.ID]
		result = models.Metric{
			ID:    metric.ID,
			MType: models.MetricGauge,
			Delta: nil,
			Value: &val,
		}
	case models.MetricCounter:
		m.Counter[metric.ID] += *metric.Delta
		val := m.Counter[metric.ID]
		result = models.Metric{
			ID:    metric.ID,
			MType: models.MetricCounter,
			Delta: &val,
			Value: nil,
		}
	}
	return result
}

func (m *MemStorage) GetMetricsValue(metric models.Metric) (*models.Metric, error) {
	var result models.Metric
	switch metric.MType {
	case models.MetricGauge:
		value, ok := m.Gauge[metric.ID]
		if !ok {
			return nil, errors.New("metric not found")
		}
		result = models.Metric{
			ID:    metric.ID,
			MType: models.MetricGauge,
			Delta: nil,
			Value: &value,
		}
	case models.MetricCounter:
		value, ok := m.Counter[metric.ID]
		if !ok {
			return nil, errors.New("metric not found")
		}
		result = models.Metric{
			ID:    metric.ID,
			MType: models.MetricCounter,
			Delta: &value,
			Value: nil,
		}
	}
	return &result, nil
}

func (m *MemStorage) GetMetrics() []models.Metric {
	result := make([]models.Metric, 0)
	for key, value := range m.Gauge {
		value := value
		result = append(result, models.Metric{
			ID:    key,
			MType: models.MetricGauge,
			Delta: nil,
			Value: &value,
		})
	}
	for key, value := range m.Counter {
		value := value
		result = append(result, models.Metric{
			ID:    key,
			MType: models.MetricCounter,
			Delta: &value,
			Value: nil,
		})
	}
	return result
}
