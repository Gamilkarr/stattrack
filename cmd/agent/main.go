package main

import (
	"fmt"
	"time"

	"github.com/go-resty/resty/v2"

	"github.com/Gamilkarr/stattrack/internal/service"
)

var UpdateMetricsURL = "http://%s/update/{type}/{name}/{value}"

type Agent struct {
	Client  *resty.Client
	Metrics Metrics
}

type Metrics interface {
	UpdateMetrics()
	GetAllMetrics() []map[string]string
}

func (a *Agent) SendMetrics() {
	mSlice := a.Metrics.GetAllMetrics()
	for _, metric := range mSlice {
		_, err := a.Client.R().
			SetHeader("Content-Type", "text/plain").
			SetPathParams(metric).
			Post(fmt.Sprintf(UpdateMetricsURL, flagRunAddr))
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

	parseFlags()

	for {
		time.Sleep(pollInterval)
		agent.Metrics.UpdateMetrics()

		time.Sleep(reportInterval)
		agent.SendMetrics()
	}
}
