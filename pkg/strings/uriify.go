package strings

// Replace all the spaces in a string with %20
func URLify(str string) string {
	runes := []rune(str)
	space := rune(' ')

	for idx := range runes {
		if runes[idx] == space {
			runes = append(runes[:idx], append([]rune{'%', '2', '0'}, runes[idx+1:]...)...)
		}
	}
	return string(runes)
}

// This one is faster because it uses a single array
// and then does in place replacements of runes.
func URLifyImproved(str string) string {
	// count the spaces
	var spaces int
	for _, c := range str {
		if c == ' ' {
			spaces++
		}
	}

	// make a new array with the extra space in it
	output := make([]rune, len(str)+3*spaces)
	outIdx := len(output) - 1

	// go backwards through the old array setting the last
	// rune of str to the last position of output. When
	// we get a space, put all 3 chars in there. We know
	// there will always be space for the extras.
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] == ' ' {
			output[outIdx-3] = '%'
			output[outIdx-2] = '2'
			output[outIdx-1] = '0'
			outIdx = outIdx - 3
		} else {
			output[outIdx-1] = rune(str[i])
			outIdx--
		}
	}

	return string(output)
}
