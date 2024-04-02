package main

import (
	"time"

	config "github.com/Gamilkarr/stattrack/configs/agent"
	"github.com/Gamilkarr/stattrack/internal/client"
	log "github.com/sirupsen/logrus"
)

func main() {
	agent, err := client.NewAgent()
	if err != nil {
		log.WithField("err", err).Fatal("client error")
	}

	cfg, err := config.NewConfig()
	if err != nil {
		log.WithField("err", err).Fatal("config error")
	}

	for {
		time.Sleep(cfg.PollInterval)
		agent.UpdateMetrics()

		time.Sleep(cfg.ReportInterval)
		err := agent.SendMetrics(cfg.Address)
		if err != nil {
			log.WithField("err", err).Error("send metrics error")
		}
	}
}
