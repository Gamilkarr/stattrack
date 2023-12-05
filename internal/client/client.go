package client

import (
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	Client         http.Client
	Service        Service
	CounterMetrics map[string]int64
	GaugeMetrics   map[string]float64
}

type Service interface {
	GetGaugeMetricsMap() map[string]float64
}

const (
	UpdateMetricsGaugeURL   = "http://localhost:8080/update/gauge/%s/%.2f"
	UpdateMetricsCounterURL = "http://localhost:8080/update/counter/%s/%d"
	pollInterval            = 2
	reportInterval          = 10
)

func (c *Client) SendMetrics() {
	time.Sleep(reportInterval * time.Second)
	for name, value := range c.GaugeMetrics {
		url := fmt.Sprintf(UpdateMetricsGaugeURL, name, value)
		resp, err := c.Client.Post(url, "text/plain", nil)
		if err != nil {
			fmt.Println(err)
		}
		resp.Body.Close()
	}

	for name, value := range c.CounterMetrics {
		url := fmt.Sprintf(UpdateMetricsCounterURL, name, value)
		resp, err := c.Client.Post(url, "text/plain", nil)
		if err != nil {
			fmt.Println(err)
		}
		resp.Body.Close()
	}
}

func (c *Client) GetMetrics() {
	time.Sleep(pollInterval * time.Second)

	c.GaugeMetrics = c.Service.GetGaugeMetricsMap()
	if c.CounterMetrics == nil {
		c.CounterMetrics = map[string]int64{"PollCount": 0}
	}
	c.CounterMetrics["PollCount"] += 1
}
