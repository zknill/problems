package list

import (
	"reflect"
	"testing"
)

func TestStrLinked_Partition(t *testing.T) {
	tests := []struct {
		name string
		l    *StrLinked
		v    string
		want *StrLinked
	}{
		{
			name: "success end to front",
			l:    NewStrLinked("a", "c", "d", "b", "g", "a"),
			v:    "b",
			want: NewStrLinked("a", "a", "c", "d", "b","g"),
		},
		{
			name: "success first to last",
			l:    NewStrLinked("c", "a", "a", "a", "b", "a"),
			v:    "c",
			want: NewStrLinked("a", "b", "a", "a", "a","c"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.l.Partition(tt.v)
			if !reflect.DeepEqual(tt.want, tt.l) {
				t.Errorf("StrLinked.Partition()\nwant\t%+v\ngot\t\t%+v", tt.want, tt.l)
			}
		})
	}
}
