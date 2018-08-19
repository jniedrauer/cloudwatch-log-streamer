package main

import (
	"github.com/jniedrauer/cloudwatch-log-streamer/pkg/config"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init() {
	err := config.SetLogLevel(log)
	if err != nil {
		log.Error(err)
	}
}

func main() {
	log.Info("Hello world!")
}
