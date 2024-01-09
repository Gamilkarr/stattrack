package configs

import (
	"flag"

	"github.com/Gamilkarr/stattrack/configs"
)

type flags struct {
	flagRunAddr    configs.NetAddress
	reportInterval int64
	pollInterval   int64
}

func parseFlags() (flags, error) {
	addr := configs.NetAddress{
		Host: "localhost",
		Port: 8080,
	}
	_ = flag.Value(&addr)
	flag.Var(&addr, "a", "address and port to run server")
	report := flag.Int64("r", 10, "metrics sending interval in seconds")
	poll := flag.Int64("p", 2, "metrics update interval in seconds")
	flag.Parse()

	return flags{
		flagRunAddr:    addr,
		reportInterval: *report,
		pollInterval:   *poll,
	}, nil
}
