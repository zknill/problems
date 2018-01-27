package strings

import "testing"

func TestURLify(t *testing.T) {
	tests := []struct {
		name string
		input string
		want string
	}{
		{
			name: "simple single replace",
			input: "hello world",
			want: "hello%20world",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := URLify(tt.input); got != tt.want {
				t.Errorf("URLify() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestURLifyImproved(t *testing.T) {
	tests := []struct {
		name string
		input string
		want string
	}{
		{
			name: "simple single replace",
			input: "hello world",
			want: "hello%20world",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := URLifyImproved(tt.input); got != tt.want {
				t.Errorf("URLifyImproved() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func BenchmarkURLify(b *testing.B) {
	for i := 0; i < b.N; i++ {
		URLify("hello world")
	}
}

func Benchmark(b *testing.B) {
	benchmarks := []struct {
		name string
		f func(string)string
	}{
		{name:"URLify", f: URLify},
		{name:"URLifyImproved", f: URLifyImproved},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				bm.f("hello world")
			}
		})
	}
}
