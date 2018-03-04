package array

import (
	"fmt"
	"io"
	"math"
	"strconv"
)

type matrixer interface {
	valueWidth() int
	width() int
	height() int
	cell(x, y int) int
}

func writeMatrix(w io.Writer, m matrixer) {
	f := " %0" + strconv.Itoa(m.valueWidth()) + "d "

	for y := 0; y < m.height(); y++ {
		io.WriteString(w, "\n")
		for x := 0; x < m.width(); x++ {
			io.WriteString(w, fmt.Sprintf(f, m.cell(x, y)))
		}
		io.WriteString(w, "\n")
	}
}

func valueWidth(v []int) int {
	var max float64
	for _, vv := range v {
		max = math.Max(max, math.Log10(float64(vv)))
	}
	return int(math.Ceil(max))
}
