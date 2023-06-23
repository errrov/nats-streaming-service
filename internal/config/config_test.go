package config

import (
	"testing"
)

func TestPsqlDefaultConfig(t *testing.T) {
	gotConfig := InitPsqlConfig()
	wantConfig := PostgresConfig{
		User:     "postgres",
		Password: "postgres",
		Host:     "localhost",
		Port:     "5432",
		Name:     "postgres",
	}
	if *gotConfig != wantConfig {
		t.Errorf("Basic POSTGRES config creation is wrong, got: %v, want: %v", *gotConfig, wantConfig)
	}
}

func TestNatsDefaultConfig(t *testing.T) {
	gotConfig := InitNatsConfig()
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