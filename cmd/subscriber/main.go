package main

import (
	"context"
	"log"
	"nats-streaming-service/internal/broker"
	"nats-streaming-service/internal/server"
	"os"
	"os/signal"
	"syscall"

)

func main() {
	l := log.New(os.Stdout, "nats-subscriber ", log.LstdFlags)
	srv, err := server.NewServer(l)
	if err != nil {
		l.Println("Error creating server")
		os.Exit(1)
	}
	sc, err := broker.CreateSubscriber(srv.Cache, l)
	if err != nil {
		l.Println("Error connecting to Nats-streaming")
		os.Exit(1)
	}
	signalChan := make(chan os.Signal, 1)
	cleanupDone := make(chan bool)
	signal.Notify(signalChan, os.Interrupt)
	signal.Notify(signalChan, syscall.SIGTERM)
	defer srv.Cache.Postgres.Db.Close()
	err = sc.Subscribe()
	if err != nil {
		l.Println("Error subscribing to nats-streaming", err)
		os.Exit(1)
	}
	go func() {
		err := srv.Srv.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server : %s\n", err)
			os.Exit(1)
		}
	}()
	go func() {
		for range signalChan {
			l.Printf("\nReceived an interrupt, unsubscribing and closing connection...\n\n")
			sc.Conn.Close()
			srv.Srv.Shutdown(context.Background())
			cleanupDone <- true
		}
	}()
	<-cleanupDone

}
