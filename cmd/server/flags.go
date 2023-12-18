package main

import (
	"errors"
	"flag"
	"fmt"
	"strconv"
	"strings"
)

type NetAddress struct {
	Host string
	Port int
}

func (n *NetAddress) String() string {
	return fmt.Sprintf("%s:%d", n.Host, n.Port)
}

func (n *NetAddress) Set(s string) error {
	host, strPort, _ := strings.Cut(s, ":")

	if host != "" {
		n.Host = host
	}

	if strPort == "" {
		return nil
	}
	port, parseErr := strconv.Atoi(strPort)
	if parseErr != nil {
		return errors.New("need address in a form host:port")
	}
	n.Port = port
	return nil
}

type flags struct {
	flagRunAddr NetAddress
}

func parseFlags() flags {
	addr := NetAddress{
		Host: "localhost",
		Port: 8080,
	}
	_ = flag.Value(&addr)
	flag.Var(&addr, "a", "address and port to run server")
	flag.Parse()
	return flags{flagRunAddr: addr}
}
