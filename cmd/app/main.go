package main

import (
	"context"
	"flag"
	"fmt"
	"home/pkg/routers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	var wait time.Duration
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()

	server := &http.Server{
		Addr:         "127.0.0.1:8000",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      routers.CreateRouter(),
	}

	fmt.Println("REST Server")

	go func() {
		for {
			if err := server.ListenAndServe(); err != nil {
				log.Println(err)
			}
		}
	}()

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)

	//конструкция <- дожидается сигнала и блокирует код дальше, пока не прийдет сигнал
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)

	defer cancel()

	log.Println("Server is shutting down...")
	server.Shutdown(ctx)
	os.Exit(0)
}
