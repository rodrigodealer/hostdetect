package main

import "time"

type State int

const (
	open State = iota
	closed
	failure
)

type Check struct {
	host  Host
	state State
}

type Host struct {
	Network string `yaml:"network"`
	Address string `yaml:"address"`
}

type Config struct {
	Timeout  time.Duration `yaml:"timeout"`
	Interval time.Duration `yaml:"interval"`
	Hosts    []Host        `yaml:"hosts"`
}
