package cache

import (
	"generate_stream_currency/internal/app/contract"
	"generate_stream_currency/pkg/model"
	"sync"
)

type Cache struct {
	sync.Mutex
	currency model.Currency
}

func NewCache() contract.ApiService {
	return &Cache{}
}

func (c *Cache) GetCurrency() model.Currency {
	c.Lock()
	defer c.Unlock()
	return c.currency
}

func (c *Cache) Run() error {
	return c.currency.RunGenerate()
}
