package main

import (
	"log"
	"net"
	"strings"
	"time"
)

func check(host *Host, timeout *time.Duration) Check {
	var state State

	conn, err := net.DialTimeout(host.Network, host.Address, *timeout)
	if err != nil {
		if strings.Contains(err.Error(), "timeout") {
			state = closed
		} else {
			log.Printf("failure when trying to access host %s:%s: %s", host.Network, host.Address, err.Error())
			state = failure
		}
	} else {
		defer conn.Close()
		state = open
	}

	return Check{host: *host, state: state}
}

func scan(config *Config) ([]Check, error) {
	var res []Check

	for _, host := range config.Hosts {
		res = append(res, check(&host, &config.Timeout))
	}

	return res, nil
}
