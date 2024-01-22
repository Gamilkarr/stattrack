package client

import (
	"encoding/json"
	"fmt"
	"log"

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

func (a *Agent) SendMetrics(adr string) {
	url := fmt.Sprintf("http://%s/update/", adr)
	mSlice := a.GetAllMetrics()
	for _, metric := range mSlice {
		msg, _ := json.Marshal(metric)
		_, err := a.Client.R().
			SetHeader("Content-Type", "application/json").
			SetBody(msg).
			Post(url)
		if err != nil {
			log.Println(err)
		}
	}
}
