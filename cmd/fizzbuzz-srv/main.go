package main

import (
	"flag"
	"github.com/mtebourbi/lbc-fizzbuzz/pkg/server"
	"github.com/sirupsen/logrus"
)

func main() {
	var listenPort = 8080
	flag.IntVar(&listenPort, "port", 8080, "server http port")
	flag.Parse()

	logrus.Info("FizzBuzz web service")
	if err := server.ListenAndServe(listenPort); err != nil {
		logrus.WithError(err).Fatal("main: failed to start http server")
	}
}
