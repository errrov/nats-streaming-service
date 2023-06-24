package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"wildberries_L0/internal/broker"
	"wildberries_L0/internal/model"
	inMemoryStorage "wildberries_L0/internal/storage/in_memory"

	"github.com/nats-io/stan.go"
)

func main() {
	sc, err := broker.ConnectToNats("test-cluster", "order-subscriber")
	if err != nil {
		log.Println(err)
	}
	cache := inMemoryStorage.NewInMemory()
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
		if err = cache.Add(&Order); err != nil {
			log.Printf("Error adding order: %v with orderUID %v", err, Order.OrderUID)
			return
		}
		log.Println("Got orderUID and added to cache", Order.OrderUID)
	}, stan.SetManualAckMode())
	defer sub.Close()
	<-cleanupDone
}
