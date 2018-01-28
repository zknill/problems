package strings

import "testing"

func TestOneAway(t *testing.T) {
	type args struct {
		originalStr string
		editStr     string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "remove one char (a)",
			args: args{originalStr: "pale", editStr: "ple"},
			want: true,
		},
		{
			name: "remove one char (s)",
			args: args{originalStr: "pales", editStr: "pale"},
			want: true,
		},
		{
			name: "change one char (p -> b)",
			args: args{originalStr: "pale", editStr: "bale"},
			want: true,
		},
		{
			name: "change two chars",
			args: args{originalStr: "pale", editStr: "bake"},
			want: false,
		},
		{
			name: "change two chars",
			args: args{originalStr: "paled", editStr: "baler"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := OneAway(tt.args.originalStr, tt.args.editStr); got != tt.want {
				t.Errorf("OneAway() = %v, want %v", got, tt.want)
			}
		})
	}
}
