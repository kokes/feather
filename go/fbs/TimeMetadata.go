// automatically generated, do not modify

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)
type TimeMetadata struct {
	_tab flatbuffers.Table
}

func (rcv *TimeMetadata) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *TimeMetadata) Unit() int8 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt8(o + rcv._tab.Pos)
	}
	return 0
}

func TimeMetadataStart(builder *flatbuffers.Builder) { builder.StartObject(1) }
func TimeMetadataAddUnit(builder *flatbuffers.Builder, unit int8) { builder.PrependInt8Slot(0, unit, 0) }
func TimeMetadataEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT { return builder.EndObject() }
