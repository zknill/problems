package strings

// Check if any of the permutations of a string are palindromes.
func PalindromePermutation(str string) bool {
	counts := make(map[rune]int)

	for _, c := range str {
		if c == ' ' {
			continue
		}
		counts[c]++
	}

	var odd bool
	for _, num := range counts {
		if num%2 == 1 {
			if odd {
				// there is already one count with an
				// odd number of runes.
				return false
			}
			odd = true
		}
	}

	return true
}

func PalindromePermutationImproved(str string) bool {
	var bitMap int
	for _, c := range str {
		if c == ' ' {
			continue
		}
		bitMap = toggle(bitMap, uint(c))
	}

	return bitMap == 0 || ((bitMap & (bitMap - 1)) == 0)
}

func toggle(bitMap int, idx uint) int {
	if idx == 0 {
		return bitMap
	}

	mask := 1 << idx
	if bitMap&mask == 0 {
		bitMap |= mask
	} else {
		bitMap &= mask
	}

	return bitMap
}
