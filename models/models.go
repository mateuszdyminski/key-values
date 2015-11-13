package models

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type Msg struct {
	data                []byte
	traceType           int
	direction           int
	interfaceName       int
	protocol            int
	protocolFormat      int
	protocolMessageType int
}

func (m *Msg) Bytes() []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, m.traceType)
	binary.Write(buf, binary.LittleEndian, m.direction)
	binary.Write(buf, binary.LittleEndian, m.interfaceName)
	binary.Write(buf, binary.LittleEndian, m.protocol)
	binary.Write(buf, binary.LittleEndian, m.protocolFormat)
	binary.Write(buf, binary.LittleEndian, m.protocolMessageType)
	binary.Write(buf, binary.LittleEndian, m.data)

	return buf.Bytes()
}

func MakeMsg() *Msg {
	return &Msg{
		data:                []byte("00010060000008000000020000000100010000024002000000420004000000000035000E00001B00090100616263640000000068000100006B0004200080000028002100AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"),
		traceType:           1,
		direction:           2,
		interfaceName:       3,
		protocol:            4,
		protocolFormat:      5,
		protocolMessageType: 6,
	}
}

type Id struct {
	trs uint32
	t   uint32
	e   uint64
	tm  uint64
}

func (i *Id) String() string {
	return fmt.Sprintf("%d#%d#%d#%d", i.trs, i.t, i.e, i.tm)
}

func (i *Id) Bytes() []byte {
	arr := make([]byte, 24)

	// littleEndian
	arr[0] = byte(i.trs)
	arr[1] = byte(i.trs >> 8)
	arr[2] = byte(i.trs >> 16)
	arr[3] = byte(i.trs >> 24)

	arr[4] = byte(i.t)
	arr[5] = byte(i.t >> 8)
	arr[6] = byte(i.t >> 16)
	arr[7] = byte(i.t >> 24)

	arr[8] = byte(i.e)
	arr[9] = byte(i.e >> 8)
	arr[10] = byte(i.e >> 16)
	arr[11] = byte(i.e >> 24)
	arr[12] = byte(i.e >> 32)
	arr[13] = byte(i.e >> 40)
	arr[14] = byte(i.e >> 48)
	arr[15] = byte(i.e >> 56)

	arr[16] = byte(i.tm)
	arr[17] = byte(i.tm >> 8)
	arr[18] = byte(i.tm >> 16)
	arr[19] = byte(i.tm >> 24)
	arr[20] = byte(i.tm >> 32)
	arr[21] = byte(i.tm >> 40)
	arr[22] = byte(i.tm >> 48)
	arr[23] = byte(i.tm >> 56)

	return arr
}

func MakeID(i int) *Id {
	return &Id{
		trs: uint32(i),
		t:   980,
		e:   15698745,
		tm:  12345678,
	}
}
