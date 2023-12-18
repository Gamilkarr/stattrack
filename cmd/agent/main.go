package main

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"

	"github.com/Gamilkarr/stattrack/internal/service"
)

type Agent struct {
	Client *resty.Client
	Metrics
}

type Metrics interface {
	UpdateMetrics()
	GetAllMetrics() []map[string]string
}

func (a *Agent) SendMetrics(url string) {
	mSlice := a.GetAllMetrics()
	for _, metric := range mSlice {
		_, err := a.Client.R().
			SetHeader("Content-Type", "text/plain").
			SetPathParams(metric).
			Post(url)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func main() {
	agent := Agent{
		Client:  resty.New(),
		Metrics: new(service.Metrics),
	}

	cfg := newConfig()

	for {
		time.Sleep(cfg.pollInterval)
		agent.UpdateMetrics()

		time.Sleep(cfg.reportInterval)
		agent.SendMetrics(getUpdateURL(cfg.address))
	}
}
