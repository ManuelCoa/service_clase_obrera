package utils

import (
	"sort"
	"time"
)

// CacheItem representa un elemento en la caché.
type CacheItem struct {
	Value       interface{}
	Size        int
	AccessTime  time.Time
	AccessCount int
}

// EvictLRU elimina los elementos menos usados recientemente.
func EvictLRU(storage map[string]CacheItem, maxSize int) {
	type kv struct {
		Key   string
		Value CacheItem
	}
	var sortedStorage []kv
	for k, v := range storage {
		sortedStorage = append(sortedStorage, kv{k, v})
	}
	sort.Slice(sortedStorage, func(i, j int) bool {
		return sortedStorage[i].Value.AccessTime.Before(sortedStorage[j].Value.AccessTime)
	})
	for i := 0; i < len(sortedStorage)-maxSize+1; i++ {
		delete(storage, sortedStorage[i].Key)
	}
}

// EvictLFU elimina los elementos con menos accesos históricos.
func EvictLFU(storage map[string]CacheItem, maxSize int) {
	type kv struct {
		Key   string
		Value CacheItem
	}
	var sortedStorage []kv
	for k, v := range storage {
		sortedStorage = append(sortedStorage, kv{k, v})
	}
	sort.Slice(sortedStorage, func(i, j int) bool {
		return sortedStorage[i].Value.AccessCount < sortedStorage[j].Value.AccessCount
	})
	for i := 0; i < len(sortedStorage)-maxSize+1; i++ {
		delete(storage, sortedStorage[i].Key)
	}
}

// EvictFIFO elimina los elementos más antiguos.
func EvictFIFO(storage map[string]CacheItem, maxSize int) {
	type kv struct {
		Key   string
		Value CacheItem
	}
	var sortedStorage []kv
	for k, v := range storage {
		sortedStorage = append(sortedStorage, kv{k, v})
	}
	sort.Slice(sortedStorage, func(i, j int) bool {
		return sortedStorage[i].Value.AccessTime.Before(sortedStorage[j].Value.AccessTime)
	})
	for i := 0; i < len(sortedStorage)-maxSize+1; i++ {
		delete(storage, sortedStorage[i].Key)
	}
}

// EvictMRU elimina los elementos más usados recientemente.
func EvictMRU(storage map[string]CacheItem, maxSize int) {
	type kv struct {
		Key   string
		Value CacheItem
	}
	var sortedStorage []kv
	for k, v := range storage {
		sortedStorage = append(sortedStorage, kv{k, v})
	}
	sort.Slice(sortedStorage, func(i, j int) bool {
		return sortedStorage[i].Value.AccessTime.After(sortedStorage[j].Value.AccessTime)
	})
	for i := 0; i < len(sortedStorage)-maxSize+1; i++ {
		delete(storage, sortedStorage[i].Key)
	}
}

// EvictKnapsack elimina elementos basado en el algoritmo de la mochila.
func EvictKnapsack(storage map[string]CacheItem, maxSize int) {
	type kv struct {
		Key   string
		Value CacheItem
	}
	var sortedStorage []kv
	for k, v := range storage {
		sortedStorage = append(sortedStorage, kv{k, v})
	}
	sort.Slice(sortedStorage, func(i, j int) bool {
		return int64(sortedStorage[i].Value.AccessCount)*sortedStorage[i].Value.AccessTime.Unix() < int64(sortedStorage[j].Value.AccessCount)*sortedStorage[j].Value.AccessTime.Unix()
	})
	for i := 0; i < len(sortedStorage)-maxSize+1; i++ {
		delete(storage, sortedStorage[i].Key)
	}
}
