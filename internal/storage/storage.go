package storage

import (
	"log"
	"nats-streaming-service/internal/model"
	"nats-streaming-service/internal/storage/memcache"
	"nats-streaming-service/internal/storage/psql"
)

type Storage struct {
	Mem_cache memcache.MemoryStorage
	Postgres  *psql.Postgresql
	Logger    *log.Logger
}

func StorageInit(logger *log.Logger) (*Storage, error) {
	logger.Println("Trying to create connection to postgres")
	postgres, err := psql.Connect(logger)
	if err != nil {
		return nil, err
	}
	logger.Println("Created connection to postgres")
	cache, err := postgres.FindAll()
	if err != nil {
		return nil, err
	}
	return &Storage{
		Mem_cache: memcache.MemoryStorage{MemoryMap: cache},
		Postgres:  postgres,
		Logger:    logger,
	}, nil
}

func (s *Storage) AddToStorage(order *model.Order) error {
	err := s.Mem_cache.Add(order)
	if err != nil {
		return err
	}
	s.Logger.Printf("Added order to cache storage with UID %v", order.OrderUID)
	err = s.Postgres.InsertOrder(order)
	if err != nil {
		return err
	}
	s.Logger.Printf("Added order to database with UID %v", order.OrderUID)
	return nil
}

func (s *Storage) FindByUID(uid string) (*model.Order, error) {
	order, err := s.Mem_cache.FindByUID(uid)
	if err != nil {
		return nil, model.ErrNotFound
	}
	return order, nil
}
