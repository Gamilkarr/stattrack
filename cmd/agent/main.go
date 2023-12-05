package main

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"

	"github.com/Gamilkarr/stattrack/internal/service"
)

const (
	UpdateMetricsURL = "http://localhost:8080/update/{type}/{name}/{value}"
)

type Agent struct {
	Client  *resty.Client
	Metrics Metrics
}

type Metrics interface {
	UpdateMetrics()
	GetAllMetrics() []map[string]string
}

func (a *Agent) SendMetrics() {
	time.Sleep(10 * time.Second)
	mSlice := a.Metrics.GetAllMetrics()
	for _, metric := range mSlice {
		_, err := a.Client.R().
			SetHeader("Content-Type", "text/plain").
			SetPathParams(metric).
			Post(UpdateMetricsURL)
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

	for {
		agent.Metrics.UpdateMetrics()
		agent.SendMetrics()
	}
}
