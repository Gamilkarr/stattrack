package repository

import (
	"errors"
	"fmt"

	"github.com/Gamilkarr/stattrack/internal/models"
)

type MemStorage struct {
	Gauge        map[string]float64
	Counter      map[string]int64
	BackUPPeriod int64
	BackUPPath   string
}

func NewRepo(period int64, path string) *MemStorage {
	return &MemStorage{
		Gauge:        make(map[string]float64),
		Counter:      make(map[string]int64),
		BackUPPeriod: period,
		BackUPPath:   path,
	}
}

func (m *MemStorage) UpdateMetrics(metric models.Metric) (models.Metric, error) {
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
	if m.BackUPPeriod == 0 {
		err := m.backUP()
		if err != nil {
			return result, fmt.Errorf("storage error: %w", err)
		}
	}
	return result, nil
}

func (m *MemStorage) GetMetricsValue(metric models.Metric) (*models.Metric, error) {
	var result models.Metric
	switch metric.MType {
	case models.MetricGauge:
		value, ok := m.Gauge[metric.ID]
		if !ok {
			return nil, errors.New("storage error: gauge metric not found")
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
			return nil, errors.New("storage error: counter metric not found")
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
