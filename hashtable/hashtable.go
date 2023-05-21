package hashtable

import (
	"sync"
)

type HashTable interface {
	Size() int
	Get(key uint64) ([]uint64, bool)
	Set(key uint64, size int) []uint64
}

type hashTable struct {
	lock *sync.RWMutex
	m    map[uint64][]uint64
}

func NewHashTable(sizeLog int) HashTable {
	return &hashTable{lock: &sync.RWMutex{}, m: make(map[uint64][]uint64, 1<<sizeLog)}
}

func (h *hashTable) Size() int {
	h.lock.RLock()
	size := len(h.m)
	h.lock.RUnlock()
	return size
}

func (h *hashTable) Get(key uint64) ([]uint64, bool) {
	h.lock.RLock()
	value, ok := h.m[key]
	h.lock.RUnlock()
	return value, ok
}

func (h *hashTable) Set(key uint64, size int) []uint64 {
	buffer := make([]uint64, size)
	h.lock.Lock()
	h.m[key] = buffer
	h.lock.Unlock()
	return buffer
}
