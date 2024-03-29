// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package packets

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

type Vec2 struct {
	_tab flatbuffers.Struct
}

func (rcv *Vec2) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Vec2) Table() flatbuffers.Table {
	return rcv._tab.Table
}

func (rcv *Vec2) X() float32 {
	return rcv._tab.GetFloat32(rcv._tab.Pos + flatbuffers.UOffsetT(0))
}
func (rcv *Vec2) MutateX(n float32) bool {
	return rcv._tab.MutateFloat32(rcv._tab.Pos+flatbuffers.UOffsetT(0), n)
}

func (rcv *Vec2) Y() float32 {
	return rcv._tab.GetFloat32(rcv._tab.Pos + flatbuffers.UOffsetT(4))
}
func (rcv *Vec2) MutateY(n float32) bool {
	return rcv._tab.MutateFloat32(rcv._tab.Pos+flatbuffers.UOffsetT(4), n)
}

func CreateVec2(builder *flatbuffers.Builder, x float32, y float32) flatbuffers.UOffsetT {
	builder.Prep(4, 8)
	builder.PrependFloat32(y)
	builder.PrependFloat32(x)
	return builder.Offset()
}
