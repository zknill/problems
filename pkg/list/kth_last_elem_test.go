package list

import (
	"reflect"
	"testing"
)

func TestStrLinked_KthLastElem(t *testing.T) {
	tests := []struct {
		name   string
		l *StrLinked
		k int
		want   Node
	}{
		{
			name: "0th last",
			l: NewStrLinked("a", "b", "c", "d", "e", "f"),
			want: NewStrLinked("f").head,
		},
		{
			name: "1st last",
			l: NewStrLinked("a", "b", "c", "d", "e", "f"),
			k:1,
			want: NewStrLinked("e", "f").head,
		},
		{
			name: "2nd last",
			l: NewStrLinked("a", "b", "c", "d", "e", "f"),
			k:2,
			want: NewStrLinked("d", "e", "f").head,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l.KthLastElem(tt.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StrLinked.KthLastElem() = %v, want %v", got, tt.want)
			}
		})
	}
}
