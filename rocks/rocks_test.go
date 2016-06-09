package rocks

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/mateuszdyminski/key-values/models"
	"testing"
	"time"
)

func TestInsertGetV1(t *testing.T) {
	db := newDB()
	db.setup("/tmp/rocks-test-go")
	defer db.close()

	msg := models.MakeMsg().BytesV1()
	id := models.MakeID().BytesV1()
	if err := db.Insert(id, msg); err != nil {
		t.Errorf("insert should not return any error. err: %v", err)
	}

	returned, err := db.Get(id)
	if err != nil {
		t.Errorf("get should not return any error. err: %v", err)
	}

	if !bytes.Equal(msg, returned) {
		t.Errorf("data value is different. is(%d) should be(%d)", returned, msg)
	}
}

var slow = flag.Bool("slow", false, "Flag whether run slow tests or not")

func TestInsertTenMillions(t *testing.T) {
	if !*slow {
		t.SkipNow()
	}

	db := newDB()
	db.setup("/tmp/rocks-test-go")
	defer db.close()

	limit := 10000000
	msg := models.MakeMsg().BytesV1()

	fmt.Printf("Msg length: %d\n", len(msg))
	fmt.Printf("Id length: %d\n", len(models.MakeID().BytesV4()))

	now := time.Now()
	for i := 0; i < limit; i++ {
		db.Insert([]byte(models.MakeID().BytesV1()), msg)

		if i%100000 == 0 {
			fmt.Printf("%d messages inserted\n", i)
		}
	}

	fmt.Printf("%d messages inserted in: %v\n", limit, time.Now().Sub(now))

	acc := 0
	it := db.db.NewIterator(db.ro)
	it.SeekToFirst()
	for it = it; it.Valid(); it.Next() {
		acc++
	}
	if err := it.Err(); err != nil {
		t.Errorf("iterating should not return any error. err: %v", err)
	}

	fmt.Printf("total number of elements in rocksdb: %d\n", acc)
}
