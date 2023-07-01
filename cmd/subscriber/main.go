package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"nats-streaming-service/internal/broker"
	"nats-streaming-service/internal/model"
	"nats-streaming-service/internal/server"
	"os"
	"os/signal"
	"syscall"

	"github.com/nats-io/stan.go"
)

func main() {
	l := log.New(os.Stdout, "nats-subscriber ", log.LstdFlags)
	sc, err := broker.ConnectToNats("test-cluster", "order-subscriber")
	if err != nil {
		l.Println("Error connecting to Nats-streaming")
		os.Exit(1)
	}
	defer sc.Close()
	l.Println("Connected to NATS | ")
	if err != nil {
		log.Println(err)
	}
	if err != nil {
		l.Printf("Got error init storage: %v", err)
	}
	l.Println("Connected to DB | ")
	var Order model.Order
	signalChan := make(chan os.Signal, 1)
	cleanupDone := make(chan bool)
	signal.Notify(signalChan, os.Interrupt)
	signal.Notify(signalChan, syscall.SIGTERM)
	srv := server.NewServer(l)
	defer srv.Cache.Postgres.Db.Close()
	go func() {
		err := srv.Srv.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server : %s\n", err)
			os.Exit(1)
		}
	}()
	sub, _ := sc.Subscribe("testing", func(msg *stan.Msg) {
		if err := msg.Ack(); err != nil {
			l.Println(err)
			return
		}
		if err := json.Unmarshal(msg.Data, &Order); err != nil {
			l.Println(err)
			return
		}
		if err = srv.Cache.AddToStorage(&Order); err != nil {
			l.Printf("Error adding order: %v with orderUID %v", err, Order.OrderUID)
			return
		}
	}, stan.SetManualAckMode())
	defer sub.Close()
	go func() {
		for range signalChan {
			fmt.Printf("\nReceived an interrupt, unsubscribing and closing connection...\n\n")
			sc.Close()
			srv.Srv.Shutdown(context.Background())
			cleanupDone <- true
		}
	}()
	<-cleanupDone

}
