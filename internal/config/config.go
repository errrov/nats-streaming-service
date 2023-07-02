package config

import (
	"log"
	"os"
)

type PostgresConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	Name     string
}

type NatsConfig struct {
	ClusterID    string
	PublsherID   string
	SubscriberID string
	Subject      string
}

func InitPsqlConfig(l *log.Logger) PostgresConfig {
	User, exist := os.LookupEnv("POSTGRES_USER")
	if !exist {
		l.Println("couldn't find postgres user, using default [postgres]")
		User = "postgres"
	}
	Password, exist := os.LookupEnv("POSTGRES_PASSWORD")
	if !exist {
		l.Println("couldn't find postgres password, using default [postgrespw]")
		Password = "postgrespw"
	}
	Host, exist := os.LookupEnv("POSTGRES_HOST")
	if !exist {
		l.Println("couldn't find postgres host, using default [localhost]")
		Host = "localhost"
	}
	Port, exist := os.LookupEnv("POSTGRES_PORT")
	if !exist {
		l.Println("couldn't find postgres port, using default [5432]")
		Port = "5432"
	}
	Name, exist := os.LookupEnv("POSTGRES_DB")
	if !exist {
		l.Println("couldn't find postgres name, using default [postgres]")
		Name = "postgres"
	}
	return PostgresConfig{
		User:     User,
		Password: Password,
		Host:     Host,
		Port:     Port,
		Name:     Name,
	}
}

func InitNatsConfig(l *log.Logger) *NatsConfig {
	clusterId, exist := os.LookupEnv("NATS_CLUSTERID")
	if !exist {
		l.Println("couldn't find nats cluster id using default [test-cluster]")
		clusterId = "test-cluster"
	}

	publishedID, exist := os.LookupEnv("NATS_PUBLISHER")
	if !exist {
		l.Println("couldn't find nats publisher id using default [order-publisher]")
		publishedID = "order-publisher"
	}

	subscriberID, exist := os.LookupEnv("NATS_SUBSCRIBER")
	if !exist {
		l.Println("couldn't find nats subscriber id using default [order-subscriber]")
		subscriberID = "order-subscriber"
	}

	subject, exist := os.LookupEnv("NATS_SUBJECT")
	if !exist {
		l.Println("couldn't find nats subject using default [orders]")
		subject = "orders"
	}

	return &NatsConfig{
		ClusterID:    clusterId,
		PublsherID:   publishedID,
		SubscriberID: subscriberID,
		Subject:      subject,
	}
}
