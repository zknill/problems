package list


// KthLastElem returns the kth element from the end of the list
// The last element is considered k = 0.
func (l *StrLinked) KthLastElem(k int) Node {
	_, ok := l.head.next()
	curr := l.head

	m := make(map[int]Node)

	idx := 0
	for ok {
		m[idx] = curr

		idx++
		curr, ok = curr.next()
	}

	return m[len(m) - 1 - k]
}
