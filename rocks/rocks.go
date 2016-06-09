package rocks

import (
	"github.com/tecbot/gorocksdb"
	"log"
	"os"
)

type rocks struct {
	path string
	db   *gorocksdb.DB
	wo   *gorocksdb.WriteOptions
	ro   *gorocksdb.ReadOptions
}

func newDB() *rocks {
	return &rocks{}
}

func (i *rocks) setup(path string) {
	i.path = path
	opts := gorocksdb.NewDefaultOptions()
	opts.SetCreateIfMissing(true)

	var err error
	i.db, err = gorocksdb.OpenDb(opts, i.path)
	if err != nil {
		log.Fatal(err)
	}

	i.wo = gorocksdb.NewDefaultWriteOptions()
	i.ro = gorocksdb.NewDefaultReadOptions()
}

func (i *rocks) Insert(key []byte, msg []byte) error {
	return i.db.Put(i.wo, key, msg)
}

func (i *rocks) Get(key []byte) ([]byte, error) {
	data, err := i.db.Get(i.ro, key)
	if err != nil {
		return nil, err
	}

	return data.Data(), nil
}

func (i *rocks) close() {
	i.wo.Destroy()
	i.ro.Destroy()
	i.db.Close()
	os.RemoveAll(i.path)
}
