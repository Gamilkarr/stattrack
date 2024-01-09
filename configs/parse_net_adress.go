package configs

import (
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
	host, port, err := parseHost(s)
	if err != nil {
		return err
	}

	if host != "" {
		n.Host = host
	}

	n.Port = port
	return nil
}

func parseHost(s string) (string, int, error) {
	before, after, found := strings.Cut(s, ":")
	if !found {
		port, err := strconv.Atoi(before)
		return "", port, err
	}
	port, err := strconv.Atoi(after)
	return before, port, err
}
