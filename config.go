package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
)

const dockerCloudConfigLocation string = "/etc/dockercloud/agent/dockercloud-agent.conf"

type (
	// Config comment pending
	Config struct {
		AppEnv              string
		DockerCloudAPIKey   string
		DockerCloudNodeUUID string
	}

	// DockerCloudConfig comment pending
	DockerCloudConfig struct {
		UUID string `json:"UUID"`
	}
)

// NewConfig comment pending
func NewConfig() (*Config, error) {
	var config Config

	appEnv, err := appEnv()
	if err != nil {
		return nil, err
	}
	log.Printf("Running on env: '%s'", appEnv)

	dockerCloudAPIKey, err := dockerCloudAPIKey()
	if err != nil {
		return nil, err
	}
	log.Printf(
		"Docker Cloud API key starts with: '%s...'",
		string(dockerCloudAPIKey[0:4]),
	)

	dockerCloudNodeUUID, err := dockerCloudNodeUUID(appEnv)
	if err != nil {
		errorMsg := fmt.Sprintf("Error when reading UUID from Docker Cloud config file: %s", err)
		return nil, errors.New(errorMsg)
	}
	log.Printf("Docker Cloud node UUID: '%s'", dockerCloudNodeUUID)

	config.AppEnv = appEnv
	config.DockerCloudAPIKey = dockerCloudAPIKey
	config.DockerCloudNodeUUID = dockerCloudNodeUUID

	return &config, nil
}

func appEnv() (string, error) {
	if os.Getenv("APP_ENV") == "" {
		return "", errors.New("'APP_ENV' env variable does not exist")
	}
	return os.Getenv("APP_ENV"), nil
}

func dockerCloudAPIKey() (string, error) {
	if os.Getenv("DC_API_KEY") == "" {
		return "", errors.New("'DC_API_KEY' env variable does not exist")
	}
	return os.Getenv("DC_API_KEY"), nil
}

func dockerCloudNodeUUID(appEnv string) (string, error) {
	fileHandler, err := os.Open(dockerCloudConfigLocation)
	if err != nil {
		return "", err
	}

	var config DockerCloudConfig

	jsonParser := json.NewDecoder(fileHandler)
	if err = jsonParser.Decode(&config); err != nil {
		return "", err
	}

	return config.UUID, nil
}
