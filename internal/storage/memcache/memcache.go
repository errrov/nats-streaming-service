package memcache

import (
	"log"
	"nats-streaming-service/internal/model"
)

type MemoryStorage struct {
	MemoryMap map[string]*model.Order
}

func NewInMemory() *MemoryStorage {
	return &MemoryStorage{
		MemoryMap: make(map[string]*model.Order),
	}
}

func (ms *MemoryStorage) Add(order *model.Order) error {
	if _, ok := ms.MemoryMap[order.OrderUID]; ok {
		return model.ErrAlreadyExist
	}
	ms.MemoryMap[order.OrderUID] = order
	return nil
}

func (ms *MemoryStorage) FindByUID(uid string) (*model.Order, error) {
	if order, ok := ms.MemoryMap[uid]; ok {
		return order, nil
	}
	return nil, model.ErrNotFound
}

func (ms *MemoryStorage) ListAll() {
	for k := range ms.MemoryMap {
		log.Println(k)
	}

}
