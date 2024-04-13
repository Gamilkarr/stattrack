package repository

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/Gamilkarr/stattrack/internal/models"
)

func (m *MemStorage) RunBackUP() {
	for {
		time.Sleep(time.Duration(m.BackUPPeriod) * time.Second)
		if err := m.backUP(); err != nil {
			log.WithField("error", err).Error("backup error")
		}
	}
}

func (m *MemStorage) backUP() error {
	file, err := os.OpenFile(m.BackUPPath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.WithField("storage error", err).Error("file closing error")
		}
	}(file)

	enc := json.NewEncoder(file)

	if err = enc.Encode(m.GetMetrics()); err != nil {
		return err
	}
	return nil
}

func (m *MemStorage) Uploading(path string) error {
	file, err := os.OpenFile(path, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		return fmt.Errorf("storage error: %w", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.WithField("storage error", err).Error("file closing error")
		}
	}(file)
	dec := json.NewDecoder(file)
	result := make([]models.Metric, 0)
	if err = dec.Decode(&result); err != nil {
		return fmt.Errorf("storage error: %w", err)
	}
	for _, metric := range result {
		_, err := m.UpdateMetrics(metric)
		if err != nil {
			return fmt.Errorf("storage error: %w", err)
		}
	}
	return nil
}
