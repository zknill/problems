package heap

import (
	"reflect"
	"testing"
)

func TestNewMinHeap(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want *Binary
	}{
		{
			name: "success",
			vals: []int{90, 55, 50, 4, 7, 87, 2, 1},
			want: &Binary{heap: []int{1, 2, 4, 7, 50, 87, 55, 90}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMinHeap(tt.vals...); !reflect.DeepEqual(got.heap, tt.want.heap) {
				t.Errorf("NewMinHeap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewMaxHeap(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want []int
	}{
		{
			name: "success",
			vals: []int{1, 2, 3, 4, 5, 6, 7},
			want: []int{7, 4, 6, 1, 3, 2, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMaxHeap(tt.vals...); !reflect.DeepEqual(got.heap, tt.want) {
				t.Errorf("NewMaxHeap() = %v, want %v", got.heap, tt.want)
			}
		})
	}
}
