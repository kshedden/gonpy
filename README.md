gonpy : Read Numpy binary files (.npy files) in Golang
======================================================

__gonpy__ reads and writes Numpy binary array data to/from Go slices

The npy file specification is here:

http://docs.scipy.org/doc/numpy/neps/npy-format.html

The documentation for this package can be found here:

https://godoc.org/github.com/kshedden/gonpy

When reading a multidimensional array, the data are returned as a
one-dimensional slice that has been flattened as specified by the
ColumnMajor field in the NpyReader struct.  If ColumnMajor is true,
the array is flattened in column major order.  If ColumnMajor is
false, the array is flattened in row major order.

The writer defaults to writing the data as a vector.  The Shape field
of the NpyWriter struct can be used to set other shapes.

Unsigned numeric data types, fixed-width string types, compound
dtypes, record arrays, and time dtypes are not supported (file an
issue if any of these are needed).  Python object types are of course
not supported.

The following example shows how to read a npy file into a slice of
float64 values (for clarity error handling is omitted).

```
r, _ := gonpy.NewFileReader("data.npy")
data, _ := r.GetFloat64()
```

The reader value has fields such sas `r.Shape` that provide additional
information about the array.

The following example shows how to write a slice of int32 values to a
npy file (again, error handling is omitted).

```
w, _ := gonpy.NewFileWriter("data.npy")
_ = w.WriteFloat64(data)
```

To specify a shape or other attributes, modify the writer object
before writing the array, for example:

```
w, _ := gonpy.NewFileWriter("data.npy")
w.Shape = []int32{50, 2}
w.Version = 2
_ = w.WriteFloat64(data)
```

To write to a stream, say a gzip stream, use the following:

```
f, _ := os.Create("data.npy.gz")
g := gzip.NewWriter(f)
w, _ := gonpy.WriterFromStream(g)
_ = w.WriteFloat64(data)
f.Close()
```
