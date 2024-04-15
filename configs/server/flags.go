package configs

import (
	"flag"
	"fmt"
	"strings"
)

type flags struct {
	flagRunAddr     string
	fileStoragePath string
	storeInterval   int64
	restore         bool
	databaseDSN     string
}

func parseFlags() (flags, error) {
	addr := flag.String("a", "localhost:8080", "address and port to run server")
	fileStoragePath := flag.String(
		"f",
		"tmp/metrics-db.json",
		"path for saving server data to disk",
	)
	storeInterval := flag.Int64("i", 300, "time interval for saving server readings to disk")
	restore := flag.Bool("r", true, "Is load previously saved values")
	databaseDSN := flag.String(
		"d",
		"",
		"database connection address",
	)

	flag.Parse()
	return flags{
		flagRunAddr:     correctAddr(addr),
		storeInterval:   *storeInterval,
		fileStoragePath: *fileStoragePath,
		restore:         *restore,
		databaseDSN:     *databaseDSN,
	}, nil
}

func correctAddr(addr *string) string {
	before, _, found := strings.Cut(*addr, ":")
	if !found {
		return fmt.Sprintf("localhost:%s", before)
	}
	return *addr
}
