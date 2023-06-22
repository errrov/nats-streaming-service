package broker

import (
	"testing"
	"github.com/nats-io/nats.go"
)

func TestPublisher(t *testing.T) {
	sc, err := ConnectToNats("test-cluster", "order-publisher")
	if err != nil {
		t.Fatalf("Error creating connection to NATS: %v", err)
	}
	defer sc.Close()
	if sc.NatsConn().Status() != nats.CONNECTED {
		t.Errorf("Error connecting to NATS: %v", sc.NatsConn().Status())
	}

}
