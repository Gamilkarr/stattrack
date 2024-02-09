package configs

import (
	"github.com/caarlos0/env/v6"
)

type envVar struct {
	Address         string `env:"ADDRESS"`
	FileStoragePath string `env:"FILE_STORAGE_PATH"`
	StoreInterval   int64  `env:"STORE_INTERVAL"`
	Restore         string `env:"RESTORE"`
}

func getEnvVar() (*envVar, error) {
	eVar := new(envVar)

	if err := env.Parse(eVar); err != nil {
		return nil, err
	}
	return eVar, nil
}
