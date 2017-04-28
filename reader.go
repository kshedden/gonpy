package gonpy

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// NpyReader can read data from a Numpy binary array into a Go slice.
type NpyReader struct {
	dtype  string
	endian binary.ByteOrder

	// The shape of the array as specified in the file.
	Shape []int

	r io.Reader

	// If true, the data are flattened in column-major order,
	// otherwise they are flattened in row-major order.
	ColumnMajor bool

	n_elt int
}

// NewFileReader returns a NpyReader that can be used to obtain array
// data from the given file.  Call one of the GetXX methods to obtain
// the slice.
func NewFileReader(f string) (*NpyReader, error) {

	fid, err := os.Open(f)
	if err != nil {
		return nil, err
	}
	r, err := NewReader(fid)
	return r, err
}

// NewReader returns a NpyReader that can be used to obtain array data
// as a Go slice.  The Go slice has a type matching the dtype in the
// Numpy file.  Call one of the GetXX methods to obtain the slice.
func NewReader(r io.Reader) (*NpyReader, error) {

	b := make([]byte, 6)
	n, err := r.Read(b)
	if err != nil {
		return nil, err
	} else if n != 6 {
		return nil, fmt.Errorf("Input appears to be truncated")
	} else if string(b) != "\x93NUMPY" {
		return nil, fmt.Errorf("Not npy format data (wrong magic number)")
	}

	var version uint8
	err = binary.Read(r, binary.LittleEndian, &version)
	if err != nil {
		return nil, err
	} else if version != 1 {
		return nil, fmt.Errorf("Can only read npy version 1")
	}

	var minor uint8
	err = binary.Read(r, binary.LittleEndian, &minor)
	if err != nil {
		return nil, err
	} else if minor != 0 {
		return nil, fmt.Errorf("Can only read npy version 1")
	}

	var header_length uint16
	err = binary.Read(r, binary.LittleEndian, &header_length)
	if err != nil {
		return nil, err
	}

	header_bytes := make([]byte, header_length)
	n, err = r.Read(header_bytes)
	if err != nil {
		return nil, err
	} else if uint16(n) != header_length {
		return nil, fmt.Errorf("Input appears to be truncated")
	}

	// Get the dtype
	re := regexp.MustCompile(`'descr':\s*'([^']*)'`)
	ma := re.FindSubmatch(header_bytes)
	if ma == nil {
		return nil, fmt.Errorf("dtype description not found in header")
	}
	dtype := string(ma[1])

	// Get the order information
	re = regexp.MustCompile(`'fortran_order':\s*(False|True)`)
	ma = re.FindSubmatch(header_bytes)
	if ma == nil {
		return nil, fmt.Errorf("fortran_order not found in header")
	}
	fortran_order := string(ma[1])

	// Get the shape information
	re = regexp.MustCompile(`'shape':\s*\(([^\(]*)\)`)
	ma = re.FindSubmatch(header_bytes)
	if ma == nil {
		return nil, fmt.Errorf("fortran_order not found in header")
	}
	shape_string := string(ma[1])
	svals := strings.Split(shape_string, ",")
	shape := make([]int, 0)
	n_elt := 1
	for _, s := range svals {
		s = strings.Trim(s, " ")
		if len(s) == 0 {
			break
		}
		x, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		n_elt *= x
		shape = append(shape, x)
	}

	var endian binary.ByteOrder
	if dtype == ">" {
		endian = binary.BigEndian
	} else {
		endian = binary.LittleEndian
	}

	rdr := &NpyReader{dtype: dtype[1:], ColumnMajor: fortran_order == "True",
		Shape: shape, endian: endian, n_elt: n_elt, r: r}

	return rdr, nil

}

// ReaderFromStream returns a NpyReader that can be used to obtain the
// array data as a Go slice of a specific numeric type.
//
// Deprecated: use NewReader instead (April, 2017)
func ReaderFromStream(stream io.Reader) (*NpyReader, error) {
	rdr, err := NewReader(stream)
	return rdr, err
}

// GetFloat64 returns the array data from the npy file as a slice of
// float64 values.
func (rdr *NpyReader) GetFloat64() ([]float64, error) {
	if rdr.dtype != "f8" {
		return nil, fmt.Errorf("Reader does not contain float64 data")
	}
	data := make([]float64, rdr.n_elt)
	for k := 0; k < rdr.n_elt; k++ {
		err := binary.Read(rdr.r, rdr.endian, &data[k])
		if err != nil {
			return nil, err
		}
	}

	return data, nil
}

// GetFloat32 returns the array data from the npy file as a slice of
// float32 values.
func (rdr *NpyReader) GetFloat32() ([]float32, error) {
	if rdr.dtype != "f4" {
		return nil, fmt.Errorf("Reader does not contain float32 data")
	}
	data := make([]float32, rdr.n_elt)
	for k := 0; k < rdr.n_elt; k++ {
		err := binary.Read(rdr.r, rdr.endian, &data[k])
		if err != nil {
			return nil, err
		}
	}

	return data, nil
}

// GetInt8 returns the array data from the npy file as a slice of int8
// values.
func (rdr *NpyReader) GetInt8() ([]int8, error) {
	if rdr.dtype != "i1" {
		return nil, fmt.Errorf("Reader does not contain int8 data")
	}
	data := make([]int8, rdr.n_elt)
	for k := 0; k < rdr.n_elt; k++ {
		err := binary.Read(rdr.r, rdr.endian, &data[k])
		if err != nil {
			return nil, err
		}
	}

	return data, nil
}

// GetInt16 returns the array data from the npy file as a slice of int16
// values.
func (rdr *NpyReader) GetInt16() ([]int16, error) {
	if rdr.dtype != "i2" {
		return nil, fmt.Errorf("Reader does not contain int16 data")
	}
	data := make([]int16, rdr.n_elt)
	for k := 0; k < rdr.n_elt; k++ {
		err := binary.Read(rdr.r, rdr.endian, &data[k])
		if err != nil {
			return nil, err
		}
	}

	return data, nil
}

// GetInt32 returns the array data from the npy file as a slice of int32
// values.
func (rdr *NpyReader) GetInt32() ([]int32, error) {
	if rdr.dtype != "i4" {
		return nil, fmt.Errorf("Reader does not contain int32 data")
	}
	data := make([]int32, rdr.n_elt)
	for k := 0; k < rdr.n_elt; k++ {
		err := binary.Read(rdr.r, rdr.endian, &data[k])
		if err != nil {
			return nil, err
		}
	}

	return data, nil
}

// GetInt64 returns the array data from the npy file as a slice of int64
// values.
func (rdr *NpyReader) GetInt64() ([]int64, error) {
	if rdr.dtype != "i8" {
		return nil, fmt.Errorf("Reader does not contain int64 data")
	}
	data := make([]int64, rdr.n_elt)
	for k := 0; k < rdr.n_elt; k++ {
		err := binary.Read(rdr.r, rdr.endian, &data[k])
		if err != nil {
			return nil, err
		}
	}

	return data, nil
}
