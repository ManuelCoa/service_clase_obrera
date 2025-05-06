package interfaz

import (
	"claseobrera/caching/operations"
	"claseobrera/caching/utils"
)

// Cache es la estructura principal que maneja la caché.
type Cache struct {
	storage map[string]utils.CacheItem
	policy  string
	maxSize int
}

// NewCache crea una nueva instancia de Cache.
func NewCache(policy string, maxSize int) *Cache {
	return &Cache{
		storage: make(map[string]utils.CacheItem),
		policy:  policy,
		maxSize: maxSize,
	}
}

// Get obtiene un valor de la caché.
func (c *Cache) Get(key string) (interface{}, bool) {
	return operations.Get(c.storage, key)
}

// Set almacena un valor en la caché.
func (c *Cache) Set(key string, value interface{}, size int) {
	operations.Set(c.storage, key, value, size, c.policy, c.maxSize)
}

// Delete elimina un valor de la caché.
func (c *Cache) Delete(key string) {
	operations.Delete(c.storage, key)
}
