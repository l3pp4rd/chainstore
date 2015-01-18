package chainstore

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
	"testing"
	"time"
)

type tStore interface {
	Del(key string) error
}

type tChain struct {
	stores []tStore
}

func tNew(stores ...tStore) tStore {
	return &tChain{stores: stores}
}

// same can be applied for all chain store actions
func (c *tChain) Del(key string) error {
	var wg sync.WaitGroup
	var errs []string
	wg.Add(len(c.stores))
	for _, s := range c.stores {
		go func(st tStore) {
			defer wg.Done()
			if err := st.Del(key); err != nil {
				errs = append(errs, err.Error())
			}
		}(s)
	}
	wg.Wait()
	if len(errs) > 0 {
		return fmt.Errorf(strings.Join(errs, ", "))
	}
	return nil
}

type mockAsyncStore struct {
	Name string
}

func (s mockAsyncStore) Del(key string) error {
	rand.Seed(time.Now().UnixNano())
	ms := time.Duration(rand.Int31n(1000)) * time.Millisecond
	time.Sleep(ms)
	fmt.Printf("Deleting %s for store %s\n", key, s.Name)
	if ms < time.Duration(500*time.Millisecond) {
		return fmt.Errorf("Some error of store: %s", s.Name)
	}
	return nil
}

func TestShouldRunAllAsync(t *testing.T) {
	a := mockAsyncStore{Name: "A"}
	b := mockAsyncStore{Name: "B"}
	c := mockAsyncStore{Name: "C"}

	chain := tNew(a, b, c)
	chain.Del("key")
}

func TestShouldHandleNestedAsyncStoreChain(t *testing.T) {

	a := mockAsyncStore{Name: "A"}
	b := mockAsyncStore{Name: "B"}
	c := mockAsyncStore{Name: "C"}
	d := mockAsyncStore{Name: "D"}
	e := mockAsyncStore{Name: "E"}

	ab := tNew(a, b)
	abc := tNew(ab, c)
	de := tNew(d, e)
	abcde := tNew(abc, de)
	if err := abcde.Del("key"); err != nil {
		fmt.Println(err)
	}
}
