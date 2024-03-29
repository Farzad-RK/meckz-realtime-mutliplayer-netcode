// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package conPackets

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type ConPacket struct {
	_tab flatbuffers.Table
}

func GetRootAsConPacket(buf []byte, offset flatbuffers.UOffsetT) *ConPacket {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &ConPacket{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *ConPacket) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *ConPacket) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *ConPacket) Type() Type {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt8(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *ConPacket) MutateType(n Type) bool {
	return rcv._tab.MutateInt8Slot(4, n)
}

func (rcv *ConPacket) SequenceNumber() uint32 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint32(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *ConPacket) MutateSequenceNumber(n uint32) bool {
	return rcv._tab.MutateUint32Slot(6, n)
}

func (rcv *ConPacket) Token() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *ConPacket) EncryptionKey() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func ConPacketStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func ConPacketAddType(builder *flatbuffers.Builder, type_ int8) {
	builder.PrependInt8Slot(0, type_, 0)
}
func ConPacketAddSequenceNumber(builder *flatbuffers.Builder, sequenceNumber uint32) {
	builder.PrependUint32Slot(1, sequenceNumber, 0)
}
func ConPacketAddToken(builder *flatbuffers.Builder, token flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(token), 0)
}
func ConPacketAddEncryptionKey(builder *flatbuffers.Builder, encryptionKey flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(encryptionKey), 0)
}
func ConPacketEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
