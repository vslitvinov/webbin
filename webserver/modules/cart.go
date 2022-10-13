package modules

import (
	"sync"

	"github.com/google/uuid"
)


type Product struct {
	Name string 
	Price float64
	Сategory string
}

type ItemCart struct {
	Product Product
	Count int64
}

// записи товара 

type Cart struct {
	sync.RWMutex
	items map[string]*ItemCart
}



func NewCart() *Cart {
	items := make(map[string]*ItemCart)
	return &Cart{
		items: items,
	}
}

func (c *Cart) EditCount(uuid string, n int64) bool{
	c.Lock()
	defer  c.Unlock()

	item, found := c.items[uuid]
	if !found {
		return  false
	} 

	item.Count = item.Count + n
	c.items[uuid] = item
	return  true 
}

func (c *Cart) Set(value *ItemCart) string {
	c.Lock()
	defer  c.Unlock()
	key := uuid.NewString()
	c.items[key] = value
	return key
}

func (c *Cart) Get(key string) (*ItemCart, bool) {
	c.RLock()
	defer c.RUnlock()
	if item, found := c.items[key]; !found {
		return nil, false
	} else {
		return item, true
	}
}

func (c *Cart) Delete(key string) {
	c.Lock()
	defer c.Unlock()
	delete(c.items, key)
}

func (c *Cart) GetAllIDs() []string {
	var ids []string

	c.RLock()
	defer c.RUnlock()

	for key := range c.items {
		ids = append(ids, key)
	}

	return ids
}