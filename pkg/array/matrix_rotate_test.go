package array_test

import (
	"reflect"
	"testing"

	"github.com/zknill/problems/pkg/array"
)

func TestMatrix4_Rotate(t *testing.T) {
	tests := []struct {
		name    string
		degrees int
		want    *array.IntMatrixN
	}{
		{
			name:    "rotate zero",
			degrees: 0,
			want:    array.NewIntMatrixN([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}),
		},
		{
			name:    "rotate 90",
			degrees: 90,
			want:    array.NewIntMatrixN([]int{12, 8, 4, 0, 13, 9, 5, 1, 14, 10, 6, 2, 15, 11, 7, 3}),
		},
		{
			name:    "rotate 180",
			degrees: 180,
			want:    array.NewIntMatrixN([]int{15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}),
		},
		{
			name:    "rotate -90",
			degrees: -90,
			want:    array.NewIntMatrixN([]int{3, 7, 11, 15, 2, 6, 10, 14, 1, 5, 9, 13, 0, 4, 8, 12}),
		},
		{
			name:    "rotate -180",
			degrees: -180,
			want:    array.NewIntMatrixN([]int{15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
			m := array.NewIntMatrixN(v)
			m.Rotate(tt.degrees)

			if !reflect.DeepEqual(m, tt.want) {
				t.Errorf("Expected matching matricies:\ngot:\n%s\nwant:\n%s", m, tt.want)
			}
		})
	}
}

func TestMatrix3_Rotate(t *testing.T) {
	tests := []struct {
		name    string
		degrees int
		want    *array.IntMatrixN
	}{
		{
			name:    "rotate zero",
			degrees: 0,
			want:    array.NewIntMatrixN([]int{0, 1, 2, 3, 4, 5, 6, 7, 8}),
		},
		{
			name:    "rotate 90",
			degrees: 90,
			want:    array.NewIntMatrixN([]int{6, 3, 0, 7, 4, 1, 8, 5, 2}),
		},
		{
			name:    "rotate 180",
			degrees: 180,
			want:    array.NewIntMatrixN([]int{8, 7, 6, 5, 4, 3, 2, 1, 0}),
		},
		{
			name:    "rotate -90",
			degrees: -90,
			want:    array.NewIntMatrixN([]int{2, 5, 8, 1, 4, 7, 0, 3, 6}),
		},
		{
			name:    "rotate -180",
			degrees: -180,
			want:    array.NewIntMatrixN([]int{8, 7, 6, 5, 4, 3, 2, 1, 0}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
			m := array.NewIntMatrixN(v)
			m.Rotate(tt.degrees)

			if !reflect.DeepEqual(m, tt.want) {
				t.Errorf("%q. Expected matching matricies:\ngot:\n%s\nwant:\n%s", tt.name, m, tt.want)
			}
		})
	}
}
