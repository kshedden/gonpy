import numpy as np

"""
Generate data files for testing.
"""

for dtype in ("i1", "i2", "i4", "i8", "u1", "u2", "u4", "u8", "f4", "f8", "c8", "c16"):
    ii = 1
    for dims in 1,2:
        for order in "CF":

            data = np.arange(8, dtype=dtype)

            if dims == 2:
                if order == "C":
                    data = data.reshape((4, 2))
                elif order == "F":
                    data = data.reshape((2, 4)).T


            fname = "data/%s_%d.npy" % (dtype, ii)
            np.save(fname, data)
            ii += 1
