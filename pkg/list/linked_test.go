package list

import (
	"reflect"
	"testing"
)

func TestNewStrLinked(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want *StrLinked
	}{
		{
			name: "success",
			args: []string{"a", "b", "c"},
			want: &StrLinked{
				head: &strNode{v: "a", n: &strNode{v: "b", n: &strNode{v: "c"}}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewStrLinked(tt.args...); !reflect.DeepEqual(tt.want, got) {
				t.Errorf("NewStrLinked() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStrLinked_RemoveDupes(t *testing.T) {
	tests := []struct {
		name string
		head node
		want *StrLinked
	}{
		{
			name: "has dupes",
			head: &strNode{v: "1", n: &strNode{v: "2", n: &strNode{v: "2", n: &strNode{v: "3"}}}},
			want: &StrLinked{head: &strNode{v: "1", n: &strNode{v: "2", n: &strNode{v: "3"}}}},
		},
		{
			name: "multiple dupes",
			head: &strNode{
				v: "1",
				n: &strNode{v: "2", n: &strNode{
					v: "2", n: &strNode{
						v: "3", n: &strNode{v: "3", n: &strNode{v: "4"}},
					},
				}},
			},
			want: &StrLinked{head: &strNode{v: "1", n: &strNode{v: "2", n: &strNode{v: "3", n: &strNode{v: "4"}}}}},
		},
		{
			name: "no dupes",
			head: &strNode{v: "1", n: &strNode{v: "2", n: &strNode{v: "3"}}},
			want: &StrLinked{head: &strNode{v: "1", n: &strNode{v: "2", n: &strNode{v: "3"}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &StrLinked{
				head: tt.head,
			}
			l.RemoveDupes()
			if !reflect.DeepEqual(tt.want, l) {
				t.Errorf("StrLinked.RemoveDupes() want %+v\ngot %+v", tt.want, l)
			}
		})
	}
}

func TestStrLinked_RemoveDupesNoBuf(t *testing.T) {
	tests := []struct {
		name string
		head node
		want *StrLinked
	}{
		{
			name: "has dupes",
			head: &strNode{v: "1", n: &strNode{v: "2", n: &strNode{v: "2", n: &strNode{v: "3"}}}},
			want: &StrLinked{head: &strNode{v: "1", n: &strNode{v: "2", n: &strNode{v: "3"}}}},
		},
		{
			name: "multiple dupes",
			head: &strNode{
				v: "1",
				n: &strNode{v: "2", n: &strNode{
					v: "2", n: &strNode{
						v: "3", n: &strNode{v: "3", n: &strNode{v: "4"}},
					},
				}},
			},
			want: &StrLinked{head: &strNode{v: "1", n: &strNode{v: "2", n: &strNode{v: "3", n: &strNode{v: "4"}}}}},
		},
		{
			name: "no dupes",
			head: &strNode{v: "1", n: &strNode{v: "2", n: &strNode{v: "3"}}},
			want: &StrLinked{head: &strNode{v: "1", n: &strNode{v: "2", n: &strNode{v: "3"}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &StrLinked{
				head: tt.head,
			}
			l.RemoveDupesNoBuf()
			if !reflect.DeepEqual(tt.want, l) {
				t.Errorf("StrLinked.RemoveDupes() want %+v\ngot %+v", tt.want, l)
			}
		})
	}
}

func Test_scanningNoder_hasNode(t *testing.T) {
	tests := []struct {
		name string
		head  node
		value node
		want bool
	}{
		{name: "first", head: &strNode{v: "1", n:&strNode{v:"1"}}, value: &strNode{v:"1"}, want: true},
		{name: "second", head: &strNode{v: "1", n: &strNode{v:"2", n: &strNode{v:"2"}}}, value: &strNode{v:"2"}, want: true},
		{name: "missing", head: &strNode{v: "1", n: &strNode{v:"2"}}, value: &strNode{v:"3"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := (&scanningNoder{head: tt.head}).hasNode(tt.value); got != tt.want {
				t.Errorf("containsStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
