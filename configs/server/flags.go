package configs

import (
	"flag"

	"github.com/Gamilkarr/stattrack/configs"
)

type flags struct {
	flagRunAddr configs.NetAddress
}

func parseFlags() (flags, error) {
	addr := configs.NetAddress{
		Host: "localhost",
		Port: 8080,
	}
	_ = flag.Value(&addr)
	flag.Var(&addr, "a", "address and port to run server")
	flag.Parse()
	return flags{flagRunAddr: addr}, nil
}
