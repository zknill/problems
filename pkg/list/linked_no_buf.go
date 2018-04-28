package list

import "reflect"

func (l *StrLinked) RemoveDupesNoBuf() {
	l.head = removeDupesNoBuf(l.head, containsStr(l.head))
}

func removeDupesNoBuf(head node, c dupeFunc) node {
	n := head
	_, ok := head.next()
	prev := head

	for ok {

		if c(n) {
			prev.setNext(nextNode(n))
		}

		prev = n
		n, ok = n.next()
	}

	return head
}

// TODO (zak): Refactor this into something that doesn't look like spaghetti!
func containsStr(head node) dupeFunc {
	return func(needle node) bool {
		var f func(n node, num uint) bool

		f = func(n node, num uint) bool {
			if n == nil {
				return false
			}

			sameNode := reflect.DeepEqual(n, needle)
			matching := n.value() == needle.value()

			if !sameNode && matching {
				return true
			}

			n, ok := n.next()
			if !ok {
				return false
			}

			return f(n, num)
		}

		return f(head, 0)
	}
}
