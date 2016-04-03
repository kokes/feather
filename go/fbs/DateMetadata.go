// automatically generated, do not modify

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)
type DateMetadata struct {
	_tab flatbuffers.Table
}

func (rcv *DateMetadata) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func DateMetadataStart(builder *flatbuffers.Builder) { builder.StartObject(0) }
func DateMetadataEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT { return builder.EndObject() }
