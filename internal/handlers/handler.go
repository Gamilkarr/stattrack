package handlers

import (
	"github.com/Gamilkarr/stattrack/internal/models"
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
