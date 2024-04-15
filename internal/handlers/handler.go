package handlers

import (
	"github.com/Gamilkarr/stattrack/internal/models"
)

type Handler struct {
	Repo Repository
}

type Repository interface {
	UpdateMetrics(metric models.Metric) (models.Metric, error)
	GetMetricsValue(metrics models.Metric) (*models.Metric, error)
	GetMetrics() []models.Metric
	Ping() error
}

func NewHandler(repo Repository) *Handler {
	return &Handler{
		Repo: repo,
	}
}
