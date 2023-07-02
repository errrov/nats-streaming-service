package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"nats-streaming-service/internal/broker"
	"nats-streaming-service/internal/config"
	"nats-streaming-service/internal/model"
	"os"
)

func main() {

	l := log.New(os.Stdout, "nats-subscriber ", log.LstdFlags)
	config := config.InitNatsConfig(l)

	var n int

	flag.IntVar(&n, "n", 1, "number of orders to publish")
	flag.Parse()
	l.Println(n)
	sc, err := broker.ConnectToNats(config.ClusterID, config.PublsherID)
	if err != nil {
		log.Fatalf("Error connecting to nats %v", err)
	}
	defer sc.Close()
	modelJSON, err := os.ReadFile("model.json")
	if err != nil {
		log.Fatalf("Error reading json: %v", err)
	}
	var sendOrder model.Order
	err = json.Unmarshal(modelJSON, &sendOrder)
	if err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
	}
	for i := 0; i < n; i++ {
		sendOrder.OrderUID = fmt.Sprintf("testing-ordering%d", i)
		log.Println(sendOrder.OrderUID)
		sendingData, err := json.Marshal(sendOrder)
		if err != nil {
			log.Fatalf("Error marshalling order %v", err)
		}
		if err = sc.Publish(config.Subject, sendingData); err != nil {
			log.Fatalf("Error sending Order %v", err)
		}
	}
}
