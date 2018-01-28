package strings

import "testing"

func TestPalindromePermutation(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{name: "success", input: "tact coa", want: true},
		{name: "failure", input: "tact coat", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PalindromePermutation(tt.input); got != tt.want {
				t.Errorf("PalindromePermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPalindromePermutationImproved(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{name: "success", input: "tact coa", want: true},
		{name: "failure", input: "tact coat", want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PalindromePermutationImproved(tt.input); got != tt.want {
				t.Errorf("PalindromePermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkPalindromePermutation(b *testing.B) {
	benchmarks := []struct {
		name string
		f    func(string) bool
	}{
		{name: "PalindromePermutation", f: PalindromePermutation},
		{name: "PalindromePermutationImproved", f: PalindromePermutation},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				bm.f("tact coa")
			}
		})
	}
}
