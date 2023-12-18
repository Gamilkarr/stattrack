package main

import (
	"fmt"
	"time"
)

type config struct {
	address        string
	pollInterval   time.Duration
	reportInterval time.Duration
}

func newConfig() *config {
	cfg := new(config)

	eVar := getEnvVar()
	flag := parseFlags()

	if eVar.Address != "" {
		cfg.address = eVar.Address
	} else {
		cfg.address = flag.flagRunAddr.String()
	}

	if eVar.PoolInterval != 0 {
		cfg.pollInterval = time.Duration(eVar.PoolInterval) * time.Second
	} else {
		cfg.pollInterval = time.Duration(flag.pollInterval) * time.Second
	}

	if eVar.ReportInterval != 0 {
		cfg.pollInterval = time.Duration(eVar.ReportInterval) * time.Second
	} else {
		cfg.pollInterval = time.Duration(flag.reportInterval) * time.Second
	}

	return cfg
}

func getUpdateURL(addr string) string {
	return fmt.Sprintf("http://%s/update/{type}/{name}/{value}", addr)
}
