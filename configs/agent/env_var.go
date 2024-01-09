package configs

import (
	"github.com/caarlos0/env/v6"
)

type envVar struct {
	Address        string `env:"ADDRESS"`
	PollInterval   int64  `env:"POLL_INTERVAL"`
	ReportInterval int64  `env:"REPORT_INTERVAL"`
}

func getEnvVar() (*envVar, error) {
	eVar := new(envVar)

	if err := env.Parse(eVar); err != nil {
		return nil, err
	}
	return eVar, nil
}
