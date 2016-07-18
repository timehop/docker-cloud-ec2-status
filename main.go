package main

import (
	"log"

	"github.com/docker/go-dockercloud/dockercloud"
	"github.com/go-errors/errors"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Println(errors.Wrap(err, 2).ErrorStack())
		}
	}()

	config, err := NewConfig()
	if err != nil {
		log.Fatalf(err.Error())
	}

	dockercloud.User = config.DockerCloudUser
	dockercloud.ApiKey = config.DockerCloudAPIKey

	poller := NewPoller(config)
	poller.Start()
}
