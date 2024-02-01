package configs

import (
	"flag"

	"github.com/Gamilkarr/stattrack/configs"
)

type flags struct {
	flagRunAddr     configs.NetAddress
	fileStoragePath string
	storeInterval   int64
	restore         bool
}

func parseFlags() (flags, error) {
	addr := configs.NetAddress{
		Host: "localhost",
		Port: 8080,
	}
	_ = flag.Value(&addr)
	flag.Var(&addr, "a", "address and port to run server")

	fileStoragePath := flag.String("f", "/Users/takhabarova/learningProjects/stattrack/tmp/metrics-db.json", "path for saving server data to disk")
	storeInterval := flag.Int64("i", 300, "time interval for saving server readings to disk")
	restore := flag.Bool("r", true, "Is load previously saved values")

	flag.Parse()
	return flags{
		flagRunAddr:     addr,
		storeInterval:   *storeInterval,
		fileStoragePath: *fileStoragePath,
		restore:         *restore,
	}, nil
}
