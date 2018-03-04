package array

import "bytes"

// ZeroMatrix is a matrix that can zero values within itself.
// The matrix can be NxM
type ZeroMatrix struct {
	v []int
	w int
}

// NewZeroMatrix is a constructor for a ZeroMatrix, accepts NxM matrices.
func NewZeroMatrix(w int, v []int) *ZeroMatrix {
	return &ZeroMatrix{w: w, v: v}
}

// Zero considers it's matrix and will search for a zero in the matrix.
// If a zero is found, all values on that row and column in the matrix
// are set to zero.
func (z *ZeroMatrix) Zero() {
	var cols, rows []int
	for x := 0; x < z.width(); x++ {
		for y := 0; y < z.height(); y++ {
			if z.cell(x, y) == 0 {
				cols = append(cols, x)
				rows = append(rows, y)
			}
		}
	}
	for _, x := range cols {
		z.zeroCol(x)
	}
	for _, y := range rows {
		z.zeroRow(y)
	}
}

func (z *ZeroMatrix) zeroRow(y int) {
	for i := 0; i < z.width(); i++ {
		z.v[i+y*z.width()] = 0
	}
}

func (z *ZeroMatrix) zeroCol(x int) {
	for i := 0; i < z.height(); i++ {
		z.v[i*z.width()+x] = 0
	}
}

func (z *ZeroMatrix) valueWidth() int {
	return valueWidth(z.v)
}

func (z *ZeroMatrix) width() int {
	return z.w
}

func (z *ZeroMatrix) height() int {
	return len(z.v) / z.w
}

func (z *ZeroMatrix) cell(x, y int) int {
	return z.v[x+y*z.width()]
}

// String prints the matrix in a human readable format
func (z *ZeroMatrix) String() string {
	w := &bytes.Buffer{}
	writeMatrix(w, z)
	return w.String()
}
