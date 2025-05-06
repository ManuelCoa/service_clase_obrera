package operations

import (
	"claseobrera/caching/utils"
	"time"
)

// Get obtiene un valor de la caché.
func Get(storage map[string]utils.CacheItem, key string) (interface{}, bool) {
	item, exists := storage[key]
	if exists {
		// Actualizar el tiempo de acceso para LRU
		item.AccessTime = time.Now()
		item.AccessCount++
		storage[key] = item
		return item.Value, true
	}
	return nil, false
}

// Set almacena un valor en la caché.
func Set(storage map[string]utils.CacheItem, key string, value interface{}, size int, policy string, maxSize int) {
	if len(storage) >= maxSize {
		switch policy {
		case "LRU":
			utils.EvictLRU(storage, maxSize)
		case "LFU":
			utils.EvictLFU(storage, maxSize)
		case "FIFO":
			utils.EvictFIFO(storage, maxSize)
		case "MRU":
			utils.EvictMRU(storage, maxSize)
		case "Knapsack":
			utils.EvictKnapsack(storage, maxSize)
		}
	}
	storage[key] = utils.CacheItem{Value: value, Size: size, AccessTime: time.Now(), AccessCount: 1}
}

// Delete elimina un valor de la caché.
func Delete(storage map[string]utils.CacheItem, key string) {
	delete(storage, key)
}
