package main

import "sync"

type SynMap[K comparable, V any] struct {
	mu   sync.RWMutex
	data map[K]V
}

func (m *SynMap[K, V]) put(key K, val V) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[key] = val
}

func NewSyncMap[K comparable, V any]() *SynMap[K, V] {
	return &SynMap[K, V]{
		data: make(map[K]V),
	}
}

func main() {
	m := NewSyncMap[string, int]()
	m.put("a", 1)
}
