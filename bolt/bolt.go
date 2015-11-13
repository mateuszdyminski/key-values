package bolt

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/mateuszdyminski/key-values/models"
	"log"
	"os"
	"time"
)

func Insert(insertsNum int) {
	// init
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		os.Remove("my.db")
		db.Close()
	}()

	msg := models.MakeMsg()
	msgBytes := msg.Bytes()

	// when
	t := time.Now()
	db.Batch(func(tx *bolt.Tx) error {
		b, e := tx.CreateBucket([]byte("test"))
		if e != nil {
			fmt.Printf("Can't create bucket: %v", e)
			return e
		}

		for i := 0; i < insertsNum; i++ {
			b.Put(models.MakeID(i).Bytes(), msgBytes)
		}

		return nil
	})
	elapsed := time.Since(t)

	// Result
	fmt.Printf("Bolt db inserted %d messages in: %s\n", insertsNum, elapsed)
}
