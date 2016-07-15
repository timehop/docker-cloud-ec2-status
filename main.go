package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-errors/errors"
)

const logLocation string = "/var/log/docker-cloud-ec2-status.log"

type ()

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Println(errors.Wrap(err, 2).ErrorStack())
		}
	}()

	setupLogging()

	config, err := NewConfig()
	if err != nil {
		log.Fatalf(err.Error())
	}

	poller := NewPoller(config)
	poller.Start()
}

func setupLogging() {
	fileHandler, err := os.OpenFile(
		logLocation, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666,
	)
	if err != nil {
		fmt.Printf("Error opening log file: %v", err)
	}

	defer fileHandler.Close()
	log.SetOutput(fileHandler)
}
