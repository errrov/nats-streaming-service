package main

import (
	"encoding/json"
	"log"
	"reflect"
	"wildberries_L0/internal/broker"

	"fmt"
	"os"
	"wildberries_L0/internal/model"
)

func main() {
	sc, err := broker.ConnectToNats("test-cluster", "order-publisher")
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
	fmt.Println(sendOrder)

	sendOrder.OrderUID = "ABCDTESTING123"
	sendingData, err := json.Marshal(sendOrder)
	if err != nil {
		log.Fatalf("Error marshalling order %v", err)
	}
	fmt.Println(reflect.TypeOf(sendingData))
	if err = sc.Publish("testing", sendingData); err != nil {
		log.Fatalf("Error sending Order %v", err)
	}

	/*
		modelJSON, err := os.ReadFile("model.json")
		if err != nil {
			log.Println(err)
		}
		var order model.Order
		err = json.Unmarshal(modelJSON, &order)
		if err != nil {
			log.Println(err)
		}
		order.OrderUID = "TESTING"
		data, err := json.Marshal(order)
		if err != nil {
			log.Println(err)
		}

		if err := pb.Publish("test-test", data); err != nil {
			log.Println(err)
		}

		log.Printf("Published order with UID = %v", order.OrderUID)
	*/
}
