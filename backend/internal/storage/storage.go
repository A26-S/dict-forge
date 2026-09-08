package storage

import (
	"fmt"
)

// Storage 数据存储接口
type Storage interface {
	Set(key string, value []byte) error
	Get(key string) ([]byte, error)
	Delete(key string) error
	Close() error
}

// Config 存储配置
type Config struct {
	Type     string // leveldb, sqlite, memory
	Path     string
	MaxSize  int64
}

// New 创建存储实例
func New(config Config) (Storage, error) {
	switch config.Type {
	case "leveldb":
		return NewLevelDBStorage(config.Path)
	case "memory":
		return NewMemoryStorage()
	default:
		return nil, fmt.Errorf("unsupported storage type: %s", config.Type)
	}
}

// MemoryStorage 内存存储
type MemoryStorage struct {
	data map[string][]byte
}

// NewMemoryStorage 创建内存存储
func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		data: make(map[string][]byte),
	}
}

func (m *MemoryStorage) Set(key string, value []byte) error {
	m.data[key] = value
	return nil
}

func (m *MemoryStorage) Get(key string) ([]byte, error) {
	if val, exists := m.data[key]; exists {
		return val, nil
	}
	return nil, fmt.Errorf("key not found: %s", key)
}

func (m *MemoryStorage) Delete(key string) error {
	delete(m.data, key)
	return nil
}

func (m *MemoryStorage) Close() error {
	return nil
}
