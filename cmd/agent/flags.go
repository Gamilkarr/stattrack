package main

import (
	"errors"
	"flag"
	"fmt"
	"strconv"
	"strings"
	"time"
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

var (
	flagRunAddr    string
	reportInterval time.Duration
	pollInterval   time.Duration
)

func parseFlags() {
	addr := NetAddress{
		Host: "localhost",
		Port: 8080,
	}
	_ = flag.Value(&addr)
	flag.Var(&addr, "a", "address and port to run server")
	report := flag.Int64("r", 10, "metrics sending interval in seconds")
	poll := flag.Int64("p", 2, "metrics update interval in seconds")
	flag.Parse()

	flagRunAddr = addr.String()

	reportInterval = time.Duration(*report) * time.Second
	pollInterval = time.Duration(*poll) * time.Second
}
