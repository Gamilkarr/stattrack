package configs

import "time"

type Config struct {
	Address         string
	FileStoragePath string
	StoreInterval   time.Duration
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
		StoreInterval:   time.Duration(eVar.StoreInterval) * time.Second,
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
		cfg.Address = flag.flagRunAddr.String()
	}

	if cfg.FileStoragePath == "" {
		cfg.FileStoragePath = flag.fileStoragePath
	}

	if cfg.StoreInterval == 0 {
		cfg.StoreInterval = time.Duration(flag.storeInterval) * time.Second
	}

	return cfg, nil
}
