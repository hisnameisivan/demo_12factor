package main

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	log.SetOutput(os.Stdout)

	log.Info("start app")

	port := os.Getenv("PORT")
	if len(port) == 0 {
		log.Fatal("port is empty")
	}

	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	serv := http.Server{
		Addr:    net.JoinHostPort("", port),
		Handler: router,
	}

	go func() {
		serv.ListenAndServe()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM) // kill -TERM <pid>

	<-quit
	log.Info("stopping app")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := serv.Shutdown(ctx)
	if err != nil {
		log.Errorf("error when shuddown app : %v", err)
	}

	log.Info("exit")
}
