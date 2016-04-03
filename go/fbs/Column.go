// automatically generated, do not modify

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)
type Column struct {
	_tab flatbuffers.Table
}

func (rcv *Column) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Column) Name() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *Column) Values(obj *PrimitiveArray) *PrimitiveArray {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
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

func (rcv *Column) MetadataType() byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetByte(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Column) Metadata(obj *flatbuffers.Table) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		rcv._tab.Union(obj, o)
		return true
	}
	return false
}

/// This should (probably) be JSON
func (rcv *Column) UserMetadata() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func ColumnStart(builder *flatbuffers.Builder) { builder.StartObject(5) }
func ColumnAddName(builder *flatbuffers.Builder, name flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(name), 0) }
func ColumnAddValues(builder *flatbuffers.Builder, values flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(values), 0) }
func ColumnAddMetadataType(builder *flatbuffers.Builder, metadataType byte) { builder.PrependByteSlot(2, metadataType, 0) }
func ColumnAddMetadata(builder *flatbuffers.Builder, metadata flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(metadata), 0) }
func ColumnAddUserMetadata(builder *flatbuffers.Builder, userMetadata flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(4, flatbuffers.UOffsetT(userMetadata), 0) }
func ColumnEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT { return builder.EndObject() }
