// automatically generated, do not modify

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)
type CategoryMetadata struct {
	_tab flatbuffers.Table
}

func (rcv *CategoryMetadata) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

/// The category codes are presumed to be integers that are valid indexes into
/// the levels array
func (rcv *CategoryMetadata) Levels(obj *PrimitiveArray) *PrimitiveArray {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(PrimitiveArray)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *CategoryMetadata) Ordered() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func CategoryMetadataStart(builder *flatbuffers.Builder) { builder.StartObject(2) }
func CategoryMetadataAddLevels(builder *flatbuffers.Builder, levels flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(levels), 0) }
func CategoryMetadataAddOrdered(builder *flatbuffers.Builder, ordered byte) { builder.PrependByteSlot(1, ordered, 0) }
func CategoryMetadataEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT { return builder.EndObject() }
