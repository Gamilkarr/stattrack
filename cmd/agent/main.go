package main

import (
	"fmt"
	"log"
	"time"

	"github.com/caarlos0/env/v6"
	"github.com/go-resty/resty/v2"

	"github.com/Gamilkarr/stattrack/internal/service"
)

var UpdateMetricsURL = "http://%s/update/{type}/{name}/{value}"

type Config struct {
	Address        string `env:"ADDRESS"`
	PoolInterval   int64  `env:"POLL_INTERVAL"`
	ReportInterval int64  `env:"REPORT_INTERVAL"`
}

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

	var cfg Config

	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	parseFlags()

	if cfg.Address != "" {
		flagRunAddr = cfg.Address
	}
	if cfg.PoolInterval != 0 {
		pollInterval = time.Duration(cfg.PoolInterval) * time.Second
	}
	if cfg.ReportInterval != 0 {
		reportInterval = time.Duration(cfg.ReportInterval) * time.Second
	}

	for {
		time.Sleep(pollInterval)
		agent.Metrics.UpdateMetrics()

		time.Sleep(reportInterval)
		agent.SendMetrics()
	}
}
