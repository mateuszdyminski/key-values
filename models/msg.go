package models

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type Msg struct {
	Data                []byte
	TraceType           int32
	Direction           int32
	InterfaceName       int32
	Protocol            int32
	ProtocolFormat      int32
	ProtocolMessageType int32
}

func (m *Msg) BytesV1() []byte {
	arr := make([]byte, 24+len(m.Data), 24+len(m.Data))

	// littleEndian
	arr[0] = byte(m.TraceType)
	arr[1] = byte(m.TraceType >> 8)
	arr[2] = byte(m.TraceType >> 16)
	arr[3] = byte(m.TraceType >> 24)

	arr[4] = byte(m.Direction)
	arr[5] = byte(m.Direction >> 8)
	arr[6] = byte(m.Direction >> 16)
	arr[7] = byte(m.Direction >> 24)

	arr[8] = byte(m.InterfaceName)
	arr[9] = byte(m.InterfaceName >> 8)
	arr[10] = byte(m.InterfaceName >> 16)
	arr[11] = byte(m.InterfaceName >> 24)

	arr[12] = byte(m.Protocol)
	arr[13] = byte(m.Protocol >> 8)
	arr[14] = byte(m.Protocol >> 16)
	arr[15] = byte(m.Protocol >> 24)

	arr[16] = byte(m.ProtocolFormat)
	arr[17] = byte(m.ProtocolFormat >> 8)
	arr[18] = byte(m.ProtocolFormat >> 16)
	arr[19] = byte(m.ProtocolFormat >> 24)

	arr[20] = byte(m.ProtocolMessageType)
	arr[21] = byte(m.ProtocolMessageType >> 8)
	arr[22] = byte(m.ProtocolMessageType >> 16)
	arr[23] = byte(m.ProtocolMessageType >> 24)

	for i := 0; i < len(m.Data); i++ {
		arr[24+i] = m.Data[i]
	}

	return arr
}

func (m *Msg) BytesV2() []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, m.TraceType)
	binary.Write(buf, binary.LittleEndian, m.Direction)
	binary.Write(buf, binary.LittleEndian, m.InterfaceName)
	binary.Write(buf, binary.LittleEndian, m.Protocol)
	binary.Write(buf, binary.LittleEndian, m.ProtocolFormat)
	binary.Write(buf, binary.LittleEndian, m.ProtocolMessageType)
	binary.Write(buf, binary.LittleEndian, m.Data)

	return buf.Bytes()
}

var payload []byte = []byte("00010060000008000000020000000100010000024002000000420004000000000035000E00001B00090100616263640000000068000100006B0004200080000028002100AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")

func MakeMsg() *Msg {
	return &Msg{
		Data:                payload,
		TraceType:           1,
		Direction:           2,
		InterfaceName:       3,
		Protocol:            4,
		ProtocolFormat:      5,
		ProtocolMessageType: 6,
	}
}

func MsgFromBytes(data []byte) (*Msg, error) {
	if len(data) < 24 {
		return nil, fmt.Errorf("wrong payload length(%d)", len(data))
	}
	buf := bytes.NewReader(data)
	var msg Msg

	if err := binary.Read(buf, binary.LittleEndian, &msg.TraceType); err != nil {
		return nil, err
	}

	if err := binary.Read(buf, binary.LittleEndian, &msg.Direction); err != nil {
		return nil, err
	}

	if err := binary.Read(buf, binary.LittleEndian, &msg.InterfaceName); err != nil {
		return nil, err
	}

	if err := binary.Read(buf, binary.LittleEndian, &msg.Protocol); err != nil {
		return nil, err
	}

	if err := binary.Read(buf, binary.LittleEndian, &msg.ProtocolFormat); err != nil {
		return nil, err
	}

	if err := binary.Read(buf, binary.LittleEndian, &msg.ProtocolMessageType); err != nil {
		return nil, err
	}

	if err := binary.Read(buf, binary.LittleEndian, &msg.Data); err != nil {
		return nil, err
	}

	return &msg, nil
}
