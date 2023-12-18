package main

import (
	"github.com/caarlos0/env/v6"
	"log"
)

type envVar struct {
	Address        string `env:"ADDRESS"`
	PoolInterval   int64  `env:"POLL_INTERVAL"`
	ReportInterval int64  `env:"REPORT_INTERVAL"`
}

func getEnvVar() *envVar {
	eVar := new(envVar)

	err := env.Parse(eVar)
	if err != nil {
		log.Fatal(err)
	}
	return eVar
}
