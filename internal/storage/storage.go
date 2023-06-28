package storage

import (
	"log"
	"wildberries_L0/internal/model"
	inMemoryStorage "wildberries_L0/internal/storage/in_memory"
	"wildberries_L0/internal/storage/psql"
)

type Storage struct {
	Mem_cache inMemoryStorage.MemoryStorage
	Postgres  *psql.Postgresql
	Logger    *log.Logger
}

func StorageInit(logger *log.Logger) (*Storage, error) {
	logger.Println("Trying to create connection to postgres")
	postgres := psql.Connect()
	logger.Println("Created connection to postgres")
	cache, err := postgres.FindAll()
	logger.Println("Found all", cache, err)
	if err != nil {
		return nil, err
	}
	return &Storage{
		Mem_cache: inMemoryStorage.MemoryStorage{MemoryMap: cache},
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
