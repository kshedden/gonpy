package gonpy

/*
  gonpy reads and writes between numpy binary array data and Go slices

  The file specification is here:

  http://docs.scipy.org/doc/numpy/neps/npy-format.html

  When reading a multidimensional array, the data are returned as a
  one-dimensional slice that has been flattened as specified by the
  ColumnMajor field in the NpyReader struct.  If ColumnMajor is true,
  the array is in column major order.  If ColumnMajor is false, the
  array is in row major order.

  The writer defaults to writing the data as a vector.  The Shape
  field of the NpyWriter struct can be used to set other shapes.

  Unsigned numeric data types, fixed-width string types, compound
  dtypes, record arrays, and time types are not supported (file an
  issue if any of these are needed).  Python object types are not
  supported.
*/
