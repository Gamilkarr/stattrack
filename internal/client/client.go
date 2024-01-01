package client

import (
	"fmt"
	"github.com/Gamilkarr/stattrack/internal/service"
	"github.com/go-resty/resty/v2"
	"log"
)

type Agent struct {
	Client *resty.Client
	Metrics
}

type Metrics interface {
	UpdateMetrics()
	GetAllMetrics() []map[string]string
}

func NewAgent() (*Agent, error) {
	return &Agent{
		Client:  resty.New(),
		Metrics: new(service.Metrics),
	}, nil
}

func (a *Agent) SendMetrics(adr string) {
	url := fmt.Sprintf("http://%s/update/{type}/{name}/{value}", adr)
	mSlice := a.GetAllMetrics()
	for _, metric := range mSlice {
		_, err := a.Client.R().
			SetHeader("Content-Type", "text/plain").
			SetPathParams(metric).
			Post(url)
		if err != nil {
			log.Println(err)
		}
	}
}
