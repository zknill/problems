package strings

import "strings"

// given a string "waterbottle", check if another string "erbottlewat"
// is a rotate of the first string, using only a single call to `isSubstring`

// IsRotation takes two strings, an original and a rotated string.
// It checks if the rotated string really is a rotated version of the original
// string. It is only allowed a single call to strings.Contains(...).
//
//	e.g. "waterbottle" & "erbottlewate" returns true
//		while "waterbottle" & "bottlewatef" returns false.
func IsRotation(original, rotated string) bool {
	if len(original) != len(rotated) {
		return false
	}

	idx, ri, ok := findFirstOverlap(original, rotated)
	if !ok {
		return false
	}

	diff := ri-idx
	for x := idx; x < len(rotated) - diff; x++ {
		if original[x] != rotated[x+diff] {
			return false
		}
	}

	return strings.Contains(original, rotated[:ri])
}

func findFirstOverlap(original, rotated string) (i, ri int, found bool) {
	for i, c := range original {
		for ri, rc := range rotated {
			if c == rc {
				return i, ri, true
			}
		}
	}
	return
}
