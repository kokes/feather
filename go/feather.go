package main

import (
	"bytes"
	"fmt"
	"os"
	//"io/ioutil"
	"./fbs"
	"encoding/binary"
	"errors"
	//	flatbuffers "github.com/google/flatbuffers/go"
)

type Frame struct {
	File *os.File
	FSize int64
	Cols []string
	Meta map[string]*Column
}

type Column struct {
	Name string
	Type int8
	Length int
	Offset int
	TotalBytes int
	NullCount int
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

	fr.readMetadata()

	return fr, nil
}

// TODO: returns?
func (fr *Frame) readMetadata() {
	mtsb := make([]byte, 4) // metadata size
	fr.File.ReadAt(mtsb, fr.FSize-4-4)  // minus magic bytes and metadata size

	var mts int32
	binary.Read(bytes.NewBuffer(mtsb), binary.LittleEndian, &mts)

	mtd := make([]byte, mts)
	fr.File.ReadAt(mtd, fr.FSize-4-4-int64(mts))

	meta := fbs.GetRootAsCTable(mtd, 0)
	// fmt.Println(meta.ColumnsLength())
	// fmt.Println(meta.NumRows())
	// fmt.Println(meta.Description())

	nc := meta.ColumnsLength()
	for j := 0; j < nc; j++ {
		cl := new(fbs.Column)
		blr := meta.Columns(cl, j)
		_ = blr
		vl := new(fbs.PrimitiveArray)
		fmt.Println(string(cl.Name()))
		cl.Values(vl)
		
		fmt.Println(vl.Type()) // up above in constants
		fmt.Println(vl.Length())
		fmt.Println(vl.NullCount())
		fmt.Println(vl.Offset())
		fmt.Println(vl.TotalBytes())

		fmt.Println(cl.MetadataType())
		// TypeMetadataNONE = 0
		// TypeMetadataCategoryMetadata = 1
		// TypeMetadataTimestampMetadata = 2
		// TypeMetadataDateMetadata = 3
		// TypeMetadataTimeMetadata = 4



	}
}

// TODO: returns?
func (f *Frame) Close() {
	f.File.Close()
}

func main() {

	fn := "../testdata/minwage.fth"

	f, _ := Open(fn)
	defer f.Close()





}
