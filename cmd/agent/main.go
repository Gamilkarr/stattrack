package main

import (
	config "github.com/Gamilkarr/stattrack/configs/agent"
	"github.com/Gamilkarr/stattrack/internal/client"
	"log"
	"time"
)

func main() {
	agent, err := client.NewAgent()
	if err != nil {
		log.Fatal(err)
	}

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	for {
		time.Sleep(cfg.PollInterval)
		agent.UpdateMetrics()

		time.Sleep(cfg.ReportInterval)
		agent.SendMetrics(cfg.Address)
	}
}
