package cache

import "sync"

type MemoryCache struct {
	store sync.Map
}

var _ Driver = (*MemoryCache)(nil)

func NewMemoryCache() Driver {
	return &MemoryCache{}
}

func (c *MemoryCache) Get(key string) (interface{}, bool) {
	return c.store.Load(key)
}

func (c *MemoryCache) Gets(keys []string, prefix string) (map[string]interface{}, []string) {
	var miss []string
	values := make(map[string]interface{}, len(keys))
	for _, key := range keys {
		if value, ok := c.store.Load(prefix + key); ok {
			values[key] = value
		} else {
			miss = append(miss, key)
		}
	}

	return values, miss
}

func (c *MemoryCache) Set(key string, value interface{}, ttl int) error {
	c.store.Store(key, value)
	return nil
}

func (c *MemoryCache) Sets(values map[string]interface{}, prefix string) error {
	for key, value := range values {
		c.store.Store(prefix+key, value)
	}

	return nil
}

func (c *MemoryCache) Delete(key string) error {
	c.store.Delete(key)
	return nil
}

func (c *MemoryCache) Deletes(keys []string, prefix string) error {
	for _, key := range keys {
		c.store.Delete(prefix + key)
	}

	return nil
}
