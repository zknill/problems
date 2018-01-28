package strings

// Check if needle is any permutation of haystack
func IsPermutation(needle, haystack string) bool {
	chars := make(map[rune]int)

	for _, c := range haystack {
		chars[c]++
	}

	var permutation bool
	for _, n := range needle {
		if (chars[n] - 1) >= 0 {
			permutation = true
		} else {
			return false
		}

		chars[n]--
	}

	return permutation
}
