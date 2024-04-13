package client

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"

	"github.com/Gamilkarr/stattrack/internal/models"
	"github.com/Gamilkarr/stattrack/internal/service"
	"github.com/go-resty/resty/v2"
)

type Agent struct {
	Client *resty.Client
	Metrics
}

type Metrics interface {
	UpdateMetrics()
	GetAllMetrics() []models.Metric
}

func NewAgent() (*Agent, error) {
	return &Agent{
		Client:  resty.New(),
		Metrics: new(service.Metrics),
	}, nil
}

func (a *Agent) SendMetrics(adr string) error {
	url := fmt.Sprintf("http://%s/update/", adr)
	mSlice := a.GetAllMetrics()
	for _, metric := range mSlice {
		msg, err := json.Marshal(metric)
		if err != nil {
			return fmt.Errorf("sending error: %w", err)
		}
		cMsg, err := Compress(msg)
		if err != nil {
			return fmt.Errorf("sending error: %w", err)
		}
		_, err = a.Client.R().
			SetHeader("Content-Type", "application/json").
			SetHeader("Content-Encoding", "gzip").
			SetHeader("Accept-Encoding", "gzip").
			SetBody(cMsg).
			Post(url)
		if err != nil {
			return fmt.Errorf("sending error: %w", err)
		}
	}
	return nil
}

// Compress сжимает слайс байт.
func Compress(data []byte) ([]byte, error) {
	buf := new(bytes.Buffer)
	zb := gzip.NewWriter(buf)
	_, err := zb.Write(data)
	if err != nil {
		return nil, err
	}
	err = zb.Close()
	return buf.Bytes(), err
}
