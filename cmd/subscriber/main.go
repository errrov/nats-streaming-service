package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"wildberries_L0/internal/broker"
	"wildberries_L0/internal/model"
	"wildberries_L0/internal/storage"

	"github.com/nats-io/stan.go"
)

func main() {
	sc, err := broker.ConnectToNats("test-cluster", "order-subscriber")
	l := log.New(os.Stdout, "nats-subscriber ", log.LstdFlags)
	if err != nil {
		log.Println(err)
	}
	storage, err := storage.StorageInit(l)
	defer storage.Postgres.Db.Close()
	if err != nil {
		l.Printf("Got error init storage: %v", err)
	}
	var Order model.Order
	log.Printf("Connected")
	signalChan := make(chan os.Signal, 1)
	cleanupDone := make(chan bool)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		for range signalChan {
			fmt.Printf("\nReceived an interrupt, unsubscribing and closing connection...\n\n")
			sc.Close()
			cleanupDone <- true
		}
	}()
	sub, _ := sc.Subscribe("testing", func(msg *stan.Msg) {
		if err := msg.Ack(); err != nil {
			log.Println(err)
			return
		}
		if err := json.Unmarshal(msg.Data, &Order); err != nil {
			log.Println(err)
			return
		}
		if err = storage.AddToStorage(&Order); err != nil {
			log.Printf("Error adding order: %v with orderUID %v", err, Order.OrderUID)
			return
		}
		log.Println("Got orderUID and added to cache", Order.OrderUID)
	}, stan.SetManualAckMode())
	defer sub.Close()
	<-cleanupDone
}
