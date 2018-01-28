package strings

// OneAway reports if the edit string is a single edit away
// from the original string. An edit is defined as:
//	1. Add a char
//	2. Remove a char
// 	3. Change a char
// This is has a O(2N + M) runtime; where
// N = len(originalStr) and M = len(editStr).
// You could do it in O(N) where N = len(shorter string).
func OneAway(originalStr, editStr string) bool {
	orig := make(map[rune]int)
	edit := make(map[rune]int)

	for _, c := range originalStr {
		orig[c]++
	}

	for _, c := range editStr {
		edit[c]++
	}

	for k := range orig {
		if edit[k] > 0 {
			remove(edit, k)
		}
		if orig[k] > 0 {
			remove(orig, k)
		}
	}

	return len(orig)+len(edit) <= 1
}

func remove(m map[rune]int, k rune) {
	if m[k] == 1 {
		delete(m, k)
		return
	}
	m[k]--
}
