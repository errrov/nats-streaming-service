package broker

import (
	"log"
	"nats-streaming-service/internal/storage"
	"os"
	"testing"

	"github.com/nats-io/nats.go"
)

func TestConnection(t *testing.T) {
	l := log.New(os.Stdout, "testing", log.LstdFlags)
	s := &storage.Storage{}
	sub, err := CreateSubscriber(s, l)
	if err != nil {
		t.Errorf("Error creating subscriber %v", err)
	}
	defer sub.Conn.Close()
	if sub.Conn.NatsConn().Status() != nats.CONNECTED {
		t.Errorf("Not connected %v", sub.Conn.NatsConn().Status())
	}
}
