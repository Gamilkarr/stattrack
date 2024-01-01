package configs

import (
	"errors"
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
		return errors.New("need address in a form host:port")
	}
	n.port = port
	return nil
}

type flags struct {
	flagRunAddr    netAddress
	reportInterval int64
	pollInterval   int64
}

func parseFlags() (flags, error) {
	addr := netAddress{
		host: "localhost",
		port: 8080,
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
