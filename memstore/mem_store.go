package memstore

import (
	"sync"

	"github.com/pressly/chainstore"
	"github.com/pressly/chainstore/lrumgr"
)

type memStore struct {
	sync.Mutex
	data map[string][]byte
}

func New(capacity int64) chainstore.Store {
	memStore := &memStore{data: make(map[string][]byte, 1000)}
	store := lrumgr.New(capacity, memStore)
	return store
}

func (s *memStore) Open() (err error)  { return }
func (s *memStore) Close() (err error) { return }

func (s *memStore) Put(key string, val []byte) (err error) {
	s.Lock()
	s.data[key] = val
	s.Unlock()
	return nil
}

func (s *memStore) Get(key string) ([]byte, error) {
	return s.data[key], nil
}

func (s *memStore) Del(key string) (err error) {
	s.Lock()
	delete(s.data, key)
	s.Unlock()
	return
}
