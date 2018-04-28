package list

import "reflect"

func (l *StrLinked) RemoveDupesNoBuf() {
	sn := &scanningNoder{head: l.head}
	l.head = removeDupesNoBuf(l.head, sn)
}

func removeDupesNoBuf(head Node, sn hasNoder) Node {
	n := head
	_, ok := head.next()
	prev := head

	for ok {

		if sn.hasNode(n) {
			prev.setNext(nextNode(n))
		}

		prev = n
		n, ok = n.next()
	}

	return head
}

type hasNoder interface {
	hasNode(n Node) bool
}

type scanningNoder struct {
	head Node
}

func (s scanningNoder) hasNode(n Node) bool {
	_, ok := s.head.next()
	curr := s.head

	for ok {
		if sameNode(curr, n) {
			return true
		}

		curr, ok = curr.next()
	}

	return false
}

func sameNode(n, needle Node) bool {
	if n == nil {
		return false
	}

	sameNode := reflect.DeepEqual(n, needle)
	matching := n.value() == needle.value()

	if !sameNode && matching {
		return true
	}

	return false
}
