package main

import (
	"github.com/caarlos0/env/v6"
	"log"
)

type envVar struct {
	Address string `env:"ADDRESS"`
}

func getEnvVar() *envVar {
	eVar := new(envVar)

	err := env.Parse(eVar)
	if err != nil {
		log.Fatal(err)
	}
	return eVar
}
