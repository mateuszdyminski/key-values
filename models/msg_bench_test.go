package models

import "testing"

func BenchmarkMsgToBytesV1(b *testing.B) {
	msg := MakeMsg()

	for i := 0; i < b.N; i++ {
		msg.BytesV1()
	}
}

func BenchmarkMsgToBytesV2(b *testing.B) {
	msg := MakeMsg()

	for i := 0; i < b.N; i++ {
		msg.BytesV2()
	}
}
