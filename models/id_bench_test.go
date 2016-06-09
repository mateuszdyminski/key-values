package models

import (
	"testing"
)

var result []byte

func BenchmarkIdToBytesV1(b *testing.B) {
	id := MakeID()

	for i := 0; i < b.N; i++ {
		result = id.BytesV1()
	}
}

func BenchmarkIdToBytesV2(b *testing.B) {
	id := MakeID()

	for i := 0; i < b.N; i++ {
		result = id.BytesV2()
	}
}

func BenchmarkIdToBytesV3(b *testing.B) {
	id := MakeID()

	for i := 0; i < b.N; i++ {
		result = id.BytesV3()
	}
}

func BenchmarkIdToBytesV4(b *testing.B) {
	id := MakeID()

	for i := 0; i < b.N; i++ {
		result = id.BytesV4()
	}
}
