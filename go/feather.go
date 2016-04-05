package feather

import (
	"./fbs"
	"bytes"
	"os"
	// "log"
	"encoding/binary"
	"errors"
	// flatbuffers "github.com/google/flatbuffers/go"
)

type Frame struct {
	File    *os.File
	FSize   int64
	NumRows int64
	Cols    []string // to preserve order
	Meta    map[string]*Column
}

type Column struct {
	Name       string
	Type       ctype // TODO: add a stringed version for further use
	Length     int64
	Offset     int64
	TotalBytes int64
	NullCount  int64
	NullMap    []bool // store bitarray of nulls (TODO: return interface{} with nils instead?)
	// TODO: this only gets populated when column gets read
	Encoding int8
	// Metadata
	// MetadataType
}

// column types (as saved in metadata)
type ctype int8

const (
	T_BOOL ctype = iota
	T_INT8
	T_INT16
	T_INT32
	T_INT64
	T_UINT8
	T_UINT16
	T_UINT32
	T_UINT64
	T_FLOAT
	T_DOUBLE
	T_UTF8
	T_BINARY
	T_CATEGORY
	T_TIMESTAMP
	T_DATE
	T_TIME
)

func Open(fn string) (*Frame, error) {
	fr := new(Frame)
	f, err := os.Open(fn)
	if err != nil {
		return new(Frame), err
	}
	st, _ := f.Stat()
	fr.FSize = st.Size()

	hd := make([]byte, 4)
	ft := make([]byte, 4)
	f.Read(hd)
	f.ReadAt(ft, fr.FSize-4)

	mby := "FEA1" // required magic string

	if string(hd) != mby {
		return nil, errors.New("Header incorrect")
	}
	if string(ft) != mby {
		return nil, errors.New("Footer incorrect")
	}

	fr.File = f
	fr.Meta = make(map[string]*Column, 0)

	fr.readMetadata()

	return fr, nil
}

// TODO: returns?
func (fr *Frame) readMetadata() {
	mtsb := make([]byte, 4)            // metadata size
	fr.File.ReadAt(mtsb, fr.FSize-4-4) // minus magic bytes and metadata size

	var mts int32
	binary.Read(bytes.NewBuffer(mtsb), binary.LittleEndian, &mts)

	mtd := make([]byte, mts)
	fr.File.ReadAt(mtd, fr.FSize-4-4-int64(mts))

	meta := fbs.GetRootAsCTable(mtd, 0)
	// fmt.Println(meta.ColumnsLength())
	fr.NumRows = meta.NumRows()
	// fmt.Println(meta.Description())

	nc := meta.ColumnsLength()
	for j := 0; j < nc; j++ {
		cl := new(fbs.Column)
		blr := meta.Columns(cl, j)
		_ = blr
		vl := new(fbs.PrimitiveArray)
		cl.Values(vl)

		nm := string(cl.Name())
		cln := Column{
			Name:       nm,
			Type:       ctype(vl.Type()),
			Length:     vl.Length(),
			Offset:     vl.Offset(),
			TotalBytes: vl.TotalBytes(),
			NullCount:  vl.NullCount(),
			Encoding:   vl.Encoding(),
		}

		fr.Meta[nm] = &cln
		fr.Cols = append(fr.Cols, nm) // TODO: buffer this slice, marginal speedup

		if cl.MetadataType() != 0 {
			panic("Column metadata for category/timestamp/date/time not supported")
		}
		// TypeMetadataNONE = 0
		// TypeMetadataCategoryMetadata = 1
		// TypeMetadataTimestampMetadata = 2
		// TypeMetadataDateMetadata = 3
		// TypeMetadataTimeMetadata = 4
	}
}

func (fr *Frame) Read(cl string) interface{} {
	cln, ok := fr.Meta[cl]
	if !ok {
		panic("Column not found")
	}

	ncb := int64(0) // bytes occupied by the bitmask
	if cln.NullCount != 0 {
		ncb = cln.Length / 8
		if cln.Length%8 > 0 {
			ncb += 1
		}

		bt := make([]byte, ncb)
		fr.File.ReadAt(bt, cln.Offset)
		cln.NullMap = make([]bool, cln.Length)

		maxb := 8
		for j, b := range bt {
			// handling a loose byte
			if int64(j) == ncb-1 && cln.Length%8 > 0 {
				maxb = int(cln.Length % 8)
			}

			for k := 0; k < maxb; k++ {
				cln.NullMap[j*8+k] = b != b|uint8(1)<<uint8(k)
			}
		}

	}
	if cln.Encoding != 0 {
		panic("Can't do dictionaries just yet") // TODO
	}

	bt := make([]byte, cln.TotalBytes-ncb)
	fr.File.ReadAt(bt, cln.Offset+ncb)

	buf := bytes.NewBuffer(bt)

	switch cln.Type {
	case T_BOOL:
		ret := make([]bool, cln.Length)
		for j, b := range bt {
			// TODO: will fail on a loose bit array
			// OPTIM: trivial to unroll - test performance
			for k := 0; k < 8; k++ {
				ret[j*8+k] = b == b|uint8(1)<<uint8(k)
			}
		}
		return ret

	// TODO: recast this slice into the target type
	//		 involves unsafe changes to slice properties
	case T_INT8:
		ret := make([]int8, cln.Length)
		binary.Read(buf, binary.LittleEndian, &ret)
		return ret
	case T_INT16:
		ret := make([]int16, cln.Length)
		binary.Read(buf, binary.LittleEndian, &ret)
		return ret
	case T_INT32:
		ret := make([]int32, cln.Length)
		binary.Read(buf, binary.LittleEndian, &ret)
		return ret
	case T_INT64:
		ret := make([]int64, cln.Length)
		binary.Read(buf, binary.LittleEndian, &ret)
		return ret

	case T_UINT8:
		ret := make([]uint8, cln.Length)
		binary.Read(buf, binary.LittleEndian, &ret)
		return ret
	case T_UINT16:
		ret := make([]uint16, cln.Length)
		binary.Read(buf, binary.LittleEndian, &ret)
		return ret
	case T_UINT32:
		ret := make([]uint32, cln.Length)
		binary.Read(buf, binary.LittleEndian, &ret)
		return ret
	case T_UINT64:
		ret := make([]uint64, cln.Length)
		binary.Read(buf, binary.LittleEndian, &ret)
		return ret

	case T_FLOAT:
		ret := make([]float32, cln.Length)
		binary.Read(buf, binary.LittleEndian, &ret)
		return ret
	case T_DOUBLE:
		ret := make([]float64, cln.Length)
		binary.Read(buf, binary.LittleEndian, &ret)
		return ret

	case T_UTF8:
		// offsets
		off := make([]uint32, cln.Length+1)
		binary.Read(buf, binary.LittleEndian, &off)

		// actual strings
		ret := make([]string, cln.Length)

		bt = bt[4*(cln.Length+1):] // cut off the offsets

		for j := 0; j < int(cln.Length); j++ {
			ret[j] = string(bt[off[j]:off[j+1]])
		}

		return ret

	// case T_BINARY:

	// case T_CATEGORY
	// case T_TIMESTAMP
	// case T_DATE
	// case T_TIME
	default:
		panic("Unsupported format") // TODO

	}

	return struct{}{}
}

func (fr *Frame) ReadBool(cl string) []bool {
	arr := fr.Read(cl)
	return arr.([]bool)
}
func (fr *Frame) ReadDouble(cl string) []float64 {
	arr := fr.Read(cl)
	return arr.([]float64)
}
func (fr *Frame) ReadString(cl string) []string {
	arr := fr.Read(cl)
	return arr.([]string)
}

// TODO: returns?
func (f *Frame) Close() {
	f.File.Close()
}
