package strings

import "strconv"

// CompressString returns the compressed version
// of the input string.
//		aaabbc -> a3b2c1
// If the compressed string is not shorter than
// the original then the original is returned.
func CompressString(original string) string {
	var compressed string
	var consecutive int
	for idx, char := range original {
		consecutive++
		if idx+1 >= len(original) || original[idx] != original[idx+1] {
			compressed += string(char)
			compressed += strconv.Itoa(consecutive)
			consecutive = 0
		}
	}
	if len(compressed) >= len(original) {
		return original
	}

	return compressed
}
