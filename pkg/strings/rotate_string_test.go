package strings

import "testing"

func TestIsRotation(t *testing.T) {
	tests := []struct {
		name string
		original string
		rotated  string
		want bool
	}{
		{
			name: "same string",
			original: "waterbottle",
			rotated: "waterbottle",
			want: true,
		},
		{
			name: "one place rotation",
			original: "waterbottle",
			rotated: "ewaterbottl",
			want: true,
		},
		{
			name: "two places rotation",
			original: "waterbottle",
			rotated: "lewaterbott",
			want: true,
		},
		{
			name: "three places rotation",
			original: "waterbottle",
			rotated: "tlewaterbot",
			want: true,
		},
		{
			name: "four places rotation",
			original: "waterbottle",
			rotated: "ttlewaterbo",
			want: true,
		},
		{
			name: "five  places rotation",
			original: "waterbottle",
			rotated: "ottlewaterb",
			want: true,
		},
		{
			name: "full rotation",
			original: "waterbottle",
			rotated: "bottlewater",
			want: true,
		},
		{
			name: "not same words",
			original: "waterbottle",
			rotated: "houseboat",
			want: false,
		},
		{
			name: "not fully same words",
			original: "waterbottle",
			rotated: "botterwater",
			want: false,
		},
		{
			name: "different letter",
			original: "waterbottle",
			rotated: "bottlewatef",
			want: false,
		},
		{
			name: "different letter",
			original: "waterbottle",
			rotated: "fbottlewate",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsRotation(tt.original, tt.rotated); got != tt.want {
				t.Errorf("IsRotation() = %v, want %v", got, tt.want)
			}
		})
	}
}
