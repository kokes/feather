// automatically generated, do not modify

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)
type TimestampMetadata struct {
	_tab flatbuffers.Table
}

func (rcv *TimestampMetadata) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *TimestampMetadata) Unit() int8 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt8(o + rcv._tab.Pos)
	}
	return 0
}

/// Timestamp data is assumed to be UTC, but the time zone is stored here for
/// presentation as localized
func (rcv *TimestampMetadata) Timezone() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func TimestampMetadataStart(builder *flatbuffers.Builder) { builder.StartObject(2) }
func TimestampMetadataAddUnit(builder *flatbuffers.Builder, unit int8) { builder.PrependInt8Slot(0, unit, 0) }
func TimestampMetadataAddTimezone(builder *flatbuffers.Builder, timezone flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(timezone), 0) }
func TimestampMetadataEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT { return builder.EndObject() }
