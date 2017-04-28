package gonpy

import (
	"io/ioutil"
	"os"
	"path"
	"strings"
	"testing"
)

func check_int8(data []int8) bool {
	for k := 0; k < len(data); k++ {
		if data[k] != int8(k) {
			return false
		}
	}
	return true
}

func check_int64(data []int64) bool {
	for k := 0; k < len(data); k++ {
		if data[k] != int64(k) {
			return false
		}
	}
	return true
}

func check_int32(data []int32) bool {
	for k := 0; k < len(data); k++ {
		if data[k] != int32(k) {
			return false
		}
	}
	return true
}

func check_int16(data []int16) bool {
	for k := 0; k < len(data); k++ {
		if data[k] != int16(k) {
			return false
		}
	}
	return true
}

func check_float64(data []float64) bool {
	for k := 0; k < len(data); k++ {
		if data[k] != float64(k) {
			return false
		}
	}
	return true
}

func check_float32(data []float32) bool {
	for k := 0; k < len(data); k++ {
		if data[k] != float32(k) {
			return false
		}
	}
	return true
}

func get_flist(dtype string) []string {

	infolist, err := ioutil.ReadDir("data")
	if err != nil {
		panic(err)
	}
	files := make([]string, 0)
	for _, v := range infolist {
		f := v.Name()
		if strings.Contains(f, dtype) {
			files = append(files, f)
		}
	}

	return files
}

func Test_i1(t *testing.T) {

	files := get_flist("i1")

	for _, fname := range files {

		fid, err := os.Open(path.Join("data", fname))
		if err != nil {
			panic(err)
		}

		rdr, err := NewReader(fid)
		if err != nil {
			panic(err)
		}
		data, err := rdr.GetInt8()
		if err != nil {
			panic(err)
		}

		if !check_int8(data) {
			t.Fail()
		}
	}
}

func Test_f8(t *testing.T) {

	files := get_flist("f8")

	for _, fname := range files {

		fid, err := os.Open(path.Join("data", fname))
		if err != nil {
			panic(err)
		}

		rdr, err := NewReader(fid)
		if err != nil {
			panic(err)
		}
		data, err := rdr.GetFloat64()
		if err != nil {
			panic(err)
		}

		if !check_float64(data) {
			t.Fail()
		}
	}
}

func Test_f8_file(t *testing.T) {

	files := get_flist("f8")

	for _, fname := range files {

		rdr, err := NewFileReader(path.Join("data", fname))
		if err != nil {
			panic(err)
		}
		data, err := rdr.GetFloat64()
		if err != nil {
			panic(err)
		}

		if !check_float64(data) {
			t.Fail()
		}
	}
}

func Test_f4(t *testing.T) {

	files := get_flist("f4")

	for _, fname := range files {

		fid, err := os.Open(path.Join("data", fname))
		if err != nil {
			panic(err)
		}

		rdr, err := NewReader(fid)
		if err != nil {
			panic(err)
		}
		data, err := rdr.GetFloat32()
		if err != nil {
			panic(err)
		}

		if !check_float32(data) {
			t.Fail()
		}
	}
}

func Test_i8(t *testing.T) {

	files := get_flist("i8")

	for _, fname := range files {

		fid, err := os.Open(path.Join("data", fname))
		if err != nil {
			panic(err)
		}

		rdr, err := NewReader(fid)
		if err != nil {
			panic(err)
		}
		data, err := rdr.GetInt64()
		if err != nil {
			panic(err)
		}

		if !check_int64(data) {
			t.Fail()
		}
	}

}

func Test_i4(t *testing.T) {

	files := get_flist("i4")

	for _, fname := range files {

		fid, err := os.Open(path.Join("data", fname))
		if err != nil {
			panic(err)
		}

		rdr, err := NewReader(fid)
		if err != nil {
			panic(err)
		}
		data, err := rdr.GetInt32()
		if err != nil {
			panic(err)
		}

		if !check_int32(data) {
			t.Fail()
		}
	}

}

func Test_i2(t *testing.T) {

	files := get_flist("i2")

	for _, fname := range files {

		fid, err := os.Open(path.Join("data", fname))
		if err != nil {
			panic(err)
		}

		rdr, err := NewReader(fid)
		if err != nil {
			panic(err)
		}
		data, err := rdr.GetInt16()
		if err != nil {
			panic(err)
		}

		if !check_int16(data) {
			t.Fail()
		}
	}
}
