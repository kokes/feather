package feather

import (
	"testing"
)

// TODO: get feather files from the test suite in the
//       original repo

func TestOpen(t *testing.T) {
	fns := []string{
		"../testdata/minwage.fth",
		"../testdata/gdp-level.fth",
	}
	for _, fn := range fns {
		_, err := Open(fn)
		if err != nil {
			t.Error(err)
		}
	}
}

func TestMetadata(t *testing.T) {
	fn := "../testdata/minwage.fth"
	f, _ := Open(fn)

	if f.NumRows != 160 {
		t.Errorf("Incorrect number of rows")
	}
}

func TestBool(t *testing.T) {
	fn := "../testdata/minwage.fth"
	f, _ := Open(fn)

	r := f.ReadBool("mw")

	vl := map[int]bool{
		0:   true,
		20: false,
		34: false,
		91: false,
		99: false,
		100: false,
		143: true,
		159: true,
	}

	for j, v := range vl {
		if r[j] != v {
			t.Errorf("Unexpected value")
		}
	}
}

func TestDouble(t *testing.T) {
	fn := "../testdata/minwage.fth"
	f, _ := Open(fn)

	r := f.ReadDouble("unem")

	vl := map[int]float64{
		0:   8.5,
		159: 4.1,
	}

	for j, v := range vl {
		if r[j] != v {
			t.Errorf("Unexpected value")
		}
	}
}

func TestUTF(t *testing.T) {
	fn := "../testdata/minwage.fth"
	f, _ := Open(fn)

	r := f.Read("c").([]string)

	vl := map[int]string{
		0:   "AFG",
		159: "ZWE",
	}

	for j, v := range vl {
		if r[j] != v {
			t.Errorf("Unexpected value")
		}
	}
}
