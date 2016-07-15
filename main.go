package main

import (
	"log"

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

	poller := NewPoller(config)
	poller.Start()
}
