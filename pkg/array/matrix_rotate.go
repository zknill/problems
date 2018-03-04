package array

import (
	"bytes"
	"math"
)

// IntMatrixN represents a NxN matrix storing ints
type IntMatrixN struct {
	v []int
	h int
}

// NewIntMatrixN constructs a new matrix, assuming that it's
// an NxN matrix and calculating the height accordingly.
func NewIntMatrixN(v []int) *IntMatrixN {
	return &IntMatrixN{
		v: v,
		h: int(math.Sqrt(float64(len(v)))),
	}
}

// Rotates the matrix by factor of 90 degrees in +/- directions
func (m *IntMatrixN) Rotate(degrees int) {
	r := degrees / 90

	if r < 0 {
		r += 4
	}

	for ; r > 0; r-- {
		m.rotate90()
	}
}

func (m *IntMatrixN) rotate90() {
	v := make([]int, len(m.v))
	for row := 1; row <= m.height(); row++ {
		for col := 1; col <= m.width(); col++ {
			c := m.cell(col-1, row-1)
			w := m.width()
			i := (w * col) - row
			v[i] = c
		}
	}
	m.v = v
}

func (m *IntMatrixN) width() int {
	return len(m.v) / m.h
}

func (m *IntMatrixN) height() int {
	return m.width()
}

func (m *IntMatrixN) cell(x, y int) int {
	return m.v[x+y*m.width()]
}

// String prints the matrix
func (m *IntMatrixN) String() string {
	w := &bytes.Buffer{}
	writeMatrix(w, m)
	return w.String()
}

func (m *IntMatrixN) valueWidth() int {
	return valueWidth(m.v)
}
