package chainstore

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/pressly/chainstore/filestore"
	"github.com/pressly/chainstore/memstore"
	. "github.com/smartystreets/goconvey/convey"
)

func TestBasicChain(t *testing.T) {
	var ms, fs, chain Store
	var err error
	dir, _ := ioutil.TempDir("", "chainstore-")

	logger := log.New(os.Stdout, "", log.LstdFlags)

	Convey("Basic chain", t, func() {
		err = nil

		ms = memstore.New(100)
		fs = filestore.New(dir+"/filestore", 0755)

		chain = Loggable(logger, "", New(
			ms,
			fs,
		))
		err = chain.Open()
		So(err, ShouldEqual, nil)

		Convey("Put/Get/Del", func() {
			v := []byte("value")
			err = chain.Put("k", v)
			So(err, ShouldEqual, nil)

			val, err := chain.Get("k")
			So(err, ShouldEqual, nil)
			So(v, ShouldResemble, v)

			val, err = ms.Get("k")
			So(err, ShouldEqual, nil)
			So(val, ShouldResemble, v)

			val, err = fs.Get("k")
			So(err, ShouldEqual, nil)
			So(val, ShouldResemble, v)

			err = chain.Del("k")
			So(err, ShouldEqual, nil)

			val, err = fs.Get("k")
			So(err, ShouldEqual, nil)
			So(len(val), ShouldEqual, 0)
		})
	})
}

// func TestAsyncChain(t *testing.T) {
// 	var ms, fs, bs, chain Store
// 	var err error
// 	dir, _ := ioutil.TempDir("", "chainstore-")

// 	logger := log.New(os.Stdout, "", log.LstdFlags)

// 	Convey("Async chain", t, func() {
// 		err = nil

// 		ms = memstore.New(100)
// 		fs = filestore.New(dir+"/filestore", 0755)
// 		bs = boltstore.New(dir+"/boltstore/bolt.db", "test")

// 		chain = Loggable(logger, "", New(
// 			ms,
// 			Loggable(logger, "async", Async(
// 				metricsmgr.New("chaintest", nil,
// 					fs,
// 					lrumgr.New(100, bs),
// 				),
// 			)),
// 		))
// 		err = chain.Open()
// 		So(err, ShouldEqual, nil)

// 		Convey("Put/Get/Del", func() {
// 			v := []byte("value")
// 			err = chain.Put("k", v)
// 			So(err, ShouldEqual, nil)

// 			val, err := chain.Get("k")
// 			So(err, ShouldEqual, nil)
// 			So(v, ShouldResemble, v)

// 			val, err = ms.Get("k")
// 			So(err, ShouldEqual, nil)
// 			So(val, ShouldResemble, v)

// 			time.Sleep(10e6) // wait for async operation..

// 			val, err = fs.Get("k")
// 			So(err, ShouldEqual, nil)
// 			So(val, ShouldResemble, v)

// 			val, err = bs.Get("k")
// 			So(err, ShouldEqual, nil)
// 			So(val, ShouldResemble, v)
// 		})
// 	})

// }

/*
c := chainstore.New(
	logmgr.New(l, ""),
	memstore.New(1000),
	chainstore.Async(
		logmgr.New(l, "async"),
		metricsmgr.New(
			"bolt", &metricsmgr.Config{a, b, c},
			batchmgr.New(10),
			lrumgr.New(5000, boltstore.New("/tmp/bolt.db", 0755)),
		),
		metricsmgr.New(
			"s3", &metricsmgr.Config{a, b, c}
			s3store.New("bucket", "u", "p")
		)
	)
)

*/
