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
		DockerCloudUser     string
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

	dockerCloudUser, err := fetchEnvVar("DC_USER")
	if err != nil {
		return nil, err
	}
	log.Printf("Docker Cloud user: '%s'", dockerCloudUser)

	dockerCloudAPIKey, err := fetchEnvVar("DC_API_KEY")
	if err != nil {
		return nil, err
	}
	log.Printf(
		"Docker Cloud API key starts with: '%s...'",
		string(dockerCloudAPIKey[0:4]),
	)

	dockerCloudNodeUUID, err := dockerCloudNodeUUID()
	if err != nil {
		errorMsg := fmt.Sprintf("Error when reading UUID from Docker Cloud config file: %s", err)
		return nil, errors.New(errorMsg)
	}
	log.Printf("Docker Cloud node UUID: '%s'", dockerCloudNodeUUID)

	config.DockerCloudUser = dockerCloudUser
	config.DockerCloudAPIKey = dockerCloudAPIKey
	config.DockerCloudNodeUUID = dockerCloudNodeUUID

	return &config, nil
}

func dockerCloudNodeUUID() (string, error) {
	if os.Getenv("DC_NODE_UUID") != "" {
		return os.Getenv("DC_NODE_UUID"), nil
	}

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

func fetchEnvVar(key string) (string, error) {
	if os.Getenv(key) == "" {
		errorMsg := fmt.Sprintf("'%s' env variable does not exist", key)
		return "", errors.New(errorMsg)
	}
	return os.Getenv(key), nil
}
