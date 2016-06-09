package models

import (
	"testing"
)

func TestIdSerDeserV1(t *testing.T) {
	id := createId()

	ser := id.BytesV1()
	deser, err := IdFromBytes(ser)
	if err != nil {
		t.Fatal("deserialization should not cause any errors. err: %v", err)
	}

	checkId(t, id, *deser)
}

func TestIdSerDeserV2(t *testing.T) {
	id := createId()

	ser := id.BytesV2()
	deser, err := IdFromBytes(ser)
	if err != nil {
		t.Fatal("deserialization should not cause any errors. err: %v", err)
	}

	checkId(t, id, *deser)
}

func checkId(t *testing.T, id, deser Id) {
	if id.Trs != deser.Trs {
		t.Errorf("trs value is different. is(%d) should be(%d)", deser.Trs, id.Trs)
	}
	if id.T != deser.T {
		t.Errorf("t value is different. is(%d) should be(%d)", deser.T, id.T)
	}
	if id.E != deser.E {
		t.Errorf("e value is different. is(%d) should be(%d)", deser.E, id.E)
	}
	if id.Tm != deser.Tm {
		t.Errorf("tm value is different. is(%d) should be(%d)", deser.Tm, id.Tm)
	}
}

func createId() Id {
	id := Id{}
	id.Trs = 2
	id.E = 4
	id.T = 19
	id.Tm = 12

	return id
}
