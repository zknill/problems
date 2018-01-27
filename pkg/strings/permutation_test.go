package strings

import "testing"

func TestIsPermutation(t *testing.T) {
	type args struct {
		needle   string
		haystack string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "simple permutation",
			args: args{needle: "ab", haystack: "abc"},
			want: true,
		},
		{
			name: "permutation duplicated chars",
			args: args{needle: "aa", haystack: "abca"},
			want: true,
		},
		{
			name: "not permutation duplicated chars",
			args: args{needle: "aa", haystack: "abc"},
			want: false,
		},
		{
			name: "permutation many duplicated chars",
			args: args{needle: "aabhdcg", haystack: "aabbccddeeffgghh"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPermutation(tt.args.needle, tt.args.haystack); got != tt.want {
				t.Errorf("IsPermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
