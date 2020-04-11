package access

import (
	"sync"
)

// ReliableMap is a map can be used in some goroutines safely
type ReliableMap struct {
	sync.RWMutex
	dictionary map[string]string
}
// New return new pointer of ReliableMap
func New() *ReliableMap {
	return &ReliableMap{
		dictionary: make(map[string]string),
	}
}
// Get return the value corresponding to given key
func (m *ReliableMap) Get(key string) string {
	m.RLock()
	defer m.Unlock()
	value := m.dictionary[key]
	return value
}
// Set sets key and value to ReliableMap
func (m *ReliableMap) Set(key string, value string) {
	m.Lock()
	defer m.Unlock()
	m.dictionary[key] = value
}