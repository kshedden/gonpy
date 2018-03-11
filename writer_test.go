package gonpy

import (
	"os"
	"testing"
)

func Test_float64(t *testing.T) {

	data := []float64{0, 1, 2, 3, 4, 5, 6, 7}

	for shape := 0; shape < 4; shape++ {

		wtr, err := NewFileWriter("data/tmp.npy")
		if err != nil {
			panic(err)
		}

		switch shape {
		case 0:
			wtr.Shape = []int{4, 2}
		case 1:
			wtr.Shape = []int{8, 1}
		case 2:
			wtr.Shape = []int{1, 8}
		case 3:
			wtr.Shape = nil
		}

		err = wtr.WriteFloat64(data)
		if err != nil {
			panic(err)
		}

		r, err := os.Open("data/tmp.npy")
		if err != nil {
			panic(err)
		}
		rdr, err := NewReader(r)
		if err != nil {
			panic(err)
		}
		data, err = rdr.GetFloat64()
		if err != nil {
			panic(err)
		}
		if !check_float64(data) {
			t.Fail()
		}
	}
}

func Test_float32(t *testing.T) {

	data := []float32{0, 1, 2, 3, 4, 5, 6, 7}

	wtr, err := NewFileWriter("data/tmp.npy")
	if err != nil {
		panic(err)
	}
	err = wtr.WriteFloat32(data)
	if err != nil {
		panic(err)
	}

	r, err := os.Open("data/tmp.npy")
	if err != nil {
		panic(err)
	}
	rdr, err := NewReader(r)
	if err != nil {
		panic(err)
	}
	data, err = rdr.GetFloat32()
	if err != nil {
		panic(err)
	}
	if !check_float32(data) {
		t.Fail()
	}
}

func Test_int64(t *testing.T) {

	data := []int64{0, 1, 2, 3, 4, 5, 6, 7}

	wtr, err := NewFileWriter("data/tmp.npy")
	if err != nil {
		panic(err)
	}
	err = wtr.WriteInt64(data)
	if err != nil {
		panic(err)
	}

	fid, err := os.Open("data/tmp.npy")
	if err != nil {
		panic(err)
	}
	rdr, err := NewReader(fid)
	if err != nil {
		panic(err)
	}
	data, err = rdr.GetInt64()
	if err != nil {
		panic(err)
	}
	if !check_int64(data) {
		t.Fail()
	}
}

func Test_int32(t *testing.T) {

	data := []int32{0, 1, 2, 3, 4, 5, 6, 7}

	wtr, err := NewFileWriter("data/tmp.npy")
	if err != nil {
		panic(err)
	}
	err = wtr.WriteInt32(data)
	if err != nil {
		panic(err)
	}

	r, err := os.Open("data/tmp.npy")
	if err != nil {
		panic(err)
	}
	rdr, err := NewReader(r)
	if err != nil {
		panic(err)
	}
	data, err = rdr.GetInt32()
	if err != nil {
		panic(err)
	}
	if !check_int32(data) {
		t.Fail()
	}
}

func Test_int16(t *testing.T) {

	data := []int16{0, 1, 2, 3, 4, 5, 6, 7}

	wtr, err := NewFileWriter("data/tmp.npy")
	if err != nil {
		panic(err)
	}
	err = wtr.WriteInt16(data)
	if err != nil {
		panic(err)
	}

	r, err := os.Open("data/tmp.npy")
	if err != nil {
		panic(err)
	}
	rdr, err := NewReader(r)
	if err != nil {
		panic(err)
	}
	data, err = rdr.GetInt16()
	if err != nil {
		panic(err)
	}
	if !check_int16(data) {
		t.Fail()
	}
}

func Test_int8(t *testing.T) {

	data := []int8{0, 1, 2, 3, 4, 5, 6, 7}

	wtr, err := NewFileWriter("data/tmp.npy")
	if err != nil {
		panic(err)
	}
	err = wtr.WriteInt8(data)
	if err != nil {
		panic(err)
	}

	r, err := os.Open("data/tmp.npy")
	if err != nil {
		panic(err)
	}
	rdr, err := NewReader(r)
	if err != nil {
		panic(err)
	}
	data, err = rdr.GetInt8()
	if err != nil {
		panic(err)
	}
	if !check_int8(data) {
		t.Fail()
	}
}
