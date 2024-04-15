package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/Gamilkarr/stattrack/internal/models"
)

type DataBase struct {
	*sql.DB
}

func NewDataBase(db *sql.DB) (*DataBase, error) {
	createQuery := "CREATE TABLE IF NOT EXISTS metrics (name varchar(30), gauge double precision, counter int);"
	_, err := db.ExecContext(context.Background(), createQuery)
	if err != nil {
		return nil, err
	}
	return &DataBase{
		db,
	}, nil
}

func (db *DataBase) UpdateMetrics(metric models.Metric) (models.Metric, error) {
	return models.Metric{}, nil
}
func (db *DataBase) GetMetricsValue(metrics models.Metric) (*models.Metric, error) {
	return nil, nil
}
func (db *DataBase) GetMetrics() []models.Metric {
	return nil
}

func (db *DataBase) Ping() error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		return err
	}
	return nil
}
