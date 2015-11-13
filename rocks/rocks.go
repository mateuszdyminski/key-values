package rocks

import (
	"fmt"
	"os"
	"time"

	"github.com/influxdb/rocksdb"
	"github.com/mateuszdyminski/key-values/models"
)

func Insert(insertsNum int) {
	// initialization
	msg := models.MakeMsg()
	msgBytes := msg.Bytes()

	opts := rocksdb.NewOptions()
	bo := rocksdb.NewBlockBasedOptions()
	bo.SetCache(rocksdb.NewLRUCache(3 << 30))
	opts.SetBlockBasedTableFactory(bo)
	opts.SetCreateIfMissing(true)
	db, err := rocksdb.Open("/tmp/rocks-test", opts)
	if err != nil {
		panic("can't open db")
	}

	wo := rocksdb.NewWriteOptions()

	defer func() {
		wo.Close()
		db.Close()
		os.RemoveAll("/tmp/rocks-test")
	}()

	// test
	t := time.Now()
	for i := 0; i < insertsNum; i++ {
		db.Put(wo, models.MakeID(i).Bytes(), msgBytes)
	}

	// result
	fmt.Printf("Rocks db inserted %d messages in: %s \n", insertsNum, time.Since(t))
}
