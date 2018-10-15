// GENERATED CODE, DO NOT EDIT

package gonpy

import (
	"encoding/binary"
	"fmt"
)

// GetFloat64 returns the array data as a slice of float64 values.
func (rdr *NpyReader) GetFloat64() ([]float64, error) {

	if rdr.Dtype != "f8" {
		return nil, fmt.Errorf("Reader does not contain float64 data")
	}

	data := make([]float64, rdr.n_elt)
	err := binary.Read(rdr.r, rdr.Endian, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// GetFloat32 returns the array data as a slice of float32 values.
func (rdr *NpyReader) GetFloat32() ([]float32, error) {

	if rdr.Dtype != "f4" {
		return nil, fmt.Errorf("Reader does not contain float32 data")
	}

	data := make([]float32, rdr.n_elt)
	err := binary.Read(rdr.r, rdr.Endian, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// GetUint64 returns the array data as a slice of uint64 values.
func (rdr *NpyReader) GetUint64() ([]uint64, error) {

	if rdr.Dtype != "u8" {
		return nil, fmt.Errorf("Reader does not contain uint64 data")
	}

	data := make([]uint64, rdr.n_elt)
	err := binary.Read(rdr.r, rdr.Endian, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// GetUint32 returns the array data as a slice of uint32 values.
func (rdr *NpyReader) GetUint32() ([]uint32, error) {

	if rdr.Dtype != "u4" {
		return nil, fmt.Errorf("Reader does not contain uint32 data")
	}

	data := make([]uint32, rdr.n_elt)
	err := binary.Read(rdr.r, rdr.Endian, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// GetUint16 returns the array data as a slice of uint16 values.
func (rdr *NpyReader) GetUint16() ([]uint16, error) {

	if rdr.Dtype != "u2" {
		return nil, fmt.Errorf("Reader does not contain uint16 data")
	}

	data := make([]uint16, rdr.n_elt)
	err := binary.Read(rdr.r, rdr.Endian, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// GetUint8 returns the array data as a slice of uint8 values.
func (rdr *NpyReader) GetUint8() ([]uint8, error) {

	if rdr.Dtype != "u1" {
		return nil, fmt.Errorf("Reader does not contain uint8 data")
	}

	data := make([]uint8, rdr.n_elt)
	err := binary.Read(rdr.r, rdr.Endian, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// GetInt64 returns the array data as a slice of int64 values.
func (rdr *NpyReader) GetInt64() ([]int64, error) {

	if rdr.Dtype != "i8" {
		return nil, fmt.Errorf("Reader does not contain int64 data")
	}

	data := make([]int64, rdr.n_elt)
	err := binary.Read(rdr.r, rdr.Endian, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// GetInt32 returns the array data as a slice of int32 values.
func (rdr *NpyReader) GetInt32() ([]int32, error) {

	if rdr.Dtype != "i4" {
		return nil, fmt.Errorf("Reader does not contain int32 data")
	}

	data := make([]int32, rdr.n_elt)
	err := binary.Read(rdr.r, rdr.Endian, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// GetInt16 returns the array data as a slice of int16 values.
func (rdr *NpyReader) GetInt16() ([]int16, error) {

	if rdr.Dtype != "i2" {
		return nil, fmt.Errorf("Reader does not contain int16 data")
	}

	data := make([]int16, rdr.n_elt)
	err := binary.Read(rdr.r, rdr.Endian, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// GetInt8 returns the array data as a slice of int8 values.
func (rdr *NpyReader) GetInt8() ([]int8, error) {

	if rdr.Dtype != "i1" {
		return nil, fmt.Errorf("Reader does not contain int8 data")
	}

	data := make([]int8, rdr.n_elt)
	err := binary.Read(rdr.r, rdr.Endian, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// WriteFloat64 writes a slice of float64 values in npy format.
func (wtr *NpyWriter) WriteFloat64(data []float64) error {

	err := wtr.write_header("f8", len(data))
	if err != nil {
		return err
	}

	err = binary.Write(wtr.w, wtr.Endian, data)
	if err != nil {
		return err
	}

	wtr.w.Close()

	return nil
}

// WriteFloat32 writes a slice of float32 values in npy format.
func (wtr *NpyWriter) WriteFloat32(data []float32) error {

	err := wtr.write_header("f4", len(data))
	if err != nil {
		return err
	}

	err = binary.Write(wtr.w, wtr.Endian, data)
	if err != nil {
		return err
	}

	wtr.w.Close()

	return nil
}

// WriteUint64 writes a slice of uint64 values in npy format.
func (wtr *NpyWriter) WriteUint64(data []uint64) error {

	err := wtr.write_header("u8", len(data))
	if err != nil {
		return err
	}

	err = binary.Write(wtr.w, wtr.Endian, data)
	if err != nil {
		return err
	}

	wtr.w.Close()

	return nil
}

// WriteUint32 writes a slice of uint32 values in npy format.
func (wtr *NpyWriter) WriteUint32(data []uint32) error {

	err := wtr.write_header("u4", len(data))
	if err != nil {
		return err
	}

	err = binary.Write(wtr.w, wtr.Endian, data)
	if err != nil {
		return err
	}

	wtr.w.Close()

	return nil
}

// WriteUint16 writes a slice of uint16 values in npy format.
func (wtr *NpyWriter) WriteUint16(data []uint16) error {

	err := wtr.write_header("u2", len(data))
	if err != nil {
		return err
	}

	err = binary.Write(wtr.w, wtr.Endian, data)
	if err != nil {
		return err
	}

	wtr.w.Close()

	return nil
}

// WriteUint8 writes a slice of uint8 values in npy format.
func (wtr *NpyWriter) WriteUint8(data []uint8) error {

	err := wtr.write_header("u1", len(data))
	if err != nil {
		return err
	}

	err = binary.Write(wtr.w, wtr.Endian, data)
	if err != nil {
		return err
	}

	wtr.w.Close()

	return nil
}

// WriteInt64 writes a slice of int64 values in npy format.
func (wtr *NpyWriter) WriteInt64(data []int64) error {

	err := wtr.write_header("i8", len(data))
	if err != nil {
		return err
	}

	err = binary.Write(wtr.w, wtr.Endian, data)
	if err != nil {
		return err
	}

	wtr.w.Close()

	return nil
}

// WriteInt32 writes a slice of int32 values in npy format.
func (wtr *NpyWriter) WriteInt32(data []int32) error {

	err := wtr.write_header("i4", len(data))
	if err != nil {
		return err
	}

	err = binary.Write(wtr.w, wtr.Endian, data)
	if err != nil {
		return err
	}

	wtr.w.Close()

	return nil
}

// WriteInt16 writes a slice of int16 values in npy format.
func (wtr *NpyWriter) WriteInt16(data []int16) error {

	err := wtr.write_header("i2", len(data))
	if err != nil {
		return err
	}

	err = binary.Write(wtr.w, wtr.Endian, data)
	if err != nil {
		return err
	}

	wtr.w.Close()

	return nil
}

// WriteInt8 writes a slice of int8 values in npy format.
func (wtr *NpyWriter) WriteInt8(data []int8) error {

	err := wtr.write_header("i1", len(data))
	if err != nil {
		return err
	}

	err = binary.Write(wtr.w, wtr.Endian, data)
	if err != nil {
		return err
	}

	wtr.w.Close()

	return nil
}
