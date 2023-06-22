package broker

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nats-io/stan.go"
)

type ProductInfo struct {
	ClusterID string
	CLientID  string
	Channel   string
}

type SubscriberInfo struct {
	ClusterID string
	ClientID  string
	Channel   string
	conn      *pgxpool.Pool
}

func ConnectToNats(ClusterId, ClientID string) (stan.Conn, error) {
	sc, err := stan.Connect(ClusterId, ClientID)
	if err != nil {
		return nil, err
	}
	return sc, nil
}
