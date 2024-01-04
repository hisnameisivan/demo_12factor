package main

import (
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("hello world")

	port := os.Getenv("PORT")
	if len(port) == 0 {
		logrus.Fatal("port is empty")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	addr := ":" + port
	http.ListenAndServe(addr, nil)
}
