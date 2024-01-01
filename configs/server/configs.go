package configs

type Config struct {
	Address string
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
		Address: eVar.Address,
	}

	if cfg.Address == "" {
		cfg.Address = flag.flagRunAddr.String()
	}
	return cfg, nil
}
