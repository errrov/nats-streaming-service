package config

import (
	"log"
	"testing"
	"os"
)

func TestPsqlDefaultConfig(t *testing.T) {
	l := log.New(os.Stdout, "testing", log.LstdFlags)
	gotConfig := InitPsqlConfig(l)
	wantConfig := PostgresConfig{
		User:     "postgres",
		Password: "postgres",
		Host:     "localhost",
		Port:     "5432",
		Name:     "postgres",
	}
	if gotConfig != wantConfig {
		t.Errorf("Basic POSTGRES config creation is wrong, got: %v, want: %v", gotConfig, wantConfig)
	}
}

func TestNatsDefaultConfig(t *testing.T) {
	l := log.New(os.Stdout, "testing", log.LstdFlags)
	gotConfig := InitNatsConfig(l)
	wantConfig := NatsConfig{
		ClusterID: "test-cluster",
		PublsherID: "order-publisher",
		SubscriberID: "order-subscriber",
		Subject: "orders",
	}
	if *gotConfig != wantConfig {
		t.Errorf("Basic NATS config creation is wrong, got %v, want %v ", *gotConfig, wantConfig)
	}
}