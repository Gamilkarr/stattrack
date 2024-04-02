package repository

import (
	"encoding/json"
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/Gamilkarr/stattrack/internal/models"
)

func (m *MemStorage) RunBackUP() {
	for {
		time.Sleep(time.Duration(m.BackUPPeriod) * time.Second)
		m.backUP()
	}
}

func (m *MemStorage) backUP() {
	file, err := os.OpenFile(m.BackUPPath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.WithField("backup error", err)
	}
	defer file.Close()

	enc := json.NewEncoder(file)

	if err = enc.Encode(m.GetMetrics()); err != nil {
		log.WithField("backup error", err)
	}
}

func (m *MemStorage) Uploading(path string) error {
	file, err := os.OpenFile(path, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	dec := json.NewDecoder(file)
	result := make([]models.Metric, 0)
	if err = dec.Decode(&result); err != nil {
		return err
	}
	for _, metric := range result {
		m.UpdateMetrics(metric)
	}
	return nil
}
