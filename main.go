package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-errors/errors"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Println(errors.Wrap(err, 2).ErrorStack())
		}
	}()

	appEnv := appEnv()
	dcAPIKey := dcAPIKey()

	log.Printf("Running on env: '%s'", appEnv)
	log.Printf("Docker Cloud API key starts with: '%s...'", string(dcAPIKey[0:4]))
	log.Print("Service starting...")

	setupLogging(appEnv)

	for {

	}
}

func appEnv() string {
	if os.Getenv("APP_ENV") == "" {
		return "development"
	}
	return os.Getenv("APP_ENV")
}

func dcAPIKey() string {
	if os.Getenv("DC_API_KEY") == "" {
		panic(
			errors.New("'DC_API_KEY' env variable does not exist"),
		)
	}
	return os.Getenv("DC_API_KEY")
}

func setupLogging(appEnv string) {
	if appEnv == "development" {
		return
	}

	location := "/var/log/docker-cloud-ec2-status.log"

	fileHandler, err := os.OpenFile(
		location, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666,
	)
	if err != nil {
		fmt.Printf("Error opening log file: %v", err)
	}

	defer fileHandler.Close()
	log.SetOutput(fileHandler)
}
