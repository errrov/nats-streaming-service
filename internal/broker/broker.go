package broker

import (
	"encoding/json"
	"log"
	"nats-streaming-service/internal/config"
	"nats-streaming-service/internal/model"
	"nats-streaming-service/internal/storage"

	"github.com/go-playground/validator/v10"
	"github.com/nats-io/stan.go"
)

type Subscriber struct {
	Conn      stan.Conn
	cache     *storage.Storage
	validator *validator.Validate
	NatsConf  *config.NatsConfig
	l         *log.Logger
}

type Publsiher struct {
	Conn stan.Conn
}

func CreatePublisher(l *log.Logger) (stan.Conn, error) {
	conf := config.InitNatsConfig(l)
	conn, err := stan.Connect(conf.ClusterID, conf.PublsherID)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func CreateSubscriber(s *storage.Storage, l *log.Logger) (*Subscriber, error) {
	conf := config.InitNatsConfig(l)
	conn, err := stan.Connect(conf.ClusterID, conf.SubscriberID)
	if err != nil {
		return nil, err
	}
	return &Subscriber{
		Conn:      conn,
		cache:     s,
		validator: validator.New(),
		NatsConf:  conf,
		l:         l,
	}, nil
}

func (s *Subscriber) Subscribe() error {
	_, err := s.Conn.Subscribe(s.NatsConf.Subject, func(msg *stan.Msg) {
		var Order model.Order
		if err := msg.Ack(); err != nil {
			s.l.Println(err)
			return
		}
		if err := json.Unmarshal(msg.Data, &Order); err != nil {
			s.l.Println(err)
			return
		}
		if err := s.validator.Struct(Order); err != nil {
			s.l.Printf("Error validating struct: %v", err)
			return
		}
		for _, v := range Order.Items {
			if err := s.validator.Struct(v); err != nil {
				s.l.Printf("Error validating item from order %v", err)
				return
			}
		}
		if err := s.cache.AddToStorage(&Order); err != nil {
			s.l.Printf("Error adding order: %v with orderUID %v", err, Order.OrderUID)
			return
		}
	}, stan.SetManualAckMode())

	if err != nil {
		s.l.Println(err)
	}
	return err
}
