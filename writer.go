package gonpy

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"strings"
)

// NpyWriter can write data from a Go slice to a Numpy binary array.
type NpyWriter struct {

	// Defaults to little endian, but can be set to
	// binary.BigEndian before writing data
	Endian binary.ByteOrder

	// Defaults to nx1, where n is the length of data.  Can be set
	// to any shape with any number of dimensions.  The shape is
	// not checked for compatibility with the data.
	Shape []int

	// Defaults to false (row major order), can be set to true
	// (column major order) before writing the data.
	ColumnMajor bool

	w io.WriteCloser
}

// NewWriter returns a NpyWriter that can be used to write data to a
// Numpy binary format file.  After creation, call one of the WriteXX
// methods to write array data to the file.  The file is automatically
// closed at the end of that call.  Only one array can be written.
//
// Note: this may be changed to take a io.Writer in the future.
// Consider using NewFileWriter or WriterFromStream instead.
func NewWriter(fname string) (*NpyWriter, error) {

	fid, err := os.Create(fname)
	if err != nil {
		return nil, err
	}

	return WriterFromStream(fid)
}

// NewFileWriter returns a NpyWriter that can be used to write data to
// a Numpy binary format file.  After creation, call one of the
// WriteXX methods to write array data to the file.  The file is
// automatically closed at the end of that call.  Only one array can
// be written.
func NewFileWriter(fname string) (*NpyWriter, error) {

	fid, err := os.Create(fname)
	if err != nil {
		return nil, err
	}

	return WriterFromStream(fid)
}

// NewWriter returns a NpyWriter that can be used to write data to an
// io.WriteCloser is the Numpy binary format.  After creation, call
// one of the WriteXX methods to write array data to the writer.  The
// file is automatically closed at the end of that call.  Only one
// array can be written.
func WriterFromStream(w io.WriteCloser) (*NpyWriter, error) {

	wtr := &NpyWriter{w: w, Endian: binary.LittleEndian}
	return wtr, nil
}

// WriteFloat64 writes a float64 slice to the npy format file.
func (wtr *NpyWriter) WriteFloat64(data []float64) error {

	err := wtr.write_header("f8", len(data))
	if err != nil {
		return err
	}

	for _, v := range data {
		err := binary.Write(wtr.w, wtr.Endian, v)
		if err != nil {
			return err
		}
	}

	wtr.w.Close()

	return nil
}

// WriteFloat32 writes a float32 slice to the npy format file.
func (wtr *NpyWriter) WriteFloat32(data []float32) error {

	err := wtr.write_header("f4", len(data))
	if err != nil {
		return err
	}

	for _, v := range data {
		err := binary.Write(wtr.w, wtr.Endian, v)
		if err != nil {
			return err
		}
	}

	wtr.w.Close()

	return nil
}

// WriteInt64 writes a int64 slice to the npy format file.
func (wtr *NpyWriter) WriteInt64(data []int64) error {

	err := wtr.write_header("i8", len(data))
	if err != nil {
		return err
	}

	for _, v := range data {
		err := binary.Write(wtr.w, wtr.Endian, v)
		if err != nil {
			return err
		}
	}

	wtr.w.Close()

	return nil
}

// WriteInt32 writes a int32 slice to the npy format file.
func (wtr *NpyWriter) WriteInt32(data []int32) error {

	err := wtr.write_header("i4", len(data))
	if err != nil {
		return err
	}

	for _, v := range data {
		err := binary.Write(wtr.w, wtr.Endian, v)
		if err != nil {
			return err
		}
	}

	wtr.w.Close()

	return nil
}

// WriteInt16 writes a int16 slice to the npy format file.
func (wtr *NpyWriter) WriteInt16(data []int16) error {

	err := wtr.write_header("i2", len(data))
	if err != nil {
		return err
	}

	for _, v := range data {
		err := binary.Write(wtr.w, wtr.Endian, v)
		if err != nil {
			return err
		}
	}

	wtr.w.Close()

	return nil
}

// WriteInt8 writes a int8 slice to the npy format file.
func (wtr *NpyWriter) WriteInt8(data []int8) error {

	err := wtr.write_header("i1", len(data))
	if err != nil {
		return err
	}

	for _, v := range data {
		err := binary.Write(wtr.w, wtr.Endian, v)
		if err != nil {
			return err
		}
	}

	wtr.w.Close()

	return nil
}

func (wtr *NpyWriter) write_header(dtype string, length int) error {

	n, err := wtr.w.Write([]byte("\x93NUMPY"))
	if err != nil {
		return err
	} else if n != 6 {
		return fmt.Errorf("unable to write magic number")
	}
	err = binary.Write(wtr.w, binary.LittleEndian, uint8(1))
	if err != nil {
		return err
	}
	err = binary.Write(wtr.w, binary.LittleEndian, uint8(0))
	if err != nil {
		return err
	}

	if wtr.Endian == binary.LittleEndian {
		dtype = "<" + dtype
	} else {
		dtype = ">" + dtype
	}

	var shape_string string
	if wtr.Shape != nil {
		shape_string = ""
		for _, v := range wtr.Shape {
			shape_string = shape_string + fmt.Sprintf("%d,", v)
		}
		shape_string = "(" + shape_string + ")"
	} else {
		shape_string = fmt.Sprintf("(%d,)", length)
	}

	cmaj := "False"
	if wtr.ColumnMajor {
		cmaj = "True"
	}

	header := fmt.Sprintf("{'descr': '%s', 'fortran_order': %s, 'shape': %s,}",
		dtype, cmaj, shape_string)
	pad := 16 - ((10 + len(header)) % 16)
	if pad > 0 {
		header = header + strings.Repeat(" ", pad)
	}

	err = binary.Write(wtr.w, binary.LittleEndian, uint16(len(header)))
	if err != nil {
		return err
	}
	n, err = wtr.w.Write([]byte(header))
	if err != nil {
		return err
	} else if n != len(header) {
		return fmt.Errorf("unable to write header")
	}

	return nil
}
