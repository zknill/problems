package strings

import "testing"

func TestUnique(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{name: "unique", input: "abcdefg", want: true},
		{name: "duplicates", input: "abcdefgg", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Unique(tt.input); got != tt.want {
				t.Errorf("Unique() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUniqueNoMap(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{name: "unique", input: "abcdefg", want: true},
		{name: "duplicates", input: "abcdefgg", want: false},
		{name: "duplicates_seperated", input: "hgjklyhsna", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Unique(tt.input); got != tt.want {
				t.Errorf("Unique() = %v, want %v", got, tt.want)
			}
		})
	}
}
