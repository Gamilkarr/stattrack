package configs

import (
	"time"
)

type Config struct {
	Address        string
	PollInterval   time.Duration
	ReportInterval time.Duration
}

func NewConfig() (*Config, error) {
	eVar, err := getEnvVar()
	if err != nil {
		return nil, err
	}
	flag, err := parseFlags()
	if err != nil {
		return nil, err
	}

	cfg := &Config{
		Address:        eVar.address,
		PollInterval:   time.Duration(eVar.pollInterval) * time.Second,
		ReportInterval: time.Duration(eVar.reportInterval) * time.Second,
	}

	if cfg.Address == "" {
		cfg.Address = flag.flagRunAddr.String()
	}

	if cfg.PollInterval == 0 {
		cfg.PollInterval = time.Duration(flag.pollInterval) * time.Second
	}

	if cfg.ReportInterval == 0 {
		cfg.ReportInterval = time.Duration(flag.reportInterval) * time.Second
	}

	return cfg, nil
}
