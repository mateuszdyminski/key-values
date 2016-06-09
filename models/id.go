package models

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"
	"time"
)

type Id struct {
	Trs int32
	T   int32
	E   int64
	Tm  int64
}

func (i *Id) BytesV1() []byte {
	arr := make([]byte, 24, 24)

	// littleEndian
	arr[0] = byte(i.Tm)
	arr[1] = byte(i.Tm >> 8)
	arr[2] = byte(i.Tm >> 16)
	arr[3] = byte(i.Tm >> 24)
	arr[4] = byte(i.Tm >> 32)
	arr[5] = byte(i.Tm >> 40)
	arr[6] = byte(i.Tm >> 48)
	arr[7] = byte(i.Tm >> 56)

	arr[8] = byte(i.E)
	arr[9] = byte(i.E >> 8)
	arr[10] = byte(i.E >> 16)
	arr[11] = byte(i.E >> 24)
	arr[12] = byte(i.E >> 32)
	arr[13] = byte(i.E >> 40)
	arr[14] = byte(i.E >> 48)
	arr[15] = byte(i.E >> 56)

	arr[16] = byte(i.Trs)
	arr[17] = byte(i.Trs >> 8)
	arr[18] = byte(i.Trs >> 16)
	arr[19] = byte(i.Trs >> 24)

	arr[20] = byte(i.T)
	arr[21] = byte(i.T >> 8)
	arr[22] = byte(i.T >> 16)
	arr[23] = byte(i.T >> 24)

	return arr
}

func (i *Id) BytesV2() []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.LittleEndian, i.Tm)
	binary.Write(buf, binary.LittleEndian, i.E)
	binary.Write(buf, binary.LittleEndian, i.Trs)
	binary.Write(buf, binary.LittleEndian, i.T)

	return buf.Bytes()
}

func (i *Id) BytesV3() []byte {
	return []byte(fmt.Sprintf("%d#%d#%d#%d", i.Tm, i.E, i.Trs, i.T))
}

func (i *Id) BytesV4() []byte {
	buf := &bytes.Buffer{}
	buf.WriteString(strconv.FormatInt(i.Tm, 10))
	buf.WriteByte('#')
	buf.WriteString(strconv.FormatInt(i.E, 10))
	buf.WriteByte('#')
	buf.WriteString(strconv.FormatInt(int64(i.Trs), 10))
	buf.WriteByte('#')
	buf.WriteString(strconv.FormatInt(int64(i.T), 10))

	return buf.Bytes()
}

func IdFromBytes(data []byte) (*Id, error) {
	if len(data) != 24 {
		return nil, fmt.Errorf("Wrong pyload length")
	}
	buf := bytes.NewReader(data)
	var id Id

	if err := binary.Read(buf, binary.LittleEndian, &id.Tm); err != nil {
		return nil, err
	}

	if err := binary.Read(buf, binary.LittleEndian, &id.E); err != nil {
		return nil, err
	}

	if err := binary.Read(buf, binary.LittleEndian, &id.Trs); err != nil {
		return nil, err
	}

	if err := binary.Read(buf, binary.LittleEndian, &id.T); err != nil {
		return nil, err
	}

	return &id, nil
}

func MakeID() *Id {
	return &Id{
		Trs: 678,
		T:   980,
		E:   12345,
		Tm:  time.Now().UnixNano(),
	}
}
