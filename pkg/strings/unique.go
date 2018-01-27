package strings

import (
	"sort"
)

// Check if all the chars in a string are unique
func Unique(str string) bool {
	chars := make(map[rune]bool)

	for _, c := range str {
		if chars[c] {
			return false
		}
		chars[c] = true
	}

	return true
}

func UniqueNoMap(str string) bool {
	runes := []rune(str)

	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})

	for idx := range runes {
		if idx == 0 {
			continue
		}

		if runes[idx-1] == runes[idx] {
			return false
		}
	}
	return true
}
