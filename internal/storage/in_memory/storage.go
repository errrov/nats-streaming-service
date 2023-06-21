package inMemoryStorage

import "wildberries_L0/internal/model"

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
