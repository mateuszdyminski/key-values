package bolt

import (
	"github.com/mateuszdyminski/key-values/models"
	"testing"
)

func BenchmarkInsertV1(b *testing.B) {
	db := newDB()
	db.setup("/tmp/bolt-test-go", "test")
	defer db.close()

	msg := models.MakeMsg().BytesV1()
	for i := 0; i < b.N; i++ {
		db.Insert(models.MakeID().BytesV1(), msg)
	}
}

func BenchmarkInsertV2(b *testing.B) {
	db := newDB()
	db.setup("/tmp/bolt-test-go", "test")
	defer db.close()

	msg := models.MakeMsg().BytesV1()
	for i := 0; i < b.N; i++ {
		db.Insert(models.MakeID().BytesV2(), msg)
	}
}

func BenchmarkInsertV3(b *testing.B) {
	db := newDB()
	db.setup("/tmp/bolt-test-go", "test")
	defer db.close()

	msg := models.MakeMsg().BytesV1()
	for i := 0; i < b.N; i++ {
		db.Insert(models.MakeID().BytesV3(), msg)
	}
}

func BenchmarkInsertV4(b *testing.B) {
	db := newDB()
	db.setup("/tmp/bolt-test-go", "test")
	defer db.close()

	msg := models.MakeMsg().BytesV1()
	for i := 0; i < b.N; i++ {
		db.Insert(models.MakeID().BytesV4(), msg)
	}
}
