package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"reflect"

	"github.com/nats-io/stan.go"
)

func main() {
	sc, err := stan.Connect("test-cluster", "order-subscriber", stan.NatsURL(stan.DefaultNatsURL))
	if err != nil {
		log.Println(err)
	}
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
	_, _ = sc.Subscribe("testing", func(msg *stan.Msg) {
		if err := msg.Ack(); err != nil {
			log.Println(err)
			return
		}
		log.Println(string(msg.Data))
		log.Println(reflect.TypeOf(msg.Data))
	}, stan.SetManualAckMode())

	<-cleanupDone
	/*
		fmt.Println("Hello world")
		srv := server.NewServer()
		go func() {
			srv.Server.ErrorLog.Println("Started server on port :8080")
			if err := srv.Server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
				srv.Server.ErrorLog.Printf("Error starting server: %s\n", err)
				os.Exit(1)
			}
		}()

		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt)
		signal.Notify(quit, syscall.SIGTERM)

		sig := <-quit
		log.Printf("Got signal : %v", sig)

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		if err := srv.Server.Shutdown(ctx); err != nil {
			log.Fatalf("Server shutdown failed: %v\n", err)
		}
	*/
}
