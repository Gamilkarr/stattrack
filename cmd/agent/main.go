package main

import (
	"github.com/Gamilkarr/stattrack/internal/client"
	"net/http"

	"github.com/Gamilkarr/stattrack/internal/service"
)

func main() {
	c := client.Client{
		Client:         http.Client{},
		Service:        &service.Service{},
		CounterMetrics: nil,
		GaugeMetrics:   nil,
	}

	for {
		c.GetMetrics()
		c.SendMetrics()
	}
}
