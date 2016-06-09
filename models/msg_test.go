package models

import (
	"bytes"
	"testing"
)

func TestMsgSerDeserV1(t *testing.T) {
	msg := createMsg()

	ser := msg.BytesV1()
	deser, err := MsgFromBytes(ser)
	if err != nil {
		t.Fatal("deserialization should not cause any errors. err: %v", err)
	}

	checkMsg(t, msg, *deser)
}

func TestMsgSerDeserV2(t *testing.T) {
	msg := createMsg()

	ser := msg.BytesV2()
	deser, err := MsgFromBytes(ser)
	if err != nil {
		t.Fatal("deserialization should not cause any errors. err: %v", err)
	}

	checkMsg(t, msg, *deser)
}

func createMsg() Msg {
	msg := Msg{}
	msg.Data = []byte("some content")
	msg.Direction = 14
	msg.InterfaceName = 19
	msg.Protocol = 12
	msg.ProtocolFormat = 45
	msg.ProtocolMessageType = 99
	msg.TraceType = 49

	return msg
}

func checkMsg(t *testing.T, msg, deser Msg) {
	if msg.Direction != deser.Direction {
		t.Errorf("direction value is different. is(%d) should be(%d)", deser.Direction, msg.Direction)
	}
	if msg.InterfaceName != deser.InterfaceName {
		t.Errorf("interfaceName value is different. is(%d) should be(%d)", deser.InterfaceName, msg.InterfaceName)
	}
	if msg.Protocol != deser.Protocol {
		t.Errorf("protocol value is different. is(%d) should be(%d)", deser.Protocol, msg.Protocol)
	}
	if msg.ProtocolFormat != deser.ProtocolFormat {
		t.Errorf("protocolFormat value is different. is(%d) should be(%d)", deser.ProtocolFormat, msg.ProtocolFormat)
	}
	if msg.ProtocolMessageType != deser.ProtocolMessageType {
		t.Errorf("protocolMessageType value is different. is(%d) should be(%d)", deser.ProtocolMessageType, msg.ProtocolMessageType)
	}
	if msg.TraceType != deser.TraceType {
		t.Errorf("traceType value is different. is(%d) should be(%d)", deser.TraceType, msg.TraceType)
	}
	if bytes.Equal(msg.Data, deser.Data) {
		t.Errorf("data value is different. is(%d) should be(%d)", deser.Data, msg.Data)
	}
}
