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

func InitPsqlConfig() *PostgresConfig {
	User, exist := os.LookupEnv("POSTGRES_USER")
	if !exist {
		log.Println("[config] COULDN'T FIND POSTGRES USER, USING DEFAULT postgres")
		User = "postgres"
	}
	Password, exist := os.LookupEnv("POSTGRES_PASSWORD")
	if !exist {
		log.Println("[config] COULDN'T FIND POSTGRES PASSWORD, USING DEFAULT postgres")
		Password = "postgres"
	}
	Host, exist := os.LookupEnv("POSTGRES_HOST")
	if !exist {
		log.Println("[config] COULDN'T FIND POSTGRES HOST, USING DEFAULT localhost")
		Host = "localhost"
	}
	Port, exist := os.LookupEnv("POSTGRES_PORT")
	if !exist {
		log.Println("[config] COULDN'T FIND POSTGRES PORT, USING DEFAULT 5432")
		Port = "5432"
	}
	Name, exist := os.LookupEnv("POSTGRES_NAME")
	if !exist {
		log.Println("[config] COULDN'T FIND POSTGRES NAME, USING DEFAULT POSTGRES")
		Name = "postgres"
	}
	return &PostgresConfig{
		User:     User,
		Password: Password,
		Host:     Host,
		Port:     Port,
		Name:     Name,
	}
}

func InitNatsConfig() *NatsConfig {
	clusterId, exist := os.LookupEnv("NATS_CLUSTERID")
	if !exist {
		log.Println("[config] COULDN'T FIND NATS CLUSTER ID USING DEFAULT test-cluster")
		clusterId = "test-cluster"
	}

	publishedID, exist := os.LookupEnv("NATS_PUBLISHER")
	if !exist {
		log.Println("[config] COULDN'T FIND NATS PUBLISHER ID USING DEFAULT order-publisher")
		publishedID = "order-publisher"
	}

	subscriberID, exist := os.LookupEnv("NATS_SUBSCRIBER")
	if !exist {
		log.Println("[config] COULDN'T FIND NATS SUBSCRIBER ID USING DEFAULT order-subscriber")
		subscriberID = "order-subscriber"
	}

	subject, exist := os.LookupEnv("NATS_SUBJECT")
	if !exist {
		log.Println("[config] COULDN'T FIND NATS SUBJECT USING DEFAULT orders")
		subject = "orders"
	}

	return &NatsConfig{
		ClusterID:    clusterId,
		PublsherID:   publishedID,
		SubscriberID: subscriberID,
		Subject:      subject,
	}
}