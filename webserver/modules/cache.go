package modules

import (
	"sync"
)

type Cache struct {
	sync.RWMutex
	data map[string]interface{}
}

func NewCache() *Cache {
	data := make(map[string]interface{})
	return &Cache{
		data: data,
	}
}

func (c *Cache) Set(key string, value interface{}) {
	c.Lock()
	defer  c.Unlock()
	c.data[key] = value
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.RLock()
	defer c.RUnlock()
	if item, found := c.data[key]; !found {
		return nil, false
	} else {
		return item, true
	}
}

func (c *Cache) Delete(key string) {
	c.Lock()
	defer c.Unlock()
	delete(c.data, key)
}

func (c *Cache) GetAllIDs() []string {
	var ids []string

	c.RLock()
	defer c.RUnlock()

	for key := range c.data {
		ids = append(ids, key)
	}

	return ids
}