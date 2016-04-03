// automatically generated, do not modify

package fbs

import (
	flatbuffers "github.com/google/flatbuffers/go"
)
type CTable struct {
	_tab flatbuffers.Table
}

func GetRootAsCTable(buf []byte, offset flatbuffers.UOffsetT) *CTable {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &CTable{}
	x.Init(buf, n + offset)
	return x
}

func (rcv *CTable) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

/// Some text (or a name) metadata about what the file is, optional
func (rcv *CTable) Description() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *CTable) NumRows() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *CTable) Columns(obj *Column, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
	if obj == nil {
		obj = new(Column)
	}
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *CTable) ColumnsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func CTableStart(builder *flatbuffers.Builder) { builder.StartObject(3) }
func CTableAddDescription(builder *flatbuffers.Builder, description flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(description), 0) }
func CTableAddNumRows(builder *flatbuffers.Builder, numRows int64) { builder.PrependInt64Slot(1, numRows, 0) }
func CTableAddColumns(builder *flatbuffers.Builder, columns flatbuffers.UOffsetT) { builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(columns), 0) }
func CTableStartColumnsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT { return builder.StartVector(4, numElems, 4)
}
func CTableEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT { return builder.EndObject() }
