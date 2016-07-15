package main

import (
	"log"
	"time"
)

type (
	// Poller comment pending
	Poller struct {
		config *Config
	}
)

// NewPoller comment pending
func NewPoller(config *Config) Poller {
	return Poller{
		config: config,
	}
}

// Start comment pending
func (p Poller) Start() {
	log.Print("Service starting polling...")

	for {
		log.Print(".")
		time.Sleep(1 * time.Second)
	}
}
