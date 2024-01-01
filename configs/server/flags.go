package configs

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

type netAddress struct {
	host string
	port int
}

func (n *netAddress) String() string {
	return fmt.Sprintf("%s:%d", n.host, n.port)
}

func (n *netAddress) Set(s string) error {
	host, strPort, _ := strings.Cut(s, ":")

	if host != "" {
		n.host = host
	}

	if strPort == "" {
		return nil
	}
	port, parseErr := strconv.Atoi(strPort)
	if parseErr != nil {
		return parseErr
	}
	n.port = port
	return nil
}

type flags struct {
	flagRunAddr netAddress
}

func parseFlags() (flags, error) {
	addr := netAddress{
		host: "localhost",
		port: 8080,
	}
	_ = flag.Value(&addr)
	flag.Var(&addr, "a", "address and port to run server")
	flag.Parse()
	return flags{flagRunAddr: addr}, nil
}
