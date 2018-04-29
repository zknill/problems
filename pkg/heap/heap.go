package heap

// Binary is a binary heap it supports insertions according
// to the comparator passed. This allows for construction of
// a min or max binary heap.
type Binary struct {
	cmp  comparator
	heap []int
}

type comparator interface {
	compare(int, int) bool
}

type compareFn func(int, int) bool

func (f compareFn) compare(i, j int) bool {
	return f(i, j)
}

// NewMinHeap constructs a new min binary heap from the ints passed
func NewMinHeap(vals ...int) *Binary {
	min := &Binary{cmp: compareFn(func(i, j int) bool { return i > j })}

	if len(vals) == 0 {
		return min
	}

	if len(vals) == 1 {
		min.heap = append(min.heap, vals[0])
		return min
	}

	min.AddAll(vals[0:]...)

	return min
}

// NewMaxHeap constructs a new max binary heap from the ints passed
func NewMaxHeap(vals ...int) *Binary {
	max := &Binary{cmp: compareFn(func(i, j int) bool { return i < j })}

	if len(vals) == 0 {
		return max
	}

	if len(vals) == 1 {
		max.heap = append(max.heap, vals[0])
	}

	max.AddAll(vals[0:]...)

	return max
}

// AddAll adds all the ints to the heap, ensuring
// that the heap is balanced according to the comparator.
func (b *Binary) AddAll(vals ...int) {
	if len(vals) == 0 {
		return
	}

	for _, vv := range vals {
		b.heap = append(b.heap, vv)
		b.balance()
	}
}

func (b *Binary) balance() {
	i := len(b.heap) - 1
	for i != 0 && b.cmp.compare(b.heap[parent(i)], b.heap[i]) {
		b.heap[parent(i)], b.heap[i] = b.heap[i], b.heap[parent(i)]
		i = parent(i)
	}
}

func parent(idx int) int {
	return (idx - 1) / 2
}
