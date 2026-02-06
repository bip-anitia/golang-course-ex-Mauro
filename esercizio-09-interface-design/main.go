package main

import (
	"errors"
	"fmt"
	"sort"
	"sync"
)

var ErrNotFound = errors.New("key not found")

type Storage interface {
	Get(key string) ([]byte, error)
	Put(key string, value []byte) error
	Delete(key string) error
	List() ([]string, error)
	Close() error
}

type MemoryStorage struct {
	mu     sync.RWMutex
	data   map[string][]byte
	closed bool
}

func (m *MemoryStorage) Get(key string) ([]byte, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if m.closed {
		return nil, errors.New("storage closed")
	}

	value, ok := m.data[key]
	if !ok {
		return nil, ErrNotFound
	}

	// copy difensiva: evita che il caller modifichi lo stato interno
	copied := append([]byte(nil), value...)
	return copied, nil
}

func (m *MemoryStorage) Put(key string, value []byte) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.closed {
		return errors.New("storage closed")
	}

	m.data[key] = append([]byte(nil), value...)
	return nil
}

func (m *MemoryStorage) Delete(key string) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.closed {
		return errors.New("storage closed")
	}

	if _, ok := m.data[key]; !ok {
		return ErrNotFound
	}
	delete(m.data, key)
	return nil
}

func (m *MemoryStorage) List() ([]string, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if m.closed {
		return nil, errors.New("storage closed")
	}

	keys := make([]string, 0, len(m.data))
	for key := range m.data {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys, nil
}

func (m *MemoryStorage) Close() error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.closed {
		return nil
	}
	m.closed = true
	return nil
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		data: make(map[string][]byte),
	}
}

type CachedStorage struct {
	backend Storage
	cache   map[string][]byte
	mu      sync.RWMutex
}

func main() {
	// TODO: Implementare il sistema di storage con interfacce
	fmt.Println("Interface Design - Storage System")
}
