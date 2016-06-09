package bolt

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
	"os"
)

type boltDb struct {
	path       string
	bucketName string
	db         *bolt.DB
	bucket     *bolt.Bucket
}

func newDB() *boltDb {
	return &boltDb{}
}

func (i *boltDb) setup(path, bucketName string) {
	i.bucketName = bucketName
	i.path = path

	var err error
	i.db, err = bolt.Open(i.path, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	err = i.db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucket([]byte(i.bucketName))
		return err
	})
	if err != nil {
		log.Fatal(err)
	}
}

func (i *boltDb) Insert(key []byte, msg []byte) error {
	return i.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(i.bucketName))
		return b.Put(key, msg)
	})
}

func (i *boltDb) Get(key []byte) ([]byte, error) {
	var res []byte

	i.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(i.bucketName))
		res = b.Get(key)

		return nil
	})

	if res == nil {
		return nil, fmt.Errorf("no such key in db. key: %v", key)
	}

	return res, nil
}

func (i *boltDb) close() {
	i.db.Close()
	os.Remove(i.path)
}
