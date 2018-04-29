package list

// Partition the list around the value v.
// All values smaller than v must appear before v,
// if v appears in the list, it does not need to appear
// before all the values that are bigger than it.
// Only that all values smaller than v, appear before v.
func (l *StrLinked) Partition(v string) {
	_, ok := l.head.next()
	curr := l.head
	head := l.head
	var prev Node

	for ok {
		if curr.value() < v {
			if prev != nil {
				var next Node
				next, ok = curr.next()

				updatePrevSkipCurr(prev, curr)
				head = pushOnHead(head, curr)

				curr = next
				continue
			}
		}

		prev = curr
		curr, ok = curr.next()
	}
	l.head = head
}

func pushOnHead(head, curr Node) Node {
	curr.setNext(head)
	return curr
}

func updatePrevSkipCurr(prev, curr Node) {
	n, _ := curr.next()
	prev.setNext(n)
}
