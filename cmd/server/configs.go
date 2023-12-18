package main

type config struct {
	address string
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

	return cfg
}
