package list

import (
	"fmt"
)

// StrLinked is an implementation of a
// linked list of strings.
type StrLinked struct {
	head Node
}

// NewStrLinked is a constructor for
// a new linked list of strings.
func NewStrLinked(strs ...string) *StrLinked {
	if len(strs) < 1 {
		return &StrLinked{}
	}

	var next Node
	for i := len(strs); i > 0; i-- {
		n := &strNode{
			n: next,
			v: strs[i-1],
		}
		next = n
	}

	return &StrLinked{head: next}
}

// RemoveDupes removes the duplicate strings
// from the string linked list. It uses a temp
// map to deduplicate the strings in the list.
// Memory is O(n) for recursive function stack
// frames and O(n) for the unique buffer.
// The computational complexity is O(n).
func (l *StrLinked) RemoveDupes() {
	tmp := make(tmp)

	l.head = remove(tmp.contains, l.head)
}

func remove(isDupe dupeFunc, n Node) Node {
	if n == nil {
		return nil
	}

	nn := nextNode(n)

	if isDupe(n) {
		return remove(isDupe, nn)
	}

	n.setNext(remove(isDupe, nn))
	return n
}

func nextNode(n Node) Node {
	nn, ok := n.next()
	if !ok {
		return nil
	}
	return nn
}

type tmp map[string]struct{}

func (t tmp) contains(n Node) bool {
	if _, exists := t[n.value()]; exists {
		return true
	}

	t[n.value()] = struct{}{}
	return false
}


type dupeFunc func (node Node) bool

// String returns a prettified version
// of the string linked list.
func (l *StrLinked) String() string {
	return fmt.Sprintf("String Linked: %v", l.head)
}

type Node interface {
	setNext(Node)
	next() (Node, bool)
	value() string

	fmt.Stringer
}

type strNode struct {
	n Node
	v string
}

func (s *strNode) String() string {
	return fmt.Sprintf("{value: %q, next: %v}", s.v, s.n)
}

func (s *strNode) setNext(n Node) {
	s.n = n
}

func (s *strNode) next() (Node, bool) {
	return s.n, s.n != nil
}

func (s *strNode) value() string {
	return s.v
}