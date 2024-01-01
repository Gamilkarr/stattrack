package configs

import (
	"github.com/caarlos0/env/v6"
)

type envVar struct {
	address        string `env:"ADDRESS"`
	pollInterval   int64  `env:"POLL_INTERVAL"`
	reportInterval int64  `env:"REPORT_INTERVAL"`
}

func getEnvVar() (*envVar, error) {
	eVar := new(envVar)

	if err := env.Parse(eVar); err != nil {
		return nil, err
	}
	return eVar, nil
}
