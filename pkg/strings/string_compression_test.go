package strings

import "testing"

func TestCompressString(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{name: "compressed", input: "aaabbbc", want: "a3b3c1"},
		{name: "original shorter", input: "abc", want: "abc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompressString(tt.input); got != tt.want {
				t.Errorf("CompressString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkCompressString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CompressString("aaabbbccc")
	}
}
