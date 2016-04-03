// automatically generated, do not modify

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)
type PrimitiveArray struct {
	_tab flatbuffers.Table
}

func (rcv *PrimitiveArray) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *PrimitiveArray) Type() int8 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt8(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *PrimitiveArray) Encoding() int8 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt8(o + rcv._tab.Pos)
	}
	return 0
}

/// Relative memory offset of the start of the array data excluding the size
/// of the metadata
func (rcv *PrimitiveArray) Offset() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

/// The number of logical values in the array
func (rcv *PrimitiveArray) Length() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

/// The number of observed nulls
func (rcv *PrimitiveArray) NullCount() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(12))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

/// The total size of the actual data in the file
func (rcv *PrimitiveArray) TotalBytes() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(14))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func PrimitiveArrayStart(builder *flatbuffers.Builder) { builder.StartObject(6) }
func PrimitiveArrayAddType(builder *flatbuffers.Builder, dtype int8) { builder.PrependInt8Slot(0, dtype, 0) }
func PrimitiveArrayAddEncoding(builder *flatbuffers.Builder, encoding int8) { builder.PrependInt8Slot(1, encoding, 0) }
func PrimitiveArrayAddOffset(builder *flatbuffers.Builder, offset int64) { builder.PrependInt64Slot(2, offset, 0) }
func PrimitiveArrayAddLength(builder *flatbuffers.Builder, length int64) { builder.PrependInt64Slot(3, length, 0) }
func PrimitiveArrayAddNullCount(builder *flatbuffers.Builder, nullCount int64) { builder.PrependInt64Slot(4, nullCount, 0) }
func PrimitiveArrayAddTotalBytes(builder *flatbuffers.Builder, totalBytes int64) { builder.PrependInt64Slot(5, totalBytes, 0) }
func PrimitiveArrayEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT { return builder.EndObject() }
