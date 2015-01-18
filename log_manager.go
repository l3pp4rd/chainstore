package chainstore

import (
	"fmt"
	"log"
)

type logManager struct {
	store  Store
	logger *log.Logger
	tag    string
}

func Loggable(logger *log.Logger, tag string, store Store) Store {
	if tag != "" {
		tag = fmt.Sprintf(" [%s]", tag)
	}
	return &logManager{
		store:  store,
		logger: logger,
		tag:    tag,
	}
}

func (m *logManager) Open() error {
	return m.store.Open()
}

func (m *logManager) Close() error {
	return m.store.Open()
}

func (m *logManager) Put(key string, val []byte) error {
	m.logger.Printf("chainstore%s: Put %s of %d bytes", m.tag, key, len(val))
	return m.store.Put(key, val)
}

func (m *logManager) Get(key string) ([]byte, error) {
	m.logger.Printf("chainstore%s: Get %s", m.tag, key)
	return m.store.Get(key)
}

func (m *logManager) Del(key string) error {
	m.logger.Printf("chainstore%s: Del %s", m.tag, key)
	return m.store.Del(key)
}
