package configs

import (
	"flag"
	"fmt"
	"strings"
)

type flags struct {
	flagRunAddr    string
	reportInterval int64
	pollInterval   int64
}

func parseFlags() (flags, error) {
	addr := flag.String("a", "localhost:8080", "address and port to run server")
	report := flag.Int64("r", 10, "metrics sending interval in seconds")
	poll := flag.Int64("p", 2, "metrics update interval in seconds")
	flag.Parse()

	return flags{
		flagRunAddr:    correctAddr(addr),
		reportInterval: *report,
		pollInterval:   *poll,
	}, nil
}

func correctAddr(addr *string) string {
	before, _, found := strings.Cut(*addr, ":")
	if !found {
		return fmt.Sprintf("localhost:%s", before)
	}
	return *addr
}
