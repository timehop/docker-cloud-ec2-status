package main

import (
	"log"
	"time"

	"github.com/docker/go-dockercloud/dockercloud"
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
		if !p.nodeHealth() {
			log.Print("Node is unhealthy...")
		}

		log.Print(".")

		time.Sleep(15 * time.Second)
	}
}

func (p Poller) nodeHealth() bool {
	for i := 0; i < 3; i++ {
		node, err := dockercloud.GetNode(p.config.DockerCloudNodeUUID)

		if err != nil {
			continue
		}

		s := node.State

		if s == "Deploying" || s == "Deployed" || s == "Upgrading" {
			return true
		}
	}

	return false
}
