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

	first, ok := findFirstOverlap(original, rotated)
	if !ok {
		return false
	}

	for x := 0; x < len(rotated)-first; x++ {
		if original[x] != rotated[x+first] {
			return false
		}
	}

	return strings.Contains(original, rotated[:first])
}

func findFirstOverlap(original, rotated string) (ri int, found bool) {
	c := original[0]
	for ri, rc := range rotated {
		if rune(c) == rc {
			return ri, true
		}
	}
	return
}
