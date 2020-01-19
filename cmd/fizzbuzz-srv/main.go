package main

import (
	"github.com/mtebourbi/lbc-fizzbuzz/pkg/server"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("FizzBuzz web service")
	if err := server.ListenAndServe(); err != nil {
		logrus.WithError(err).Fatal("main: failed to start http server")
	}
}
