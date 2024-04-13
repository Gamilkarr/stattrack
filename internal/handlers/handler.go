package handlers

import (
	"database/sql"

	"github.com/Gamilkarr/stattrack/internal/models"
)

type Handler struct {
	Repo repo
	db   *sql.DB
}

type repo interface {
	UpdateMetrics(metric models.Metric) (models.Metric, error)
	GetMetricsValue(metrics models.Metric) (*models.Metric, error)
	GetMetrics() []models.Metric
}

func NewHandler(repo repo, db *sql.DB) *Handler {
	return &Handler{
		Repo: repo,
		db:   db,
	}
}
