package filestore

import (
	"io/ioutil"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFileStore(t *testing.T) {
	dir, _ := ioutil.TempDir("", "chainstore-")
	var err error

	Convey("Fsdb Open", t, func() {
		store := New(dir, 0755)
		err = nil
		So(err, ShouldEqual, nil)

		Convey("Put/Get/Del basic data", func() {
			err = store.Put("test.txt", []byte{1, 2, 3, 4})
			So(err, ShouldEqual, nil)

			data, err := store.Get("test.txt")
			So(err, ShouldEqual, nil)
			So(data, ShouldResemble, []byte{1, 2, 3, 4})
		})

		Convey("Auto-creating directories on put", func() {
			err = store.Put("hello/there/everyone.txt", []byte{1, 2, 3, 4})
			So(err, ShouldEqual, nil)
		})

	})
}
