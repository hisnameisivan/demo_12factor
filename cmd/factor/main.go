package main

import (
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("hello world")

	port := os.Getenv("PORT")
	if len(port) == 0 {
		logrus.Fatal("port is empty")
	}

}
