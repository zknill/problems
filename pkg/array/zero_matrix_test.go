package array_test

import (
	"reflect"
	"testing"

	"github.com/zknill/problems/pkg/array"
)

func TestZeroMatrix_Zero(t *testing.T) {
	tests := []struct {
		name string
		m    *array.ZeroMatrix
		want *array.ZeroMatrix
	}{
		{
			name: "zero zeros NxN",
			m:    array.NewZeroMatrix(3, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}),
			want: array.NewZeroMatrix(3, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}),
		},
		{
			name: "zero zeros NxM",
			m:    array.NewZeroMatrix(2, []int{1, 2, 3, 4, 5, 6, 7, 8}),
			want: array.NewZeroMatrix(2, []int{1, 2, 3, 4, 5, 6, 7, 8}),
		},
		{
			name: "one zero NxN",
			m:    array.NewZeroMatrix(3, []int{1, 0, 3, 4, 5, 6, 7, 8, 9}),
			want: array.NewZeroMatrix(3, []int{0, 0, 0, 4, 0, 6, 7, 0, 9}),
		},
		{
			name: "one zero NxM",
			m:    array.NewZeroMatrix(2, []int{1, 0, 3, 4, 5, 6, 7, 8}),
			want: array.NewZeroMatrix(2, []int{0, 0, 3, 0, 5, 0, 7, 0}),
		},
		{
			name: "two zeros NxN",
			m:    array.NewZeroMatrix(3, []int{1, 0, 3, 4, 5, 6, 7, 8, 0}),
			want: array.NewZeroMatrix(3, []int{0, 0, 0, 4, 0, 0, 0, 0, 0}),
		},
		{
			name: "two zeros NxM",
			m:    array.NewZeroMatrix(2, []int{1, 0, 3, 4, 5, 6, 7, 0}),
			want: array.NewZeroMatrix(2, []int{0, 0, 3, 0, 5, 0, 0, 0}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.Zero()
			if !reflect.DeepEqual(tt.m, tt.want) {
				t.Errorf("Expected matching matrices: got =\n%s\nwant=\n%s", tt.m, tt.want)
			}
		})
	}
}
