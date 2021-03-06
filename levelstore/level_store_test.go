package levelstore_test

import (
	"testing"

	"github.com/pressly/chainstore"
	"github.com/pressly/chainstore/levelstore"
	. "github.com/smartystreets/goconvey/convey"
)

func TestLevelStore(t *testing.T) {
	var store chainstore.Store
	var err error

	store = levelstore.New(chainstore.TempDir())
	err = store.Open()
	if err != nil {
		t.Error(err)
	}
	defer store.Close()

	Convey("Leveldb Open", t, func() {

		Convey("Put a bunch of objects", func() {
			e1 := store.Put("hi", []byte{1, 2, 3})
			e2 := store.Put("bye", []byte{4, 5, 6})
			So(e1, ShouldEqual, nil)
			So(e2, ShouldEqual, nil)
		})

		Convey("Get those objects", func() {
			v1, _ := store.Get("hi")
			v2, _ := store.Get("bye")
			So(v1, ShouldResemble, []byte{1, 2, 3})
			So(v2, ShouldResemble, []byte{4, 5, 6})
		})

		Convey("Delete those objects", func() {
			e1 := store.Del("hi")
			e2 := store.Del("bye")
			So(e1, ShouldEqual, nil)
			So(e2, ShouldEqual, nil)
		})

	})
}
