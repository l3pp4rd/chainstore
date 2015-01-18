package chainstore

import (
	"container/list"
	"io/ioutil"
	"testing"

	"github.com/pressly/chainstore/filestore"
	. "github.com/smartystreets/goconvey/convey"
)

func TestLRUManager(t *testing.T) {
	var err error
	dir, _ := ioutil.TempDir("", "chainstore-")
	var store Store = filestore.New(dir, 0755)
	var lru *lruManager
	var capacity int64 = 20

	Convey("LRUManager", t, func() {
		lru = &lruManager{
			store:    store,
			capacity: capacity,
			cushion:  int64(float64(capacity) * 0.1),
			items:    make(map[string]*lruItem, 10000),
			list:     list.New(),
		}

		// based on 10% cushion
		lru.Put("peter", []byte{1, 2, 3})
		lru.Put("jeff", []byte{4})
		lru.Put("julia", []byte{5, 6, 7, 8, 9, 10})
		lru.Put("janet", []byte{11, 12, 13})
		lru.Put("ted", []byte{14, 15, 16, 17, 18})

		remaining := capacity - 18
		So(len(lru.capacity), ShouldEqual, remaining)

		remaining = remaining + 4
		err = lru.Put("agnes", []byte{20, 21, 22, 23, 24, 25})
		So(lru.capacity, ShouldEqual, remaining)
		So(err, ShouldEqual, nil)
	})

	Convey("Should wrap store", t, func() {
		wrapped := LruCacheable(capacity, store)
		wrapped.Put("key", []byte("val"))

		So(store.Get("key"), ShouldEqual, "val")
	})
}
