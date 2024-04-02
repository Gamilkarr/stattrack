package configs

type Config struct {
	Address         string
	FileStoragePath string
	StoreInterval   int64
	Restore         bool
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
		Address:         eVar.Address,
		FileStoragePath: eVar.FileStoragePath,
		StoreInterval:   eVar.StoreInterval,
	}

	switch eVar.Restore {
	case "true":
		cfg.Restore = true
	case "false":
		cfg.Restore = false
	case "":
		cfg.Restore = flag.restore
	}

	if cfg.Address == "" {
		cfg.Address = flag.flagRunAddr
	}

	if cfg.FileStoragePath == "" {
		cfg.FileStoragePath = flag.fileStoragePath
	}

	if eVar.StoreInterval == -1 {
		cfg.StoreInterval = flag.storeInterval
	}

	return cfg, nil
}
